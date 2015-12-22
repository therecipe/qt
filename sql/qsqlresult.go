package sql

//#include "sql.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QSqlResult struct {
	ptr unsafe.Pointer
}

type QSqlResult_ITF interface {
	QSqlResult_PTR() *QSqlResult
}

func (p *QSqlResult) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSqlResult) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSqlResult(ptr QSqlResult_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSqlResult_PTR().Pointer()
	}
	return nil
}

func NewQSqlResultFromPointer(ptr unsafe.Pointer) *QSqlResult {
	var n = new(QSqlResult)
	n.SetPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QSqlResult_") {
		n.SetObjectNameAbs("QSqlResult_" + qt.Identifier())
	}
	return n
}

func (ptr *QSqlResult) QSqlResult_PTR() *QSqlResult {
	return ptr
}

//QSqlResult::BindingSyntax
type QSqlResult__BindingSyntax int64

const (
	QSqlResult__PositionalBinding = QSqlResult__BindingSyntax(0)
	QSqlResult__NamedBinding      = QSqlResult__BindingSyntax(1)
)

func (ptr *QSqlResult) Handle() *core.QVariant {
	defer qt.Recovering("QSqlResult::handle")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QSqlResult_Handle(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSqlResult) ConnectSetActive(f func(active bool)) {
	defer qt.Recovering("connect QSqlResult::setActive")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "setActive", f)
	}
}

func (ptr *QSqlResult) DisconnectSetActive() {
	defer qt.Recovering("disconnect QSqlResult::setActive")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "setActive")
	}
}

//export callbackQSqlResultSetActive
func callbackQSqlResultSetActive(ptrName *C.char, active C.int) bool {
	defer qt.Recovering("callback QSqlResult::setActive")

	if signal := qt.GetSignal(C.GoString(ptrName), "setActive"); signal != nil {
		signal.(func(bool))(int(active) != 0)
		return true
	}
	return false

}

func (ptr *QSqlResult) ConnectSetAt(f func(index int)) {
	defer qt.Recovering("connect QSqlResult::setAt")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "setAt", f)
	}
}

func (ptr *QSqlResult) DisconnectSetAt() {
	defer qt.Recovering("disconnect QSqlResult::setAt")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "setAt")
	}
}

//export callbackQSqlResultSetAt
func callbackQSqlResultSetAt(ptrName *C.char, index C.int) bool {
	defer qt.Recovering("callback QSqlResult::setAt")

	if signal := qt.GetSignal(C.GoString(ptrName), "setAt"); signal != nil {
		signal.(func(int))(int(index))
		return true
	}
	return false

}

func (ptr *QSqlResult) ConnectSetForwardOnly(f func(forward bool)) {
	defer qt.Recovering("connect QSqlResult::setForwardOnly")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "setForwardOnly", f)
	}
}

func (ptr *QSqlResult) DisconnectSetForwardOnly() {
	defer qt.Recovering("disconnect QSqlResult::setForwardOnly")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "setForwardOnly")
	}
}

//export callbackQSqlResultSetForwardOnly
func callbackQSqlResultSetForwardOnly(ptrName *C.char, forward C.int) bool {
	defer qt.Recovering("callback QSqlResult::setForwardOnly")

	if signal := qt.GetSignal(C.GoString(ptrName), "setForwardOnly"); signal != nil {
		signal.(func(bool))(int(forward) != 0)
		return true
	}
	return false

}

func (ptr *QSqlResult) ConnectSetQuery(f func(query string)) {
	defer qt.Recovering("connect QSqlResult::setQuery")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "setQuery", f)
	}
}

func (ptr *QSqlResult) DisconnectSetQuery() {
	defer qt.Recovering("disconnect QSqlResult::setQuery")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "setQuery")
	}
}

//export callbackQSqlResultSetQuery
func callbackQSqlResultSetQuery(ptrName *C.char, query *C.char) bool {
	defer qt.Recovering("callback QSqlResult::setQuery")

	if signal := qt.GetSignal(C.GoString(ptrName), "setQuery"); signal != nil {
		signal.(func(string))(C.GoString(query))
		return true
	}
	return false

}

func (ptr *QSqlResult) DestroyQSqlResult() {
	defer qt.Recovering("QSqlResult::~QSqlResult")

	if ptr.Pointer() != nil {
		C.QSqlResult_DestroyQSqlResult(ptr.Pointer())
	}
}

func (ptr *QSqlResult) ObjectNameAbs() string {
	defer qt.Recovering("QSqlResult::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QSqlResult_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QSqlResult) SetObjectNameAbs(name string) {
	defer qt.Recovering("QSqlResult::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QSqlResult_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}
