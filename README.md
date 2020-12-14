# bar
process bar

# Usage
```go
package main

import (
	"time"

	"github.com/ruyewuxin/bar"
)

func main() {
	var bar bar.Bar
	bar.NewOption(0, 100)
	for i := 0; i <= 100; i++ {
		time.Sleep(100 * time.Millisecond)
		bar.Play(i)
	}
	bar.Finish()
}

```
