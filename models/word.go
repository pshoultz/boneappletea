package models

type Word struct {
	Root     string   `json:"root,omitempty" bson:"root,omitempty"`
	Values   []string `json:"values,omitempty" bson:"values,omitempty"`
	Flag     bool     `json:"flag,omitempty" bson:"flag,omitempty"`
	Likes    int      `json:"likes,omitempty" bson:"likes,omitempty"`
	Dislikes int      `json:"dislikes,omitempty" bson:"dislikes,omitempty"`
}

/*NOTE: Values could be an array of structs?
this is something I have to talk to kristina about, how the application is going to interface with its users and how things can be flagged for obvious reasons like racism ect... Maybe even return ones that have a higher likelyhood of being popular to make the app better rather than raw rng?

each value would consist of its own value, likes, dislikes, flag?

type Word struct {
	Root     string   `json:"root,omitempty" bson:"root,omitempty"`
	Values   []string `json:"values,omitempty" bson:"values,omitempty"`
	Flag     bool     `json:"flag,omitempty" bson:"flag,omitempty"`
	Likes    int      `json:"likes,omitempty" bson:"likes,omitempty"`
	Dislikes int      `json:"dislikes,omitempty" bson:"dislikes,omitempty"`
}
===============================================================================

type Value struct{
	value: string,
	flag: bool,
	likes: 0,
	dislikes: 0
}

type Word struct {
	Root     string   `json:"root,omitempty" bson:"root,omitempty"`
	Values []Value{{
		{
			value1 struct
		},
		{
			value2 struct
		},
	}
	Flag     bool     `json:"flag,omitempty" bson:"flag,omitempty"`
	Likes    int      `json:"likes,omitempty" bson:"likes,omitempty"`
	Dislikes int      `json:"dislikes,omitempty" bson:"dislikes,omitempty"`
}
*/
