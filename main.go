package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/dsociative/evego/api"
	"os"
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

func Print() {
	api := GetApi()
	fmt.Println(api.SkillTree())
	characters := api.Characters()
	fmt.Println(characters)
	for _, char := range characters {
		fmt.Println(char.Name, api.Queue(&char))
	}
}

func Dump() {
	api := GetApi()

	_, db := DialTestDB()
	manager := ManagerNew(db)
	manager.Process(api)
}

func main() {
	var print = flag.Bool("print", false, "print api data")
	var dump = flag.Bool("dump", false, "dump api data")
	flag.Parse()

	if *print {
		Print()
	}
	if *dump {
		Dump()
	}
}
