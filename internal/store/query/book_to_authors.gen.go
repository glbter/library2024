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

func newBookToAuthor(db *gorm.DB, opts ...gen.DOOption) bookToAuthor {
	_bookToAuthor := bookToAuthor{}

	_bookToAuthor.bookToAuthorDo.UseDB(db, opts...)
	_bookToAuthor.bookToAuthorDo.UseModel(&model.BookToAuthor{})

	tableName := _bookToAuthor.bookToAuthorDo.TableName()
	_bookToAuthor.ALL = field.NewAsterisk(tableName)
	_bookToAuthor.BookID = field.NewInt64(tableName, "book_id")
	_bookToAuthor.AuthorID = field.NewInt64(tableName, "author_id")

	_bookToAuthor.fillFieldMap()

	return _bookToAuthor
}

type bookToAuthor struct {
	bookToAuthorDo bookToAuthorDo

	ALL      field.Asterisk
	BookID   field.Int64
	AuthorID field.Int64

	fieldMap map[string]field.Expr
}

func (b bookToAuthor) Table(newTableName string) *bookToAuthor {
	b.bookToAuthorDo.UseTable(newTableName)
	return b.updateTableName(newTableName)
}

func (b bookToAuthor) As(alias string) *bookToAuthor {
	b.bookToAuthorDo.DO = *(b.bookToAuthorDo.As(alias).(*gen.DO))
	return b.updateTableName(alias)
}

func (b *bookToAuthor) updateTableName(table string) *bookToAuthor {
	b.ALL = field.NewAsterisk(table)
	b.BookID = field.NewInt64(table, "book_id")
	b.AuthorID = field.NewInt64(table, "author_id")

	b.fillFieldMap()

	return b
}

func (b *bookToAuthor) WithContext(ctx context.Context) IBookToAuthorDo {
	return b.bookToAuthorDo.WithContext(ctx)
}

func (b bookToAuthor) TableName() string { return b.bookToAuthorDo.TableName() }

func (b bookToAuthor) Alias() string { return b.bookToAuthorDo.Alias() }

func (b bookToAuthor) Columns(cols ...field.Expr) gen.Columns {
	return b.bookToAuthorDo.Columns(cols...)
}

func (b *bookToAuthor) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := b.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (b *bookToAuthor) fillFieldMap() {
	b.fieldMap = make(map[string]field.Expr, 2)
	b.fieldMap["book_id"] = b.BookID
	b.fieldMap["author_id"] = b.AuthorID
}

func (b bookToAuthor) clone(db *gorm.DB) bookToAuthor {
	b.bookToAuthorDo.ReplaceConnPool(db.Statement.ConnPool)
	return b
}

func (b bookToAuthor) replaceDB(db *gorm.DB) bookToAuthor {
	b.bookToAuthorDo.ReplaceDB(db)
	return b
}

type bookToAuthorDo struct{ gen.DO }

type IBookToAuthorDo interface {
	gen.SubQuery
	Debug() IBookToAuthorDo
	WithContext(ctx context.Context) IBookToAuthorDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IBookToAuthorDo
	WriteDB() IBookToAuthorDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IBookToAuthorDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IBookToAuthorDo
	Not(conds ...gen.Condition) IBookToAuthorDo
	Or(conds ...gen.Condition) IBookToAuthorDo
	Select(conds ...field.Expr) IBookToAuthorDo
	Where(conds ...gen.Condition) IBookToAuthorDo
	Order(conds ...field.Expr) IBookToAuthorDo
	Distinct(cols ...field.Expr) IBookToAuthorDo
	Omit(cols ...field.Expr) IBookToAuthorDo
	Join(table schema.Tabler, on ...field.Expr) IBookToAuthorDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IBookToAuthorDo
	RightJoin(table schema.Tabler, on ...field.Expr) IBookToAuthorDo
	Group(cols ...field.Expr) IBookToAuthorDo
	Having(conds ...gen.Condition) IBookToAuthorDo
	Limit(limit int) IBookToAuthorDo
	Offset(offset int) IBookToAuthorDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IBookToAuthorDo
	Unscoped() IBookToAuthorDo
	Create(values ...*model.BookToAuthor) error
	CreateInBatches(values []*model.BookToAuthor, batchSize int) error
	Save(values ...*model.BookToAuthor) error
	First() (*model.BookToAuthor, error)
	Take() (*model.BookToAuthor, error)
	Last() (*model.BookToAuthor, error)
	Find() ([]*model.BookToAuthor, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.BookToAuthor, err error)
	FindInBatches(result *[]*model.BookToAuthor, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.BookToAuthor) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IBookToAuthorDo
	Assign(attrs ...field.AssignExpr) IBookToAuthorDo
	Joins(fields ...field.RelationField) IBookToAuthorDo
	Preload(fields ...field.RelationField) IBookToAuthorDo
	FirstOrInit() (*model.BookToAuthor, error)
	FirstOrCreate() (*model.BookToAuthor, error)
	FindByPage(offset int, limit int) (result []*model.BookToAuthor, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IBookToAuthorDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (b bookToAuthorDo) Debug() IBookToAuthorDo {
	return b.withDO(b.DO.Debug())
}

func (b bookToAuthorDo) WithContext(ctx context.Context) IBookToAuthorDo {
	return b.withDO(b.DO.WithContext(ctx))
}

func (b bookToAuthorDo) ReadDB() IBookToAuthorDo {
	return b.Clauses(dbresolver.Read)
}

func (b bookToAuthorDo) WriteDB() IBookToAuthorDo {
	return b.Clauses(dbresolver.Write)
}

func (b bookToAuthorDo) Session(config *gorm.Session) IBookToAuthorDo {
	return b.withDO(b.DO.Session(config))
}

func (b bookToAuthorDo) Clauses(conds ...clause.Expression) IBookToAuthorDo {
	return b.withDO(b.DO.Clauses(conds...))
}

func (b bookToAuthorDo) Returning(value interface{}, columns ...string) IBookToAuthorDo {
	return b.withDO(b.DO.Returning(value, columns...))
}

func (b bookToAuthorDo) Not(conds ...gen.Condition) IBookToAuthorDo {
	return b.withDO(b.DO.Not(conds...))
}

func (b bookToAuthorDo) Or(conds ...gen.Condition) IBookToAuthorDo {
	return b.withDO(b.DO.Or(conds...))
}

func (b bookToAuthorDo) Select(conds ...field.Expr) IBookToAuthorDo {
	return b.withDO(b.DO.Select(conds...))
}

func (b bookToAuthorDo) Where(conds ...gen.Condition) IBookToAuthorDo {
	return b.withDO(b.DO.Where(conds...))
}

func (b bookToAuthorDo) Order(conds ...field.Expr) IBookToAuthorDo {
	return b.withDO(b.DO.Order(conds...))
}

func (b bookToAuthorDo) Distinct(cols ...field.Expr) IBookToAuthorDo {
	return b.withDO(b.DO.Distinct(cols...))
}

func (b bookToAuthorDo) Omit(cols ...field.Expr) IBookToAuthorDo {
	return b.withDO(b.DO.Omit(cols...))
}

func (b bookToAuthorDo) Join(table schema.Tabler, on ...field.Expr) IBookToAuthorDo {
	return b.withDO(b.DO.Join(table, on...))
}

func (b bookToAuthorDo) LeftJoin(table schema.Tabler, on ...field.Expr) IBookToAuthorDo {
	return b.withDO(b.DO.LeftJoin(table, on...))
}

func (b bookToAuthorDo) RightJoin(table schema.Tabler, on ...field.Expr) IBookToAuthorDo {
	return b.withDO(b.DO.RightJoin(table, on...))
}

func (b bookToAuthorDo) Group(cols ...field.Expr) IBookToAuthorDo {
	return b.withDO(b.DO.Group(cols...))
}

func (b bookToAuthorDo) Having(conds ...gen.Condition) IBookToAuthorDo {
	return b.withDO(b.DO.Having(conds...))
}

func (b bookToAuthorDo) Limit(limit int) IBookToAuthorDo {
	return b.withDO(b.DO.Limit(limit))
}

func (b bookToAuthorDo) Offset(offset int) IBookToAuthorDo {
	return b.withDO(b.DO.Offset(offset))
}

func (b bookToAuthorDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IBookToAuthorDo {
	return b.withDO(b.DO.Scopes(funcs...))
}

func (b bookToAuthorDo) Unscoped() IBookToAuthorDo {
	return b.withDO(b.DO.Unscoped())
}

func (b bookToAuthorDo) Create(values ...*model.BookToAuthor) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Create(values)
}

func (b bookToAuthorDo) CreateInBatches(values []*model.BookToAuthor, batchSize int) error {
	return b.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (b bookToAuthorDo) Save(values ...*model.BookToAuthor) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Save(values)
}

func (b bookToAuthorDo) First() (*model.BookToAuthor, error) {
	if result, err := b.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.BookToAuthor), nil
	}
}

func (b bookToAuthorDo) Take() (*model.BookToAuthor, error) {
	if result, err := b.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.BookToAuthor), nil
	}
}

func (b bookToAuthorDo) Last() (*model.BookToAuthor, error) {
	if result, err := b.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.BookToAuthor), nil
	}
}

func (b bookToAuthorDo) Find() ([]*model.BookToAuthor, error) {
	result, err := b.DO.Find()
	return result.([]*model.BookToAuthor), err
}

func (b bookToAuthorDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.BookToAuthor, err error) {
	buf := make([]*model.BookToAuthor, 0, batchSize)
	err = b.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (b bookToAuthorDo) FindInBatches(result *[]*model.BookToAuthor, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return b.DO.FindInBatches(result, batchSize, fc)
}

func (b bookToAuthorDo) Attrs(attrs ...field.AssignExpr) IBookToAuthorDo {
	return b.withDO(b.DO.Attrs(attrs...))
}

func (b bookToAuthorDo) Assign(attrs ...field.AssignExpr) IBookToAuthorDo {
	return b.withDO(b.DO.Assign(attrs...))
}

func (b bookToAuthorDo) Joins(fields ...field.RelationField) IBookToAuthorDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Joins(_f))
	}
	return &b
}

func (b bookToAuthorDo) Preload(fields ...field.RelationField) IBookToAuthorDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Preload(_f))
	}
	return &b
}

func (b bookToAuthorDo) FirstOrInit() (*model.BookToAuthor, error) {
	if result, err := b.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.BookToAuthor), nil
	}
}

func (b bookToAuthorDo) FirstOrCreate() (*model.BookToAuthor, error) {
	if result, err := b.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.BookToAuthor), nil
	}
}

func (b bookToAuthorDo) FindByPage(offset int, limit int) (result []*model.BookToAuthor, count int64, err error) {
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

func (b bookToAuthorDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = b.Count()
	if err != nil {
		return
	}

	err = b.Offset(offset).Limit(limit).Scan(result)
	return
}

func (b bookToAuthorDo) Scan(result interface{}) (err error) {
	return b.DO.Scan(result)
}

func (b bookToAuthorDo) Delete(models ...*model.BookToAuthor) (result gen.ResultInfo, err error) {
	return b.DO.Delete(models)
}

func (b *bookToAuthorDo) withDO(do gen.Dao) *bookToAuthorDo {
	b.DO = *do.(*gen.DO)
	return b
}
