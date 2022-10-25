package luaparser

import "github.com/Shopify/go-lua"

type Parser struct {
}

func (r Parser) InitState() *lua.State {
	l := lua.NewState()
	lua.OpenLibraries(l)
	return l
}

func (r Parser) Execute(luaScript string, state *lua.State) (*lua.State, error) {
	if err := lua.DoString(state, luaScript); err != nil {
		return nil, err
	}

	return state, nil
}

func New() Parser {
	return Parser{}
}
