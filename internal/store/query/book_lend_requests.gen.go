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

func newBookLendRequest(db *gorm.DB, opts ...gen.DOOption) bookLendRequest {
	_bookLendRequest := bookLendRequest{}

	_bookLendRequest.bookLendRequestDo.UseDB(db, opts...)
	_bookLendRequest.bookLendRequestDo.UseModel(&model.BookLendRequest{})

	tableName := _bookLendRequest.bookLendRequestDo.TableName()
	_bookLendRequest.ALL = field.NewAsterisk(tableName)
	_bookLendRequest.ID = field.NewField(tableName, "id")
	_bookLendRequest.UserID = field.NewInt64(tableName, "user_id")
	_bookLendRequest.BookID = field.NewInt64(tableName, "book_id")

	_bookLendRequest.fillFieldMap()

	return _bookLendRequest
}

type bookLendRequest struct {
	bookLendRequestDo bookLendRequestDo

	ALL    field.Asterisk
	ID     field.Field
	UserID field.Int64
	BookID field.Int64

	fieldMap map[string]field.Expr
}

func (b bookLendRequest) Table(newTableName string) *bookLendRequest {
	b.bookLendRequestDo.UseTable(newTableName)
	return b.updateTableName(newTableName)
}

func (b bookLendRequest) As(alias string) *bookLendRequest {
	b.bookLendRequestDo.DO = *(b.bookLendRequestDo.As(alias).(*gen.DO))
	return b.updateTableName(alias)
}

func (b *bookLendRequest) updateTableName(table string) *bookLendRequest {
	b.ALL = field.NewAsterisk(table)
	b.ID = field.NewField(table, "id")
	b.UserID = field.NewInt64(table, "user_id")
	b.BookID = field.NewInt64(table, "book_id")

	b.fillFieldMap()

	return b
}

func (b *bookLendRequest) WithContext(ctx context.Context) IBookLendRequestDo {
	return b.bookLendRequestDo.WithContext(ctx)
}

func (b bookLendRequest) TableName() string { return b.bookLendRequestDo.TableName() }

func (b bookLendRequest) Alias() string { return b.bookLendRequestDo.Alias() }

func (b bookLendRequest) Columns(cols ...field.Expr) gen.Columns {
	return b.bookLendRequestDo.Columns(cols...)
}

func (b *bookLendRequest) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := b.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (b *bookLendRequest) fillFieldMap() {
	b.fieldMap = make(map[string]field.Expr, 3)
	b.fieldMap["id"] = b.ID
	b.fieldMap["user_id"] = b.UserID
	b.fieldMap["book_id"] = b.BookID
}

func (b bookLendRequest) clone(db *gorm.DB) bookLendRequest {
	b.bookLendRequestDo.ReplaceConnPool(db.Statement.ConnPool)
	return b
}

func (b bookLendRequest) replaceDB(db *gorm.DB) bookLendRequest {
	b.bookLendRequestDo.ReplaceDB(db)
	return b
}

type bookLendRequestDo struct{ gen.DO }

type IBookLendRequestDo interface {
	gen.SubQuery
	Debug() IBookLendRequestDo
	WithContext(ctx context.Context) IBookLendRequestDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IBookLendRequestDo
	WriteDB() IBookLendRequestDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IBookLendRequestDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IBookLendRequestDo
	Not(conds ...gen.Condition) IBookLendRequestDo
	Or(conds ...gen.Condition) IBookLendRequestDo
	Select(conds ...field.Expr) IBookLendRequestDo
	Where(conds ...gen.Condition) IBookLendRequestDo
	Order(conds ...field.Expr) IBookLendRequestDo
	Distinct(cols ...field.Expr) IBookLendRequestDo
	Omit(cols ...field.Expr) IBookLendRequestDo
	Join(table schema.Tabler, on ...field.Expr) IBookLendRequestDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IBookLendRequestDo
	RightJoin(table schema.Tabler, on ...field.Expr) IBookLendRequestDo
	Group(cols ...field.Expr) IBookLendRequestDo
	Having(conds ...gen.Condition) IBookLendRequestDo
	Limit(limit int) IBookLendRequestDo
	Offset(offset int) IBookLendRequestDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IBookLendRequestDo
	Unscoped() IBookLendRequestDo
	Create(values ...*model.BookLendRequest) error
	CreateInBatches(values []*model.BookLendRequest, batchSize int) error
	Save(values ...*model.BookLendRequest) error
	First() (*model.BookLendRequest, error)
	Take() (*model.BookLendRequest, error)
	Last() (*model.BookLendRequest, error)
	Find() ([]*model.BookLendRequest, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.BookLendRequest, err error)
	FindInBatches(result *[]*model.BookLendRequest, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.BookLendRequest) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IBookLendRequestDo
	Assign(attrs ...field.AssignExpr) IBookLendRequestDo
	Joins(fields ...field.RelationField) IBookLendRequestDo
	Preload(fields ...field.RelationField) IBookLendRequestDo
	FirstOrInit() (*model.BookLendRequest, error)
	FirstOrCreate() (*model.BookLendRequest, error)
	FindByPage(offset int, limit int) (result []*model.BookLendRequest, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IBookLendRequestDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (b bookLendRequestDo) Debug() IBookLendRequestDo {
	return b.withDO(b.DO.Debug())
}

func (b bookLendRequestDo) WithContext(ctx context.Context) IBookLendRequestDo {
	return b.withDO(b.DO.WithContext(ctx))
}

func (b bookLendRequestDo) ReadDB() IBookLendRequestDo {
	return b.Clauses(dbresolver.Read)
}

func (b bookLendRequestDo) WriteDB() IBookLendRequestDo {
	return b.Clauses(dbresolver.Write)
}

func (b bookLendRequestDo) Session(config *gorm.Session) IBookLendRequestDo {
	return b.withDO(b.DO.Session(config))
}

func (b bookLendRequestDo) Clauses(conds ...clause.Expression) IBookLendRequestDo {
	return b.withDO(b.DO.Clauses(conds...))
}

func (b bookLendRequestDo) Returning(value interface{}, columns ...string) IBookLendRequestDo {
	return b.withDO(b.DO.Returning(value, columns...))
}

func (b bookLendRequestDo) Not(conds ...gen.Condition) IBookLendRequestDo {
	return b.withDO(b.DO.Not(conds...))
}

func (b bookLendRequestDo) Or(conds ...gen.Condition) IBookLendRequestDo {
	return b.withDO(b.DO.Or(conds...))
}

func (b bookLendRequestDo) Select(conds ...field.Expr) IBookLendRequestDo {
	return b.withDO(b.DO.Select(conds...))
}

func (b bookLendRequestDo) Where(conds ...gen.Condition) IBookLendRequestDo {
	return b.withDO(b.DO.Where(conds...))
}

func (b bookLendRequestDo) Order(conds ...field.Expr) IBookLendRequestDo {
	return b.withDO(b.DO.Order(conds...))
}

func (b bookLendRequestDo) Distinct(cols ...field.Expr) IBookLendRequestDo {
	return b.withDO(b.DO.Distinct(cols...))
}

func (b bookLendRequestDo) Omit(cols ...field.Expr) IBookLendRequestDo {
	return b.withDO(b.DO.Omit(cols...))
}

func (b bookLendRequestDo) Join(table schema.Tabler, on ...field.Expr) IBookLendRequestDo {
	return b.withDO(b.DO.Join(table, on...))
}

func (b bookLendRequestDo) LeftJoin(table schema.Tabler, on ...field.Expr) IBookLendRequestDo {
	return b.withDO(b.DO.LeftJoin(table, on...))
}

func (b bookLendRequestDo) RightJoin(table schema.Tabler, on ...field.Expr) IBookLendRequestDo {
	return b.withDO(b.DO.RightJoin(table, on...))
}

func (b bookLendRequestDo) Group(cols ...field.Expr) IBookLendRequestDo {
	return b.withDO(b.DO.Group(cols...))
}

func (b bookLendRequestDo) Having(conds ...gen.Condition) IBookLendRequestDo {
	return b.withDO(b.DO.Having(conds...))
}

func (b bookLendRequestDo) Limit(limit int) IBookLendRequestDo {
	return b.withDO(b.DO.Limit(limit))
}

func (b bookLendRequestDo) Offset(offset int) IBookLendRequestDo {
	return b.withDO(b.DO.Offset(offset))
}

func (b bookLendRequestDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IBookLendRequestDo {
	return b.withDO(b.DO.Scopes(funcs...))
}

func (b bookLendRequestDo) Unscoped() IBookLendRequestDo {
	return b.withDO(b.DO.Unscoped())
}

func (b bookLendRequestDo) Create(values ...*model.BookLendRequest) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Create(values)
}

func (b bookLendRequestDo) CreateInBatches(values []*model.BookLendRequest, batchSize int) error {
	return b.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (b bookLendRequestDo) Save(values ...*model.BookLendRequest) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Save(values)
}

func (b bookLendRequestDo) First() (*model.BookLendRequest, error) {
	if result, err := b.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.BookLendRequest), nil
	}
}

func (b bookLendRequestDo) Take() (*model.BookLendRequest, error) {
	if result, err := b.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.BookLendRequest), nil
	}
}

func (b bookLendRequestDo) Last() (*model.BookLendRequest, error) {
	if result, err := b.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.BookLendRequest), nil
	}
}

func (b bookLendRequestDo) Find() ([]*model.BookLendRequest, error) {
	result, err := b.DO.Find()
	return result.([]*model.BookLendRequest), err
}

func (b bookLendRequestDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.BookLendRequest, err error) {
	buf := make([]*model.BookLendRequest, 0, batchSize)
	err = b.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (b bookLendRequestDo) FindInBatches(result *[]*model.BookLendRequest, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return b.DO.FindInBatches(result, batchSize, fc)
}

func (b bookLendRequestDo) Attrs(attrs ...field.AssignExpr) IBookLendRequestDo {
	return b.withDO(b.DO.Attrs(attrs...))
}

func (b bookLendRequestDo) Assign(attrs ...field.AssignExpr) IBookLendRequestDo {
	return b.withDO(b.DO.Assign(attrs...))
}

func (b bookLendRequestDo) Joins(fields ...field.RelationField) IBookLendRequestDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Joins(_f))
	}
	return &b
}

func (b bookLendRequestDo) Preload(fields ...field.RelationField) IBookLendRequestDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Preload(_f))
	}
	return &b
}

func (b bookLendRequestDo) FirstOrInit() (*model.BookLendRequest, error) {
	if result, err := b.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.BookLendRequest), nil
	}
}

func (b bookLendRequestDo) FirstOrCreate() (*model.BookLendRequest, error) {
	if result, err := b.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.BookLendRequest), nil
	}
}

func (b bookLendRequestDo) FindByPage(offset int, limit int) (result []*model.BookLendRequest, count int64, err error) {
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

func (b bookLendRequestDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = b.Count()
	if err != nil {
		return
	}

	err = b.Offset(offset).Limit(limit).Scan(result)
	return
}

func (b bookLendRequestDo) Scan(result interface{}) (err error) {
	return b.DO.Scan(result)
}

func (b bookLendRequestDo) Delete(models ...*model.BookLendRequest) (result gen.ResultInfo, err error) {
	return b.DO.Delete(models)
}

func (b *bookLendRequestDo) withDO(do gen.Dao) *bookLendRequestDo {
	b.DO = *do.(*gen.DO)
	return b
}
