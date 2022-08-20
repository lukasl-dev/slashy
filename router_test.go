package slashy

import (
	"github.com/bwmarrin/discordgo"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

type runner struct {
	mock.Mock
}

var _ Runner = (*runner)(nil)

func (r *runner) Run(ctx *Context, resp *Response) error {
	return r.Called(ctx, resp).Error(0)
}

type completableRunner struct {
	mock.Mock
}

var (
	_ Runner        = (*completableRunner)(nil)
	_ AutoCompleter = (*completableRunner)(nil)
)

func (c *completableRunner) Run(ctx *Context, resp *Response) error {
	return c.Called(ctx, resp).Error(0)
}

func (c *completableRunner) AutoComplete(ctx *Context) []*discordgo.ApplicationCommandOptionChoice {
	return c.Called(ctx).Get(0).([]*discordgo.ApplicationCommandOptionChoice)
}

type errorResponder struct {
	mock.Mock
}

var _ ErrorResponder = (*errorResponder)(nil)

func (e *errorResponder) RespondError(ctx *Context, resp *Response, err error) {
	e.Called(ctx, resp, err)
}

func TestNewRouter_WithoutErrorResponder(t *testing.T) {
	r := require.New(t)

	router := NewRouter(nil)

	r.NotNil(router)
	r.Empty(router.commands)
	r.NotNil(router.opts.ErrorResponder)
}

func TestNewRouter_WithCustomErrorResponder(t *testing.T) {
	errResp := new(errorResponder)

	r := require.New(t)

	router := NewRouter(&RouterOptions{ErrorResponder: errResp})

	r.NotNil(router)
	r.Empty(router.commands)
	r.Equal(router.opts.ErrorResponder, errResp)
}

func TestRouter_Bind_WithoutName(t *testing.T) {
	r := require.New(t)

	router := NewRouter(nil)

	r.Panics(func() {
		router.Bind("", &CommandRoute{Runner: new(runner)})
	})
	r.Empty(router.commands)
}

func TestRouter_Bind_WithoutCommand(t *testing.T) {
	r := require.New(t)

	router := NewRouter(nil)

	r.Panics(func() {
		router.Bind("test", nil)
	})
	r.Empty(router.commands)
}

func TestRouter_Bind_WithoutRunner(t *testing.T) {
	r := require.New(t)

	router := NewRouter(nil)

	r.Panics(func() {
		router.Bind("test", new(CommandRoute))
	})
	r.Empty(router.commands)
}

func TestRouter_Bind_Valid(t *testing.T) {
	const name = "TEst"

	r := require.New(t)

	router := NewRouter(nil)

	r.NotPanics(func() {
		router.Bind("test", &CommandRoute{Runner: new(runner)})
	})
	r.Contains(router.commands, strings.ToLower(name))
}

func TestRouter_AutoBind_WithoutName(t *testing.T) {
	r := require.New(t)

	router := NewRouter(nil)

	r.Panics(func() {
		router.AutoBind("", new(runner))
	})
}

func TestRouter_AutoBind_WithoutRunner(t *testing.T) {
	r := require.New(t)

	router := NewRouter(nil)

	r.Panics(func() {
		router.AutoBind("test", nil)
	})
	r.Empty(router.commands)
}

func TestRouter_AutoBind_OnlyRunner(t *testing.T) {
	const name = "test"

	r := require.New(t)

	router := NewRouter(nil)

	r.NotPanics(func() {
		router.AutoBind(name, new(runner))
	})
	r.Contains(router.commands, strings.ToLower(name))
	r.NotNil(router.commands[strings.ToLower(name)])
	r.NotNil(router.commands[strings.ToLower(name)].Runner)
}

func TestRouter_AutoBind_WithCompleter(t *testing.T) {
	const name = "test"

	r := require.New(t)

	router := NewRouter(nil)

	r.NotPanics(func() {
		router.AutoBind(name, new(completableRunner))
	})
	r.Contains(router.commands, strings.ToLower(name))
	r.NotNil(router.commands[strings.ToLower(name)])
	r.NotNil(router.commands[strings.ToLower(name)].Runner)
	r.NotNil(router.commands[strings.ToLower(name)].AutoCompleter)
}
