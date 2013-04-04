package main

import (
	"fmt"
	"github.com/hoisie/web"
)

func hello(ctx *web.Context, val string) (html string) {
	for k, v := range ctx.Params {
		html += fmt.Sprintf("%v=%v", k, v)
	}
	return
}

func main() {
	web.Get("/(.*)", hello)
	web.Run("0.0.0.0:9999")
}
