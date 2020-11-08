package state

// State

import "final/model"

type Context struct {
	User          *model.User
	IsAuthorized  bool
}

type State interface {
	Execute(c *Context)
}

type UnauthorizedState struct {
	User *model.User
}

func (UnauthorizedState) Execute(c *Context) {
	c.IsAuthorized = false
	c.User = nil
}

type AuthorizedState struct {
	User *model.User
}

func (a *AuthorizedState) Execute(c *Context) {
	c.IsAuthorized = true
	c.User = a.User
}

