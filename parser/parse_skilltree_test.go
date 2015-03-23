package parser

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

const SkillTreeRaw = `<?xml version='1.0' encoding='UTF-8'?>
<eveapi version="2">
  <currentTime>2010-12-21 14:33:30</currentTime>
  <result>
    <rowset columns="groupName,groupID" key="groupID" name="skillGroups">
      <row groupID="266" groupName="Corporation Management">
        <rowset columns="typeName,groupID,typeID,published" key="typeID" name="skills">
          <row groupID="266" published="1" typeID="11584" typeName="Anchoring">
            <description>Skill at Anchoring Deployables. Can not be trained on Trial Accounts.</description>
            <rank>3</rank>
            <rowset columns="typeID,skillLevel" key="typeID" name="requiredSkills"/>
            <requiredAttributes>
              <secondaryAttribute>charisma</secondaryAttribute>
              <primaryAttribute>memory</primaryAttribute>
            </requiredAttributes>
            <rowset columns="bonusType,bonusValue" key="bonusType" name="skillBonusCollection">
              <row bonusType="canNotBeTrainedOnTrial" bonusValue="1"/>
            </rowset>
          </row>
          <row groupID="266" published="0" typeID="3369" typeName="CFO Training">
            <description>Skill at managing corp finances. 5% discount on all fees at non-hostile NPC station if acting as CFO of a corp. </description>
            <rank>3</rank>
            <rowset columns="typeID,skillLevel" key="typeID" name="requiredSkills">
              <row skillLevel="2" typeID="3363"/>
              <row skillLevel="3" typeID="3444"/>
            </rowset>
            <requiredAttributes>
              <secondaryAttribute>charisma</secondaryAttribute>
              <primaryAttribute>memory</primaryAttribute>
            </requiredAttributes>
            <rowset columns="bonusType,bonusValue" key="bonusType" name="skillBonusCollection"/>
          </row>
        </rowset>
      </row>
    </rowset>
  </result>
  <cachedUntil>2007-12-23 21:51:40</cachedUntil>
</eveapi>`



func TestParseSkillTree(t *testing.T) {
    tree := ParseSkillTree([]byte(SkillTreeRaw)).(Tree)
    group := []Group{Group{Name:"Corporation Management", Skill:[]GroupSkill{GroupSkill{Name:"Anchoring"}, GroupSkill{Name:"CFO Training"}}}}
    assert.Equal(t, tree.Group, group)
}

func BenchmarkParser(b *testing.B) {
    ParseSkillTree([]byte(SkillTreeRaw))
}