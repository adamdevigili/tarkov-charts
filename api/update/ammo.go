package update

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	. "github.com/adamdevigili/tarkov-charts-models"
	"github.com/kelseyhightower/envconfig"
	"github.com/mitchellh/mapstructure"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)
func UpdateAmmo(w http.ResponseWriter, r *http.Request) {
	var config Config
	err := envconfig.Process("", &config)
	if err != nil {
		log.Fatal(err.Error())
	}

	if config.VERCEL_ENV == "development" {
		log.Printf("%+v\n", config)
	}

	if config.VERCEL_ENV != "development" {
		if r.Header.Get(APIKeyHeader) != config.TC_API_KEY {
			log.Printf("incoming request API key invalid: %s", r.Header.Get(APIKeyHeader))
			fmt.Fprint(w, "not authorized")
	
			return
		}
	}

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// This is the map we will build of all ammo and relevant information throughout this function
	// We will eventually write this to our data store
	parsedAmmo := map[string]*Ammo{}
	ammoByCaliber := map[string]map[string]*Ammo{}

	client := &http.Client{Timeout: time.Second * 10}

	// Build GraphQL query to fetch only ammo items
	jsonValue, _ := json.Marshal(map[string]string{
		"query": `
			{
				itemsByType(type: ammo) {
					id
					name
					shortName
					iconLink
				}
			}
        `,
	})

	// Fetch all ammo IDs, as well as names, short names, and icon links
	request, _ := http.NewRequest("POST", "https://tarkov-tools.com/graphql", bytes.NewBuffer(jsonValue))
	response, err := client.Do(request)
	if err != nil {
		log.Printf("GraphQL request failed: %s\n", err)
	} else {
		log.Println("succesfully fetched ammo IDs")
	}
	defer response.Body.Close()

	data, _ := ioutil.ReadAll(response.Body)
	graphQLResp := &GraphQLResponse{}
	json.Unmarshal(data, graphQLResp)

	// Fetch current pen/damage data from BSG API
	request, _ = http.NewRequest(http.MethodGet, "https://tarkov-market.com/api/v1/bsg/items/all", nil)
	request.Header.Set("x-api-key", config.TM_API_KEY)
	response, err = client.Do(request)
	if err != nil || response.StatusCode != http.StatusOK {
		log.Fatalf("failed to fetch pen/damage data. Code: %d", response.StatusCode)
	} else {
		log.Println("succesfully fetched pen/damage data")
	}
	data, _ = ioutil.ReadAll(response.Body)

	var f interface{}
	err = json.Unmarshal(data, &f)
	if err != nil {
		log.Fatal("error parsing JSON: ", err)
	}

	// Need to do some Go magic to consume the BSG API properly. Also leverage the mapstructure package
	itemsMap := f.(map[string]interface{})
	var result BSGItem
	for _, item := range graphQLResp.Data.ItemsByType {
		err = mapstructure.Decode(itemsMap[item.ID], &result)
		if err != nil {
			log.Print("mapstructure error: ", err, item)
		}

		// When querying for ammo types, we currently get grenades and ammo boxes. Ignore them
		if !strings.Contains(result.Props.Name, "grenade") && 
			!strings.Contains(result.Props.Name, "pack") &&
			!strings.Contains(result.Props.Caliber, "Caliber40x46") {
			// Initialize the final map with BSG data
			parsedAmmo[item.ID] = &Ammo{
				Caliber:     result.Props.Caliber,
				Name:        result.Props.ShortName,
				Damage:      result.Props.Damage,
				Penetration: result.Props.PenetrationPower,
			}

			if ammoByCaliber[result.Props.Caliber] == nil {
				ammoByCaliber[result.Props.Caliber] = make(map[string]*Ammo)
			}
			ammoByCaliber[result.Props.Caliber][item.ID] = &Ammo{
				Caliber:     result.Props.Caliber,
				Name:        result.Props.ShortName,
				Damage:      result.Props.Damage,
				Penetration: result.Props.PenetrationPower,
			}
		}
	}

	// Fetch current prices of all items and parse.
	// Other option would be to fetch all 100+ ammo types individually, no thanks.
	// Also no option to fetch only ammo items via this API :(

	// NOTE: All the API requests can be done in parallel, however since this is intended to be run
	// periodically (so performance isn't that important), and as a lambda function (where memory
	// usage is important), we run these in sequence
	request, _ = http.NewRequest(http.MethodGet, "https://tarkov-market.com/api/v1/items/all", nil)
	request.Header.Set("x-api-key", config.TM_API_KEY)
	response, err = client.Do(request)
	if err != nil || response.StatusCode != http.StatusOK {
		log.Fatalf("failed to fetch ammo prices. Code: %d", response.StatusCode)
	} else {
		log.Println("succesfully fetched ammo prices")
	}
	data, _ = ioutil.ReadAll(response.Body)

	var fleaMarketData TarkovMarketItems
	err = json.Unmarshal(data, &fleaMarketData)
	if err != nil {
		log.Fatal("error parsing JSON: ", err)
	}

	// Since we have all items in Tarkov returned here, and no O(1) access by ID,
	// iterate over all entries, and update the relevant fields in our target map with
	// the average 24hr price
	for _, item := range fleaMarketData {
		if _, found := parsedAmmo[item.BsgID]; found {
			parsedAmmo[item.BsgID].Price = item.Avg24HPrice

			for _, ammoMap := range(ammoByCaliber) {
				if _, found := ammoMap[item.BsgID]; found {
					ammoMap[item.BsgID].Price = item.Avg24HPrice
				}
			}
		}
	}

	/*
	JSONBin integration deprecated. Leaving for later reference or...something 
	*/

	// Post the resulting JSON to jsonbin.io. We will probably want to store this in a more
	// mature data store (DynamoDB) in the future, but for now this is a good tool for rapid
	// development


	// binID := config.JSONBIN_BIN_ID
	// binAPIKEY := config.JSONBIN_API_KEY
	// binURL := fmt.Sprintf("https://api.jsonbin.io/v3/b/%s", binID)

	// req, _ := http.NewRequest(http.MethodPut, binURL, bytes.NewBuffer(parsedJSON))
	// req.Header.Set("X-Master-Key", binAPIKEY)
	// req.Header.Set("Content-Type", "application/json")

	// resp, err := client.Do(req)
	// if err != nil || response.StatusCode != http.StatusOK {
	// 	log.Fatalf("failed to write to the data store. Code: %d", response.StatusCode)
	// } else {
	// 	log.Println("succesfully wrote to the data store")
	// }
	// defer resp.Body.Close()



	clientOptions := options.Client().ApplyURI(fmt.Sprintf(
		"mongodb+srv://%s:%s@%s/%s?retryWrites=true&w=majority",
		config.MONGO_USER,
		config.MONGO_PASSWORD,
		config.MONGO_CLUSTER_PATH,
		config.MONGO_DB_NAME,
	))
	
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second); defer cancel()

	mongoClient, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	defer mongoClient.Disconnect(ctx)

	// ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second); defer cancel()

	if err = mongoClient.Ping(ctx, readpref.Primary()); err != nil {
        log.Fatal(err)
    }

	log.Print("successfully connected to database")


	items := mongoClient.Database(config.MONGO_DB_NAME).Collection("ammo")

	
	// parsedBSON, err := bson.Marshal(ammoByCaliber)
	// if err != nil {
	// 	log.Fatal("error marshalling BSON", err)
	// }

	log.Print("attempting to write updated data to database")

	res, err := items.ReplaceOne(
		ctx, 
		bson.M{"_name": "ammo_data"},
		bson.D{
			{"_name", "ammo_data"},
			{"_updated_at", time.Now().Format(time.RFC822)}, 
			{"data", ammoByCaliber}})
	if err != nil {
		log.Fatal("error writing to database", err)
	}

	log.Printf("successfully updated ammo data, number modified: %d", res.ModifiedCount)
	
	fmt.Fprint(w, "success")
}
