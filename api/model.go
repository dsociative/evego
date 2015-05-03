package api

import (
	"encoding/xml"
	"gopkg.in/mgo.v2/bson"
)

type WithKey interface {
	FormKey() bson.M
}

type WithItems interface {
	Items() []WithKey
}

type Model interface {
	GetError() APIError
}

type ModelStore struct {
	Kills []WithKey
}

type APIError struct {
	Code string `xml:"code,attr"`
	Msg  string `xml:",chardata"`
}

type EVEAPI struct {
	XMLName xml.Name `xml:"eveapi"`
	Time    string   `xml:"currentTime"`
	Error   APIError `xml:"error"`
}

func (c EVEAPI) GetError() APIError {
	return c.Error
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
	EVEAPI
	Group []Group `xml:"result>rowset>row"`
}

type Kills struct {
	EVEAPI
	ModelStore
	Kills []Kill `xml:"result>rowset>row"`
}

type Kill struct {
	KillID string `xml:"killID,attr"`
	Victim Victim `xml:"victim"`
}

func (k Kills) Items() []WithKey {
	var items []WithKey = []WithKey{}
	for _, kill := range k.Kills {
		items = append(items, WithKey(kill))
	}
	return items
}

func (c Kill) FormKey() bson.M {
	return bson.M{"killid": c.KillID}
}

type Victim struct {
	Name            string `xml:"characterName,attr"`
	CharacterID     string `xml:"characterID,attr"`
	CorporationID   string `xml:"corporationID,attr"`
	CorporationName string `xml:"corporationName,attr"`
}
