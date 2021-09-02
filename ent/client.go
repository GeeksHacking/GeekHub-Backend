// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/geekshacking/geekhub-backend/ent/migrate"

	"github.com/geekshacking/geekhub-backend/ent/language"
	"github.com/geekshacking/geekhub-backend/ent/project"
	"github.com/geekshacking/geekhub-backend/ent/tag"
	"github.com/geekshacking/geekhub-backend/ent/ticket"
	"github.com/geekshacking/geekhub-backend/ent/user"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Language is the client for interacting with the Language builders.
	Language *LanguageClient
	// Project is the client for interacting with the Project builders.
	Project *ProjectClient
	// Tag is the client for interacting with the Tag builders.
	Tag *TagClient
	// Ticket is the client for interacting with the Ticket builders.
	Ticket *TicketClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Language = NewLanguageClient(c.config)
	c.Project = NewProjectClient(c.config)
	c.Tag = NewTagClient(c.config)
	c.Ticket = NewTicketClient(c.config)
	c.User = NewUserClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:      ctx,
		config:   cfg,
		Language: NewLanguageClient(cfg),
		Project:  NewProjectClient(cfg),
		Tag:      NewTagClient(cfg),
		Ticket:   NewTicketClient(cfg),
		User:     NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		config:   cfg,
		Language: NewLanguageClient(cfg),
		Project:  NewProjectClient(cfg),
		Tag:      NewTagClient(cfg),
		Ticket:   NewTicketClient(cfg),
		User:     NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Language.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Language.Use(hooks...)
	c.Project.Use(hooks...)
	c.Tag.Use(hooks...)
	c.Ticket.Use(hooks...)
	c.User.Use(hooks...)
}

// LanguageClient is a client for the Language schema.
type LanguageClient struct {
	config
}

// NewLanguageClient returns a client for the Language from the given config.
func NewLanguageClient(c config) *LanguageClient {
	return &LanguageClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `language.Hooks(f(g(h())))`.
func (c *LanguageClient) Use(hooks ...Hook) {
	c.hooks.Language = append(c.hooks.Language, hooks...)
}

// Create returns a create builder for Language.
func (c *LanguageClient) Create() *LanguageCreate {
	mutation := newLanguageMutation(c.config, OpCreate)
	return &LanguageCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Language entities.
func (c *LanguageClient) CreateBulk(builders ...*LanguageCreate) *LanguageCreateBulk {
	return &LanguageCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Language.
func (c *LanguageClient) Update() *LanguageUpdate {
	mutation := newLanguageMutation(c.config, OpUpdate)
	return &LanguageUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *LanguageClient) UpdateOne(l *Language) *LanguageUpdateOne {
	mutation := newLanguageMutation(c.config, OpUpdateOne, withLanguage(l))
	return &LanguageUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *LanguageClient) UpdateOneID(id int) *LanguageUpdateOne {
	mutation := newLanguageMutation(c.config, OpUpdateOne, withLanguageID(id))
	return &LanguageUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Language.
func (c *LanguageClient) Delete() *LanguageDelete {
	mutation := newLanguageMutation(c.config, OpDelete)
	return &LanguageDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *LanguageClient) DeleteOne(l *Language) *LanguageDeleteOne {
	return c.DeleteOneID(l.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *LanguageClient) DeleteOneID(id int) *LanguageDeleteOne {
	builder := c.Delete().Where(language.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &LanguageDeleteOne{builder}
}

// Query returns a query builder for Language.
func (c *LanguageClient) Query() *LanguageQuery {
	return &LanguageQuery{
		config: c.config,
	}
}

// Get returns a Language entity by its id.
func (c *LanguageClient) Get(ctx context.Context, id int) (*Language, error) {
	return c.Query().Where(language.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *LanguageClient) GetX(ctx context.Context, id int) *Language {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryProjects queries the projects edge of a Language.
func (c *LanguageClient) QueryProjects(l *Language) *ProjectQuery {
	query := &ProjectQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := l.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(language.Table, language.FieldID, id),
			sqlgraph.To(project.Table, project.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, language.ProjectsTable, language.ProjectsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(l.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *LanguageClient) Hooks() []Hook {
	return c.hooks.Language
}

// ProjectClient is a client for the Project schema.
type ProjectClient struct {
	config
}

// NewProjectClient returns a client for the Project from the given config.
func NewProjectClient(c config) *ProjectClient {
	return &ProjectClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `project.Hooks(f(g(h())))`.
func (c *ProjectClient) Use(hooks ...Hook) {
	c.hooks.Project = append(c.hooks.Project, hooks...)
}

// Create returns a create builder for Project.
func (c *ProjectClient) Create() *ProjectCreate {
	mutation := newProjectMutation(c.config, OpCreate)
	return &ProjectCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Project entities.
func (c *ProjectClient) CreateBulk(builders ...*ProjectCreate) *ProjectCreateBulk {
	return &ProjectCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Project.
func (c *ProjectClient) Update() *ProjectUpdate {
	mutation := newProjectMutation(c.config, OpUpdate)
	return &ProjectUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ProjectClient) UpdateOne(pr *Project) *ProjectUpdateOne {
	mutation := newProjectMutation(c.config, OpUpdateOne, withProject(pr))
	return &ProjectUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ProjectClient) UpdateOneID(id int) *ProjectUpdateOne {
	mutation := newProjectMutation(c.config, OpUpdateOne, withProjectID(id))
	return &ProjectUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Project.
func (c *ProjectClient) Delete() *ProjectDelete {
	mutation := newProjectMutation(c.config, OpDelete)
	return &ProjectDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ProjectClient) DeleteOne(pr *Project) *ProjectDeleteOne {
	return c.DeleteOneID(pr.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ProjectClient) DeleteOneID(id int) *ProjectDeleteOne {
	builder := c.Delete().Where(project.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ProjectDeleteOne{builder}
}

// Query returns a query builder for Project.
func (c *ProjectClient) Query() *ProjectQuery {
	return &ProjectQuery{
		config: c.config,
	}
}

// Get returns a Project entity by its id.
func (c *ProjectClient) Get(ctx context.Context, id int) (*Project, error) {
	return c.Query().Where(project.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ProjectClient) GetX(ctx context.Context, id int) *Project {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUsers queries the users edge of a Project.
func (c *ProjectClient) QueryUsers(pr *Project) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(project.Table, project.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, project.UsersTable, project.UsersPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(pr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryTags queries the tags edge of a Project.
func (c *ProjectClient) QueryTags(pr *Project) *TagQuery {
	query := &TagQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(project.Table, project.FieldID, id),
			sqlgraph.To(tag.Table, tag.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, project.TagsTable, project.TagsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(pr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryTickets queries the tickets edge of a Project.
func (c *ProjectClient) QueryTickets(pr *Project) *TicketQuery {
	query := &TicketQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(project.Table, project.FieldID, id),
			sqlgraph.To(ticket.Table, ticket.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, project.TicketsTable, project.TicketsColumn),
		)
		fromV = sqlgraph.Neighbors(pr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryLanguages queries the languages edge of a Project.
func (c *ProjectClient) QueryLanguages(pr *Project) *LanguageQuery {
	query := &LanguageQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(project.Table, project.FieldID, id),
			sqlgraph.To(language.Table, language.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, project.LanguagesTable, project.LanguagesPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(pr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryOwner queries the owner edge of a Project.
func (c *ProjectClient) QueryOwner(pr *Project) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(project.Table, project.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, project.OwnerTable, project.OwnerColumn),
		)
		fromV = sqlgraph.Neighbors(pr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ProjectClient) Hooks() []Hook {
	return c.hooks.Project
}

// TagClient is a client for the Tag schema.
type TagClient struct {
	config
}

// NewTagClient returns a client for the Tag from the given config.
func NewTagClient(c config) *TagClient {
	return &TagClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `tag.Hooks(f(g(h())))`.
func (c *TagClient) Use(hooks ...Hook) {
	c.hooks.Tag = append(c.hooks.Tag, hooks...)
}

// Create returns a create builder for Tag.
func (c *TagClient) Create() *TagCreate {
	mutation := newTagMutation(c.config, OpCreate)
	return &TagCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Tag entities.
func (c *TagClient) CreateBulk(builders ...*TagCreate) *TagCreateBulk {
	return &TagCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Tag.
func (c *TagClient) Update() *TagUpdate {
	mutation := newTagMutation(c.config, OpUpdate)
	return &TagUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TagClient) UpdateOne(t *Tag) *TagUpdateOne {
	mutation := newTagMutation(c.config, OpUpdateOne, withTag(t))
	return &TagUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TagClient) UpdateOneID(id int) *TagUpdateOne {
	mutation := newTagMutation(c.config, OpUpdateOne, withTagID(id))
	return &TagUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Tag.
func (c *TagClient) Delete() *TagDelete {
	mutation := newTagMutation(c.config, OpDelete)
	return &TagDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *TagClient) DeleteOne(t *Tag) *TagDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *TagClient) DeleteOneID(id int) *TagDeleteOne {
	builder := c.Delete().Where(tag.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TagDeleteOne{builder}
}

// Query returns a query builder for Tag.
func (c *TagClient) Query() *TagQuery {
	return &TagQuery{
		config: c.config,
	}
}

// Get returns a Tag entity by its id.
func (c *TagClient) Get(ctx context.Context, id int) (*Tag, error) {
	return c.Query().Where(tag.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TagClient) GetX(ctx context.Context, id int) *Tag {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryProjects queries the projects edge of a Tag.
func (c *TagClient) QueryProjects(t *Tag) *ProjectQuery {
	query := &ProjectQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(tag.Table, tag.FieldID, id),
			sqlgraph.To(project.Table, project.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, tag.ProjectsTable, tag.ProjectsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *TagClient) Hooks() []Hook {
	return c.hooks.Tag
}

// TicketClient is a client for the Ticket schema.
type TicketClient struct {
	config
}

// NewTicketClient returns a client for the Ticket from the given config.
func NewTicketClient(c config) *TicketClient {
	return &TicketClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `ticket.Hooks(f(g(h())))`.
func (c *TicketClient) Use(hooks ...Hook) {
	c.hooks.Ticket = append(c.hooks.Ticket, hooks...)
}

// Create returns a create builder for Ticket.
func (c *TicketClient) Create() *TicketCreate {
	mutation := newTicketMutation(c.config, OpCreate)
	return &TicketCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Ticket entities.
func (c *TicketClient) CreateBulk(builders ...*TicketCreate) *TicketCreateBulk {
	return &TicketCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Ticket.
func (c *TicketClient) Update() *TicketUpdate {
	mutation := newTicketMutation(c.config, OpUpdate)
	return &TicketUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TicketClient) UpdateOne(t *Ticket) *TicketUpdateOne {
	mutation := newTicketMutation(c.config, OpUpdateOne, withTicket(t))
	return &TicketUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TicketClient) UpdateOneID(id int) *TicketUpdateOne {
	mutation := newTicketMutation(c.config, OpUpdateOne, withTicketID(id))
	return &TicketUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Ticket.
func (c *TicketClient) Delete() *TicketDelete {
	mutation := newTicketMutation(c.config, OpDelete)
	return &TicketDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *TicketClient) DeleteOne(t *Ticket) *TicketDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *TicketClient) DeleteOneID(id int) *TicketDeleteOne {
	builder := c.Delete().Where(ticket.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TicketDeleteOne{builder}
}

// Query returns a query builder for Ticket.
func (c *TicketClient) Query() *TicketQuery {
	return &TicketQuery{
		config: c.config,
	}
}

// Get returns a Ticket entity by its id.
func (c *TicketClient) Get(ctx context.Context, id int) (*Ticket, error) {
	return c.Query().Where(ticket.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TicketClient) GetX(ctx context.Context, id int) *Ticket {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryProject queries the project edge of a Ticket.
func (c *TicketClient) QueryProject(t *Ticket) *ProjectQuery {
	query := &ProjectQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(ticket.Table, ticket.FieldID, id),
			sqlgraph.To(project.Table, project.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ticket.ProjectTable, ticket.ProjectColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryReporter queries the reporter edge of a Ticket.
func (c *TicketClient) QueryReporter(t *Ticket) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(ticket.Table, ticket.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, ticket.ReporterTable, ticket.ReporterPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryAssignee queries the assignee edge of a Ticket.
func (c *TicketClient) QueryAssignee(t *Ticket) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(ticket.Table, ticket.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, ticket.AssigneeTable, ticket.AssigneePrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryParent queries the parent edge of a Ticket.
func (c *TicketClient) QueryParent(t *Ticket) *TicketQuery {
	query := &TicketQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(ticket.Table, ticket.FieldID, id),
			sqlgraph.To(ticket.Table, ticket.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, ticket.ParentTable, ticket.ParentColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *TicketClient) Hooks() []Hook {
	return c.hooks.Ticket
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Create returns a create builder for User.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryProjects queries the projects edge of a User.
func (c *UserClient) QueryProjects(u *User) *ProjectQuery {
	query := &ProjectQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(project.Table, project.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, user.ProjectsTable, user.ProjectsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryReportedTickets queries the reported_tickets edge of a User.
func (c *UserClient) QueryReportedTickets(u *User) *TicketQuery {
	query := &TicketQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(ticket.Table, ticket.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, user.ReportedTicketsTable, user.ReportedTicketsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryAssignedTickets queries the assigned_tickets edge of a User.
func (c *UserClient) QueryAssignedTickets(u *User) *TicketQuery {
	query := &TicketQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(ticket.Table, ticket.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, user.AssignedTicketsTable, user.AssignedTicketsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}