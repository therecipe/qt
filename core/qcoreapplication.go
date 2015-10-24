package core

//#include "qcoreapplication.h"
import "C"
import (
	"github.com/therecipe/qt"
	"strings"
	"unsafe"
)

type QCoreApplication struct {
	QObject
}

type QCoreApplicationITF interface {
	QObjectITF
	QCoreApplicationPTR() *QCoreApplication
}

func PointerFromQCoreApplication(ptr QCoreApplicationITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCoreApplicationPTR().Pointer()
	}
	return nil
}

func QCoreApplicationFromPointer(ptr unsafe.Pointer) *QCoreApplication {
	var n = new(QCoreApplication)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QCoreApplication_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QCoreApplication) QCoreApplicationPTR() *QCoreApplication {
	return ptr
}

func QCoreApplication_ApplicationName() string {
	return C.GoString(C.QCoreApplication_QCoreApplication_ApplicationName())
}

func QCoreApplication_ApplicationVersion() string {
	return C.GoString(C.QCoreApplication_QCoreApplication_ApplicationVersion())
}

func QCoreApplication_OrganizationDomain() string {
	return C.GoString(C.QCoreApplication_QCoreApplication_OrganizationDomain())
}

func QCoreApplication_OrganizationName() string {
	return C.GoString(C.QCoreApplication_QCoreApplication_OrganizationName())
}

func QCoreApplication_SetApplicationName(application string) {
	C.QCoreApplication_QCoreApplication_SetApplicationName(C.CString(application))
}

func QCoreApplication_SetApplicationVersion(version string) {
	C.QCoreApplication_QCoreApplication_SetApplicationVersion(C.CString(version))
}

func QCoreApplication_SetOrganizationDomain(orgDomain string) {
	C.QCoreApplication_QCoreApplication_SetOrganizationDomain(C.CString(orgDomain))
}

func QCoreApplication_SetOrganizationName(orgName string) {
	C.QCoreApplication_QCoreApplication_SetOrganizationName(C.CString(orgName))
}

func NewQCoreApplication(argc int, argv string) *QCoreApplication {
	return QCoreApplicationFromPointer(unsafe.Pointer(C.QCoreApplication_NewQCoreApplication(C.int(argc), C.CString(argv))))
}

func (ptr *QCoreApplication) ConnectAboutToQuit(f func()) {
	if ptr.Pointer() != nil {
		C.QCoreApplication_ConnectAboutToQuit(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "aboutToQuit", f)
	}
}

func (ptr *QCoreApplication) DisconnectAboutToQuit() {
	if ptr.Pointer() != nil {
		C.QCoreApplication_DisconnectAboutToQuit(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "aboutToQuit")
	}
}

//export callbackQCoreApplicationAboutToQuit
func callbackQCoreApplicationAboutToQuit(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "aboutToQuit").(func())()
}

func QCoreApplication_AddLibraryPath(path string) {
	C.QCoreApplication_QCoreApplication_AddLibraryPath(C.CString(path))
}

func QCoreApplication_ApplicationDirPath() string {
	return C.GoString(C.QCoreApplication_QCoreApplication_ApplicationDirPath())
}

func QCoreApplication_ApplicationFilePath() string {
	return C.GoString(C.QCoreApplication_QCoreApplication_ApplicationFilePath())
}

func QCoreApplication_Arguments() []string {
	return strings.Split(C.GoString(C.QCoreApplication_QCoreApplication_Arguments()), "|")
}

func QCoreApplication_ClosingDown() bool {
	return C.QCoreApplication_QCoreApplication_ClosingDown() != 0
}

func QCoreApplication_EventDispatcher() *QAbstractEventDispatcher {
	return QAbstractEventDispatcherFromPointer(unsafe.Pointer(C.QCoreApplication_QCoreApplication_EventDispatcher()))
}

func QCoreApplication_Exec() int {
	return int(C.QCoreApplication_QCoreApplication_Exec())
}

func QCoreApplication_Exit(returnCode int) {
	C.QCoreApplication_QCoreApplication_Exit(C.int(returnCode))
}

func QCoreApplication_Flush() {
	C.QCoreApplication_QCoreApplication_Flush()
}

func (ptr *QCoreApplication) InstallNativeEventFilter(filterObj QAbstractNativeEventFilterITF) {
	if ptr.Pointer() != nil {
		C.QCoreApplication_InstallNativeEventFilter(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQAbstractNativeEventFilter(filterObj)))
	}
}

func QCoreApplication_InstallTranslator(translationFile QTranslatorITF) bool {
	return C.QCoreApplication_QCoreApplication_InstallTranslator(C.QtObjectPtr(PointerFromQTranslator(translationFile))) != 0
}

func QCoreApplication_Instance() *QCoreApplication {
	return QCoreApplicationFromPointer(unsafe.Pointer(C.QCoreApplication_QCoreApplication_Instance()))
}

func QCoreApplication_IsQuitLockEnabled() bool {
	return C.QCoreApplication_QCoreApplication_IsQuitLockEnabled() != 0
}

func QCoreApplication_IsSetuidAllowed() bool {
	return C.QCoreApplication_QCoreApplication_IsSetuidAllowed() != 0
}

func QCoreApplication_LibraryPaths() []string {
	return strings.Split(C.GoString(C.QCoreApplication_QCoreApplication_LibraryPaths()), "|")
}

func (ptr *QCoreApplication) Notify(receiver QObjectITF, event QEventITF) bool {
	if ptr.Pointer() != nil {
		return C.QCoreApplication_Notify(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQObject(receiver)), C.QtObjectPtr(PointerFromQEvent(event))) != 0
	}
	return false
}

func QCoreApplication_PostEvent(receiver QObjectITF, event QEventITF, priority int) {
	C.QCoreApplication_QCoreApplication_PostEvent(C.QtObjectPtr(PointerFromQObject(receiver)), C.QtObjectPtr(PointerFromQEvent(event)), C.int(priority))
}

func QCoreApplication_ProcessEvents(flags QEventLoop__ProcessEventsFlag) {
	C.QCoreApplication_QCoreApplication_ProcessEvents(C.int(flags))
}

func QCoreApplication_ProcessEvents2(flags QEventLoop__ProcessEventsFlag, maxtime int) {
	C.QCoreApplication_QCoreApplication_ProcessEvents2(C.int(flags), C.int(maxtime))
}

func QCoreApplication_Quit() {
	C.QCoreApplication_QCoreApplication_Quit()
}

func QCoreApplication_RemoveLibraryPath(path string) {
	C.QCoreApplication_QCoreApplication_RemoveLibraryPath(C.CString(path))
}

func (ptr *QCoreApplication) RemoveNativeEventFilter(filterObject QAbstractNativeEventFilterITF) {
	if ptr.Pointer() != nil {
		C.QCoreApplication_RemoveNativeEventFilter(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQAbstractNativeEventFilter(filterObject)))
	}
}

func QCoreApplication_RemovePostedEvents(receiver QObjectITF, eventType int) {
	C.QCoreApplication_QCoreApplication_RemovePostedEvents(C.QtObjectPtr(PointerFromQObject(receiver)), C.int(eventType))
}

func QCoreApplication_RemoveTranslator(translationFile QTranslatorITF) bool {
	return C.QCoreApplication_QCoreApplication_RemoveTranslator(C.QtObjectPtr(PointerFromQTranslator(translationFile))) != 0
}

func QCoreApplication_SendEvent(receiver QObjectITF, event QEventITF) bool {
	return C.QCoreApplication_QCoreApplication_SendEvent(C.QtObjectPtr(PointerFromQObject(receiver)), C.QtObjectPtr(PointerFromQEvent(event))) != 0
}

func QCoreApplication_SendPostedEvents(receiver QObjectITF, event_type int) {
	C.QCoreApplication_QCoreApplication_SendPostedEvents(C.QtObjectPtr(PointerFromQObject(receiver)), C.int(event_type))
}

func QCoreApplication_SetAttribute(attribute Qt__ApplicationAttribute, on bool) {
	C.QCoreApplication_QCoreApplication_SetAttribute(C.int(attribute), C.int(qt.GoBoolToInt(on)))
}

func QCoreApplication_SetEventDispatcher(eventDispatcher QAbstractEventDispatcherITF) {
	C.QCoreApplication_QCoreApplication_SetEventDispatcher(C.QtObjectPtr(PointerFromQAbstractEventDispatcher(eventDispatcher)))
}

func QCoreApplication_SetLibraryPaths(paths []string) {
	C.QCoreApplication_QCoreApplication_SetLibraryPaths(C.CString(strings.Join(paths, "|")))
}

func QCoreApplication_SetQuitLockEnabled(enabled bool) {
	C.QCoreApplication_QCoreApplication_SetQuitLockEnabled(C.int(qt.GoBoolToInt(enabled)))
}

func QCoreApplication_SetSetuidAllowed(allow bool) {
	C.QCoreApplication_QCoreApplication_SetSetuidAllowed(C.int(qt.GoBoolToInt(allow)))
}

func QCoreApplication_StartingUp() bool {
	return C.QCoreApplication_QCoreApplication_StartingUp() != 0
}

func QCoreApplication_TestAttribute(attribute Qt__ApplicationAttribute) bool {
	return C.QCoreApplication_QCoreApplication_TestAttribute(C.int(attribute)) != 0
}

func QCoreApplication_Translate(context string, sourceText string, disambiguation string, n int) string {
	return C.GoString(C.QCoreApplication_QCoreApplication_Translate(C.CString(context), C.CString(sourceText), C.CString(disambiguation), C.int(n)))
}

func (ptr *QCoreApplication) DestroyQCoreApplication() {
	if ptr.Pointer() != nil {
		C.QCoreApplication_DestroyQCoreApplication(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
