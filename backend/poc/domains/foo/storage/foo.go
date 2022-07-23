package storage

type FooStorage struct {
	saver
	updater
	deleter
}

func (fs FooStorage) FindByID(id uint32) entities.Foo {
	return entities.Foo{}
}
