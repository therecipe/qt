package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QGraphicsEffect struct {
	core.QObject
}

type QGraphicsEffect_ITF interface {
	core.QObject_ITF
	QGraphicsEffect_PTR() *QGraphicsEffect
}

func PointerFromQGraphicsEffect(ptr QGraphicsEffect_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsEffect_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsEffectFromPointer(ptr unsafe.Pointer) *QGraphicsEffect {
	var n = new(QGraphicsEffect)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QGraphicsEffect_") {
		n.SetObjectName("QGraphicsEffect_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QGraphicsEffect) QGraphicsEffect_PTR() *QGraphicsEffect {
	return ptr
}

//QGraphicsEffect::ChangeFlag
type QGraphicsEffect__ChangeFlag int64

const (
	QGraphicsEffect__SourceAttached            = QGraphicsEffect__ChangeFlag(0x1)
	QGraphicsEffect__SourceDetached            = QGraphicsEffect__ChangeFlag(0x2)
	QGraphicsEffect__SourceBoundingRectChanged = QGraphicsEffect__ChangeFlag(0x4)
	QGraphicsEffect__SourceInvalidated         = QGraphicsEffect__ChangeFlag(0x8)
)

//QGraphicsEffect::PixmapPadMode
type QGraphicsEffect__PixmapPadMode int64

const (
	QGraphicsEffect__NoPad                      = QGraphicsEffect__PixmapPadMode(0)
	QGraphicsEffect__PadToTransparentBorder     = QGraphicsEffect__PixmapPadMode(1)
	QGraphicsEffect__PadToEffectiveBoundingRect = QGraphicsEffect__PixmapPadMode(2)
)
