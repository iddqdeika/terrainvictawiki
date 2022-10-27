package project

func NewObject(data *Metadata) *Object {
	return &Object{
		Metadata: data,
		Locale:   &Locale{},
	}
}

type Object struct {
	Metadata *Metadata
	Locale   *Locale
}
