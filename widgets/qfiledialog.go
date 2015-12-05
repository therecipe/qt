package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
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
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDialog::QFileDialog")
		}
	}()

	return NewQFileDialogFromPointer(C.QFileDialog_NewQFileDialog(PointerFromQWidget(parent), C.int(flags)))
}

func (ptr *QFileDialog) AcceptMode() QFileDialog__AcceptMode {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDialog::acceptMode")
		}
	}()

	if ptr.Pointer() != nil {
		return QFileDialog__AcceptMode(C.QFileDialog_AcceptMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFileDialog) ConfirmOverwrite() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDialog::confirmOverwrite")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QFileDialog_ConfirmOverwrite(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QFileDialog) DefaultSuffix() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDialog::defaultSuffix")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QFileDialog_DefaultSuffix(ptr.Pointer()))
	}
	return ""
}

func (ptr *QFileDialog) FileMode() QFileDialog__FileMode {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDialog::fileMode")
		}
	}()

	if ptr.Pointer() != nil {
		return QFileDialog__FileMode(C.QFileDialog_FileMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFileDialog) IsNameFilterDetailsVisible() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDialog::isNameFilterDetailsVisible")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QFileDialog_IsNameFilterDetailsVisible(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QFileDialog) IsReadOnly() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDialog::isReadOnly")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QFileDialog_IsReadOnly(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QFileDialog) Options() QFileDialog__Option {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDialog::options")
		}
	}()

	if ptr.Pointer() != nil {
		return QFileDialog__Option(C.QFileDialog_Options(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFileDialog) ResolveSymlinks() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDialog::resolveSymlinks")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QFileDialog_ResolveSymlinks(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QFileDialog) SetAcceptMode(mode QFileDialog__AcceptMode) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDialog::setAcceptMode")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFileDialog_SetAcceptMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QFileDialog) SetConfirmOverwrite(enabled bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDialog::setConfirmOverwrite")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFileDialog_SetConfirmOverwrite(ptr.Pointer(), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QFileDialog) SetDefaultSuffix(suffix string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDialog::setDefaultSuffix")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFileDialog_SetDefaultSuffix(ptr.Pointer(), C.CString(suffix))
	}
}

func (ptr *QFileDialog) SetFileMode(mode QFileDialog__FileMode) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDialog::setFileMode")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFileDialog_SetFileMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QFileDialog) SetNameFilterDetailsVisible(enabled bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDialog::setNameFilterDetailsVisible")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFileDialog_SetNameFilterDetailsVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QFileDialog) SetOptions(options QFileDialog__Option) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDialog::setOptions")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFileDialog_SetOptions(ptr.Pointer(), C.int(options))
	}
}

func (ptr *QFileDialog) SetReadOnly(enabled bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDialog::setReadOnly")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFileDialog_SetReadOnly(ptr.Pointer(), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QFileDialog) SetResolveSymlinks(enabled bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDialog::setResolveSymlinks")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFileDialog_SetResolveSymlinks(ptr.Pointer(), C.int(qt.GoBoolToInt(enabled)))
	}
}

func (ptr *QFileDialog) SetViewMode(mode QFileDialog__ViewMode) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDialog::setViewMode")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFileDialog_SetViewMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QFileDialog) ViewMode() QFileDialog__ViewMode {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDialog::viewMode")
		}
	}()

	if ptr.Pointer() != nil {
		return QFileDialog__ViewMode(C.QFileDialog_ViewMode(ptr.Pointer()))
	}
	return 0
}

func NewQFileDialog2(parent QWidget_ITF, caption string, directory string, filter string) *QFileDialog {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDialog::QFileDialog")
		}
	}()

	return NewQFileDialogFromPointer(C.QFileDialog_NewQFileDialog2(PointerFromQWidget(parent), C.CString(caption), C.CString(directory), C.CString(filter)))
}

func (ptr *QFileDialog) ConnectCurrentChanged(f func(path string)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDialog::currentChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFileDialog_ConnectCurrentChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "currentChanged", f)
	}
}

func (ptr *QFileDialog) DisconnectCurrentChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDialog::currentChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFileDialog_DisconnectCurrentChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "currentChanged")
	}
}

//export callbackQFileDialogCurrentChanged
func callbackQFileDialogCurrentChanged(ptrName *C.char, path *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDialog::currentChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "currentChanged").(func(string))(C.GoString(path))
}

func (ptr *QFileDialog) Directory() *core.QDir {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDialog::directory")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQDirFromPointer(C.QFileDialog_Directory(ptr.Pointer()))
	}
	return nil
}

func (ptr *QFileDialog) ConnectDirectoryEntered(f func(directory string)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDialog::directoryEntered")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFileDialog_ConnectDirectoryEntered(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "directoryEntered", f)
	}
}

func (ptr *QFileDialog) DisconnectDirectoryEntered() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDialog::directoryEntered")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFileDialog_DisconnectDirectoryEntered(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "directoryEntered")
	}
}

//export callbackQFileDialogDirectoryEntered
func callbackQFileDialogDirectoryEntered(ptrName *C.char, directory *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDialog::directoryEntered")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "directoryEntered").(func(string))(C.GoString(directory))
}

func (ptr *QFileDialog) ConnectFileSelected(f func(file string)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDialog::fileSelected")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFileDialog_ConnectFileSelected(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "fileSelected", f)
	}
}

func (ptr *QFileDialog) DisconnectFileSelected() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDialog::fileSelected")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFileDialog_DisconnectFileSelected(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "fileSelected")
	}
}

//export callbackQFileDialogFileSelected
func callbackQFileDialogFileSelected(ptrName *C.char, file *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDialog::fileSelected")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "fileSelected").(func(string))(C.GoString(file))
}

func (ptr *QFileDialog) ConnectFilesSelected(f func(selected []string)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDialog::filesSelected")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFileDialog_ConnectFilesSelected(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "filesSelected", f)
	}
}

func (ptr *QFileDialog) DisconnectFilesSelected() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDialog::filesSelected")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFileDialog_DisconnectFilesSelected(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "filesSelected")
	}
}

//export callbackQFileDialogFilesSelected
func callbackQFileDialogFilesSelected(ptrName *C.char, selected *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDialog::filesSelected")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "filesSelected").(func([]string))(strings.Split(C.GoString(selected), ",,,"))
}

func (ptr *QFileDialog) Filter() core.QDir__Filter {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDialog::filter")
		}
	}()

	if ptr.Pointer() != nil {
		return core.QDir__Filter(C.QFileDialog_Filter(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFileDialog) ConnectFilterSelected(f func(filter string)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDialog::filterSelected")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFileDialog_ConnectFilterSelected(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "filterSelected", f)
	}
}

func (ptr *QFileDialog) DisconnectFilterSelected() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDialog::filterSelected")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFileDialog_DisconnectFilterSelected(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "filterSelected")
	}
}

//export callbackQFileDialogFilterSelected
func callbackQFileDialogFilterSelected(ptrName *C.char, filter *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDialog::filterSelected")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "filterSelected").(func(string))(C.GoString(filter))
}

func QFileDialog_GetExistingDirectory(parent QWidget_ITF, caption string, dir string, options QFileDialog__Option) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDialog::getExistingDirectory")
		}
	}()

	return C.GoString(C.QFileDialog_QFileDialog_GetExistingDirectory(PointerFromQWidget(parent), C.CString(caption), C.CString(dir), C.int(options)))
}

func QFileDialog_GetOpenFileName(parent QWidget_ITF, caption string, dir string, filter string, selectedFilter string, options QFileDialog__Option) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDialog::getOpenFileName")
		}
	}()

	return C.GoString(C.QFileDialog_QFileDialog_GetOpenFileName(PointerFromQWidget(parent), C.CString(caption), C.CString(dir), C.CString(filter), C.CString(selectedFilter), C.int(options)))
}

func QFileDialog_GetOpenFileNames(parent QWidget_ITF, caption string, dir string, filter string, selectedFilter string, options QFileDialog__Option) []string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDialog::getOpenFileNames")
		}
	}()

	return strings.Split(C.GoString(C.QFileDialog_QFileDialog_GetOpenFileNames(PointerFromQWidget(parent), C.CString(caption), C.CString(dir), C.CString(filter), C.CString(selectedFilter), C.int(options))), ",,,")
}

func QFileDialog_GetSaveFileName(parent QWidget_ITF, caption string, dir string, filter string, selectedFilter string, options QFileDialog__Option) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDialog::getSaveFileName")
		}
	}()

	return C.GoString(C.QFileDialog_QFileDialog_GetSaveFileName(PointerFromQWidget(parent), C.CString(caption), C.CString(dir), C.CString(filter), C.CString(selectedFilter), C.int(options)))
}

func (ptr *QFileDialog) History() []string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDialog::history")
		}
	}()

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QFileDialog_History(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QFileDialog) IconProvider() *QFileIconProvider {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDialog::iconProvider")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQFileIconProviderFromPointer(C.QFileDialog_IconProvider(ptr.Pointer()))
	}
	return nil
}

func (ptr *QFileDialog) ItemDelegate() *QAbstractItemDelegate {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDialog::itemDelegate")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQAbstractItemDelegateFromPointer(C.QFileDialog_ItemDelegate(ptr.Pointer()))
	}
	return nil
}

func (ptr *QFileDialog) LabelText(label QFileDialog__DialogLabel) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDialog::labelText")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QFileDialog_LabelText(ptr.Pointer(), C.int(label)))
	}
	return ""
}

func (ptr *QFileDialog) MimeTypeFilters() []string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDialog::mimeTypeFilters")
		}
	}()

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QFileDialog_MimeTypeFilters(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QFileDialog) NameFilters() []string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDialog::nameFilters")
		}
	}()

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QFileDialog_NameFilters(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QFileDialog) Open(receiver core.QObject_ITF, member string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDialog::open")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFileDialog_Open(ptr.Pointer(), core.PointerFromQObject(receiver), C.CString(member))
	}
}

func (ptr *QFileDialog) ProxyModel() *core.QAbstractProxyModel {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDialog::proxyModel")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQAbstractProxyModelFromPointer(C.QFileDialog_ProxyModel(ptr.Pointer()))
	}
	return nil
}

func (ptr *QFileDialog) RestoreState(state core.QByteArray_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDialog::restoreState")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QFileDialog_RestoreState(ptr.Pointer(), core.PointerFromQByteArray(state)) != 0
	}
	return false
}

func (ptr *QFileDialog) SaveState() *core.QByteArray {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDialog::saveState")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QFileDialog_SaveState(ptr.Pointer()))
	}
	return nil
}

func (ptr *QFileDialog) SelectFile(filename string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDialog::selectFile")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFileDialog_SelectFile(ptr.Pointer(), C.CString(filename))
	}
}

func (ptr *QFileDialog) SelectMimeTypeFilter(filter string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDialog::selectMimeTypeFilter")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFileDialog_SelectMimeTypeFilter(ptr.Pointer(), C.CString(filter))
	}
}

func (ptr *QFileDialog) SelectNameFilter(filter string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDialog::selectNameFilter")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFileDialog_SelectNameFilter(ptr.Pointer(), C.CString(filter))
	}
}

func (ptr *QFileDialog) SelectUrl(url core.QUrl_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDialog::selectUrl")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFileDialog_SelectUrl(ptr.Pointer(), core.PointerFromQUrl(url))
	}
}

func (ptr *QFileDialog) SelectedFiles() []string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDialog::selectedFiles")
		}
	}()

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QFileDialog_SelectedFiles(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QFileDialog) SelectedNameFilter() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDialog::selectedNameFilter")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QFileDialog_SelectedNameFilter(ptr.Pointer()))
	}
	return ""
}

func (ptr *QFileDialog) SetDirectory2(directory core.QDir_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDialog::setDirectory")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFileDialog_SetDirectory2(ptr.Pointer(), core.PointerFromQDir(directory))
	}
}

func (ptr *QFileDialog) SetDirectory(directory string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDialog::setDirectory")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFileDialog_SetDirectory(ptr.Pointer(), C.CString(directory))
	}
}

func (ptr *QFileDialog) SetDirectoryUrl(directory core.QUrl_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDialog::setDirectoryUrl")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFileDialog_SetDirectoryUrl(ptr.Pointer(), core.PointerFromQUrl(directory))
	}
}

func (ptr *QFileDialog) SetFilter(filters core.QDir__Filter) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDialog::setFilter")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFileDialog_SetFilter(ptr.Pointer(), C.int(filters))
	}
}

func (ptr *QFileDialog) SetHistory(paths []string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDialog::setHistory")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFileDialog_SetHistory(ptr.Pointer(), C.CString(strings.Join(paths, ",,,")))
	}
}

func (ptr *QFileDialog) SetIconProvider(provider QFileIconProvider_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDialog::setIconProvider")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFileDialog_SetIconProvider(ptr.Pointer(), PointerFromQFileIconProvider(provider))
	}
}

func (ptr *QFileDialog) SetItemDelegate(delegate QAbstractItemDelegate_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDialog::setItemDelegate")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFileDialog_SetItemDelegate(ptr.Pointer(), PointerFromQAbstractItemDelegate(delegate))
	}
}

func (ptr *QFileDialog) SetLabelText(label QFileDialog__DialogLabel, text string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDialog::setLabelText")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFileDialog_SetLabelText(ptr.Pointer(), C.int(label), C.CString(text))
	}
}

func (ptr *QFileDialog) SetMimeTypeFilters(filters []string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDialog::setMimeTypeFilters")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFileDialog_SetMimeTypeFilters(ptr.Pointer(), C.CString(strings.Join(filters, ",,,")))
	}
}

func (ptr *QFileDialog) SetNameFilter(filter string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDialog::setNameFilter")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFileDialog_SetNameFilter(ptr.Pointer(), C.CString(filter))
	}
}

func (ptr *QFileDialog) SetNameFilters(filters []string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDialog::setNameFilters")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFileDialog_SetNameFilters(ptr.Pointer(), C.CString(strings.Join(filters, ",,,")))
	}
}

func (ptr *QFileDialog) SetOption(option QFileDialog__Option, on bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDialog::setOption")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFileDialog_SetOption(ptr.Pointer(), C.int(option), C.int(qt.GoBoolToInt(on)))
	}
}

func (ptr *QFileDialog) SetProxyModel(proxyModel core.QAbstractProxyModel_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDialog::setProxyModel")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFileDialog_SetProxyModel(ptr.Pointer(), core.PointerFromQAbstractProxyModel(proxyModel))
	}
}

func (ptr *QFileDialog) SetVisible(visible bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDialog::setVisible")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFileDialog_SetVisible(ptr.Pointer(), C.int(qt.GoBoolToInt(visible)))
	}
}

func (ptr *QFileDialog) TestOption(option QFileDialog__Option) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDialog::testOption")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QFileDialog_TestOption(ptr.Pointer(), C.int(option)) != 0
	}
	return false
}

func (ptr *QFileDialog) DestroyQFileDialog() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QFileDialog::~QFileDialog")
		}
	}()

	if ptr.Pointer() != nil {
		C.QFileDialog_DestroyQFileDialog(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
