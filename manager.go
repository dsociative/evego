package main

import (
    "github.com/dsociative/evego/api"
    "gopkg.in/mgo.v2"
)


type Manager struct {

}

func ManagerNew(db *mgo.Database) Manager {
    return Manager{}
}

func (m *Manager) Process(ap ...api.APIFace) {

}