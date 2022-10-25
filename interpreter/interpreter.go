package interpreter

import (
	"github.com/anuarkaliyev23/mkle-go/logger"
	lua "github.com/yuin/gopher-lua"
)

type LuaInterpreter struct {
	logger logger.Logger
}

func (r LuaInterpreter) Execute(luaString string) (*lua.LState, error) {
	r.logger.Debug("executing lua script", map[string]any{
		"script": luaString,
	})

	state := lua.NewState()
	if err := state.DoString(luaString); err != nil {
		return nil, err
	}

	return state, nil
}

func (r LuaInterpreter) ExecuteAndCloseState(luaString string) error {
	r.logger.Debug("executing lua script", map[string]any{
		"script": luaString,
	})

	state := lua.NewState()
	defer state.Close()

	if err := state.DoString(luaString); err != nil {
		return err
	}

	return nil
}

func NewLuaInterpreter() LuaInterpreter {
	return LuaInterpreter{
		logger: logger.NewZerologLogger(),
	}
}
