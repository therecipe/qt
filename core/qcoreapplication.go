package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"log"
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
		n.SetObjectName("QCoreApplication_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QCoreApplication) QCoreApplication_PTR() *QCoreApplication {
	return ptr
}

func QCoreApplication_ApplicationName() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCoreApplication::applicationName")
		}
	}()

	return C.GoString(C.QCoreApplication_QCoreApplication_ApplicationName())
}

func QCoreApplication_ApplicationVersion() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCoreApplication::applicationVersion")
		}
	}()

	return C.GoString(C.QCoreApplication_QCoreApplication_ApplicationVersion())
}

func QCoreApplication_OrganizationDomain() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCoreApplication::organizationDomain")
		}
	}()

	return C.GoString(C.QCoreApplication_QCoreApplication_OrganizationDomain())
}

func QCoreApplication_OrganizationName() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCoreApplication::organizationName")
		}
	}()

	return C.GoString(C.QCoreApplication_QCoreApplication_OrganizationName())
}

func QCoreApplication_SetApplicationName(application string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCoreApplication::setApplicationName")
		}
	}()

	C.QCoreApplication_QCoreApplication_SetApplicationName(C.CString(application))
}

func QCoreApplication_SetApplicationVersion(version string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCoreApplication::setApplicationVersion")
		}
	}()

	C.QCoreApplication_QCoreApplication_SetApplicationVersion(C.CString(version))
}

func QCoreApplication_SetOrganizationDomain(orgDomain string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCoreApplication::setOrganizationDomain")
		}
	}()

	C.QCoreApplication_QCoreApplication_SetOrganizationDomain(C.CString(orgDomain))
}

func QCoreApplication_SetOrganizationName(orgName string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCoreApplication::setOrganizationName")
		}
	}()

	C.QCoreApplication_QCoreApplication_SetOrganizationName(C.CString(orgName))
}

func NewQCoreApplication(argc int, argv []string) *QCoreApplication {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCoreApplication::QCoreApplication")
		}
	}()

	return NewQCoreApplicationFromPointer(C.QCoreApplication_NewQCoreApplication(C.int(argc), C.CString(strings.Join(argv, ",,,"))))
}

func (ptr *QCoreApplication) ConnectAboutToQuit(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCoreApplication::aboutToQuit")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCoreApplication_ConnectAboutToQuit(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "aboutToQuit", f)
	}
}

func (ptr *QCoreApplication) DisconnectAboutToQuit() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCoreApplication::aboutToQuit")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCoreApplication_DisconnectAboutToQuit(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "aboutToQuit")
	}
}

//export callbackQCoreApplicationAboutToQuit
func callbackQCoreApplicationAboutToQuit(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCoreApplication::aboutToQuit")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "aboutToQuit").(func())()
}

func QCoreApplication_AddLibraryPath(path string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCoreApplication::addLibraryPath")
		}
	}()

	C.QCoreApplication_QCoreApplication_AddLibraryPath(C.CString(path))
}

func QCoreApplication_ApplicationDirPath() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCoreApplication::applicationDirPath")
		}
	}()

	return C.GoString(C.QCoreApplication_QCoreApplication_ApplicationDirPath())
}

func QCoreApplication_ApplicationFilePath() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCoreApplication::applicationFilePath")
		}
	}()

	return C.GoString(C.QCoreApplication_QCoreApplication_ApplicationFilePath())
}

func QCoreApplication_Arguments() []string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCoreApplication::arguments")
		}
	}()

	return strings.Split(C.GoString(C.QCoreApplication_QCoreApplication_Arguments()), ",,,")
}

func QCoreApplication_ClosingDown() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCoreApplication::closingDown")
		}
	}()

	return C.QCoreApplication_QCoreApplication_ClosingDown() != 0
}

func QCoreApplication_EventDispatcher() *QAbstractEventDispatcher {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCoreApplication::eventDispatcher")
		}
	}()

	return NewQAbstractEventDispatcherFromPointer(C.QCoreApplication_QCoreApplication_EventDispatcher())
}

func QCoreApplication_Exec() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCoreApplication::exec")
		}
	}()

	return int(C.QCoreApplication_QCoreApplication_Exec())
}

func QCoreApplication_Exit(returnCode int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCoreApplication::exit")
		}
	}()

	C.QCoreApplication_QCoreApplication_Exit(C.int(returnCode))
}

func QCoreApplication_Flush() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCoreApplication::flush")
		}
	}()

	C.QCoreApplication_QCoreApplication_Flush()
}

func (ptr *QCoreApplication) InstallNativeEventFilter(filterObj QAbstractNativeEventFilter_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCoreApplication::installNativeEventFilter")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCoreApplication_InstallNativeEventFilter(ptr.Pointer(), PointerFromQAbstractNativeEventFilter(filterObj))
	}
}

func QCoreApplication_InstallTranslator(translationFile QTranslator_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCoreApplication::installTranslator")
		}
	}()

	return C.QCoreApplication_QCoreApplication_InstallTranslator(PointerFromQTranslator(translationFile)) != 0
}

func QCoreApplication_Instance() *QCoreApplication {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCoreApplication::instance")
		}
	}()

	return NewQCoreApplicationFromPointer(C.QCoreApplication_QCoreApplication_Instance())
}

func QCoreApplication_IsQuitLockEnabled() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCoreApplication::isQuitLockEnabled")
		}
	}()

	return C.QCoreApplication_QCoreApplication_IsQuitLockEnabled() != 0
}

func QCoreApplication_IsSetuidAllowed() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCoreApplication::isSetuidAllowed")
		}
	}()

	return C.QCoreApplication_QCoreApplication_IsSetuidAllowed() != 0
}

func QCoreApplication_LibraryPaths() []string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCoreApplication::libraryPaths")
		}
	}()

	return strings.Split(C.GoString(C.QCoreApplication_QCoreApplication_LibraryPaths()), ",,,")
}

func (ptr *QCoreApplication) Notify(receiver QObject_ITF, event QEvent_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCoreApplication::notify")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QCoreApplication_Notify(ptr.Pointer(), PointerFromQObject(receiver), PointerFromQEvent(event)) != 0
	}
	return false
}

func QCoreApplication_PostEvent(receiver QObject_ITF, event QEvent_ITF, priority int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCoreApplication::postEvent")
		}
	}()

	C.QCoreApplication_QCoreApplication_PostEvent(PointerFromQObject(receiver), PointerFromQEvent(event), C.int(priority))
}

func QCoreApplication_ProcessEvents(flags QEventLoop__ProcessEventsFlag) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCoreApplication::processEvents")
		}
	}()

	C.QCoreApplication_QCoreApplication_ProcessEvents(C.int(flags))
}

func QCoreApplication_ProcessEvents2(flags QEventLoop__ProcessEventsFlag, maxtime int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCoreApplication::processEvents")
		}
	}()

	C.QCoreApplication_QCoreApplication_ProcessEvents2(C.int(flags), C.int(maxtime))
}

func QCoreApplication_Quit() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCoreApplication::quit")
		}
	}()

	C.QCoreApplication_QCoreApplication_Quit()
}

func QCoreApplication_RemoveLibraryPath(path string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCoreApplication::removeLibraryPath")
		}
	}()

	C.QCoreApplication_QCoreApplication_RemoveLibraryPath(C.CString(path))
}

func (ptr *QCoreApplication) RemoveNativeEventFilter(filterObject QAbstractNativeEventFilter_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCoreApplication::removeNativeEventFilter")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCoreApplication_RemoveNativeEventFilter(ptr.Pointer(), PointerFromQAbstractNativeEventFilter(filterObject))
	}
}

func QCoreApplication_RemovePostedEvents(receiver QObject_ITF, eventType int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCoreApplication::removePostedEvents")
		}
	}()

	C.QCoreApplication_QCoreApplication_RemovePostedEvents(PointerFromQObject(receiver), C.int(eventType))
}

func QCoreApplication_RemoveTranslator(translationFile QTranslator_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCoreApplication::removeTranslator")
		}
	}()

	return C.QCoreApplication_QCoreApplication_RemoveTranslator(PointerFromQTranslator(translationFile)) != 0
}

func QCoreApplication_SendEvent(receiver QObject_ITF, event QEvent_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCoreApplication::sendEvent")
		}
	}()

	return C.QCoreApplication_QCoreApplication_SendEvent(PointerFromQObject(receiver), PointerFromQEvent(event)) != 0
}

func QCoreApplication_SendPostedEvents(receiver QObject_ITF, event_type int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCoreApplication::sendPostedEvents")
		}
	}()

	C.QCoreApplication_QCoreApplication_SendPostedEvents(PointerFromQObject(receiver), C.int(event_type))
}

func QCoreApplication_SetAttribute(attribute Qt__ApplicationAttribute, on bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCoreApplication::setAttribute")
		}
	}()

	C.QCoreApplication_QCoreApplication_SetAttribute(C.int(attribute), C.int(qt.GoBoolToInt(on)))
}

func QCoreApplication_SetEventDispatcher(eventDispatcher QAbstractEventDispatcher_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCoreApplication::setEventDispatcher")
		}
	}()

	C.QCoreApplication_QCoreApplication_SetEventDispatcher(PointerFromQAbstractEventDispatcher(eventDispatcher))
}

func QCoreApplication_SetLibraryPaths(paths []string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCoreApplication::setLibraryPaths")
		}
	}()

	C.QCoreApplication_QCoreApplication_SetLibraryPaths(C.CString(strings.Join(paths, ",,,")))
}

func QCoreApplication_SetQuitLockEnabled(enabled bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCoreApplication::setQuitLockEnabled")
		}
	}()

	C.QCoreApplication_QCoreApplication_SetQuitLockEnabled(C.int(qt.GoBoolToInt(enabled)))
}

func QCoreApplication_SetSetuidAllowed(allow bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCoreApplication::setSetuidAllowed")
		}
	}()

	C.QCoreApplication_QCoreApplication_SetSetuidAllowed(C.int(qt.GoBoolToInt(allow)))
}

func QCoreApplication_StartingUp() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCoreApplication::startingUp")
		}
	}()

	return C.QCoreApplication_QCoreApplication_StartingUp() != 0
}

func QCoreApplication_TestAttribute(attribute Qt__ApplicationAttribute) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCoreApplication::testAttribute")
		}
	}()

	return C.QCoreApplication_QCoreApplication_TestAttribute(C.int(attribute)) != 0
}

func QCoreApplication_Translate(context string, sourceText string, disambiguation string, n int) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCoreApplication::translate")
		}
	}()

	return C.GoString(C.QCoreApplication_QCoreApplication_Translate(C.CString(context), C.CString(sourceText), C.CString(disambiguation), C.int(n)))
}

func (ptr *QCoreApplication) DestroyQCoreApplication() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QCoreApplication::~QCoreApplication")
		}
	}()

	if ptr.Pointer() != nil {
		C.QCoreApplication_DestroyQCoreApplication(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
