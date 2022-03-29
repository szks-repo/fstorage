# fstorage
golang file storage utilities

## create file with directory

```go
package main

import (
	"fmt"
	"github.com/szks-repo/fstorage"
	"log"
	"strings"
)

func main() {
	sc, _ := fstorage2.New("/path/to/storage")
	if err := sc.SaveAll("test/test1.txt", strings.NewReader("test"), nil); err != nil {
		log.Fatal(err)
	}

	f, err := sc.Get("/test/test1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	fmt.Println(f.String())
}
```