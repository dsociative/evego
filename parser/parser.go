package parser

import "encoding/xml"
import "fmt"

func ParseCharacters(raw []byte) Characters {
    characters := Characters{}
    err := xml.Unmarshal(raw, &characters)
    if err != nil {
        fmt.Printf("error: %v", err)
    }
    return characters
}

func ParseSkillQueue(raw []byte) SkillQueue {
    fmt.Println(string(raw))
    queue := SkillQueue{}
    err := xml.Unmarshal(raw, &queue)
    if err != nil {
        fmt.Printf("error: %v", err)
    }
    return queue
}