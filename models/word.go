package models

//NOTE: flag is lets the word pass through after one of us authorizes it good to go.  This prevents racist/bad shit from getting into the db.

type Word struct {
	//Flag   bool     `json:"flag,omitempty" bson:"flag,omitempty"`
	Flag   bool
	ID     string
	Root   string   `json:"root,omitempty" bson:"root,omitempty"`
	Values []string `json:"values,omitempty" bson:"values,omitempty"`
}
