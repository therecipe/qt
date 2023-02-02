package main

import (
	"os"
	"time"

	"github.com/bluszcz/cutego/core"
	"github.com/bluszcz/cutego/gui"
	"github.com/bluszcz/cutego/quick"
)

func main() {
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	gui.NewQGuiApplication(len(os.Args), os.Args)

	var view = quick.NewQQuickView(nil)

	var model = NewPersonModel(nil)

	var p = NewPerson(nil)
	p.SetFirstName("john")
	p.SetLastName("doe")
	model.SetPeople([]*Person{p})

	view.RootContext().SetContextProperty("PersonModel", model)

	view.SetSource(core.NewQUrl3("qrc:/qml/main.qml", 0))
	view.SetResizeMode(quick.QQuickView__SizeRootObjectToView)
	view.Show()

	go func() {

		time.Sleep(5 * time.Second)

		//add person
		for i := 0; i < 3; i++ {
			var p = NewPerson(nil)
			p.MoveToThread(view.Engine().Thread())
			p.SetFirstName("hello")
			p.SetLastName("world")
			model.AddPerson(p)
		}

		//edit person
		model.EditPerson(1, "bob", "")
		model.EditPerson(3, "", "john")

		//remove person
		model.RemovePerson(2)

	}()

	gui.QGuiApplication_Exec()
}
