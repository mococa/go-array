# Go Array

### Description

As many people know, golang lacks some slice methods compared to JavaScript or Dart, for example. So coming from these languages can be a little tricky and sometimes frustating because you'll write the same code many times. And the goal of this library is to mock the way JavaScript does things, using Generic types introducted on Go v1.18

### Installation

```bash
go get github.com/mococa/go-array
```

### Usage

```go
import (
    "fmt"
    "strings"

    "github.com/mococa/go-array"
)

type Message struct {
    body string
}

func main() {
    // If there is no default slice, it needs the type inference: array := NewGoArray[Message]()
    array := NewGoArray([]Message{{body: "I'm an optional default slice"}})

    array.Push(Message{ body: "Hello" })
    array.Push(Message{ body: "World" })

    array.Find(func(item Message) bool {
		return item.body == "World"
	}) // Returns array.data[2]

    array.FindIndex(func(item Message) bool {
		return item.body == "World"
	}) // 2

    array.At(2) // Message{ body: "World" }

    array.RemoveAt(2) // Removes Message{ body: "World" } from array

    array.RemoveWhere(func(item Message) bool {
		return strings.Contains(item.body, "l")
	}) // Removes all items that the body contains "l"

    array.Filter(func(item Message) bool {
		return item.body == "World"
	}) // Returns all item that the body is equals to "World"

    array.Map(func(item Message) any {
		return fmt.Sprintf("The message body is: %s", item.body)
	}) // Returns a modified version of the array
}
```
