package models

import "gopkg.in/mgo.v2/bson"

type Post struct {
	Id              bson.ObjectId `json:"id" bson:"_id"`
	UserId          string        `json:"userid" bson:"_userid"`
	Caption         string        `json:"caption" bson:"caption"`
	ImageUrl        string        `json:"ImageUrl" bson:"ImageUrl"`
	PostedTimestamp string        `json:"PostedTimestamp" bson:"PostedTimestamp"`
}
