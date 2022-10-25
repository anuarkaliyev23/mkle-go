package luaparser

import (
	"github.com/Shopify/go-lua"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

func TestParser_Execute(t *testing.T) {
	t.Run("test lua", func(t *testing.T) {
		state := lua.NewState()
		lua.OpenLibraries(state)
		err := lua.DoString(state, "a = 1 + 1")
		require.NoError(t, err)
	})

	t.Run("string result", func(t *testing.T) {
		parser := New()
		state := parser.InitState()
		state, err := parser.Execute(`s = "Hello " .. "World"`, state)
		require.NoError(t, err)
		state.PushGlobalTable()
		state.Pop(2)
		for state.Next(-2) {
			key := lua.CheckString(state, -2)
			val := lua.CheckString(state, -1)
			log.Printf("key = %s val = %s", key, val)
			state.Pop(1) // Remove val, but need key for the next iter.
		}
	})

}
