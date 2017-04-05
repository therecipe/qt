//this file is just here to let moc and qtmoc generate some files

package quickflux

import (
	"github.com/therecipe/qt/core"
	_ "github.com/therecipe/qt/quick"
)

type stub struct {
	core.QObject
}

//go:generate /usr/local/Qt5.8.0/5.8/clang_64/bin/qmake quickflux.pro
//go:generate make mocables

//go:generate qtmoc
//go:generate go install
