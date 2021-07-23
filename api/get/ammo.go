package get

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	. "github.com/adamdevigili/tarkov-charts-models"
	"github.com/kelseyhightower/envconfig"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)
func GetAmmo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

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
	// parsedAmmo := map[string]*Ammo{}
	// ammoByCaliber := map[string]map[string]*Ammo{}


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

	log.Print("attempting to read from database")

	var ammo bson.M
	err = items.FindOne(
		ctx,
		bson.M{"_name": "ammo_data"},
	).Decode(&ammo)
	if err != nil {
		log.Fatal("error fetching from database", err)
	}

	log.Printf("successfully fetched data")

	w.WriteHeader(http.StatusOK)

	// Cache response in CDN for 30 minutes
	w.Header().Set("Cache-Control", "s-maxage=1800")

	json.NewEncoder(w).Encode(ammo)
	// jsonString, _ := json.Marshal(ammo["data"])
	// fmt.Fprint(w, string(jsonString))
}
