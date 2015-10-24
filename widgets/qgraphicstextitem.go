package widgets

//#include "qgraphicstextitem.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QGraphicsTextItem struct {
	QGraphicsObject
}

type QGraphicsTextItemITF interface {
	QGraphicsObjectITF
	QGraphicsTextItemPTR() *QGraphicsTextItem
}

func PointerFromQGraphicsTextItem(ptr QGraphicsTextItemITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsTextItemPTR().Pointer()
	}
	return nil
}

func QGraphicsTextItemFromPointer(ptr unsafe.Pointer) *QGraphicsTextItem {
	var n = new(QGraphicsTextItem)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QGraphicsTextItem_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QGraphicsTextItem) QGraphicsTextItemPTR() *QGraphicsTextItem {
	return ptr
}

func (ptr *QGraphicsTextItem) OpenExternalLinks() bool {
	if ptr.Pointer() != nil {
		return C.QGraphicsTextItem_OpenExternalLinks(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QGraphicsTextItem) SetOpenExternalLinks(open bool) {
	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_SetOpenExternalLinks(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(open)))
	}
}

func (ptr *QGraphicsTextItem) SetTextCursor(cursor gui.QTextCursorITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_SetTextCursor(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQTextCursor(cursor)))
	}
}

func NewQGraphicsTextItem(parent QGraphicsItemITF) *QGraphicsTextItem {
	return QGraphicsTextItemFromPointer(unsafe.Pointer(C.QGraphicsTextItem_NewQGraphicsTextItem(C.QtObjectPtr(PointerFromQGraphicsItem(parent)))))
}

func NewQGraphicsTextItem2(text string, parent QGraphicsItemITF) *QGraphicsTextItem {
	return QGraphicsTextItemFromPointer(unsafe.Pointer(C.QGraphicsTextItem_NewQGraphicsTextItem2(C.CString(text), C.QtObjectPtr(PointerFromQGraphicsItem(parent)))))
}

func (ptr *QGraphicsTextItem) AdjustSize() {
	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_AdjustSize(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QGraphicsTextItem) Contains(point core.QPointFITF) bool {
	if ptr.Pointer() != nil {
		return C.QGraphicsTextItem_Contains(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPointF(point))) != 0
	}
	return false
}

func (ptr *QGraphicsTextItem) Document() *gui.QTextDocument {
	if ptr.Pointer() != nil {
		return gui.QTextDocumentFromPointer(unsafe.Pointer(C.QGraphicsTextItem_Document(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QGraphicsTextItem) IsObscuredBy(item QGraphicsItemITF) bool {
	if ptr.Pointer() != nil {
		return C.QGraphicsTextItem_IsObscuredBy(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQGraphicsItem(item))) != 0
	}
	return false
}

func (ptr *QGraphicsTextItem) ConnectLinkActivated(f func(link string)) {
	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_ConnectLinkActivated(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "linkActivated", f)
	}
}

func (ptr *QGraphicsTextItem) DisconnectLinkActivated() {
	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_DisconnectLinkActivated(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "linkActivated")
	}
}

//export callbackQGraphicsTextItemLinkActivated
func callbackQGraphicsTextItemLinkActivated(ptrName *C.char, link *C.char) {
	qt.GetSignal(C.GoString(ptrName), "linkActivated").(func(string))(C.GoString(link))
}

func (ptr *QGraphicsTextItem) ConnectLinkHovered(f func(link string)) {
	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_ConnectLinkHovered(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "linkHovered", f)
	}
}

func (ptr *QGraphicsTextItem) DisconnectLinkHovered() {
	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_DisconnectLinkHovered(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "linkHovered")
	}
}

//export callbackQGraphicsTextItemLinkHovered
func callbackQGraphicsTextItemLinkHovered(ptrName *C.char, link *C.char) {
	qt.GetSignal(C.GoString(ptrName), "linkHovered").(func(string))(C.GoString(link))
}

func (ptr *QGraphicsTextItem) Paint(painter gui.QPainterITF, option QStyleOptionGraphicsItemITF, widget QWidgetITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_Paint(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQPainter(painter)), C.QtObjectPtr(PointerFromQStyleOptionGraphicsItem(option)), C.QtObjectPtr(PointerFromQWidget(widget)))
	}
}

func (ptr *QGraphicsTextItem) SetDefaultTextColor(col gui.QColorITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_SetDefaultTextColor(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQColor(col)))
	}
}

func (ptr *QGraphicsTextItem) SetDocument(document gui.QTextDocumentITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_SetDocument(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQTextDocument(document)))
	}
}

func (ptr *QGraphicsTextItem) SetFont(font gui.QFontITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_SetFont(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQFont(font)))
	}
}

func (ptr *QGraphicsTextItem) SetHtml(text string) {
	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_SetHtml(C.QtObjectPtr(ptr.Pointer()), C.CString(text))
	}
}

func (ptr *QGraphicsTextItem) SetPlainText(text string) {
	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_SetPlainText(C.QtObjectPtr(ptr.Pointer()), C.CString(text))
	}
}

func (ptr *QGraphicsTextItem) SetTabChangesFocus(b bool) {
	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_SetTabChangesFocus(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(b)))
	}
}

func (ptr *QGraphicsTextItem) SetTextInteractionFlags(flags core.Qt__TextInteractionFlag) {
	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_SetTextInteractionFlags(C.QtObjectPtr(ptr.Pointer()), C.int(flags))
	}
}

func (ptr *QGraphicsTextItem) TabChangesFocus() bool {
	if ptr.Pointer() != nil {
		return C.QGraphicsTextItem_TabChangesFocus(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QGraphicsTextItem) TextInteractionFlags() core.Qt__TextInteractionFlag {
	if ptr.Pointer() != nil {
		return core.Qt__TextInteractionFlag(C.QGraphicsTextItem_TextInteractionFlags(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGraphicsTextItem) ToHtml() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QGraphicsTextItem_ToHtml(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QGraphicsTextItem) ToPlainText() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QGraphicsTextItem_ToPlainText(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QGraphicsTextItem) Type() int {
	if ptr.Pointer() != nil {
		return int(C.QGraphicsTextItem_Type(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGraphicsTextItem) DestroyQGraphicsTextItem() {
	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_DestroyQGraphicsTextItem(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
