package api

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var CharactersRaw = []byte(`<?xml version='1.0' encoding='UTF-8'?>
<eveapi version="2">
  <currentTime>2015-03-19 09:01:39</currentTime>
  <result>
    <rowset name="characters" key="characterID" columns="name,characterID,corporationName,corporationID,allianceID,allianceName,factionID,factionName">
      <row name="Superjoint Ritual" characterID="95225333" corporationName="Hedion University" corporationID="1000165" allianceID="0" allianceName="" factionID="0" factionName="" />
      <row name="DISSNET" characterID="129943370" corporationName="Worst Player Ever." corporationID="98148709" allianceID="0" allianceName="" factionID="0" factionName="" />
      <row name="DISSTORG" characterID="402705959" corporationName="School of Applied Knowledge" corporationID="1000044" allianceID="0" allianceName="" factionID="0" factionName="" />
    </rowset>
  </result>
  <cachedUntil>2015-03-19 09:55:07</cachedUntil>
</eveapi>`)

func TestParseCharacters(t *testing.T) {
	characters := Characters{}
	err := Parse(CharactersRaw, &characters)
	assert.NoError(t, err)
	assert.Equal(t, characters.Character[1], Character{Name: "DISSNET", CharacterID: "129943370", CorporationName: "Worst Player Ever."})
}
