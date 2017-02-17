// +build !minimal

package opengl

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "opengl.h"
import "C"
import (
	"unsafe"
)

func cGoUnpackString(s C.struct_QtOpenGL_PackedString) string {
	if len := int(s.len); len == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}

type QGL struct {
	ptr unsafe.Pointer
}

type QGL_ITF interface {
	QGL_PTR() *QGL
}

func (ptr *QGL) QGL_PTR() *QGL {
	return ptr
}

func (ptr *QGL) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QGL) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQGL(ptr QGL_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGL_PTR().Pointer()
	}
	return nil
}

func NewQGLFromPointer(ptr unsafe.Pointer) *QGL {
	var n = new(QGL)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGL) DestroyQGL() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//go:generate stringer -type=QGL__FormatOption
//QGL::FormatOption
type QGL__FormatOption int64

var (
	QGL__DoubleBuffer          QGL__FormatOption = QGL__FormatOption(0x0001)
	QGL__DepthBuffer           QGL__FormatOption = QGL__FormatOption(0x0002)
	QGL__Rgba                  QGL__FormatOption = QGL__FormatOption(0x0004)
	QGL__AlphaChannel          QGL__FormatOption = QGL__FormatOption(0x0008)
	QGL__AccumBuffer           QGL__FormatOption = QGL__FormatOption(0x0010)
	QGL__StencilBuffer         QGL__FormatOption = QGL__FormatOption(0x0020)
	QGL__StereoBuffers         QGL__FormatOption = QGL__FormatOption(0x0040)
	QGL__DirectRendering       QGL__FormatOption = QGL__FormatOption(0x0080)
	QGL__HasOverlay            QGL__FormatOption = QGL__FormatOption(0x0100)
	QGL__SampleBuffers         QGL__FormatOption = QGL__FormatOption(0x0200)
	QGL__DeprecatedFunctions   QGL__FormatOption = QGL__FormatOption(0x0400)
	QGL__SingleBuffer          QGL__FormatOption = QGL__FormatOption(C.QGL_SingleBuffer_Type())
	QGL__NoDepthBuffer         QGL__FormatOption = QGL__FormatOption(C.QGL_NoDepthBuffer_Type())
	QGL__ColorIndex            QGL__FormatOption = QGL__FormatOption(C.QGL_ColorIndex_Type())
	QGL__NoAlphaChannel        QGL__FormatOption = QGL__FormatOption(C.QGL_NoAlphaChannel_Type())
	QGL__NoAccumBuffer         QGL__FormatOption = QGL__FormatOption(C.QGL_NoAccumBuffer_Type())
	QGL__NoStencilBuffer       QGL__FormatOption = QGL__FormatOption(C.QGL_NoStencilBuffer_Type())
	QGL__NoStereoBuffers       QGL__FormatOption = QGL__FormatOption(C.QGL_NoStereoBuffers_Type())
	QGL__IndirectRendering     QGL__FormatOption = QGL__FormatOption(C.QGL_IndirectRendering_Type())
	QGL__NoOverlay             QGL__FormatOption = QGL__FormatOption(C.QGL_NoOverlay_Type())
	QGL__NoSampleBuffers       QGL__FormatOption = QGL__FormatOption(C.QGL_NoSampleBuffers_Type())
	QGL__NoDeprecatedFunctions QGL__FormatOption = QGL__FormatOption(C.QGL_NoDeprecatedFunctions_Type())
)
