# Calendar for Golang

[![Build Status](https://travis-ci.org/binzume/go-calendar.svg)](https://travis-ci.org/binzume/go-calendar)
[![codecov](https://codecov.io/gh/binzume/go-calendar/branch/master/graph/badge.svg)](https://codecov.io/gh/binzume/go-calendar)

Output HTML and Markdown calendar.

## Usage

```go
package main

import (
	"fmt"
	"github.com/binzume/go-calendar"
)

func main() {
	fmt.Println(calendar.NewCalendar().Markdown())
}
```

### Output example

```
| Sun | Mon | Tue | Wed | Thu | Fri | Sat |
|----:|----:|----:|----:|----:|----:|----:|
|     |     |     |   1 |   2 |   3 |   4 |
|   5 |   6 |   7 |   8 |   9 |  10 |  11 |
|  12 |  13 |  14 |  15 |  16 |  17 |  18 |
|  19 |  20 |  21 |  22 |  23 |  24 |  25 |
|  26 |  27 |  28 |  29 |  30 |     |     |
```

| Sun | Mon | Tue | Wed | Thu | Fri | Sat |
|----:|----:|----:|----:|----:|----:|----:|
|     |     |     |   1 |   2 |   3 |   4 |
|   5 |   6 |   7 |   8 |   9 |  10 |  11 |
|  12 |  13 |  14 |  15 |  16 |  17 |  18 |
|  19 |  20 |  21 |  22 |  23 |  24 |  25 |
|  26 |  27 |  28 |  29 |  30 |     |     |

# License

These codes are licensed under CC0.

[![CC0](http://i.creativecommons.org/p/zero/1.0/88x31.png "CC0")](http://creativecommons.org/publicdomain/zero/1.0/deed.ja)
