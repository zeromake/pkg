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
    textconv.CamelCase("camel case") // `camelCase`
    textconv.PascalCase("pascal case") // `PascalCase`
    textconv.CapitalCase("capital case") // `Capital Case`
    textconv.PathCase("path case") // `path/case`
    textconv.SnakeCase("snake case") // `snake_case`
    textconv.NoCase("no case") // `no case`
}
```
