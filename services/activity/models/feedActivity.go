package models

type FeedActivity struct {
	*ActivityBase
	FedBy string `json:"fed_by"`
}

func NewFeedActivity(animalID string, fedBy string) *FeedActivity {
	return &FeedActivity{
		ActivityBase: newActivity(animalID, "feed"),
		FedBy:        fedBy,
	}
}
