package multimedia

//#include "qmedianetworkaccesscontrol.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QMediaNetworkAccessControl struct {
	QMediaControl
}

type QMediaNetworkAccessControlITF interface {
	QMediaControlITF
	QMediaNetworkAccessControlPTR() *QMediaNetworkAccessControl
}

func PointerFromQMediaNetworkAccessControl(ptr QMediaNetworkAccessControlITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaNetworkAccessControlPTR().Pointer()
	}
	return nil
}

func QMediaNetworkAccessControlFromPointer(ptr unsafe.Pointer) *QMediaNetworkAccessControl {
	var n = new(QMediaNetworkAccessControl)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QMediaNetworkAccessControl_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QMediaNetworkAccessControl) QMediaNetworkAccessControlPTR() *QMediaNetworkAccessControl {
	return ptr
}

func (ptr *QMediaNetworkAccessControl) DestroyQMediaNetworkAccessControl() {
	if ptr.Pointer() != nil {
		C.QMediaNetworkAccessControl_DestroyQMediaNetworkAccessControl(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
