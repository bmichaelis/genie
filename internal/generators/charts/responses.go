package charts

type Responses struct {
	Enable         bool
	ArtifactoryUrl string
	Author         string
	Email          string
}

func NewResponses() *Responses {
	return &Responses{}
}
