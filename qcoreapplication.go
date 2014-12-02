package qt

//#include "qcoreapplication.h"
import "C"

type qcoreapplication struct {
	qobject
}

type QCoreApplication interface {
	QObject
}

func (p *qcoreapplication) Pointer() (ptr C.QtObjectPtr) {
	return p.ptr
}

func (p *qcoreapplication) SetPointer(ptr C.QtObjectPtr) {
	p.ptr = ptr
}

func NewQCoreApplication_Int_String(argc int, argv string) QCoreApplication {
	var qcoreapplication = new(qcoreapplication)
	qcoreapplication.SetPointer(C.QCoreApplication_New_Int_String(C.int(argc), C.CString(argv)))
	qcoreapplication.SetObjectName_String("QCoreApplication_" + randomIdentifier())
	return qcoreapplication
}

func (p *qcoreapplication) Destroy() {
	if p.Pointer() != nil {
		getSignal(p.ObjectName(), "destroyed")()
		C.QCoreApplication_Destroy(p.Pointer())
		p.SetPointer(nil)
	}
}

func QCoreApplication_AddLibraryPath_String(path string) {
	C.QCoreApplication_AddLibraryPath_String(C.CString(path))
}

func QCoreApplication_ApplicationDirPath() string {
	return C.GoString(C.QCoreApplication_ApplicationDirPath())
}

func QCoreApplication_ApplicationFilePath() string {
	return C.GoString(C.QCoreApplication_ApplicationFilePath())
}

func QCoreApplication_ApplicationName() string {
	return C.GoString(C.QCoreApplication_ApplicationName())
}

func QCoreApplication_ApplicationVersion() string {
	return C.GoString(C.QCoreApplication_ApplicationVersion())
}

func QCoreApplication_ClosingDown() bool {
	return C.QCoreApplication_ClosingDown() != 0
}

func QCoreApplication_Exec() int {
	return int(C.QCoreApplication_Exec())
}

func QCoreApplication_Exit_Int(returnCode int) {
	C.QCoreApplication_Exit_Int(C.int(returnCode))
}

func QCoreApplication_Instance() QCoreApplication {
	var qcoreapplication = new(qcoreapplication)
	qcoreapplication.SetPointer(C.QCoreApplication_Instance())
	if qcoreapplication.ObjectName() == "" {
		qcoreapplication.SetObjectName_String("QCoreApplication_" + randomIdentifier())
	}
	return qcoreapplication
}

func QCoreApplication_IsQuitLockEnabled() bool {
	return C.QCoreApplication_IsQuitLockEnabled() != 0
}

func QCoreApplication_IsSetuidAllowed() bool {
	return C.QCoreApplication_IsSetuidAllowed() != 0
}

func QCoreApplication_OrganizationDomain() string {
	return C.GoString(C.QCoreApplication_OrganizationDomain())
}

func QCoreApplication_OrganizationName() string {
	return C.GoString(C.QCoreApplication_OrganizationName())
}

func QCoreApplication_RemoveLibraryPath_String(path string) {
	C.QCoreApplication_RemoveLibraryPath_String(C.CString(path))
}

func QCoreApplication_RemovePostedEvents_QObject_Int(receiver QObject, eventType int) {
	var receiverPtr C.QtObjectPtr = nil
	if receiver != nil {
		receiverPtr = receiver.Pointer()
	}
	C.QCoreApplication_RemovePostedEvents_QObject_Int(receiverPtr, C.int(eventType))
}

func QCoreApplication_SendPostedEvents_QObject_Int(receiver QObject, event_type int) {
	var receiverPtr C.QtObjectPtr = nil
	if receiver != nil {
		receiverPtr = receiver.Pointer()
	}
	C.QCoreApplication_SendPostedEvents_QObject_Int(receiverPtr, C.int(event_type))
}

func QCoreApplication_SetApplicationName_String(application string) {
	C.QCoreApplication_SetApplicationName_String(C.CString(application))
}

func QCoreApplication_SetApplicationVersion_String(version string) {
	C.QCoreApplication_SetApplicationVersion_String(C.CString(version))
}

func QCoreApplication_SetAttribute_ApplicationAttribute_Bool(attribute ApplicationAttribute, on bool) {
	C.QCoreApplication_SetAttribute_ApplicationAttribute_Bool(C.int(attribute), goBoolToCInt(on))
}

func QCoreApplication_SetOrganizationDomain_String(orgDomain string) {
	C.QCoreApplication_SetOrganizationDomain_String(C.CString(orgDomain))
}

func QCoreApplication_SetOrganizationName_String(orgName string) {
	C.QCoreApplication_SetOrganizationName_String(C.CString(orgName))
}

func QCoreApplication_SetQuitLockEnabled_Bool(enabled bool) {
	C.QCoreApplication_SetQuitLockEnabled_Bool(goBoolToCInt(enabled))
}

func QCoreApplication_SetSetuidAllowed_Bool(allow bool) {
	C.QCoreApplication_SetSetuidAllowed_Bool(goBoolToCInt(allow))
}

func QCoreApplication_StartingUp() bool {
	return C.QCoreApplication_StartingUp() != 0
}

func QCoreApplication_TestAttribute_ApplicationAttribute(attribute ApplicationAttribute) bool {
	return C.QCoreApplication_TestAttribute_ApplicationAttribute(C.int(attribute)) != 0
}

func QCoreApplication_Translate_String_String_String_Int(context string, sourceText string, disambiguation string, n int) string {
	return C.GoString(C.QCoreApplication_Translate_String_String_String_Int(C.CString(context), C.CString(sourceText), C.CString(disambiguation), C.int(n)))
}
