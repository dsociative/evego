package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/dsociative/evego/api"
	"log"
	"os"
)

var (
	print_data = flag.Bool("print", true, "print api data")
	dump       = flag.Bool("dump", false, "dump api data to mongodb")
	killlog    = flag.Bool("killlog", false, "killlog print")
)

type Config struct {
	Code string
	ID   string
}

func ReadConfig() Config {
	config := Config{}
	file, error := os.Open("config.json")

	if error != nil {
		panic(error)
	}

	decoder := json.NewDecoder(file)
	error = decoder.Decode(&config)

	if error != nil {
		panic(error)
	}
	return config
}

func GetApi() api.API {
	config := ReadConfig()
	return api.New(config.Code, config.ID)
}

func tryPrint(err error, d ...interface{}) {
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println(d...)
	}

}

func Print() {
	api := GetApi()
	fmt.Println(api.SkillTree())
	characters, err := api.Characters()
	tryPrint(err, characters)
	for _, char := range characters {
		queue, err := api.Queue(&char)
		tryPrint(err, char.Name, queue)

		if *killlog {
			kills, err := api.KillLog(&char)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(kills)
			}
		}
	}
}

func Dump() {
	api := GetApi()

	_, db := DialTestDB()
	manager := ManagerNew(db)
	manager.Process(api)
}

func main() {
	flag.Parse()

	if *print_data {
		Print()
	}
	if *dump {
		Dump()
	}

}
