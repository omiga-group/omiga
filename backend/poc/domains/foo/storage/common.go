package storage

import (
	"context"

	"gorm.io/gorm"
)

type withWritesKey int

var withWrites withWritesKey

func WithWrites(ctx context.Context) context.Context {
	if hasWrites(ctx) {
		return ctx
	}
	return context.WithValue(ctx, withWrites, true)
}

func hasWrites(ctx context.Context) bool {
	v, ok := ctx.Value(withWrites).(bool)
	return ok && v
}

type gormConnections struct {
	writerDB *gorm.DB
	readerDB *gorm.DB
}

func (g gormConnections) reader() *gorm.DB {
	return g.readerDB
}

func (g gormConnections) writer() *gorm.DB {
	return g.writerDB
}

func (g gormConnections) benchReader(ctx context.Context) *gorm.DB {
	if hasWrites(ctx) {
		return g.benchWriter(ctx)
	}
	return g.reader()
}

func (g gormConnections) benchWriter(ctx context.Context) *gorm.DB {
	return g.writer()
}

type saver struct {
	*gormConnections
}

func (s saver) Save(ctx context.Context, i interface{}) error {
	return s.benchWriter(ctx).Create(i).Error
}

type updater struct {
	*gormConnections
}

func (s updater) Update(ctx context.Context, i interface{}) error {
	return s.benchWriter(ctx).Save(i).Error
}

type deleter struct {
	*gormConnections
}

func (d deleter) Delete(ctx context.Context, i interface{}) error {
	if err := d.benchWriter(ctx).Delete(i).Error; err != nil {
		return err
	}
	return nil
}
