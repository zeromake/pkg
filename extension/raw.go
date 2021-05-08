package extension

import (
	"bytes"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
	east "github.com/zeromake/pkg/goldmark/ast"
	"github.com/zeromake/pkg/strutil"
)

type RawOption struct {
	IsBlock               bool   `json:"is_block"`                 // 是否为块级
	Tag                   string `json:"tag"`                      // 捕捉后生成的 html 标签
	ClassName             string `json:"class_name"`               // 生成的标签的 class 值
	KeepMark              bool   `json:"keep_mark"`                // 是否保留匹配用标识
	PrefixMark            string `json:"prefix_mark"`              // 开头匹配
	SuffixMark            string `json:"suffix_mark"`              // 结尾匹配
	Trigger               string `json:"trigger"`                  // 触发字节
	ParserPriority        *int   `json:"parser_priority"`          // 注册的解析器优先级
	RendererPriority      *int   `json:"renderer_priority"`        // 注册的渲染器优先级
	CanInterruptParagraph bool   `json:"can_interrupt_paragraph"`  // 块级解析器是否可以中断段落
	CanAcceptIndentedLine bool   `json:"can_accept_indented_line"` // 块级解析器是否可以渲染缩进线
}

type RawBlockParser struct {
	Option RawOption
}

func (s *RawBlockParser) Open(
	_ ast.Node,
	reader text.Reader,
	_ parser.Context,
) (ast.Node, parser.State) {
	line, segment := reader.PeekLine()
	if !bytes.HasPrefix(
		line,
		strutil.StringToBytes(s.Option.PrefixMark),
	) {
		return nil, parser.NoChildren
	}
	node := east.NewRawBlock()
	if s.Option.KeepMark {
		node.Lines().Append(segment)
	}
	reader.Advance(segment.Len() - 1)
	return node, parser.NoChildren
}

func (s *RawBlockParser) Continue(
	node ast.Node,
	reader text.Reader,
	_ parser.Context,
) parser.State {
	line, segment := reader.PeekLine()
	if bytes.HasPrefix(
		line,
		strutil.StringToBytes(s.Option.SuffixMark),
	) {
		if s.Option.KeepMark {
			node.Lines().Append(segment)
		}
		reader.Advance(segment.Len())
		return parser.Close
	}
	node.Lines().Append(segment)
	reader.Advance(segment.Len() - 1)
	return parser.Continue | parser.NoChildren
}

func (s *RawBlockParser) Close(_ ast.Node, _ text.Reader, _ parser.Context) {}

func (s *RawBlockParser) CanInterruptParagraph() bool {
	return s.Option.CanInterruptParagraph
}

func (s *RawBlockParser) CanAcceptIndentedLine() bool {
	return s.Option.CanAcceptIndentedLine
}

func (s *RawBlockParser) Trigger() []byte {
	return []byte(s.Option.Trigger)
}

func NewRawBlockParser(option RawOption) parser.BlockParser {
	return &RawBlockParser{
		Option: option,
	}
}

type RawInlineParser struct {
	Option RawOption
}

func (r *RawInlineParser) Trigger() []byte {
	return strutil.StringToBytes(r.Option.Trigger)
}

func (r *RawInlineParser) Parse(_ ast.Node, reader text.Reader, _ parser.Context) ast.Node {
	line, _ := reader.PeekLine()
	prefixCount := len(r.Option.PrefixMark)
	if !bytes.HasPrefix(line, strutil.StringToBytes(r.Option.PrefixMark)) {
		return nil
	}
	suffixCount := len(r.Option.SuffixMark)
	index := bytes.Index(line[prefixCount:], strutil.StringToBytes(r.Option.SuffixMark))
	if index > suffixCount {
		index += suffixCount
		if line[index-suffixCount] != '\\' {
			reader.Advance(index + suffixCount)
			start := prefixCount
			end := index
			if r.Option.KeepMark {
				start = 0
				end += suffixCount
			}
			node := east.NewRawInline(line[start:end])
			return node
		}
	}
	return nil
}

func NewRawInlineParser(option RawOption) parser.InlineParser {
	return &RawInlineParser{
		Option: option,
	}
}

type RawBlockHTMLRenderer struct {
	html.Config
	Option RawOption
}

func (t *RawBlockHTMLRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(east.KindRawBlock, func(w util.BufWriter, source []byte, n ast.Node, entering bool) (ast.WalkStatus, error) {
		if entering {
			_ = w.WriteByte('<')
			_, _ = w.WriteString(t.Option.Tag)
			if t.Option.ClassName != "" {
				_, _ = w.WriteString(" class=\"")
				_, _ = w.WriteString(t.Option.ClassName)
				_ = w.WriteByte('"')
			}
			_, _ = w.WriteString(">\n")
		} else {
			_, _ = w.WriteString("</")
			_, _ = w.WriteString(t.Option.Tag)
			_ = w.WriteByte('>')
		}
		return ast.WalkContinue, nil
	})
}

func NewRawBlockHTMLRenderer(option RawOption) renderer.NodeRenderer {
	return &RawBlockHTMLRenderer{
		Option: option,
	}
}

type RawInlineHTMLRenderer struct {
	html.Config
	Option RawOption
}

func (r *RawInlineHTMLRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(east.KindRawInline, func(w util.BufWriter, source []byte, n ast.Node, entering bool) (ast.WalkStatus, error) {
		if entering {
			_ = w.WriteByte('<')
			_, _ = w.WriteString(r.Option.Tag)
			if r.Option.ClassName != "" {
				_, _ = w.WriteString(" class=\"")
				_, _ = w.WriteString(r.Option.ClassName)
				_ = w.WriteByte('"')
			}
			_ = w.WriteByte('>')
			_, _ = w.Write(n.(*east.RawInline).Content)
		} else {
			_, _ = w.WriteString("</")
			_, _ = w.WriteString(r.Option.Tag)
			_ = w.WriteByte('>')
		}
		return ast.WalkContinue, nil
	})
}

func NewRawInlineHTMLRenderer(option RawOption) renderer.NodeRenderer {
	return &RawInlineHTMLRenderer{
		Option: option,
	}
}

type RawBlock struct {
	Option RawOption
}

var DefaultPriority = 8999

func (e *RawBlock) Extend(m goldmark.Markdown) {
	if e.Option.Tag == "" {
		return
	}
	if e.Option.RendererPriority == nil {
		e.Option.RendererPriority = &DefaultPriority
	}
	if e.Option.ParserPriority == nil {
		e.Option.ParserPriority = &DefaultPriority
	}
	m.Parser().AddOptions(
		parser.WithBlockParsers(
			util.Prioritized(
				NewRawBlockParser(e.Option),
				*e.Option.ParserPriority,
			),
		),
	)
	m.Renderer().AddOptions(renderer.WithNodeRenderers(
		util.Prioritized(
			NewRawBlockHTMLRenderer(e.Option),
			*e.Option.RendererPriority,
		),
	))
}

type RawInline struct {
	Option RawOption
}

func (e *RawInline) Extend(m goldmark.Markdown) {
	if e.Option.Tag == "" {
		return
	}
	if e.Option.RendererPriority == nil {
		e.Option.RendererPriority = &DefaultPriority
	}
	if e.Option.ParserPriority == nil {
		e.Option.ParserPriority = &DefaultPriority
	}
	m.Parser().AddOptions(
		parser.WithInlineParsers(
			util.Prioritized(
				NewRawInlineParser(e.Option),
				*e.Option.ParserPriority,
			),
		),
	)
	m.Renderer().AddOptions(renderer.WithNodeRenderers(
		util.Prioritized(
			NewRawInlineHTMLRenderer(e.Option),
			*e.Option.RendererPriority,
		),
	))
}

func NewRawExtension(option RawOption) goldmark.Extender {
	if option.IsBlock {
		return &RawBlock{
			Option: option,
		}
	} else {
		return &RawInline{
			Option: option,
		}
	}
}
