package repository

type FindOption struct {
	limit        int
	offset       int
	field, group string
	join         []string
	where        string
	whereArgs    []interface{}
	count        *int64
	preload      []string
	alias        string
	distinct     []interface{}
}

func NewFindOption() *FindOption {
	return &FindOption{}
}

func (f *FindOption) Distinct() []interface{} {
	return f.distinct
}

func (f *FindOption) SetDistinct(distinct ...interface{}) *FindOption {
	f.distinct = distinct
	return f
}

func (f *FindOption) SetAlias(alias string) *FindOption {
	f.alias = alias
	return f
}

func (f *FindOption) Alias() string {
	return f.alias
}

func (f *FindOption) Limit() int {
	return f.limit
}

func (f *FindOption) SetLimit(limit int) *FindOption {
	f.limit = limit
	return f
}

func (f *FindOption) Offset() int {
	return f.offset
}

func (f *FindOption) SetOffset(offset int) *FindOption {
	f.offset = offset
	return f
}

func (f *FindOption) Field() string {
	return f.field
}

func (f *FindOption) SetField(field string) *FindOption {
	f.field = field
	return f
}

func (f *FindOption) Group() string {
	return f.group
}

func (f *FindOption) SetGroup(group string) *FindOption {
	f.group = group
	return f
}

func (f *FindOption) Join() []string {
	return f.join
}

func (f *FindOption) SetJoin(join []string) *FindOption {
	f.join = join
	return f
}

func (f *FindOption) Where() string {
	return f.where
}

func (f *FindOption) SetWhere(where string, args ...interface{}) *FindOption {
	f.where = where
	f.SetWhereArgs(args...)
	return f
}

func (f *FindOption) WhereArgs() []interface{} {
	return f.whereArgs
}

func (f *FindOption) SetWhereArgs(whereArgs ...interface{}) *FindOption {
	f.whereArgs = whereArgs
	return f
}

func (f *FindOption) Count() *int64 {
	return f.count
}

func (f *FindOption) SetCount(count *int64) *FindOption {
	f.count = count
	return f
}

func (f *FindOption) Preload() []string {
	return f.preload
}

func (f *FindOption) SetPreload(preload ...string) *FindOption {
	f.preload = append(f.preload, preload...)
	return f
}
