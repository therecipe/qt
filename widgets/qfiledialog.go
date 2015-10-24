package widgets

//#include "qfiledialog.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"strings"
	"unsafe"
)

type QFileDialog struct {
	QDialog
}

type QFileDialogITF interface {
	QDialogITF
	QFileDialogPTR() *QFileDialog
}

func PointerFromQFileDialog(ptr QFileDialogITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QFileDialogPTR().Pointer()
	}
	return nil
}

func QFileDialogFromPointer(ptr unsafe.Pointer) *QFileDialog {
	var n = new(QFileDialog)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QFileDialog_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QFileDialog) QFileDialogPTR() *QFileDialog {
	return ptr
}

//QFileDialog::AcceptMode
type QFileDialog__AcceptMode int

var (
	QFileDialog__AcceptOpen = QFileDialog__AcceptMode(0)
	QFileDialog__AcceptSave = QFileDialog__AcceptMode(1)
)

//QFileDialog::DialogLabel
type QFileDialog__DialogLabel int

var (
	QFileDialog__LookIn   = QFileDialog__DialogLabel(0)
	QFileDialog__FileName = QFileDialog__DialogLabel(1)
	QFileDialog__FileType = QFileDialog__DialogLabel(2)
	QFileDialog__Accept   = QFileDialog__DialogLabel(3)
	QFileDialog__Reject   = QFileDialog__DialogLabel(4)
)

//QFileDialog::FileMode
type QFileDialog__FileMode int

var (
	QFileDialog__AnyFile       = QFileDialog__FileMode(0)
	QFileDialog__ExistingFile  = QFileDialog__FileMode(1)
	QFileDialog__Directory     = QFileDialog__FileMode(2)
	QFileDialog__ExistingFiles = QFileDialog__FileMode(3)
	QFileDialog__DirectoryOnly = QFileDialog__FileMode(4)
)

//QFileDialog::Option
type QFileDialog__Option int

var (
	QFileDialog__ShowDirsOnly                = QFileDialog__Option(0x00000001)
	QFileDialog__DontResolveSymlinks         = QFileDialog__Option(0x00000002)
	QFileDialog__DontConfirmOverwrite        = QFileDialog__Option(0x00000004)
	QFileDialog__DontUseSheet                = QFileDialog__Option(0x00000008)
	QFileDialog__DontUseNativeDialog         = QFileDialog__Option(0x00000010)
	QFileDialog__ReadOnly                    = QFileDialog__Option(0x00000020)
	QFileDialog__HideNameFilterDetails       = QFileDialog__Option(0x00000040)
	QFileDialog__DontUseCustomDirectoryIcons = QFileDialog__Option(0x00000080)
)

//QFileDialog::ViewMode
type QFileDialog__ViewMode int

var (
	QFileDialog__Detail = QFileDialog__ViewMode(0)
	QFileDialog__List   = QFileDialog__ViewMode(1)
)

func NewQFileDialog(parent QWidgetITF, flags core.Qt__WindowType) *QFileDialog {
	return QFileDialogFromPointer(unsafe.Pointer(C.QFileDialog_NewQFileDialog(C.QtObjectPtr(PointerFromQWidget(parent)), C.int(flags))))
}

func (ptr *QFileDialog) AcceptMode() QFileDialog__AcceptMode {
	if ptr.Pointer() != nil {
		return QFileDialog__AcceptMode(C.QFileDialog_AcceptMode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QFileDialog) ConfirmOverwrite() bool {
	if ptr.Pointer() != nil {
		return C.QFileDialog_ConfirmOverwrite(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QFileDialog) DefaultSuffix() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QFileDialog_DefaultSuffix(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QFileDialog) FileMode() QFileDialog__FileMode {
	if ptr.Pointer() != nil {
		return QFileDialog__FileMode(C.QFileDialog_FileMode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QFileDialog) IsNameFilterDetailsVisible() bool {
	if ptr.Pointer() != nil {
		return C.QFileDialog_IsNameFilterDetailsVisible(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QFileDialog) IsReadOnly() bool {
	if ptr.Pointer() != nil {
		return C.QFileDialog_IsReadOnly(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QFileDialog) Options() QFileDialog__Option {
	if ptr.Pointer() != nil {
		return QFileDialog__Option(C.QFileDialog_Options(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QFileDialog) ResolveSymlinks() bool {
	if ptr.Pointer() != nil {
		return C.QFileDialog_ResolveSymlinks(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QFileDialog) SetAcceptMode(mode QFileDialog__AcceptMode) {
	if ptr.Pointer() != nil {
		C.QFileDialog_SetAcceptMode(C.QtObjectPtr(ptr.Pointer()), C.int(mode))
	}
}

func (ptr *QFileDialog) SetConfirmOverwrite(enabled bool) {
	if ptr.Pointer() != nil {
		C.QFileDialog_SetConfirmOverwrite(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QFileDialog) SetDefaultSuffix(suffix string) {
	if ptr.Pointer() != nil {
		C.QFileDialog_SetDefaultSuffix(C.QtObjectPtr(ptr.Pointer()), C.CString(suffix))
	}
}

func (ptr *QFileDialog) SetFileMode(mode QFileDialog__FileMode) {
	if ptr.Pointer() != nil {
		C.QFileDialog_SetFileMode(C.QtObjectPtr(ptr.Pointer()), C.int(mode))
	}
}

func (ptr *QFileDialog) SetNameFilterDetailsVisible(enabled bool) {
	if ptr.Pointer() != nil {
		C.QFileDialog_SetNameFilterDetailsVisible(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QFileDialog) SetOptions(options QFileDialog__Option) {
	if ptr.Pointer() != nil {
		C.QFileDialog_SetOptions(C.QtObjectPtr(ptr.Pointer()), C.int(options))
	}
}

func (ptr *QFileDialog) SetReadOnly(enabled bool) {
	if ptr.Pointer() != nil {
		C.QFileDialog_SetReadOnly(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QFileDialog) SetResolveSymlinks(enabled bool) {
	if ptr.Pointer() != nil {
		C.QFileDialog_SetResolveSymlinks(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QFileDialog) SetViewMode(mode QFileDialog__ViewMode) {
	if ptr.Pointer() != nil {
		C.QFileDialog_SetViewMode(C.QtObjectPtr(ptr.Pointer()), C.int(mode))
	}
}

func (ptr *QFileDialog) ViewMode() QFileDialog__ViewMode {
	if ptr.Pointer() != nil {
		return QFileDialog__ViewMode(C.QFileDialog_ViewMode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func NewQFileDialog2(parent QWidgetITF, caption string, directory string, filter string) *QFileDialog {
	return QFileDialogFromPointer(unsafe.Pointer(C.QFileDialog_NewQFileDialog2(C.QtObjectPtr(PointerFromQWidget(parent)), C.CString(caption), C.CString(directory), C.CString(filter))))
}

func (ptr *QFileDialog) ConnectCurrentChanged(f func(path string)) {
	if ptr.Pointer() != nil {
		C.QFileDialog_ConnectCurrentChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "currentChanged", f)
	}
}

func (ptr *QFileDialog) DisconnectCurrentChanged() {
	if ptr.Pointer() != nil {
		C.QFileDialog_DisconnectCurrentChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "currentChanged")
	}
}

//export callbackQFileDialogCurrentChanged
func callbackQFileDialogCurrentChanged(ptrName *C.char, path *C.char) {
	qt.GetSignal(C.GoString(ptrName), "currentChanged").(func(string))(C.GoString(path))
}

func (ptr *QFileDialog) ConnectCurrentUrlChanged(f func(url string)) {
	if ptr.Pointer() != nil {
		C.QFileDialog_ConnectCurrentUrlChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "currentUrlChanged", f)
	}
}

func (ptr *QFileDialog) DisconnectCurrentUrlChanged() {
	if ptr.Pointer() != nil {
		C.QFileDialog_DisconnectCurrentUrlChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "currentUrlChanged")
	}
}

//export callbackQFileDialogCurrentUrlChanged
func callbackQFileDialogCurrentUrlChanged(ptrName *C.char, url *C.char) {
	qt.GetSignal(C.GoString(ptrName), "currentUrlChanged").(func(string))(C.GoString(url))
}

func (ptr *QFileDialog) ConnectDirectoryEntered(f func(directory string)) {
	if ptr.Pointer() != nil {
		C.QFileDialog_ConnectDirectoryEntered(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "directoryEntered", f)
	}
}

func (ptr *QFileDialog) DisconnectDirectoryEntered() {
	if ptr.Pointer() != nil {
		C.QFileDialog_DisconnectDirectoryEntered(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "directoryEntered")
	}
}

//export callbackQFileDialogDirectoryEntered
func callbackQFileDialogDirectoryEntered(ptrName *C.char, directory *C.char) {
	qt.GetSignal(C.GoString(ptrName), "directoryEntered").(func(string))(C.GoString(directory))
}

func (ptr *QFileDialog) DirectoryUrl() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QFileDialog_DirectoryUrl(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QFileDialog) ConnectDirectoryUrlEntered(f func(directory string)) {
	if ptr.Pointer() != nil {
		C.QFileDialog_ConnectDirectoryUrlEntered(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "directoryUrlEntered", f)
	}
}

func (ptr *QFileDialog) DisconnectDirectoryUrlEntered() {
	if ptr.Pointer() != nil {
		C.QFileDialog_DisconnectDirectoryUrlEntered(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "directoryUrlEntered")
	}
}

//export callbackQFileDialogDirectoryUrlEntered
func callbackQFileDialogDirectoryUrlEntered(ptrName *C.char, directory *C.char) {
	qt.GetSignal(C.GoString(ptrName), "directoryUrlEntered").(func(string))(C.GoString(directory))
}

func (ptr *QFileDialog) ConnectFileSelected(f func(file string)) {
	if ptr.Pointer() != nil {
		C.QFileDialog_ConnectFileSelected(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "fileSelected", f)
	}
}

func (ptr *QFileDialog) DisconnectFileSelected() {
	if ptr.Pointer() != nil {
		C.QFileDialog_DisconnectFileSelected(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "fileSelected")
	}
}

//export callbackQFileDialogFileSelected
func callbackQFileDialogFileSelected(ptrName *C.char, file *C.char) {
	qt.GetSignal(C.GoString(ptrName), "fileSelected").(func(string))(C.GoString(file))
}

func (ptr *QFileDialog) ConnectFilesSelected(f func(selected []string)) {
	if ptr.Pointer() != nil {
		C.QFileDialog_ConnectFilesSelected(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "filesSelected", f)
	}
}

func (ptr *QFileDialog) DisconnectFilesSelected() {
	if ptr.Pointer() != nil {
		C.QFileDialog_DisconnectFilesSelected(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "filesSelected")
	}
}

//export callbackQFileDialogFilesSelected
func callbackQFileDialogFilesSelected(ptrName *C.char, selected *C.char) {
	qt.GetSignal(C.GoString(ptrName), "filesSelected").(func([]string))(strings.Split(C.GoString(selected), "|"))
}

func (ptr *QFileDialog) Filter() core.QDir__Filter {
	if ptr.Pointer() != nil {
		return core.QDir__Filter(C.QFileDialog_Filter(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QFileDialog) ConnectFilterSelected(f func(filter string)) {
	if ptr.Pointer() != nil {
		C.QFileDialog_ConnectFilterSelected(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "filterSelected", f)
	}
}

func (ptr *QFileDialog) DisconnectFilterSelected() {
	if ptr.Pointer() != nil {
		C.QFileDialog_DisconnectFilterSelected(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "filterSelected")
	}
}

//export callbackQFileDialogFilterSelected
func callbackQFileDialogFilterSelected(ptrName *C.char, filter *C.char) {
	qt.GetSignal(C.GoString(ptrName), "filterSelected").(func(string))(C.GoString(filter))
}

func QFileDialog_GetExistingDirectory(parent QWidgetITF, caption string, dir string, options QFileDialog__Option) string {
	return C.GoString(C.QFileDialog_QFileDialog_GetExistingDirectory(C.QtObjectPtr(PointerFromQWidget(parent)), C.CString(caption), C.CString(dir), C.int(options)))
}

func QFileDialog_GetExistingDirectoryUrl(parent QWidgetITF, caption string, dir string, options QFileDialog__Option, supportedSchemes []string) string {
	return C.GoString(C.QFileDialog_QFileDialog_GetExistingDirectoryUrl(C.QtObjectPtr(PointerFromQWidget(parent)), C.CString(caption), C.CString(dir), C.int(options), C.CString(strings.Join(supportedSchemes, "|"))))
}

func QFileDialog_GetOpenFileName(parent QWidgetITF, caption string, dir string, filter string, selectedFilter string, options QFileDialog__Option) string {
	return C.GoString(C.QFileDialog_QFileDialog_GetOpenFileName(C.QtObjectPtr(PointerFromQWidget(parent)), C.CString(caption), C.CString(dir), C.CString(filter), C.CString(selectedFilter), C.int(options)))
}

func QFileDialog_GetOpenFileNames(parent QWidgetITF, caption string, dir string, filter string, selectedFilter string, options QFileDialog__Option) []string {
	return strings.Split(C.GoString(C.QFileDialog_QFileDialog_GetOpenFileNames(C.QtObjectPtr(PointerFromQWidget(parent)), C.CString(caption), C.CString(dir), C.CString(filter), C.CString(selectedFilter), C.int(options))), "|")
}

func QFileDialog_GetOpenFileUrl(parent QWidgetITF, caption string, dir string, filter string, selectedFilter string, options QFileDialog__Option, supportedSchemes []string) string {
	return C.GoString(C.QFileDialog_QFileDialog_GetOpenFileUrl(C.QtObjectPtr(PointerFromQWidget(parent)), C.CString(caption), C.CString(dir), C.CString(filter), C.CString(selectedFilter), C.int(options), C.CString(strings.Join(supportedSchemes, "|"))))
}

func QFileDialog_GetSaveFileName(parent QWidgetITF, caption string, dir string, filter string, selectedFilter string, options QFileDialog__Option) string {
	return C.GoString(C.QFileDialog_QFileDialog_GetSaveFileName(C.QtObjectPtr(PointerFromQWidget(parent)), C.CString(caption), C.CString(dir), C.CString(filter), C.CString(selectedFilter), C.int(options)))
}

func QFileDialog_GetSaveFileUrl(parent QWidgetITF, caption string, dir string, filter string, selectedFilter string, options QFileDialog__Option, supportedSchemes []string) string {
	return C.GoString(C.QFileDialog_QFileDialog_GetSaveFileUrl(C.QtObjectPtr(PointerFromQWidget(parent)), C.CString(caption), C.CString(dir), C.CString(filter), C.CString(selectedFilter), C.int(options), C.CString(strings.Join(supportedSchemes, "|"))))
}

func (ptr *QFileDialog) History() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QFileDialog_History(C.QtObjectPtr(ptr.Pointer()))), "|")
	}
	return make([]string, 0)
}

func (ptr *QFileDialog) IconProvider() *QFileIconProvider {
	if ptr.Pointer() != nil {
		return QFileIconProviderFromPointer(unsafe.Pointer(C.QFileDialog_IconProvider(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QFileDialog) ItemDelegate() *QAbstractItemDelegate {
	if ptr.Pointer() != nil {
		return QAbstractItemDelegateFromPointer(unsafe.Pointer(C.QFileDialog_ItemDelegate(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QFileDialog) LabelText(label QFileDialog__DialogLabel) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QFileDialog_LabelText(C.QtObjectPtr(ptr.Pointer()), C.int(label)))
	}
	return ""
}

func (ptr *QFileDialog) MimeTypeFilters() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QFileDialog_MimeTypeFilters(C.QtObjectPtr(ptr.Pointer()))), "|")
	}
	return make([]string, 0)
}

func (ptr *QFileDialog) NameFilters() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QFileDialog_NameFilters(C.QtObjectPtr(ptr.Pointer()))), "|")
	}
	return make([]string, 0)
}

func (ptr *QFileDialog) Open(receiver core.QObjectITF, member string) {
	if ptr.Pointer() != nil {
		C.QFileDialog_Open(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQObject(receiver)), C.CString(member))
	}
}

func (ptr *QFileDialog) ProxyModel() *core.QAbstractProxyModel {
	if ptr.Pointer() != nil {
		return core.QAbstractProxyModelFromPointer(unsafe.Pointer(C.QFileDialog_ProxyModel(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QFileDialog) RestoreState(state core.QByteArrayITF) bool {
	if ptr.Pointer() != nil {
		return C.QFileDialog_RestoreState(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQByteArray(state))) != 0
	}
	return false
}

func (ptr *QFileDialog) SelectFile(filename string) {
	if ptr.Pointer() != nil {
		C.QFileDialog_SelectFile(C.QtObjectPtr(ptr.Pointer()), C.CString(filename))
	}
}

func (ptr *QFileDialog) SelectMimeTypeFilter(filter string) {
	if ptr.Pointer() != nil {
		C.QFileDialog_SelectMimeTypeFilter(C.QtObjectPtr(ptr.Pointer()), C.CString(filter))
	}
}

func (ptr *QFileDialog) SelectNameFilter(filter string) {
	if ptr.Pointer() != nil {
		C.QFileDialog_SelectNameFilter(C.QtObjectPtr(ptr.Pointer()), C.CString(filter))
	}
}

func (ptr *QFileDialog) SelectUrl(url string) {
	if ptr.Pointer() != nil {
		C.QFileDialog_SelectUrl(C.QtObjectPtr(ptr.Pointer()), C.CString(url))
	}
}

func (ptr *QFileDialog) SelectedFiles() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QFileDialog_SelectedFiles(C.QtObjectPtr(ptr.Pointer()))), "|")
	}
	return make([]string, 0)
}

func (ptr *QFileDialog) SelectedNameFilter() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QFileDialog_SelectedNameFilter(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QFileDialog) SetDirectory2(directory core.QDirITF) {
	if ptr.Pointer() != nil {
		C.QFileDialog_SetDirectory2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQDir(directory)))
	}
}

func (ptr *QFileDialog) SetDirectory(directory string) {
	if ptr.Pointer() != nil {
		C.QFileDialog_SetDirectory(C.QtObjectPtr(ptr.Pointer()), C.CString(directory))
	}
}

func (ptr *QFileDialog) SetDirectoryUrl(directory string) {
	if ptr.Pointer() != nil {
		C.QFileDialog_SetDirectoryUrl(C.QtObjectPtr(ptr.Pointer()), C.CString(directory))
	}
}

func (ptr *QFileDialog) SetFilter(filters core.QDir__Filter) {
	if ptr.Pointer() != nil {
		C.QFileDialog_SetFilter(C.QtObjectPtr(ptr.Pointer()), C.int(filters))
	}
}

func (ptr *QFileDialog) SetHistory(paths []string) {
	if ptr.Pointer() != nil {
		C.QFileDialog_SetHistory(C.QtObjectPtr(ptr.Pointer()), C.CString(strings.Join(paths, "|")))
	}
}

func (ptr *QFileDialog) SetIconProvider(provider QFileIconProviderITF) {
	if ptr.Pointer() != nil {
		C.QFileDialog_SetIconProvider(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQFileIconProvider(provider)))
	}
}

func (ptr *QFileDialog) SetItemDelegate(delegate QAbstractItemDelegateITF) {
	if ptr.Pointer() != nil {
		C.QFileDialog_SetItemDelegate(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQAbstractItemDelegate(delegate)))
	}
}

func (ptr *QFileDialog) SetLabelText(label QFileDialog__DialogLabel, text string) {
	if ptr.Pointer() != nil {
		C.QFileDialog_SetLabelText(C.QtObjectPtr(ptr.Pointer()), C.int(label), C.CString(text))
	}
}

func (ptr *QFileDialog) SetMimeTypeFilters(filters []string) {
	if ptr.Pointer() != nil {
		C.QFileDialog_SetMimeTypeFilters(C.QtObjectPtr(ptr.Pointer()), C.CString(strings.Join(filters, "|")))
	}
}

func (ptr *QFileDialog) SetNameFilter(filter string) {
	if ptr.Pointer() != nil {
		C.QFileDialog_SetNameFilter(C.QtObjectPtr(ptr.Pointer()), C.CString(filter))
	}
}

func (ptr *QFileDialog) SetNameFilters(filters []string) {
	if ptr.Pointer() != nil {
		C.QFileDialog_SetNameFilters(C.QtObjectPtr(ptr.Pointer()), C.CString(strings.Join(filters, "|")))
	}
}

func (ptr *QFileDialog) SetOption(option QFileDialog__Option, on bool) {
	if ptr.Pointer() != nil {
		C.QFileDialog_SetOption(C.QtObjectPtr(ptr.Pointer()), C.int(option), C.int(qt.GoBoolToInt(on)))
	}
}

func (ptr *QFileDialog) SetProxyModel(proxyModel core.QAbstractProxyModelITF) {
	if ptr.Pointer() != nil {
		C.QFileDialog_SetProxyModel(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQAbstractProxyModel(proxyModel)))
	}
}

func (ptr *QFileDialog) SetVisible(visible bool) {
	if ptr.Pointer() != nil {
		C.QFileDialog_SetVisible(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QFileDialog) TestOption(option QFileDialog__Option) bool {
	if ptr.Pointer() != nil {
		return C.QFileDialog_TestOption(C.QtObjectPtr(ptr.Pointer()), C.int(option)) != 0
	}
	return false
}

func (ptr *QFileDialog) ConnectUrlSelected(f func(url string)) {
	if ptr.Pointer() != nil {
		C.QFileDialog_ConnectUrlSelected(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "urlSelected", f)
	}
}

func (ptr *QFileDialog) DisconnectUrlSelected() {
	if ptr.Pointer() != nil {
		C.QFileDialog_DisconnectUrlSelected(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "urlSelected")
	}
}

//export callbackQFileDialogUrlSelected
func callbackQFileDialogUrlSelected(ptrName *C.char, url *C.char) {
	qt.GetSignal(C.GoString(ptrName), "urlSelected").(func(string))(C.GoString(url))
}

func (ptr *QFileDialog) DestroyQFileDialog() {
	if ptr.Pointer() != nil {
		C.QFileDialog_DestroyQFileDialog(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
