# ENTRY
Package for checking whether an element or an array of elements is included in the checked set
## Examples

```go
package main

import (
	"fmt"
	"github.com/AndrivA89/entry"
)

func main() {
	fmt.Println(entry.IsContains(3, []int{13, 14, 3}))
	// true
	
	fmt.Println(entry.IsContains("one", []string{"one", "two", "six"}))
	// true

	fmt.Println(entry.IsContains([]int{3, 4, 5}, []uint32{13, 14, 3, 4, 5}))
	// true

	fmt.Println(entry.IsContains([]int{3, 4, 6}, []string{"13", "14", "3", "4", "5"}))
	// false

	fmt.Println(entry.IsContains([]int{3, 4, 5}, []string{"13", "14", "3", "4", "5"}))
	// true
}
```