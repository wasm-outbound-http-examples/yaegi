package main

import (
	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

func main() {
	interpreter := interp.New(interp.Options{})

	interpreter.Use(stdlib.Symbols)

	src := `
package main

import (
  "fmt"; "io"; "net/http"
)

func main() {
  resp, err := http.Get("https://httpbin.org/anything")
  if err != nil {
    panic(err)
  }
  defer resp.Body.Close()
  body, _ := io.ReadAll(resp.Body)
  fmt.Println(string(body))
}
`
	_, err := interpreter.Eval(src)
	if err != nil {
		panic(err)
	}
}
