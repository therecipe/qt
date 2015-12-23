package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"strings"
	"unsafe"
)

type QCoreApplication struct {
	QObject
}

type QCoreApplication_ITF interface {
	QObject_ITF
	QCoreApplication_PTR() *QCoreApplication
}

func PointerFromQCoreApplication(ptr QCoreApplication_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCoreApplication_PTR().Pointer()
	}
	return nil
}

func NewQCoreApplicationFromPointer(ptr unsafe.Pointer) *QCoreApplication {
	var n = new(QCoreApplication)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QCoreApplication_") {
		n.SetObjectName("QCoreApplication_" + qt.Identifier())
	}
	return n
}

func (ptr *QCoreApplication) QCoreApplication_PTR() *QCoreApplication {
	return ptr
}

func QCoreApplication_ApplicationName() string {
	defer qt.Recovering("QCoreApplication::applicationName")

	return C.GoString(C.QCoreApplication_QCoreApplication_ApplicationName())
}

func QCoreApplication_ApplicationVersion() string {
	defer qt.Recovering("QCoreApplication::applicationVersion")

	return C.GoString(C.QCoreApplication_QCoreApplication_ApplicationVersion())
}

func QCoreApplication_OrganizationDomain() string {
	defer qt.Recovering("QCoreApplication::organizationDomain")

	return C.GoString(C.QCoreApplication_QCoreApplication_OrganizationDomain())
}

func QCoreApplication_OrganizationName() string {
	defer qt.Recovering("QCoreApplication::organizationName")

	return C.GoString(C.QCoreApplication_QCoreApplication_OrganizationName())
}

func QCoreApplication_SetApplicationName(application string) {
	defer qt.Recovering("QCoreApplication::setApplicationName")

	C.QCoreApplication_QCoreApplication_SetApplicationName(C.CString(application))
}

func QCoreApplication_SetApplicationVersion(version string) {
	defer qt.Recovering("QCoreApplication::setApplicationVersion")

	C.QCoreApplication_QCoreApplication_SetApplicationVersion(C.CString(version))
}

func QCoreApplication_SetOrganizationDomain(orgDomain string) {
	defer qt.Recovering("QCoreApplication::setOrganizationDomain")

	C.QCoreApplication_QCoreApplication_SetOrganizationDomain(C.CString(orgDomain))
}

func QCoreApplication_SetOrganizationName(orgName string) {
	defer qt.Recovering("QCoreApplication::setOrganizationName")

	C.QCoreApplication_QCoreApplication_SetOrganizationName(C.CString(orgName))
}

func NewQCoreApplication(argc int, argv []string) *QCoreApplication {
	defer qt.Recovering("QCoreApplication::QCoreApplication")

	return NewQCoreApplicationFromPointer(C.QCoreApplication_NewQCoreApplication(C.int(argc), C.CString(strings.Join(argv, ",,,"))))
}

func (ptr *QCoreApplication) ConnectAboutToQuit(f func()) {
	defer qt.Recovering("connect QCoreApplication::aboutToQuit")

	if ptr.Pointer() != nil {
		C.QCoreApplication_ConnectAboutToQuit(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "aboutToQuit", f)
	}
}

func (ptr *QCoreApplication) DisconnectAboutToQuit() {
	defer qt.Recovering("disconnect QCoreApplication::aboutToQuit")

	if ptr.Pointer() != nil {
		C.QCoreApplication_DisconnectAboutToQuit(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "aboutToQuit")
	}
}

//export callbackQCoreApplicationAboutToQuit
func callbackQCoreApplicationAboutToQuit(ptrName *C.char) {
	defer qt.Recovering("callback QCoreApplication::aboutToQuit")

	if signal := qt.GetSignal(C.GoString(ptrName), "aboutToQuit"); signal != nil {
		signal.(func())()
	}

}

func QCoreApplication_AddLibraryPath(path string) {
	defer qt.Recovering("QCoreApplication::addLibraryPath")

	C.QCoreApplication_QCoreApplication_AddLibraryPath(C.CString(path))
}

func QCoreApplication_ApplicationDirPath() string {
	defer qt.Recovering("QCoreApplication::applicationDirPath")

	return C.GoString(C.QCoreApplication_QCoreApplication_ApplicationDirPath())
}

func QCoreApplication_ApplicationFilePath() string {
	defer qt.Recovering("QCoreApplication::applicationFilePath")

	return C.GoString(C.QCoreApplication_QCoreApplication_ApplicationFilePath())
}

func QCoreApplication_ApplicationPid() int64 {
	defer qt.Recovering("QCoreApplication::applicationPid")

	return int64(C.QCoreApplication_QCoreApplication_ApplicationPid())
}

func QCoreApplication_Arguments() []string {
	defer qt.Recovering("QCoreApplication::arguments")

	return strings.Split(C.GoString(C.QCoreApplication_QCoreApplication_Arguments()), ",,,")
}

func QCoreApplication_ClosingDown() bool {
	defer qt.Recovering("QCoreApplication::closingDown")

	return C.QCoreApplication_QCoreApplication_ClosingDown() != 0
}

func QCoreApplication_EventDispatcher() *QAbstractEventDispatcher {
	defer qt.Recovering("QCoreApplication::eventDispatcher")

	return NewQAbstractEventDispatcherFromPointer(C.QCoreApplication_QCoreApplication_EventDispatcher())
}

func QCoreApplication_Exec() int {
	defer qt.Recovering("QCoreApplication::exec")

	return int(C.QCoreApplication_QCoreApplication_Exec())
}

func QCoreApplication_Exit(returnCode int) {
	defer qt.Recovering("QCoreApplication::exit")

	C.QCoreApplication_QCoreApplication_Exit(C.int(returnCode))
}

func QCoreApplication_Flush() {
	defer qt.Recovering("QCoreApplication::flush")

	C.QCoreApplication_QCoreApplication_Flush()
}

func (ptr *QCoreApplication) InstallNativeEventFilter(filterObj QAbstractNativeEventFilter_ITF) {
	defer qt.Recovering("QCoreApplication::installNativeEventFilter")

	if ptr.Pointer() != nil {
		C.QCoreApplication_InstallNativeEventFilter(ptr.Pointer(), PointerFromQAbstractNativeEventFilter(filterObj))
	}
}

func QCoreApplication_InstallTranslator(translationFile QTranslator_ITF) bool {
	defer qt.Recovering("QCoreApplication::installTranslator")

	return C.QCoreApplication_QCoreApplication_InstallTranslator(PointerFromQTranslator(translationFile)) != 0
}

func QCoreApplication_Instance() *QCoreApplication {
	defer qt.Recovering("QCoreApplication::instance")

	return NewQCoreApplicationFromPointer(C.QCoreApplication_QCoreApplication_Instance())
}

func QCoreApplication_IsQuitLockEnabled() bool {
	defer qt.Recovering("QCoreApplication::isQuitLockEnabled")

	return C.QCoreApplication_QCoreApplication_IsQuitLockEnabled() != 0
}

func QCoreApplication_IsSetuidAllowed() bool {
	defer qt.Recovering("QCoreApplication::isSetuidAllowed")

	return C.QCoreApplication_QCoreApplication_IsSetuidAllowed() != 0
}

func QCoreApplication_LibraryPaths() []string {
	defer qt.Recovering("QCoreApplication::libraryPaths")

	return strings.Split(C.GoString(C.QCoreApplication_QCoreApplication_LibraryPaths()), ",,,")
}

func (ptr *QCoreApplication) Notify(receiver QObject_ITF, event QEvent_ITF) bool {
	defer qt.Recovering("QCoreApplication::notify")

	if ptr.Pointer() != nil {
		return C.QCoreApplication_Notify(ptr.Pointer(), PointerFromQObject(receiver), PointerFromQEvent(event)) != 0
	}
	return false
}

func QCoreApplication_PostEvent(receiver QObject_ITF, event QEvent_ITF, priority int) {
	defer qt.Recovering("QCoreApplication::postEvent")

	C.QCoreApplication_QCoreApplication_PostEvent(PointerFromQObject(receiver), PointerFromQEvent(event), C.int(priority))
}

func QCoreApplication_ProcessEvents(flags QEventLoop__ProcessEventsFlag) {
	defer qt.Recovering("QCoreApplication::processEvents")

	C.QCoreApplication_QCoreApplication_ProcessEvents(C.int(flags))
}

func QCoreApplication_ProcessEvents2(flags QEventLoop__ProcessEventsFlag, maxtime int) {
	defer qt.Recovering("QCoreApplication::processEvents")

	C.QCoreApplication_QCoreApplication_ProcessEvents2(C.int(flags), C.int(maxtime))
}

func QCoreApplication_Quit() {
	defer qt.Recovering("QCoreApplication::quit")

	C.QCoreApplication_QCoreApplication_Quit()
}

func QCoreApplication_RemoveLibraryPath(path string) {
	defer qt.Recovering("QCoreApplication::removeLibraryPath")

	C.QCoreApplication_QCoreApplication_RemoveLibraryPath(C.CString(path))
}

func (ptr *QCoreApplication) RemoveNativeEventFilter(filterObject QAbstractNativeEventFilter_ITF) {
	defer qt.Recovering("QCoreApplication::removeNativeEventFilter")

	if ptr.Pointer() != nil {
		C.QCoreApplication_RemoveNativeEventFilter(ptr.Pointer(), PointerFromQAbstractNativeEventFilter(filterObject))
	}
}

func QCoreApplication_RemovePostedEvents(receiver QObject_ITF, eventType int) {
	defer qt.Recovering("QCoreApplication::removePostedEvents")

	C.QCoreApplication_QCoreApplication_RemovePostedEvents(PointerFromQObject(receiver), C.int(eventType))
}

func QCoreApplication_RemoveTranslator(translationFile QTranslator_ITF) bool {
	defer qt.Recovering("QCoreApplication::removeTranslator")

	return C.QCoreApplication_QCoreApplication_RemoveTranslator(PointerFromQTranslator(translationFile)) != 0
}

func QCoreApplication_SendEvent(receiver QObject_ITF, event QEvent_ITF) bool {
	defer qt.Recovering("QCoreApplication::sendEvent")

	return C.QCoreApplication_QCoreApplication_SendEvent(PointerFromQObject(receiver), PointerFromQEvent(event)) != 0
}

func QCoreApplication_SendPostedEvents(receiver QObject_ITF, event_type int) {
	defer qt.Recovering("QCoreApplication::sendPostedEvents")

	C.QCoreApplication_QCoreApplication_SendPostedEvents(PointerFromQObject(receiver), C.int(event_type))
}

func QCoreApplication_SetAttribute(attribute Qt__ApplicationAttribute, on bool) {
	defer qt.Recovering("QCoreApplication::setAttribute")

	C.QCoreApplication_QCoreApplication_SetAttribute(C.int(attribute), C.int(qt.GoBoolToInt(on)))
}

func QCoreApplication_SetEventDispatcher(eventDispatcher QAbstractEventDispatcher_ITF) {
	defer qt.Recovering("QCoreApplication::setEventDispatcher")

	C.QCoreApplication_QCoreApplication_SetEventDispatcher(PointerFromQAbstractEventDispatcher(eventDispatcher))
}

func QCoreApplication_SetLibraryPaths(paths []string) {
	defer qt.Recovering("QCoreApplication::setLibraryPaths")

	C.QCoreApplication_QCoreApplication_SetLibraryPaths(C.CString(strings.Join(paths, ",,,")))
}

func QCoreApplication_SetQuitLockEnabled(enabled bool) {
	defer qt.Recovering("QCoreApplication::setQuitLockEnabled")

	C.QCoreApplication_QCoreApplication_SetQuitLockEnabled(C.int(qt.GoBoolToInt(enabled)))
}

func QCoreApplication_SetSetuidAllowed(allow bool) {
	defer qt.Recovering("QCoreApplication::setSetuidAllowed")

	C.QCoreApplication_QCoreApplication_SetSetuidAllowed(C.int(qt.GoBoolToInt(allow)))
}

func QCoreApplication_StartingUp() bool {
	defer qt.Recovering("QCoreApplication::startingUp")

	return C.QCoreApplication_QCoreApplication_StartingUp() != 0
}

func QCoreApplication_TestAttribute(attribute Qt__ApplicationAttribute) bool {
	defer qt.Recovering("QCoreApplication::testAttribute")

	return C.QCoreApplication_QCoreApplication_TestAttribute(C.int(attribute)) != 0
}

func QCoreApplication_Translate(context string, sourceText string, disambiguation string, n int) string {
	defer qt.Recovering("QCoreApplication::translate")

	return C.GoString(C.QCoreApplication_QCoreApplication_Translate(C.CString(context), C.CString(sourceText), C.CString(disambiguation), C.int(n)))
}

func (ptr *QCoreApplication) DestroyQCoreApplication() {
	defer qt.Recovering("QCoreApplication::~QCoreApplication")

	if ptr.Pointer() != nil {
		C.QCoreApplication_DestroyQCoreApplication(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QCoreApplication) ConnectTimerEvent(f func(event *QTimerEvent)) {
	defer qt.Recovering("connect QCoreApplication::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QCoreApplication) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QCoreApplication::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQCoreApplicationTimerEvent
func callbackQCoreApplicationTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QCoreApplication::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*QTimerEvent))(NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QCoreApplication) ConnectChildEvent(f func(event *QChildEvent)) {
	defer qt.Recovering("connect QCoreApplication::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QCoreApplication) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QCoreApplication::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQCoreApplicationChildEvent
func callbackQCoreApplicationChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QCoreApplication::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*QChildEvent))(NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QCoreApplication) ConnectCustomEvent(f func(event *QEvent)) {
	defer qt.Recovering("connect QCoreApplication::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QCoreApplication) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QCoreApplication::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQCoreApplicationCustomEvent
func callbackQCoreApplicationCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QCoreApplication::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*QEvent))(NewQEventFromPointer(event))
		return true
	}
	return false

}
