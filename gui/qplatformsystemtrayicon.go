package gui

//#include "qplatformsystemtrayicon.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QPlatformSystemTrayIcon struct {
	core.QObject
}

type QPlatformSystemTrayIconITF interface {
	core.QObjectITF
	QPlatformSystemTrayIconPTR() *QPlatformSystemTrayIcon
}

func PointerFromQPlatformSystemTrayIcon(ptr QPlatformSystemTrayIconITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlatformSystemTrayIconPTR().Pointer()
	}
	return nil
}

func QPlatformSystemTrayIconFromPointer(ptr unsafe.Pointer) *QPlatformSystemTrayIcon {
	var n = new(QPlatformSystemTrayIcon)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QPlatformSystemTrayIcon_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QPlatformSystemTrayIcon) QPlatformSystemTrayIconPTR() *QPlatformSystemTrayIcon {
	return ptr
}

//QPlatformSystemTrayIcon::ActivationReason
type QPlatformSystemTrayIcon__ActivationReason int

var (
	QPlatformSystemTrayIcon__Unknown     = QPlatformSystemTrayIcon__ActivationReason(0)
	QPlatformSystemTrayIcon__Context     = QPlatformSystemTrayIcon__ActivationReason(1)
	QPlatformSystemTrayIcon__DoubleClick = QPlatformSystemTrayIcon__ActivationReason(2)
	QPlatformSystemTrayIcon__Trigger     = QPlatformSystemTrayIcon__ActivationReason(3)
	QPlatformSystemTrayIcon__MiddleClick = QPlatformSystemTrayIcon__ActivationReason(4)
)

//QPlatformSystemTrayIcon::MessageIcon
type QPlatformSystemTrayIcon__MessageIcon int

var (
	QPlatformSystemTrayIcon__NoIcon      = QPlatformSystemTrayIcon__MessageIcon(0)
	QPlatformSystemTrayIcon__Information = QPlatformSystemTrayIcon__MessageIcon(1)
	QPlatformSystemTrayIcon__Warning     = QPlatformSystemTrayIcon__MessageIcon(2)
	QPlatformSystemTrayIcon__Critical    = QPlatformSystemTrayIcon__MessageIcon(3)
)
