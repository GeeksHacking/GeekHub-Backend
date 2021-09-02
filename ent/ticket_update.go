// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/geekshacking/geekhub-backend/ent/predicate"
	"github.com/geekshacking/geekhub-backend/ent/project"
	"github.com/geekshacking/geekhub-backend/ent/ticket"
	"github.com/geekshacking/geekhub-backend/ent/user"
)

// TicketUpdate is the builder for updating Ticket entities.
type TicketUpdate struct {
	config
	hooks    []Hook
	mutation *TicketMutation
}

// Where appends a list predicates to the TicketUpdate builder.
func (tu *TicketUpdate) Where(ps ...predicate.Ticket) *TicketUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetName sets the "name" field.
func (tu *TicketUpdate) SetName(s string) *TicketUpdate {
	tu.mutation.SetName(s)
	return tu
}

// SetDescription sets the "description" field.
func (tu *TicketUpdate) SetDescription(s string) *TicketUpdate {
	tu.mutation.SetDescription(s)
	return tu
}

// SetType sets the "type" field.
func (tu *TicketUpdate) SetType(t ticket.Type) *TicketUpdate {
	tu.mutation.SetType(t)
	return tu
}

// SetStatus sets the "status" field.
func (tu *TicketUpdate) SetStatus(t ticket.Status) *TicketUpdate {
	tu.mutation.SetStatus(t)
	return tu
}

// SetProjectID sets the "project" edge to the Project entity by ID.
func (tu *TicketUpdate) SetProjectID(id int) *TicketUpdate {
	tu.mutation.SetProjectID(id)
	return tu
}

// SetNillableProjectID sets the "project" edge to the Project entity by ID if the given value is not nil.
func (tu *TicketUpdate) SetNillableProjectID(id *int) *TicketUpdate {
	if id != nil {
		tu = tu.SetProjectID(*id)
	}
	return tu
}

// SetProject sets the "project" edge to the Project entity.
func (tu *TicketUpdate) SetProject(p *Project) *TicketUpdate {
	return tu.SetProjectID(p.ID)
}

// AddReporterIDs adds the "reporter" edge to the User entity by IDs.
func (tu *TicketUpdate) AddReporterIDs(ids ...int) *TicketUpdate {
	tu.mutation.AddReporterIDs(ids...)
	return tu
}

// AddReporter adds the "reporter" edges to the User entity.
func (tu *TicketUpdate) AddReporter(u ...*User) *TicketUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tu.AddReporterIDs(ids...)
}

// AddAssigneeIDs adds the "assignee" edge to the User entity by IDs.
func (tu *TicketUpdate) AddAssigneeIDs(ids ...int) *TicketUpdate {
	tu.mutation.AddAssigneeIDs(ids...)
	return tu
}

// AddAssignee adds the "assignee" edges to the User entity.
func (tu *TicketUpdate) AddAssignee(u ...*User) *TicketUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tu.AddAssigneeIDs(ids...)
}

// SetParentID sets the "parent" edge to the Ticket entity by ID.
func (tu *TicketUpdate) SetParentID(id int) *TicketUpdate {
	tu.mutation.SetParentID(id)
	return tu
}

// SetNillableParentID sets the "parent" edge to the Ticket entity by ID if the given value is not nil.
func (tu *TicketUpdate) SetNillableParentID(id *int) *TicketUpdate {
	if id != nil {
		tu = tu.SetParentID(*id)
	}
	return tu
}

// SetParent sets the "parent" edge to the Ticket entity.
func (tu *TicketUpdate) SetParent(t *Ticket) *TicketUpdate {
	return tu.SetParentID(t.ID)
}

// Mutation returns the TicketMutation object of the builder.
func (tu *TicketUpdate) Mutation() *TicketMutation {
	return tu.mutation
}

// ClearProject clears the "project" edge to the Project entity.
func (tu *TicketUpdate) ClearProject() *TicketUpdate {
	tu.mutation.ClearProject()
	return tu
}

// ClearReporter clears all "reporter" edges to the User entity.
func (tu *TicketUpdate) ClearReporter() *TicketUpdate {
	tu.mutation.ClearReporter()
	return tu
}

// RemoveReporterIDs removes the "reporter" edge to User entities by IDs.
func (tu *TicketUpdate) RemoveReporterIDs(ids ...int) *TicketUpdate {
	tu.mutation.RemoveReporterIDs(ids...)
	return tu
}

// RemoveReporter removes "reporter" edges to User entities.
func (tu *TicketUpdate) RemoveReporter(u ...*User) *TicketUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tu.RemoveReporterIDs(ids...)
}

// ClearAssignee clears all "assignee" edges to the User entity.
func (tu *TicketUpdate) ClearAssignee() *TicketUpdate {
	tu.mutation.ClearAssignee()
	return tu
}

// RemoveAssigneeIDs removes the "assignee" edge to User entities by IDs.
func (tu *TicketUpdate) RemoveAssigneeIDs(ids ...int) *TicketUpdate {
	tu.mutation.RemoveAssigneeIDs(ids...)
	return tu
}

// RemoveAssignee removes "assignee" edges to User entities.
func (tu *TicketUpdate) RemoveAssignee(u ...*User) *TicketUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tu.RemoveAssigneeIDs(ids...)
}

// ClearParent clears the "parent" edge to the Ticket entity.
func (tu *TicketUpdate) ClearParent() *TicketUpdate {
	tu.mutation.ClearParent()
	return tu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TicketUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tu.hooks) == 0 {
		if err = tu.check(); err != nil {
			return 0, err
		}
		affected, err = tu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TicketMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tu.check(); err != nil {
				return 0, err
			}
			tu.mutation = mutation
			affected, err = tu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tu.hooks) - 1; i >= 0; i-- {
			if tu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TicketUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TicketUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TicketUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *TicketUpdate) check() error {
	if v, ok := tu.mutation.GetType(); ok {
		if err := ticket.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf("ent: validator failed for field \"type\": %w", err)}
		}
	}
	if v, ok := tu.mutation.Status(); ok {
		if err := ticket.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf("ent: validator failed for field \"status\": %w", err)}
		}
	}
	return nil
}

func (tu *TicketUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   ticket.Table,
			Columns: ticket.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: ticket.FieldID,
			},
		},
	}
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: ticket.FieldName,
		})
	}
	if value, ok := tu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: ticket.FieldDescription,
		})
	}
	if value, ok := tu.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: ticket.FieldType,
		})
	}
	if value, ok := tu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: ticket.FieldStatus,
		})
	}
	if tu.mutation.ProjectCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ticket.ProjectTable,
			Columns: []string{ticket.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: project.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.ProjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ticket.ProjectTable,
			Columns: []string{ticket.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: project.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.ReporterCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   ticket.ReporterTable,
			Columns: ticket.ReporterPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedReporterIDs(); len(nodes) > 0 && !tu.mutation.ReporterCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   ticket.ReporterTable,
			Columns: ticket.ReporterPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.ReporterIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   ticket.ReporterTable,
			Columns: ticket.ReporterPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.AssigneeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   ticket.AssigneeTable,
			Columns: ticket.AssigneePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedAssigneeIDs(); len(nodes) > 0 && !tu.mutation.AssigneeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   ticket.AssigneeTable,
			Columns: ticket.AssigneePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.AssigneeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   ticket.AssigneeTable,
			Columns: ticket.AssigneePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   ticket.ParentTable,
			Columns: []string{ticket.ParentColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: ticket.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   ticket.ParentTable,
			Columns: []string{ticket.ParentColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: ticket.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{ticket.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// TicketUpdateOne is the builder for updating a single Ticket entity.
type TicketUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TicketMutation
}

// SetName sets the "name" field.
func (tuo *TicketUpdateOne) SetName(s string) *TicketUpdateOne {
	tuo.mutation.SetName(s)
	return tuo
}

// SetDescription sets the "description" field.
func (tuo *TicketUpdateOne) SetDescription(s string) *TicketUpdateOne {
	tuo.mutation.SetDescription(s)
	return tuo
}

// SetType sets the "type" field.
func (tuo *TicketUpdateOne) SetType(t ticket.Type) *TicketUpdateOne {
	tuo.mutation.SetType(t)
	return tuo
}

// SetStatus sets the "status" field.
func (tuo *TicketUpdateOne) SetStatus(t ticket.Status) *TicketUpdateOne {
	tuo.mutation.SetStatus(t)
	return tuo
}

// SetProjectID sets the "project" edge to the Project entity by ID.
func (tuo *TicketUpdateOne) SetProjectID(id int) *TicketUpdateOne {
	tuo.mutation.SetProjectID(id)
	return tuo
}

// SetNillableProjectID sets the "project" edge to the Project entity by ID if the given value is not nil.
func (tuo *TicketUpdateOne) SetNillableProjectID(id *int) *TicketUpdateOne {
	if id != nil {
		tuo = tuo.SetProjectID(*id)
	}
	return tuo
}

// SetProject sets the "project" edge to the Project entity.
func (tuo *TicketUpdateOne) SetProject(p *Project) *TicketUpdateOne {
	return tuo.SetProjectID(p.ID)
}

// AddReporterIDs adds the "reporter" edge to the User entity by IDs.
func (tuo *TicketUpdateOne) AddReporterIDs(ids ...int) *TicketUpdateOne {
	tuo.mutation.AddReporterIDs(ids...)
	return tuo
}

// AddReporter adds the "reporter" edges to the User entity.
func (tuo *TicketUpdateOne) AddReporter(u ...*User) *TicketUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tuo.AddReporterIDs(ids...)
}

// AddAssigneeIDs adds the "assignee" edge to the User entity by IDs.
func (tuo *TicketUpdateOne) AddAssigneeIDs(ids ...int) *TicketUpdateOne {
	tuo.mutation.AddAssigneeIDs(ids...)
	return tuo
}

// AddAssignee adds the "assignee" edges to the User entity.
func (tuo *TicketUpdateOne) AddAssignee(u ...*User) *TicketUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tuo.AddAssigneeIDs(ids...)
}

// SetParentID sets the "parent" edge to the Ticket entity by ID.
func (tuo *TicketUpdateOne) SetParentID(id int) *TicketUpdateOne {
	tuo.mutation.SetParentID(id)
	return tuo
}

// SetNillableParentID sets the "parent" edge to the Ticket entity by ID if the given value is not nil.
func (tuo *TicketUpdateOne) SetNillableParentID(id *int) *TicketUpdateOne {
	if id != nil {
		tuo = tuo.SetParentID(*id)
	}
	return tuo
}

// SetParent sets the "parent" edge to the Ticket entity.
func (tuo *TicketUpdateOne) SetParent(t *Ticket) *TicketUpdateOne {
	return tuo.SetParentID(t.ID)
}

// Mutation returns the TicketMutation object of the builder.
func (tuo *TicketUpdateOne) Mutation() *TicketMutation {
	return tuo.mutation
}

// ClearProject clears the "project" edge to the Project entity.
func (tuo *TicketUpdateOne) ClearProject() *TicketUpdateOne {
	tuo.mutation.ClearProject()
	return tuo
}

// ClearReporter clears all "reporter" edges to the User entity.
func (tuo *TicketUpdateOne) ClearReporter() *TicketUpdateOne {
	tuo.mutation.ClearReporter()
	return tuo
}

// RemoveReporterIDs removes the "reporter" edge to User entities by IDs.
func (tuo *TicketUpdateOne) RemoveReporterIDs(ids ...int) *TicketUpdateOne {
	tuo.mutation.RemoveReporterIDs(ids...)
	return tuo
}

// RemoveReporter removes "reporter" edges to User entities.
func (tuo *TicketUpdateOne) RemoveReporter(u ...*User) *TicketUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tuo.RemoveReporterIDs(ids...)
}

// ClearAssignee clears all "assignee" edges to the User entity.
func (tuo *TicketUpdateOne) ClearAssignee() *TicketUpdateOne {
	tuo.mutation.ClearAssignee()
	return tuo
}

// RemoveAssigneeIDs removes the "assignee" edge to User entities by IDs.
func (tuo *TicketUpdateOne) RemoveAssigneeIDs(ids ...int) *TicketUpdateOne {
	tuo.mutation.RemoveAssigneeIDs(ids...)
	return tuo
}

// RemoveAssignee removes "assignee" edges to User entities.
func (tuo *TicketUpdateOne) RemoveAssignee(u ...*User) *TicketUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tuo.RemoveAssigneeIDs(ids...)
}

// ClearParent clears the "parent" edge to the Ticket entity.
func (tuo *TicketUpdateOne) ClearParent() *TicketUpdateOne {
	tuo.mutation.ClearParent()
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TicketUpdateOne) Select(field string, fields ...string) *TicketUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Ticket entity.
func (tuo *TicketUpdateOne) Save(ctx context.Context) (*Ticket, error) {
	var (
		err  error
		node *Ticket
	)
	if len(tuo.hooks) == 0 {
		if err = tuo.check(); err != nil {
			return nil, err
		}
		node, err = tuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TicketMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tuo.check(); err != nil {
				return nil, err
			}
			tuo.mutation = mutation
			node, err = tuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tuo.hooks) - 1; i >= 0; i-- {
			if tuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TicketUpdateOne) SaveX(ctx context.Context) *Ticket {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TicketUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TicketUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TicketUpdateOne) check() error {
	if v, ok := tuo.mutation.GetType(); ok {
		if err := ticket.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf("ent: validator failed for field \"type\": %w", err)}
		}
	}
	if v, ok := tuo.mutation.Status(); ok {
		if err := ticket.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf("ent: validator failed for field \"status\": %w", err)}
		}
	}
	return nil
}

func (tuo *TicketUpdateOne) sqlSave(ctx context.Context) (_node *Ticket, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   ticket.Table,
			Columns: ticket.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: ticket.FieldID,
			},
		},
	}
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Ticket.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, ticket.FieldID)
		for _, f := range fields {
			if !ticket.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != ticket.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: ticket.FieldName,
		})
	}
	if value, ok := tuo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: ticket.FieldDescription,
		})
	}
	if value, ok := tuo.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: ticket.FieldType,
		})
	}
	if value, ok := tuo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: ticket.FieldStatus,
		})
	}
	if tuo.mutation.ProjectCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ticket.ProjectTable,
			Columns: []string{ticket.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: project.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.ProjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ticket.ProjectTable,
			Columns: []string{ticket.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: project.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.ReporterCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   ticket.ReporterTable,
			Columns: ticket.ReporterPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedReporterIDs(); len(nodes) > 0 && !tuo.mutation.ReporterCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   ticket.ReporterTable,
			Columns: ticket.ReporterPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.ReporterIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   ticket.ReporterTable,
			Columns: ticket.ReporterPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.AssigneeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   ticket.AssigneeTable,
			Columns: ticket.AssigneePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedAssigneeIDs(); len(nodes) > 0 && !tuo.mutation.AssigneeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   ticket.AssigneeTable,
			Columns: ticket.AssigneePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.AssigneeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   ticket.AssigneeTable,
			Columns: ticket.AssigneePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   ticket.ParentTable,
			Columns: []string{ticket.ParentColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: ticket.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   ticket.ParentTable,
			Columns: []string{ticket.ParentColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: ticket.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Ticket{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{ticket.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
