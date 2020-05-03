package models

type Word struct {
	Root   string `json:"root,omitempty" bson:"root,omitempty"`
	Values []Value
}

type Value struct {
	Flag        bool
	Replacement string
}
