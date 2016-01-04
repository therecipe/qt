package xml

//#include "xml.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QXmlInputSource struct {
	ptr unsafe.Pointer
}

type QXmlInputSource_ITF interface {
	QXmlInputSource_PTR() *QXmlInputSource
}

func (p *QXmlInputSource) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QXmlInputSource) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQXmlInputSource(ptr QXmlInputSource_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlInputSource_PTR().Pointer()
	}
	return nil
}

func NewQXmlInputSourceFromPointer(ptr unsafe.Pointer) *QXmlInputSource {
	var n = new(QXmlInputSource)
	n.SetPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QXmlInputSource_") {
		n.SetObjectNameAbs("QXmlInputSource_" + qt.Identifier())
	}
	return n
}

func (ptr *QXmlInputSource) QXmlInputSource_PTR() *QXmlInputSource {
	return ptr
}

func NewQXmlInputSource() *QXmlInputSource {
	defer qt.Recovering("QXmlInputSource::QXmlInputSource")

	return NewQXmlInputSourceFromPointer(C.QXmlInputSource_NewQXmlInputSource())
}

func NewQXmlInputSource2(dev core.QIODevice_ITF) *QXmlInputSource {
	defer qt.Recovering("QXmlInputSource::QXmlInputSource")

	return NewQXmlInputSourceFromPointer(C.QXmlInputSource_NewQXmlInputSource2(core.PointerFromQIODevice(dev)))
}

func (ptr *QXmlInputSource) Data() string {
	defer qt.Recovering("QXmlInputSource::data")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlInputSource_Data(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlInputSource) ConnectFetchData(f func()) {
	defer qt.Recovering("connect QXmlInputSource::fetchData")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "fetchData", f)
	}
}

func (ptr *QXmlInputSource) DisconnectFetchData() {
	defer qt.Recovering("disconnect QXmlInputSource::fetchData")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "fetchData")
	}
}

//export callbackQXmlInputSourceFetchData
func callbackQXmlInputSourceFetchData(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QXmlInputSource::fetchData")

	if signal := qt.GetSignal(C.GoString(ptrName), "fetchData"); signal != nil {
		signal.(func())()
	} else {
		NewQXmlInputSourceFromPointer(ptr).FetchDataDefault()
	}
}

func (ptr *QXmlInputSource) FetchData() {
	defer qt.Recovering("QXmlInputSource::fetchData")

	if ptr.Pointer() != nil {
		C.QXmlInputSource_FetchData(ptr.Pointer())
	}
}

func (ptr *QXmlInputSource) FetchDataDefault() {
	defer qt.Recovering("QXmlInputSource::fetchData")

	if ptr.Pointer() != nil {
		C.QXmlInputSource_FetchDataDefault(ptr.Pointer())
	}
}

func (ptr *QXmlInputSource) FromRawData(data core.QByteArray_ITF, beginning bool) string {
	defer qt.Recovering("QXmlInputSource::fromRawData")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlInputSource_FromRawData(ptr.Pointer(), core.PointerFromQByteArray(data), C.int(qt.GoBoolToInt(beginning))))
	}
	return ""
}

func (ptr *QXmlInputSource) ConnectReset(f func()) {
	defer qt.Recovering("connect QXmlInputSource::reset")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "reset", f)
	}
}

func (ptr *QXmlInputSource) DisconnectReset() {
	defer qt.Recovering("disconnect QXmlInputSource::reset")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "reset")
	}
}

//export callbackQXmlInputSourceReset
func callbackQXmlInputSourceReset(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QXmlInputSource::reset")

	if signal := qt.GetSignal(C.GoString(ptrName), "reset"); signal != nil {
		signal.(func())()
	} else {
		NewQXmlInputSourceFromPointer(ptr).ResetDefault()
	}
}

func (ptr *QXmlInputSource) Reset() {
	defer qt.Recovering("QXmlInputSource::reset")

	if ptr.Pointer() != nil {
		C.QXmlInputSource_Reset(ptr.Pointer())
	}
}

func (ptr *QXmlInputSource) ResetDefault() {
	defer qt.Recovering("QXmlInputSource::reset")

	if ptr.Pointer() != nil {
		C.QXmlInputSource_ResetDefault(ptr.Pointer())
	}
}

func (ptr *QXmlInputSource) ConnectSetData(f func(dat string)) {
	defer qt.Recovering("connect QXmlInputSource::setData")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "setData", f)
	}
}

func (ptr *QXmlInputSource) DisconnectSetData() {
	defer qt.Recovering("disconnect QXmlInputSource::setData")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "setData")
	}
}

//export callbackQXmlInputSourceSetData
func callbackQXmlInputSourceSetData(ptr unsafe.Pointer, ptrName *C.char, dat *C.char) {
	defer qt.Recovering("callback QXmlInputSource::setData")

	if signal := qt.GetSignal(C.GoString(ptrName), "setData"); signal != nil {
		signal.(func(string))(C.GoString(dat))
	} else {
		NewQXmlInputSourceFromPointer(ptr).SetDataDefault(C.GoString(dat))
	}
}

func (ptr *QXmlInputSource) SetData(dat string) {
	defer qt.Recovering("QXmlInputSource::setData")

	if ptr.Pointer() != nil {
		C.QXmlInputSource_SetData(ptr.Pointer(), C.CString(dat))
	}
}

func (ptr *QXmlInputSource) SetDataDefault(dat string) {
	defer qt.Recovering("QXmlInputSource::setData")

	if ptr.Pointer() != nil {
		C.QXmlInputSource_SetDataDefault(ptr.Pointer(), C.CString(dat))
	}
}

func (ptr *QXmlInputSource) DestroyQXmlInputSource() {
	defer qt.Recovering("QXmlInputSource::~QXmlInputSource")

	if ptr.Pointer() != nil {
		C.QXmlInputSource_DestroyQXmlInputSource(ptr.Pointer())
	}
}

func (ptr *QXmlInputSource) ObjectNameAbs() string {
	defer qt.Recovering("QXmlInputSource::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlInputSource_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlInputSource) SetObjectNameAbs(name string) {
	defer qt.Recovering("QXmlInputSource::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QXmlInputSource_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}
