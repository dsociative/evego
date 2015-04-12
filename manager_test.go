package main

import (
	"github.com/stretchr/testify/suite"
	"testing"
	// "gopkg.in/mgo.v2"
	. "github.com/dsociative/evego/api"
)

type ManagerTests struct {
	DumperTests
	manager Manager
}

func (s *ManagerTests) SetupTest() {
	s.DumperTests.SetupTest()
	s.manager = ManagerNew(s.DumperTests.db)
}

func (s *ManagerTests) TestProcess() {
	data := []Character{
		Character{Name: "DISSNET", CharacterID: "1"},
	}
	api := FakeApi{CharactersData: data}
	s.manager.Process(api)

	s.Equal(data, s.DumperTests.GetAllCharacters())
}

func TestManagerTests(t *testing.T) {
	suite.Run(t, new(ManagerTests))
}
