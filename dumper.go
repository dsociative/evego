package main

import (
	"github.com/dsociative/evego/api"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Dumper struct {
	db         *mgo.Database
	characters *mgo.Collection
	queue      *mgo.Collection
}

func (d *Dumper) Dump(collection *mgo.Collection, key bson.M, model api.Model) error {
	err := collection.Update(key, model)
	if err == mgo.ErrNotFound {
		return collection.Insert(model)
	}
	return err
}

func (d *Dumper) Characters(characters ...api.Character) error {
	for _, character := range characters {
		err := d.Dump(d.characters, character.FormKey(), character)
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *Dumper) Queue(queue api.SkillQueue) error {
	return d.Dump(d.queue, queue.FormKey(), queue)
}

func DumperNew(db *mgo.Database) Dumper {
	return Dumper{
		db:         db,
		characters: db.C("characters"),
		queue:      db.C("queue"),
	}
}
