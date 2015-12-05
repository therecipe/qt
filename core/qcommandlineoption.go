package core

//#include "core.h"
import "C"
import (
	"log"
	"strings"
	"unsafe"
)

type QCommandLineOption struct {
	ptr unsafe.Pointer
}

type QCommandLineOption_ITF interface {
	QCommandLineOption_PTR() *QCommandLineOption
}

func (p *QCommandLineOption) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QCommandLineOption) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQCommandLineOption(ptr QCommandLineOption_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCommandLineOption_PTR().Pointer()
	}
	return nil
}

func NewQCommandLineOptionFromPointer(ptr unsafe.Pointer) *QCommandLineOption {
	var n = new(QCommandLineOption)
	n.SetPointer(ptr)
	return n
}

func (ptr *QCommandLineOption) QCommandLineOption_PTR() *QCommandLineOption {
	return ptr
}

func NewQCommandLineOption5(other QCommandLineOption_ITF) *QCommandLineOption {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCommandLineOption::QCommandLineOption")
		}
	}()

	return NewQCommandLineOptionFromPointer(C.QCommandLineOption_NewQCommandLineOption5(PointerFromQCommandLineOption(other)))
}

func NewQCommandLineOption(name string) *QCommandLineOption {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCommandLineOption::QCommandLineOption")
		}
	}()

	return NewQCommandLineOptionFromPointer(C.QCommandLineOption_NewQCommandLineOption(C.CString(name)))
}

func NewQCommandLineOption3(name string, description string, valueName string, defaultValue string) *QCommandLineOption {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCommandLineOption::QCommandLineOption")
		}
	}()

	return NewQCommandLineOptionFromPointer(C.QCommandLineOption_NewQCommandLineOption3(C.CString(name), C.CString(description), C.CString(valueName), C.CString(defaultValue)))
}

func NewQCommandLineOption2(names []string) *QCommandLineOption {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCommandLineOption::QCommandLineOption")
		}
	}()

	return NewQCommandLineOptionFromPointer(C.QCommandLineOption_NewQCommandLineOption2(C.CString(strings.Join(names, ",,,"))))
}

func NewQCommandLineOption4(names []string, description string, valueName string, defaultValue string) *QCommandLineOption {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCommandLineOption::QCommandLineOption")
		}
	}()

	return NewQCommandLineOptionFromPointer(C.QCommandLineOption_NewQCommandLineOption4(C.CString(strings.Join(names, ",,,")), C.CString(description), C.CString(valueName), C.CString(defaultValue)))
}

func (ptr *QCommandLineOption) DefaultValues() []string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCommandLineOption::defaultValues")
		}
	}()

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QCommandLineOption_DefaultValues(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QCommandLineOption) Description() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCommandLineOption::description")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QCommandLineOption_Description(ptr.Pointer()))
	}
	return ""
}

func (ptr *QCommandLineOption) Names() []string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCommandLineOption::names")
		}
	}()

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QCommandLineOption_Names(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QCommandLineOption) SetDefaultValue(defaultValue string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCommandLineOption::setDefaultValue")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCommandLineOption_SetDefaultValue(ptr.Pointer(), C.CString(defaultValue))
	}
}

func (ptr *QCommandLineOption) SetDefaultValues(defaultValues []string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCommandLineOption::setDefaultValues")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCommandLineOption_SetDefaultValues(ptr.Pointer(), C.CString(strings.Join(defaultValues, ",,,")))
	}
}

func (ptr *QCommandLineOption) SetDescription(description string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCommandLineOption::setDescription")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCommandLineOption_SetDescription(ptr.Pointer(), C.CString(description))
	}
}

func (ptr *QCommandLineOption) SetValueName(valueName string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCommandLineOption::setValueName")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCommandLineOption_SetValueName(ptr.Pointer(), C.CString(valueName))
	}
}

func (ptr *QCommandLineOption) Swap(other QCommandLineOption_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCommandLineOption::swap")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCommandLineOption_Swap(ptr.Pointer(), PointerFromQCommandLineOption(other))
	}
}

func (ptr *QCommandLineOption) ValueName() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCommandLineOption::valueName")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QCommandLineOption_ValueName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QCommandLineOption) DestroyQCommandLineOption() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCommandLineOption::~QCommandLineOption")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCommandLineOption_DestroyQCommandLineOption(ptr.Pointer())
	}
}
