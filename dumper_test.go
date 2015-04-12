package main

import (
	"github.com/dsociative/evego/parser"
	"github.com/stretchr/testify/suite"

	"gopkg.in/mgo.v2"

	"testing"
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

	s.dumper = New(s.db)
	s.characters = s.dumper.characters
	s.queue = s.dumper.queue
	s.db.DropDatabase()

}

func (s *BaseTests) GetAllCharacters() []parser.Character {
	var characters []parser.Character
	iter := s.characters.Find(nil).Iter()
	iter.All(&characters)
	return characters
}

func (s *BaseTests) GetAllQueue() []parser.SkillQueue {
	var model []parser.SkillQueue
	iter := s.queue.Find(nil).Iter()
	iter.All(&model)
	return model
}

func (s *DumperTests) TestCharacters() {
	data := []parser.Character{
		parser.Character{Name: "DISSNET", CharacterID: "123"},
		parser.Character{Name: "DISSTORG", CharacterID: "345"},
	}
	s.Equal(nil, s.dumper.Characters(data...))
	s.Equal(data, s.GetAllCharacters())
}

func (s *DumperTests) TestCharacterUpdate() {
	data := []parser.Character{
		parser.Character{Name: "DISSNET", CharacterID: "123"},
		parser.Character{Name: "DISSTORG", CharacterID: "142"},
	}
	s.dumper.Characters(data...)
	s.dumper.Characters(parser.Character{Name: "NewNick", CharacterID: "142"})

	s.Equal([]parser.Character{
		parser.Character{Name: "DISSNET", CharacterID: "123"},
		parser.Character{Name: "NewNick", CharacterID: "142"},
	}, s.GetAllCharacters())
}

func (s *DumperTests) TestSkillQueue() {
	data := parser.SkillQueue{
		CharacterID: "123",
		Skill: []parser.Skill{
			parser.Skill{TypeID: 32},
		},
	}
	dumpAndCheck := func() {
		s.dumper.Queue(data)
		s.Equal([]parser.SkillQueue{data}, s.GetAllQueue())
	}

	dumpAndCheck()

	data.Skill = append(data.Skill, parser.Skill{TypeID: 99})
	dumpAndCheck()

	data.Skill = []parser.Skill{}
	dumpAndCheck()
}

func TestDumperTests(t *testing.T) {
	suite.Run(t, new(DumperTests))
}
