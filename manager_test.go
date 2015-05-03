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
	charactes := []Character{
		Character{Name: "DISSNET", CharacterID: "1"},
	}

	kills := Kills{Kills: killsExample}

	api := FakeApi{
		CharactersData: charactes,
		KillsData:      kills,
	}
	s.manager.Process(api)

	s.Equal(charactes, s.DumperTests.GetAllCharacters())
	s.Equal(kills.Kills, s.DumperTests.GetAllKills())
}

func TestManagerTests(t *testing.T) {
	suite.Run(t, new(ManagerTests))
}
