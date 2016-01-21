package controllers

import (
	"net/http"

	"github.com/alkchr/pongoal"

	"github.com/colegion/contrib/controllers/requests"
	"github.com/colegion/contrib/controllers/sessions"
)

// Controllers is a struct that should be embedded into every controller
// of your app to make methods and fields provided by standard controllers available.
type Controllers struct {
	*requests.Requests
	*pongoal.Pongo2
	*sessions.Sessions
}

// Before is a magic action that is executed on every request
// before any other action.
//
// Only structures with at least one action are treated as controllers.
// So, do not delete this method.
func (c *Controllers) Before() http.Handler {
	return nil
}