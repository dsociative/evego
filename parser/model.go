package parser

import "encoding/xml"

type EVEAPI struct {
    XMLName xml.Name `xml:"eveapi"`
    Time string `xml:"currentTime"`
}

type Character struct {
    Name string `xml:"name,attr"`
    CharacterID string `xml:"characterID,attr"`
    CorporationName string `xml:"corporationName,attr"`
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
    Skill []Skill `xml:"result>rowset>row"`
}