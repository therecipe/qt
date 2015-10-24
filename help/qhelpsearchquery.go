package help

//#include "qhelpsearchquery.h"
import "C"
import (
	"strings"
	"unsafe"
)

type QHelpSearchQuery struct {
	ptr unsafe.Pointer
}

type QHelpSearchQueryITF interface {
	QHelpSearchQueryPTR() *QHelpSearchQuery
}

func (p *QHelpSearchQuery) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QHelpSearchQuery) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQHelpSearchQuery(ptr QHelpSearchQueryITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpSearchQueryPTR().Pointer()
	}
	return nil
}

func QHelpSearchQueryFromPointer(ptr unsafe.Pointer) *QHelpSearchQuery {
	var n = new(QHelpSearchQuery)
	n.SetPointer(ptr)
	return n
}

func (ptr *QHelpSearchQuery) QHelpSearchQueryPTR() *QHelpSearchQuery {
	return ptr
}

//QHelpSearchQuery::FieldName
type QHelpSearchQuery__FieldName int

var (
	QHelpSearchQuery__DEFAULT = QHelpSearchQuery__FieldName(0)
	QHelpSearchQuery__FUZZY   = QHelpSearchQuery__FieldName(1)
	QHelpSearchQuery__WITHOUT = QHelpSearchQuery__FieldName(2)
	QHelpSearchQuery__PHRASE  = QHelpSearchQuery__FieldName(3)
	QHelpSearchQuery__ALL     = QHelpSearchQuery__FieldName(4)
	QHelpSearchQuery__ATLEAST = QHelpSearchQuery__FieldName(5)
)

func NewQHelpSearchQuery() *QHelpSearchQuery {
	return QHelpSearchQueryFromPointer(unsafe.Pointer(C.QHelpSearchQuery_NewQHelpSearchQuery()))
}

func NewQHelpSearchQuery2(field QHelpSearchQuery__FieldName, wordList []string) *QHelpSearchQuery {
	return QHelpSearchQueryFromPointer(unsafe.Pointer(C.QHelpSearchQuery_NewQHelpSearchQuery2(C.int(field), C.CString(strings.Join(wordList, "|")))))
}
