package gui

//#include "qpainterpath.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QPainterPath struct {
	ptr unsafe.Pointer
}

type QPainterPathITF interface {
	QPainterPathPTR() *QPainterPath
}

func (p *QPainterPath) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QPainterPath) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQPainterPath(ptr QPainterPathITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPainterPathPTR().Pointer()
	}
	return nil
}

func QPainterPathFromPointer(ptr unsafe.Pointer) *QPainterPath {
	var n = new(QPainterPath)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPainterPath) QPainterPathPTR() *QPainterPath {
	return ptr
}

//QPainterPath::ElementType
type QPainterPath__ElementType int

var (
	QPainterPath__MoveToElement      = QPainterPath__ElementType(0)
	QPainterPath__LineToElement      = QPainterPath__ElementType(1)
	QPainterPath__CurveToElement     = QPainterPath__ElementType(2)
	QPainterPath__CurveToDataElement = QPainterPath__ElementType(3)
)

func NewQPainterPath3(path QPainterPathITF) *QPainterPath {
	return QPainterPathFromPointer(unsafe.Pointer(C.QPainterPath_NewQPainterPath3(C.QtObjectPtr(PointerFromQPainterPath(path)))))
}

func (ptr *QPainterPath) AddEllipse(boundingRectangle core.QRectFITF) {
	if ptr.Pointer() != nil {
		C.QPainterPath_AddEllipse(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRectF(boundingRectangle)))
	}
}

func (ptr *QPainterPath) AddPath(path QPainterPathITF) {
	if ptr.Pointer() != nil {
		C.QPainterPath_AddPath(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPainterPath(path)))
	}
}

func (ptr *QPainterPath) AddRect(rectangle core.QRectFITF) {
	if ptr.Pointer() != nil {
		C.QPainterPath_AddRect(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRectF(rectangle)))
	}
}

func (ptr *QPainterPath) AddText(point core.QPointFITF, font QFontITF, text string) {
	if ptr.Pointer() != nil {
		C.QPainterPath_AddText(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPointF(point)), C.QtObjectPtr(PointerFromQFont(font)), C.CString(text))
	}
}

func (ptr *QPainterPath) ConnectPath(path QPainterPathITF) {
	if ptr.Pointer() != nil {
		C.QPainterPath_ConnectPath(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPainterPath(path)))
	}
}

func (ptr *QPainterPath) Contains(point core.QPointFITF) bool {
	if ptr.Pointer() != nil {
		return C.QPainterPath_Contains(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPointF(point))) != 0
	}
	return false
}

func (ptr *QPainterPath) Contains2(rectangle core.QRectFITF) bool {
	if ptr.Pointer() != nil {
		return C.QPainterPath_Contains2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRectF(rectangle))) != 0
	}
	return false
}

func (ptr *QPainterPath) CubicTo(c1 core.QPointFITF, c2 core.QPointFITF, endPoint core.QPointFITF) {
	if ptr.Pointer() != nil {
		C.QPainterPath_CubicTo(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPointF(c1)), C.QtObjectPtr(core.PointerFromQPointF(c2)), C.QtObjectPtr(core.PointerFromQPointF(endPoint)))
	}
}

func (ptr *QPainterPath) ElementCount() int {
	if ptr.Pointer() != nil {
		return int(C.QPainterPath_ElementCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QPainterPath) Intersects(rectangle core.QRectFITF) bool {
	if ptr.Pointer() != nil {
		return C.QPainterPath_Intersects(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRectF(rectangle))) != 0
	}
	return false
}

func (ptr *QPainterPath) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return C.QPainterPath_IsEmpty(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QPainterPath) LineTo(endPoint core.QPointFITF) {
	if ptr.Pointer() != nil {
		C.QPainterPath_LineTo(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPointF(endPoint)))
	}
}

func (ptr *QPainterPath) MoveTo(point core.QPointFITF) {
	if ptr.Pointer() != nil {
		C.QPainterPath_MoveTo(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPointF(point)))
	}
}

func (ptr *QPainterPath) QuadTo(c core.QPointFITF, endPoint core.QPointFITF) {
	if ptr.Pointer() != nil {
		C.QPainterPath_QuadTo(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPointF(c)), C.QtObjectPtr(core.PointerFromQPointF(endPoint)))
	}
}

func (ptr *QPainterPath) SetFillRule(fillRule core.Qt__FillRule) {
	if ptr.Pointer() != nil {
		C.QPainterPath_SetFillRule(C.QtObjectPtr(ptr.Pointer()), C.int(fillRule))
	}
}

func NewQPainterPath() *QPainterPath {
	return QPainterPathFromPointer(unsafe.Pointer(C.QPainterPath_NewQPainterPath()))
}

func NewQPainterPath2(startPoint core.QPointFITF) *QPainterPath {
	return QPainterPathFromPointer(unsafe.Pointer(C.QPainterPath_NewQPainterPath2(C.QtObjectPtr(core.PointerFromQPointF(startPoint)))))
}

func (ptr *QPainterPath) AddPolygon(polygon QPolygonFITF) {
	if ptr.Pointer() != nil {
		C.QPainterPath_AddPolygon(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPolygonF(polygon)))
	}
}

func (ptr *QPainterPath) AddRegion(region QRegionITF) {
	if ptr.Pointer() != nil {
		C.QPainterPath_AddRegion(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQRegion(region)))
	}
}

func (ptr *QPainterPath) CloseSubpath() {
	if ptr.Pointer() != nil {
		C.QPainterPath_CloseSubpath(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QPainterPath) Contains3(p QPainterPathITF) bool {
	if ptr.Pointer() != nil {
		return C.QPainterPath_Contains3(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPainterPath(p))) != 0
	}
	return false
}

func (ptr *QPainterPath) FillRule() core.Qt__FillRule {
	if ptr.Pointer() != nil {
		return core.Qt__FillRule(C.QPainterPath_FillRule(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QPainterPath) Intersects2(p QPainterPathITF) bool {
	if ptr.Pointer() != nil {
		return C.QPainterPath_Intersects2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPainterPath(p))) != 0
	}
	return false
}

func (ptr *QPainterPath) Swap(other QPainterPathITF) {
	if ptr.Pointer() != nil {
		C.QPainterPath_Swap(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPainterPath(other)))
	}
}

func (ptr *QPainterPath) Translate2(offset core.QPointFITF) {
	if ptr.Pointer() != nil {
		C.QPainterPath_Translate2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPointF(offset)))
	}
}

func (ptr *QPainterPath) DestroyQPainterPath() {
	if ptr.Pointer() != nil {
		C.QPainterPath_DestroyQPainterPath(C.QtObjectPtr(ptr.Pointer()))
	}
}
