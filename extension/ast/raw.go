package ast

import (
	gast "github.com/yuin/goldmark/ast"
)

type RawBlock struct {
	gast.BaseBlock
}

// Dump implements Node.Dump.
func (n *RawBlock) Dump(source []byte, level int) {
	gast.DumpHelper(n, source, level, nil, nil)
}

// KindRawBlock is a NodeKind of the KindRawBlock node.
var KindRawBlock = gast.NewNodeKind("RawBlock")

// Kind implements Node.Kind.
func (n *RawBlock) Kind() gast.NodeKind {
	return KindRawBlock
}

// NewRawBlock returns a new RawBlock node.
func NewRawBlock() *RawBlock {
	return &RawBlock{}
}

// A RawInline struct represents a checkbox of a task list.
type RawInline struct {
	gast.BaseInline
	Content []byte
}

// Dump implements Node.Dump.
func (n *RawInline) Dump(source []byte, level int) {
	gast.DumpHelper(n, source, level, nil, nil)
}

// KindRawInline is a NodeKind of the RawInline node.
var KindRawInline = gast.NewNodeKind("RawInline")

// Kind implements Node.Kind.
func (n *RawInline) Kind() gast.NodeKind {
	return KindRawInline
}

// NewRawInline returns a new RawInline node.
func NewRawInline(content []byte) *RawInline {
	return &RawInline{
		Content: content,
	}
}
