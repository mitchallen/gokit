# github.com/mitchallen/gokit

## Usage

### initialize your module

```
$ go mod init example.com/my-gokit-demo
```

### Get thie gokit module

Note that you need to include the **v** in the version tag.

```
$ go get github.com/mitchallen/gokit@v0.3.0
```

```go
package main

import (
	"fmt"

	"github.com/mitchallen/gokit"
)

func main() {
	fmt.Println(gokit.Info())
}
```

## Init

How this repo was intialized as a **go** (golang) module

```
$ go mod init github.com/mitchallen/gokit
```

## Testing

```
$ go test
```

```
$ go test ./calc
```

```
$ cd (package)
$ go test
```

## Tagging

```
$ git tag v0.1.0
```

## Reference

* https://eli.thegreenplace.net/2019/simple-go-project-layout-with-modules/


## Package

* https://pkg.go.dev/search?q=mitchallen
