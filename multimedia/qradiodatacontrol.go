package multimedia

//#include "qradiodatacontrol.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QRadioDataControl struct {
	QMediaControl
}

type QRadioDataControlITF interface {
	QMediaControlITF
	QRadioDataControlPTR() *QRadioDataControl
}

func PointerFromQRadioDataControl(ptr QRadioDataControlITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRadioDataControlPTR().Pointer()
	}
	return nil
}

func QRadioDataControlFromPointer(ptr unsafe.Pointer) *QRadioDataControl {
	var n = new(QRadioDataControl)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QRadioDataControl_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QRadioDataControl) QRadioDataControlPTR() *QRadioDataControl {
	return ptr
}

func (ptr *QRadioDataControl) ConnectAlternativeFrequenciesEnabledChanged(f func(enabled bool)) {
	if ptr.Pointer() != nil {
		C.QRadioDataControl_ConnectAlternativeFrequenciesEnabledChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "alternativeFrequenciesEnabledChanged", f)
	}
}

func (ptr *QRadioDataControl) DisconnectAlternativeFrequenciesEnabledChanged() {
	if ptr.Pointer() != nil {
		C.QRadioDataControl_DisconnectAlternativeFrequenciesEnabledChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "alternativeFrequenciesEnabledChanged")
	}
}

//export callbackQRadioDataControlAlternativeFrequenciesEnabledChanged
func callbackQRadioDataControlAlternativeFrequenciesEnabledChanged(ptrName *C.char, enabled C.int) {
	qt.GetSignal(C.GoString(ptrName), "alternativeFrequenciesEnabledChanged").(func(bool))(int(enabled) != 0)
}

func (ptr *QRadioDataControl) Error() QRadioData__Error {
	if ptr.Pointer() != nil {
		return QRadioData__Error(C.QRadioDataControl_Error(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QRadioDataControl) ErrorString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QRadioDataControl_ErrorString(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QRadioDataControl) IsAlternativeFrequenciesEnabled() bool {
	if ptr.Pointer() != nil {
		return C.QRadioDataControl_IsAlternativeFrequenciesEnabled(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QRadioDataControl) ProgramType() QRadioData__ProgramType {
	if ptr.Pointer() != nil {
		return QRadioData__ProgramType(C.QRadioDataControl_ProgramType(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QRadioDataControl) ConnectProgramTypeChanged(f func(programType QRadioData__ProgramType)) {
	if ptr.Pointer() != nil {
		C.QRadioDataControl_ConnectProgramTypeChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "programTypeChanged", f)
	}
}

func (ptr *QRadioDataControl) DisconnectProgramTypeChanged() {
	if ptr.Pointer() != nil {
		C.QRadioDataControl_DisconnectProgramTypeChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "programTypeChanged")
	}
}

//export callbackQRadioDataControlProgramTypeChanged
func callbackQRadioDataControlProgramTypeChanged(ptrName *C.char, programType C.int) {
	qt.GetSignal(C.GoString(ptrName), "programTypeChanged").(func(QRadioData__ProgramType))(QRadioData__ProgramType(programType))
}

func (ptr *QRadioDataControl) ProgramTypeName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QRadioDataControl_ProgramTypeName(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QRadioDataControl) ConnectProgramTypeNameChanged(f func(programTypeName string)) {
	if ptr.Pointer() != nil {
		C.QRadioDataControl_ConnectProgramTypeNameChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "programTypeNameChanged", f)
	}
}

func (ptr *QRadioDataControl) DisconnectProgramTypeNameChanged() {
	if ptr.Pointer() != nil {
		C.QRadioDataControl_DisconnectProgramTypeNameChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "programTypeNameChanged")
	}
}

//export callbackQRadioDataControlProgramTypeNameChanged
func callbackQRadioDataControlProgramTypeNameChanged(ptrName *C.char, programTypeName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "programTypeNameChanged").(func(string))(C.GoString(programTypeName))
}

func (ptr *QRadioDataControl) RadioText() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QRadioDataControl_RadioText(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QRadioDataControl) ConnectRadioTextChanged(f func(radioText string)) {
	if ptr.Pointer() != nil {
		C.QRadioDataControl_ConnectRadioTextChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "radioTextChanged", f)
	}
}

func (ptr *QRadioDataControl) DisconnectRadioTextChanged() {
	if ptr.Pointer() != nil {
		C.QRadioDataControl_DisconnectRadioTextChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "radioTextChanged")
	}
}

//export callbackQRadioDataControlRadioTextChanged
func callbackQRadioDataControlRadioTextChanged(ptrName *C.char, radioText *C.char) {
	qt.GetSignal(C.GoString(ptrName), "radioTextChanged").(func(string))(C.GoString(radioText))
}

func (ptr *QRadioDataControl) SetAlternativeFrequenciesEnabled(enabled bool) {
	if ptr.Pointer() != nil {
		C.QRadioDataControl_SetAlternativeFrequenciesEnabled(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QRadioDataControl) StationId() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QRadioDataControl_StationId(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QRadioDataControl) ConnectStationIdChanged(f func(stationId string)) {
	if ptr.Pointer() != nil {
		C.QRadioDataControl_ConnectStationIdChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "stationIdChanged", f)
	}
}

func (ptr *QRadioDataControl) DisconnectStationIdChanged() {
	if ptr.Pointer() != nil {
		C.QRadioDataControl_DisconnectStationIdChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "stationIdChanged")
	}
}

//export callbackQRadioDataControlStationIdChanged
func callbackQRadioDataControlStationIdChanged(ptrName *C.char, stationId *C.char) {
	qt.GetSignal(C.GoString(ptrName), "stationIdChanged").(func(string))(C.GoString(stationId))
}

func (ptr *QRadioDataControl) StationName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QRadioDataControl_StationName(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QRadioDataControl) ConnectStationNameChanged(f func(stationName string)) {
	if ptr.Pointer() != nil {
		C.QRadioDataControl_ConnectStationNameChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "stationNameChanged", f)
	}
}

func (ptr *QRadioDataControl) DisconnectStationNameChanged() {
	if ptr.Pointer() != nil {
		C.QRadioDataControl_DisconnectStationNameChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "stationNameChanged")
	}
}

//export callbackQRadioDataControlStationNameChanged
func callbackQRadioDataControlStationNameChanged(ptrName *C.char, stationName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "stationNameChanged").(func(string))(C.GoString(stationName))
}

func (ptr *QRadioDataControl) DestroyQRadioDataControl() {
	if ptr.Pointer() != nil {
		C.QRadioDataControl_DestroyQRadioDataControl(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
