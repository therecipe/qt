package gui

//#include "qopengldebugmessage.h"
import "C"
import (
	"unsafe"
)

type QOpenGLDebugMessage struct {
	ptr unsafe.Pointer
}

type QOpenGLDebugMessageITF interface {
	QOpenGLDebugMessagePTR() *QOpenGLDebugMessage
}

func (p *QOpenGLDebugMessage) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QOpenGLDebugMessage) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQOpenGLDebugMessage(ptr QOpenGLDebugMessageITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QOpenGLDebugMessagePTR().Pointer()
	}
	return nil
}

func QOpenGLDebugMessageFromPointer(ptr unsafe.Pointer) *QOpenGLDebugMessage {
	var n = new(QOpenGLDebugMessage)
	n.SetPointer(ptr)
	return n
}

func (ptr *QOpenGLDebugMessage) QOpenGLDebugMessagePTR() *QOpenGLDebugMessage {
	return ptr
}

//QOpenGLDebugMessage::Severity
type QOpenGLDebugMessage__Severity int

var (
	QOpenGLDebugMessage__InvalidSeverity      = QOpenGLDebugMessage__Severity(0x00000000)
	QOpenGLDebugMessage__HighSeverity         = QOpenGLDebugMessage__Severity(0x00000001)
	QOpenGLDebugMessage__MediumSeverity       = QOpenGLDebugMessage__Severity(0x00000002)
	QOpenGLDebugMessage__LowSeverity          = QOpenGLDebugMessage__Severity(0x00000004)
	QOpenGLDebugMessage__NotificationSeverity = QOpenGLDebugMessage__Severity(0x00000008)
	QOpenGLDebugMessage__LastSeverity         = QOpenGLDebugMessage__Severity(QOpenGLDebugMessage__NotificationSeverity)
	QOpenGLDebugMessage__AnySeverity          = QOpenGLDebugMessage__Severity(0xffffffff)
)

//QOpenGLDebugMessage::Source
type QOpenGLDebugMessage__Source int

var (
	QOpenGLDebugMessage__InvalidSource        = QOpenGLDebugMessage__Source(0x00000000)
	QOpenGLDebugMessage__APISource            = QOpenGLDebugMessage__Source(0x00000001)
	QOpenGLDebugMessage__WindowSystemSource   = QOpenGLDebugMessage__Source(0x00000002)
	QOpenGLDebugMessage__ShaderCompilerSource = QOpenGLDebugMessage__Source(0x00000004)
	QOpenGLDebugMessage__ThirdPartySource     = QOpenGLDebugMessage__Source(0x00000008)
	QOpenGLDebugMessage__ApplicationSource    = QOpenGLDebugMessage__Source(0x00000010)
	QOpenGLDebugMessage__OtherSource          = QOpenGLDebugMessage__Source(0x00000020)
	QOpenGLDebugMessage__LastSource           = QOpenGLDebugMessage__Source(QOpenGLDebugMessage__OtherSource)
	QOpenGLDebugMessage__AnySource            = QOpenGLDebugMessage__Source(0xffffffff)
)

//QOpenGLDebugMessage::Type
type QOpenGLDebugMessage__Type int

var (
	QOpenGLDebugMessage__InvalidType            = QOpenGLDebugMessage__Type(0x00000000)
	QOpenGLDebugMessage__ErrorType              = QOpenGLDebugMessage__Type(0x00000001)
	QOpenGLDebugMessage__DeprecatedBehaviorType = QOpenGLDebugMessage__Type(0x00000002)
	QOpenGLDebugMessage__UndefinedBehaviorType  = QOpenGLDebugMessage__Type(0x00000004)
	QOpenGLDebugMessage__PortabilityType        = QOpenGLDebugMessage__Type(0x00000008)
	QOpenGLDebugMessage__PerformanceType        = QOpenGLDebugMessage__Type(0x00000010)
	QOpenGLDebugMessage__OtherType              = QOpenGLDebugMessage__Type(0x00000020)
	QOpenGLDebugMessage__MarkerType             = QOpenGLDebugMessage__Type(0x00000040)
	QOpenGLDebugMessage__GroupPushType          = QOpenGLDebugMessage__Type(0x00000080)
	QOpenGLDebugMessage__GroupPopType           = QOpenGLDebugMessage__Type(0x00000100)
	QOpenGLDebugMessage__LastType               = QOpenGLDebugMessage__Type(QOpenGLDebugMessage__GroupPopType)
	QOpenGLDebugMessage__AnyType                = QOpenGLDebugMessage__Type(0xffffffff)
)

func NewQOpenGLDebugMessage() *QOpenGLDebugMessage {
	return QOpenGLDebugMessageFromPointer(unsafe.Pointer(C.QOpenGLDebugMessage_NewQOpenGLDebugMessage()))
}

func NewQOpenGLDebugMessage2(debugMessage QOpenGLDebugMessageITF) *QOpenGLDebugMessage {
	return QOpenGLDebugMessageFromPointer(unsafe.Pointer(C.QOpenGLDebugMessage_NewQOpenGLDebugMessage2(C.QtObjectPtr(PointerFromQOpenGLDebugMessage(debugMessage)))))
}

func (ptr *QOpenGLDebugMessage) Message() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QOpenGLDebugMessage_Message(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QOpenGLDebugMessage) Severity() QOpenGLDebugMessage__Severity {
	if ptr.Pointer() != nil {
		return QOpenGLDebugMessage__Severity(C.QOpenGLDebugMessage_Severity(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QOpenGLDebugMessage) Source() QOpenGLDebugMessage__Source {
	if ptr.Pointer() != nil {
		return QOpenGLDebugMessage__Source(C.QOpenGLDebugMessage_Source(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QOpenGLDebugMessage) Swap(debugMessage QOpenGLDebugMessageITF) {
	if ptr.Pointer() != nil {
		C.QOpenGLDebugMessage_Swap(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQOpenGLDebugMessage(debugMessage)))
	}
}

func (ptr *QOpenGLDebugMessage) Type() QOpenGLDebugMessage__Type {
	if ptr.Pointer() != nil {
		return QOpenGLDebugMessage__Type(C.QOpenGLDebugMessage_Type(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QOpenGLDebugMessage) DestroyQOpenGLDebugMessage() {
	if ptr.Pointer() != nil {
		C.QOpenGLDebugMessage_DestroyQOpenGLDebugMessage(C.QtObjectPtr(ptr.Pointer()))
	}
}
