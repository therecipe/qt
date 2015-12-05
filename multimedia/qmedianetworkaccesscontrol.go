package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"log"
	"unsafe"
)

type QMediaNetworkAccessControl struct {
	QMediaControl
}

type QMediaNetworkAccessControl_ITF interface {
	QMediaControl_ITF
	QMediaNetworkAccessControl_PTR() *QMediaNetworkAccessControl
}

func PointerFromQMediaNetworkAccessControl(ptr QMediaNetworkAccessControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaNetworkAccessControl_PTR().Pointer()
	}
	return nil
}

func NewQMediaNetworkAccessControlFromPointer(ptr unsafe.Pointer) *QMediaNetworkAccessControl {
	var n = new(QMediaNetworkAccessControl)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QMediaNetworkAccessControl_") {
		n.SetObjectName("QMediaNetworkAccessControl_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QMediaNetworkAccessControl) QMediaNetworkAccessControl_PTR() *QMediaNetworkAccessControl {
	return ptr
}

func (ptr *QMediaNetworkAccessControl) DestroyQMediaNetworkAccessControl() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QMediaNetworkAccessControl::~QMediaNetworkAccessControl")
		}
	}()

	if ptr.Pointer() != nil {
		C.QMediaNetworkAccessControl_DestroyQMediaNetworkAccessControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
