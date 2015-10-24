package core

//#include "qcommandlineoption.h"
import "C"
import (
	"strings"
	"unsafe"
)

type QCommandLineOption struct {
	ptr unsafe.Pointer
}

type QCommandLineOptionITF interface {
	QCommandLineOptionPTR() *QCommandLineOption
}

func (p *QCommandLineOption) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QCommandLineOption) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQCommandLineOption(ptr QCommandLineOptionITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCommandLineOptionPTR().Pointer()
	}
	return nil
}

func QCommandLineOptionFromPointer(ptr unsafe.Pointer) *QCommandLineOption {
	var n = new(QCommandLineOption)
	n.SetPointer(ptr)
	return n
}

func (ptr *QCommandLineOption) QCommandLineOptionPTR() *QCommandLineOption {
	return ptr
}

func NewQCommandLineOption5(other QCommandLineOptionITF) *QCommandLineOption {
	return QCommandLineOptionFromPointer(unsafe.Pointer(C.QCommandLineOption_NewQCommandLineOption5(C.QtObjectPtr(PointerFromQCommandLineOption(other)))))
}

func NewQCommandLineOption(name string) *QCommandLineOption {
	return QCommandLineOptionFromPointer(unsafe.Pointer(C.QCommandLineOption_NewQCommandLineOption(C.CString(name))))
}

func NewQCommandLineOption3(name string, description string, valueName string, defaultValue string) *QCommandLineOption {
	return QCommandLineOptionFromPointer(unsafe.Pointer(C.QCommandLineOption_NewQCommandLineOption3(C.CString(name), C.CString(description), C.CString(valueName), C.CString(defaultValue))))
}

func NewQCommandLineOption2(names []string) *QCommandLineOption {
	return QCommandLineOptionFromPointer(unsafe.Pointer(C.QCommandLineOption_NewQCommandLineOption2(C.CString(strings.Join(names, "|")))))
}

func NewQCommandLineOption4(names []string, description string, valueName string, defaultValue string) *QCommandLineOption {
	return QCommandLineOptionFromPointer(unsafe.Pointer(C.QCommandLineOption_NewQCommandLineOption4(C.CString(strings.Join(names, "|")), C.CString(description), C.CString(valueName), C.CString(defaultValue))))
}

func (ptr *QCommandLineOption) DefaultValues() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QCommandLineOption_DefaultValues(C.QtObjectPtr(ptr.Pointer()))), "|")
	}
	return make([]string, 0)
}

func (ptr *QCommandLineOption) Description() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QCommandLineOption_Description(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QCommandLineOption) Names() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QCommandLineOption_Names(C.QtObjectPtr(ptr.Pointer()))), "|")
	}
	return make([]string, 0)
}

func (ptr *QCommandLineOption) SetDefaultValue(defaultValue string) {
	if ptr.Pointer() != nil {
		C.QCommandLineOption_SetDefaultValue(C.QtObjectPtr(ptr.Pointer()), C.CString(defaultValue))
	}
}

func (ptr *QCommandLineOption) SetDefaultValues(defaultValues []string) {
	if ptr.Pointer() != nil {
		C.QCommandLineOption_SetDefaultValues(C.QtObjectPtr(ptr.Pointer()), C.CString(strings.Join(defaultValues, "|")))
	}
}

func (ptr *QCommandLineOption) SetDescription(description string) {
	if ptr.Pointer() != nil {
		C.QCommandLineOption_SetDescription(C.QtObjectPtr(ptr.Pointer()), C.CString(description))
	}
}

func (ptr *QCommandLineOption) SetValueName(valueName string) {
	if ptr.Pointer() != nil {
		C.QCommandLineOption_SetValueName(C.QtObjectPtr(ptr.Pointer()), C.CString(valueName))
	}
}

func (ptr *QCommandLineOption) Swap(other QCommandLineOptionITF) {
	if ptr.Pointer() != nil {
		C.QCommandLineOption_Swap(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQCommandLineOption(other)))
	}
}

func (ptr *QCommandLineOption) ValueName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QCommandLineOption_ValueName(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QCommandLineOption) DestroyQCommandLineOption() {
	if ptr.Pointer() != nil {
		C.QCommandLineOption_DestroyQCommandLineOption(C.QtObjectPtr(ptr.Pointer()))
	}
}
