// +build !minimal

package quickcontrols2

import (
	"github.com/bluszcz/cutego"
	"github.com/bluszcz/cutego/core"
	"strings"
	"unsafe"
)

type QQuickStyle struct {
	internal.Internal
}

type QQuickStyle_ITF interface {
	QQuickStyle_PTR() *QQuickStyle
}

func (ptr *QQuickStyle) QQuickStyle_PTR() *QQuickStyle {
	return ptr
}

func (ptr *QQuickStyle) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QQuickStyle) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQQuickStyle(ptr QQuickStyle_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQuickStyle_PTR().Pointer()
	}
	return nil
}

func (n *QQuickStyle) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQQuickStyleFromPointer(ptr unsafe.Pointer) (n *QQuickStyle) {
	n = new(QQuickStyle)
	n.InitFromInternal(uintptr(ptr), "quickcontrols2.QQuickStyle")
	return
}

func (ptr *QQuickStyle) DestroyQQuickStyle() {
}

func QQuickStyle_AddStylePath(path string) {

	internal.CallLocalFunction([]interface{}{"", "", "quickcontrols2.QQuickStyle_AddStylePath", "", path})
}

func (ptr *QQuickStyle) AddStylePath(path string) {

	internal.CallLocalFunction([]interface{}{"", "", "quickcontrols2.QQuickStyle_AddStylePath", "", path})
}

func QQuickStyle_AvailableStyles() []string {

	return internal.CallLocalFunction([]interface{}{"", "", "quickcontrols2.QQuickStyle_AvailableStyles", ""}).([]string)
}

func (ptr *QQuickStyle) AvailableStyles() []string {

	return internal.CallLocalFunction([]interface{}{"", "", "quickcontrols2.QQuickStyle_AvailableStyles", ""}).([]string)
}

func QQuickStyle_Name() string {

	return internal.CallLocalFunction([]interface{}{"", "", "quickcontrols2.QQuickStyle_Name", ""}).(string)
}

func (ptr *QQuickStyle) Name() string {

	return internal.CallLocalFunction([]interface{}{"", "", "quickcontrols2.QQuickStyle_Name", ""}).(string)
}

func QQuickStyle_Path() string {

	return internal.CallLocalFunction([]interface{}{"", "", "quickcontrols2.QQuickStyle_Path", ""}).(string)
}

func (ptr *QQuickStyle) Path() string {

	return internal.CallLocalFunction([]interface{}{"", "", "quickcontrols2.QQuickStyle_Path", ""}).(string)
}

func QQuickStyle_SetFallbackStyle(style string) {

	internal.CallLocalFunction([]interface{}{"", "", "quickcontrols2.QQuickStyle_SetFallbackStyle", "", style})
}

func (ptr *QQuickStyle) SetFallbackStyle(style string) {

	internal.CallLocalFunction([]interface{}{"", "", "quickcontrols2.QQuickStyle_SetFallbackStyle", "", style})
}

func QQuickStyle_SetStyle(style string) {

	internal.CallLocalFunction([]interface{}{"", "", "quickcontrols2.QQuickStyle_SetStyle", "", style})
}

func (ptr *QQuickStyle) SetStyle(style string) {

	internal.CallLocalFunction([]interface{}{"", "", "quickcontrols2.QQuickStyle_SetStyle", "", style})
}

func QQuickStyle_StylePathList() []string {

	return internal.CallLocalFunction([]interface{}{"", "", "quickcontrols2.QQuickStyle_StylePathList", ""}).([]string)
}

func (ptr *QQuickStyle) StylePathList() []string {

	return internal.CallLocalFunction([]interface{}{"", "", "quickcontrols2.QQuickStyle_StylePathList", ""}).([]string)
}

func init() {
	internal.ConstructorTable["quickcontrols2.QQuickStyle"] = NewQQuickStyleFromPointer
}
