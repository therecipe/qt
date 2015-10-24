package xml

//#include "qdomcomment.h"
import "C"
import (
	"unsafe"
)

type QDomComment struct {
	QDomCharacterData
}

type QDomCommentITF interface {
	QDomCharacterDataITF
	QDomCommentPTR() *QDomComment
}

func PointerFromQDomComment(ptr QDomCommentITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDomCommentPTR().Pointer()
	}
	return nil
}

func QDomCommentFromPointer(ptr unsafe.Pointer) *QDomComment {
	var n = new(QDomComment)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDomComment) QDomCommentPTR() *QDomComment {
	return ptr
}

func NewQDomComment() *QDomComment {
	return QDomCommentFromPointer(unsafe.Pointer(C.QDomComment_NewQDomComment()))
}

func NewQDomComment2(x QDomCommentITF) *QDomComment {
	return QDomCommentFromPointer(unsafe.Pointer(C.QDomComment_NewQDomComment2(C.QtObjectPtr(PointerFromQDomComment(x)))))
}

func (ptr *QDomComment) NodeType() QDomNode__NodeType {
	if ptr.Pointer() != nil {
		return QDomNode__NodeType(C.QDomComment_NodeType(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}
