package models

type Word struct {
	Root     string   `bson:"root,omitempty" json:"root,omitempty"`
	Values   []string `bson:"values,omitempty" json:"values,omitempty"`
	Flag     bool     `bson:"flag,omitempty" json:"flag,omitempty"`
	Likes    int      `bson:"likes,omitempty" json:"likes,omitempty"`
	Dislikes int      `bson:"dislikes,omitempty" json:"dislikes,omitempty"`
}
