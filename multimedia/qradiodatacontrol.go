package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QRadioDataControl struct {
	QMediaControl
}

type QRadioDataControl_ITF interface {
	QMediaControl_ITF
	QRadioDataControl_PTR() *QRadioDataControl
}

func PointerFromQRadioDataControl(ptr QRadioDataControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRadioDataControl_PTR().Pointer()
	}
	return nil
}

func NewQRadioDataControlFromPointer(ptr unsafe.Pointer) *QRadioDataControl {
	var n = new(QRadioDataControl)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QRadioDataControl_") {
		n.SetObjectName("QRadioDataControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QRadioDataControl) QRadioDataControl_PTR() *QRadioDataControl {
	return ptr
}

func (ptr *QRadioDataControl) ConnectAlternativeFrequenciesEnabledChanged(f func(enabled bool)) {
	defer qt.Recovering("connect QRadioDataControl::alternativeFrequenciesEnabledChanged")

	if ptr.Pointer() != nil {
		C.QRadioDataControl_ConnectAlternativeFrequenciesEnabledChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "alternativeFrequenciesEnabledChanged", f)
	}
}

func (ptr *QRadioDataControl) DisconnectAlternativeFrequenciesEnabledChanged() {
	defer qt.Recovering("disconnect QRadioDataControl::alternativeFrequenciesEnabledChanged")

	if ptr.Pointer() != nil {
		C.QRadioDataControl_DisconnectAlternativeFrequenciesEnabledChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "alternativeFrequenciesEnabledChanged")
	}
}

//export callbackQRadioDataControlAlternativeFrequenciesEnabledChanged
func callbackQRadioDataControlAlternativeFrequenciesEnabledChanged(ptrName *C.char, enabled C.int) {
	defer qt.Recovering("callback QRadioDataControl::alternativeFrequenciesEnabledChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "alternativeFrequenciesEnabledChanged"); signal != nil {
		signal.(func(bool))(int(enabled) != 0)
	}

}

func (ptr *QRadioDataControl) ConnectError2(f func(error QRadioData__Error)) {
	defer qt.Recovering("connect QRadioDataControl::error")

	if ptr.Pointer() != nil {
		C.QRadioDataControl_ConnectError2(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error2", f)
	}
}

func (ptr *QRadioDataControl) DisconnectError2() {
	defer qt.Recovering("disconnect QRadioDataControl::error")

	if ptr.Pointer() != nil {
		C.QRadioDataControl_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error2")
	}
}

//export callbackQRadioDataControlError2
func callbackQRadioDataControlError2(ptrName *C.char, error C.int) {
	defer qt.Recovering("callback QRadioDataControl::error")

	if signal := qt.GetSignal(C.GoString(ptrName), "error2"); signal != nil {
		signal.(func(QRadioData__Error))(QRadioData__Error(error))
	}

}

func (ptr *QRadioDataControl) Error() QRadioData__Error {
	defer qt.Recovering("QRadioDataControl::error")

	if ptr.Pointer() != nil {
		return QRadioData__Error(C.QRadioDataControl_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRadioDataControl) ErrorString() string {
	defer qt.Recovering("QRadioDataControl::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QRadioDataControl_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QRadioDataControl) IsAlternativeFrequenciesEnabled() bool {
	defer qt.Recovering("QRadioDataControl::isAlternativeFrequenciesEnabled")

	if ptr.Pointer() != nil {
		return C.QRadioDataControl_IsAlternativeFrequenciesEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QRadioDataControl) ProgramType() QRadioData__ProgramType {
	defer qt.Recovering("QRadioDataControl::programType")

	if ptr.Pointer() != nil {
		return QRadioData__ProgramType(C.QRadioDataControl_ProgramType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRadioDataControl) ConnectProgramTypeChanged(f func(programType QRadioData__ProgramType)) {
	defer qt.Recovering("connect QRadioDataControl::programTypeChanged")

	if ptr.Pointer() != nil {
		C.QRadioDataControl_ConnectProgramTypeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "programTypeChanged", f)
	}
}

func (ptr *QRadioDataControl) DisconnectProgramTypeChanged() {
	defer qt.Recovering("disconnect QRadioDataControl::programTypeChanged")

	if ptr.Pointer() != nil {
		C.QRadioDataControl_DisconnectProgramTypeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "programTypeChanged")
	}
}

//export callbackQRadioDataControlProgramTypeChanged
func callbackQRadioDataControlProgramTypeChanged(ptrName *C.char, programType C.int) {
	defer qt.Recovering("callback QRadioDataControl::programTypeChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "programTypeChanged"); signal != nil {
		signal.(func(QRadioData__ProgramType))(QRadioData__ProgramType(programType))
	}

}

func (ptr *QRadioDataControl) ProgramTypeName() string {
	defer qt.Recovering("QRadioDataControl::programTypeName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QRadioDataControl_ProgramTypeName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QRadioDataControl) ConnectProgramTypeNameChanged(f func(programTypeName string)) {
	defer qt.Recovering("connect QRadioDataControl::programTypeNameChanged")

	if ptr.Pointer() != nil {
		C.QRadioDataControl_ConnectProgramTypeNameChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "programTypeNameChanged", f)
	}
}

func (ptr *QRadioDataControl) DisconnectProgramTypeNameChanged() {
	defer qt.Recovering("disconnect QRadioDataControl::programTypeNameChanged")

	if ptr.Pointer() != nil {
		C.QRadioDataControl_DisconnectProgramTypeNameChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "programTypeNameChanged")
	}
}

//export callbackQRadioDataControlProgramTypeNameChanged
func callbackQRadioDataControlProgramTypeNameChanged(ptrName *C.char, programTypeName *C.char) {
	defer qt.Recovering("callback QRadioDataControl::programTypeNameChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "programTypeNameChanged"); signal != nil {
		signal.(func(string))(C.GoString(programTypeName))
	}

}

func (ptr *QRadioDataControl) RadioText() string {
	defer qt.Recovering("QRadioDataControl::radioText")

	if ptr.Pointer() != nil {
		return C.GoString(C.QRadioDataControl_RadioText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QRadioDataControl) ConnectRadioTextChanged(f func(radioText string)) {
	defer qt.Recovering("connect QRadioDataControl::radioTextChanged")

	if ptr.Pointer() != nil {
		C.QRadioDataControl_ConnectRadioTextChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "radioTextChanged", f)
	}
}

func (ptr *QRadioDataControl) DisconnectRadioTextChanged() {
	defer qt.Recovering("disconnect QRadioDataControl::radioTextChanged")

	if ptr.Pointer() != nil {
		C.QRadioDataControl_DisconnectRadioTextChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "radioTextChanged")
	}
}

//export callbackQRadioDataControlRadioTextChanged
func callbackQRadioDataControlRadioTextChanged(ptrName *C.char, radioText *C.char) {
	defer qt.Recovering("callback QRadioDataControl::radioTextChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "radioTextChanged"); signal != nil {
		signal.(func(string))(C.GoString(radioText))
	}

}

func (ptr *QRadioDataControl) SetAlternativeFrequenciesEnabled(enabled bool) {
	defer qt.Recovering("QRadioDataControl::setAlternativeFrequenciesEnabled")

	if ptr.Pointer() != nil {
		C.QRadioDataControl_SetAlternativeFrequenciesEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QRadioDataControl) StationId() string {
	defer qt.Recovering("QRadioDataControl::stationId")

	if ptr.Pointer() != nil {
		return C.GoString(C.QRadioDataControl_StationId(ptr.Pointer()))
	}
	return ""
}

func (ptr *QRadioDataControl) ConnectStationIdChanged(f func(stationId string)) {
	defer qt.Recovering("connect QRadioDataControl::stationIdChanged")

	if ptr.Pointer() != nil {
		C.QRadioDataControl_ConnectStationIdChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stationIdChanged", f)
	}
}

func (ptr *QRadioDataControl) DisconnectStationIdChanged() {
	defer qt.Recovering("disconnect QRadioDataControl::stationIdChanged")

	if ptr.Pointer() != nil {
		C.QRadioDataControl_DisconnectStationIdChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stationIdChanged")
	}
}

//export callbackQRadioDataControlStationIdChanged
func callbackQRadioDataControlStationIdChanged(ptrName *C.char, stationId *C.char) {
	defer qt.Recovering("callback QRadioDataControl::stationIdChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "stationIdChanged"); signal != nil {
		signal.(func(string))(C.GoString(stationId))
	}

}

func (ptr *QRadioDataControl) StationName() string {
	defer qt.Recovering("QRadioDataControl::stationName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QRadioDataControl_StationName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QRadioDataControl) ConnectStationNameChanged(f func(stationName string)) {
	defer qt.Recovering("connect QRadioDataControl::stationNameChanged")

	if ptr.Pointer() != nil {
		C.QRadioDataControl_ConnectStationNameChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stationNameChanged", f)
	}
}

func (ptr *QRadioDataControl) DisconnectStationNameChanged() {
	defer qt.Recovering("disconnect QRadioDataControl::stationNameChanged")

	if ptr.Pointer() != nil {
		C.QRadioDataControl_DisconnectStationNameChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stationNameChanged")
	}
}

//export callbackQRadioDataControlStationNameChanged
func callbackQRadioDataControlStationNameChanged(ptrName *C.char, stationName *C.char) {
	defer qt.Recovering("callback QRadioDataControl::stationNameChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "stationNameChanged"); signal != nil {
		signal.(func(string))(C.GoString(stationName))
	}

}

func (ptr *QRadioDataControl) DestroyQRadioDataControl() {
	defer qt.Recovering("QRadioDataControl::~QRadioDataControl")

	if ptr.Pointer() != nil {
		C.QRadioDataControl_DestroyQRadioDataControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
