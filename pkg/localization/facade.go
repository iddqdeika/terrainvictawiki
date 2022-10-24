package localization

func NewStorageFacade(s *Storage, root string) *StorageFacade{
	return &StorageFacade{
		root: root,
		s:    s,
	}
}

// StorageFacade provides easy shortcuts
type StorageFacade struct {
	root string
	s *Storage
}

func (f *StorageFacade) DisplayName(name string) string {
	return f.s.Get(f.root, "displayName", name)
}

func (f *StorageFacade) Summary(name string) string {
	return f.s.Get(f.root, "summary", name)
}

func (f *StorageFacade) Quote(name string) string {
	return f.s.Get(f.root, "quote", name)
}

func (f *StorageFacade) Description(name string) string {
	return f.s.Get(f.root, "description", name)
}
