package rssclient

type FeedType int8

const (
	FeedTypeUnknown FeedType = iota
	FeedTypeRSS
	FeedTypeAtom
)
