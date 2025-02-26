package event

type IndexRefreshTokenEvent struct {
}

func (*IndexRefreshTokenEvent) EventName() string {
	return "IndexRefreshTokenEvent"
}
