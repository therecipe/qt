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

type QHelpEngineCoreITF interface {
	core.QObjectITF
	QHelpEngineCorePTR() *QHelpEngineCore
}

func PointerFromQHelpEngineCore(ptr QHelpEngineCoreITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpEngineCorePTR().Pointer()
	}
	return nil
}

func QHelpEngineCoreFromPointer(ptr unsafe.Pointer) *QHelpEngineCore {
	var n = new(QHelpEngineCore)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QHelpEngineCore_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QHelpEngineCore) QHelpEngineCorePTR() *QHelpEngineCore {
	return ptr
}

func (ptr *QHelpEngineCore) AutoSaveFilter() bool {
	if ptr.Pointer() != nil {
		return C.QHelpEngineCore_AutoSaveFilter(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QHelpEngineCore) CollectionFile() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QHelpEngineCore_CollectionFile(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QHelpEngineCore) CurrentFilter() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QHelpEngineCore_CurrentFilter(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QHelpEngineCore) SetAutoSaveFilter(save bool) {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore_SetAutoSaveFilter(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(save)))
	}
}

func (ptr *QHelpEngineCore) SetCollectionFile(fileName string) {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore_SetCollectionFile(C.QtObjectPtr(ptr.Pointer()), C.CString(fileName))
	}
}

func (ptr *QHelpEngineCore) SetCurrentFilter(filterName string) {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore_SetCurrentFilter(C.QtObjectPtr(ptr.Pointer()), C.CString(filterName))
	}
}

func NewQHelpEngineCore(collectionFile string, parent core.QObjectITF) *QHelpEngineCore {
	return QHelpEngineCoreFromPointer(unsafe.Pointer(C.QHelpEngineCore_NewQHelpEngineCore(C.CString(collectionFile), C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QHelpEngineCore) AddCustomFilter(filterName string, attributes []string) bool {
	if ptr.Pointer() != nil {
		return C.QHelpEngineCore_AddCustomFilter(C.QtObjectPtr(ptr.Pointer()), C.CString(filterName), C.CString(strings.Join(attributes, "|"))) != 0
	}
	return false
}

func (ptr *QHelpEngineCore) CopyCollectionFile(fileName string) bool {
	if ptr.Pointer() != nil {
		return C.QHelpEngineCore_CopyCollectionFile(C.QtObjectPtr(ptr.Pointer()), C.CString(fileName)) != 0
	}
	return false
}

func (ptr *QHelpEngineCore) ConnectCurrentFilterChanged(f func(newFilter string)) {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore_ConnectCurrentFilterChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "currentFilterChanged", f)
	}
}

func (ptr *QHelpEngineCore) DisconnectCurrentFilterChanged() {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore_DisconnectCurrentFilterChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "currentFilterChanged")
	}
}

//export callbackQHelpEngineCoreCurrentFilterChanged
func callbackQHelpEngineCoreCurrentFilterChanged(ptrName *C.char, newFilter *C.char) {
	qt.GetSignal(C.GoString(ptrName), "currentFilterChanged").(func(string))(C.GoString(newFilter))
}

func (ptr *QHelpEngineCore) CustomFilters() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QHelpEngineCore_CustomFilters(C.QtObjectPtr(ptr.Pointer()))), "|")
	}
	return make([]string, 0)
}

func (ptr *QHelpEngineCore) CustomValue(key string, defaultValue string) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QHelpEngineCore_CustomValue(C.QtObjectPtr(ptr.Pointer()), C.CString(key), C.CString(defaultValue)))
	}
	return ""
}

func (ptr *QHelpEngineCore) DocumentationFileName(namespaceName string) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QHelpEngineCore_DocumentationFileName(C.QtObjectPtr(ptr.Pointer()), C.CString(namespaceName)))
	}
	return ""
}

func (ptr *QHelpEngineCore) Error() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QHelpEngineCore_Error(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QHelpEngineCore) FilterAttributes() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QHelpEngineCore_FilterAttributes(C.QtObjectPtr(ptr.Pointer()))), "|")
	}
	return make([]string, 0)
}

func (ptr *QHelpEngineCore) FilterAttributes2(filterName string) []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QHelpEngineCore_FilterAttributes2(C.QtObjectPtr(ptr.Pointer()), C.CString(filterName))), "|")
	}
	return make([]string, 0)
}

func (ptr *QHelpEngineCore) FindFile(url string) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QHelpEngineCore_FindFile(C.QtObjectPtr(ptr.Pointer()), C.CString(url)))
	}
	return ""
}

func QHelpEngineCore_MetaData(documentationFileName string, name string) string {
	return C.GoString(C.QHelpEngineCore_QHelpEngineCore_MetaData(C.CString(documentationFileName), C.CString(name)))
}

func QHelpEngineCore_NamespaceName(documentationFileName string) string {
	return C.GoString(C.QHelpEngineCore_QHelpEngineCore_NamespaceName(C.CString(documentationFileName)))
}

func (ptr *QHelpEngineCore) ConnectReadersAboutToBeInvalidated(f func()) {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore_ConnectReadersAboutToBeInvalidated(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "readersAboutToBeInvalidated", f)
	}
}

func (ptr *QHelpEngineCore) DisconnectReadersAboutToBeInvalidated() {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore_DisconnectReadersAboutToBeInvalidated(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "readersAboutToBeInvalidated")
	}
}

//export callbackQHelpEngineCoreReadersAboutToBeInvalidated
func callbackQHelpEngineCoreReadersAboutToBeInvalidated(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "readersAboutToBeInvalidated").(func())()
}

func (ptr *QHelpEngineCore) RegisterDocumentation(documentationFileName string) bool {
	if ptr.Pointer() != nil {
		return C.QHelpEngineCore_RegisterDocumentation(C.QtObjectPtr(ptr.Pointer()), C.CString(documentationFileName)) != 0
	}
	return false
}

func (ptr *QHelpEngineCore) RegisteredDocumentations() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QHelpEngineCore_RegisteredDocumentations(C.QtObjectPtr(ptr.Pointer()))), "|")
	}
	return make([]string, 0)
}

func (ptr *QHelpEngineCore) RemoveCustomFilter(filterName string) bool {
	if ptr.Pointer() != nil {
		return C.QHelpEngineCore_RemoveCustomFilter(C.QtObjectPtr(ptr.Pointer()), C.CString(filterName)) != 0
	}
	return false
}

func (ptr *QHelpEngineCore) RemoveCustomValue(key string) bool {
	if ptr.Pointer() != nil {
		return C.QHelpEngineCore_RemoveCustomValue(C.QtObjectPtr(ptr.Pointer()), C.CString(key)) != 0
	}
	return false
}

func (ptr *QHelpEngineCore) SetCustomValue(key string, value string) bool {
	if ptr.Pointer() != nil {
		return C.QHelpEngineCore_SetCustomValue(C.QtObjectPtr(ptr.Pointer()), C.CString(key), C.CString(value)) != 0
	}
	return false
}

func (ptr *QHelpEngineCore) SetupData() bool {
	if ptr.Pointer() != nil {
		return C.QHelpEngineCore_SetupData(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QHelpEngineCore) ConnectSetupFinished(f func()) {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore_ConnectSetupFinished(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "setupFinished", f)
	}
}

func (ptr *QHelpEngineCore) DisconnectSetupFinished() {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore_DisconnectSetupFinished(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "setupFinished")
	}
}

//export callbackQHelpEngineCoreSetupFinished
func callbackQHelpEngineCoreSetupFinished(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "setupFinished").(func())()
}

func (ptr *QHelpEngineCore) ConnectSetupStarted(f func()) {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore_ConnectSetupStarted(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "setupStarted", f)
	}
}

func (ptr *QHelpEngineCore) DisconnectSetupStarted() {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore_DisconnectSetupStarted(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "setupStarted")
	}
}

//export callbackQHelpEngineCoreSetupStarted
func callbackQHelpEngineCoreSetupStarted(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "setupStarted").(func())()
}

func (ptr *QHelpEngineCore) UnregisterDocumentation(namespaceName string) bool {
	if ptr.Pointer() != nil {
		return C.QHelpEngineCore_UnregisterDocumentation(C.QtObjectPtr(ptr.Pointer()), C.CString(namespaceName)) != 0
	}
	return false
}

func (ptr *QHelpEngineCore) ConnectWarning(f func(msg string)) {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore_ConnectWarning(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "warning", f)
	}
}

func (ptr *QHelpEngineCore) DisconnectWarning() {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore_DisconnectWarning(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "warning")
	}
}

//export callbackQHelpEngineCoreWarning
func callbackQHelpEngineCoreWarning(ptrName *C.char, msg *C.char) {
	qt.GetSignal(C.GoString(ptrName), "warning").(func(string))(C.GoString(msg))
}

func (ptr *QHelpEngineCore) DestroyQHelpEngineCore() {
	if ptr.Pointer() != nil {
		C.QHelpEngineCore_DestroyQHelpEngineCore(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
