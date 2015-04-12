package main

import (
	"github.com/dsociative/evego/parser"
	"gopkg.in/mgo.v2"
)

type FakeApi struct {
	CharactersData []parser.Character
}

func (api FakeApi) Characters() []parser.Character {
	return api.CharactersData
}

func DialTestDB() (*mgo.Session, *mgo.Database) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	db := session.DB("test")
	return session, db
}
