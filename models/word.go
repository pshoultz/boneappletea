package models

type Word struct {
	Root   string   `json:"root,omitempty" bson:"root,omitempty"`
	Values []string `json:"values,omitempty" bson:"values,omitempty"`
	Flag   bool     `json:"flag,omitempty" bson:"flag,omitempty"`
}

//NOTE: flag is lets the word pass through after one of us authorizes it good to go.  This prevents racist/bad shit from getting into the db.
