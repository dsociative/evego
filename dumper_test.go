package main

import (
	"github.com/dsociative/evego/api"
	"github.com/stretchr/testify/suite"

	"gopkg.in/mgo.v2"
	"testing"
)

var (
	killsExample = []api.Kill{
		api.Kill{
			KillID: "1",
			Victim: api.Victim{Name: "Victim1", CharacterID: "1", CorporationID: "1", CorporationName: "Corp1"},
		},
		api.Kill{
			KillID: "2",
			Victim: api.Victim{Name: "Victim2", CharacterID: "2", CorporationID: "2", CorporationName: "Corp2"},
		},
	}
)

type BaseTests struct {
	suite.Suite
	session    *mgo.Session
	characters *mgo.Collection
	queue      *mgo.Collection
	db         *mgo.Database
	dumper     Dumper
}

type DumperTests struct {
	BaseTests
}

func (s *BaseTests) SetupTest() {
	s.session, s.db = DialTestDB()

	s.dumper = DumperNew(s.db)
	s.characters = s.dumper.characters
	s.queue = s.dumper.queue
	s.db.DropDatabase()

}

func (s *BaseTests) GetAllCharacters() []api.Character {
	var characters []api.Character
	iter := s.characters.Find(nil).Iter()
	iter.All(&characters)
	return characters
}

func (s *BaseTests) GetAllQueue() []api.SkillQueue {
	var model []api.SkillQueue
	iter := s.queue.Find(nil).Iter()
	iter.All(&model)
	return model
}

func (s *DumperTests) TestCharacters() {
	data := []api.Character{
		api.Character{Name: "DISSNET", CharacterID: "123"},
		api.Character{Name: "DISSTORG", CharacterID: "345"},
	}
	s.Equal(nil, s.dumper.Characters(data...))
	s.Equal(data, s.GetAllCharacters())
}

func (s *DumperTests) TestCharacterUpdate() {
	data := []api.Character{
		api.Character{Name: "DISSNET", CharacterID: "123"},
		api.Character{Name: "DISSTORG", CharacterID: "142"},
	}
	s.dumper.Characters(data...)
	s.dumper.Characters(api.Character{Name: "NewNick", CharacterID: "142"})

	s.Equal([]api.Character{
		api.Character{Name: "DISSNET", CharacterID: "123"},
		api.Character{Name: "NewNick", CharacterID: "142"},
	}, s.GetAllCharacters())
}

func (s *DumperTests) TestSkillQueue() {
	data := api.SkillQueue{
		CharacterID: "123",
		Skill: []api.Skill{
			api.Skill{TypeID: 32},
		},
	}
	dumpAndCheck := func() {
		s.dumper.Queue(data)
		s.Equal([]api.SkillQueue{data}, s.GetAllQueue())
	}

	dumpAndCheck()

	data.Skill = append(data.Skill, api.Skill{TypeID: 99})
	dumpAndCheck()

	data.Skill = []api.Skill{}
	dumpAndCheck()
}

func (s *BaseTests) GetAllKills() (kills []api.Kill) {
	iter := s.dumper.kills.Find(nil).Iter()
	iter.All(&kills)
	return
}

func (s *DumperTests) TestKills() {
	s.dumper.Kills(api.Kills{Kills: killsExample})
	s.Equal(killsExample, s.GetAllKills())
}

func TestDumperTests(t *testing.T) {
	suite.Run(t, new(DumperTests))
}
