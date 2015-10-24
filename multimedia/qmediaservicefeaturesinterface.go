package multimedia

//#include "qmediaservicefeaturesinterface.h"
import "C"
import (
	"unsafe"
)

type QMediaServiceFeaturesInterface struct {
	ptr unsafe.Pointer
}

type QMediaServiceFeaturesInterfaceITF interface {
	QMediaServiceFeaturesInterfacePTR() *QMediaServiceFeaturesInterface
}

func (p *QMediaServiceFeaturesInterface) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QMediaServiceFeaturesInterface) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQMediaServiceFeaturesInterface(ptr QMediaServiceFeaturesInterfaceITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaServiceFeaturesInterfacePTR().Pointer()
	}
	return nil
}

func QMediaServiceFeaturesInterfaceFromPointer(ptr unsafe.Pointer) *QMediaServiceFeaturesInterface {
	var n = new(QMediaServiceFeaturesInterface)
	n.SetPointer(ptr)
	return n
}

func (ptr *QMediaServiceFeaturesInterface) QMediaServiceFeaturesInterfacePTR() *QMediaServiceFeaturesInterface {
	return ptr
}

func (ptr *QMediaServiceFeaturesInterface) DestroyQMediaServiceFeaturesInterface() {
	if ptr.Pointer() != nil {
		C.QMediaServiceFeaturesInterface_DestroyQMediaServiceFeaturesInterface(C.QtObjectPtr(ptr.Pointer()))
	}
}
