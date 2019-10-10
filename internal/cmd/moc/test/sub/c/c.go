package c

import (
	"github.com/therecipe/qt/core"

	_ "github.com/therecipe/qt/internal/cmd/moc/test/sub/conf"
)

type StructSubGoC struct{}
type StructSubMocC struct{ core.QObject }
