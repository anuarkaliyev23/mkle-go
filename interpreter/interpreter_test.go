package interpreter

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	lua "github.com/yuin/gopher-lua"
	"testing"
)

func TestParser_Execute(t *testing.T) {
	t.Run("string result", func(t *testing.T) {
		parser := NewLuaInterpreter()
		state, err := parser.Execute(`s = "Hello " .. "World"`)
		defer state.Close()

		require.NoError(t, err)

		s := state.GetGlobal("s").String()
		assert.Equal(t, "Hello World", s)
	})

	t.Run("number result", func(t *testing.T) {
		parser := NewLuaInterpreter()
		state, err := parser.Execute(`a = 1 + 1`)
		require.NoError(t, err)
		a := state.GetGlobal("a").(lua.LNumber)
		require.Equal(t, 2.0, float64(a))
	})
}

func TestGopherLua(t *testing.T) {
	t.Run("read global string value", func(t *testing.T) {
		l := lua.NewState()
		defer l.Close()
		err := l.DoString(`s = "Hello " .. "World"`)
		require.NoError(t, err)
		value := l.GetGlobal("s").String()
		require.Equal(t, "Hello World", value)
	})

	t.Run("read global number (float) value", func(t *testing.T) {
		l := lua.NewState()
		defer l.Close()
		err := l.DoString(`a = 1 + 1`)
		require.NoError(t, err)
		value := l.GetGlobal("a")
		require.Equal(t, 2.0, float64(value.(lua.LNumber)))
	})

	t.Run("read boolean value", func(t *testing.T) {
		l := lua.NewState()
		defer l.Close()
		err := l.DoString(`a = true`)
		require.NoError(t, err)
		value := l.GetGlobal("a")
		require.Equal(t, lua.LTrue, value)
	})

	t.Run("custom function", func(t *testing.T) {
		luaFunctionDouble := func(state *lua.LState) int {
			lv := state.ToInt(1)          // get argument
			result := lua.LNumber(lv * 2) // calculate * 2
			state.Push(result)            // push on stack result
			return 1                      // number of results
		}

		state := lua.NewState()
		defer state.Close()

		state.SetGlobal("double", state.NewFunction(luaFunctionDouble))
		err := state.DoString(`a = double(20)`) // use our new function
		require.NoError(t, err)
		value := state.GetGlobal("a")
		require.Equal(t, 40.0, float64(value.(lua.LNumber)))
	})
}
