package widgets

//#include "qlabel.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QLabel struct {
	QFrame
}

type QLabelITF interface {
	QFrameITF
	QLabelPTR() *QLabel
}

func PointerFromQLabel(ptr QLabelITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLabelPTR().Pointer()
	}
	return nil
}

func QLabelFromPointer(ptr unsafe.Pointer) *QLabel {
	var n = new(QLabel)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QLabel_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QLabel) QLabelPTR() *QLabel {
	return ptr
}

func (ptr *QLabel) Alignment() core.Qt__AlignmentFlag {
	if ptr.Pointer() != nil {
		return core.Qt__AlignmentFlag(C.QLabel_Alignment(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QLabel) HasScaledContents() bool {
	if ptr.Pointer() != nil {
		return C.QLabel_HasScaledContents(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QLabel) HasSelectedText() bool {
	if ptr.Pointer() != nil {
		return C.QLabel_HasSelectedText(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QLabel) Indent() int {
	if ptr.Pointer() != nil {
		return int(C.QLabel_Indent(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QLabel) Margin() int {
	if ptr.Pointer() != nil {
		return int(C.QLabel_Margin(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QLabel) OpenExternalLinks() bool {
	if ptr.Pointer() != nil {
		return C.QLabel_OpenExternalLinks(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QLabel) Pixmap() *gui.QPixmap {
	if ptr.Pointer() != nil {
		return gui.QPixmapFromPointer(unsafe.Pointer(C.QLabel_Pixmap(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QLabel) SelectedText() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QLabel_SelectedText(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QLabel) SetAlignment(v core.Qt__AlignmentFlag) {
	if ptr.Pointer() != nil {
		C.QLabel_SetAlignment(C.QtObjectPtr(ptr.Pointer()), C.int(v))
	}
}

func (ptr *QLabel) SetIndent(v int) {
	if ptr.Pointer() != nil {
		C.QLabel_SetIndent(C.QtObjectPtr(ptr.Pointer()), C.int(v))
	}
}

func (ptr *QLabel) SetMargin(v int) {
	if ptr.Pointer() != nil {
		C.QLabel_SetMargin(C.QtObjectPtr(ptr.Pointer()), C.int(v))
	}
}

func (ptr *QLabel) SetOpenExternalLinks(open bool) {
	if ptr.Pointer() != nil {
		C.QLabel_SetOpenExternalLinks(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(open)))
	}
}

func (ptr *QLabel) SetPixmap(v gui.QPixmapITF) {
	if ptr.Pointer() != nil {
		C.QLabel_SetPixmap(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQPixmap(v)))
	}
}

func (ptr *QLabel) SetScaledContents(v bool) {
	if ptr.Pointer() != nil {
		C.QLabel_SetScaledContents(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QLabel) SetText(v string) {
	if ptr.Pointer() != nil {
		C.QLabel_SetText(C.QtObjectPtr(ptr.Pointer()), C.CString(v))
	}
}

func (ptr *QLabel) SetTextFormat(v core.Qt__TextFormat) {
	if ptr.Pointer() != nil {
		C.QLabel_SetTextFormat(C.QtObjectPtr(ptr.Pointer()), C.int(v))
	}
}

func (ptr *QLabel) SetTextInteractionFlags(flags core.Qt__TextInteractionFlag) {
	if ptr.Pointer() != nil {
		C.QLabel_SetTextInteractionFlags(C.QtObjectPtr(ptr.Pointer()), C.int(flags))
	}
}

func (ptr *QLabel) SetWordWrap(on bool) {
	if ptr.Pointer() != nil {
		C.QLabel_SetWordWrap(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(on)))
	}
}

func (ptr *QLabel) Text() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QLabel_Text(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QLabel) TextFormat() core.Qt__TextFormat {
	if ptr.Pointer() != nil {
		return core.Qt__TextFormat(C.QLabel_TextFormat(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QLabel) TextInteractionFlags() core.Qt__TextInteractionFlag {
	if ptr.Pointer() != nil {
		return core.Qt__TextInteractionFlag(C.QLabel_TextInteractionFlags(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QLabel) WordWrap() bool {
	if ptr.Pointer() != nil {
		return C.QLabel_WordWrap(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func NewQLabel(parent QWidgetITF, f core.Qt__WindowType) *QLabel {
	return QLabelFromPointer(unsafe.Pointer(C.QLabel_NewQLabel(C.QtObjectPtr(PointerFromQWidget(parent)), C.int(f))))
}

func NewQLabel2(text string, parent QWidgetITF, f core.Qt__WindowType) *QLabel {
	return QLabelFromPointer(unsafe.Pointer(C.QLabel_NewQLabel2(C.CString(text), C.QtObjectPtr(PointerFromQWidget(parent)), C.int(f))))
}

func (ptr *QLabel) Buddy() *QWidget {
	if ptr.Pointer() != nil {
		return QWidgetFromPointer(unsafe.Pointer(C.QLabel_Buddy(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QLabel) Clear() {
	if ptr.Pointer() != nil {
		C.QLabel_Clear(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QLabel) HeightForWidth(w int) int {
	if ptr.Pointer() != nil {
		return int(C.QLabel_HeightForWidth(C.QtObjectPtr(ptr.Pointer()), C.int(w)))
	}
	return 0
}

func (ptr *QLabel) ConnectLinkActivated(f func(link string)) {
	if ptr.Pointer() != nil {
		C.QLabel_ConnectLinkActivated(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "linkActivated", f)
	}
}

func (ptr *QLabel) DisconnectLinkActivated() {
	if ptr.Pointer() != nil {
		C.QLabel_DisconnectLinkActivated(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "linkActivated")
	}
}

//export callbackQLabelLinkActivated
func callbackQLabelLinkActivated(ptrName *C.char, link *C.char) {
	qt.GetSignal(C.GoString(ptrName), "linkActivated").(func(string))(C.GoString(link))
}

func (ptr *QLabel) ConnectLinkHovered(f func(link string)) {
	if ptr.Pointer() != nil {
		C.QLabel_ConnectLinkHovered(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "linkHovered", f)
	}
}

func (ptr *QLabel) DisconnectLinkHovered() {
	if ptr.Pointer() != nil {
		C.QLabel_DisconnectLinkHovered(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "linkHovered")
	}
}

//export callbackQLabelLinkHovered
func callbackQLabelLinkHovered(ptrName *C.char, link *C.char) {
	qt.GetSignal(C.GoString(ptrName), "linkHovered").(func(string))(C.GoString(link))
}

func (ptr *QLabel) Movie() *gui.QMovie {
	if ptr.Pointer() != nil {
		return gui.QMovieFromPointer(unsafe.Pointer(C.QLabel_Movie(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QLabel) Picture() *gui.QPicture {
	if ptr.Pointer() != nil {
		return gui.QPictureFromPointer(unsafe.Pointer(C.QLabel_Picture(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QLabel) SelectionStart() int {
	if ptr.Pointer() != nil {
		return int(C.QLabel_SelectionStart(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QLabel) SetBuddy(buddy QWidgetITF) {
	if ptr.Pointer() != nil {
		C.QLabel_SetBuddy(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQWidget(buddy)))
	}
}

func (ptr *QLabel) SetMovie(movie gui.QMovieITF) {
	if ptr.Pointer() != nil {
		C.QLabel_SetMovie(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQMovie(movie)))
	}
}

func (ptr *QLabel) SetNum(num int) {
	if ptr.Pointer() != nil {
		C.QLabel_SetNum(C.QtObjectPtr(ptr.Pointer()), C.int(num))
	}
}

func (ptr *QLabel) SetPicture(picture gui.QPictureITF) {
	if ptr.Pointer() != nil {
		C.QLabel_SetPicture(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(gui.PointerFromQPicture(picture)))
	}
}

func (ptr *QLabel) SetSelection(start int, length int) {
	if ptr.Pointer() != nil {
		C.QLabel_SetSelection(C.QtObjectPtr(ptr.Pointer()), C.int(start), C.int(length))
	}
}

func (ptr *QLabel) DestroyQLabel() {
	if ptr.Pointer() != nil {
		C.QLabel_DestroyQLabel(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
