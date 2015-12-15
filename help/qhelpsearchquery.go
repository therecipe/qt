package help

//#include "help.h"
import "C"
import (
	"github.com/therecipe/qt"
	"strings"
	"unsafe"
)

type QHelpSearchQuery struct {
	ptr unsafe.Pointer
}

type QHelpSearchQuery_ITF interface {
	QHelpSearchQuery_PTR() *QHelpSearchQuery
}

func (p *QHelpSearchQuery) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QHelpSearchQuery) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQHelpSearchQuery(ptr QHelpSearchQuery_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpSearchQuery_PTR().Pointer()
	}
	return nil
}

func NewQHelpSearchQueryFromPointer(ptr unsafe.Pointer) *QHelpSearchQuery {
	var n = new(QHelpSearchQuery)
	n.SetPointer(ptr)
	return n
}

func (ptr *QHelpSearchQuery) QHelpSearchQuery_PTR() *QHelpSearchQuery {
	return ptr
}

//QHelpSearchQuery::FieldName
type QHelpSearchQuery__FieldName int64

const (
	QHelpSearchQuery__DEFAULT = QHelpSearchQuery__FieldName(0)
	QHelpSearchQuery__FUZZY   = QHelpSearchQuery__FieldName(1)
	QHelpSearchQuery__WITHOUT = QHelpSearchQuery__FieldName(2)
	QHelpSearchQuery__PHRASE  = QHelpSearchQuery__FieldName(3)
	QHelpSearchQuery__ALL     = QHelpSearchQuery__FieldName(4)
	QHelpSearchQuery__ATLEAST = QHelpSearchQuery__FieldName(5)
)

func NewQHelpSearchQuery() *QHelpSearchQuery {
	defer qt.Recovering("QHelpSearchQuery::QHelpSearchQuery")

	return NewQHelpSearchQueryFromPointer(C.QHelpSearchQuery_NewQHelpSearchQuery())
}

func NewQHelpSearchQuery2(field QHelpSearchQuery__FieldName, wordList []string) *QHelpSearchQuery {
	defer qt.Recovering("QHelpSearchQuery::QHelpSearchQuery")

	return NewQHelpSearchQueryFromPointer(C.QHelpSearchQuery_NewQHelpSearchQuery2(C.int(field), C.CString(strings.Join(wordList, ",,,"))))
}
