package fetcher

import (
    "testing"
    "github.com/stretchr/testify/suite"
    "github.com/stretchr/testify/assert"

    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"

    "github.com/dsociative/evego/parser"
)

type FakeApi struct {}

func (api FakeApi) Characters() []parser.Character {
    characters := []parser.Character{parser.Character{Name: "DISSNET"}}
    return characters
}

type FetcherTests struct {
    suite.Suite
    session *mgo.Session
    characters *mgo.Collection
    db *mgo.Database
    fetcher Fetcher
}

func (s *FetcherTests) SetupTest() {
    session, err := mgo.Dial("localhost")
    if err != nil {
        panic(err)
    }
    s.session = session
    s.db = session.DB("test")

    s.fetcher = New(s.db)
}

func (s *FetcherTests) TestExample() {
    api := FakeApi{}
    s.fetcher.Dump(api)

    character := parser.Character{}
    s.fetcher.characters.Find(bson.M{"name": "DISSNET"}).One(&character)

    assert.Equal(s.T(), character, parser.Character{Name: "DISSNET"})
}

func TestFetcherTests(t *testing.T) {
    suite.Run(t, new(FetcherTests))
}