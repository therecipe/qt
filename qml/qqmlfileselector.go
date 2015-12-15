package qml

//#include "qml.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

type QQmlFileSelector struct {
	core.QObject
}

type QQmlFileSelector_ITF interface {
	core.QObject_ITF
	QQmlFileSelector_PTR() *QQmlFileSelector
}

func PointerFromQQmlFileSelector(ptr QQmlFileSelector_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlFileSelector_PTR().Pointer()
	}
	return nil
}

func NewQQmlFileSelectorFromPointer(ptr unsafe.Pointer) *QQmlFileSelector {
	var n = new(QQmlFileSelector)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QQmlFileSelector_") {
		n.SetObjectName("QQmlFileSelector_" + qt.Identifier())
	}
	return n
}

func (ptr *QQmlFileSelector) QQmlFileSelector_PTR() *QQmlFileSelector {
	return ptr
}

func NewQQmlFileSelector(engine QQmlEngine_ITF, parent core.QObject_ITF) *QQmlFileSelector {
	defer qt.Recovering("QQmlFileSelector::QQmlFileSelector")

	return NewQQmlFileSelectorFromPointer(C.QQmlFileSelector_NewQQmlFileSelector(PointerFromQQmlEngine(engine), core.PointerFromQObject(parent)))
}

func QQmlFileSelector_Get(engine QQmlEngine_ITF) *QQmlFileSelector {
	defer qt.Recovering("QQmlFileSelector::get")

	return NewQQmlFileSelectorFromPointer(C.QQmlFileSelector_QQmlFileSelector_Get(PointerFromQQmlEngine(engine)))
}

func (ptr *QQmlFileSelector) SetExtraSelectors(strin []string) {
	defer qt.Recovering("QQmlFileSelector::setExtraSelectors")

	if ptr.Pointer() != nil {
		C.QQmlFileSelector_SetExtraSelectors(ptr.Pointer(), C.CString(strings.Join(strin, ",,,")))
	}
}

func (ptr *QQmlFileSelector) SetExtraSelectors2(strin []string) {
	defer qt.Recovering("QQmlFileSelector::setExtraSelectors")

	if ptr.Pointer() != nil {
		C.QQmlFileSelector_SetExtraSelectors2(ptr.Pointer(), C.CString(strings.Join(strin, ",,,")))
	}
}

func (ptr *QQmlFileSelector) SetSelector(selector core.QFileSelector_ITF) {
	defer qt.Recovering("QQmlFileSelector::setSelector")

	if ptr.Pointer() != nil {
		C.QQmlFileSelector_SetSelector(ptr.Pointer(), core.PointerFromQFileSelector(selector))
	}
}

func (ptr *QQmlFileSelector) DestroyQQmlFileSelector() {
	defer qt.Recovering("QQmlFileSelector::~QQmlFileSelector")

	if ptr.Pointer() != nil {
		C.QQmlFileSelector_DestroyQQmlFileSelector(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
