package main

import (
	"github.com/StarAurryon/qt/core"
)

type Person struct {
	core.QObject

	_ string `property:"name"`
	_ int    `property:"shoeSize"`
}
