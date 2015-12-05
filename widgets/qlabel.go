package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"log"
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
	for len(n.ObjectName()) < len("QLabel_") {
		n.SetObjectName("QLabel_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QLabel) QLabel_PTR() *QLabel {
	return ptr
}

func (ptr *QLabel) Alignment() core.Qt__AlignmentFlag {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLabel::alignment")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__AlignmentFlag(C.QLabel_Alignment(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLabel) HasScaledContents() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLabel::hasScaledContents")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QLabel_HasScaledContents(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLabel) HasSelectedText() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLabel::hasSelectedText")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QLabel_HasSelectedText(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLabel) Indent() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLabel::indent")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QLabel_Indent(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLabel) Margin() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLabel::margin")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QLabel_Margin(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLabel) OpenExternalLinks() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLabel::openExternalLinks")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QLabel_OpenExternalLinks(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLabel) Pixmap() *gui.QPixmap {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLabel::pixmap")
		}
	}()

	if ptr.Pointer() != nil {
		return gui.NewQPixmapFromPointer(C.QLabel_Pixmap(ptr.Pointer()))
	}
	return nil
}

func (ptr *QLabel) SelectedText() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLabel::selectedText")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QLabel_SelectedText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLabel) SetAlignment(v core.Qt__AlignmentFlag) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLabel::setAlignment")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLabel_SetAlignment(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QLabel) SetIndent(v int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLabel::setIndent")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLabel_SetIndent(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QLabel) SetMargin(v int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLabel::setMargin")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLabel_SetMargin(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QLabel) SetOpenExternalLinks(open bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLabel::setOpenExternalLinks")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLabel_SetOpenExternalLinks(ptr.Pointer(), C.int(qt.GoBoolToInt(open)))
	}
}

func (ptr *QLabel) SetPixmap(v gui.QPixmap_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLabel::setPixmap")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLabel_SetPixmap(ptr.Pointer(), gui.PointerFromQPixmap(v))
	}
}

func (ptr *QLabel) SetScaledContents(v bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLabel::setScaledContents")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLabel_SetScaledContents(ptr.Pointer(), C.int(qt.GoBoolToInt(v)))
	}
}

func (ptr *QLabel) SetText(v string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLabel::setText")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLabel_SetText(ptr.Pointer(), C.CString(v))
	}
}

func (ptr *QLabel) SetTextFormat(v core.Qt__TextFormat) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLabel::setTextFormat")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLabel_SetTextFormat(ptr.Pointer(), C.int(v))
	}
}

func (ptr *QLabel) SetTextInteractionFlags(flags core.Qt__TextInteractionFlag) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLabel::setTextInteractionFlags")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLabel_SetTextInteractionFlags(ptr.Pointer(), C.int(flags))
	}
}

func (ptr *QLabel) SetWordWrap(on bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLabel::setWordWrap")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLabel_SetWordWrap(ptr.Pointer(), C.int(qt.GoBoolToInt(on)))
	}
}

func (ptr *QLabel) Text() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLabel::text")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QLabel_Text(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLabel) TextFormat() core.Qt__TextFormat {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLabel::textFormat")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__TextFormat(C.QLabel_TextFormat(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLabel) TextInteractionFlags() core.Qt__TextInteractionFlag {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLabel::textInteractionFlags")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__TextInteractionFlag(C.QLabel_TextInteractionFlags(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLabel) WordWrap() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLabel::wordWrap")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QLabel_WordWrap(ptr.Pointer()) != 0
	}
	return false
}

func NewQLabel(parent QWidget_ITF, f core.Qt__WindowType) *QLabel {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLabel::QLabel")
		}
	}()

	return NewQLabelFromPointer(C.QLabel_NewQLabel(PointerFromQWidget(parent), C.int(f)))
}

func NewQLabel2(text string, parent QWidget_ITF, f core.Qt__WindowType) *QLabel {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLabel::QLabel")
		}
	}()

	return NewQLabelFromPointer(C.QLabel_NewQLabel2(C.CString(text), PointerFromQWidget(parent), C.int(f)))
}

func (ptr *QLabel) Buddy() *QWidget {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLabel::buddy")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQWidgetFromPointer(C.QLabel_Buddy(ptr.Pointer()))
	}
	return nil
}

func (ptr *QLabel) Clear() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLabel::clear")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLabel_Clear(ptr.Pointer())
	}
}

func (ptr *QLabel) HeightForWidth(w int) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLabel::heightForWidth")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QLabel_HeightForWidth(ptr.Pointer(), C.int(w)))
	}
	return 0
}

func (ptr *QLabel) ConnectLinkActivated(f func(link string)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLabel::linkActivated")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLabel_ConnectLinkActivated(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "linkActivated", f)
	}
}

func (ptr *QLabel) DisconnectLinkActivated() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLabel::linkActivated")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLabel_DisconnectLinkActivated(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "linkActivated")
	}
}

//export callbackQLabelLinkActivated
func callbackQLabelLinkActivated(ptrName *C.char, link *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLabel::linkActivated")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "linkActivated").(func(string))(C.GoString(link))
}

func (ptr *QLabel) ConnectLinkHovered(f func(link string)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLabel::linkHovered")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLabel_ConnectLinkHovered(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "linkHovered", f)
	}
}

func (ptr *QLabel) DisconnectLinkHovered() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLabel::linkHovered")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLabel_DisconnectLinkHovered(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "linkHovered")
	}
}

//export callbackQLabelLinkHovered
func callbackQLabelLinkHovered(ptrName *C.char, link *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLabel::linkHovered")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "linkHovered").(func(string))(C.GoString(link))
}

func (ptr *QLabel) Movie() *gui.QMovie {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLabel::movie")
		}
	}()

	if ptr.Pointer() != nil {
		return gui.NewQMovieFromPointer(C.QLabel_Movie(ptr.Pointer()))
	}
	return nil
}

func (ptr *QLabel) Picture() *gui.QPicture {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLabel::picture")
		}
	}()

	if ptr.Pointer() != nil {
		return gui.NewQPictureFromPointer(C.QLabel_Picture(ptr.Pointer()))
	}
	return nil
}

func (ptr *QLabel) SelectionStart() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLabel::selectionStart")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QLabel_SelectionStart(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLabel) SetBuddy(buddy QWidget_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLabel::setBuddy")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLabel_SetBuddy(ptr.Pointer(), PointerFromQWidget(buddy))
	}
}

func (ptr *QLabel) SetMovie(movie gui.QMovie_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLabel::setMovie")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLabel_SetMovie(ptr.Pointer(), gui.PointerFromQMovie(movie))
	}
}

func (ptr *QLabel) SetNum(num int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLabel::setNum")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLabel_SetNum(ptr.Pointer(), C.int(num))
	}
}

func (ptr *QLabel) SetPicture(picture gui.QPicture_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLabel::setPicture")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLabel_SetPicture(ptr.Pointer(), gui.PointerFromQPicture(picture))
	}
}

func (ptr *QLabel) SetSelection(start int, length int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLabel::setSelection")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLabel_SetSelection(ptr.Pointer(), C.int(start), C.int(length))
	}
}

func (ptr *QLabel) DestroyQLabel() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLabel::~QLabel")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLabel_DestroyQLabel(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
