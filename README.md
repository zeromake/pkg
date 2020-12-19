# pkg
golang library

## textconv
> word conv

``` go
package main

import (
    "github.com/zeromake/pkg/textconv"
)

func main() {
    textconv.CamelCase("camel case")        // `camelCase`
    textconv.PascalCase("pascal case")      // `PascalCase`
    textconv.CapitalCase("capital case")    // `Capital Case`
    textconv.HeaderCase("header case")      // `Header-Case`
    textconv.TitleCase("title case")        // `Title Case`
    textconv.PathCase("path case")          // `path/case`
    textconv.ParamCase("param case")        // `param-case`
    textconv.SnakeCase("snake case")        // `snake_case`
    textconv.DotCase("dot case")            // `dot.case`
    textconv.NoCase("no case")              // `no case`
    textconv.ConstantCase("constant case")  // `CONSTANT_CASE`
    textconv.SentenceCase("sentence case")  // `SentenceCase`
}
```
