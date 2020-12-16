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
	bar.NewBar(0, 100)
	for i := 0; i <= 100; i++ {
		time.Sleep(100 * time.Millisecond)
		bar.Increment()
		bar.Play()
	}
	bar.Finish()
}

```
