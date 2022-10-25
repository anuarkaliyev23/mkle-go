package markdown

import "github.com/yuin/goldmark/ast"

var KindLuaInline = ast.NewNodeKind("LuaInline")

type luaInline struct {
	ast.BaseInline
}

func (r *luaInline) IsRaw() bool {
	return false
}

func (r *luaInline) Dump(source []byte, level int) {
	ast.DumpHelper(r, source, level, nil, nil)
}

func (r *luaInline) Kind() ast.NodeKind {
	return KindLuaInline
}

func NewLuaInline() ast.Node {
	return &luaInline{
		BaseInline: ast.BaseInline{
			BaseNode: ast.BaseNode{},
		},
	}
}
