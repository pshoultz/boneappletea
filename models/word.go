package models

//NOTE:original version
//NOTE: flag lets the word pass through after one of us authorizes it good to go.  This prevents racist/bad shit from getting into the db.
type Word struct {
	Root   string `json:"root,omitempty" bson:"root,omitempty"`
	Values []Value
}

type Value struct {
	Flag        bool
	Replacement string
}

//NOTE: new version with a nested struct
//type Word struct {
//	//Flag   bool     `json:"flag,omitempty" bson:"flag,omitempty"`
//	ID     string
//	Root   string `json:"root,omitempty" bson:"root,omitempty"`
//	Values []struct {
//		Value string `json:"value,omitempty" bson:"value,omitempty"`
//		Flag  bool
//	}
//}
