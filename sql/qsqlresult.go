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

func (ptr *QSqlResult) Exec() bool {
	defer qt.Recovering("QSqlResult::exec")

	if ptr.Pointer() != nil {
		return C.QSqlResult_Exec(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlResult) FetchNext() bool {
	defer qt.Recovering("QSqlResult::fetchNext")

	if ptr.Pointer() != nil {
		return C.QSqlResult_FetchNext(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlResult) FetchPrevious() bool {
	defer qt.Recovering("QSqlResult::fetchPrevious")

	if ptr.Pointer() != nil {
		return C.QSqlResult_FetchPrevious(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSqlResult) Handle() *core.QVariant {
	defer qt.Recovering("QSqlResult::handle")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QSqlResult_Handle(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSqlResult) LastInsertId() *core.QVariant {
	defer qt.Recovering("QSqlResult::lastInsertId")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QSqlResult_LastInsertId(ptr.Pointer()))
	}
	return nil
}

func (ptr *QSqlResult) Prepare(query string) bool {
	defer qt.Recovering("QSqlResult::prepare")

	if ptr.Pointer() != nil {
		return C.QSqlResult_Prepare(ptr.Pointer(), C.CString(query)) != 0
	}
	return false
}

func (ptr *QSqlResult) SavePrepare(query string) bool {
	defer qt.Recovering("QSqlResult::savePrepare")

	if ptr.Pointer() != nil {
		return C.QSqlResult_SavePrepare(ptr.Pointer(), C.CString(query)) != 0
	}
	return false
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
func callbackQSqlResultSetActive(ptr unsafe.Pointer, ptrName *C.char, active C.int) {
	defer qt.Recovering("callback QSqlResult::setActive")

	if signal := qt.GetSignal(C.GoString(ptrName), "setActive"); signal != nil {
		signal.(func(bool))(int(active) != 0)
	} else {
		NewQSqlResultFromPointer(ptr).SetActiveDefault(int(active) != 0)
	}
}

func (ptr *QSqlResult) SetActive(active bool) {
	defer qt.Recovering("QSqlResult::setActive")

	if ptr.Pointer() != nil {
		C.QSqlResult_SetActive(ptr.Pointer(), C.int(qt.GoBoolToInt(active)))
	}
}

func (ptr *QSqlResult) SetActiveDefault(active bool) {
	defer qt.Recovering("QSqlResult::setActive")

	if ptr.Pointer() != nil {
		C.QSqlResult_SetActiveDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(active)))
	}
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
func callbackQSqlResultSetAt(ptr unsafe.Pointer, ptrName *C.char, index C.int) {
	defer qt.Recovering("callback QSqlResult::setAt")

	if signal := qt.GetSignal(C.GoString(ptrName), "setAt"); signal != nil {
		signal.(func(int))(int(index))
	} else {
		NewQSqlResultFromPointer(ptr).SetAtDefault(int(index))
	}
}

func (ptr *QSqlResult) SetAt(index int) {
	defer qt.Recovering("QSqlResult::setAt")

	if ptr.Pointer() != nil {
		C.QSqlResult_SetAt(ptr.Pointer(), C.int(index))
	}
}

func (ptr *QSqlResult) SetAtDefault(index int) {
	defer qt.Recovering("QSqlResult::setAt")

	if ptr.Pointer() != nil {
		C.QSqlResult_SetAtDefault(ptr.Pointer(), C.int(index))
	}
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
func callbackQSqlResultSetForwardOnly(ptr unsafe.Pointer, ptrName *C.char, forward C.int) {
	defer qt.Recovering("callback QSqlResult::setForwardOnly")

	if signal := qt.GetSignal(C.GoString(ptrName), "setForwardOnly"); signal != nil {
		signal.(func(bool))(int(forward) != 0)
	} else {
		NewQSqlResultFromPointer(ptr).SetForwardOnlyDefault(int(forward) != 0)
	}
}

func (ptr *QSqlResult) SetForwardOnly(forward bool) {
	defer qt.Recovering("QSqlResult::setForwardOnly")

	if ptr.Pointer() != nil {
		C.QSqlResult_SetForwardOnly(ptr.Pointer(), C.int(qt.GoBoolToInt(forward)))
	}
}

func (ptr *QSqlResult) SetForwardOnlyDefault(forward bool) {
	defer qt.Recovering("QSqlResult::setForwardOnly")

	if ptr.Pointer() != nil {
		C.QSqlResult_SetForwardOnlyDefault(ptr.Pointer(), C.int(qt.GoBoolToInt(forward)))
	}
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
func callbackQSqlResultSetQuery(ptr unsafe.Pointer, ptrName *C.char, query *C.char) {
	defer qt.Recovering("callback QSqlResult::setQuery")

	if signal := qt.GetSignal(C.GoString(ptrName), "setQuery"); signal != nil {
		signal.(func(string))(C.GoString(query))
	} else {
		NewQSqlResultFromPointer(ptr).SetQueryDefault(C.GoString(query))
	}
}

func (ptr *QSqlResult) SetQuery(query string) {
	defer qt.Recovering("QSqlResult::setQuery")

	if ptr.Pointer() != nil {
		C.QSqlResult_SetQuery(ptr.Pointer(), C.CString(query))
	}
}

func (ptr *QSqlResult) SetQueryDefault(query string) {
	defer qt.Recovering("QSqlResult::setQuery")

	if ptr.Pointer() != nil {
		C.QSqlResult_SetQueryDefault(ptr.Pointer(), C.CString(query))
	}
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
