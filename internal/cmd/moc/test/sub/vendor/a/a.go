package a

import (
	"github.com/therecipe/qt/core"

	_ "github.com/therecipe/qt/internal/cmd/moc/test/sub/conf"
)

type StructSubGoA struct{}
type StructSubMocA struct{ core.QObject }
