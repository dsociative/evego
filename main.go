package main

import (
    "github.com/dsociative/evego/api"
    "fmt"
    "os"
    "encoding/json"
)

type Config struct {
    Code string
    ID string
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

func main() {
    config := ReadConfig()
    api := api.New(config.Code, config.ID)
    fmt.Println(api.SkillTree())
    characters := api.Characters()
    fmt.Println(characters)
    for _, char := range characters.Character {
        fmt.Println(char.Name, api.Queue(&char))
    }
}
