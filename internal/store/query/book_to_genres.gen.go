// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"library/internal/store/model"
)

func newBookToGenre(db *gorm.DB, opts ...gen.DOOption) bookToGenre {
	_bookToGenre := bookToGenre{}

	_bookToGenre.bookToGenreDo.UseDB(db, opts...)
	_bookToGenre.bookToGenreDo.UseModel(&model.BookToGenre{})

	tableName := _bookToGenre.bookToGenreDo.TableName()
	_bookToGenre.ALL = field.NewAsterisk(tableName)
	_bookToGenre.BookID = field.NewInt64(tableName, "book_id")
	_bookToGenre.GenreID = field.NewInt16(tableName, "genre_id")

	_bookToGenre.fillFieldMap()

	return _bookToGenre
}

type bookToGenre struct {
	bookToGenreDo

	ALL     field.Asterisk
	BookID  field.Int64
	GenreID field.Int16

	fieldMap map[string]field.Expr
}

func (b bookToGenre) Table(newTableName string) *bookToGenre {
	b.bookToGenreDo.UseTable(newTableName)
	return b.updateTableName(newTableName)
}

func (b bookToGenre) As(alias string) *bookToGenre {
	b.bookToGenreDo.DO = *(b.bookToGenreDo.As(alias).(*gen.DO))
	return b.updateTableName(alias)
}

func (b *bookToGenre) updateTableName(table string) *bookToGenre {
	b.ALL = field.NewAsterisk(table)
	b.BookID = field.NewInt64(table, "book_id")
	b.GenreID = field.NewInt16(table, "genre_id")

	b.fillFieldMap()

	return b
}

func (b *bookToGenre) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := b.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (b *bookToGenre) fillFieldMap() {
	b.fieldMap = make(map[string]field.Expr, 2)
	b.fieldMap["book_id"] = b.BookID
	b.fieldMap["genre_id"] = b.GenreID
}

func (b bookToGenre) clone(db *gorm.DB) bookToGenre {
	b.bookToGenreDo.ReplaceConnPool(db.Statement.ConnPool)
	return b
}

func (b bookToGenre) replaceDB(db *gorm.DB) bookToGenre {
	b.bookToGenreDo.ReplaceDB(db)
	return b
}

type bookToGenreDo struct{ gen.DO }

type IBookToGenreDo interface {
	gen.SubQuery
	Debug() IBookToGenreDo
	WithContext(ctx context.Context) IBookToGenreDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IBookToGenreDo
	WriteDB() IBookToGenreDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IBookToGenreDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IBookToGenreDo
	Not(conds ...gen.Condition) IBookToGenreDo
	Or(conds ...gen.Condition) IBookToGenreDo
	Select(conds ...field.Expr) IBookToGenreDo
	Where(conds ...gen.Condition) IBookToGenreDo
	Order(conds ...field.Expr) IBookToGenreDo
	Distinct(cols ...field.Expr) IBookToGenreDo
	Omit(cols ...field.Expr) IBookToGenreDo
	Join(table schema.Tabler, on ...field.Expr) IBookToGenreDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IBookToGenreDo
	RightJoin(table schema.Tabler, on ...field.Expr) IBookToGenreDo
	Group(cols ...field.Expr) IBookToGenreDo
	Having(conds ...gen.Condition) IBookToGenreDo
	Limit(limit int) IBookToGenreDo
	Offset(offset int) IBookToGenreDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IBookToGenreDo
	Unscoped() IBookToGenreDo
	Create(values ...*model.BookToGenre) error
	CreateInBatches(values []*model.BookToGenre, batchSize int) error
	Save(values ...*model.BookToGenre) error
	First() (*model.BookToGenre, error)
	Take() (*model.BookToGenre, error)
	Last() (*model.BookToGenre, error)
	Find() ([]*model.BookToGenre, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.BookToGenre, err error)
	FindInBatches(result *[]*model.BookToGenre, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.BookToGenre) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IBookToGenreDo
	Assign(attrs ...field.AssignExpr) IBookToGenreDo
	Joins(fields ...field.RelationField) IBookToGenreDo
	Preload(fields ...field.RelationField) IBookToGenreDo
	FirstOrInit() (*model.BookToGenre, error)
	FirstOrCreate() (*model.BookToGenre, error)
	FindByPage(offset int, limit int) (result []*model.BookToGenre, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IBookToGenreDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (b bookToGenreDo) Debug() IBookToGenreDo {
	return b.withDO(b.DO.Debug())
}

func (b bookToGenreDo) WithContext(ctx context.Context) IBookToGenreDo {
	return b.withDO(b.DO.WithContext(ctx))
}

func (b bookToGenreDo) ReadDB() IBookToGenreDo {
	return b.Clauses(dbresolver.Read)
}

func (b bookToGenreDo) WriteDB() IBookToGenreDo {
	return b.Clauses(dbresolver.Write)
}

func (b bookToGenreDo) Session(config *gorm.Session) IBookToGenreDo {
	return b.withDO(b.DO.Session(config))
}

func (b bookToGenreDo) Clauses(conds ...clause.Expression) IBookToGenreDo {
	return b.withDO(b.DO.Clauses(conds...))
}

func (b bookToGenreDo) Returning(value interface{}, columns ...string) IBookToGenreDo {
	return b.withDO(b.DO.Returning(value, columns...))
}

func (b bookToGenreDo) Not(conds ...gen.Condition) IBookToGenreDo {
	return b.withDO(b.DO.Not(conds...))
}

func (b bookToGenreDo) Or(conds ...gen.Condition) IBookToGenreDo {
	return b.withDO(b.DO.Or(conds...))
}

func (b bookToGenreDo) Select(conds ...field.Expr) IBookToGenreDo {
	return b.withDO(b.DO.Select(conds...))
}

func (b bookToGenreDo) Where(conds ...gen.Condition) IBookToGenreDo {
	return b.withDO(b.DO.Where(conds...))
}

func (b bookToGenreDo) Order(conds ...field.Expr) IBookToGenreDo {
	return b.withDO(b.DO.Order(conds...))
}

func (b bookToGenreDo) Distinct(cols ...field.Expr) IBookToGenreDo {
	return b.withDO(b.DO.Distinct(cols...))
}

func (b bookToGenreDo) Omit(cols ...field.Expr) IBookToGenreDo {
	return b.withDO(b.DO.Omit(cols...))
}

func (b bookToGenreDo) Join(table schema.Tabler, on ...field.Expr) IBookToGenreDo {
	return b.withDO(b.DO.Join(table, on...))
}

func (b bookToGenreDo) LeftJoin(table schema.Tabler, on ...field.Expr) IBookToGenreDo {
	return b.withDO(b.DO.LeftJoin(table, on...))
}

func (b bookToGenreDo) RightJoin(table schema.Tabler, on ...field.Expr) IBookToGenreDo {
	return b.withDO(b.DO.RightJoin(table, on...))
}

func (b bookToGenreDo) Group(cols ...field.Expr) IBookToGenreDo {
	return b.withDO(b.DO.Group(cols...))
}

func (b bookToGenreDo) Having(conds ...gen.Condition) IBookToGenreDo {
	return b.withDO(b.DO.Having(conds...))
}

func (b bookToGenreDo) Limit(limit int) IBookToGenreDo {
	return b.withDO(b.DO.Limit(limit))
}

func (b bookToGenreDo) Offset(offset int) IBookToGenreDo {
	return b.withDO(b.DO.Offset(offset))
}

func (b bookToGenreDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IBookToGenreDo {
	return b.withDO(b.DO.Scopes(funcs...))
}

func (b bookToGenreDo) Unscoped() IBookToGenreDo {
	return b.withDO(b.DO.Unscoped())
}

func (b bookToGenreDo) Create(values ...*model.BookToGenre) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Create(values)
}

func (b bookToGenreDo) CreateInBatches(values []*model.BookToGenre, batchSize int) error {
	return b.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (b bookToGenreDo) Save(values ...*model.BookToGenre) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Save(values)
}

func (b bookToGenreDo) First() (*model.BookToGenre, error) {
	if result, err := b.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.BookToGenre), nil
	}
}

func (b bookToGenreDo) Take() (*model.BookToGenre, error) {
	if result, err := b.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.BookToGenre), nil
	}
}

func (b bookToGenreDo) Last() (*model.BookToGenre, error) {
	if result, err := b.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.BookToGenre), nil
	}
}

func (b bookToGenreDo) Find() ([]*model.BookToGenre, error) {
	result, err := b.DO.Find()
	return result.([]*model.BookToGenre), err
}

func (b bookToGenreDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.BookToGenre, err error) {
	buf := make([]*model.BookToGenre, 0, batchSize)
	err = b.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (b bookToGenreDo) FindInBatches(result *[]*model.BookToGenre, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return b.DO.FindInBatches(result, batchSize, fc)
}

func (b bookToGenreDo) Attrs(attrs ...field.AssignExpr) IBookToGenreDo {
	return b.withDO(b.DO.Attrs(attrs...))
}

func (b bookToGenreDo) Assign(attrs ...field.AssignExpr) IBookToGenreDo {
	return b.withDO(b.DO.Assign(attrs...))
}

func (b bookToGenreDo) Joins(fields ...field.RelationField) IBookToGenreDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Joins(_f))
	}
	return &b
}

func (b bookToGenreDo) Preload(fields ...field.RelationField) IBookToGenreDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Preload(_f))
	}
	return &b
}

func (b bookToGenreDo) FirstOrInit() (*model.BookToGenre, error) {
	if result, err := b.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.BookToGenre), nil
	}
}

func (b bookToGenreDo) FirstOrCreate() (*model.BookToGenre, error) {
	if result, err := b.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.BookToGenre), nil
	}
}

func (b bookToGenreDo) FindByPage(offset int, limit int) (result []*model.BookToGenre, count int64, err error) {
	result, err = b.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = b.Offset(-1).Limit(-1).Count()
	return
}

func (b bookToGenreDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = b.Count()
	if err != nil {
		return
	}

	err = b.Offset(offset).Limit(limit).Scan(result)
	return
}

func (b bookToGenreDo) Scan(result interface{}) (err error) {
	return b.DO.Scan(result)
}

func (b bookToGenreDo) Delete(models ...*model.BookToGenre) (result gen.ResultInfo, err error) {
	return b.DO.Delete(models)
}

func (b *bookToGenreDo) withDO(do gen.Dao) *bookToGenreDo {
	b.DO = *do.(*gen.DO)
	return b
}
