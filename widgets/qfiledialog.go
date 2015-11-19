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

type QFileDialog_ITF interface {
	QDialog_ITF
	QFileDialog_PTR() *QFileDialog
}

func PointerFromQFileDialog(ptr QFileDialog_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QFileDialog_PTR().Pointer()
	}
	return nil
}

func NewQFileDialogFromPointer(ptr unsafe.Pointer) *QFileDialog {
	var n = new(QFileDialog)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QFileDialog_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QFileDialog) QFileDialog_PTR() *QFileDialog {
	return ptr
}

//QFileDialog::AcceptMode
type QFileDialog__AcceptMode int64

const (
	QFileDialog__AcceptOpen = QFileDialog__AcceptMode(0)
	QFileDialog__AcceptSave = QFileDialog__AcceptMode(1)
)

//QFileDialog::DialogLabel
type QFileDialog__DialogLabel int64

const (
	QFileDialog__LookIn   = QFileDialog__DialogLabel(0)
	QFileDialog__FileName = QFileDialog__DialogLabel(1)
	QFileDialog__FileType = QFileDialog__DialogLabel(2)
	QFileDialog__Accept   = QFileDialog__DialogLabel(3)
	QFileDialog__Reject   = QFileDialog__DialogLabel(4)
)

//QFileDialog::FileMode
type QFileDialog__FileMode int64

const (
	QFileDialog__AnyFile       = QFileDialog__FileMode(0)
	QFileDialog__ExistingFile  = QFileDialog__FileMode(1)
	QFileDialog__Directory     = QFileDialog__FileMode(2)
	QFileDialog__ExistingFiles = QFileDialog__FileMode(3)
	QFileDialog__DirectoryOnly = QFileDialog__FileMode(4)
)

//QFileDialog::Option
type QFileDialog__Option int64

const (
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
type QFileDialog__ViewMode int64

const (
	QFileDialog__Detail = QFileDialog__ViewMode(0)
	QFileDialog__List   = QFileDialog__ViewMode(1)
)

func NewQFileDialog(parent QWidget_ITF, flags core.Qt__WindowType) *QFileDialog {
	return NewQFileDialogFromPointer(C.QFileDialog_NewQFileDialog(PointerFromQWidget(parent), C.int(flags)))
}

func (ptr *QFileDialog) AcceptMode() QFileDialog__AcceptMode {
	if ptr.Pointer() != nil {
		return QFileDialog__AcceptMode(C.QFileDialog_AcceptMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFileDialog) ConfirmOverwrite() bool {
	if ptr.Pointer() != nil {
		return C.QFileDialog_ConfirmOverwrite(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QFileDialog) DefaultSuffix() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QFileDialog_DefaultSuffix(ptr.Pointer()))
	}
	return ""
}

func (ptr *QFileDialog) FileMode() QFileDialog__FileMode {
	if ptr.Pointer() != nil {
		return QFileDialog__FileMode(C.QFileDialog_FileMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFileDialog) IsNameFilterDetailsVisible() bool {
	if ptr.Pointer() != nil {
		return C.QFileDialog_IsNameFilterDetailsVisible(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QFileDialog) IsReadOnly() bool {
	if ptr.Pointer() != nil {
		return C.QFileDialog_IsReadOnly(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QFileDialog) Options() QFileDialog__Option {
	if ptr.Pointer() != nil {
		return QFileDialog__Option(C.QFileDialog_Options(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFileDialog) ResolveSymlinks() bool {
	if ptr.Pointer() != nil {
		return C.QFileDialog_ResolveSymlinks(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QFileDialog) SetAcceptMode(mode QFileDialog__AcceptMode) {
	if ptr.Pointer() != nil {
		C.QFileDialog_SetAcceptMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QFileDialog) SetConfirmOverwrite(enabled bool) {
	if ptr.Pointer() != nil {
		C.QFileDialog_SetConfirmOverwrite(ptr.Pointer(), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QFileDialog) SetDefaultSuffix(suffix string) {
	if ptr.Pointer() != nil {
		C.QFileDialog_SetDefaultSuffix(ptr.Pointer(), C.CString(suffix))
	}
}

func (ptr *QFileDialog) SetFileMode(mode QFileDialog__FileMode) {
	if ptr.Pointer() != nil {
		C.QFileDialog_SetFileMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QFileDialog) SetNameFilterDetailsVisible(enabled bool) {
	if ptr.Pointer() != nil {
		C.QFileDialog_SetNameFilterDetailsVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QFileDialog) SetOptions(options QFileDialog__Option) {
	if ptr.Pointer() != nil {
		C.QFileDialog_SetOptions(ptr.Pointer(), C.int(options))
	}
}

func (ptr *QFileDialog) SetReadOnly(enabled bool) {
	if ptr.Pointer() != nil {
		C.QFileDialog_SetReadOnly(ptr.Pointer(), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QFileDialog) SetResolveSymlinks(enabled bool) {
	if ptr.Pointer() != nil {
		C.QFileDialog_SetResolveSymlinks(ptr.Pointer(), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QFileDialog) SetViewMode(mode QFileDialog__ViewMode) {
	if ptr.Pointer() != nil {
		C.QFileDialog_SetViewMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QFileDialog) ViewMode() QFileDialog__ViewMode {
	if ptr.Pointer() != nil {
		return QFileDialog__ViewMode(C.QFileDialog_ViewMode(ptr.Pointer()))
	}
	return 0
}

func NewQFileDialog2(parent QWidget_ITF, caption string, directory string, filter string) *QFileDialog {
	return NewQFileDialogFromPointer(C.QFileDialog_NewQFileDialog2(PointerFromQWidget(parent), C.CString(caption), C.CString(directory), C.CString(filter)))
}

func (ptr *QFileDialog) ConnectCurrentChanged(f func(path string)) {
	if ptr.Pointer() != nil {
		C.QFileDialog_ConnectCurrentChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "currentChanged", f)
	}
}

func (ptr *QFileDialog) DisconnectCurrentChanged() {
	if ptr.Pointer() != nil {
		C.QFileDialog_DisconnectCurrentChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "currentChanged")
	}
}

//export callbackQFileDialogCurrentChanged
func callbackQFileDialogCurrentChanged(ptrName *C.char, path *C.char) {
	qt.GetSignal(C.GoString(ptrName), "currentChanged").(func(string))(C.GoString(path))
}

func (ptr *QFileDialog) Directory() *core.QDir {
	if ptr.Pointer() != nil {
		return core.NewQDirFromPointer(C.QFileDialog_Directory(ptr.Pointer()))
	}
	return nil
}

func (ptr *QFileDialog) ConnectDirectoryEntered(f func(directory string)) {
	if ptr.Pointer() != nil {
		C.QFileDialog_ConnectDirectoryEntered(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "directoryEntered", f)
	}
}

func (ptr *QFileDialog) DisconnectDirectoryEntered() {
	if ptr.Pointer() != nil {
		C.QFileDialog_DisconnectDirectoryEntered(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "directoryEntered")
	}
}

//export callbackQFileDialogDirectoryEntered
func callbackQFileDialogDirectoryEntered(ptrName *C.char, directory *C.char) {
	qt.GetSignal(C.GoString(ptrName), "directoryEntered").(func(string))(C.GoString(directory))
}

func (ptr *QFileDialog) ConnectFileSelected(f func(file string)) {
	if ptr.Pointer() != nil {
		C.QFileDialog_ConnectFileSelected(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "fileSelected", f)
	}
}

func (ptr *QFileDialog) DisconnectFileSelected() {
	if ptr.Pointer() != nil {
		C.QFileDialog_DisconnectFileSelected(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "fileSelected")
	}
}

//export callbackQFileDialogFileSelected
func callbackQFileDialogFileSelected(ptrName *C.char, file *C.char) {
	qt.GetSignal(C.GoString(ptrName), "fileSelected").(func(string))(C.GoString(file))
}

func (ptr *QFileDialog) ConnectFilesSelected(f func(selected []string)) {
	if ptr.Pointer() != nil {
		C.QFileDialog_ConnectFilesSelected(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "filesSelected", f)
	}
}

func (ptr *QFileDialog) DisconnectFilesSelected() {
	if ptr.Pointer() != nil {
		C.QFileDialog_DisconnectFilesSelected(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "filesSelected")
	}
}

//export callbackQFileDialogFilesSelected
func callbackQFileDialogFilesSelected(ptrName *C.char, selected *C.char) {
	qt.GetSignal(C.GoString(ptrName), "filesSelected").(func([]string))(strings.Split(C.GoString(selected), "|"))
}

func (ptr *QFileDialog) Filter() core.QDir__Filter {
	if ptr.Pointer() != nil {
		return core.QDir__Filter(C.QFileDialog_Filter(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFileDialog) ConnectFilterSelected(f func(filter string)) {
	if ptr.Pointer() != nil {
		C.QFileDialog_ConnectFilterSelected(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "filterSelected", f)
	}
}

func (ptr *QFileDialog) DisconnectFilterSelected() {
	if ptr.Pointer() != nil {
		C.QFileDialog_DisconnectFilterSelected(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "filterSelected")
	}
}

//export callbackQFileDialogFilterSelected
func callbackQFileDialogFilterSelected(ptrName *C.char, filter *C.char) {
	qt.GetSignal(C.GoString(ptrName), "filterSelected").(func(string))(C.GoString(filter))
}

func QFileDialog_GetExistingDirectory(parent QWidget_ITF, caption string, dir string, options QFileDialog__Option) string {
	return C.GoString(C.QFileDialog_QFileDialog_GetExistingDirectory(PointerFromQWidget(parent), C.CString(caption), C.CString(dir), C.int(options)))
}

func QFileDialog_GetOpenFileName(parent QWidget_ITF, caption string, dir string, filter string, selectedFilter string, options QFileDialog__Option) string {
	return C.GoString(C.QFileDialog_QFileDialog_GetOpenFileName(PointerFromQWidget(parent), C.CString(caption), C.CString(dir), C.CString(filter), C.CString(selectedFilter), C.int(options)))
}

func QFileDialog_GetOpenFileNames(parent QWidget_ITF, caption string, dir string, filter string, selectedFilter string, options QFileDialog__Option) []string {
	return strings.Split(C.GoString(C.QFileDialog_QFileDialog_GetOpenFileNames(PointerFromQWidget(parent), C.CString(caption), C.CString(dir), C.CString(filter), C.CString(selectedFilter), C.int(options))), "|")
}

func QFileDialog_GetSaveFileName(parent QWidget_ITF, caption string, dir string, filter string, selectedFilter string, options QFileDialog__Option) string {
	return C.GoString(C.QFileDialog_QFileDialog_GetSaveFileName(PointerFromQWidget(parent), C.CString(caption), C.CString(dir), C.CString(filter), C.CString(selectedFilter), C.int(options)))
}

func (ptr *QFileDialog) History() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QFileDialog_History(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QFileDialog) IconProvider() *QFileIconProvider {
	if ptr.Pointer() != nil {
		return NewQFileIconProviderFromPointer(C.QFileDialog_IconProvider(ptr.Pointer()))
	}
	return nil
}

func (ptr *QFileDialog) ItemDelegate() *QAbstractItemDelegate {
	if ptr.Pointer() != nil {
		return NewQAbstractItemDelegateFromPointer(C.QFileDialog_ItemDelegate(ptr.Pointer()))
	}
	return nil
}

func (ptr *QFileDialog) LabelText(label QFileDialog__DialogLabel) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QFileDialog_LabelText(ptr.Pointer(), C.int(label)))
	}
	return ""
}

func (ptr *QFileDialog) MimeTypeFilters() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QFileDialog_MimeTypeFilters(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QFileDialog) NameFilters() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QFileDialog_NameFilters(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QFileDialog) Open(receiver core.QObject_ITF, member string) {
	if ptr.Pointer() != nil {
		C.QFileDialog_Open(ptr.Pointer(), core.PointerFromQObject(receiver), C.CString(member))
	}
}

func (ptr *QFileDialog) ProxyModel() *core.QAbstractProxyModel {
	if ptr.Pointer() != nil {
		return core.NewQAbstractProxyModelFromPointer(C.QFileDialog_ProxyModel(ptr.Pointer()))
	}
	return nil
}

func (ptr *QFileDialog) RestoreState(state core.QByteArray_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QFileDialog_RestoreState(ptr.Pointer(), core.PointerFromQByteArray(state)) != 0
	}
	return false
}

func (ptr *QFileDialog) SaveState() *core.QByteArray {
	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QFileDialog_SaveState(ptr.Pointer()))
	}
	return nil
}

func (ptr *QFileDialog) SelectFile(filename string) {
	if ptr.Pointer() != nil {
		C.QFileDialog_SelectFile(ptr.Pointer(), C.CString(filename))
	}
}

func (ptr *QFileDialog) SelectMimeTypeFilter(filter string) {
	if ptr.Pointer() != nil {
		C.QFileDialog_SelectMimeTypeFilter(ptr.Pointer(), C.CString(filter))
	}
}

func (ptr *QFileDialog) SelectNameFilter(filter string) {
	if ptr.Pointer() != nil {
		C.QFileDialog_SelectNameFilter(ptr.Pointer(), C.CString(filter))
	}
}

func (ptr *QFileDialog) SelectUrl(url core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QFileDialog_SelectUrl(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QFileDialog) SelectedFiles() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QFileDialog_SelectedFiles(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QFileDialog) SelectedNameFilter() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QFileDialog_SelectedNameFilter(ptr.Pointer()))
	}
	return ""
}

func (ptr *QFileDialog) SetDirectory2(directory core.QDir_ITF) {
	if ptr.Pointer() != nil {
		C.QFileDialog_SetDirectory2(ptr.Pointer(), core.PointerFromQDir(directory))
	}
}

func (ptr *QFileDialog) SetDirectory(directory string) {
	if ptr.Pointer() != nil {
		C.QFileDialog_SetDirectory(ptr.Pointer(), C.CString(directory))
	}
}

func (ptr *QFileDialog) SetDirectoryUrl(directory core.QUrl_ITF) {
	if ptr.Pointer() != nil {
		C.QFileDialog_SetDirectoryUrl(ptr.Pointer(), core.PointerFromQUrl(directory))
	}
}

func (ptr *QFileDialog) SetFilter(filters core.QDir__Filter) {
	if ptr.Pointer() != nil {
		C.QFileDialog_SetFilter(ptr.Pointer(), C.int(filters))
	}
}

func (ptr *QFileDialog) SetHistory(paths []string) {
	if ptr.Pointer() != nil {
		C.QFileDialog_SetHistory(ptr.Pointer(), C.CString(strings.Join(paths, "|")))
	}
}

func (ptr *QFileDialog) SetIconProvider(provider QFileIconProvider_ITF) {
	if ptr.Pointer() != nil {
		C.QFileDialog_SetIconProvider(ptr.Pointer(), PointerFromQFileIconProvider(provider))
	}
}

func (ptr *QFileDialog) SetItemDelegate(delegate QAbstractItemDelegate_ITF) {
	if ptr.Pointer() != nil {
		C.QFileDialog_SetItemDelegate(ptr.Pointer(), PointerFromQAbstractItemDelegate(delegate))
	}
}

func (ptr *QFileDialog) SetLabelText(label QFileDialog__DialogLabel, text string) {
	if ptr.Pointer() != nil {
		C.QFileDialog_SetLabelText(ptr.Pointer(), C.int(label), C.CString(text))
	}
}

func (ptr *QFileDialog) SetMimeTypeFilters(filters []string) {
	if ptr.Pointer() != nil {
		C.QFileDialog_SetMimeTypeFilters(ptr.Pointer(), C.CString(strings.Join(filters, "|")))
	}
}

func (ptr *QFileDialog) SetNameFilter(filter string) {
	if ptr.Pointer() != nil {
		C.QFileDialog_SetNameFilter(ptr.Pointer(), C.CString(filter))
	}
}

func (ptr *QFileDialog) SetNameFilters(filters []string) {
	if ptr.Pointer() != nil {
		C.QFileDialog_SetNameFilters(ptr.Pointer(), C.CString(strings.Join(filters, "|")))
	}
}

func (ptr *QFileDialog) SetOption(option QFileDialog__Option, on bool) {
	if ptr.Pointer() != nil {
		C.QFileDialog_SetOption(ptr.Pointer(), C.int(option), C.int(qt.GoBoolToInt(on)))
	}
}

func (ptr *QFileDialog) SetProxyModel(proxyModel core.QAbstractProxyModel_ITF) {
	if ptr.Pointer() != nil {
		C.QFileDialog_SetProxyModel(ptr.Pointer(), core.PointerFromQAbstractProxyModel(proxyModel))
	}
}

func (ptr *QFileDialog) SetVisible(visible bool) {
	if ptr.Pointer() != nil {
		C.QFileDialog_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QFileDialog) TestOption(option QFileDialog__Option) bool {
	if ptr.Pointer() != nil {
		return C.QFileDialog_TestOption(ptr.Pointer(), C.int(option)) != 0
	}
	return false
}

func (ptr *QFileDialog) DestroyQFileDialog() {
	if ptr.Pointer() != nil {
		C.QFileDialog_DestroyQFileDialog(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
