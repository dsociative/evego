package main

import (
	"github.com/dsociative/evego/api"
	"gopkg.in/mgo.v2"
)

type FakeApi struct {
	CharactersData []api.Character
	KillsData      api.Kills
}

func (api FakeApi) Characters() ([]api.Character, error) {
	return api.CharactersData, nil
}

func (api FakeApi) KillLog(_ *api.Character) (api.Kills, error) {
	return api.KillsData, nil
}

func DialTestDB() (*mgo.Session, *mgo.Database) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	db := session.DB("test")
	return session, db
}
