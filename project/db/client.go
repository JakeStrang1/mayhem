package db

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type Client struct {
	session *mgo.Session
}

func NewClient(url string) (*Client, error) {
	session, err := mgo.Dial(url)
	if err != nil {
		return nil, err
	}
	return &Client{session: session}, nil
}

func (c *Client) Create(collection string, document interface{}, result interface{}) error {
	s := c.session.Copy()
	change := mgo.Change{Update: bson.M{"$setOnInsert": document}, Upsert: true, ReturnNew: true}
	_, err := s.DB("").C(collection).Find(document).Apply(change, result)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) EnsureUniqueIndex(collection string, key []string) error {
	s := c.session.Copy()
	err := s.DB("").C(collection).EnsureIndex(mgo.Index{
		Key:    key,
		Unique: true,
	})
	return err
}

func (c *Client) Close() {
	c.session.Close()
}
