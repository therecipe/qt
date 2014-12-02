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
	SetAlignment_AlignmentFlag(alignment AlignmentFlag)
	SetBuddy_QWidget(buddy QWidget)
	SetIndent_Int(indent int)
	SetMargin_Int(margin int)
	SetOpenExternalLinks_Bool(open bool)
	SetScaledContents_Bool(scaledContents bool)
	SetSelection_Int_Int(start int, length int)
	SetTextFormat_TextFormat(textFormat TextFormat)
	SetTextInteractionFlags_TextInteractionFlag(flags TextInteractionFlag)
	SetWordWrap_Bool(on bool)
	Text() string
	TextFormat() TextFormat
	TextInteractionFlags() TextInteractionFlag
	WordWrap() bool
	ConnectSlotClear()
	DisconnectSlotClear()
	SlotClear()
	ConnectSlotSetText()
	DisconnectSlotSetText()
	SlotSetText_String(text string)
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

func NewQLabel_QWidget_WindowType(parent QWidget, f WindowType) QLabel {
	var parentPtr C.QtObjectPtr = nil
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qlabel = new(qlabel)
	qlabel.SetPointer(C.QLabel_New_QWidget_WindowType(parentPtr, C.int(f)))
	qlabel.SetObjectName_String("QLabel_" + randomIdentifier())
	return qlabel
}

func NewQLabel_String_QWidget_WindowType(text string, parent QWidget, f WindowType) QLabel {
	var parentPtr C.QtObjectPtr = nil
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qlabel = new(qlabel)
	qlabel.SetPointer(C.QLabel_New_String_QWidget_WindowType(C.CString(text), parentPtr, C.int(f)))
	qlabel.SetObjectName_String("QLabel_" + randomIdentifier())
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
	} else {
		return AlignmentFlag(C.QLabel_Alignment(p.Pointer()))
	}
}

func (p *qlabel) Buddy() QWidget {
	if p.Pointer() == nil {
		return nil
	} else {
		var qwidget = new(qwidget)
		qwidget.SetPointer(C.QLabel_Buddy(p.Pointer()))
		if qwidget.ObjectName() == "" {
			qwidget.SetObjectName_String("QWidget_" + randomIdentifier())
		}
		return qwidget
	}
}

func (p *qlabel) HasScaledContents() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QLabel_HasScaledContents(p.Pointer()) != 0
	}
}

func (p *qlabel) HasSelectedText() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QLabel_HasSelectedText(p.Pointer()) != 0
	}
}

func (p *qlabel) Indent() int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QLabel_Indent(p.Pointer()))
	}
}

func (p *qlabel) Margin() int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QLabel_Margin(p.Pointer()))
	}
}

func (p *qlabel) OpenExternalLinks() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QLabel_OpenExternalLinks(p.Pointer()) != 0
	}
}

func (p *qlabel) SelectedText() string {
	if p.Pointer() == nil {
		return ""
	} else {
		return C.GoString(C.QLabel_SelectedText(p.Pointer()))
	}
}

func (p *qlabel) SelectionStart() int {
	if p.Pointer() == nil {
		return 0
	} else {
		return int(C.QLabel_SelectionStart(p.Pointer()))
	}
}

func (p *qlabel) SetAlignment_AlignmentFlag(alignment AlignmentFlag) {
	if p.Pointer() != nil {
		C.QLabel_SetAlignment_AlignmentFlag(p.Pointer(), C.int(alignment))
	}
}

func (p *qlabel) SetBuddy_QWidget(buddy QWidget) {
	if p.Pointer() == nil {
	} else {
		var buddyPtr C.QtObjectPtr = nil
		if buddy != nil {
			buddyPtr = buddy.Pointer()
		}
		C.QLabel_SetBuddy_QWidget(p.Pointer(), buddyPtr)
	}
}

func (p *qlabel) SetIndent_Int(indent int) {
	if p.Pointer() != nil {
		C.QLabel_SetIndent_Int(p.Pointer(), C.int(indent))
	}
}

func (p *qlabel) SetMargin_Int(margin int) {
	if p.Pointer() != nil {
		C.QLabel_SetMargin_Int(p.Pointer(), C.int(margin))
	}
}

func (p *qlabel) SetOpenExternalLinks_Bool(open bool) {
	if p.Pointer() != nil {
		C.QLabel_SetOpenExternalLinks_Bool(p.Pointer(), goBoolToCInt(open))
	}
}

func (p *qlabel) SetScaledContents_Bool(scaledContents bool) {
	if p.Pointer() != nil {
		C.QLabel_SetScaledContents_Bool(p.Pointer(), goBoolToCInt(scaledContents))
	}
}

func (p *qlabel) SetSelection_Int_Int(start int, length int) {
	if p.Pointer() != nil {
		C.QLabel_SetSelection_Int_Int(p.Pointer(), C.int(start), C.int(length))
	}
}

func (p *qlabel) SetTextFormat_TextFormat(textFormat TextFormat) {
	if p.Pointer() != nil {
		C.QLabel_SetTextFormat_TextFormat(p.Pointer(), C.int(textFormat))
	}
}

func (p *qlabel) SetTextInteractionFlags_TextInteractionFlag(flags TextInteractionFlag) {
	if p.Pointer() != nil {
		C.QLabel_SetTextInteractionFlags_TextInteractionFlag(p.Pointer(), C.int(flags))
	}
}

func (p *qlabel) SetWordWrap_Bool(on bool) {
	if p.Pointer() != nil {
		C.QLabel_SetWordWrap_Bool(p.Pointer(), goBoolToCInt(on))
	}
}

func (p *qlabel) Text() string {
	if p.Pointer() == nil {
		return ""
	} else {
		return C.GoString(C.QLabel_Text(p.Pointer()))
	}
}

func (p *qlabel) TextFormat() TextFormat {
	if p.Pointer() == nil {
		return 0
	} else {
		return TextFormat(C.QLabel_TextFormat(p.Pointer()))
	}
}

func (p *qlabel) TextInteractionFlags() TextInteractionFlag {
	if p.Pointer() == nil {
		return 0
	} else {
		return TextInteractionFlag(C.QLabel_TextInteractionFlags(p.Pointer()))
	}
}

func (p *qlabel) WordWrap() bool {
	if p.Pointer() == nil {
		return false
	} else {
		return C.QLabel_WordWrap(p.Pointer()) != 0
	}
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

func (p *qlabel) SlotSetText_String(text string) {
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
