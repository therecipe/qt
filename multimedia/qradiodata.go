package multimedia

//#include "qradiodata.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QRadioData struct {
	core.QObject
	QMediaBindableInterface
}

type QRadioData_ITF interface {
	core.QObject_ITF
	QMediaBindableInterface_ITF
	QRadioData_PTR() *QRadioData
}

func (p *QRadioData) Pointer() unsafe.Pointer {
	return p.QObject_PTR().Pointer()
}

func (p *QRadioData) SetPointer(ptr unsafe.Pointer) {
	p.QObject_PTR().SetPointer(ptr)
	p.QMediaBindableInterface_PTR().SetPointer(ptr)
}

func PointerFromQRadioData(ptr QRadioData_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRadioData_PTR().Pointer()
	}
	return nil
}

func NewQRadioDataFromPointer(ptr unsafe.Pointer) *QRadioData {
	var n = new(QRadioData)
	n.SetPointer(ptr)
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QRadioData_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QRadioData) QRadioData_PTR() *QRadioData {
	return ptr
}

//QRadioData::Error
type QRadioData__Error int64

const (
	QRadioData__NoError         = QRadioData__Error(0)
	QRadioData__ResourceError   = QRadioData__Error(1)
	QRadioData__OpenError       = QRadioData__Error(2)
	QRadioData__OutOfRangeError = QRadioData__Error(3)
)

//QRadioData::ProgramType
type QRadioData__ProgramType int64

const (
	QRadioData__Undefined           = QRadioData__ProgramType(0)
	QRadioData__News                = QRadioData__ProgramType(1)
	QRadioData__CurrentAffairs      = QRadioData__ProgramType(2)
	QRadioData__Information         = QRadioData__ProgramType(3)
	QRadioData__Sport               = QRadioData__ProgramType(4)
	QRadioData__Education           = QRadioData__ProgramType(5)
	QRadioData__Drama               = QRadioData__ProgramType(6)
	QRadioData__Culture             = QRadioData__ProgramType(7)
	QRadioData__Science             = QRadioData__ProgramType(8)
	QRadioData__Varied              = QRadioData__ProgramType(9)
	QRadioData__PopMusic            = QRadioData__ProgramType(10)
	QRadioData__RockMusic           = QRadioData__ProgramType(11)
	QRadioData__EasyListening       = QRadioData__ProgramType(12)
	QRadioData__LightClassical      = QRadioData__ProgramType(13)
	QRadioData__SeriousClassical    = QRadioData__ProgramType(14)
	QRadioData__OtherMusic          = QRadioData__ProgramType(15)
	QRadioData__Weather             = QRadioData__ProgramType(16)
	QRadioData__Finance             = QRadioData__ProgramType(17)
	QRadioData__ChildrensProgrammes = QRadioData__ProgramType(18)
	QRadioData__SocialAffairs       = QRadioData__ProgramType(19)
	QRadioData__Religion            = QRadioData__ProgramType(20)
	QRadioData__PhoneIn             = QRadioData__ProgramType(21)
	QRadioData__Travel              = QRadioData__ProgramType(22)
	QRadioData__Leisure             = QRadioData__ProgramType(23)
	QRadioData__JazzMusic           = QRadioData__ProgramType(24)
	QRadioData__CountryMusic        = QRadioData__ProgramType(25)
	QRadioData__NationalMusic       = QRadioData__ProgramType(26)
	QRadioData__OldiesMusic         = QRadioData__ProgramType(27)
	QRadioData__FolkMusic           = QRadioData__ProgramType(28)
	QRadioData__Documentary         = QRadioData__ProgramType(29)
	QRadioData__AlarmTest           = QRadioData__ProgramType(30)
	QRadioData__Alarm               = QRadioData__ProgramType(31)
	QRadioData__Talk                = QRadioData__ProgramType(32)
	QRadioData__ClassicRock         = QRadioData__ProgramType(33)
	QRadioData__AdultHits           = QRadioData__ProgramType(34)
	QRadioData__SoftRock            = QRadioData__ProgramType(35)
	QRadioData__Top40               = QRadioData__ProgramType(36)
	QRadioData__Soft                = QRadioData__ProgramType(37)
	QRadioData__Nostalgia           = QRadioData__ProgramType(38)
	QRadioData__Classical           = QRadioData__ProgramType(39)
	QRadioData__RhythmAndBlues      = QRadioData__ProgramType(40)
	QRadioData__SoftRhythmAndBlues  = QRadioData__ProgramType(41)
	QRadioData__Language            = QRadioData__ProgramType(42)
	QRadioData__ReligiousMusic      = QRadioData__ProgramType(43)
	QRadioData__ReligiousTalk       = QRadioData__ProgramType(44)
	QRadioData__Personality         = QRadioData__ProgramType(45)
	QRadioData__Public              = QRadioData__ProgramType(46)
	QRadioData__College             = QRadioData__ProgramType(47)
)

func (ptr *QRadioData) IsAlternativeFrequenciesEnabled() bool {
	if ptr.Pointer() != nil {
		return C.QRadioData_IsAlternativeFrequenciesEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QRadioData) ProgramType() QRadioData__ProgramType {
	if ptr.Pointer() != nil {
		return QRadioData__ProgramType(C.QRadioData_ProgramType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRadioData) ProgramTypeName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QRadioData_ProgramTypeName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QRadioData) RadioText() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QRadioData_RadioText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QRadioData) SetAlternativeFrequenciesEnabled(enabled bool) {
	if ptr.Pointer() != nil {
		C.QRadioData_SetAlternativeFrequenciesEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QRadioData) StationId() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QRadioData_StationId(ptr.Pointer()))
	}
	return ""
}

func (ptr *QRadioData) StationName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QRadioData_StationName(ptr.Pointer()))
	}
	return ""
}

func NewQRadioData(mediaObject QMediaObject_ITF, parent core.QObject_ITF) *QRadioData {
	return NewQRadioDataFromPointer(C.QRadioData_NewQRadioData(PointerFromQMediaObject(mediaObject), core.PointerFromQObject(parent)))
}

func (ptr *QRadioData) ConnectAlternativeFrequenciesEnabledChanged(f func(enabled bool)) {
	if ptr.Pointer() != nil {
		C.QRadioData_ConnectAlternativeFrequenciesEnabledChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "alternativeFrequenciesEnabledChanged", f)
	}
}

func (ptr *QRadioData) DisconnectAlternativeFrequenciesEnabledChanged() {
	if ptr.Pointer() != nil {
		C.QRadioData_DisconnectAlternativeFrequenciesEnabledChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "alternativeFrequenciesEnabledChanged")
	}
}

//export callbackQRadioDataAlternativeFrequenciesEnabledChanged
func callbackQRadioDataAlternativeFrequenciesEnabledChanged(ptrName *C.char, enabled C.int) {
	qt.GetSignal(C.GoString(ptrName), "alternativeFrequenciesEnabledChanged").(func(bool))(int(enabled) != 0)
}

func (ptr *QRadioData) Error() QRadioData__Error {
	if ptr.Pointer() != nil {
		return QRadioData__Error(C.QRadioData_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRadioData) ErrorString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QRadioData_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QRadioData) MediaObject() *QMediaObject {
	if ptr.Pointer() != nil {
		return NewQMediaObjectFromPointer(C.QRadioData_MediaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QRadioData) ConnectProgramTypeChanged(f func(programType QRadioData__ProgramType)) {
	if ptr.Pointer() != nil {
		C.QRadioData_ConnectProgramTypeChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "programTypeChanged", f)
	}
}

func (ptr *QRadioData) DisconnectProgramTypeChanged() {
	if ptr.Pointer() != nil {
		C.QRadioData_DisconnectProgramTypeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "programTypeChanged")
	}
}

//export callbackQRadioDataProgramTypeChanged
func callbackQRadioDataProgramTypeChanged(ptrName *C.char, programType C.int) {
	qt.GetSignal(C.GoString(ptrName), "programTypeChanged").(func(QRadioData__ProgramType))(QRadioData__ProgramType(programType))
}

func (ptr *QRadioData) ConnectProgramTypeNameChanged(f func(programTypeName string)) {
	if ptr.Pointer() != nil {
		C.QRadioData_ConnectProgramTypeNameChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "programTypeNameChanged", f)
	}
}

func (ptr *QRadioData) DisconnectProgramTypeNameChanged() {
	if ptr.Pointer() != nil {
		C.QRadioData_DisconnectProgramTypeNameChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "programTypeNameChanged")
	}
}

//export callbackQRadioDataProgramTypeNameChanged
func callbackQRadioDataProgramTypeNameChanged(ptrName *C.char, programTypeName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "programTypeNameChanged").(func(string))(C.GoString(programTypeName))
}

func (ptr *QRadioData) ConnectRadioTextChanged(f func(radioText string)) {
	if ptr.Pointer() != nil {
		C.QRadioData_ConnectRadioTextChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "radioTextChanged", f)
	}
}

func (ptr *QRadioData) DisconnectRadioTextChanged() {
	if ptr.Pointer() != nil {
		C.QRadioData_DisconnectRadioTextChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "radioTextChanged")
	}
}

//export callbackQRadioDataRadioTextChanged
func callbackQRadioDataRadioTextChanged(ptrName *C.char, radioText *C.char) {
	qt.GetSignal(C.GoString(ptrName), "radioTextChanged").(func(string))(C.GoString(radioText))
}

func (ptr *QRadioData) ConnectStationIdChanged(f func(stationId string)) {
	if ptr.Pointer() != nil {
		C.QRadioData_ConnectStationIdChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stationIdChanged", f)
	}
}

func (ptr *QRadioData) DisconnectStationIdChanged() {
	if ptr.Pointer() != nil {
		C.QRadioData_DisconnectStationIdChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stationIdChanged")
	}
}

//export callbackQRadioDataStationIdChanged
func callbackQRadioDataStationIdChanged(ptrName *C.char, stationId *C.char) {
	qt.GetSignal(C.GoString(ptrName), "stationIdChanged").(func(string))(C.GoString(stationId))
}

func (ptr *QRadioData) ConnectStationNameChanged(f func(stationName string)) {
	if ptr.Pointer() != nil {
		C.QRadioData_ConnectStationNameChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stationNameChanged", f)
	}
}

func (ptr *QRadioData) DisconnectStationNameChanged() {
	if ptr.Pointer() != nil {
		C.QRadioData_DisconnectStationNameChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stationNameChanged")
	}
}

//export callbackQRadioDataStationNameChanged
func callbackQRadioDataStationNameChanged(ptrName *C.char, stationName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "stationNameChanged").(func(string))(C.GoString(stationName))
}

func (ptr *QRadioData) DestroyQRadioData() {
	if ptr.Pointer() != nil {
		C.QRadioData_DestroyQRadioData(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
