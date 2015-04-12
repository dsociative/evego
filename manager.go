package main

import (
	"github.com/dsociative/evego/api"
	"gopkg.in/mgo.v2"
)

type Manager struct {
	dumper Dumper
}

func ManagerNew(db *mgo.Database) Manager {
	return Manager{dumper: DumperNew(db)}
}

func (m *Manager) Process(apies ...api.APIFace) {
	for _, api := range apies {
		m.dumper.Characters(api.Characters()...)
	}
}
