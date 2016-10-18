// Author: Ian McLoughlin
// Date: October 6th, 2016
// Adapted from: https://go-macaron.com

package main

import "gopkg.in/macaron.v1"

// Function taken from: github.com/golang/example
func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func main() {
	m := macaron.Classic()
	m.Use(macaron.Renderer())

	m.Get("/bookaroom", func() string {
		return "Book a room here!"
	})

	m.Get("/reverse/:name", func(ctx *macaron.Context) {
		// Adapted from: https://go-macaron.com/docs/middlewares/templating
		ctx.Data["Name"] = reverse(ctx.Params(":name"))
		ctx.HTML(200, "hello")
	})

	m.Run(4001)
}
