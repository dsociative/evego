package api

import (
	"encoding/xml"
	"gopkg.in/mgo.v2/bson"
)

type Model interface {
	FormKey() bson.M
}

type EVEAPI struct {
	XMLName xml.Name `xml:"eveapi"`
	Time    string   `xml:"currentTime"`
}

type Character struct {
	Name            string `xml:"name,attr"`
	CharacterID     string `xml:"characterID,attr"`
	CorporationName string `xml:"corporationName,attr"`
}

func (c Character) FormKey() bson.M {
	return bson.M{"characterid": c.CharacterID}
}

type Characters struct {
	Model
	EVEAPI
	Character []Character `xml:"result>rowset>row"`
}

type Skill struct {
	TypeID int `xml:"typeID,attr"`
}

type SkillQueue struct {
	EVEAPI
	CharacterID string
	Skill       []Skill `xml:"result>rowset>row"`
}

func (c SkillQueue) FormKey() bson.M {
	return bson.M{"characterid": c.CharacterID}
}

type GroupSkill struct {
	Name string `xml:"typeName,attr"`
}

type Group struct {
	Name  string       `xml:"groupName,attr"`
	Skill []GroupSkill `xml:"rowset>row"`
}

type Tree struct {
	Model
	EVEAPI
	Group []Group `xml:"result>rowset>row"`
}
