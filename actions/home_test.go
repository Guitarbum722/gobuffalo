package actions_test

import (
	"testing"

	"github.com/gobuffalo/gobuffalo/actions"
	"github.com/markbates/willie"
	"github.com/stretchr/testify/require"
)

func Test_HomeHandler(t *testing.T) {
	r := require.New(t)

	w := willie.New(actions.App())
	res := w.Request("/").Get()

	r.Equal(302, res.Code)
	r.Equal("/docs/getting-started", res.Location())
	// r.Equal(200, res.Code)
	// r.Contains(res.Body.String(), "Welcome to Buffalo!")
}
