package main

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"os"
	"strconv"
	"strings"
)

const NO_POKER_HAND = "no poker hand"
const PAIR_ONE = "1 pair"
const PAIR_TWO = "2 pair"
const THREE_OF_A_KIND = "3 of a kind"
const FOUR_OF_A_KIND = "4 of a kind"
const FIVE_OF_A_KIND = "5 of a kind"
const SIX_OF_A_KIND = "6 of a kind"
const STRAIGHT_5 = "straight 5 in a row"
const STRAIGHT_6 = "straight 6 in a row"
const FULL_HOUSE_3_2 = "3-2 full house"
const FULL_HOUSE_3_3 = "3-3 full house"
const FULL_HOUSE_4_2 = "4-2 full house"
const FULL_HOUSE_4_3 = "4-3 full house"
const FULL_HOUSE_5_2 = "5-2 full house"

const TWINS = "twins"
const TRIPLETS = "triplets"

type Monkey struct {
	id          int
	hat         int
	fur         int
	clothes     int
	eyes        int
	earring     int
	mouth       int
	background  int
	trait_count int
	color_match string
	mouth_match string
	zeros       int
	nips        string
	xplets      []int
	poker_hands []string
}

func GetOnChainMonkeys() []Monkey {
	err := DownloadFile("/tmp/meta.csv", "https://raw.githubusercontent.com/metagood/OnChainMonkeyData/main/OCM_meta_traits.csv")
	if err != nil {
		log.Fatal(err)
	}
	err = DownloadFile("/tmp/twins.json", "https://raw.githubusercontent.com/metagood/OnChainMonkeyData/main/meta_traits.json")
	if err != nil {
		log.Fatal(err)
	}

	monkeys, err := buildMonkeysFromCsv("/tmp/ocm.csv")
	if err != nil {
		log.Fatal(err)
	}

	updateMonkeysWithTwinMetaFromJson("/tmp/twins.json", monkeys)
	return monkeys
}

func buildMonkeysFromCsv(file string) ([]Monkey, error) {
	var monkeys []Monkey
	csv_file, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	r := csv.NewReader(strings.NewReader(string(csv_file)))
	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	for i, record := range records {
		if i == 0 { //ignore header
			continue
		}
		monkey := Monkey{
			id:          GetFirstIntReturn(strconv.Atoi(record[0])),
			hat:         GetFirstIntReturn(strconv.Atoi(record[1])),
			fur:         GetFirstIntReturn(strconv.Atoi(record[2])),
			clothes:     GetFirstIntReturn(strconv.Atoi(record[3])),
			eyes:        GetFirstIntReturn(strconv.Atoi(record[4])),
			earring:     GetFirstIntReturn(strconv.Atoi(record[5])),
			mouth:       GetFirstIntReturn(strconv.Atoi(record[6])),
			background:  GetFirstIntReturn(strconv.Atoi(record[7])),
			trait_count: GetFirstIntReturn(strconv.Atoi(record[8])),
			color_match: record[9],
			mouth_match: record[10],
			zeros:       GetFirstIntReturn(strconv.Atoi(record[11])),
			nips:        record[12],
			xplets:      []int{},
		}
		monkeys = append(monkeys, monkey)
	}
	return monkeys, nil
}

func updateMonkeysWithTwinMetaFromJson(file string, monkeys []Monkey) {
	j, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	var meta_traits map[string]interface{}
	json.Unmarshal([]byte(j), &meta_traits)

	addMetaTraitToMonkeys(monkeys, meta_traits, NO_POKER_HAND)
	addMetaTraitToMonkeys(monkeys, meta_traits, PAIR_ONE)
	addMetaTraitToMonkeys(monkeys, meta_traits, PAIR_TWO)
	addMetaTraitToMonkeys(monkeys, meta_traits, THREE_OF_A_KIND)
	addMetaTraitToMonkeys(monkeys, meta_traits, FOUR_OF_A_KIND)
	addMetaTraitToMonkeys(monkeys, meta_traits, FIVE_OF_A_KIND)
	addMetaTraitToMonkeys(monkeys, meta_traits, SIX_OF_A_KIND)
	addMetaTraitToMonkeys(monkeys, meta_traits, STRAIGHT_5)
	addMetaTraitToMonkeys(monkeys, meta_traits, STRAIGHT_6)
	addMetaTraitToMonkeys(monkeys, meta_traits, FULL_HOUSE_3_2)
	addMetaTraitToMonkeys(monkeys, meta_traits, FULL_HOUSE_3_3)
	addMetaTraitToMonkeys(monkeys, meta_traits, FULL_HOUSE_4_2)
	addMetaTraitToMonkeys(monkeys, meta_traits, FULL_HOUSE_4_3)
	addMetaTraitToMonkeys(monkeys, meta_traits, FULL_HOUSE_5_2)

	addTwinsFromXpletsJson(monkeys, meta_traits, TWINS)
	addTwinsFromXpletsJson(monkeys, meta_traits, TRIPLETS)

}

func addMetaTraitToMonkeys(monkeys []Monkey, meta_traits map[string]interface{}, trait string) {
	meta_trait, ok := meta_traits[trait].([]interface{})
	if !ok {
		log.Fatalln("Unknown meta trait property")
	}
	for _, v := range meta_trait {
		id := int(v.(float64))
		monkeys[id-1].poker_hands = append(monkeys[id-1].poker_hands, trait)
	}
}

func addTwinsFromXpletsJson(monkeys []Monkey, meta_traits map[string]interface{}, trait string) {
	meta_trait, ok := meta_traits[trait].([]interface{})
	if !ok {
		log.Fatalln("Unknown meta trait property")
	}
	for _, id := range meta_trait {
		obj, ok := id.(map[string]interface{})
		if !ok {
			log.Fatalln("Invalid format: Not a twin object")
		}
		for _, ids := range obj {
			id_array := ids.([]interface{})
			xplets := []int{}
			for _, twin := range id_array {
				xplets = append(xplets, int(twin.(float64)))
			}
			for _, twin := range xplets {
				filtered := Filter(xplets, func(i int) bool {
					return i != twin
				})
				monkeys[twin-1].xplets = filtered
			}
		}
	}
}