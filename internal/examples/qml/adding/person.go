package main

import (
	"github.com/bluszcz/cutego/core"
)

type Person struct {
	core.QObject

	_ string `property:"name"`
	_ int    `property:"shoeSize"`
}
