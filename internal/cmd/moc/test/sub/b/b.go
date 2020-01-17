package b

import (
	"github.com/therecipe/qt/core"

	_ "github.com/therecipe/qt/internal/cmd/moc/test/sub/conf"
)

type StructSubGoB struct{}
type StructSubMocB struct{ core.QObject }
