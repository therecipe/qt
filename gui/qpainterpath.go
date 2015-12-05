package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QPainterPath struct {
	ptr unsafe.Pointer
}

type QPainterPath_ITF interface {
	QPainterPath_PTR() *QPainterPath
}

func (p *QPainterPath) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QPainterPath) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQPainterPath(ptr QPainterPath_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPainterPath_PTR().Pointer()
	}
	return nil
}

func NewQPainterPathFromPointer(ptr unsafe.Pointer) *QPainterPath {
	var n = new(QPainterPath)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPainterPath) QPainterPath_PTR() *QPainterPath {
	return ptr
}

//QPainterPath::ElementType
type QPainterPath__ElementType int64

const (
	QPainterPath__MoveToElement      = QPainterPath__ElementType(0)
	QPainterPath__LineToElement      = QPainterPath__ElementType(1)
	QPainterPath__CurveToElement     = QPainterPath__ElementType(2)
	QPainterPath__CurveToDataElement = QPainterPath__ElementType(3)
)

func NewQPainterPath3(path QPainterPath_ITF) *QPainterPath {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainterPath::QPainterPath")
		}
	}()

	return NewQPainterPathFromPointer(C.QPainterPath_NewQPainterPath3(PointerFromQPainterPath(path)))
}

func (ptr *QPainterPath) AddEllipse(boundingRectangle core.QRectF_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainterPath::addEllipse")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainterPath_AddEllipse(ptr.Pointer(), core.PointerFromQRectF(boundingRectangle))
	}
}

func (ptr *QPainterPath) AddPath(path QPainterPath_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainterPath::addPath")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainterPath_AddPath(ptr.Pointer(), PointerFromQPainterPath(path))
	}
}

func (ptr *QPainterPath) AddRect(rectangle core.QRectF_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainterPath::addRect")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainterPath_AddRect(ptr.Pointer(), core.PointerFromQRectF(rectangle))
	}
}

func (ptr *QPainterPath) AddText(point core.QPointF_ITF, font QFont_ITF, text string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainterPath::addText")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainterPath_AddText(ptr.Pointer(), core.PointerFromQPointF(point), PointerFromQFont(font), C.CString(text))
	}
}

func (ptr *QPainterPath) ArcMoveTo(rectangle core.QRectF_ITF, angle float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainterPath::arcMoveTo")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainterPath_ArcMoveTo(ptr.Pointer(), core.PointerFromQRectF(rectangle), C.double(angle))
	}
}

func (ptr *QPainterPath) ArcTo(rectangle core.QRectF_ITF, startAngle float64, sweepLength float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainterPath::arcTo")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainterPath_ArcTo(ptr.Pointer(), core.PointerFromQRectF(rectangle), C.double(startAngle), C.double(sweepLength))
	}
}

func (ptr *QPainterPath) ConnectPath(path QPainterPath_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainterPath::connectPath")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainterPath_ConnectPath(ptr.Pointer(), PointerFromQPainterPath(path))
	}
}

func (ptr *QPainterPath) Contains(point core.QPointF_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainterPath::contains")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QPainterPath_Contains(ptr.Pointer(), core.PointerFromQPointF(point)) != 0
	}
	return false
}

func (ptr *QPainterPath) Contains2(rectangle core.QRectF_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainterPath::contains")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QPainterPath_Contains2(ptr.Pointer(), core.PointerFromQRectF(rectangle)) != 0
	}
	return false
}

func (ptr *QPainterPath) CubicTo(c1 core.QPointF_ITF, c2 core.QPointF_ITF, endPoint core.QPointF_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainterPath::cubicTo")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainterPath_CubicTo(ptr.Pointer(), core.PointerFromQPointF(c1), core.PointerFromQPointF(c2), core.PointerFromQPointF(endPoint))
	}
}

func (ptr *QPainterPath) ElementCount() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainterPath::elementCount")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QPainterPath_ElementCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPainterPath) Intersects(rectangle core.QRectF_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainterPath::intersects")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QPainterPath_Intersects(ptr.Pointer(), core.PointerFromQRectF(rectangle)) != 0
	}
	return false
}

func (ptr *QPainterPath) IsEmpty() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainterPath::isEmpty")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QPainterPath_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPainterPath) LineTo(endPoint core.QPointF_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainterPath::lineTo")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainterPath_LineTo(ptr.Pointer(), core.PointerFromQPointF(endPoint))
	}
}

func (ptr *QPainterPath) MoveTo(point core.QPointF_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainterPath::moveTo")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainterPath_MoveTo(ptr.Pointer(), core.PointerFromQPointF(point))
	}
}

func (ptr *QPainterPath) QuadTo(c core.QPointF_ITF, endPoint core.QPointF_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainterPath::quadTo")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainterPath_QuadTo(ptr.Pointer(), core.PointerFromQPointF(c), core.PointerFromQPointF(endPoint))
	}
}

func (ptr *QPainterPath) SetElementPositionAt(index int, x float64, y float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainterPath::setElementPositionAt")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainterPath_SetElementPositionAt(ptr.Pointer(), C.int(index), C.double(x), C.double(y))
	}
}

func (ptr *QPainterPath) SetFillRule(fillRule core.Qt__FillRule) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainterPath::setFillRule")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainterPath_SetFillRule(ptr.Pointer(), C.int(fillRule))
	}
}

func NewQPainterPath() *QPainterPath {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainterPath::QPainterPath")
		}
	}()

	return NewQPainterPathFromPointer(C.QPainterPath_NewQPainterPath())
}

func NewQPainterPath2(startPoint core.QPointF_ITF) *QPainterPath {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainterPath::QPainterPath")
		}
	}()

	return NewQPainterPathFromPointer(C.QPainterPath_NewQPainterPath2(core.PointerFromQPointF(startPoint)))
}

func (ptr *QPainterPath) AddEllipse3(center core.QPointF_ITF, rx float64, ry float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainterPath::addEllipse")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainterPath_AddEllipse3(ptr.Pointer(), core.PointerFromQPointF(center), C.double(rx), C.double(ry))
	}
}

func (ptr *QPainterPath) AddEllipse2(x float64, y float64, width float64, height float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainterPath::addEllipse")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainterPath_AddEllipse2(ptr.Pointer(), C.double(x), C.double(y), C.double(width), C.double(height))
	}
}

func (ptr *QPainterPath) AddPolygon(polygon QPolygonF_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainterPath::addPolygon")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainterPath_AddPolygon(ptr.Pointer(), PointerFromQPolygonF(polygon))
	}
}

func (ptr *QPainterPath) AddRect2(x float64, y float64, width float64, height float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainterPath::addRect")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainterPath_AddRect2(ptr.Pointer(), C.double(x), C.double(y), C.double(width), C.double(height))
	}
}

func (ptr *QPainterPath) AddRegion(region QRegion_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainterPath::addRegion")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainterPath_AddRegion(ptr.Pointer(), PointerFromQRegion(region))
	}
}

func (ptr *QPainterPath) AddRoundedRect(rect core.QRectF_ITF, xRadius float64, yRadius float64, mode core.Qt__SizeMode) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainterPath::addRoundedRect")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainterPath_AddRoundedRect(ptr.Pointer(), core.PointerFromQRectF(rect), C.double(xRadius), C.double(yRadius), C.int(mode))
	}
}

func (ptr *QPainterPath) AddRoundedRect2(x float64, y float64, w float64, h float64, xRadius float64, yRadius float64, mode core.Qt__SizeMode) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainterPath::addRoundedRect")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainterPath_AddRoundedRect2(ptr.Pointer(), C.double(x), C.double(y), C.double(w), C.double(h), C.double(xRadius), C.double(yRadius), C.int(mode))
	}
}

func (ptr *QPainterPath) AddText2(x float64, y float64, font QFont_ITF, text string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainterPath::addText")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainterPath_AddText2(ptr.Pointer(), C.double(x), C.double(y), PointerFromQFont(font), C.CString(text))
	}
}

func (ptr *QPainterPath) AngleAtPercent(t float64) float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainterPath::angleAtPercent")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QPainterPath_AngleAtPercent(ptr.Pointer(), C.double(t)))
	}
	return 0
}

func (ptr *QPainterPath) ArcMoveTo2(x float64, y float64, width float64, height float64, angle float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainterPath::arcMoveTo")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainterPath_ArcMoveTo2(ptr.Pointer(), C.double(x), C.double(y), C.double(width), C.double(height), C.double(angle))
	}
}

func (ptr *QPainterPath) ArcTo2(x float64, y float64, width float64, height float64, startAngle float64, sweepLength float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainterPath::arcTo")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainterPath_ArcTo2(ptr.Pointer(), C.double(x), C.double(y), C.double(width), C.double(height), C.double(startAngle), C.double(sweepLength))
	}
}

func (ptr *QPainterPath) CloseSubpath() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainterPath::closeSubpath")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainterPath_CloseSubpath(ptr.Pointer())
	}
}

func (ptr *QPainterPath) Contains3(p QPainterPath_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainterPath::contains")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QPainterPath_Contains3(ptr.Pointer(), PointerFromQPainterPath(p)) != 0
	}
	return false
}

func (ptr *QPainterPath) CubicTo2(c1X float64, c1Y float64, c2X float64, c2Y float64, endPointX float64, endPointY float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainterPath::cubicTo")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainterPath_CubicTo2(ptr.Pointer(), C.double(c1X), C.double(c1Y), C.double(c2X), C.double(c2Y), C.double(endPointX), C.double(endPointY))
	}
}

func (ptr *QPainterPath) FillRule() core.Qt__FillRule {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainterPath::fillRule")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__FillRule(C.QPainterPath_FillRule(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPainterPath) Intersects2(p QPainterPath_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainterPath::intersects")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QPainterPath_Intersects2(ptr.Pointer(), PointerFromQPainterPath(p)) != 0
	}
	return false
}

func (ptr *QPainterPath) Length() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainterPath::length")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QPainterPath_Length(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPainterPath) LineTo2(x float64, y float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainterPath::lineTo")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainterPath_LineTo2(ptr.Pointer(), C.double(x), C.double(y))
	}
}

func (ptr *QPainterPath) MoveTo2(x float64, y float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainterPath::moveTo")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainterPath_MoveTo2(ptr.Pointer(), C.double(x), C.double(y))
	}
}

func (ptr *QPainterPath) PercentAtLength(len float64) float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainterPath::percentAtLength")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QPainterPath_PercentAtLength(ptr.Pointer(), C.double(len)))
	}
	return 0
}

func (ptr *QPainterPath) QuadTo2(cx float64, cy float64, endPointX float64, endPointY float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainterPath::quadTo")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainterPath_QuadTo2(ptr.Pointer(), C.double(cx), C.double(cy), C.double(endPointX), C.double(endPointY))
	}
}

func (ptr *QPainterPath) SlopeAtPercent(t float64) float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainterPath::slopeAtPercent")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QPainterPath_SlopeAtPercent(ptr.Pointer(), C.double(t)))
	}
	return 0
}

func (ptr *QPainterPath) Swap(other QPainterPath_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainterPath::swap")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainterPath_Swap(ptr.Pointer(), PointerFromQPainterPath(other))
	}
}

func (ptr *QPainterPath) Translate2(offset core.QPointF_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainterPath::translate")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainterPath_Translate2(ptr.Pointer(), core.PointerFromQPointF(offset))
	}
}

func (ptr *QPainterPath) Translate(dx float64, dy float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainterPath::translate")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainterPath_Translate(ptr.Pointer(), C.double(dx), C.double(dy))
	}
}

func (ptr *QPainterPath) DestroyQPainterPath() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainterPath::~QPainterPath")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainterPath_DestroyQPainterPath(ptr.Pointer())
	}
}
