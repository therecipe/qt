package qt

//#include "qfiledialog.h"
import "C"

import "strings"

type qfiledialog struct {
	qdialog
}

type QFileDialog interface {
	QDialog
	AcceptMode() AcceptMode
	DefaultSuffix() string
	FileMode() FileMode
	History() []string
	LabelText(label DialogLabel) string
	MimeTypeFilters() []string
	NameFilters() []string
	Open(receiver QObject, member string)
	Options() Option
	SelectFile(filename string)
	SelectMimeTypeFilter(filter string)
	SelectNameFilter(filter string)
	SelectedFiles() []string
	SelectedNameFilter() string
	SetAcceptMode(mode AcceptMode)
	SetDefaultSuffix(suffix string)
	SetDirectory(directory string)
	SetFileMode(mode FileMode)
	SetHistory(paths []string)
	SetLabelText(label DialogLabel, text string)
	SetMimeTypeFilters(filters []string)
	SetNameFilter(filter string)
	SetNameFilters(filters []string)
	SetOption(option Option, on bool)
	SetOptions(options Option)
	SetViewMode(mode ViewMode)
	TestOption(option Option) bool
	ViewMode() ViewMode
	ConnectSignalCurrentChanged(f func())
	DisconnectSignalCurrentChanged()
	SignalCurrentChanged() func()
	ConnectSignalDirectoryEntered(f func())
	DisconnectSignalDirectoryEntered()
	SignalDirectoryEntered() func()
	ConnectSignalFileSelected(f func())
	DisconnectSignalFileSelected()
	SignalFileSelected() func()
	ConnectSignalFilesSelected(f func())
	DisconnectSignalFilesSelected()
	SignalFilesSelected() func()
	ConnectSignalFilterSelected(f func())
	DisconnectSignalFilterSelected()
	SignalFilterSelected() func()
}

func (p *qfiledialog) Pointer() (ptr C.QtObjectPtr) {
	return p.ptr
}

func (p *qfiledialog) SetPointer(ptr C.QtObjectPtr) {
	p.ptr = ptr
}

//AcceptMode
type AcceptMode int

var (
	ACCEPTOPEN = AcceptMode(C.QFileDialog_AcceptOpen())
	ACCEPTSAVE = AcceptMode(C.QFileDialog_AcceptSave())
)

//DialogLabel
type DialogLabel int

var (
	LOOKIN   = DialogLabel(C.QFileDialog_LookIn())
	FILENAME = DialogLabel(C.QFileDialog_FileName())
	FILETYPE = DialogLabel(C.QFileDialog_FileType())
	ACCEPT   = DialogLabel(C.QFileDialog_Accept())
	REJECT   = DialogLabel(C.QFileDialog_Reject())
)

//FileMode
type FileMode int

var (
	ANYFILE       = FileMode(C.QFileDialog_AnyFile())
	EXISTINGFILE  = FileMode(C.QFileDialog_ExistingFile())
	DIRECTORY     = FileMode(C.QFileDialog_Directory())
	EXISTINGFILES = FileMode(C.QFileDialog_ExistingFiles())
	DIRECTORYONLY = FileMode(C.QFileDialog_DirectoryOnly())
)

//Option
type Option int

var (
	SHOWDIRSONLY                = Option(C.QFileDialog_ShowDirsOnly())
	DONTRESOLVESYMLINKS         = Option(C.QFileDialog_DontResolveSymlinks())
	DONTCONFIRMOVERWRITE        = Option(C.QFileDialog_DontConfirmOverwrite())
	DONTUSENATIVEDIALOG         = Option(C.QFileDialog_DontUseNativeDialog())
	READONLY                    = Option(C.QFileDialog_ReadOnly())
	HIDENAMEFILTERDETAILS       = Option(C.QFileDialog_HideNameFilterDetails())
	DONTUSESHEET                = Option(C.QFileDialog_DontUseSheet())
	DONTUSECUSTOMDIRECTORYICONS = Option(C.QFileDialog_DontUseCustomDirectoryIcons())
)

//ViewMode
type ViewMode int

var (
	DETAIL = ViewMode(C.QFileDialog_Detail())
	LIST   = ViewMode(C.QFileDialog_List())
)

func NewQFileDialog(parent QWidget, caption string, directory string, filter string) QFileDialog {
	var parentPtr C.QtObjectPtr
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qfiledialog = new(qfiledialog)
	qfiledialog.SetPointer(C.QFileDialog_New_QWidget_String_String_String(parentPtr, C.CString(caption), C.CString(directory), C.CString(filter)))
	qfiledialog.SetObjectName("QFileDialog_" + randomIdentifier())
	return qfiledialog
}

func (p *qfiledialog) Destroy() {
	if p.Pointer() != nil {
		getSignal(p.ObjectName(), "destroyed")()
		C.QFileDialog_Destroy(p.Pointer())
		p.SetPointer(nil)
	}
}

func (p *qfiledialog) AcceptMode() AcceptMode {
	if p.Pointer() == nil {
		return 0
	}
	return AcceptMode(C.QFileDialog_AcceptMode(p.Pointer()))
}

func (p *qfiledialog) DefaultSuffix() string {
	if p.Pointer() == nil {
		return ""
	}
	return C.GoString(C.QFileDialog_DefaultSuffix(p.Pointer()))
}

func (p *qfiledialog) FileMode() FileMode {
	if p.Pointer() == nil {
		return 0
	}
	return FileMode(C.QFileDialog_FileMode(p.Pointer()))
}

func (p *qfiledialog) History() []string {
	if p.Pointer() == nil {
		return []string{""}
	}
	return strings.Split(C.GoString(C.QFileDialog_History(p.Pointer())), "|")
}

func (p *qfiledialog) LabelText(label DialogLabel) string {
	if p.Pointer() == nil {
		return ""
	}
	return C.GoString(C.QFileDialog_LabelText_DialogLabel(p.Pointer(), C.int(label)))
}

func (p *qfiledialog) MimeTypeFilters() []string {
	if p.Pointer() == nil {
		return []string{""}
	}
	return strings.Split(C.GoString(C.QFileDialog_MimeTypeFilters(p.Pointer())), "|")
}

func (p *qfiledialog) NameFilters() []string {
	if p.Pointer() == nil {
		return []string{""}
	}
	return strings.Split(C.GoString(C.QFileDialog_NameFilters(p.Pointer())), "|")
}

func (p *qfiledialog) Open(receiver QObject, member string) {
	if p.Pointer() != nil {
		var receiverPtr C.QtObjectPtr
		if receiver != nil {
			receiverPtr = receiver.Pointer()
		}
		C.QFileDialog_Open_QObject_String(p.Pointer(), receiverPtr, C.CString(member))
	}
}

func (p *qfiledialog) Options() Option {
	if p.Pointer() == nil {
		return 0
	}
	return Option(C.QFileDialog_Options(p.Pointer()))
}

func (p *qfiledialog) SelectFile(filename string) {
	if p.Pointer() != nil {
		C.QFileDialog_SelectFile_String(p.Pointer(), C.CString(filename))
	}
}

func (p *qfiledialog) SelectMimeTypeFilter(filter string) {
	if p.Pointer() != nil {
		C.QFileDialog_SelectMimeTypeFilter_String(p.Pointer(), C.CString(filter))
	}
}

func (p *qfiledialog) SelectNameFilter(filter string) {
	if p.Pointer() != nil {
		C.QFileDialog_SelectNameFilter_String(p.Pointer(), C.CString(filter))
	}
}

func (p *qfiledialog) SelectedFiles() []string {
	if p.Pointer() == nil {
		return []string{""}
	}
	return strings.Split(C.GoString(C.QFileDialog_SelectedFiles(p.Pointer())), "|")
}

func (p *qfiledialog) SelectedNameFilter() string {
	if p.Pointer() == nil {
		return ""
	}
	return C.GoString(C.QFileDialog_SelectedNameFilter(p.Pointer()))
}

func (p *qfiledialog) SetAcceptMode(mode AcceptMode) {
	if p.Pointer() != nil {
		C.QFileDialog_SetAcceptMode_AcceptMode(p.Pointer(), C.int(mode))
	}
}

func (p *qfiledialog) SetDefaultSuffix(suffix string) {
	if p.Pointer() != nil {
		C.QFileDialog_SetDefaultSuffix_String(p.Pointer(), C.CString(suffix))
	}
}

func (p *qfiledialog) SetDirectory(directory string) {
	if p.Pointer() != nil {
		C.QFileDialog_SetDirectory_String(p.Pointer(), C.CString(directory))
	}
}

func (p *qfiledialog) SetFileMode(mode FileMode) {
	if p.Pointer() != nil {
		C.QFileDialog_SetFileMode_FileMode(p.Pointer(), C.int(mode))
	}
}

func (p *qfiledialog) SetHistory(paths []string) {
	if p.Pointer() != nil {
		C.QFileDialog_SetHistory_QStringList(p.Pointer(), C.CString(strings.Join(paths, "|")))
	}
}

func (p *qfiledialog) SetLabelText(label DialogLabel, text string) {
	if p.Pointer() != nil {
		C.QFileDialog_SetLabelText_DialogLabel_String(p.Pointer(), C.int(label), C.CString(text))
	}
}

func (p *qfiledialog) SetMimeTypeFilters(filters []string) {
	if p.Pointer() != nil {
		C.QFileDialog_SetMimeTypeFilters_QStringList(p.Pointer(), C.CString(strings.Join(filters, "|")))
	}
}

func (p *qfiledialog) SetNameFilter(filter string) {
	if p.Pointer() != nil {
		C.QFileDialog_SetNameFilter_String(p.Pointer(), C.CString(filter))
	}
}

func (p *qfiledialog) SetNameFilters(filters []string) {
	if p.Pointer() != nil {
		C.QFileDialog_SetNameFilters_QStringList(p.Pointer(), C.CString(strings.Join(filters, "|")))
	}
}

func (p *qfiledialog) SetOption(option Option, on bool) {
	if p.Pointer() != nil {
		C.QFileDialog_SetOption_Option_Bool(p.Pointer(), C.int(option), goBoolToCInt(on))
	}
}

func (p *qfiledialog) SetOptions(options Option) {
	if p.Pointer() != nil {
		C.QFileDialog_SetOptions_Option(p.Pointer(), C.int(options))
	}
}

func (p *qfiledialog) SetViewMode(mode ViewMode) {
	if p.Pointer() != nil {
		C.QFileDialog_SetViewMode_ViewMode(p.Pointer(), C.int(mode))
	}
}

func (p *qfiledialog) TestOption(option Option) bool {
	if p.Pointer() == nil {
		return false
	}
	return C.QFileDialog_TestOption_Option(p.Pointer(), C.int(option)) != 0
}

func (p *qfiledialog) ViewMode() ViewMode {
	if p.Pointer() == nil {
		return 0
	}
	return ViewMode(C.QFileDialog_ViewMode(p.Pointer()))
}

func (p *qfiledialog) ConnectSignalCurrentChanged(f func()) {
	C.QFileDialog_ConnectSignalCurrentChanged(p.Pointer())
	connectSignal(p.ObjectName(), "currentChanged", f)
}

func (p *qfiledialog) DisconnectSignalCurrentChanged() {
	C.QFileDialog_DisconnectSignalCurrentChanged(p.Pointer())
	disconnectSignal(p.ObjectName(), "currentChanged")
}

func (p *qfiledialog) SignalCurrentChanged() func() {
	return getSignal(p.ObjectName(), "currentChanged")
}

func (p *qfiledialog) ConnectSignalDirectoryEntered(f func()) {
	C.QFileDialog_ConnectSignalDirectoryEntered(p.Pointer())
	connectSignal(p.ObjectName(), "directoryEntered", f)
}

func (p *qfiledialog) DisconnectSignalDirectoryEntered() {
	C.QFileDialog_DisconnectSignalDirectoryEntered(p.Pointer())
	disconnectSignal(p.ObjectName(), "directoryEntered")
}

func (p *qfiledialog) SignalDirectoryEntered() func() {
	return getSignal(p.ObjectName(), "directoryEntered")
}

func (p *qfiledialog) ConnectSignalFileSelected(f func()) {
	C.QFileDialog_ConnectSignalFileSelected(p.Pointer())
	connectSignal(p.ObjectName(), "fileSelected", f)
}

func (p *qfiledialog) DisconnectSignalFileSelected() {
	C.QFileDialog_DisconnectSignalFileSelected(p.Pointer())
	disconnectSignal(p.ObjectName(), "fileSelected")
}

func (p *qfiledialog) SignalFileSelected() func() {
	return getSignal(p.ObjectName(), "fileSelected")
}

func (p *qfiledialog) ConnectSignalFilesSelected(f func()) {
	C.QFileDialog_ConnectSignalFilesSelected(p.Pointer())
	connectSignal(p.ObjectName(), "filesSelected", f)
}

func (p *qfiledialog) DisconnectSignalFilesSelected() {
	C.QFileDialog_DisconnectSignalFilesSelected(p.Pointer())
	disconnectSignal(p.ObjectName(), "filesSelected")
}

func (p *qfiledialog) SignalFilesSelected() func() {
	return getSignal(p.ObjectName(), "filesSelected")
}

func (p *qfiledialog) ConnectSignalFilterSelected(f func()) {
	C.QFileDialog_ConnectSignalFilterSelected(p.Pointer())
	connectSignal(p.ObjectName(), "filterSelected", f)
}

func (p *qfiledialog) DisconnectSignalFilterSelected() {
	C.QFileDialog_DisconnectSignalFilterSelected(p.Pointer())
	disconnectSignal(p.ObjectName(), "filterSelected")
}

func (p *qfiledialog) SignalFilterSelected() func() {
	return getSignal(p.ObjectName(), "filterSelected")
}

func QFileDialog_GetExistingDirectory(parent QWidget, caption string, dir string, options Option) string {
	var parentPtr C.QtObjectPtr
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	return C.GoString(C.QFileDialog_GetExistingDirectory_QWidget_String_String_Option(parentPtr, C.CString(caption), C.CString(dir), C.int(options)))
}

func QFileDialog_GetOpenFileName(parent QWidget, caption string, dir string, filter string, selectedFilter string, options Option) string {
	var parentPtr C.QtObjectPtr
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	return C.GoString(C.QFileDialog_GetOpenFileName_QWidget_String_String_String_String_Option(parentPtr, C.CString(caption), C.CString(dir), C.CString(filter), C.CString(selectedFilter), C.int(options)))
}

func QFileDialog_GetOpenFileNames(parent QWidget, caption string, dir string, filter string, selectedFilter string, options Option) []string {
	var parentPtr C.QtObjectPtr
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	return strings.Split(C.GoString(C.QFileDialog_GetOpenFileNames_QWidget_String_String_String_String_Option(parentPtr, C.CString(caption), C.CString(dir), C.CString(filter), C.CString(selectedFilter), C.int(options))), "|")
}

func QFileDialog_GetSaveFileName(parent QWidget, caption string, dir string, filter string, selectedFilter string, options Option) string {
	var parentPtr C.QtObjectPtr
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	return C.GoString(C.QFileDialog_GetSaveFileName_QWidget_String_String_String_String_Option(parentPtr, C.CString(caption), C.CString(dir), C.CString(filter), C.CString(selectedFilter), C.int(options)))
}

//export callbackQFileDialog
func callbackQFileDialog(ptr C.QtObjectPtr, msg *C.char) {
	var qfiledialog = new(qfiledialog)
	qfiledialog.SetPointer(ptr)
	getSignal(qfiledialog.ObjectName(), C.GoString(msg))()
}
