package c

import (
	"github.com/akiyosi/qt/core"

	_ "github.com/akiyosi/qt/internal/cmd/moc/test/sub/conf"
)

type StructSubGoC struct{}
type StructSubMocC struct{ core.QObject }
