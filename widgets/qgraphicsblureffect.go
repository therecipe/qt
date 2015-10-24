package widgets

//#include "qgraphicsblureffect.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QGraphicsBlurEffect struct {
	QGraphicsEffect
}

type QGraphicsBlurEffectITF interface {
	QGraphicsEffectITF
	QGraphicsBlurEffectPTR() *QGraphicsBlurEffect
}

func PointerFromQGraphicsBlurEffect(ptr QGraphicsBlurEffectITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsBlurEffectPTR().Pointer()
	}
	return nil
}

func QGraphicsBlurEffectFromPointer(ptr unsafe.Pointer) *QGraphicsBlurEffect {
	var n = new(QGraphicsBlurEffect)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QGraphicsBlurEffect_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QGraphicsBlurEffect) QGraphicsBlurEffectPTR() *QGraphicsBlurEffect {
	return ptr
}

//QGraphicsBlurEffect::BlurHint
type QGraphicsBlurEffect__BlurHint int

var (
	QGraphicsBlurEffect__PerformanceHint = QGraphicsBlurEffect__BlurHint(0x00)
	QGraphicsBlurEffect__QualityHint     = QGraphicsBlurEffect__BlurHint(0x01)
	QGraphicsBlurEffect__AnimationHint   = QGraphicsBlurEffect__BlurHint(0x02)
)

func (ptr *QGraphicsBlurEffect) BlurHints() QGraphicsBlurEffect__BlurHint {
	if ptr.Pointer() != nil {
		return QGraphicsBlurEffect__BlurHint(C.QGraphicsBlurEffect_BlurHints(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGraphicsBlurEffect) SetBlurHints(hints QGraphicsBlurEffect__BlurHint) {
	if ptr.Pointer() != nil {
		C.QGraphicsBlurEffect_SetBlurHints(C.QtObjectPtr(ptr.Pointer()), C.int(hints))
	}
}

func NewQGraphicsBlurEffect(parent core.QObjectITF) *QGraphicsBlurEffect {
	return QGraphicsBlurEffectFromPointer(unsafe.Pointer(C.QGraphicsBlurEffect_NewQGraphicsBlurEffect(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QGraphicsBlurEffect) ConnectBlurHintsChanged(f func(hints QGraphicsBlurEffect__BlurHint)) {
	if ptr.Pointer() != nil {
		C.QGraphicsBlurEffect_ConnectBlurHintsChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "blurHintsChanged", f)
	}
}

func (ptr *QGraphicsBlurEffect) DisconnectBlurHintsChanged() {
	if ptr.Pointer() != nil {
		C.QGraphicsBlurEffect_DisconnectBlurHintsChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "blurHintsChanged")
	}
}

//export callbackQGraphicsBlurEffectBlurHintsChanged
func callbackQGraphicsBlurEffectBlurHintsChanged(ptrName *C.char, hints C.int) {
	qt.GetSignal(C.GoString(ptrName), "blurHintsChanged").(func(QGraphicsBlurEffect__BlurHint))(QGraphicsBlurEffect__BlurHint(hints))
}

func (ptr *QGraphicsBlurEffect) DestroyQGraphicsBlurEffect() {
	if ptr.Pointer() != nil {
		C.QGraphicsBlurEffect_DestroyQGraphicsBlurEffect(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
