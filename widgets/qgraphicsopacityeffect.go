package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"log"
	"unsafe"
)

type QGraphicsOpacityEffect struct {
	QGraphicsEffect
}

type QGraphicsOpacityEffect_ITF interface {
	QGraphicsEffect_ITF
	QGraphicsOpacityEffect_PTR() *QGraphicsOpacityEffect
}

func PointerFromQGraphicsOpacityEffect(ptr QGraphicsOpacityEffect_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsOpacityEffect_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsOpacityEffectFromPointer(ptr unsafe.Pointer) *QGraphicsOpacityEffect {
	var n = new(QGraphicsOpacityEffect)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QGraphicsOpacityEffect_") {
		n.SetObjectName("QGraphicsOpacityEffect_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QGraphicsOpacityEffect) QGraphicsOpacityEffect_PTR() *QGraphicsOpacityEffect {
	return ptr
}

func (ptr *QGraphicsOpacityEffect) Opacity() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsOpacityEffect::opacity")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QGraphicsOpacityEffect_Opacity(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsOpacityEffect) OpacityMask() *gui.QBrush {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsOpacityEffect::opacityMask")
		}
	}()

	if ptr.Pointer() != nil {
		return gui.NewQBrushFromPointer(C.QGraphicsOpacityEffect_OpacityMask(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGraphicsOpacityEffect) SetOpacity(opacity float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsOpacityEffect::setOpacity")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsOpacityEffect_SetOpacity(ptr.Pointer(), C.double(opacity))
	}
}

func (ptr *QGraphicsOpacityEffect) SetOpacityMask(mask gui.QBrush_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsOpacityEffect::setOpacityMask")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsOpacityEffect_SetOpacityMask(ptr.Pointer(), gui.PointerFromQBrush(mask))
	}
}

func NewQGraphicsOpacityEffect(parent core.QObject_ITF) *QGraphicsOpacityEffect {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsOpacityEffect::QGraphicsOpacityEffect")
		}
	}()

	return NewQGraphicsOpacityEffectFromPointer(C.QGraphicsOpacityEffect_NewQGraphicsOpacityEffect(core.PointerFromQObject(parent)))
}

func (ptr *QGraphicsOpacityEffect) ConnectOpacityMaskChanged(f func(mask *gui.QBrush)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsOpacityEffect::opacityMaskChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsOpacityEffect_ConnectOpacityMaskChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "opacityMaskChanged", f)
	}
}

func (ptr *QGraphicsOpacityEffect) DisconnectOpacityMaskChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsOpacityEffect::opacityMaskChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsOpacityEffect_DisconnectOpacityMaskChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "opacityMaskChanged")
	}
}

//export callbackQGraphicsOpacityEffectOpacityMaskChanged
func callbackQGraphicsOpacityEffectOpacityMaskChanged(ptrName *C.char, mask unsafe.Pointer) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsOpacityEffect::opacityMaskChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "opacityMaskChanged").(func(*gui.QBrush))(gui.NewQBrushFromPointer(mask))
}

func (ptr *QGraphicsOpacityEffect) DestroyQGraphicsOpacityEffect() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsOpacityEffect::~QGraphicsOpacityEffect")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsOpacityEffect_DestroyQGraphicsOpacityEffect(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
