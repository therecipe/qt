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

type QLabel_ITF interface {
	QFrame_ITF
	QLabel_PTR() *QLabel
}

func PointerFromQLabel(ptr QLabel_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLabel_PTR().Pointer()
	}
	return nil
}

func NewQLabelFromPointer(ptr unsafe.Pointer) *QLabel {
	var n = new(QLabel)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QLabel_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QLabel) QLabel_PTR() *QLabel {
	return ptr
}

func (ptr *QLabel) Alignment() core.Qt__AlignmentFlag {
	if ptr.Pointer() != nil {
		return core.Qt__AlignmentFlag(C.QLabel_Alignment(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLabel) HasScaledContents() bool {
	if ptr.Pointer() != nil {
		return C.QLabel_HasScaledContents(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLabel) HasSelectedText() bool {
	if ptr.Pointer() != nil {
		return C.QLabel_HasSelectedText(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLabel) Indent() int {
	if ptr.Pointer() != nil {
		return int(C.QLabel_Indent(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLabel) Margin() int {
	if ptr.Pointer() != nil {
		return int(C.QLabel_Margin(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLabel) OpenExternalLinks() bool {
	if ptr.Pointer() != nil {
		return C.QLabel_OpenExternalLinks(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLabel) Pixmap() *gui.QPixmap {
	if ptr.Pointer() != nil {
		return gui.NewQPixmapFromPointer(C.QLabel_Pixmap(ptr.Pointer()))
	}
	return nil
}

func (ptr *QLabel) SelectedText() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QLabel_SelectedText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLabel) SetAlignment(v core.Qt__AlignmentFlag) {
	if ptr.Pointer() != nil {
		C.QLabel_SetAlignment(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QLabel) SetIndent(v int) {
	if ptr.Pointer() != nil {
		C.QLabel_SetIndent(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QLabel) SetMargin(v int) {
	if ptr.Pointer() != nil {
		C.QLabel_SetMargin(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QLabel) SetOpenExternalLinks(open bool) {
	if ptr.Pointer() != nil {
		C.QLabel_SetOpenExternalLinks(ptr.Pointer(), C.int(qt.GoBoolToInt(open)))
	}
}

func (ptr *QLabel) SetPixmap(v gui.QPixmap_ITF) {
	if ptr.Pointer() != nil {
		C.QLabel_SetPixmap(ptr.Pointer(), gui.PointerFromQPixmap(v))
	}
}

func (ptr *QLabel) SetScaledContents(v bool) {
	if ptr.Pointer() != nil {
		C.QLabel_SetScaledContents(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QLabel) SetText(v string) {
	if ptr.Pointer() != nil {
		C.QLabel_SetText(ptr.Pointer(), C.CString(v))
	}
}

func (ptr *QLabel) SetTextFormat(v core.Qt__TextFormat) {
	if ptr.Pointer() != nil {
		C.QLabel_SetTextFormat(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QLabel) SetTextInteractionFlags(flags core.Qt__TextInteractionFlag) {
	if ptr.Pointer() != nil {
		C.QLabel_SetTextInteractionFlags(ptr.Pointer(), C.int(flags))
	}
}

func (ptr *QLabel) SetWordWrap(on bool) {
	if ptr.Pointer() != nil {
		C.QLabel_SetWordWrap(ptr.Pointer(), C.int(qt.GoBoolToInt(on)))
	}
}

func (ptr *QLabel) Text() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QLabel_Text(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLabel) TextFormat() core.Qt__TextFormat {
	if ptr.Pointer() != nil {
		return core.Qt__TextFormat(C.QLabel_TextFormat(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLabel) TextInteractionFlags() core.Qt__TextInteractionFlag {
	if ptr.Pointer() != nil {
		return core.Qt__TextInteractionFlag(C.QLabel_TextInteractionFlags(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLabel) WordWrap() bool {
	if ptr.Pointer() != nil {
		return C.QLabel_WordWrap(ptr.Pointer()) != 0
	}
	return false
}

func NewQLabel(parent QWidget_ITF, f core.Qt__WindowType) *QLabel {
	return NewQLabelFromPointer(C.QLabel_NewQLabel(PointerFromQWidget(parent), C.int(f)))
}

func NewQLabel2(text string, parent QWidget_ITF, f core.Qt__WindowType) *QLabel {
	return NewQLabelFromPointer(C.QLabel_NewQLabel2(C.CString(text), PointerFromQWidget(parent), C.int(f)))
}

func (ptr *QLabel) Buddy() *QWidget {
	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QLabel_Buddy(ptr.Pointer()))
	}
	return nil
}

func (ptr *QLabel) Clear() {
	if ptr.Pointer() != nil {
		C.QLabel_Clear(ptr.Pointer())
	}
}

func (ptr *QLabel) HeightForWidth(w int) int {
	if ptr.Pointer() != nil {
		return int(C.QLabel_HeightForWidth(ptr.Pointer(), C.int(w)))
	}
	return 0
}

func (ptr *QLabel) ConnectLinkActivated(f func(link string)) {
	if ptr.Pointer() != nil {
		C.QLabel_ConnectLinkActivated(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "linkActivated", f)
	}
}

func (ptr *QLabel) DisconnectLinkActivated() {
	if ptr.Pointer() != nil {
		C.QLabel_DisconnectLinkActivated(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "linkActivated")
	}
}

//export callbackQLabelLinkActivated
func callbackQLabelLinkActivated(ptrName *C.char, link *C.char) {
	qt.GetSignal(C.GoString(ptrName), "linkActivated").(func(string))(C.GoString(link))
}

func (ptr *QLabel) ConnectLinkHovered(f func(link string)) {
	if ptr.Pointer() != nil {
		C.QLabel_ConnectLinkHovered(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "linkHovered", f)
	}
}

func (ptr *QLabel) DisconnectLinkHovered() {
	if ptr.Pointer() != nil {
		C.QLabel_DisconnectLinkHovered(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "linkHovered")
	}
}

//export callbackQLabelLinkHovered
func callbackQLabelLinkHovered(ptrName *C.char, link *C.char) {
	qt.GetSignal(C.GoString(ptrName), "linkHovered").(func(string))(C.GoString(link))
}

func (ptr *QLabel) Movie() *gui.QMovie {
	if ptr.Pointer() != nil {
		return gui.NewQMovieFromPointer(C.QLabel_Movie(ptr.Pointer()))
	}
	return nil
}

func (ptr *QLabel) Picture() *gui.QPicture {
	if ptr.Pointer() != nil {
		return gui.NewQPictureFromPointer(C.QLabel_Picture(ptr.Pointer()))
	}
	return nil
}

func (ptr *QLabel) SelectionStart() int {
	if ptr.Pointer() != nil {
		return int(C.QLabel_SelectionStart(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLabel) SetBuddy(buddy QWidget_ITF) {
	if ptr.Pointer() != nil {
		C.QLabel_SetBuddy(ptr.Pointer(), PointerFromQWidget(buddy))
	}
}

func (ptr *QLabel) SetMovie(movie gui.QMovie_ITF) {
	if ptr.Pointer() != nil {
		C.QLabel_SetMovie(ptr.Pointer(), gui.PointerFromQMovie(movie))
	}
}

func (ptr *QLabel) SetNum(num int) {
	if ptr.Pointer() != nil {
		C.QLabel_SetNum(ptr.Pointer(), C.int(num))
	}
}

func (ptr *QLabel) SetPicture(picture gui.QPicture_ITF) {
	if ptr.Pointer() != nil {
		C.QLabel_SetPicture(ptr.Pointer(), gui.PointerFromQPicture(picture))
	}
}

func (ptr *QLabel) SetSelection(start int, length int) {
	if ptr.Pointer() != nil {
		C.QLabel_SetSelection(ptr.Pointer(), C.int(start), C.int(length))
	}
}

func (ptr *QLabel) DestroyQLabel() {
	if ptr.Pointer() != nil {
		C.QLabel_DestroyQLabel(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
