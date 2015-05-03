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
		if characters, err := api.Characters(); err == nil {
			m.dumper.Characters(characters...)
			for _, character := range characters {
				if kills, err := api.KillLog(&character); err == nil {
					m.dumper.Kills(kills)
				}
			}
		}
	}
}
