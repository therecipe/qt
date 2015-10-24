package multimedia

//#include "qmediastreamscontrol.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QMediaStreamsControl struct {
	QMediaControl
}

type QMediaStreamsControlITF interface {
	QMediaControlITF
	QMediaStreamsControlPTR() *QMediaStreamsControl
}

func PointerFromQMediaStreamsControl(ptr QMediaStreamsControlITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMediaStreamsControlPTR().Pointer()
	}
	return nil
}

func QMediaStreamsControlFromPointer(ptr unsafe.Pointer) *QMediaStreamsControl {
	var n = new(QMediaStreamsControl)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QMediaStreamsControl_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QMediaStreamsControl) QMediaStreamsControlPTR() *QMediaStreamsControl {
	return ptr
}

//QMediaStreamsControl::StreamType
type QMediaStreamsControl__StreamType int

var (
	QMediaStreamsControl__UnknownStream    = QMediaStreamsControl__StreamType(0)
	QMediaStreamsControl__VideoStream      = QMediaStreamsControl__StreamType(1)
	QMediaStreamsControl__AudioStream      = QMediaStreamsControl__StreamType(2)
	QMediaStreamsControl__SubPictureStream = QMediaStreamsControl__StreamType(3)
	QMediaStreamsControl__DataStream       = QMediaStreamsControl__StreamType(4)
)

func (ptr *QMediaStreamsControl) ConnectActiveStreamsChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QMediaStreamsControl_ConnectActiveStreamsChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "activeStreamsChanged", f)
	}
}

func (ptr *QMediaStreamsControl) DisconnectActiveStreamsChanged() {
	if ptr.Pointer() != nil {
		C.QMediaStreamsControl_DisconnectActiveStreamsChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "activeStreamsChanged")
	}
}

//export callbackQMediaStreamsControlActiveStreamsChanged
func callbackQMediaStreamsControlActiveStreamsChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "activeStreamsChanged").(func())()
}

func (ptr *QMediaStreamsControl) IsActive(stream int) bool {
	if ptr.Pointer() != nil {
		return C.QMediaStreamsControl_IsActive(C.QtObjectPtr(ptr.Pointer()), C.int(stream)) != 0
	}
	return false
}

func (ptr *QMediaStreamsControl) MetaData(stream int, key string) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QMediaStreamsControl_MetaData(C.QtObjectPtr(ptr.Pointer()), C.int(stream), C.CString(key)))
	}
	return ""
}

func (ptr *QMediaStreamsControl) SetActive(stream int, state bool) {
	if ptr.Pointer() != nil {
		C.QMediaStreamsControl_SetActive(C.QtObjectPtr(ptr.Pointer()), C.int(stream), C.int(qt.GoBoolToInt(state)))
	}
}

func (ptr *QMediaStreamsControl) StreamCount() int {
	if ptr.Pointer() != nil {
		return int(C.QMediaStreamsControl_StreamCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QMediaStreamsControl) StreamType(stream int) QMediaStreamsControl__StreamType {
	if ptr.Pointer() != nil {
		return QMediaStreamsControl__StreamType(C.QMediaStreamsControl_StreamType(C.QtObjectPtr(ptr.Pointer()), C.int(stream)))
	}
	return 0
}

func (ptr *QMediaStreamsControl) ConnectStreamsChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QMediaStreamsControl_ConnectStreamsChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "streamsChanged", f)
	}
}

func (ptr *QMediaStreamsControl) DisconnectStreamsChanged() {
	if ptr.Pointer() != nil {
		C.QMediaStreamsControl_DisconnectStreamsChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "streamsChanged")
	}
}

//export callbackQMediaStreamsControlStreamsChanged
func callbackQMediaStreamsControlStreamsChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "streamsChanged").(func())()
}

func (ptr *QMediaStreamsControl) DestroyQMediaStreamsControl() {
	if ptr.Pointer() != nil {
		C.QMediaStreamsControl_DestroyQMediaStreamsControl(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
