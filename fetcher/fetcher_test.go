package fetcher

import (
    "testing"
    "io"

    . "gopkg.in/check.v1"
    "gopkg.in/mgo.v2"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type FethcerSuite struct{
    session *mgo.Session
    db *mgo.Database
}

var _ = Suite(&FethcerSuite{})

func (s *FethcerSuite) SetUpTest(c *C) {
    session, err := mgo.Dial("localhost")
    if err != nil {
        panic(err)
    }
    s.session = session
    s.db = session.DB("test")
}

func (s *FethcerSuite) TestHelloWorld(c *C) {
    c.Assert(42, Equals, "42")
    c.Assert(io.ErrClosedPipe, ErrorMatches, "io: .*on closed pipe")
    c.Check(42, Equals, 42)
}