package core

//#include "qfilesystemwatcher.h"
import "C"
import (
	"github.com/therecipe/qt"
	"strings"
	"unsafe"
)

type QFileSystemWatcher struct {
	QObject
}

type QFileSystemWatcherITF interface {
	QObjectITF
	QFileSystemWatcherPTR() *QFileSystemWatcher
}

func PointerFromQFileSystemWatcher(ptr QFileSystemWatcherITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QFileSystemWatcherPTR().Pointer()
	}
	return nil
}

func QFileSystemWatcherFromPointer(ptr unsafe.Pointer) *QFileSystemWatcher {
	var n = new(QFileSystemWatcher)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QFileSystemWatcher_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QFileSystemWatcher) QFileSystemWatcherPTR() *QFileSystemWatcher {
	return ptr
}

func (ptr *QFileSystemWatcher) Directories() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QFileSystemWatcher_Directories(C.QtObjectPtr(ptr.Pointer()))), "|")
	}
	return make([]string, 0)
}

func (ptr *QFileSystemWatcher) Files() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QFileSystemWatcher_Files(C.QtObjectPtr(ptr.Pointer()))), "|")
	}
	return make([]string, 0)
}

func NewQFileSystemWatcher(parent QObjectITF) *QFileSystemWatcher {
	return QFileSystemWatcherFromPointer(unsafe.Pointer(C.QFileSystemWatcher_NewQFileSystemWatcher(C.QtObjectPtr(PointerFromQObject(parent)))))
}

func NewQFileSystemWatcher2(paths []string, parent QObjectITF) *QFileSystemWatcher {
	return QFileSystemWatcherFromPointer(unsafe.Pointer(C.QFileSystemWatcher_NewQFileSystemWatcher2(C.CString(strings.Join(paths, "|")), C.QtObjectPtr(PointerFromQObject(parent)))))
}

func (ptr *QFileSystemWatcher) AddPath(path string) bool {
	if ptr.Pointer() != nil {
		return C.QFileSystemWatcher_AddPath(C.QtObjectPtr(ptr.Pointer()), C.CString(path)) != 0
	}
	return false
}

func (ptr *QFileSystemWatcher) AddPaths(paths []string) []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QFileSystemWatcher_AddPaths(C.QtObjectPtr(ptr.Pointer()), C.CString(strings.Join(paths, "|")))), "|")
	}
	return make([]string, 0)
}

func (ptr *QFileSystemWatcher) ConnectDirectoryChanged(f func(path string)) {
	if ptr.Pointer() != nil {
		C.QFileSystemWatcher_ConnectDirectoryChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "directoryChanged", f)
	}
}

func (ptr *QFileSystemWatcher) DisconnectDirectoryChanged() {
	if ptr.Pointer() != nil {
		C.QFileSystemWatcher_DisconnectDirectoryChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "directoryChanged")
	}
}

//export callbackQFileSystemWatcherDirectoryChanged
func callbackQFileSystemWatcherDirectoryChanged(ptrName *C.char, path *C.char) {
	qt.GetSignal(C.GoString(ptrName), "directoryChanged").(func(string))(C.GoString(path))
}

func (ptr *QFileSystemWatcher) ConnectFileChanged(f func(path string)) {
	if ptr.Pointer() != nil {
		C.QFileSystemWatcher_ConnectFileChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "fileChanged", f)
	}
}

func (ptr *QFileSystemWatcher) DisconnectFileChanged() {
	if ptr.Pointer() != nil {
		C.QFileSystemWatcher_DisconnectFileChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "fileChanged")
	}
}

//export callbackQFileSystemWatcherFileChanged
func callbackQFileSystemWatcherFileChanged(ptrName *C.char, path *C.char) {
	qt.GetSignal(C.GoString(ptrName), "fileChanged").(func(string))(C.GoString(path))
}

func (ptr *QFileSystemWatcher) RemovePath(path string) bool {
	if ptr.Pointer() != nil {
		return C.QFileSystemWatcher_RemovePath(C.QtObjectPtr(ptr.Pointer()), C.CString(path)) != 0
	}
	return false
}

func (ptr *QFileSystemWatcher) RemovePaths(paths []string) []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QFileSystemWatcher_RemovePaths(C.QtObjectPtr(ptr.Pointer()), C.CString(strings.Join(paths, "|")))), "|")
	}
	return make([]string, 0)
}

func (ptr *QFileSystemWatcher) DestroyQFileSystemWatcher() {
	if ptr.Pointer() != nil {
		C.QFileSystemWatcher_DestroyQFileSystemWatcher(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
