package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/mitchellh/mapstructure"
)

type Ammo struct {
	Caliber     string `json:"caliber"`
	Name        string `json:"name"`
	Damage      int    `json:"damage"`
	Penetration int    `json:"penetration"`
	Price       int    `json:"price"`
}

// ------- tarkov-tools graphQL models -------
type Item struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	ShortName string `json:"shortName"`
	IconLink  string `json:"iconLink"`
}

type Data struct {
	ItemsByType []Item `json:"itemsByType"`
}

type GraphQLResponse struct {
	Data Data `json:"data"`
}

// ------- tarkov-market models -------
type TarkovMarketItem struct {
	UID            string    `json:"uid"`
	Name           string    `json:"name"`
	ShortName      string    `json:"shortName"`
	Price          int       `json:"price"`
	BasePrice      int       `json:"basePrice"`
	Avg24HPrice    int       `json:"avg24hPrice"`
	Avg7DaysPrice  int       `json:"avg7daysPrice"`
	TraderName     string    `json:"traderName"`
	TraderPrice    int       `json:"traderPrice"`
	TraderPriceCur string    `json:"traderPriceCur"`
	Updated        time.Time `json:"updated"`
	Slots          int       `json:"slots"`
	Diff24H        float64   `json:"diff24h"`
	Diff7Days      float64   `json:"diff7days"`
	Icon           string    `json:"icon"`
	Link           string    `json:"link"`
	WikiLink       string    `json:"wikiLink"`
	Img            string    `json:"img"`
	ImgBig         string    `json:"imgBig"`
	BsgID          string    `json:"bsgId"`
	IsFunctional   bool      `json:"isFunctional"`
	Reference      string    `json:"reference"`
}

type TarkovMarketItems []TarkovMarketItem

// ------- BSG models -------
type Prefab struct {
	Path string `json:"path"`
	Rcid string `json:"rcid"`
}

type UsePrefab struct {
	Path string `json:"path"`
	Rcid string `json:"rcid"`
}

type Contusion struct {
	X int `json:"x"`
	Y int `json:"y"`
	Z int `json:"z"`
}

type ArmorDistanceDistanceDamage struct {
	X int `json:"x"`
	Y int `json:"y"`
	Z int `json:"z"`
}

type Blindness struct {
	X int `json:"x"`
	Y int `json:"y"`
	Z int `json:"z"`
}

type Props struct {
	Name                                   string                      `json:"Name"`
	ShortName                              string                      `json:"ShortName"`
	Description                            string                      `json:"Description"`
	Weight                                 float64                     `json:"Weight"`
	BackgroundColor                        string                      `json:"BackgroundColor"`
	Width                                  int                         `json:"Width"`
	Height                                 int                         `json:"Height"`
	StackMaxSize                           int                         `json:"StackMaxSize"`
	Rarity                                 string                      `json:"Rarity"`
	SpawnChance                            int                         `json:"SpawnChance"`
	CreditsPrice                           int                         `json:"CreditsPrice"`
	ItemSound                              string                      `json:"ItemSound"`
	Prefab                                 Prefab                      `json:"Prefab"`
	UsePrefab                              UsePrefab                   `json:"UsePrefab"`
	StackObjectsCount                      int                         `json:"StackObjectsCount"`
	NotShownInSlot                         bool                        `json:"NotShownInSlot"`
	ExaminedByDefault                      bool                        `json:"ExaminedByDefault"`
	ExamineTime                            int                         `json:"ExamineTime"`
	IsUndiscardable                        bool                        `json:"IsUndiscardable"`
	IsUnsaleable                           bool                        `json:"IsUnsaleable"`
	IsUnbuyable                            bool                        `json:"IsUnbuyable"`
	IsUngivable                            bool                        `json:"IsUngivable"`
	IsLockedafterEquip                     bool                        `json:"IsLockedafterEquip"`
	QuestItem                              bool                        `json:"QuestItem"`
	LootExperience                         int                         `json:"LootExperience"`
	ExamineExperience                      int                         `json:"ExamineExperience"`
	HideEntrails                           bool                        `json:"HideEntrails"`
	RepairCost                             int                         `json:"RepairCost"`
	RepairSpeed                            int                         `json:"RepairSpeed"`
	ExtraSizeLeft                          int                         `json:"ExtraSizeLeft"`
	ExtraSizeRight                         int                         `json:"ExtraSizeRight"`
	ExtraSizeUp                            int                         `json:"ExtraSizeUp"`
	ExtraSizeDown                          int                         `json:"ExtraSizeDown"`
	ExtraSizeForceAdd                      bool                        `json:"ExtraSizeForceAdd"`
	MergesWithChildren                     bool                        `json:"MergesWithChildren"`
	CanSellOnRagfair                       bool                        `json:"CanSellOnRagfair"`
	CanRequireOnRagfair                    bool                        `json:"CanRequireOnRagfair"`
	ConflictingItems                       []interface{}               `json:"ConflictingItems"`
	FixedPrice                             bool                        `json:"FixedPrice"`
	Unlootable                             bool                        `json:"Unlootable"`
	UnlootableFromSlot                     string                      `json:"UnlootableFromSlot"`
	UnlootableFromSide                     []interface{}               `json:"UnlootableFromSide"`
	ChangePriceCoef                        int                         `json:"ChangePriceCoef"`
	AllowSpawnOnLocations                  []interface{}               `json:"AllowSpawnOnLocations"`
	SendToClient                           bool                        `json:"SendToClient"`
	AnimationVariantsNumber                int                         `json:"AnimationVariantsNumber"`
	DiscardingBlock                        bool                        `json:"DiscardingBlock"`
	RagFairCommissionModifier              int                         `json:"RagFairCommissionModifier"`
	IsAlwaysAvailableForInsurance          bool                        `json:"IsAlwaysAvailableForInsurance"`
	StackMinRandom                         int                         `json:"StackMinRandom"`
	StackMaxRandom                         int                         `json:"StackMaxRandom"`
	AmmoType                               string                      `json:"ammoType"`
	Damage                                 int                         `json:"Damage"`
	AmmoAccr                               int                         `json:"ammoAccr"`
	AmmoRec                                int                         `json:"ammoRec"`
	AmmoDist                               int                         `json:"ammoDist"`
	BuckshotBullets                        int                         `json:"buckshotBullets"`
	PenetrationPower                       int                         `json:"PenetrationPower"`
	PenetrationPowerDiviation              int                         `json:"PenetrationPowerDiviation"`
	AmmoHear                               int                         `json:"ammoHear"`
	AmmoSfx                                string                      `json:"ammoSfx"`
	MisfireChance                          float64                     `json:"MisfireChance"`
	MinFragmentsCount                      int                         `json:"MinFragmentsCount"`
	MaxFragmentsCount                      int                         `json:"MaxFragmentsCount"`
	AmmoShiftChance                        int                         `json:"ammoShiftChance"`
	CasingName                             string                      `json:"casingName"`
	CasingEjectPower                       int                         `json:"casingEjectPower"`
	CasingMass                             float64                     `json:"casingMass"`
	CasingSounds                           string                      `json:"casingSounds"`
	ProjectileCount                        int                         `json:"ProjectileCount"`
	InitialSpeed                           int                         `json:"InitialSpeed"`
	PenetrationChance                      float64                     `json:"PenetrationChance"`
	RicochetChance                         float64                     `json:"RicochetChance"`
	FragmentationChance                    float64                     `json:"FragmentationChance"`
	BallisticCoeficient                    int                         `json:"BallisticCoeficient"`
	Deterioration                          int                         `json:"Deterioration"`
	SpeedRetardation                       float64                     `json:"SpeedRetardation"`
	Tracer                                 bool                        `json:"Tracer"`
	TracerColor                            string                      `json:"TracerColor"`
	TracerDistance                         int                         `json:"TracerDistance"`
	ArmorDamage                            int                         `json:"ArmorDamage"`
	Caliber                                string                      `json:"Caliber"`
	StaminaBurnPerDamage                   float64                     `json:"StaminaBurnPerDamage"`
	HeavyBleedingDelta                     int                         `json:"HeavyBleedingDelta"`
	LightBleedingDelta                     int                         `json:"LightBleedingDelta"`
	ShowBullet                             bool                        `json:"ShowBullet"`
	HasGrenaderComponent                   bool                        `json:"HasGrenaderComponent"`
	FuzeArmTimeSec                         int                         `json:"FuzeArmTimeSec"`
	ExplosionStrength                      int                         `json:"ExplosionStrength"`
	MinExplosionDistance                   int                         `json:"MinExplosionDistance"`
	MaxExplosionDistance                   int                         `json:"MaxExplosionDistance"`
	FragmentsCount                         int                         `json:"FragmentsCount"`
	FragmentType                           string                      `json:"FragmentType"`
	ShowHitEffectOnExplode                 bool                        `json:"ShowHitEffectOnExplode"`
	ExplosionType                          string                      `json:"ExplosionType"`
	AmmoLifeTimeSec                        int                         `json:"AmmoLifeTimeSec"`
	Contusion                              Contusion                   `json:"Contusion"`
	ArmorDistanceDistanceDamage            ArmorDistanceDistanceDamage `json:"ArmorDistanceDistanceDamage"`
	Blindness                              Blindness                   `json:"Blindness"`
	IsLightAndSoundShot                    bool                        `json:"IsLightAndSoundShot"`
	LightAndSoundShotAngle                 int                         `json:"LightAndSoundShotAngle"`
	LightAndSoundShotSelfContusionTime      int                         `json:"LightAndSoundShotSelfContusionTime"`
	LightAndSoundShotSelfContusionStrength int                         `json:"LightAndSoundShotSelfContusionStrength"`
}

type BSGItem struct {
	ID     string `mapstructure:"_id"`
	Name   string `mapstructure:"_name"`
	Parent string `mapstructure:"_parent"`
	Type   string `mapstructure:"_type"`
	Props  Props  `mapstructure:"_props"`
	Proto  string `mapstructure:"_proto"`
}

// Configuration to be filled by envconfig
type Config struct {
	JSONBIN_BIN_ID  string
	JSONBIN_API_KEY string

	TM_API_KEY string

	UPDATE_AMMO_API_KEY string

	VERCEL_ENV string
}

func Handler(w http.ResponseWriter, r *http.Request) {
	var config Config
	err := envconfig.Process("", &config)
	if err != nil {
		log.Fatal(err.Error())
	}

	if config.VERCEL_ENV == "development" {
		log.Printf("%+v\n", config)
	}

	if config.VERCEL_ENV != "development" {
		if r.Header.Get("X-Update-Ammo-API-Key") != config.UPDATE_AMMO_API_KEY {
			fmt.Fprint(w, "not authorized")
			return
		}
	}

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
			log.Fatal("mapstructure error: ", err)
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

	parsedJSON, err := json.Marshal(ammoByCaliber)
	if err != nil {
		log.Fatal("error marshalling JSON: ", err)
	}

	// Post the resulting JSON to jsonbin.io. We will probably want to store this in a more
	// mature data store (DynamoDB) in the future, but for now this is a good tool for rapid
	// development
	binID := config.JSONBIN_BIN_ID
	binAPIKEY := config.JSONBIN_API_KEY
	binURL := fmt.Sprintf("https://api.jsonbin.io/v3/b/%s", binID)

	req, _ := http.NewRequest(http.MethodPut, binURL, bytes.NewBuffer(parsedJSON))
	req.Header.Set("X-Master-Key", binAPIKEY)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil || response.StatusCode != http.StatusOK {
		log.Fatalf("failed to write to the data store. Code: %d", response.StatusCode)
	} else {
		log.Println("succesfully wrote to the data store")
	}
	defer resp.Body.Close()

	fmt.Fprint(w, "success")
}
