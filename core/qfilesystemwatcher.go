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

type QFileSystemWatcher_ITF interface {
	QObject_ITF
	QFileSystemWatcher_PTR() *QFileSystemWatcher
}

func PointerFromQFileSystemWatcher(ptr QFileSystemWatcher_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QFileSystemWatcher_PTR().Pointer()
	}
	return nil
}

func NewQFileSystemWatcherFromPointer(ptr unsafe.Pointer) *QFileSystemWatcher {
	var n = new(QFileSystemWatcher)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QFileSystemWatcher_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QFileSystemWatcher) QFileSystemWatcher_PTR() *QFileSystemWatcher {
	return ptr
}

func (ptr *QFileSystemWatcher) Directories() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QFileSystemWatcher_Directories(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QFileSystemWatcher) Files() []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QFileSystemWatcher_Files(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func NewQFileSystemWatcher(parent QObject_ITF) *QFileSystemWatcher {
	return NewQFileSystemWatcherFromPointer(C.QFileSystemWatcher_NewQFileSystemWatcher(PointerFromQObject(parent)))
}

func NewQFileSystemWatcher2(paths []string, parent QObject_ITF) *QFileSystemWatcher {
	return NewQFileSystemWatcherFromPointer(C.QFileSystemWatcher_NewQFileSystemWatcher2(C.CString(strings.Join(paths, "|")), PointerFromQObject(parent)))
}

func (ptr *QFileSystemWatcher) AddPath(path string) bool {
	if ptr.Pointer() != nil {
		return C.QFileSystemWatcher_AddPath(ptr.Pointer(), C.CString(path)) != 0
	}
	return false
}

func (ptr *QFileSystemWatcher) AddPaths(paths []string) []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QFileSystemWatcher_AddPaths(ptr.Pointer(), C.CString(strings.Join(paths, "|")))), "|")
	}
	return make([]string, 0)
}

func (ptr *QFileSystemWatcher) ConnectDirectoryChanged(f func(path string)) {
	if ptr.Pointer() != nil {
		C.QFileSystemWatcher_ConnectDirectoryChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "directoryChanged", f)
	}
}

func (ptr *QFileSystemWatcher) DisconnectDirectoryChanged() {
	if ptr.Pointer() != nil {
		C.QFileSystemWatcher_DisconnectDirectoryChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "directoryChanged")
	}
}

//export callbackQFileSystemWatcherDirectoryChanged
func callbackQFileSystemWatcherDirectoryChanged(ptrName *C.char, path *C.char) {
	qt.GetSignal(C.GoString(ptrName), "directoryChanged").(func(string))(C.GoString(path))
}

func (ptr *QFileSystemWatcher) ConnectFileChanged(f func(path string)) {
	if ptr.Pointer() != nil {
		C.QFileSystemWatcher_ConnectFileChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "fileChanged", f)
	}
}

func (ptr *QFileSystemWatcher) DisconnectFileChanged() {
	if ptr.Pointer() != nil {
		C.QFileSystemWatcher_DisconnectFileChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "fileChanged")
	}
}

//export callbackQFileSystemWatcherFileChanged
func callbackQFileSystemWatcherFileChanged(ptrName *C.char, path *C.char) {
	qt.GetSignal(C.GoString(ptrName), "fileChanged").(func(string))(C.GoString(path))
}

func (ptr *QFileSystemWatcher) RemovePath(path string) bool {
	if ptr.Pointer() != nil {
		return C.QFileSystemWatcher_RemovePath(ptr.Pointer(), C.CString(path)) != 0
	}
	return false
}

func (ptr *QFileSystemWatcher) RemovePaths(paths []string) []string {
	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QFileSystemWatcher_RemovePaths(ptr.Pointer(), C.CString(strings.Join(paths, "|")))), "|")
	}
	return make([]string, 0)
}

func (ptr *QFileSystemWatcher) DestroyQFileSystemWatcher() {
	if ptr.Pointer() != nil {
		C.QFileSystemWatcher_DestroyQFileSystemWatcher(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
