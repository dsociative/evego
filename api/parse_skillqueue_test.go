package api

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const SkillQueueRaw = `<?xml version='1.0' encoding='UTF-8'?>
<eveapi version="2">
  <currentTime>2009-03-18 13:19:43</currentTime>
  <result>
    <rowset name="skillqueue" key="queuePosition" columns="queuePosition,typeID,level,startSP,endSP,startTime,endTime">
      <row queuePosition="1" typeID="11441" level="3" startSP="7072" endSP="40000" startTime="2009-03-18 02:01:06" endTime="2009-03-18 15:19:21" />
      <row queuePosition="2" typeID="20533" level="4" startSP="112000" endSP="633542" startTime="2009-03-18 15:19:21" endTime="2009-03-30 03:16:14" />
    </rowset>
  </result>
  <cachedUntil>2009-03-18 13:34:43</cachedUntil>
</eveapi>`

func TestParseSkillQueue(t *testing.T) {
	queue := SkillQueue{}
	Parse([]byte(SkillQueueRaw), &queue)
	skills := []Skill{Skill{TypeID: 11441}, Skill{TypeID: 20533}}
	assert.Equal(t, queue.Skill, skills)
}
