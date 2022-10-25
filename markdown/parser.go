package markdown

import (
	"github.com/anuarkaliyev23/mkle-go/logger"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

type luaMdParser struct {
	logger logger.Logger
}

// Trigger format: `[lua]:`
func (r luaMdParser) Trigger() []byte {
	return []byte{'['}
}

func (r luaMdParser) Parse(parent ast.Node, block text.Reader, pc parser.Context) ast.Node {
	//todo
	return NewLuaInline()
}
