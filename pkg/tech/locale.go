package tech

type Locale struct {
	DisplayName string
	Summary     string
	Quote       string
	Description string
}

func (l *Locale) fill(name string, facade LocalizationStorageFacade) {
	l.DisplayName = facade.DisplayName(name)
	l.Summary = facade.Summary(name)
	l.Quote = facade.Quote(name)
	l.Description = facade.Description(name)
}
