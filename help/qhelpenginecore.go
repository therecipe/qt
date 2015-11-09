package help

//#include "qhelpenginecore.h"
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
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QHelpEngineCore_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QHelpEngineCore) QHelpEngineCore_PTR() *QHelpEngineCore {
	return ptr
}

func (ptr *QHelpEngineCore) AutoSaveFilter() bool {
	if ptr.Pointer() != nil {
		return C.QHelpEngineCore_AutoSaveFilter(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QHelpEngineCore) CollectionFile() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QHelpEngineCore_CollectionFile(ptr.Pointer()))
	}
	return ""
}

func (ptr *QHelpEngineCore) CurrentFilter() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QHelpEngineCore_CurrentFilter(ptr.Pointer()))
	}
	return ""
}

func (ptr *QHelpEngineCore) SetAutoSaveFilter(save bool) {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore_SetAutoSaveFilter(ptr.Pointer(), C.int(qt.GoBoolToInt(save)))
	}
}

func (ptr *QHelpEngineCore) SetCollectionFile(fileName string) {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore_SetCollectionFile(ptr.Pointer(), C.CString(fileName))
	}
}

func (ptr *QHelpEngineCore) SetCurrentFilter(filterName string) {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore_SetCurrentFilter(ptr.Pointer(), C.CString(filterName))
	}
}

func NewQHelpEngineCore(collectionFile string, parent core.QObject_ITF) *QHelpEngineCore {
	return NewQHelpEngineCoreFromPointer(C.QHelpEngineCore_NewQHelpEngineCore(C.CString(collectionFile), core.PointerFromQObject(parent)))
}

func (ptr *QHelpEngineCore) AddCustomFilter(filterName string, attributes []string) bool {
	if ptr.Pointer() != nil {
		return C.QHelpEngineCore_AddCustomFilter(ptr.Pointer(), C.CString(filterName), C.CString(strings.Join(attributes, "|"))) != 0
	}
	return false
}

func (ptr *QHelpEngineCore) CopyCollectionFile(fileName string) bool {
	if ptr.Pointer() != nil {
		return C.QHelpEngineCore_CopyCollectionFile(ptr.Pointer(), C.CString(fileName)) != 0
	}
	return false
}

func (ptr *QHelpEngineCore) ConnectCurrentFilterChanged(f func(newFilter string)) {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore_ConnectCurrentFilterChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "currentFilterChanged", f)
	}
}

func (ptr *QHelpEngineCore) DisconnectCurrentFilterChanged() {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore_DisconnectCurrentFilterChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "currentFilterChanged")
	}
}

//export callbackQHelpEngineCoreCurrentFilterChanged
func callbackQHelpEngineCoreCurrentFilterChanged(ptrName *C.char, newFilter *C.char) {
	qt.GetSignal(C.GoString(ptrName), "currentFilterChanged").(func(string))(C.GoString(newFilter))
}

func (ptr *QHelpEngineCore) CustomFilters() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QHelpEngineCore_CustomFilters(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QHelpEngineCore) CustomValue(key string, defaultValue core.QVariant_ITF) *core.QVariant {
	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QHelpEngineCore_CustomValue(ptr.Pointer(), C.CString(key), core.PointerFromQVariant(defaultValue)))
	}
	return nil
}

func (ptr *QHelpEngineCore) DocumentationFileName(namespaceName string) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QHelpEngineCore_DocumentationFileName(ptr.Pointer(), C.CString(namespaceName)))
	}
	return ""
}

func (ptr *QHelpEngineCore) Error() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QHelpEngineCore_Error(ptr.Pointer()))
	}
	return ""
}

func (ptr *QHelpEngineCore) FileData(url core.QUrl_ITF) *core.QByteArray {
	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QHelpEngineCore_FileData(ptr.Pointer(), core.PointerFromQUrl(url)))
	}
	return nil
}

func (ptr *QHelpEngineCore) FilterAttributes() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QHelpEngineCore_FilterAttributes(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QHelpEngineCore) FilterAttributes2(filterName string) []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QHelpEngineCore_FilterAttributes2(ptr.Pointer(), C.CString(filterName))), "|")
	}
	return make([]string, 0)
}

func QHelpEngineCore_MetaData(documentationFileName string, name string) *core.QVariant {
	return core.NewQVariantFromPointer(C.QHelpEngineCore_QHelpEngineCore_MetaData(C.CString(documentationFileName), C.CString(name)))
}

func QHelpEngineCore_NamespaceName(documentationFileName string) string {
	return C.GoString(C.QHelpEngineCore_QHelpEngineCore_NamespaceName(C.CString(documentationFileName)))
}

func (ptr *QHelpEngineCore) ConnectReadersAboutToBeInvalidated(f func()) {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore_ConnectReadersAboutToBeInvalidated(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "readersAboutToBeInvalidated", f)
	}
}

func (ptr *QHelpEngineCore) DisconnectReadersAboutToBeInvalidated() {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore_DisconnectReadersAboutToBeInvalidated(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "readersAboutToBeInvalidated")
	}
}

//export callbackQHelpEngineCoreReadersAboutToBeInvalidated
func callbackQHelpEngineCoreReadersAboutToBeInvalidated(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "readersAboutToBeInvalidated").(func())()
}

func (ptr *QHelpEngineCore) RegisterDocumentation(documentationFileName string) bool {
	if ptr.Pointer() != nil {
		return C.QHelpEngineCore_RegisterDocumentation(ptr.Pointer(), C.CString(documentationFileName)) != 0
	}
	return false
}

func (ptr *QHelpEngineCore) RegisteredDocumentations() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QHelpEngineCore_RegisteredDocumentations(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QHelpEngineCore) RemoveCustomFilter(filterName string) bool {
	if ptr.Pointer() != nil {
		return C.QHelpEngineCore_RemoveCustomFilter(ptr.Pointer(), C.CString(filterName)) != 0
	}
	return false
}

func (ptr *QHelpEngineCore) RemoveCustomValue(key string) bool {
	if ptr.Pointer() != nil {
		return C.QHelpEngineCore_RemoveCustomValue(ptr.Pointer(), C.CString(key)) != 0
	}
	return false
}

func (ptr *QHelpEngineCore) SetCustomValue(key string, value core.QVariant_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QHelpEngineCore_SetCustomValue(ptr.Pointer(), C.CString(key), core.PointerFromQVariant(value)) != 0
	}
	return false
}

func (ptr *QHelpEngineCore) SetupData() bool {
	if ptr.Pointer() != nil {
		return C.QHelpEngineCore_SetupData(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QHelpEngineCore) ConnectSetupFinished(f func()) {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore_ConnectSetupFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "setupFinished", f)
	}
}

func (ptr *QHelpEngineCore) DisconnectSetupFinished() {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore_DisconnectSetupFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "setupFinished")
	}
}

//export callbackQHelpEngineCoreSetupFinished
func callbackQHelpEngineCoreSetupFinished(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "setupFinished").(func())()
}

func (ptr *QHelpEngineCore) ConnectSetupStarted(f func()) {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore_ConnectSetupStarted(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "setupStarted", f)
	}
}

func (ptr *QHelpEngineCore) DisconnectSetupStarted() {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore_DisconnectSetupStarted(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "setupStarted")
	}
}

//export callbackQHelpEngineCoreSetupStarted
func callbackQHelpEngineCoreSetupStarted(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "setupStarted").(func())()
}

func (ptr *QHelpEngineCore) UnregisterDocumentation(namespaceName string) bool {
	if ptr.Pointer() != nil {
		return C.QHelpEngineCore_UnregisterDocumentation(ptr.Pointer(), C.CString(namespaceName)) != 0
	}
	return false
}

func (ptr *QHelpEngineCore) ConnectWarning(f func(msg string)) {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore_ConnectWarning(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "warning", f)
	}
}

func (ptr *QHelpEngineCore) DisconnectWarning() {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore_DisconnectWarning(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "warning")
	}
}

//export callbackQHelpEngineCoreWarning
func callbackQHelpEngineCoreWarning(ptrName *C.char, msg *C.char) {
	qt.GetSignal(C.GoString(ptrName), "warning").(func(string))(C.GoString(msg))
}

func (ptr *QHelpEngineCore) DestroyQHelpEngineCore() {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore_DestroyQHelpEngineCore(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
