package core

//#include "qurlquery.h"
import "C"
import (
	"strings"
	"unsafe"
)

type QUrlQuery struct {
	ptr unsafe.Pointer
}

type QUrlQueryITF interface {
	QUrlQueryPTR() *QUrlQuery
}

func (p *QUrlQuery) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QUrlQuery) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQUrlQuery(ptr QUrlQueryITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QUrlQueryPTR().Pointer()
	}
	return nil
}

func QUrlQueryFromPointer(ptr unsafe.Pointer) *QUrlQuery {
	var n = new(QUrlQuery)
	n.SetPointer(ptr)
	return n
}

func (ptr *QUrlQuery) QUrlQueryPTR() *QUrlQuery {
	return ptr
}

func NewQUrlQuery() *QUrlQuery {
	return QUrlQueryFromPointer(unsafe.Pointer(C.QUrlQuery_NewQUrlQuery()))
}

func NewQUrlQuery3(queryString string) *QUrlQuery {
	return QUrlQueryFromPointer(unsafe.Pointer(C.QUrlQuery_NewQUrlQuery3(C.CString(queryString))))
}

func NewQUrlQuery2(url string) *QUrlQuery {
	return QUrlQueryFromPointer(unsafe.Pointer(C.QUrlQuery_NewQUrlQuery2(C.CString(url))))
}

func NewQUrlQuery4(other QUrlQueryITF) *QUrlQuery {
	return QUrlQueryFromPointer(unsafe.Pointer(C.QUrlQuery_NewQUrlQuery4(C.QtObjectPtr(PointerFromQUrlQuery(other)))))
}

func (ptr *QUrlQuery) AddQueryItem(key string, value string) {
	if ptr.Pointer() != nil {
		C.QUrlQuery_AddQueryItem(C.QtObjectPtr(ptr.Pointer()), C.CString(key), C.CString(value))
	}
}

func (ptr *QUrlQuery) AllQueryItemValues(key string, encoding QUrl__ComponentFormattingOption) []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QUrlQuery_AllQueryItemValues(C.QtObjectPtr(ptr.Pointer()), C.CString(key), C.int(encoding))), "|")
	}
	return make([]string, 0)
}

func (ptr *QUrlQuery) Clear() {
	if ptr.Pointer() != nil {
		C.QUrlQuery_Clear(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QUrlQuery) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return C.QUrlQuery_IsEmpty(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QUrlQuery) Query(encoding QUrl__ComponentFormattingOption) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QUrlQuery_Query(C.QtObjectPtr(ptr.Pointer()), C.int(encoding)))
	}
	return ""
}

func (ptr *QUrlQuery) RemoveAllQueryItems(key string) {
	if ptr.Pointer() != nil {
		C.QUrlQuery_RemoveAllQueryItems(C.QtObjectPtr(ptr.Pointer()), C.CString(key))
	}
}

func (ptr *QUrlQuery) RemoveQueryItem(key string) {
	if ptr.Pointer() != nil {
		C.QUrlQuery_RemoveQueryItem(C.QtObjectPtr(ptr.Pointer()), C.CString(key))
	}
}

func (ptr *QUrlQuery) SetQuery(queryString string) {
	if ptr.Pointer() != nil {
		C.QUrlQuery_SetQuery(C.QtObjectPtr(ptr.Pointer()), C.CString(queryString))
	}
}

func (ptr *QUrlQuery) SetQueryDelimiters(valueDelimiter QCharITF, pairDelimiter QCharITF) {
	if ptr.Pointer() != nil {
		C.QUrlQuery_SetQueryDelimiters(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQChar(valueDelimiter)), C.QtObjectPtr(PointerFromQChar(pairDelimiter)))
	}
}

func (ptr *QUrlQuery) Swap(other QUrlQueryITF) {
	if ptr.Pointer() != nil {
		C.QUrlQuery_Swap(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQUrlQuery(other)))
	}
}

func (ptr *QUrlQuery) ToString(encoding QUrl__ComponentFormattingOption) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QUrlQuery_ToString(C.QtObjectPtr(ptr.Pointer()), C.int(encoding)))
	}
	return ""
}

func (ptr *QUrlQuery) DestroyQUrlQuery() {
	if ptr.Pointer() != nil {
		C.QUrlQuery_DestroyQUrlQuery(C.QtObjectPtr(ptr.Pointer()))
	}
}
