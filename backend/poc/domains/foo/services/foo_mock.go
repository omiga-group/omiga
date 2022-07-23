package services

type FooMockManager struct {
	CreateFn   func(entities.Foo) (entities.Foo, error)
	FindByIDFn func(id uint32) (entities.Foo, error)
}

func (fmm FooMockManager) Create(foo entities.Foo) (entities.Foo, error) {
	if fmm.CreateFn != nil {
		return fmm.CreateFn(foo)
	}

	return entities.Foo{}, nil
}

func (fmm FooMockManager) FindByID(id uint32) (entities.Foo, error) {
	if fmm.FindByIDFn != nil {
		return fmm.FindByIDFn(id)
	}
	return entities.Foo{}, nil
}
