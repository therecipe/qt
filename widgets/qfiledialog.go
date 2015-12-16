package widgets

//#include "widgets.h"
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
	for len(n.ObjectName()) < len("QFileDialog_") {
		n.SetObjectName("QFileDialog_" + qt.Identifier())
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
	defer qt.Recovering("QFileDialog::QFileDialog")

	return NewQFileDialogFromPointer(C.QFileDialog_NewQFileDialog(PointerFromQWidget(parent), C.int(flags)))
}

func (ptr *QFileDialog) AcceptMode() QFileDialog__AcceptMode {
	defer qt.Recovering("QFileDialog::acceptMode")

	if ptr.Pointer() != nil {
		return QFileDialog__AcceptMode(C.QFileDialog_AcceptMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFileDialog) ConfirmOverwrite() bool {
	defer qt.Recovering("QFileDialog::confirmOverwrite")

	if ptr.Pointer() != nil {
		return C.QFileDialog_ConfirmOverwrite(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QFileDialog) DefaultSuffix() string {
	defer qt.Recovering("QFileDialog::defaultSuffix")

	if ptr.Pointer() != nil {
		return C.GoString(C.QFileDialog_DefaultSuffix(ptr.Pointer()))
	}
	return ""
}

func (ptr *QFileDialog) FileMode() QFileDialog__FileMode {
	defer qt.Recovering("QFileDialog::fileMode")

	if ptr.Pointer() != nil {
		return QFileDialog__FileMode(C.QFileDialog_FileMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFileDialog) IsNameFilterDetailsVisible() bool {
	defer qt.Recovering("QFileDialog::isNameFilterDetailsVisible")

	if ptr.Pointer() != nil {
		return C.QFileDialog_IsNameFilterDetailsVisible(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QFileDialog) IsReadOnly() bool {
	defer qt.Recovering("QFileDialog::isReadOnly")

	if ptr.Pointer() != nil {
		return C.QFileDialog_IsReadOnly(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QFileDialog) Options() QFileDialog__Option {
	defer qt.Recovering("QFileDialog::options")

	if ptr.Pointer() != nil {
		return QFileDialog__Option(C.QFileDialog_Options(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFileDialog) ResolveSymlinks() bool {
	defer qt.Recovering("QFileDialog::resolveSymlinks")

	if ptr.Pointer() != nil {
		return C.QFileDialog_ResolveSymlinks(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QFileDialog) SetAcceptMode(mode QFileDialog__AcceptMode) {
	defer qt.Recovering("QFileDialog::setAcceptMode")

	if ptr.Pointer() != nil {
		C.QFileDialog_SetAcceptMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QFileDialog) SetConfirmOverwrite(enabled bool) {
	defer qt.Recovering("QFileDialog::setConfirmOverwrite")

	if ptr.Pointer() != nil {
		C.QFileDialog_SetConfirmOverwrite(ptr.Pointer(), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QFileDialog) SetDefaultSuffix(suffix string) {
	defer qt.Recovering("QFileDialog::setDefaultSuffix")

	if ptr.Pointer() != nil {
		C.QFileDialog_SetDefaultSuffix(ptr.Pointer(), C.CString(suffix))
	}
}

func (ptr *QFileDialog) SetFileMode(mode QFileDialog__FileMode) {
	defer qt.Recovering("QFileDialog::setFileMode")

	if ptr.Pointer() != nil {
		C.QFileDialog_SetFileMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QFileDialog) SetNameFilterDetailsVisible(enabled bool) {
	defer qt.Recovering("QFileDialog::setNameFilterDetailsVisible")

	if ptr.Pointer() != nil {
		C.QFileDialog_SetNameFilterDetailsVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QFileDialog) SetOptions(options QFileDialog__Option) {
	defer qt.Recovering("QFileDialog::setOptions")

	if ptr.Pointer() != nil {
		C.QFileDialog_SetOptions(ptr.Pointer(), C.int(options))
	}
}

func (ptr *QFileDialog) SetReadOnly(enabled bool) {
	defer qt.Recovering("QFileDialog::setReadOnly")

	if ptr.Pointer() != nil {
		C.QFileDialog_SetReadOnly(ptr.Pointer(), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QFileDialog) SetResolveSymlinks(enabled bool) {
	defer qt.Recovering("QFileDialog::setResolveSymlinks")

	if ptr.Pointer() != nil {
		C.QFileDialog_SetResolveSymlinks(ptr.Pointer(), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QFileDialog) SetViewMode(mode QFileDialog__ViewMode) {
	defer qt.Recovering("QFileDialog::setViewMode")

	if ptr.Pointer() != nil {
		C.QFileDialog_SetViewMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QFileDialog) ViewMode() QFileDialog__ViewMode {
	defer qt.Recovering("QFileDialog::viewMode")

	if ptr.Pointer() != nil {
		return QFileDialog__ViewMode(C.QFileDialog_ViewMode(ptr.Pointer()))
	}
	return 0
}

func NewQFileDialog2(parent QWidget_ITF, caption string, directory string, filter string) *QFileDialog {
	defer qt.Recovering("QFileDialog::QFileDialog")

	return NewQFileDialogFromPointer(C.QFileDialog_NewQFileDialog2(PointerFromQWidget(parent), C.CString(caption), C.CString(directory), C.CString(filter)))
}

func (ptr *QFileDialog) ConnectAccept(f func()) {
	defer qt.Recovering("connect QFileDialog::accept")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "accept", f)
	}
}

func (ptr *QFileDialog) DisconnectAccept() {
	defer qt.Recovering("disconnect QFileDialog::accept")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "accept")
	}
}

//export callbackQFileDialogAccept
func callbackQFileDialogAccept(ptrName *C.char) bool {
	defer qt.Recovering("callback QFileDialog::accept")

	var signal = qt.GetSignal(C.GoString(ptrName), "accept")
	if signal != nil {
		defer signal.(func())()
		return true
	}
	return false

}

func (ptr *QFileDialog) ConnectChangeEvent(f func(e *core.QEvent)) {
	defer qt.Recovering("connect QFileDialog::changeEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "changeEvent", f)
	}
}

func (ptr *QFileDialog) DisconnectChangeEvent() {
	defer qt.Recovering("disconnect QFileDialog::changeEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "changeEvent")
	}
}

//export callbackQFileDialogChangeEvent
func callbackQFileDialogChangeEvent(ptrName *C.char, e unsafe.Pointer) bool {
	defer qt.Recovering("callback QFileDialog::changeEvent")

	var signal = qt.GetSignal(C.GoString(ptrName), "changeEvent")
	if signal != nil {
		defer signal.(func(*core.QEvent))(core.NewQEventFromPointer(e))
		return true
	}
	return false

}

func (ptr *QFileDialog) ConnectCurrentChanged(f func(path string)) {
	defer qt.Recovering("connect QFileDialog::currentChanged")

	if ptr.Pointer() != nil {
		C.QFileDialog_ConnectCurrentChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "currentChanged", f)
	}
}

func (ptr *QFileDialog) DisconnectCurrentChanged() {
	defer qt.Recovering("disconnect QFileDialog::currentChanged")

	if ptr.Pointer() != nil {
		C.QFileDialog_DisconnectCurrentChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "currentChanged")
	}
}

//export callbackQFileDialogCurrentChanged
func callbackQFileDialogCurrentChanged(ptrName *C.char, path *C.char) {
	defer qt.Recovering("callback QFileDialog::currentChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "currentChanged")
	if signal != nil {
		signal.(func(string))(C.GoString(path))
	}

}

func (ptr *QFileDialog) ConnectCurrentUrlChanged(f func(url *core.QUrl)) {
	defer qt.Recovering("connect QFileDialog::currentUrlChanged")

	if ptr.Pointer() != nil {
		C.QFileDialog_ConnectCurrentUrlChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "currentUrlChanged", f)
	}
}

func (ptr *QFileDialog) DisconnectCurrentUrlChanged() {
	defer qt.Recovering("disconnect QFileDialog::currentUrlChanged")

	if ptr.Pointer() != nil {
		C.QFileDialog_DisconnectCurrentUrlChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "currentUrlChanged")
	}
}

//export callbackQFileDialogCurrentUrlChanged
func callbackQFileDialogCurrentUrlChanged(ptrName *C.char, url unsafe.Pointer) {
	defer qt.Recovering("callback QFileDialog::currentUrlChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "currentUrlChanged")
	if signal != nil {
		signal.(func(*core.QUrl))(core.NewQUrlFromPointer(url))
	}

}

func (ptr *QFileDialog) Directory() *core.QDir {
	defer qt.Recovering("QFileDialog::directory")

	if ptr.Pointer() != nil {
		return core.NewQDirFromPointer(C.QFileDialog_Directory(ptr.Pointer()))
	}
	return nil
}

func (ptr *QFileDialog) ConnectDirectoryEntered(f func(directory string)) {
	defer qt.Recovering("connect QFileDialog::directoryEntered")

	if ptr.Pointer() != nil {
		C.QFileDialog_ConnectDirectoryEntered(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "directoryEntered", f)
	}
}

func (ptr *QFileDialog) DisconnectDirectoryEntered() {
	defer qt.Recovering("disconnect QFileDialog::directoryEntered")

	if ptr.Pointer() != nil {
		C.QFileDialog_DisconnectDirectoryEntered(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "directoryEntered")
	}
}

//export callbackQFileDialogDirectoryEntered
func callbackQFileDialogDirectoryEntered(ptrName *C.char, directory *C.char) {
	defer qt.Recovering("callback QFileDialog::directoryEntered")

	var signal = qt.GetSignal(C.GoString(ptrName), "directoryEntered")
	if signal != nil {
		signal.(func(string))(C.GoString(directory))
	}

}

func (ptr *QFileDialog) DirectoryUrl() *core.QUrl {
	defer qt.Recovering("QFileDialog::directoryUrl")

	if ptr.Pointer() != nil {
		return core.NewQUrlFromPointer(C.QFileDialog_DirectoryUrl(ptr.Pointer()))
	}
	return nil
}

func (ptr *QFileDialog) ConnectDirectoryUrlEntered(f func(directory *core.QUrl)) {
	defer qt.Recovering("connect QFileDialog::directoryUrlEntered")

	if ptr.Pointer() != nil {
		C.QFileDialog_ConnectDirectoryUrlEntered(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "directoryUrlEntered", f)
	}
}

func (ptr *QFileDialog) DisconnectDirectoryUrlEntered() {
	defer qt.Recovering("disconnect QFileDialog::directoryUrlEntered")

	if ptr.Pointer() != nil {
		C.QFileDialog_DisconnectDirectoryUrlEntered(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "directoryUrlEntered")
	}
}

//export callbackQFileDialogDirectoryUrlEntered
func callbackQFileDialogDirectoryUrlEntered(ptrName *C.char, directory unsafe.Pointer) {
	defer qt.Recovering("callback QFileDialog::directoryUrlEntered")

	var signal = qt.GetSignal(C.GoString(ptrName), "directoryUrlEntered")
	if signal != nil {
		signal.(func(*core.QUrl))(core.NewQUrlFromPointer(directory))
	}

}

func (ptr *QFileDialog) ConnectDone(f func(result int)) {
	defer qt.Recovering("connect QFileDialog::done")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "done", f)
	}
}

func (ptr *QFileDialog) DisconnectDone() {
	defer qt.Recovering("disconnect QFileDialog::done")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "done")
	}
}

//export callbackQFileDialogDone
func callbackQFileDialogDone(ptrName *C.char, result C.int) bool {
	defer qt.Recovering("callback QFileDialog::done")

	var signal = qt.GetSignal(C.GoString(ptrName), "done")
	if signal != nil {
		defer signal.(func(int))(int(result))
		return true
	}
	return false

}

func (ptr *QFileDialog) ConnectFileSelected(f func(file string)) {
	defer qt.Recovering("connect QFileDialog::fileSelected")

	if ptr.Pointer() != nil {
		C.QFileDialog_ConnectFileSelected(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "fileSelected", f)
	}
}

func (ptr *QFileDialog) DisconnectFileSelected() {
	defer qt.Recovering("disconnect QFileDialog::fileSelected")

	if ptr.Pointer() != nil {
		C.QFileDialog_DisconnectFileSelected(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "fileSelected")
	}
}

//export callbackQFileDialogFileSelected
func callbackQFileDialogFileSelected(ptrName *C.char, file *C.char) {
	defer qt.Recovering("callback QFileDialog::fileSelected")

	var signal = qt.GetSignal(C.GoString(ptrName), "fileSelected")
	if signal != nil {
		signal.(func(string))(C.GoString(file))
	}

}

func (ptr *QFileDialog) ConnectFilesSelected(f func(selected []string)) {
	defer qt.Recovering("connect QFileDialog::filesSelected")

	if ptr.Pointer() != nil {
		C.QFileDialog_ConnectFilesSelected(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "filesSelected", f)
	}
}

func (ptr *QFileDialog) DisconnectFilesSelected() {
	defer qt.Recovering("disconnect QFileDialog::filesSelected")

	if ptr.Pointer() != nil {
		C.QFileDialog_DisconnectFilesSelected(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "filesSelected")
	}
}

//export callbackQFileDialogFilesSelected
func callbackQFileDialogFilesSelected(ptrName *C.char, selected *C.char) {
	defer qt.Recovering("callback QFileDialog::filesSelected")

	var signal = qt.GetSignal(C.GoString(ptrName), "filesSelected")
	if signal != nil {
		signal.(func([]string))(strings.Split(C.GoString(selected), ",,,"))
	}

}

func (ptr *QFileDialog) Filter() core.QDir__Filter {
	defer qt.Recovering("QFileDialog::filter")

	if ptr.Pointer() != nil {
		return core.QDir__Filter(C.QFileDialog_Filter(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFileDialog) ConnectFilterSelected(f func(filter string)) {
	defer qt.Recovering("connect QFileDialog::filterSelected")

	if ptr.Pointer() != nil {
		C.QFileDialog_ConnectFilterSelected(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "filterSelected", f)
	}
}

func (ptr *QFileDialog) DisconnectFilterSelected() {
	defer qt.Recovering("disconnect QFileDialog::filterSelected")

	if ptr.Pointer() != nil {
		C.QFileDialog_DisconnectFilterSelected(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "filterSelected")
	}
}

//export callbackQFileDialogFilterSelected
func callbackQFileDialogFilterSelected(ptrName *C.char, filter *C.char) {
	defer qt.Recovering("callback QFileDialog::filterSelected")

	var signal = qt.GetSignal(C.GoString(ptrName), "filterSelected")
	if signal != nil {
		signal.(func(string))(C.GoString(filter))
	}

}

func QFileDialog_GetExistingDirectory(parent QWidget_ITF, caption string, dir string, options QFileDialog__Option) string {
	defer qt.Recovering("QFileDialog::getExistingDirectory")

	return C.GoString(C.QFileDialog_QFileDialog_GetExistingDirectory(PointerFromQWidget(parent), C.CString(caption), C.CString(dir), C.int(options)))
}

func QFileDialog_GetExistingDirectoryUrl(parent QWidget_ITF, caption string, dir core.QUrl_ITF, options QFileDialog__Option, supportedSchemes []string) *core.QUrl {
	defer qt.Recovering("QFileDialog::getExistingDirectoryUrl")

	return core.NewQUrlFromPointer(C.QFileDialog_QFileDialog_GetExistingDirectoryUrl(PointerFromQWidget(parent), C.CString(caption), core.PointerFromQUrl(dir), C.int(options), C.CString(strings.Join(supportedSchemes, ",,,"))))
}

func QFileDialog_GetOpenFileName(parent QWidget_ITF, caption string, dir string, filter string, selectedFilter string, options QFileDialog__Option) string {
	defer qt.Recovering("QFileDialog::getOpenFileName")

	return C.GoString(C.QFileDialog_QFileDialog_GetOpenFileName(PointerFromQWidget(parent), C.CString(caption), C.CString(dir), C.CString(filter), C.CString(selectedFilter), C.int(options)))
}

func QFileDialog_GetOpenFileNames(parent QWidget_ITF, caption string, dir string, filter string, selectedFilter string, options QFileDialog__Option) []string {
	defer qt.Recovering("QFileDialog::getOpenFileNames")

	return strings.Split(C.GoString(C.QFileDialog_QFileDialog_GetOpenFileNames(PointerFromQWidget(parent), C.CString(caption), C.CString(dir), C.CString(filter), C.CString(selectedFilter), C.int(options))), ",,,")
}

func QFileDialog_GetOpenFileUrl(parent QWidget_ITF, caption string, dir core.QUrl_ITF, filter string, selectedFilter string, options QFileDialog__Option, supportedSchemes []string) *core.QUrl {
	defer qt.Recovering("QFileDialog::getOpenFileUrl")

	return core.NewQUrlFromPointer(C.QFileDialog_QFileDialog_GetOpenFileUrl(PointerFromQWidget(parent), C.CString(caption), core.PointerFromQUrl(dir), C.CString(filter), C.CString(selectedFilter), C.int(options), C.CString(strings.Join(supportedSchemes, ",,,"))))
}

func QFileDialog_GetSaveFileName(parent QWidget_ITF, caption string, dir string, filter string, selectedFilter string, options QFileDialog__Option) string {
	defer qt.Recovering("QFileDialog::getSaveFileName")

	return C.GoString(C.QFileDialog_QFileDialog_GetSaveFileName(PointerFromQWidget(parent), C.CString(caption), C.CString(dir), C.CString(filter), C.CString(selectedFilter), C.int(options)))
}

func QFileDialog_GetSaveFileUrl(parent QWidget_ITF, caption string, dir core.QUrl_ITF, filter string, selectedFilter string, options QFileDialog__Option, supportedSchemes []string) *core.QUrl {
	defer qt.Recovering("QFileDialog::getSaveFileUrl")

	return core.NewQUrlFromPointer(C.QFileDialog_QFileDialog_GetSaveFileUrl(PointerFromQWidget(parent), C.CString(caption), core.PointerFromQUrl(dir), C.CString(filter), C.CString(selectedFilter), C.int(options), C.CString(strings.Join(supportedSchemes, ",,,"))))
}

func (ptr *QFileDialog) History() []string {
	defer qt.Recovering("QFileDialog::history")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QFileDialog_History(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QFileDialog) IconProvider() *QFileIconProvider {
	defer qt.Recovering("QFileDialog::iconProvider")

	if ptr.Pointer() != nil {
		return NewQFileIconProviderFromPointer(C.QFileDialog_IconProvider(ptr.Pointer()))
	}
	return nil
}

func (ptr *QFileDialog) ItemDelegate() *QAbstractItemDelegate {
	defer qt.Recovering("QFileDialog::itemDelegate")

	if ptr.Pointer() != nil {
		return NewQAbstractItemDelegateFromPointer(C.QFileDialog_ItemDelegate(ptr.Pointer()))
	}
	return nil
}

func (ptr *QFileDialog) LabelText(label QFileDialog__DialogLabel) string {
	defer qt.Recovering("QFileDialog::labelText")

	if ptr.Pointer() != nil {
		return C.GoString(C.QFileDialog_LabelText(ptr.Pointer(), C.int(label)))
	}
	return ""
}

func (ptr *QFileDialog) MimeTypeFilters() []string {
	defer qt.Recovering("QFileDialog::mimeTypeFilters")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QFileDialog_MimeTypeFilters(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QFileDialog) NameFilters() []string {
	defer qt.Recovering("QFileDialog::nameFilters")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QFileDialog_NameFilters(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QFileDialog) Open(receiver core.QObject_ITF, member string) {
	defer qt.Recovering("QFileDialog::open")

	if ptr.Pointer() != nil {
		C.QFileDialog_Open(ptr.Pointer(), core.PointerFromQObject(receiver), C.CString(member))
	}
}

func (ptr *QFileDialog) ProxyModel() *core.QAbstractProxyModel {
	defer qt.Recovering("QFileDialog::proxyModel")

	if ptr.Pointer() != nil {
		return core.NewQAbstractProxyModelFromPointer(C.QFileDialog_ProxyModel(ptr.Pointer()))
	}
	return nil
}

func (ptr *QFileDialog) RestoreState(state core.QByteArray_ITF) bool {
	defer qt.Recovering("QFileDialog::restoreState")

	if ptr.Pointer() != nil {
		return C.QFileDialog_RestoreState(ptr.Pointer(), core.PointerFromQByteArray(state)) != 0
	}
	return false
}

func (ptr *QFileDialog) SaveState() *core.QByteArray {
	defer qt.Recovering("QFileDialog::saveState")

	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QFileDialog_SaveState(ptr.Pointer()))
	}
	return nil
}

func (ptr *QFileDialog) SelectFile(filename string) {
	defer qt.Recovering("QFileDialog::selectFile")

	if ptr.Pointer() != nil {
		C.QFileDialog_SelectFile(ptr.Pointer(), C.CString(filename))
	}
}

func (ptr *QFileDialog) SelectMimeTypeFilter(filter string) {
	defer qt.Recovering("QFileDialog::selectMimeTypeFilter")

	if ptr.Pointer() != nil {
		C.QFileDialog_SelectMimeTypeFilter(ptr.Pointer(), C.CString(filter))
	}
}

func (ptr *QFileDialog) SelectNameFilter(filter string) {
	defer qt.Recovering("QFileDialog::selectNameFilter")

	if ptr.Pointer() != nil {
		C.QFileDialog_SelectNameFilter(ptr.Pointer(), C.CString(filter))
	}
}

func (ptr *QFileDialog) SelectUrl(url core.QUrl_ITF) {
	defer qt.Recovering("QFileDialog::selectUrl")

	if ptr.Pointer() != nil {
		C.QFileDialog_SelectUrl(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QFileDialog) SelectedFiles() []string {
	defer qt.Recovering("QFileDialog::selectedFiles")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QFileDialog_SelectedFiles(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QFileDialog) SelectedNameFilter() string {
	defer qt.Recovering("QFileDialog::selectedNameFilter")

	if ptr.Pointer() != nil {
		return C.GoString(C.QFileDialog_SelectedNameFilter(ptr.Pointer()))
	}
	return ""
}

func (ptr *QFileDialog) SetDirectory2(directory core.QDir_ITF) {
	defer qt.Recovering("QFileDialog::setDirectory")

	if ptr.Pointer() != nil {
		C.QFileDialog_SetDirectory2(ptr.Pointer(), core.PointerFromQDir(directory))
	}
}

func (ptr *QFileDialog) SetDirectory(directory string) {
	defer qt.Recovering("QFileDialog::setDirectory")

	if ptr.Pointer() != nil {
		C.QFileDialog_SetDirectory(ptr.Pointer(), C.CString(directory))
	}
}

func (ptr *QFileDialog) SetDirectoryUrl(directory core.QUrl_ITF) {
	defer qt.Recovering("QFileDialog::setDirectoryUrl")

	if ptr.Pointer() != nil {
		C.QFileDialog_SetDirectoryUrl(ptr.Pointer(), core.PointerFromQUrl(directory))
	}
}

func (ptr *QFileDialog) SetFilter(filters core.QDir__Filter) {
	defer qt.Recovering("QFileDialog::setFilter")

	if ptr.Pointer() != nil {
		C.QFileDialog_SetFilter(ptr.Pointer(), C.int(filters))
	}
}

func (ptr *QFileDialog) SetHistory(paths []string) {
	defer qt.Recovering("QFileDialog::setHistory")

	if ptr.Pointer() != nil {
		C.QFileDialog_SetHistory(ptr.Pointer(), C.CString(strings.Join(paths, ",,,")))
	}
}

func (ptr *QFileDialog) SetIconProvider(provider QFileIconProvider_ITF) {
	defer qt.Recovering("QFileDialog::setIconProvider")

	if ptr.Pointer() != nil {
		C.QFileDialog_SetIconProvider(ptr.Pointer(), PointerFromQFileIconProvider(provider))
	}
}

func (ptr *QFileDialog) SetItemDelegate(delegate QAbstractItemDelegate_ITF) {
	defer qt.Recovering("QFileDialog::setItemDelegate")

	if ptr.Pointer() != nil {
		C.QFileDialog_SetItemDelegate(ptr.Pointer(), PointerFromQAbstractItemDelegate(delegate))
	}
}

func (ptr *QFileDialog) SetLabelText(label QFileDialog__DialogLabel, text string) {
	defer qt.Recovering("QFileDialog::setLabelText")

	if ptr.Pointer() != nil {
		C.QFileDialog_SetLabelText(ptr.Pointer(), C.int(label), C.CString(text))
	}
}

func (ptr *QFileDialog) SetMimeTypeFilters(filters []string) {
	defer qt.Recovering("QFileDialog::setMimeTypeFilters")

	if ptr.Pointer() != nil {
		C.QFileDialog_SetMimeTypeFilters(ptr.Pointer(), C.CString(strings.Join(filters, ",,,")))
	}
}

func (ptr *QFileDialog) SetNameFilter(filter string) {
	defer qt.Recovering("QFileDialog::setNameFilter")

	if ptr.Pointer() != nil {
		C.QFileDialog_SetNameFilter(ptr.Pointer(), C.CString(filter))
	}
}

func (ptr *QFileDialog) SetNameFilters(filters []string) {
	defer qt.Recovering("QFileDialog::setNameFilters")

	if ptr.Pointer() != nil {
		C.QFileDialog_SetNameFilters(ptr.Pointer(), C.CString(strings.Join(filters, ",,,")))
	}
}

func (ptr *QFileDialog) SetOption(option QFileDialog__Option, on bool) {
	defer qt.Recovering("QFileDialog::setOption")

	if ptr.Pointer() != nil {
		C.QFileDialog_SetOption(ptr.Pointer(), C.int(option), C.int(qt.GoBoolToInt(on)))
	}
}

func (ptr *QFileDialog) SetProxyModel(proxyModel core.QAbstractProxyModel_ITF) {
	defer qt.Recovering("QFileDialog::setProxyModel")

	if ptr.Pointer() != nil {
		C.QFileDialog_SetProxyModel(ptr.Pointer(), core.PointerFromQAbstractProxyModel(proxyModel))
	}
}

func (ptr *QFileDialog) ConnectSetVisible(f func(visible bool)) {
	defer qt.Recovering("connect QFileDialog::setVisible")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setVisible", f)
	}
}

func (ptr *QFileDialog) DisconnectSetVisible() {
	defer qt.Recovering("disconnect QFileDialog::setVisible")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setVisible")
	}
}

//export callbackQFileDialogSetVisible
func callbackQFileDialogSetVisible(ptrName *C.char, visible C.int) bool {
	defer qt.Recovering("callback QFileDialog::setVisible")

	var signal = qt.GetSignal(C.GoString(ptrName), "setVisible")
	if signal != nil {
		defer signal.(func(bool))(int(visible) != 0)
		return true
	}
	return false

}

func (ptr *QFileDialog) TestOption(option QFileDialog__Option) bool {
	defer qt.Recovering("QFileDialog::testOption")

	if ptr.Pointer() != nil {
		return C.QFileDialog_TestOption(ptr.Pointer(), C.int(option)) != 0
	}
	return false
}

func (ptr *QFileDialog) ConnectUrlSelected(f func(url *core.QUrl)) {
	defer qt.Recovering("connect QFileDialog::urlSelected")

	if ptr.Pointer() != nil {
		C.QFileDialog_ConnectUrlSelected(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "urlSelected", f)
	}
}

func (ptr *QFileDialog) DisconnectUrlSelected() {
	defer qt.Recovering("disconnect QFileDialog::urlSelected")

	if ptr.Pointer() != nil {
		C.QFileDialog_DisconnectUrlSelected(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "urlSelected")
	}
}

//export callbackQFileDialogUrlSelected
func callbackQFileDialogUrlSelected(ptrName *C.char, url unsafe.Pointer) {
	defer qt.Recovering("callback QFileDialog::urlSelected")

	var signal = qt.GetSignal(C.GoString(ptrName), "urlSelected")
	if signal != nil {
		signal.(func(*core.QUrl))(core.NewQUrlFromPointer(url))
	}

}

func (ptr *QFileDialog) DestroyQFileDialog() {
	defer qt.Recovering("QFileDialog::~QFileDialog")

	if ptr.Pointer() != nil {
		C.QFileDialog_DestroyQFileDialog(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
