# goldmark 扩展

## raw

```go
package extension

type RawOption struct {
    IsBlock               bool   `json:"is_block"`                 // 是否为块级
    Tag                   string `json:"tag"`                      // 捕捉后生成的 html 标签
    ClassName             string `json:"class_name"`               // 生成的标签的 class 值
    KeepMark              bool   `json:"keep_mark"`                // 是否保留前后标识
    PrefixMark            string `json:"prefix_mark"`              // 开头匹配
    SuffixMark            string `json:"suffix_mark"`              // 结尾匹配
    Trigger               string `json:"trigger"`                  // 触发字节
    ParserPriority        *int   `json:"parser_priority"`          // 注册的解析器优先级
    RendererPriority      *int   `json:"renderer_priority"`        // 注册的渲染器优先级
    CanInterruptParagraph bool   `json:"can_interrupt_paragraph"`  // 块级解析器是否可以中断段落
    CanAcceptIndentedLine bool   `json:"can_accept_indented_line"` // 块级解析器是否可以渲染缩进线
}
```
