package models

type Word struct {
	Root     string   `json:"root,omitempty" bson:"root,omitempty"`
	Values   []string `json:"values,omitempty" bson:"values,omitempty"`
	Flag     bool     `json:"flag,omitempty" bson:"flag,omitempty"`
	Likes    int      `json:"likes,omitempty" bson:"likes,omitempty"`
	Dislikes int      `json:"dislikes,omitempty" bson:"dislikes,omitempty"`
}
