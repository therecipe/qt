package c

import (
	"github.com/bluszcz/cutego/core"

	_ "github.com/bluszcz/cutego/internal/cmd/moc/test/sub/conf"
)

type StructSubGoC struct{}
type StructSubMocC struct{ core.QObject }
