package activity

type FeedActivity struct {
	*Activity
	FedBy string
}

func NewFeedActivity(animal ActivityPerformer, fedBy string) *FeedActivity {
	return &FeedActivity{
		Activity: newActivity("feed", animal),
		FedBy:    fedBy,
	}
}
