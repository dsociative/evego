package api

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	KillLogRaw = []byte(`<?xml version='1.0' encoding='UTF-8'?>
<eveapi version="2">
  <currentTime>2010-06-17 18:15:01</currentTime>
  <result>
    <rowset name="kills" key="killID" columns="killID,solarSystemID,killTime,moonID">
      <row killID="63" solarSystemID="30000848" killTime="2007-11-15 15:36:00" moonID="0">
        <victim characterID="1" characterName="Victim1" corporationID="1"
                corporationName="Corp1" allianceID="0"
                allianceName="" factionID="0" factionName=""
                damageTaken="6378" shipTypeID="12003" />
        <rowset name="attackers" columns="characterID,characterName,corporationID,corporationName,allianceID,allianceName,
                factionID,factionName,securityStatus,damageDone,finalBlow,weaponTypeID,shipTypeID">
          <row characterID="0" characterName="" corporationID="1000127" corporationName="Guristas"
               allianceID="0" allianceName="" factionID="0" factionName="" securityStatus="0" 
               damageDone="6313" finalBlow="1" weaponTypeID="0" shipTypeID="203" />
          <row characterID="0" characterName="" corporationID="150279367" corporationName="Starbase Anchoring Corp"
               allianceID="0" allianceName="" securityStatus="0" damageDone="65" finalBlow="0"
               weaponTypeID="0" shipTypeID="16632" />
        </rowset>
        <rowset name="items" columns="typeID,flag,qtyDropped,qtyDestroyed,singleton" />
      </row>
      <row killID="62" solarSystemID="30000848" killTime="2007-11-15 14:48:00" moonID="0">
        <victim characterID="2" characterName="Victim2" corporationID="2"
                corporationName="Corp2" allianceID="0"
                allianceName="" factionID="0" factionName=""
                damageTaken="455" shipTypeID="606" />
        <rowset name="attackers" columns="characterID,characterName,corporationID,corporationName,allianceID,allianceName,
                factionID,factionName,securityStatus,damageDone,finalBlow,weaponTypeID,shipTypeID">
          <row characterID="0" characterName="" corporationID="1000127" corporationName="Guristas"
               allianceID="0" allianceName="" factionID="0" factionName="" securityStatus="0" 
               damageDone="394" finalBlow="1" weaponTypeID="0" shipTypeID="23328" />
          <row characterID="150131146" characterName="Mark Player" corporationID="150147571"
               corporationName="Peanut Butter Jelly Time" allianceID="150148475"
               allianceName="Margaritaville" securityStatus="0.3" damageDone="0"
               finalBlow="0" weaponTypeID="25715" shipTypeID="24698" />
        </rowset>
        <rowset name="items" columns="typeID,flag,qtyDropped,qtyDestroyed,singleton">
          <row typeID="3520" flag="0" qtyDropped="3" qtyDestroyed="1" singleton="0" />
          <row typeID="12076" flag="0" qtyDropped="0" qtyDestroyed="1" singleton="0">
            <rowset name="items" columns="typeID,flag,qtyDropped,qtyDestroyed,singleton">
              <row typeID="12259" flag="0" qtyDropped="0" qtyDestroyed="1" singleton="0" />
              <row typeID="1236" flag="0" qtyDropped="2" qtyDestroyed="1" singleton="0" />
              <row typeID="2032" flag="0" qtyDropped="1" qtyDestroyed="1" singleton="0" />
            </rowset>
          </row>
          <row typeID="12814" flag="0" qtyDropped="1" qtyDestroyed="3" singleton="0" />
          <row typeID="2364" flag="0" qtyDropped="0" qtyDestroyed="3" singleton="0" />
          <row typeID="26070" flag="0" qtyDropped="0" qtyDestroyed="2" singleton="0" />
          <row typeID="2605" flag="0" qtyDropped="1" qtyDestroyed="0" singleton="0" />
        </rowset>
      </row>
    </rowset>
  </result>
  <cachedUntil>2010-06-17 19:15:01</cachedUntil>
</eveapi>`)

	KillLogErrorRaw = []byte(`<?xml version='1.0' encoding='UTF-8'?>
<eveapi version="2">
  <currentTime>2015-05-02 18:38:04</currentTime>
  <error code="119">Kill log exhausted (You can only fetch kills that are less than a month old): New kills will be accessible at: 2015-05-02 19:31:06. If you are not expecting this message it is possible that some other application is using this key!</error>
  <cachedUntil>2015-05-02 19:31:06</cachedUntil>
</eveapi>`)
)

func TestParseKillLogError(t *testing.T) {
	kills := &Kills{}
	err := Parse(KillLogErrorRaw, kills)
	assert.Equal(t, "119", kills.GetError().Code)
	assert.EqualError(t, err, "Api request error code:119")
}

func TestParseKillLog(t *testing.T) {
	expected := []Kill{
		Kill{
			KillID: "63",
			Victim: Victim{Name: "Victim1", CharacterID: "1", CorporationID: "1", CorporationName: "Corp1"},
		},
		Kill{
			KillID: "62",
			Victim: Victim{Name: "Victim2", CharacterID: "2", CorporationID: "2", CorporationName: "Corp2"},
		},
	}
	killsModel := Kills{}
	err := Parse(KillLogRaw, &killsModel)
	kills := killsModel.Kills
	assert.NoError(t, err)
	assert.Equal(t, expected, kills)
}

func TestItems(t *testing.T) {
	kill1 := Kill{
		KillID: "63",
		Victim: Victim{Name: "Victim1", CharacterID: "1", CorporationID: "1", CorporationName: "Corp1"},
	}
	kill2 := Kill{
		KillID: "62",
		Victim: Victim{Name: "Victim2", CharacterID: "2", CorporationID: "2", CorporationName: "Corp2"},
	}
	kills := Kills{Kills: []Kill{kill1, kill2}}
	assert.Equal(t, []WithKey{WithKey(kill1), WithKey(kill2)}, kills.Items())
}
