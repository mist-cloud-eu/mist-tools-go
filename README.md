# Mist Tools Go

## Installation

```shell
$ go get github.com/mist-cloud-eu/mist-tools-go@<version>
```

Where `<version>` is a tag on this repository.

## Example Usage

The example below expects JSON on the form `{ "name": "Alan" }`,
and will reply with a JSON message (depending on the river).

```go
package main

import (
    "encoding/json"
    "fmt"
    "github.com/mist-cloud-eu/mist-tools-go"
)

// data I expect to receive
type Payload struct {
    Name string `json:"name"`
}

// data I expect to send
type Reply struct {
    Msg string `json:"msg"`
}

func main() {
    mist.MistService(map[string]func(string){
        "english-greeting": func(payload string) {
            var p Payload
            json.Unmarshal([]byte(payload), &p)
            reply := Reply{
                Msg: fmt.Sprintf("Hello, %s!", p.Name),
            }
            mist.PostToRapids[Reply]("reply", reply)
        },
        "spanish-greeting": func(payload string) {
            var p Payload
            json.Unmarshal([]byte(payload), &p)
            reply := Reply{
                Msg: fmt.Sprintf("Hola, %s!", p.Name),
            }
            mist.PostToRapids[Reply]("reply", reply)
        },
    })
}
```
