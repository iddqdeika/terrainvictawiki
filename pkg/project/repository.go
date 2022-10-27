package project

import "terrainvictawiki/pkg/index"

func (o *Object) ID() string {
	return o.Metadata.DataName
}

type Storage interface {
	Put(object *Object)
	Get(key string) (*Object, bool)
	GetAll() []*Object
}

type LocalizationStorageFacade interface {
	DisplayName(name string) string
	Summary(name string) string
	Quote(name string) string
	Description(name string) string
}

func NewRepositoryFromFile(templateFileName string) (*Repository, error) {
	r := NewRepository()
	repository, err := fillRepoFromFile(templateFileName, r)
	if err != nil {
		return repository, err
	}
	return r, nil
}

func fillRepoFromFile(templateFileName string, r *Repository) (*Repository, error) {
	mdl, err := getMetadataFromFile(templateFileName)
	if err != nil {
		return nil, err
	}
	ol := mergeToObjects(mdl)
	for _, o := range ol {
		r.s.Put(o)
	}
	return nil, nil
}

func mergeToObjects(mdl []*Metadata) []*Object {
	var res []*Object
	for _, data := range mdl {
		o := NewObject(data)
		res = append(res, o)
	}
	return res
}

func NewRepository() *Repository {
	r := &Repository{
		s: index.NewStorage[string, *Object](),
	}
	return r
}

type Repository struct {
	s Storage
}

func (r *Repository) GetAll() []*Object {
	return r.s.GetAll()
}

func (r *Repository) Localize(l LocalizationStorageFacade) {
	for _, object := range r.GetAll() {
		object.Locale.fill(object.Metadata.DataName, l)
	}
}
