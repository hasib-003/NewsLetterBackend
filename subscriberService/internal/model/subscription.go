package model

type Subscription struct {
	UserId        int64 `json:"userId"`
	PublicationId int64 `json:"publicationId"`
}
