package main

import (
    "testing"
    "github.com/stretchr/testify/suite"

    "gopkg.in/mgo.v2"
    // "gopkg.in/mgo.v2/bson"

    "github.com/dsociative/evego/parser"
)


type FakeApi struct {
    CharactersData []parser.Character
}

func (api FakeApi) Characters() []parser.Character {
    return api.CharactersData
}

type DumperTests struct {
    suite.Suite
    session *mgo.Session
    characters *mgo.Collection
    db *mgo.Database
    dumper Dumper
}

func (s *DumperTests) SetupTest() {
    session, err := mgo.Dial("localhost")
    if err != nil {
        panic(err)
    }
    s.session = session
    s.db = session.DB("test")

    s.dumper = New(s.db)
}

func (s *DumperTests) TearDownTest() {
    s.db.DropDatabase()
}

func (s *DumperTests) GetAllCharacters() []parser.Character {
    var characters []parser.Character
    iter := s.dumper.characters.Find(nil).Iter()
    iter.All(&characters)
    return characters
}

func (s *DumperTests) GetAllQueue() []parser.SkillQueue {
    var model []parser.SkillQueue
    iter := s.dumper.queue.Find(nil).Iter()
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

    s.Equal(s.GetAllCharacters(), []parser.Character{
        parser.Character{Name: "DISSNET", CharacterID: "123"},
        parser.Character{Name: "NewNick", CharacterID: "142"},
    })
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