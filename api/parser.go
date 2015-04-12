package api

import "encoding/xml"
import "fmt"

func Parse(raw []byte, model Model) {
	err := xml.Unmarshal(raw, &model)
	if err != nil {
		fmt.Printf("error: %v", err)
	}
}

func ParseCharacters(raw []byte) []Character {
	characters := Characters{}
	err := xml.Unmarshal(raw, &characters)
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	return characters.Character
}

func ParseSkillQueue(raw []byte) SkillQueue {
	queue := SkillQueue{}
	Parse(raw, &queue)
	return queue
}

func ParseSkillTree(raw []byte) Model {
	tree := Tree{}
	Parse(raw, &tree)
	return Tree(tree)
}
