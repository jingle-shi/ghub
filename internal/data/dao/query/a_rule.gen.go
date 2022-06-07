// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"ghub/internal/data/dao/model"

	"gorm.io/gen"
	"gorm.io/gen/field"
)

func newRule(db *gorm.DB) rule {
	_rule := rule{}

	_rule.ruleDo.UseDB(db)
	_rule.ruleDo.UseModel(&model.Rule{})

	tableName := _rule.ruleDo.TableName()
	_rule.ALL = field.NewField(tableName, "*")
	_rule.ID = field.NewInt32(tableName, "id")
	_rule.Domain = field.NewString(tableName, "domain")
	_rule.Name = field.NewString(tableName, "name")
	_rule.Remark = field.NewString(tableName, "remark")
	_rule.CreatedAt = field.NewTime(tableName, "created_at")
	_rule.UpdatedAt = field.NewTime(tableName, "updated_at")
	_rule.DeletedAt = field.NewField(tableName, "deleted_at")

	_rule.fillFieldMap()

	return _rule
}

type rule struct {
	ruleDo ruleDo

	ALL       field.Field
	ID        field.Int32
	Domain    field.String
	Name      field.String
	Remark    field.String
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field

	fieldMap map[string]field.Expr
}

func (r rule) Table(newTableName string) *rule {
	r.ruleDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r rule) As(alias string) *rule {
	r.ruleDo.DO = *(r.ruleDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *rule) updateTableName(table string) *rule {
	r.ALL = field.NewField(table, "*")
	r.ID = field.NewInt32(table, "id")
	r.Domain = field.NewString(table, "domain")
	r.Name = field.NewString(table, "name")
	r.Remark = field.NewString(table, "remark")
	r.CreatedAt = field.NewTime(table, "created_at")
	r.UpdatedAt = field.NewTime(table, "updated_at")
	r.DeletedAt = field.NewField(table, "deleted_at")

	r.fillFieldMap()

	return r
}

func (r *rule) WithContext(ctx context.Context) *ruleDo { return r.ruleDo.WithContext(ctx) }

func (r rule) TableName() string { return r.ruleDo.TableName() }

func (r *rule) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *rule) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 7)
	r.fieldMap["id"] = r.ID
	r.fieldMap["domain"] = r.Domain
	r.fieldMap["name"] = r.Name
	r.fieldMap["remark"] = r.Remark
	r.fieldMap["created_at"] = r.CreatedAt
	r.fieldMap["updated_at"] = r.UpdatedAt
	r.fieldMap["deleted_at"] = r.DeletedAt
}

func (r rule) clone(db *gorm.DB) rule {
	r.ruleDo.ReplaceDB(db)
	return r
}

type ruleDo struct{ gen.DO }

func (r ruleDo) Debug() *ruleDo {
	return r.withDO(r.DO.Debug())
}

func (r ruleDo) WithContext(ctx context.Context) *ruleDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r ruleDo) Clauses(conds ...clause.Expression) *ruleDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r ruleDo) Returning(value interface{}, columns ...string) *ruleDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r ruleDo) Not(conds ...gen.Condition) *ruleDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r ruleDo) Or(conds ...gen.Condition) *ruleDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r ruleDo) Select(conds ...field.Expr) *ruleDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r ruleDo) Where(conds ...gen.Condition) *ruleDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r ruleDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *ruleDo {
	return r.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (r ruleDo) Order(conds ...field.Expr) *ruleDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r ruleDo) Distinct(cols ...field.Expr) *ruleDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r ruleDo) Omit(cols ...field.Expr) *ruleDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r ruleDo) Join(table schema.Tabler, on ...field.Expr) *ruleDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r ruleDo) LeftJoin(table schema.Tabler, on ...field.Expr) *ruleDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r ruleDo) RightJoin(table schema.Tabler, on ...field.Expr) *ruleDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r ruleDo) Group(cols ...field.Expr) *ruleDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r ruleDo) Having(conds ...gen.Condition) *ruleDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r ruleDo) Limit(limit int) *ruleDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r ruleDo) Offset(offset int) *ruleDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r ruleDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *ruleDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r ruleDo) Unscoped() *ruleDo {
	return r.withDO(r.DO.Unscoped())
}

func (r ruleDo) Create(values ...*model.Rule) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r ruleDo) CreateInBatches(values []*model.Rule, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r ruleDo) Save(values ...*model.Rule) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r ruleDo) First() (*model.Rule, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Rule), nil
	}
}

func (r ruleDo) Take() (*model.Rule, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Rule), nil
	}
}

func (r ruleDo) Last() (*model.Rule, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Rule), nil
	}
}

func (r ruleDo) Find() ([]*model.Rule, error) {
	result, err := r.DO.Find()
	return result.([]*model.Rule), err
}

func (r ruleDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Rule, err error) {
	buf := make([]*model.Rule, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r ruleDo) FindInBatches(result *[]*model.Rule, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r ruleDo) Attrs(attrs ...field.AssignExpr) *ruleDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r ruleDo) Assign(attrs ...field.AssignExpr) *ruleDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r ruleDo) Joins(field field.RelationField) *ruleDo {
	return r.withDO(r.DO.Joins(field))
}

func (r ruleDo) Preload(field field.RelationField) *ruleDo {
	return r.withDO(r.DO.Preload(field))
}

func (r ruleDo) FirstOrInit() (*model.Rule, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Rule), nil
	}
}

func (r ruleDo) FirstOrCreate() (*model.Rule, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Rule), nil
	}
}

func (r ruleDo) FindByPage(offset int, limit int) (result []*model.Rule, count int64, err error) {
	if limit <= 0 {
		count, err = r.Count()
		return
	}

	result, err = r.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = r.Offset(-1).Limit(-1).Count()
	return
}

func (r ruleDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r *ruleDo) withDO(do gen.Dao) *ruleDo {
	r.DO = *do.(*gen.DO)
	return r
}