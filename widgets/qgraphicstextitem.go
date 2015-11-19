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

type QGraphicsTextItem_ITF interface {
	QGraphicsObject_ITF
	QGraphicsTextItem_PTR() *QGraphicsTextItem
}

func PointerFromQGraphicsTextItem(ptr QGraphicsTextItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsTextItem_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsTextItemFromPointer(ptr unsafe.Pointer) *QGraphicsTextItem {
	var n = new(QGraphicsTextItem)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QGraphicsTextItem_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QGraphicsTextItem) QGraphicsTextItem_PTR() *QGraphicsTextItem {
	return ptr
}

func (ptr *QGraphicsTextItem) OpenExternalLinks() bool {
	if ptr.Pointer() != nil {
		return C.QGraphicsTextItem_OpenExternalLinks(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGraphicsTextItem) SetOpenExternalLinks(open bool) {
	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_SetOpenExternalLinks(ptr.Pointer(), C.int(qt.GoBoolToInt(open)))
	}
}

func (ptr *QGraphicsTextItem) SetTextCursor(cursor gui.QTextCursor_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_SetTextCursor(ptr.Pointer(), gui.PointerFromQTextCursor(cursor))
	}
}

func (ptr *QGraphicsTextItem) AdjustSize() {
	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_AdjustSize(ptr.Pointer())
	}
}

func (ptr *QGraphicsTextItem) Contains(point core.QPointF_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QGraphicsTextItem_Contains(ptr.Pointer(), core.PointerFromQPointF(point)) != 0
	}
	return false
}

func (ptr *QGraphicsTextItem) DefaultTextColor() *gui.QColor {
	if ptr.Pointer() != nil {
		return gui.NewQColorFromPointer(C.QGraphicsTextItem_DefaultTextColor(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGraphicsTextItem) Document() *gui.QTextDocument {
	if ptr.Pointer() != nil {
		return gui.NewQTextDocumentFromPointer(C.QGraphicsTextItem_Document(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGraphicsTextItem) IsObscuredBy(item QGraphicsItem_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QGraphicsTextItem_IsObscuredBy(ptr.Pointer(), PointerFromQGraphicsItem(item)) != 0
	}
	return false
}

func (ptr *QGraphicsTextItem) ConnectLinkActivated(f func(link string)) {
	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_ConnectLinkActivated(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "linkActivated", f)
	}
}

func (ptr *QGraphicsTextItem) DisconnectLinkActivated() {
	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_DisconnectLinkActivated(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "linkActivated")
	}
}

//export callbackQGraphicsTextItemLinkActivated
func callbackQGraphicsTextItemLinkActivated(ptrName *C.char, link *C.char) {
	qt.GetSignal(C.GoString(ptrName), "linkActivated").(func(string))(C.GoString(link))
}

func (ptr *QGraphicsTextItem) ConnectLinkHovered(f func(link string)) {
	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_ConnectLinkHovered(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "linkHovered", f)
	}
}

func (ptr *QGraphicsTextItem) DisconnectLinkHovered() {
	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_DisconnectLinkHovered(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "linkHovered")
	}
}

//export callbackQGraphicsTextItemLinkHovered
func callbackQGraphicsTextItemLinkHovered(ptrName *C.char, link *C.char) {
	qt.GetSignal(C.GoString(ptrName), "linkHovered").(func(string))(C.GoString(link))
}

func (ptr *QGraphicsTextItem) Paint(painter gui.QPainter_ITF, option QStyleOptionGraphicsItem_ITF, widget QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_Paint(ptr.Pointer(), gui.PointerFromQPainter(painter), PointerFromQStyleOptionGraphicsItem(option), PointerFromQWidget(widget))
	}
}

func (ptr *QGraphicsTextItem) SetDefaultTextColor(col gui.QColor_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_SetDefaultTextColor(ptr.Pointer(), gui.PointerFromQColor(col))
	}
}

func (ptr *QGraphicsTextItem) SetDocument(document gui.QTextDocument_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_SetDocument(ptr.Pointer(), gui.PointerFromQTextDocument(document))
	}
}

func (ptr *QGraphicsTextItem) SetFont(font gui.QFont_ITF) {
	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_SetFont(ptr.Pointer(), gui.PointerFromQFont(font))
	}
}

func (ptr *QGraphicsTextItem) SetHtml(text string) {
	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_SetHtml(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QGraphicsTextItem) SetPlainText(text string) {
	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_SetPlainText(ptr.Pointer(), C.CString(text))
	}
}

func (ptr *QGraphicsTextItem) SetTabChangesFocus(b bool) {
	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_SetTabChangesFocus(ptr.Pointer(), C.int(qt.GoBoolToInt(b)))
	}
}

func (ptr *QGraphicsTextItem) SetTextInteractionFlags(flags core.Qt__TextInteractionFlag) {
	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_SetTextInteractionFlags(ptr.Pointer(), C.int(flags))
	}
}

func (ptr *QGraphicsTextItem) SetTextWidth(width float64) {
	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_SetTextWidth(ptr.Pointer(), C.double(width))
	}
}

func (ptr *QGraphicsTextItem) TabChangesFocus() bool {
	if ptr.Pointer() != nil {
		return C.QGraphicsTextItem_TabChangesFocus(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGraphicsTextItem) TextInteractionFlags() core.Qt__TextInteractionFlag {
	if ptr.Pointer() != nil {
		return core.Qt__TextInteractionFlag(C.QGraphicsTextItem_TextInteractionFlags(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsTextItem) TextWidth() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QGraphicsTextItem_TextWidth(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsTextItem) ToHtml() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QGraphicsTextItem_ToHtml(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGraphicsTextItem) ToPlainText() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QGraphicsTextItem_ToPlainText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QGraphicsTextItem) Type() int {
	if ptr.Pointer() != nil {
		return int(C.QGraphicsTextItem_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsTextItem) DestroyQGraphicsTextItem() {
	if ptr.Pointer() != nil {
		C.QGraphicsTextItem_DestroyQGraphicsTextItem(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
