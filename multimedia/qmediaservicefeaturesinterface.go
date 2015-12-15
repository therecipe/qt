package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QMediaServiceFeaturesInterface struct {
	ptr unsafe.Pointer
}

type QMediaServiceFeaturesInterface_ITF interface {
	QMediaServiceFeaturesInterface_PTR() *QMediaServiceFeaturesInterface
}

func (p *QMediaServiceFeaturesInterface) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QMediaServiceFeaturesInterface) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQMediaServiceFeaturesInterface(ptr QMediaServiceFeaturesInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaServiceFeaturesInterface_PTR().Pointer()
	}
	return nil
}

func NewQMediaServiceFeaturesInterfaceFromPointer(ptr unsafe.Pointer) *QMediaServiceFeaturesInterface {
	var n = new(QMediaServiceFeaturesInterface)
	n.SetPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QMediaServiceFeaturesInterface_") {
		n.SetObjectNameAbs("QMediaServiceFeaturesInterface_" + qt.Identifier())
	}
	return n
}

func (ptr *QMediaServiceFeaturesInterface) QMediaServiceFeaturesInterface_PTR() *QMediaServiceFeaturesInterface {
	return ptr
}

func (ptr *QMediaServiceFeaturesInterface) DestroyQMediaServiceFeaturesInterface() {
	defer qt.Recovering("QMediaServiceFeaturesInterface::~QMediaServiceFeaturesInterface")

	if ptr.Pointer() != nil {
		C.QMediaServiceFeaturesInterface_DestroyQMediaServiceFeaturesInterface(ptr.Pointer())
	}
}

func (ptr *QMediaServiceFeaturesInterface) ObjectNameAbs() string {
	defer qt.Recovering("QMediaServiceFeaturesInterface::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QMediaServiceFeaturesInterface_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QMediaServiceFeaturesInterface) SetObjectNameAbs(name string) {
	defer qt.Recovering("QMediaServiceFeaturesInterface::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QMediaServiceFeaturesInterface_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}
