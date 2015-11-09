package gui

//#include "qsurfaceformat.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QSurfaceFormat struct {
	ptr unsafe.Pointer
}

type QSurfaceFormat_ITF interface {
	QSurfaceFormat_PTR() *QSurfaceFormat
}

func (p *QSurfaceFormat) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSurfaceFormat) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSurfaceFormat(ptr QSurfaceFormat_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSurfaceFormat_PTR().Pointer()
	}
	return nil
}

func NewQSurfaceFormatFromPointer(ptr unsafe.Pointer) *QSurfaceFormat {
	var n = new(QSurfaceFormat)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSurfaceFormat) QSurfaceFormat_PTR() *QSurfaceFormat {
	return ptr
}

//QSurfaceFormat::FormatOption
type QSurfaceFormat__FormatOption int64

const (
	QSurfaceFormat__StereoBuffers       = QSurfaceFormat__FormatOption(0x0001)
	QSurfaceFormat__DebugContext        = QSurfaceFormat__FormatOption(0x0002)
	QSurfaceFormat__DeprecatedFunctions = QSurfaceFormat__FormatOption(0x0004)
	QSurfaceFormat__ResetNotification   = QSurfaceFormat__FormatOption(0x0008)
)

//QSurfaceFormat::OpenGLContextProfile
type QSurfaceFormat__OpenGLContextProfile int64

const (
	QSurfaceFormat__NoProfile            = QSurfaceFormat__OpenGLContextProfile(0)
	QSurfaceFormat__CoreProfile          = QSurfaceFormat__OpenGLContextProfile(1)
	QSurfaceFormat__CompatibilityProfile = QSurfaceFormat__OpenGLContextProfile(2)
)

//QSurfaceFormat::RenderableType
type QSurfaceFormat__RenderableType int64

const (
	QSurfaceFormat__DefaultRenderableType = QSurfaceFormat__RenderableType(0x0)
	QSurfaceFormat__OpenGL                = QSurfaceFormat__RenderableType(0x1)
	QSurfaceFormat__OpenGLES              = QSurfaceFormat__RenderableType(0x2)
	QSurfaceFormat__OpenVG                = QSurfaceFormat__RenderableType(0x4)
)

//QSurfaceFormat::SwapBehavior
type QSurfaceFormat__SwapBehavior int64

const (
	QSurfaceFormat__DefaultSwapBehavior = QSurfaceFormat__SwapBehavior(0)
	QSurfaceFormat__SingleBuffer        = QSurfaceFormat__SwapBehavior(1)
	QSurfaceFormat__DoubleBuffer        = QSurfaceFormat__SwapBehavior(2)
	QSurfaceFormat__TripleBuffer        = QSurfaceFormat__SwapBehavior(3)
)

func NewQSurfaceFormat() *QSurfaceFormat {
	return NewQSurfaceFormatFromPointer(C.QSurfaceFormat_NewQSurfaceFormat())
}

func NewQSurfaceFormat2(options QSurfaceFormat__FormatOption) *QSurfaceFormat {
	return NewQSurfaceFormatFromPointer(C.QSurfaceFormat_NewQSurfaceFormat2(C.int(options)))
}

func NewQSurfaceFormat3(other QSurfaceFormat_ITF) *QSurfaceFormat {
	return NewQSurfaceFormatFromPointer(C.QSurfaceFormat_NewQSurfaceFormat3(PointerFromQSurfaceFormat(other)))
}

func (ptr *QSurfaceFormat) AlphaBufferSize() int {
	if ptr.Pointer() != nil {
		return int(C.QSurfaceFormat_AlphaBufferSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSurfaceFormat) BlueBufferSize() int {
	if ptr.Pointer() != nil {
		return int(C.QSurfaceFormat_BlueBufferSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSurfaceFormat) DepthBufferSize() int {
	if ptr.Pointer() != nil {
		return int(C.QSurfaceFormat_DepthBufferSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSurfaceFormat) GreenBufferSize() int {
	if ptr.Pointer() != nil {
		return int(C.QSurfaceFormat_GreenBufferSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSurfaceFormat) HasAlpha() bool {
	if ptr.Pointer() != nil {
		return C.QSurfaceFormat_HasAlpha(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSurfaceFormat) MajorVersion() int {
	if ptr.Pointer() != nil {
		return int(C.QSurfaceFormat_MajorVersion(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSurfaceFormat) MinorVersion() int {
	if ptr.Pointer() != nil {
		return int(C.QSurfaceFormat_MinorVersion(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSurfaceFormat) Options() QSurfaceFormat__FormatOption {
	if ptr.Pointer() != nil {
		return QSurfaceFormat__FormatOption(C.QSurfaceFormat_Options(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSurfaceFormat) Profile() QSurfaceFormat__OpenGLContextProfile {
	if ptr.Pointer() != nil {
		return QSurfaceFormat__OpenGLContextProfile(C.QSurfaceFormat_Profile(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSurfaceFormat) RedBufferSize() int {
	if ptr.Pointer() != nil {
		return int(C.QSurfaceFormat_RedBufferSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSurfaceFormat) RenderableType() QSurfaceFormat__RenderableType {
	if ptr.Pointer() != nil {
		return QSurfaceFormat__RenderableType(C.QSurfaceFormat_RenderableType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSurfaceFormat) Samples() int {
	if ptr.Pointer() != nil {
		return int(C.QSurfaceFormat_Samples(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSurfaceFormat) SetAlphaBufferSize(size int) {
	if ptr.Pointer() != nil {
		C.QSurfaceFormat_SetAlphaBufferSize(ptr.Pointer(), C.int(size))
	}
}

func (ptr *QSurfaceFormat) SetBlueBufferSize(size int) {
	if ptr.Pointer() != nil {
		C.QSurfaceFormat_SetBlueBufferSize(ptr.Pointer(), C.int(size))
	}
}

func QSurfaceFormat_SetDefaultFormat(format QSurfaceFormat_ITF) {
	C.QSurfaceFormat_QSurfaceFormat_SetDefaultFormat(PointerFromQSurfaceFormat(format))
}

func (ptr *QSurfaceFormat) SetDepthBufferSize(size int) {
	if ptr.Pointer() != nil {
		C.QSurfaceFormat_SetDepthBufferSize(ptr.Pointer(), C.int(size))
	}
}

func (ptr *QSurfaceFormat) SetGreenBufferSize(size int) {
	if ptr.Pointer() != nil {
		C.QSurfaceFormat_SetGreenBufferSize(ptr.Pointer(), C.int(size))
	}
}

func (ptr *QSurfaceFormat) SetMajorVersion(major int) {
	if ptr.Pointer() != nil {
		C.QSurfaceFormat_SetMajorVersion(ptr.Pointer(), C.int(major))
	}
}

func (ptr *QSurfaceFormat) SetMinorVersion(minor int) {
	if ptr.Pointer() != nil {
		C.QSurfaceFormat_SetMinorVersion(ptr.Pointer(), C.int(minor))
	}
}

func (ptr *QSurfaceFormat) SetOption(option QSurfaceFormat__FormatOption, on bool) {
	if ptr.Pointer() != nil {
		C.QSurfaceFormat_SetOption(ptr.Pointer(), C.int(option), C.int(qt.GoBoolToInt(on)))
	}
}

func (ptr *QSurfaceFormat) SetOptions(options QSurfaceFormat__FormatOption) {
	if ptr.Pointer() != nil {
		C.QSurfaceFormat_SetOptions(ptr.Pointer(), C.int(options))
	}
}

func (ptr *QSurfaceFormat) SetProfile(profile QSurfaceFormat__OpenGLContextProfile) {
	if ptr.Pointer() != nil {
		C.QSurfaceFormat_SetProfile(ptr.Pointer(), C.int(profile))
	}
}

func (ptr *QSurfaceFormat) SetRedBufferSize(size int) {
	if ptr.Pointer() != nil {
		C.QSurfaceFormat_SetRedBufferSize(ptr.Pointer(), C.int(size))
	}
}

func (ptr *QSurfaceFormat) SetRenderableType(ty QSurfaceFormat__RenderableType) {
	if ptr.Pointer() != nil {
		C.QSurfaceFormat_SetRenderableType(ptr.Pointer(), C.int(ty))
	}
}

func (ptr *QSurfaceFormat) SetSamples(numSamples int) {
	if ptr.Pointer() != nil {
		C.QSurfaceFormat_SetSamples(ptr.Pointer(), C.int(numSamples))
	}
}

func (ptr *QSurfaceFormat) SetStencilBufferSize(size int) {
	if ptr.Pointer() != nil {
		C.QSurfaceFormat_SetStencilBufferSize(ptr.Pointer(), C.int(size))
	}
}

func (ptr *QSurfaceFormat) SetStereo(enable bool) {
	if ptr.Pointer() != nil {
		C.QSurfaceFormat_SetStereo(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QSurfaceFormat) SetSwapBehavior(behavior QSurfaceFormat__SwapBehavior) {
	if ptr.Pointer() != nil {
		C.QSurfaceFormat_SetSwapBehavior(ptr.Pointer(), C.int(behavior))
	}
}

func (ptr *QSurfaceFormat) SetSwapInterval(interval int) {
	if ptr.Pointer() != nil {
		C.QSurfaceFormat_SetSwapInterval(ptr.Pointer(), C.int(interval))
	}
}

func (ptr *QSurfaceFormat) SetVersion(major int, minor int) {
	if ptr.Pointer() != nil {
		C.QSurfaceFormat_SetVersion(ptr.Pointer(), C.int(major), C.int(minor))
	}
}

func (ptr *QSurfaceFormat) StencilBufferSize() int {
	if ptr.Pointer() != nil {
		return int(C.QSurfaceFormat_StencilBufferSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSurfaceFormat) Stereo() bool {
	if ptr.Pointer() != nil {
		return C.QSurfaceFormat_Stereo(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSurfaceFormat) SwapBehavior() QSurfaceFormat__SwapBehavior {
	if ptr.Pointer() != nil {
		return QSurfaceFormat__SwapBehavior(C.QSurfaceFormat_SwapBehavior(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSurfaceFormat) SwapInterval() int {
	if ptr.Pointer() != nil {
		return int(C.QSurfaceFormat_SwapInterval(ptr.Pointer()))
	}
	return 0
}

func (ptr *QSurfaceFormat) TestOption(option QSurfaceFormat__FormatOption) bool {
	if ptr.Pointer() != nil {
		return C.QSurfaceFormat_TestOption(ptr.Pointer(), C.int(option)) != 0
	}
	return false
}

func (ptr *QSurfaceFormat) DestroyQSurfaceFormat() {
	if ptr.Pointer() != nil {
		C.QSurfaceFormat_DestroyQSurfaceFormat(ptr.Pointer())
	}
}
