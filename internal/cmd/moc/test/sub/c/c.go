package c

import (
	"github.com/dev-drprasad/qt/core"

	_ "github.com/dev-drprasad/qt/internal/cmd/moc/test/sub/conf"
)

type StructSubGoC struct{}
type StructSubMocC struct{ core.QObject }
