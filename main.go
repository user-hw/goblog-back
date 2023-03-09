package main

import (
	_ "goblog-end/dao"
	"goblog-end/router"
)

func main() {
	router.Start()
}
