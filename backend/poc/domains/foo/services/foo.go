package services

type FooManager struct {
	storer storage.FooStorer
}

func (fmm FooManager) Create(foo entities.Foo) (entities.Foo, error) {
	return storer.Save()
}

func (fmm FooManager) FindByID(id uint32) (entities.Foo, error) {
	return storer.FindByID()
}
