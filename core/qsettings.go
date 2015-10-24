package core

//#include "qsettings.h"
import "C"
import (
	"github.com/therecipe/qt"
	"strings"
	"unsafe"
)

type QSettings struct {
	QObject
}

type QSettingsITF interface {
	QObjectITF
	QSettingsPTR() *QSettings
}

func PointerFromQSettings(ptr QSettingsITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSettingsPTR().Pointer()
	}
	return nil
}

func QSettingsFromPointer(ptr unsafe.Pointer) *QSettings {
	var n = new(QSettings)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QSettings_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QSettings) QSettingsPTR() *QSettings {
	return ptr
}

//QSettings::Format
type QSettings__Format int

var (
	QSettings__NativeFormat   = QSettings__Format(0)
	QSettings__IniFormat      = QSettings__Format(1)
	QSettings__InvalidFormat  = QSettings__Format(16)
	QSettings__CustomFormat1  = QSettings__Format(17)
	QSettings__CustomFormat2  = QSettings__Format(18)
	QSettings__CustomFormat3  = QSettings__Format(19)
	QSettings__CustomFormat4  = QSettings__Format(20)
	QSettings__CustomFormat5  = QSettings__Format(21)
	QSettings__CustomFormat6  = QSettings__Format(22)
	QSettings__CustomFormat7  = QSettings__Format(23)
	QSettings__CustomFormat8  = QSettings__Format(24)
	QSettings__CustomFormat9  = QSettings__Format(25)
	QSettings__CustomFormat10 = QSettings__Format(26)
	QSettings__CustomFormat11 = QSettings__Format(27)
	QSettings__CustomFormat12 = QSettings__Format(28)
	QSettings__CustomFormat13 = QSettings__Format(29)
	QSettings__CustomFormat14 = QSettings__Format(30)
	QSettings__CustomFormat15 = QSettings__Format(31)
	QSettings__CustomFormat16 = QSettings__Format(32)
)

//QSettings::Scope
type QSettings__Scope int

var (
	QSettings__UserScope   = QSettings__Scope(0)
	QSettings__SystemScope = QSettings__Scope(1)
)

//QSettings::Status
type QSettings__Status int

var (
	QSettings__NoError     = QSettings__Status(0)
	QSettings__AccessError = QSettings__Status(1)
	QSettings__FormatError = QSettings__Status(2)
)

func NewQSettings3(format QSettings__Format, scope QSettings__Scope, organization string, application string, parent QObjectITF) *QSettings {
	return QSettingsFromPointer(unsafe.Pointer(C.QSettings_NewQSettings3(C.int(format), C.int(scope), C.CString(organization), C.CString(application), C.QtObjectPtr(PointerFromQObject(parent)))))
}

func NewQSettings5(parent QObjectITF) *QSettings {
	return QSettingsFromPointer(unsafe.Pointer(C.QSettings_NewQSettings5(C.QtObjectPtr(PointerFromQObject(parent)))))
}

func NewQSettings2(scope QSettings__Scope, organization string, application string, parent QObjectITF) *QSettings {
	return QSettingsFromPointer(unsafe.Pointer(C.QSettings_NewQSettings2(C.int(scope), C.CString(organization), C.CString(application), C.QtObjectPtr(PointerFromQObject(parent)))))
}

func NewQSettings4(fileName string, format QSettings__Format, parent QObjectITF) *QSettings {
	return QSettingsFromPointer(unsafe.Pointer(C.QSettings_NewQSettings4(C.CString(fileName), C.int(format), C.QtObjectPtr(PointerFromQObject(parent)))))
}

func NewQSettings(organization string, application string, parent QObjectITF) *QSettings {
	return QSettingsFromPointer(unsafe.Pointer(C.QSettings_NewQSettings(C.CString(organization), C.CString(application), C.QtObjectPtr(PointerFromQObject(parent)))))
}

func (ptr *QSettings) AllKeys() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QSettings_AllKeys(C.QtObjectPtr(ptr.Pointer()))), "|")
	}
	return make([]string, 0)
}

func (ptr *QSettings) ApplicationName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSettings_ApplicationName(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QSettings) BeginGroup(prefix string) {
	if ptr.Pointer() != nil {
		C.QSettings_BeginGroup(C.QtObjectPtr(ptr.Pointer()), C.CString(prefix))
	}
}

func (ptr *QSettings) BeginReadArray(prefix string) int {
	if ptr.Pointer() != nil {
		return int(C.QSettings_BeginReadArray(C.QtObjectPtr(ptr.Pointer()), C.CString(prefix)))
	}
	return 0
}

func (ptr *QSettings) BeginWriteArray(prefix string, size int) {
	if ptr.Pointer() != nil {
		C.QSettings_BeginWriteArray(C.QtObjectPtr(ptr.Pointer()), C.CString(prefix), C.int(size))
	}
}

func (ptr *QSettings) ChildGroups() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QSettings_ChildGroups(C.QtObjectPtr(ptr.Pointer()))), "|")
	}
	return make([]string, 0)
}

func (ptr *QSettings) ChildKeys() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QSettings_ChildKeys(C.QtObjectPtr(ptr.Pointer()))), "|")
	}
	return make([]string, 0)
}

func (ptr *QSettings) Clear() {
	if ptr.Pointer() != nil {
		C.QSettings_Clear(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QSettings) Contains(key string) bool {
	if ptr.Pointer() != nil {
		return C.QSettings_Contains(C.QtObjectPtr(ptr.Pointer()), C.CString(key)) != 0
	}
	return false
}

func QSettings_DefaultFormat() QSettings__Format {
	return QSettings__Format(C.QSettings_QSettings_DefaultFormat())
}

func (ptr *QSettings) EndArray() {
	if ptr.Pointer() != nil {
		C.QSettings_EndArray(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QSettings) EndGroup() {
	if ptr.Pointer() != nil {
		C.QSettings_EndGroup(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QSettings) FallbacksEnabled() bool {
	if ptr.Pointer() != nil {
		return C.QSettings_FallbacksEnabled(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSettings) FileName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSettings_FileName(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QSettings) Format() QSettings__Format {
	if ptr.Pointer() != nil {
		return QSettings__Format(C.QSettings_Format(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSettings) Group() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSettings_Group(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QSettings) IniCodec() *QTextCodec {
	if ptr.Pointer() != nil {
		return QTextCodecFromPointer(unsafe.Pointer(C.QSettings_IniCodec(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QSettings) IsWritable() bool {
	if ptr.Pointer() != nil {
		return C.QSettings_IsWritable(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSettings) OrganizationName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSettings_OrganizationName(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QSettings) Scope() QSettings__Scope {
	if ptr.Pointer() != nil {
		return QSettings__Scope(C.QSettings_Scope(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSettings) SetArrayIndex(i int) {
	if ptr.Pointer() != nil {
		C.QSettings_SetArrayIndex(C.QtObjectPtr(ptr.Pointer()), C.int(i))
	}
}

func QSettings_SetDefaultFormat(format QSettings__Format) {
	C.QSettings_QSettings_SetDefaultFormat(C.int(format))
}

func (ptr *QSettings) SetFallbacksEnabled(b bool) {
	if ptr.Pointer() != nil {
		C.QSettings_SetFallbacksEnabled(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(b)))
	}
}

func (ptr *QSettings) SetIniCodec(codec QTextCodecITF) {
	if ptr.Pointer() != nil {
		C.QSettings_SetIniCodec(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTextCodec(codec)))
	}
}

func (ptr *QSettings) SetIniCodec2(codecName string) {
	if ptr.Pointer() != nil {
		C.QSettings_SetIniCodec2(C.QtObjectPtr(ptr.Pointer()), C.CString(codecName))
	}
}

func QSettings_SetPath(format QSettings__Format, scope QSettings__Scope, path string) {
	C.QSettings_QSettings_SetPath(C.int(format), C.int(scope), C.CString(path))
}

func (ptr *QSettings) SetValue(key string, value string) {
	if ptr.Pointer() != nil {
		C.QSettings_SetValue(C.QtObjectPtr(ptr.Pointer()), C.CString(key), C.CString(value))
	}
}

func (ptr *QSettings) Status() QSettings__Status {
	if ptr.Pointer() != nil {
		return QSettings__Status(C.QSettings_Status(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSettings) Sync() {
	if ptr.Pointer() != nil {
		C.QSettings_Sync(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QSettings) Value(key string, defaultValue string) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSettings_Value(C.QtObjectPtr(ptr.Pointer()), C.CString(key), C.CString(defaultValue)))
	}
	return ""
}

func (ptr *QSettings) DestroyQSettings() {
	if ptr.Pointer() != nil {
		C.QSettings_DestroyQSettings(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
