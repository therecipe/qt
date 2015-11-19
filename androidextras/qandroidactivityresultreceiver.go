// +build !android

package androidextras

import "C"
import (
	"unsafe"
)

type QAndroidActivityResultReceiver struct {
	ptr unsafe.Pointer
}

type QAndroidActivityResultReceiver_ITF interface {
	QAndroidActivityResultReceiver_PTR() *QAndroidActivityResultReceiver
}

func (p *QAndroidActivityResultReceiver) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QAndroidActivityResultReceiver) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQAndroidActivityResultReceiver(ptr QAndroidActivityResultReceiver_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAndroidActivityResultReceiver_PTR().Pointer()
	}
	return nil
}

func NewQAndroidActivityResultReceiverFromPointer(ptr unsafe.Pointer) *QAndroidActivityResultReceiver {
	var n = new(QAndroidActivityResultReceiver)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAndroidActivityResultReceiver) QAndroidActivityResultReceiver_PTR() *QAndroidActivityResultReceiver {
	return ptr
}

func (ptr *QAndroidActivityResultReceiver) HandleActivityResult(receiverRequestCode int, resultCode int, data QAndroidJniObject_ITF) {

	return
}
