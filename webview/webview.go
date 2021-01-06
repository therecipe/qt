// +build !minimal

package webview

import (
	"github.com/dev-drprasad/qt/internal"
	"unsafe"
)

type QtWebView struct {
	internal.Internal
}

type QtWebView_ITF interface {
	QtWebView_PTR() *QtWebView
}

func (ptr *QtWebView) QtWebView_PTR() *QtWebView {
	return ptr
}

func (ptr *QtWebView) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QtWebView) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQtWebView(ptr QtWebView_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QtWebView_PTR().Pointer()
	}
	return nil
}

func (n *QtWebView) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQtWebViewFromPointer(ptr unsafe.Pointer) (n *QtWebView) {
	n = new(QtWebView)
	n.InitFromInternal(uintptr(ptr), "webview.QtWebView")
	return
}

func (ptr *QtWebView) DestroyQtWebView() {
}

func QtWebView_Initialize() {

	internal.CallLocalFunction([]interface{}{"", "", "webview.QtWebView_Initialize", ""})
}

func (ptr *QtWebView) Initialize() {

	internal.CallLocalFunction([]interface{}{"", "", "webview.QtWebView_Initialize", ""})
}

func init() {
	internal.ConstructorTable["webview.QtWebView"] = NewQtWebViewFromPointer
}
