package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QGroupBox struct {
	QWidget
}

type QGroupBox_ITF interface {
	QWidget_ITF
	QGroupBox_PTR() *QGroupBox
}

func PointerFromQGroupBox(ptr QGroupBox_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGroupBox_PTR().Pointer()
	}
	return nil
}

func NewQGroupBoxFromPointer(ptr unsafe.Pointer) *QGroupBox {
	var n = new(QGroupBox)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QGroupBox_") {
		n.SetObjectName("QGroupBox_" + qt.Identifier())
	}
	return n
}

func (ptr *QGroupBox) QGroupBox_PTR() *QGroupBox {
	return ptr
}

func (ptr *QGroupBox) Alignment() core.Qt__AlignmentFlag {
	defer qt.Recovering("QGroupBox::alignment")

	if ptr.Pointer() != nil {
		return core.Qt__AlignmentFlag(C.QGroupBox_Alignment(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGroupBox) IsCheckable() bool {
	defer qt.Recovering("QGroupBox::isCheckable")

	if ptr.Pointer() != nil {
		return C.QGroupBox_IsCheckable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGroupBox) IsChecked() bool {
	defer qt.Recovering("QGroupBox::isChecked")

	if ptr.Pointer() != nil {
		return C.QGroupBox_IsChecked(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGroupBox) IsFlat() bool {
	defer qt.Recovering("QGroupBox::isFlat")

	if ptr.Pointer() != nil {
		return C.QGroupBox_IsFlat(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QGroupBox) SetAlignment(alignment int) {
	defer qt.Recovering("QGroupBox::setAlignment")

	if ptr.Pointer() != nil {
		C.QGroupBox_SetAlignment(ptr.Pointer(), C.int(alignment))
	}
}

func (ptr *QGroupBox) SetCheckable(checkable bool) {
	defer qt.Recovering("QGroupBox::setCheckable")

	if ptr.Pointer() != nil {
		C.QGroupBox_SetCheckable(ptr.Pointer(), C.int(qt.GoBoolToInt(checkable)))
	}
}

func (ptr *QGroupBox) SetChecked(checked bool) {
	defer qt.Recovering("QGroupBox::setChecked")

	if ptr.Pointer() != nil {
		C.QGroupBox_SetChecked(ptr.Pointer(), C.int(qt.GoBoolToInt(checked)))
	}
}

func (ptr *QGroupBox) SetFlat(flat bool) {
	defer qt.Recovering("QGroupBox::setFlat")

	if ptr.Pointer() != nil {
		C.QGroupBox_SetFlat(ptr.Pointer(), C.int(qt.GoBoolToInt(flat)))
	}
}

func (ptr *QGroupBox) SetTitle(title string) {
	defer qt.Recovering("QGroupBox::setTitle")

	if ptr.Pointer() != nil {
		C.QGroupBox_SetTitle(ptr.Pointer(), C.CString(title))
	}
}

func (ptr *QGroupBox) Title() string {
	defer qt.Recovering("QGroupBox::title")

	if ptr.Pointer() != nil {
		return C.GoString(C.QGroupBox_Title(ptr.Pointer()))
	}
	return ""
}

func NewQGroupBox(parent QWidget_ITF) *QGroupBox {
	defer qt.Recovering("QGroupBox::QGroupBox")

	return NewQGroupBoxFromPointer(C.QGroupBox_NewQGroupBox(PointerFromQWidget(parent)))
}

func NewQGroupBox2(title string, parent QWidget_ITF) *QGroupBox {
	defer qt.Recovering("QGroupBox::QGroupBox")

	return NewQGroupBoxFromPointer(C.QGroupBox_NewQGroupBox2(C.CString(title), PointerFromQWidget(parent)))
}

func (ptr *QGroupBox) ConnectChangeEvent(f func(ev *core.QEvent)) {
	defer qt.Recovering("connect QGroupBox::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QGroupBox) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QGroupBox::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQGroupBoxChangeEvent
func callbackQGroupBoxChangeEvent(ptrName *C.char, ev unsafe.Pointer) bool {
	defer qt.Recovering("callback QGroupBox::changeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "changeEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(ev))
		return true
	}
	return false

}

func (ptr *QGroupBox) ConnectChildEvent(f func(c *core.QChildEvent)) {
	defer qt.Recovering("connect QGroupBox::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QGroupBox) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QGroupBox::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQGroupBoxChildEvent
func callbackQGroupBoxChildEvent(ptrName *C.char, c unsafe.Pointer) bool {
	defer qt.Recovering("callback QGroupBox::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(c))
		return true
	}
	return false

}

func (ptr *QGroupBox) ConnectClicked(f func(checked bool)) {
	defer qt.Recovering("connect QGroupBox::clicked")

	if ptr.Pointer() != nil {
		C.QGroupBox_ConnectClicked(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "clicked", f)
	}
}

func (ptr *QGroupBox) DisconnectClicked() {
	defer qt.Recovering("disconnect QGroupBox::clicked")

	if ptr.Pointer() != nil {
		C.QGroupBox_DisconnectClicked(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "clicked")
	}
}

//export callbackQGroupBoxClicked
func callbackQGroupBoxClicked(ptrName *C.char, checked C.int) {
	defer qt.Recovering("callback QGroupBox::clicked")

	if signal := qt.GetSignal(C.GoString(ptrName), "clicked"); signal != nil {
		signal.(func(bool))(int(checked) != 0)
	}

}

func (ptr *QGroupBox) ConnectFocusInEvent(f func(fe *gui.QFocusEvent)) {
	defer qt.Recovering("connect QGroupBox::focusInEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "focusInEvent", f)
	}
}

func (ptr *QGroupBox) DisconnectFocusInEvent() {
	defer qt.Recovering("disconnect QGroupBox::focusInEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "focusInEvent")
	}
}

//export callbackQGroupBoxFocusInEvent
func callbackQGroupBoxFocusInEvent(ptrName *C.char, fe unsafe.Pointer) bool {
	defer qt.Recovering("callback QGroupBox::focusInEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "focusInEvent"); signal != nil {
		signal.(func(*gui.QFocusEvent))(gui.NewQFocusEventFromPointer(fe))
		return true
	}
	return false

}

func (ptr *QGroupBox) MinimumSizeHint() *core.QSize {
	defer qt.Recovering("QGroupBox::minimumSizeHint")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QGroupBox_MinimumSizeHint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGroupBox) ConnectMouseMoveEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QGroupBox::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseMoveEvent", f)
	}
}

func (ptr *QGroupBox) DisconnectMouseMoveEvent() {
	defer qt.Recovering("disconnect QGroupBox::mouseMoveEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseMoveEvent")
	}
}

//export callbackQGroupBoxMouseMoveEvent
func callbackQGroupBoxMouseMoveEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGroupBox::mouseMoveEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseMoveEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QGroupBox) ConnectMousePressEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QGroupBox::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mousePressEvent", f)
	}
}

func (ptr *QGroupBox) DisconnectMousePressEvent() {
	defer qt.Recovering("disconnect QGroupBox::mousePressEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mousePressEvent")
	}
}

//export callbackQGroupBoxMousePressEvent
func callbackQGroupBoxMousePressEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGroupBox::mousePressEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mousePressEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QGroupBox) ConnectMouseReleaseEvent(f func(event *gui.QMouseEvent)) {
	defer qt.Recovering("connect QGroupBox::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "mouseReleaseEvent", f)
	}
}

func (ptr *QGroupBox) DisconnectMouseReleaseEvent() {
	defer qt.Recovering("disconnect QGroupBox::mouseReleaseEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "mouseReleaseEvent")
	}
}

//export callbackQGroupBoxMouseReleaseEvent
func callbackQGroupBoxMouseReleaseEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGroupBox::mouseReleaseEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "mouseReleaseEvent"); signal != nil {
		signal.(func(*gui.QMouseEvent))(gui.NewQMouseEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QGroupBox) ConnectPaintEvent(f func(event *gui.QPaintEvent)) {
	defer qt.Recovering("connect QGroupBox::paintEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "paintEvent", f)
	}
}

func (ptr *QGroupBox) DisconnectPaintEvent() {
	defer qt.Recovering("disconnect QGroupBox::paintEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "paintEvent")
	}
}

//export callbackQGroupBoxPaintEvent
func callbackQGroupBoxPaintEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QGroupBox::paintEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "paintEvent"); signal != nil {
		signal.(func(*gui.QPaintEvent))(gui.NewQPaintEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QGroupBox) ConnectResizeEvent(f func(e *gui.QResizeEvent)) {
	defer qt.Recovering("connect QGroupBox::resizeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "resizeEvent", f)
	}
}

func (ptr *QGroupBox) DisconnectResizeEvent() {
	defer qt.Recovering("disconnect QGroupBox::resizeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "resizeEvent")
	}
}

//export callbackQGroupBoxResizeEvent
func callbackQGroupBoxResizeEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QGroupBox::resizeEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "resizeEvent"); signal != nil {
		signal.(func(*gui.QResizeEvent))(gui.NewQResizeEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QGroupBox) ConnectToggled(f func(on bool)) {
	defer qt.Recovering("connect QGroupBox::toggled")

	if ptr.Pointer() != nil {
		C.QGroupBox_ConnectToggled(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "toggled", f)
	}
}

func (ptr *QGroupBox) DisconnectToggled() {
	defer qt.Recovering("disconnect QGroupBox::toggled")

	if ptr.Pointer() != nil {
		C.QGroupBox_DisconnectToggled(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "toggled")
	}
}

//export callbackQGroupBoxToggled
func callbackQGroupBoxToggled(ptrName *C.char, on C.int) {
	defer qt.Recovering("callback QGroupBox::toggled")

	if signal := qt.GetSignal(C.GoString(ptrName), "toggled"); signal != nil {
		signal.(func(bool))(int(on) != 0)
	}

}

func (ptr *QGroupBox) DestroyQGroupBox() {
	defer qt.Recovering("QGroupBox::~QGroupBox")

	if ptr.Pointer() != nil {
		C.QGroupBox_DestroyQGroupBox(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
