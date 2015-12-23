package help

//#include "help.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

type QHelpEngineCore struct {
	core.QObject
}

type QHelpEngineCore_ITF interface {
	core.QObject_ITF
	QHelpEngineCore_PTR() *QHelpEngineCore
}

func PointerFromQHelpEngineCore(ptr QHelpEngineCore_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpEngineCore_PTR().Pointer()
	}
	return nil
}

func NewQHelpEngineCoreFromPointer(ptr unsafe.Pointer) *QHelpEngineCore {
	var n = new(QHelpEngineCore)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QHelpEngineCore_") {
		n.SetObjectName("QHelpEngineCore_" + qt.Identifier())
	}
	return n
}

func (ptr *QHelpEngineCore) QHelpEngineCore_PTR() *QHelpEngineCore {
	return ptr
}

func (ptr *QHelpEngineCore) AutoSaveFilter() bool {
	defer qt.Recovering("QHelpEngineCore::autoSaveFilter")

	if ptr.Pointer() != nil {
		return C.QHelpEngineCore_AutoSaveFilter(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QHelpEngineCore) CollectionFile() string {
	defer qt.Recovering("QHelpEngineCore::collectionFile")

	if ptr.Pointer() != nil {
		return C.GoString(C.QHelpEngineCore_CollectionFile(ptr.Pointer()))
	}
	return ""
}

func (ptr *QHelpEngineCore) CurrentFilter() string {
	defer qt.Recovering("QHelpEngineCore::currentFilter")

	if ptr.Pointer() != nil {
		return C.GoString(C.QHelpEngineCore_CurrentFilter(ptr.Pointer()))
	}
	return ""
}

func (ptr *QHelpEngineCore) SetAutoSaveFilter(save bool) {
	defer qt.Recovering("QHelpEngineCore::setAutoSaveFilter")

	if ptr.Pointer() != nil {
		C.QHelpEngineCore_SetAutoSaveFilter(ptr.Pointer(), C.int(qt.GoBoolToInt(save)))
	}
}

func (ptr *QHelpEngineCore) SetCollectionFile(fileName string) {
	defer qt.Recovering("QHelpEngineCore::setCollectionFile")

	if ptr.Pointer() != nil {
		C.QHelpEngineCore_SetCollectionFile(ptr.Pointer(), C.CString(fileName))
	}
}

func (ptr *QHelpEngineCore) SetCurrentFilter(filterName string) {
	defer qt.Recovering("QHelpEngineCore::setCurrentFilter")

	if ptr.Pointer() != nil {
		C.QHelpEngineCore_SetCurrentFilter(ptr.Pointer(), C.CString(filterName))
	}
}

func NewQHelpEngineCore(collectionFile string, parent core.QObject_ITF) *QHelpEngineCore {
	defer qt.Recovering("QHelpEngineCore::QHelpEngineCore")

	return NewQHelpEngineCoreFromPointer(C.QHelpEngineCore_NewQHelpEngineCore(C.CString(collectionFile), core.PointerFromQObject(parent)))
}

func (ptr *QHelpEngineCore) AddCustomFilter(filterName string, attributes []string) bool {
	defer qt.Recovering("QHelpEngineCore::addCustomFilter")

	if ptr.Pointer() != nil {
		return C.QHelpEngineCore_AddCustomFilter(ptr.Pointer(), C.CString(filterName), C.CString(strings.Join(attributes, ",,,"))) != 0
	}
	return false
}

func (ptr *QHelpEngineCore) CopyCollectionFile(fileName string) bool {
	defer qt.Recovering("QHelpEngineCore::copyCollectionFile")

	if ptr.Pointer() != nil {
		return C.QHelpEngineCore_CopyCollectionFile(ptr.Pointer(), C.CString(fileName)) != 0
	}
	return false
}

func (ptr *QHelpEngineCore) ConnectCurrentFilterChanged(f func(newFilter string)) {
	defer qt.Recovering("connect QHelpEngineCore::currentFilterChanged")

	if ptr.Pointer() != nil {
		C.QHelpEngineCore_ConnectCurrentFilterChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "currentFilterChanged", f)
	}
}

func (ptr *QHelpEngineCore) DisconnectCurrentFilterChanged() {
	defer qt.Recovering("disconnect QHelpEngineCore::currentFilterChanged")

	if ptr.Pointer() != nil {
		C.QHelpEngineCore_DisconnectCurrentFilterChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "currentFilterChanged")
	}
}

//export callbackQHelpEngineCoreCurrentFilterChanged
func callbackQHelpEngineCoreCurrentFilterChanged(ptrName *C.char, newFilter *C.char) {
	defer qt.Recovering("callback QHelpEngineCore::currentFilterChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "currentFilterChanged"); signal != nil {
		signal.(func(string))(C.GoString(newFilter))
	}

}

func (ptr *QHelpEngineCore) CustomFilters() []string {
	defer qt.Recovering("QHelpEngineCore::customFilters")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QHelpEngineCore_CustomFilters(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QHelpEngineCore) CustomValue(key string, defaultValue core.QVariant_ITF) *core.QVariant {
	defer qt.Recovering("QHelpEngineCore::customValue")

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QHelpEngineCore_CustomValue(ptr.Pointer(), C.CString(key), core.PointerFromQVariant(defaultValue)))
	}
	return nil
}

func (ptr *QHelpEngineCore) DocumentationFileName(namespaceName string) string {
	defer qt.Recovering("QHelpEngineCore::documentationFileName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QHelpEngineCore_DocumentationFileName(ptr.Pointer(), C.CString(namespaceName)))
	}
	return ""
}

func (ptr *QHelpEngineCore) Error() string {
	defer qt.Recovering("QHelpEngineCore::error")

	if ptr.Pointer() != nil {
		return C.GoString(C.QHelpEngineCore_Error(ptr.Pointer()))
	}
	return ""
}

func (ptr *QHelpEngineCore) FileData(url core.QUrl_ITF) *core.QByteArray {
	defer qt.Recovering("QHelpEngineCore::fileData")

	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QHelpEngineCore_FileData(ptr.Pointer(), core.PointerFromQUrl(url)))
	}
	return nil
}

func (ptr *QHelpEngineCore) FilterAttributes() []string {
	defer qt.Recovering("QHelpEngineCore::filterAttributes")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QHelpEngineCore_FilterAttributes(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QHelpEngineCore) FilterAttributes2(filterName string) []string {
	defer qt.Recovering("QHelpEngineCore::filterAttributes")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QHelpEngineCore_FilterAttributes2(ptr.Pointer(), C.CString(filterName))), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QHelpEngineCore) FindFile(url core.QUrl_ITF) *core.QUrl {
	defer qt.Recovering("QHelpEngineCore::findFile")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QHelpEngineCore_FindFile(ptr.Pointer(), core.PointerFromQUrl(url)))
	}
	return nil
}

func QHelpEngineCore_MetaData(documentationFileName string, name string) *core.QVariant {
	defer qt.Recovering("QHelpEngineCore::metaData")

	return core.NewQVariantFromPointer(C.QHelpEngineCore_QHelpEngineCore_MetaData(C.CString(documentationFileName), C.CString(name)))
}

func QHelpEngineCore_NamespaceName(documentationFileName string) string {
	defer qt.Recovering("QHelpEngineCore::namespaceName")

	return C.GoString(C.QHelpEngineCore_QHelpEngineCore_NamespaceName(C.CString(documentationFileName)))
}

func (ptr *QHelpEngineCore) ConnectReadersAboutToBeInvalidated(f func()) {
	defer qt.Recovering("connect QHelpEngineCore::readersAboutToBeInvalidated")

	if ptr.Pointer() != nil {
		C.QHelpEngineCore_ConnectReadersAboutToBeInvalidated(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "readersAboutToBeInvalidated", f)
	}
}

func (ptr *QHelpEngineCore) DisconnectReadersAboutToBeInvalidated() {
	defer qt.Recovering("disconnect QHelpEngineCore::readersAboutToBeInvalidated")

	if ptr.Pointer() != nil {
		C.QHelpEngineCore_DisconnectReadersAboutToBeInvalidated(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "readersAboutToBeInvalidated")
	}
}

//export callbackQHelpEngineCoreReadersAboutToBeInvalidated
func callbackQHelpEngineCoreReadersAboutToBeInvalidated(ptrName *C.char) {
	defer qt.Recovering("callback QHelpEngineCore::readersAboutToBeInvalidated")

	if signal := qt.GetSignal(C.GoString(ptrName), "readersAboutToBeInvalidated"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHelpEngineCore) RegisterDocumentation(documentationFileName string) bool {
	defer qt.Recovering("QHelpEngineCore::registerDocumentation")

	if ptr.Pointer() != nil {
		return C.QHelpEngineCore_RegisterDocumentation(ptr.Pointer(), C.CString(documentationFileName)) != 0
	}
	return false
}

func (ptr *QHelpEngineCore) RegisteredDocumentations() []string {
	defer qt.Recovering("QHelpEngineCore::registeredDocumentations")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QHelpEngineCore_RegisteredDocumentations(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QHelpEngineCore) RemoveCustomFilter(filterName string) bool {
	defer qt.Recovering("QHelpEngineCore::removeCustomFilter")

	if ptr.Pointer() != nil {
		return C.QHelpEngineCore_RemoveCustomFilter(ptr.Pointer(), C.CString(filterName)) != 0
	}
	return false
}

func (ptr *QHelpEngineCore) RemoveCustomValue(key string) bool {
	defer qt.Recovering("QHelpEngineCore::removeCustomValue")

	if ptr.Pointer() != nil {
		return C.QHelpEngineCore_RemoveCustomValue(ptr.Pointer(), C.CString(key)) != 0
	}
	return false
}

func (ptr *QHelpEngineCore) SetCustomValue(key string, value core.QVariant_ITF) bool {
	defer qt.Recovering("QHelpEngineCore::setCustomValue")

	if ptr.Pointer() != nil {
		return C.QHelpEngineCore_SetCustomValue(ptr.Pointer(), C.CString(key), core.PointerFromQVariant(value)) != 0
	}
	return false
}

func (ptr *QHelpEngineCore) SetupData() bool {
	defer qt.Recovering("QHelpEngineCore::setupData")

	if ptr.Pointer() != nil {
		return C.QHelpEngineCore_SetupData(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QHelpEngineCore) ConnectSetupFinished(f func()) {
	defer qt.Recovering("connect QHelpEngineCore::setupFinished")

	if ptr.Pointer() != nil {
		C.QHelpEngineCore_ConnectSetupFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "setupFinished", f)
	}
}

func (ptr *QHelpEngineCore) DisconnectSetupFinished() {
	defer qt.Recovering("disconnect QHelpEngineCore::setupFinished")

	if ptr.Pointer() != nil {
		C.QHelpEngineCore_DisconnectSetupFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "setupFinished")
	}
}

//export callbackQHelpEngineCoreSetupFinished
func callbackQHelpEngineCoreSetupFinished(ptrName *C.char) {
	defer qt.Recovering("callback QHelpEngineCore::setupFinished")

	if signal := qt.GetSignal(C.GoString(ptrName), "setupFinished"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHelpEngineCore) ConnectSetupStarted(f func()) {
	defer qt.Recovering("connect QHelpEngineCore::setupStarted")

	if ptr.Pointer() != nil {
		C.QHelpEngineCore_ConnectSetupStarted(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "setupStarted", f)
	}
}

func (ptr *QHelpEngineCore) DisconnectSetupStarted() {
	defer qt.Recovering("disconnect QHelpEngineCore::setupStarted")

	if ptr.Pointer() != nil {
		C.QHelpEngineCore_DisconnectSetupStarted(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "setupStarted")
	}
}

//export callbackQHelpEngineCoreSetupStarted
func callbackQHelpEngineCoreSetupStarted(ptrName *C.char) {
	defer qt.Recovering("callback QHelpEngineCore::setupStarted")

	if signal := qt.GetSignal(C.GoString(ptrName), "setupStarted"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QHelpEngineCore) UnregisterDocumentation(namespaceName string) bool {
	defer qt.Recovering("QHelpEngineCore::unregisterDocumentation")

	if ptr.Pointer() != nil {
		return C.QHelpEngineCore_UnregisterDocumentation(ptr.Pointer(), C.CString(namespaceName)) != 0
	}
	return false
}

func (ptr *QHelpEngineCore) ConnectWarning(f func(msg string)) {
	defer qt.Recovering("connect QHelpEngineCore::warning")

	if ptr.Pointer() != nil {
		C.QHelpEngineCore_ConnectWarning(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "warning", f)
	}
}

func (ptr *QHelpEngineCore) DisconnectWarning() {
	defer qt.Recovering("disconnect QHelpEngineCore::warning")

	if ptr.Pointer() != nil {
		C.QHelpEngineCore_DisconnectWarning(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "warning")
	}
}

//export callbackQHelpEngineCoreWarning
func callbackQHelpEngineCoreWarning(ptrName *C.char, msg *C.char) {
	defer qt.Recovering("callback QHelpEngineCore::warning")

	if signal := qt.GetSignal(C.GoString(ptrName), "warning"); signal != nil {
		signal.(func(string))(C.GoString(msg))
	}

}

func (ptr *QHelpEngineCore) DestroyQHelpEngineCore() {
	defer qt.Recovering("QHelpEngineCore::~QHelpEngineCore")

	if ptr.Pointer() != nil {
		C.QHelpEngineCore_DestroyQHelpEngineCore(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QHelpEngineCore) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QHelpEngineCore::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QHelpEngineCore) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QHelpEngineCore::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQHelpEngineCoreTimerEvent
func callbackQHelpEngineCoreTimerEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QHelpEngineCore::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QHelpEngineCore) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QHelpEngineCore::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QHelpEngineCore) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QHelpEngineCore::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQHelpEngineCoreChildEvent
func callbackQHelpEngineCoreChildEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QHelpEngineCore::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
		return true
	}
	return false

}

func (ptr *QHelpEngineCore) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QHelpEngineCore::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QHelpEngineCore) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QHelpEngineCore::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQHelpEngineCoreCustomEvent
func callbackQHelpEngineCoreCustomEvent(ptrName *C.char, event unsafe.Pointer) bool {
	defer qt.Recovering("callback QHelpEngineCore::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
		return true
	}
	return false

}
