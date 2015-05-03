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
	kills      *mgo.Collection
}

func (d *Dumper) Dump(collection *mgo.Collection, key bson.M, model api.WithKey) error {
	err := collection.Update(key, model)
	if err == mgo.ErrNotFound {
		return collection.Insert(model)
	}
	return err
}

func (d *Dumper) MultiDump(collection *mgo.Collection, model api.WithItems) (err error) {
	for _, model := range model.Items() {
		if err = d.Dump(collection, model.FormKey(), model); err != nil {
			return err
		}
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

func (d *Dumper) Kills(kills api.Kills) error {
	return d.MultiDump(d.kills, kills)
}

func DumperNew(db *mgo.Database) Dumper {
	return Dumper{
		db:         db,
		characters: db.C("characters"),
		queue:      db.C("queue"),
		kills:      db.C("kills"),
	}
}
