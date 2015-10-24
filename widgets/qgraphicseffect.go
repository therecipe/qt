package widgets

//#include "qgraphicseffect.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QGraphicsEffect struct {
	core.QObject
}

type QGraphicsEffectITF interface {
	core.QObjectITF
	QGraphicsEffectPTR() *QGraphicsEffect
}

func PointerFromQGraphicsEffect(ptr QGraphicsEffectITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsEffectPTR().Pointer()
	}
	return nil
}

func QGraphicsEffectFromPointer(ptr unsafe.Pointer) *QGraphicsEffect {
	var n = new(QGraphicsEffect)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QGraphicsEffect_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QGraphicsEffect) QGraphicsEffectPTR() *QGraphicsEffect {
	return ptr
}

//QGraphicsEffect::ChangeFlag
type QGraphicsEffect__ChangeFlag int

var (
	QGraphicsEffect__SourceAttached            = QGraphicsEffect__ChangeFlag(0x1)
	QGraphicsEffect__SourceDetached            = QGraphicsEffect__ChangeFlag(0x2)
	QGraphicsEffect__SourceBoundingRectChanged = QGraphicsEffect__ChangeFlag(0x4)
	QGraphicsEffect__SourceInvalidated         = QGraphicsEffect__ChangeFlag(0x8)
)

//QGraphicsEffect::PixmapPadMode
type QGraphicsEffect__PixmapPadMode int

var (
	QGraphicsEffect__NoPad                      = QGraphicsEffect__PixmapPadMode(0)
	QGraphicsEffect__PadToTransparentBorder     = QGraphicsEffect__PixmapPadMode(1)
	QGraphicsEffect__PadToEffectiveBoundingRect = QGraphicsEffect__PixmapPadMode(2)
)
