package main

import "github.com/go-martini/martini"
//import "github.com/drlenux/go/app/Controllers/HomeController"
//import "github.com/codegangsta/inject"

func main() {
	m := martini.Classic()
	m.Get("/", func () string {
		return ""//HomeController.Run();
	})
	m.Run()
}
