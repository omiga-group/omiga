package storage

type Saver interface {
	Save(interface{}) (interface{}, error)
}

type Updater interface {
	Update(interface{}) (interface{}, error)
}

type Deleter interface {
	Delete(interface{}) (interface{}, error)
}

type FooStorer interface {
	Saver
	Updater
	Deleter
	FindByID(id uint32)
}
