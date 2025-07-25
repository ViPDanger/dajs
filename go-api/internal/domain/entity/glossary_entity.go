package entity

type Glossary struct {
	CreatorID string `json:"creator_id" bson:"creator_id"`
	ID        ID     `bson:"_id"`
	Text      string
}
