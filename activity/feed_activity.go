package activity

type FeedActivity struct {
	*activity
	FedBy string
}

func NewFeedActivity(animal ActivityPerformer, fedBy string) *FeedActivity {
	return &FeedActivity{
		activity: newActivity(animal),
		FedBy:    fedBy,
	}
}
