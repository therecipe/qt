package qt

//#include "qlabel.h"
import "C"

type qlabel struct {
	qframe
}

type QLabel interface {
	QFrame
	Alignment() AlignmentFlag
	Buddy() QWidget
	HasScaledContents() bool
	HasSelectedText() bool
	Indent() int
	Margin() int
	OpenExternalLinks() bool
	SelectedText() string
	SelectionStart() int
	SetAlignment(alignment AlignmentFlag)
	SetBuddy(buddy QWidget)
	SetIndent(indent int)
	SetMargin(margin int)
	SetOpenExternalLinks(open bool)
	SetScaledContents(scaledContents bool)
	SetSelection(start int, length int)
	SetTextFormat(textFormat TextFormat)
	SetTextInteractionFlags(flags TextInteractionFlag)
	SetWordWrap(on bool)
	Text() string
	TextFormat() TextFormat
	TextInteractionFlags() TextInteractionFlag
	WordWrap() bool
	ConnectSlotClear()
	DisconnectSlotClear()
	SlotClear()
	ConnectSlotSetText()
	DisconnectSlotSetText()
	SlotSetText(text string)
	ConnectSignalLinkActivated(f func())
	DisconnectSignalLinkActivated()
	SignalLinkActivated() func()
	ConnectSignalLinkHovered(f func())
	DisconnectSignalLinkHovered()
	SignalLinkHovered() func()
}

func (p *qlabel) Pointer() (ptr C.QtObjectPtr) {
	return p.ptr
}

func (p *qlabel) SetPointer(ptr C.QtObjectPtr) {
	p.ptr = ptr
}

func NewQLabel1(parent QWidget, f WindowType) QLabel {
	var parentPtr C.QtObjectPtr
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qlabel = new(qlabel)
	qlabel.SetPointer(C.QLabel_New_QWidget_WindowType(parentPtr, C.int(f)))
	qlabel.SetObjectName("QLabel_" + randomIdentifier())
	return qlabel
}

func NewQLabel2(text string, parent QWidget, f WindowType) QLabel {
	var parentPtr C.QtObjectPtr
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qlabel = new(qlabel)
	qlabel.SetPointer(C.QLabel_New_String_QWidget_WindowType(C.CString(text), parentPtr, C.int(f)))
	qlabel.SetObjectName("QLabel_" + randomIdentifier())
	return qlabel
}

func (p *qlabel) Destroy() {
	if p.Pointer() != nil {
		getSignal(p.ObjectName(), "destroyed")()
		C.QLabel_Destroy(p.Pointer())
		p.SetPointer(nil)
	}
}

func (p *qlabel) Alignment() AlignmentFlag {
	if p.Pointer() == nil {
		return 0
	}
	return AlignmentFlag(C.QLabel_Alignment(p.Pointer()))
}

func (p *qlabel) Buddy() QWidget {
	if p.Pointer() == nil {
		return nil
	} else {
		var qwidget = new(qwidget)
		qwidget.SetPointer(C.QLabel_Buddy(p.Pointer()))
		if qwidget.ObjectName() == "" {
			qwidget.SetObjectName("QWidget_" + randomIdentifier())
		}
		return qwidget
	}
}

func (p *qlabel) HasScaledContents() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QLabel_HasScaledContents(p.Pointer()) != 0
}

func (p *qlabel) HasSelectedText() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QLabel_HasSelectedText(p.Pointer()) != 0
}

func (p *qlabel) Indent() int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QLabel_Indent(p.Pointer()))
}

func (p *qlabel) Margin() int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QLabel_Margin(p.Pointer()))
}

func (p *qlabel) OpenExternalLinks() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QLabel_OpenExternalLinks(p.Pointer()) != 0
}

func (p *qlabel) SelectedText() string {
	if p.Pointer() == nil {
		return ""
	}
	return C.GoString(C.QLabel_SelectedText(p.Pointer()))
}

func (p *qlabel) SelectionStart() int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QLabel_SelectionStart(p.Pointer()))
}

func (p *qlabel) SetAlignment(alignment AlignmentFlag) {
	if p.Pointer() != nil {
		C.QLabel_SetAlignment_AlignmentFlag(p.Pointer(), C.int(alignment))
	}
}

func (p *qlabel) SetBuddy(buddy QWidget) {
	if p.Pointer() != nil {
		var buddyPtr C.QtObjectPtr
		if buddy != nil {
			buddyPtr = buddy.Pointer()
		}
		C.QLabel_SetBuddy_QWidget(p.Pointer(), buddyPtr)
	}
}

func (p *qlabel) SetIndent(indent int) {
	if p.Pointer() != nil {
		C.QLabel_SetIndent_Int(p.Pointer(), C.int(indent))
	}
}

func (p *qlabel) SetMargin(margin int) {
	if p.Pointer() != nil {
		C.QLabel_SetMargin_Int(p.Pointer(), C.int(margin))
	}
}

func (p *qlabel) SetOpenExternalLinks(open bool) {
	if p.Pointer() != nil {
		C.QLabel_SetOpenExternalLinks_Bool(p.Pointer(), goBoolToCInt(open))
	}
}

func (p *qlabel) SetScaledContents(scaledContents bool) {
	if p.Pointer() != nil {
		C.QLabel_SetScaledContents_Bool(p.Pointer(), goBoolToCInt(scaledContents))
	}
}

func (p *qlabel) SetSelection(start int, length int) {
	if p.Pointer() != nil {
		C.QLabel_SetSelection_Int_Int(p.Pointer(), C.int(start), C.int(length))
	}
}

func (p *qlabel) SetTextFormat(textFormat TextFormat) {
	if p.Pointer() != nil {
		C.QLabel_SetTextFormat_TextFormat(p.Pointer(), C.int(textFormat))
	}
}

func (p *qlabel) SetTextInteractionFlags(flags TextInteractionFlag) {
	if p.Pointer() != nil {
		C.QLabel_SetTextInteractionFlags_TextInteractionFlag(p.Pointer(), C.int(flags))
	}
}

func (p *qlabel) SetWordWrap(on bool) {
	if p.Pointer() != nil {
		C.QLabel_SetWordWrap_Bool(p.Pointer(), goBoolToCInt(on))
	}
}

func (p *qlabel) Text() string {
	if p.Pointer() == nil {
		return ""
	}
	return C.GoString(C.QLabel_Text(p.Pointer()))
}

func (p *qlabel) TextFormat() TextFormat {
	if p.Pointer() == nil {
		return 0
	}
	return TextFormat(C.QLabel_TextFormat(p.Pointer()))
}

func (p *qlabel) TextInteractionFlags() TextInteractionFlag {
	if p.Pointer() == nil {
		return 0
	}
	return TextInteractionFlag(C.QLabel_TextInteractionFlags(p.Pointer()))
}

func (p *qlabel) WordWrap() bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QLabel_WordWrap(p.Pointer()) != 0
}

func (p *qlabel) ConnectSlotClear() {
	C.QLabel_ConnectSlotClear(p.Pointer())
}

func (p *qlabel) DisconnectSlotClear() {
	C.QLabel_DisconnectSlotClear(p.Pointer())
}

func (p *qlabel) SlotClear() {
	if p.Pointer() != nil {
		C.QLabel_Clear(p.Pointer())
	}
}

func (p *qlabel) ConnectSlotSetText() {
	C.QLabel_ConnectSlotSetText(p.Pointer())
}

func (p *qlabel) DisconnectSlotSetText() {
	C.QLabel_DisconnectSlotSetText(p.Pointer())
}

func (p *qlabel) SlotSetText(text string) {
	if p.Pointer() != nil {
		C.QLabel_SetText_String(p.Pointer(), C.CString(text))
	}
}

func (p *qlabel) ConnectSignalLinkActivated(f func()) {
	C.QLabel_ConnectSignalLinkActivated(p.Pointer())
	connectSignal(p.ObjectName(), "linkActivated", f)
}

func (p *qlabel) DisconnectSignalLinkActivated() {
	C.QLabel_DisconnectSignalLinkActivated(p.Pointer())
	disconnectSignal(p.ObjectName(), "linkActivated")
}

func (p *qlabel) SignalLinkActivated() func() {
	return getSignal(p.ObjectName(), "linkActivated")
}

func (p *qlabel) ConnectSignalLinkHovered(f func()) {
	C.QLabel_ConnectSignalLinkHovered(p.Pointer())
	connectSignal(p.ObjectName(), "linkHovered", f)
}

func (p *qlabel) DisconnectSignalLinkHovered() {
	C.QLabel_DisconnectSignalLinkHovered(p.Pointer())
	disconnectSignal(p.ObjectName(), "linkHovered")
}

func (p *qlabel) SignalLinkHovered() func() {
	return getSignal(p.ObjectName(), "linkHovered")
}

//export callbackQLabel
func callbackQLabel(ptr C.QtObjectPtr, msg *C.char) {
	var qlabel = new(qlabel)
	qlabel.SetPointer(ptr)
	getSignal(qlabel.ObjectName(), C.GoString(msg))()
}
