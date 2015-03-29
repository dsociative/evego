package fetcher

import (
    // "github.com/dsociative/evego/parser"
    "github.com/dsociative/evego/api"
    "gopkg.in/mgo.v2"
)

type Character struct {
    Nick string
}

type Fetcher struct {
    db *mgo.Database
    characters *mgo.Collection
}

func (fetcher *Fetcher) Dump(api api.APIFace) {
    characters := api.Characters()
    for _, character := range characters {
        fetcher.characters.Insert(&character)
    }
}

func New(db *mgo.Database) Fetcher {
    return Fetcher{db: db, characters: db.C("characters")}
}