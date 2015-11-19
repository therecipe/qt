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

type QPlatformSystemTrayIcon_ITF interface {
	core.QObject_ITF
	QPlatformSystemTrayIcon_PTR() *QPlatformSystemTrayIcon
}

func PointerFromQPlatformSystemTrayIcon(ptr QPlatformSystemTrayIcon_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlatformSystemTrayIcon_PTR().Pointer()
	}
	return nil
}

func NewQPlatformSystemTrayIconFromPointer(ptr unsafe.Pointer) *QPlatformSystemTrayIcon {
	var n = new(QPlatformSystemTrayIcon)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QPlatformSystemTrayIcon_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QPlatformSystemTrayIcon) QPlatformSystemTrayIcon_PTR() *QPlatformSystemTrayIcon {
	return ptr
}

//QPlatformSystemTrayIcon::ActivationReason
type QPlatformSystemTrayIcon__ActivationReason int64

const (
	QPlatformSystemTrayIcon__Unknown     = QPlatformSystemTrayIcon__ActivationReason(0)
	QPlatformSystemTrayIcon__Context     = QPlatformSystemTrayIcon__ActivationReason(1)
	QPlatformSystemTrayIcon__DoubleClick = QPlatformSystemTrayIcon__ActivationReason(2)
	QPlatformSystemTrayIcon__Trigger     = QPlatformSystemTrayIcon__ActivationReason(3)
	QPlatformSystemTrayIcon__MiddleClick = QPlatformSystemTrayIcon__ActivationReason(4)
)

//QPlatformSystemTrayIcon::MessageIcon
type QPlatformSystemTrayIcon__MessageIcon int64

const (
	QPlatformSystemTrayIcon__NoIcon      = QPlatformSystemTrayIcon__MessageIcon(0)
	QPlatformSystemTrayIcon__Information = QPlatformSystemTrayIcon__MessageIcon(1)
	QPlatformSystemTrayIcon__Warning     = QPlatformSystemTrayIcon__MessageIcon(2)
	QPlatformSystemTrayIcon__Critical    = QPlatformSystemTrayIcon__MessageIcon(3)
)
