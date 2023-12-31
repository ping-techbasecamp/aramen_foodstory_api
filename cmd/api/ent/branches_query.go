// Code generated by ent, DO NOT EDIT.

package ent

import (
	"aramen-api/cmd/api/ent/branches"
	"aramen-api/cmd/api/ent/predicate"
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BranchesQuery is the builder for querying Branches entities.
type BranchesQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Branches
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BranchesQuery builder.
func (bq *BranchesQuery) Where(ps ...predicate.Branches) *BranchesQuery {
	bq.predicates = append(bq.predicates, ps...)
	return bq
}

// Limit adds a limit step to the query.
func (bq *BranchesQuery) Limit(limit int) *BranchesQuery {
	bq.limit = &limit
	return bq
}

// Offset adds an offset step to the query.
func (bq *BranchesQuery) Offset(offset int) *BranchesQuery {
	bq.offset = &offset
	return bq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (bq *BranchesQuery) Unique(unique bool) *BranchesQuery {
	bq.unique = &unique
	return bq
}

// Order adds an order step to the query.
func (bq *BranchesQuery) Order(o ...OrderFunc) *BranchesQuery {
	bq.order = append(bq.order, o...)
	return bq
}

// First returns the first Branches entity from the query.
// Returns a *NotFoundError when no Branches was found.
func (bq *BranchesQuery) First(ctx context.Context) (*Branches, error) {
	nodes, err := bq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{branches.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (bq *BranchesQuery) FirstX(ctx context.Context) *Branches {
	node, err := bq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Branches ID from the query.
// Returns a *NotFoundError when no Branches ID was found.
func (bq *BranchesQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{branches.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (bq *BranchesQuery) FirstIDX(ctx context.Context) int {
	id, err := bq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Branches entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Branches entity is found.
// Returns a *NotFoundError when no Branches entities are found.
func (bq *BranchesQuery) Only(ctx context.Context) (*Branches, error) {
	nodes, err := bq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{branches.Label}
	default:
		return nil, &NotSingularError{branches.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (bq *BranchesQuery) OnlyX(ctx context.Context) *Branches {
	node, err := bq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Branches ID in the query.
// Returns a *NotSingularError when more than one Branches ID is found.
// Returns a *NotFoundError when no entities are found.
func (bq *BranchesQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{branches.Label}
	default:
		err = &NotSingularError{branches.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (bq *BranchesQuery) OnlyIDX(ctx context.Context) int {
	id, err := bq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of BranchesSlice.
func (bq *BranchesQuery) All(ctx context.Context) ([]*Branches, error) {
	if err := bq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return bq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (bq *BranchesQuery) AllX(ctx context.Context) []*Branches {
	nodes, err := bq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Branches IDs.
func (bq *BranchesQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := bq.Select(branches.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (bq *BranchesQuery) IDsX(ctx context.Context) []int {
	ids, err := bq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (bq *BranchesQuery) Count(ctx context.Context) (int, error) {
	if err := bq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return bq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (bq *BranchesQuery) CountX(ctx context.Context) int {
	count, err := bq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (bq *BranchesQuery) Exist(ctx context.Context) (bool, error) {
	if err := bq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return bq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (bq *BranchesQuery) ExistX(ctx context.Context) bool {
	exist, err := bq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BranchesQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (bq *BranchesQuery) Clone() *BranchesQuery {
	if bq == nil {
		return nil
	}
	return &BranchesQuery{
		config:     bq.config,
		limit:      bq.limit,
		offset:     bq.offset,
		order:      append([]OrderFunc{}, bq.order...),
		predicates: append([]predicate.Branches{}, bq.predicates...),
		// clone intermediate query.
		sql:    bq.sql.Clone(),
		path:   bq.path,
		unique: bq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Branches.Query().
//		GroupBy(branches.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (bq *BranchesQuery) GroupBy(field string, fields ...string) *BranchesGroupBy {
	grbuild := &BranchesGroupBy{config: bq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := bq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return bq.sqlQuery(ctx), nil
	}
	grbuild.label = branches.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.Branches.Query().
//		Select(branches.FieldName).
//		Scan(ctx, &v)
//
func (bq *BranchesQuery) Select(fields ...string) *BranchesSelect {
	bq.fields = append(bq.fields, fields...)
	selbuild := &BranchesSelect{BranchesQuery: bq}
	selbuild.label = branches.Label
	selbuild.flds, selbuild.scan = &bq.fields, selbuild.Scan
	return selbuild
}

func (bq *BranchesQuery) prepareQuery(ctx context.Context) error {
	for _, f := range bq.fields {
		if !branches.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if bq.path != nil {
		prev, err := bq.path(ctx)
		if err != nil {
			return err
		}
		bq.sql = prev
	}
	return nil
}

func (bq *BranchesQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Branches, error) {
	var (
		nodes = []*Branches{}
		_spec = bq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Branches).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Branches{config: bq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, bq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (bq *BranchesQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := bq.querySpec()
	_spec.Node.Columns = bq.fields
	if len(bq.fields) > 0 {
		_spec.Unique = bq.unique != nil && *bq.unique
	}
	return sqlgraph.CountNodes(ctx, bq.driver, _spec)
}

func (bq *BranchesQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := bq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (bq *BranchesQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   branches.Table,
			Columns: branches.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: branches.FieldID,
			},
		},
		From:   bq.sql,
		Unique: true,
	}
	if unique := bq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := bq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, branches.FieldID)
		for i := range fields {
			if fields[i] != branches.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := bq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := bq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := bq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := bq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (bq *BranchesQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(bq.driver.Dialect())
	t1 := builder.Table(branches.Table)
	columns := bq.fields
	if len(columns) == 0 {
		columns = branches.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if bq.sql != nil {
		selector = bq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if bq.unique != nil && *bq.unique {
		selector.Distinct()
	}
	for _, p := range bq.predicates {
		p(selector)
	}
	for _, p := range bq.order {
		p(selector)
	}
	if offset := bq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := bq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// BranchesGroupBy is the group-by builder for Branches entities.
type BranchesGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bgb *BranchesGroupBy) Aggregate(fns ...AggregateFunc) *BranchesGroupBy {
	bgb.fns = append(bgb.fns, fns...)
	return bgb
}

// Scan applies the group-by query and scans the result into the given value.
func (bgb *BranchesGroupBy) Scan(ctx context.Context, v any) error {
	query, err := bgb.path(ctx)
	if err != nil {
		return err
	}
	bgb.sql = query
	return bgb.sqlScan(ctx, v)
}

func (bgb *BranchesGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range bgb.fields {
		if !branches.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := bgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (bgb *BranchesGroupBy) sqlQuery() *sql.Selector {
	selector := bgb.sql.Select()
	aggregation := make([]string, 0, len(bgb.fns))
	for _, fn := range bgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(bgb.fields)+len(bgb.fns))
		for _, f := range bgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(bgb.fields...)...)
}

// BranchesSelect is the builder for selecting fields of Branches entities.
type BranchesSelect struct {
	*BranchesQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (bs *BranchesSelect) Scan(ctx context.Context, v any) error {
	if err := bs.prepareQuery(ctx); err != nil {
		return err
	}
	bs.sql = bs.BranchesQuery.sqlQuery(ctx)
	return bs.sqlScan(ctx, v)
}

func (bs *BranchesSelect) sqlScan(ctx context.Context, v any) error {
	rows := &sql.Rows{}
	query, args := bs.sql.Query()
	if err := bs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
