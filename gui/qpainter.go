package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QPainter struct {
	ptr unsafe.Pointer
}

type QPainter_ITF interface {
	QPainter_PTR() *QPainter
}

func (p *QPainter) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QPainter) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQPainter(ptr QPainter_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPainter_PTR().Pointer()
	}
	return nil
}

func NewQPainterFromPointer(ptr unsafe.Pointer) *QPainter {
	var n = new(QPainter)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPainter) QPainter_PTR() *QPainter {
	return ptr
}

//QPainter::CompositionMode
type QPainter__CompositionMode int64

const (
	QPainter__CompositionMode_SourceOver          = QPainter__CompositionMode(0)
	QPainter__CompositionMode_DestinationOver     = QPainter__CompositionMode(1)
	QPainter__CompositionMode_Clear               = QPainter__CompositionMode(2)
	QPainter__CompositionMode_Source              = QPainter__CompositionMode(3)
	QPainter__CompositionMode_Destination         = QPainter__CompositionMode(4)
	QPainter__CompositionMode_SourceIn            = QPainter__CompositionMode(5)
	QPainter__CompositionMode_DestinationIn       = QPainter__CompositionMode(6)
	QPainter__CompositionMode_SourceOut           = QPainter__CompositionMode(7)
	QPainter__CompositionMode_DestinationOut      = QPainter__CompositionMode(8)
	QPainter__CompositionMode_SourceAtop          = QPainter__CompositionMode(9)
	QPainter__CompositionMode_DestinationAtop     = QPainter__CompositionMode(10)
	QPainter__CompositionMode_Xor                 = QPainter__CompositionMode(11)
	QPainter__CompositionMode_Plus                = QPainter__CompositionMode(12)
	QPainter__CompositionMode_Multiply            = QPainter__CompositionMode(13)
	QPainter__CompositionMode_Screen              = QPainter__CompositionMode(14)
	QPainter__CompositionMode_Overlay             = QPainter__CompositionMode(15)
	QPainter__CompositionMode_Darken              = QPainter__CompositionMode(16)
	QPainter__CompositionMode_Lighten             = QPainter__CompositionMode(17)
	QPainter__CompositionMode_ColorDodge          = QPainter__CompositionMode(18)
	QPainter__CompositionMode_ColorBurn           = QPainter__CompositionMode(19)
	QPainter__CompositionMode_HardLight           = QPainter__CompositionMode(20)
	QPainter__CompositionMode_SoftLight           = QPainter__CompositionMode(21)
	QPainter__CompositionMode_Difference          = QPainter__CompositionMode(22)
	QPainter__CompositionMode_Exclusion           = QPainter__CompositionMode(23)
	QPainter__RasterOp_SourceOrDestination        = QPainter__CompositionMode(24)
	QPainter__RasterOp_SourceAndDestination       = QPainter__CompositionMode(25)
	QPainter__RasterOp_SourceXorDestination       = QPainter__CompositionMode(26)
	QPainter__RasterOp_NotSourceAndNotDestination = QPainter__CompositionMode(27)
	QPainter__RasterOp_NotSourceOrNotDestination  = QPainter__CompositionMode(28)
	QPainter__RasterOp_NotSourceXorDestination    = QPainter__CompositionMode(29)
	QPainter__RasterOp_NotSource                  = QPainter__CompositionMode(30)
	QPainter__RasterOp_NotSourceAndDestination    = QPainter__CompositionMode(31)
	QPainter__RasterOp_SourceAndNotDestination    = QPainter__CompositionMode(32)
	QPainter__RasterOp_NotSourceOrDestination     = QPainter__CompositionMode(33)
	QPainter__RasterOp_SourceOrNotDestination     = QPainter__CompositionMode(34)
	QPainter__RasterOp_ClearDestination           = QPainter__CompositionMode(35)
	QPainter__RasterOp_SetDestination             = QPainter__CompositionMode(36)
	QPainter__RasterOp_NotDestination             = QPainter__CompositionMode(37)
)

//QPainter::PixmapFragmentHint
type QPainter__PixmapFragmentHint int64

const (
	QPainter__OpaqueHint = QPainter__PixmapFragmentHint(0x01)
)

//QPainter::RenderHint
type QPainter__RenderHint int64

const (
	QPainter__Antialiasing            = QPainter__RenderHint(0x01)
	QPainter__TextAntialiasing        = QPainter__RenderHint(0x02)
	QPainter__SmoothPixmapTransform   = QPainter__RenderHint(0x04)
	QPainter__HighQualityAntialiasing = QPainter__RenderHint(0x08)
	QPainter__NonCosmeticDefaultPen   = QPainter__RenderHint(0x10)
	QPainter__Qt4CompatiblePainting   = QPainter__RenderHint(0x20)
)

func NewQPainter2(device QPaintDevice_ITF) *QPainter {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::QPainter")
		}
	}()

	return NewQPainterFromPointer(C.QPainter_NewQPainter2(PointerFromQPaintDevice(device)))
}

func (ptr *QPainter) Begin(device QPaintDevice_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::begin")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QPainter_Begin(ptr.Pointer(), PointerFromQPaintDevice(device)) != 0
	}
	return false
}

func (ptr *QPainter) DrawArc(rectangle core.QRectF_ITF, startAngle int, spanAngle int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawArc")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawArc(ptr.Pointer(), core.PointerFromQRectF(rectangle), C.int(startAngle), C.int(spanAngle))
	}
}

func (ptr *QPainter) DrawChord(rectangle core.QRectF_ITF, startAngle int, spanAngle int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawChord")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawChord(ptr.Pointer(), core.PointerFromQRectF(rectangle), C.int(startAngle), C.int(spanAngle))
	}
}

func (ptr *QPainter) DrawConvexPolygon2(points core.QPoint_ITF, pointCount int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawConvexPolygon")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawConvexPolygon2(ptr.Pointer(), core.PointerFromQPoint(points), C.int(pointCount))
	}
}

func (ptr *QPainter) DrawConvexPolygon(points core.QPointF_ITF, pointCount int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawConvexPolygon")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawConvexPolygon(ptr.Pointer(), core.PointerFromQPointF(points), C.int(pointCount))
	}
}

func (ptr *QPainter) DrawEllipse2(rectangle core.QRect_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawEllipse")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawEllipse2(ptr.Pointer(), core.PointerFromQRect(rectangle))
	}
}

func (ptr *QPainter) DrawEllipse(rectangle core.QRectF_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawEllipse")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawEllipse(ptr.Pointer(), core.PointerFromQRectF(rectangle))
	}
}

func (ptr *QPainter) DrawGlyphRun(position core.QPointF_ITF, glyphs QGlyphRun_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawGlyphRun")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawGlyphRun(ptr.Pointer(), core.PointerFromQPointF(position), PointerFromQGlyphRun(glyphs))
	}
}

func (ptr *QPainter) DrawImage3(point core.QPointF_ITF, image QImage_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawImage")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawImage3(ptr.Pointer(), core.PointerFromQPointF(point), PointerFromQImage(image))
	}
}

func (ptr *QPainter) DrawImage(target core.QRectF_ITF, image QImage_ITF, source core.QRectF_ITF, flags core.Qt__ImageConversionFlag) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawImage")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawImage(ptr.Pointer(), core.PointerFromQRectF(target), PointerFromQImage(image), core.PointerFromQRectF(source), C.int(flags))
	}
}

func (ptr *QPainter) DrawLines2(lines core.QLine_ITF, lineCount int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawLines")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawLines2(ptr.Pointer(), core.PointerFromQLine(lines), C.int(lineCount))
	}
}

func (ptr *QPainter) DrawPicture(point core.QPointF_ITF, picture QPicture_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawPicture")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawPicture(ptr.Pointer(), core.PointerFromQPointF(point), PointerFromQPicture(picture))
	}
}

func (ptr *QPainter) DrawPie(rectangle core.QRectF_ITF, startAngle int, spanAngle int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawPie")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawPie(ptr.Pointer(), core.PointerFromQRectF(rectangle), C.int(startAngle), C.int(spanAngle))
	}
}

func (ptr *QPainter) DrawPixmap5(point core.QPointF_ITF, pixmap QPixmap_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawPixmap")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawPixmap5(ptr.Pointer(), core.PointerFromQPointF(point), PointerFromQPixmap(pixmap))
	}
}

func (ptr *QPainter) DrawPixmap(target core.QRectF_ITF, pixmap QPixmap_ITF, source core.QRectF_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawPixmap")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawPixmap(ptr.Pointer(), core.PointerFromQRectF(target), PointerFromQPixmap(pixmap), core.PointerFromQRectF(source))
	}
}

func (ptr *QPainter) DrawRects2(rectangles core.QRect_ITF, rectCount int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawRects")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawRects2(ptr.Pointer(), core.PointerFromQRect(rectangles), C.int(rectCount))
	}
}

func (ptr *QPainter) DrawRects(rectangles core.QRectF_ITF, rectCount int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawRects")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawRects(ptr.Pointer(), core.PointerFromQRectF(rectangles), C.int(rectCount))
	}
}

func (ptr *QPainter) DrawText(position core.QPointF_ITF, text string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawText")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawText(ptr.Pointer(), core.PointerFromQPointF(position), C.CString(text))
	}
}

func (ptr *QPainter) DrawText5(rectangle core.QRect_ITF, flags int, text string, boundingRect core.QRect_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawText")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawText5(ptr.Pointer(), core.PointerFromQRect(rectangle), C.int(flags), C.CString(text), core.PointerFromQRect(boundingRect))
	}
}

func (ptr *QPainter) DrawText8(rectangle core.QRectF_ITF, text string, option QTextOption_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawText")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawText8(ptr.Pointer(), core.PointerFromQRectF(rectangle), C.CString(text), PointerFromQTextOption(option))
	}
}

func (ptr *QPainter) DrawText4(rectangle core.QRectF_ITF, flags int, text string, boundingRect core.QRectF_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawText")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawText4(ptr.Pointer(), core.PointerFromQRectF(rectangle), C.int(flags), C.CString(text), core.PointerFromQRectF(boundingRect))
	}
}

func (ptr *QPainter) DrawTiledPixmap(rectangle core.QRectF_ITF, pixmap QPixmap_ITF, position core.QPointF_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawTiledPixmap")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawTiledPixmap(ptr.Pointer(), core.PointerFromQRectF(rectangle), PointerFromQPixmap(pixmap), core.PointerFromQPointF(position))
	}
}

func (ptr *QPainter) EraseRect(rectangle core.QRectF_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::eraseRect")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_EraseRect(ptr.Pointer(), core.PointerFromQRectF(rectangle))
	}
}

func (ptr *QPainter) FillRect5(rectangle core.QRect_ITF, brush QBrush_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::fillRect")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_FillRect5(ptr.Pointer(), core.PointerFromQRect(rectangle), PointerFromQBrush(brush))
	}
}

func (ptr *QPainter) FillRect6(rectangle core.QRect_ITF, color QColor_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::fillRect")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_FillRect6(ptr.Pointer(), core.PointerFromQRect(rectangle), PointerFromQColor(color))
	}
}

func (ptr *QPainter) FillRect(rectangle core.QRectF_ITF, brush QBrush_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::fillRect")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_FillRect(ptr.Pointer(), core.PointerFromQRectF(rectangle), PointerFromQBrush(brush))
	}
}

func (ptr *QPainter) FillRect7(rectangle core.QRectF_ITF, color QColor_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::fillRect")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_FillRect7(ptr.Pointer(), core.PointerFromQRectF(rectangle), PointerFromQColor(color))
	}
}

func (ptr *QPainter) Rotate(angle float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::rotate")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_Rotate(ptr.Pointer(), C.double(angle))
	}
}

func (ptr *QPainter) SetBackground(brush QBrush_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::setBackground")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_SetBackground(ptr.Pointer(), PointerFromQBrush(brush))
	}
}

func (ptr *QPainter) SetBrushOrigin(position core.QPointF_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::setBrushOrigin")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_SetBrushOrigin(ptr.Pointer(), core.PointerFromQPointF(position))
	}
}

func (ptr *QPainter) SetClipPath(path QPainterPath_ITF, operation core.Qt__ClipOperation) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::setClipPath")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_SetClipPath(ptr.Pointer(), PointerFromQPainterPath(path), C.int(operation))
	}
}

func (ptr *QPainter) SetClipRect3(rectangle core.QRect_ITF, operation core.Qt__ClipOperation) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::setClipRect")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_SetClipRect3(ptr.Pointer(), core.PointerFromQRect(rectangle), C.int(operation))
	}
}

func (ptr *QPainter) SetClipRect(rectangle core.QRectF_ITF, operation core.Qt__ClipOperation) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::setClipRect")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_SetClipRect(ptr.Pointer(), core.PointerFromQRectF(rectangle), C.int(operation))
	}
}

func (ptr *QPainter) SetClipRegion(region QRegion_ITF, operation core.Qt__ClipOperation) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::setClipRegion")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_SetClipRegion(ptr.Pointer(), PointerFromQRegion(region), C.int(operation))
	}
}

func (ptr *QPainter) SetViewport(rectangle core.QRect_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::setViewport")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_SetViewport(ptr.Pointer(), core.PointerFromQRect(rectangle))
	}
}

func (ptr *QPainter) SetWindow(rectangle core.QRect_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::setWindow")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_SetWindow(ptr.Pointer(), core.PointerFromQRect(rectangle))
	}
}

func NewQPainter() *QPainter {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::QPainter")
		}
	}()

	return NewQPainterFromPointer(C.QPainter_NewQPainter())
}

func (ptr *QPainter) Background() *QBrush {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::background")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQBrushFromPointer(C.QPainter_Background(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPainter) BackgroundMode() core.Qt__BGMode {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::backgroundMode")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__BGMode(C.QPainter_BackgroundMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPainter) BeginNativePainting() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::beginNativePainting")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_BeginNativePainting(ptr.Pointer())
	}
}

func (ptr *QPainter) Brush() *QBrush {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::brush")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQBrushFromPointer(C.QPainter_Brush(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPainter) ClipRegion() *QRegion {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::clipRegion")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQRegionFromPointer(C.QPainter_ClipRegion(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPainter) CompositionMode() QPainter__CompositionMode {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::compositionMode")
		}
	}()

	if ptr.Pointer() != nil {
		return QPainter__CompositionMode(C.QPainter_CompositionMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPainter) Device() *QPaintDevice {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::device")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQPaintDeviceFromPointer(C.QPainter_Device(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPainter) DrawArc2(rectangle core.QRect_ITF, startAngle int, spanAngle int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawArc")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawArc2(ptr.Pointer(), core.PointerFromQRect(rectangle), C.int(startAngle), C.int(spanAngle))
	}
}

func (ptr *QPainter) DrawArc3(x int, y int, width int, height int, startAngle int, spanAngle int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawArc")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawArc3(ptr.Pointer(), C.int(x), C.int(y), C.int(width), C.int(height), C.int(startAngle), C.int(spanAngle))
	}
}

func (ptr *QPainter) DrawChord2(rectangle core.QRect_ITF, startAngle int, spanAngle int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawChord")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawChord2(ptr.Pointer(), core.PointerFromQRect(rectangle), C.int(startAngle), C.int(spanAngle))
	}
}

func (ptr *QPainter) DrawChord3(x int, y int, width int, height int, startAngle int, spanAngle int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawChord")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawChord3(ptr.Pointer(), C.int(x), C.int(y), C.int(width), C.int(height), C.int(startAngle), C.int(spanAngle))
	}
}

func (ptr *QPainter) DrawConvexPolygon4(polygon QPolygon_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawConvexPolygon")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawConvexPolygon4(ptr.Pointer(), PointerFromQPolygon(polygon))
	}
}

func (ptr *QPainter) DrawConvexPolygon3(polygon QPolygonF_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawConvexPolygon")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawConvexPolygon3(ptr.Pointer(), PointerFromQPolygonF(polygon))
	}
}

func (ptr *QPainter) DrawEllipse5(center core.QPoint_ITF, rx int, ry int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawEllipse")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawEllipse5(ptr.Pointer(), core.PointerFromQPoint(center), C.int(rx), C.int(ry))
	}
}

func (ptr *QPainter) DrawEllipse4(center core.QPointF_ITF, rx float64, ry float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawEllipse")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawEllipse4(ptr.Pointer(), core.PointerFromQPointF(center), C.double(rx), C.double(ry))
	}
}

func (ptr *QPainter) DrawEllipse3(x int, y int, width int, height int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawEllipse")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawEllipse3(ptr.Pointer(), C.int(x), C.int(y), C.int(width), C.int(height))
	}
}

func (ptr *QPainter) DrawImage4(point core.QPoint_ITF, image QImage_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawImage")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawImage4(ptr.Pointer(), core.PointerFromQPoint(point), PointerFromQImage(image))
	}
}

func (ptr *QPainter) DrawImage6(point core.QPoint_ITF, image QImage_ITF, source core.QRect_ITF, flags core.Qt__ImageConversionFlag) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawImage")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawImage6(ptr.Pointer(), core.PointerFromQPoint(point), PointerFromQImage(image), core.PointerFromQRect(source), C.int(flags))
	}
}

func (ptr *QPainter) DrawImage5(point core.QPointF_ITF, image QImage_ITF, source core.QRectF_ITF, flags core.Qt__ImageConversionFlag) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawImage")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawImage5(ptr.Pointer(), core.PointerFromQPointF(point), PointerFromQImage(image), core.PointerFromQRectF(source), C.int(flags))
	}
}

func (ptr *QPainter) DrawImage8(rectangle core.QRect_ITF, image QImage_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawImage")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawImage8(ptr.Pointer(), core.PointerFromQRect(rectangle), PointerFromQImage(image))
	}
}

func (ptr *QPainter) DrawImage2(target core.QRect_ITF, image QImage_ITF, source core.QRect_ITF, flags core.Qt__ImageConversionFlag) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawImage")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawImage2(ptr.Pointer(), core.PointerFromQRect(target), PointerFromQImage(image), core.PointerFromQRect(source), C.int(flags))
	}
}

func (ptr *QPainter) DrawImage7(rectangle core.QRectF_ITF, image QImage_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawImage")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawImage7(ptr.Pointer(), core.PointerFromQRectF(rectangle), PointerFromQImage(image))
	}
}

func (ptr *QPainter) DrawImage9(x int, y int, image QImage_ITF, sx int, sy int, sw int, sh int, flags core.Qt__ImageConversionFlag) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawImage")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawImage9(ptr.Pointer(), C.int(x), C.int(y), PointerFromQImage(image), C.int(sx), C.int(sy), C.int(sw), C.int(sh), C.int(flags))
	}
}

func (ptr *QPainter) DrawLine2(line core.QLine_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawLine")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawLine2(ptr.Pointer(), core.PointerFromQLine(line))
	}
}

func (ptr *QPainter) DrawLine(line core.QLineF_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawLine")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawLine(ptr.Pointer(), core.PointerFromQLineF(line))
	}
}

func (ptr *QPainter) DrawLine3(p1 core.QPoint_ITF, p2 core.QPoint_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawLine")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawLine3(ptr.Pointer(), core.PointerFromQPoint(p1), core.PointerFromQPoint(p2))
	}
}

func (ptr *QPainter) DrawLine4(p1 core.QPointF_ITF, p2 core.QPointF_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawLine")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawLine4(ptr.Pointer(), core.PointerFromQPointF(p1), core.PointerFromQPointF(p2))
	}
}

func (ptr *QPainter) DrawLine5(x1 int, y1 int, x2 int, y2 int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawLine")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawLine5(ptr.Pointer(), C.int(x1), C.int(y1), C.int(x2), C.int(y2))
	}
}

func (ptr *QPainter) DrawLines(lines core.QLineF_ITF, lineCount int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawLines")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawLines(ptr.Pointer(), core.PointerFromQLineF(lines), C.int(lineCount))
	}
}

func (ptr *QPainter) DrawLines4(pointPairs core.QPoint_ITF, lineCount int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawLines")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawLines4(ptr.Pointer(), core.PointerFromQPoint(pointPairs), C.int(lineCount))
	}
}

func (ptr *QPainter) DrawLines3(pointPairs core.QPointF_ITF, lineCount int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawLines")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawLines3(ptr.Pointer(), core.PointerFromQPointF(pointPairs), C.int(lineCount))
	}
}

func (ptr *QPainter) DrawPath(path QPainterPath_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawPath")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawPath(ptr.Pointer(), PointerFromQPainterPath(path))
	}
}

func (ptr *QPainter) DrawPicture2(point core.QPoint_ITF, picture QPicture_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawPicture")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawPicture2(ptr.Pointer(), core.PointerFromQPoint(point), PointerFromQPicture(picture))
	}
}

func (ptr *QPainter) DrawPicture3(x int, y int, picture QPicture_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawPicture")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawPicture3(ptr.Pointer(), C.int(x), C.int(y), PointerFromQPicture(picture))
	}
}

func (ptr *QPainter) DrawPie2(rectangle core.QRect_ITF, startAngle int, spanAngle int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawPie")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawPie2(ptr.Pointer(), core.PointerFromQRect(rectangle), C.int(startAngle), C.int(spanAngle))
	}
}

func (ptr *QPainter) DrawPie3(x int, y int, width int, height int, startAngle int, spanAngle int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawPie")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawPie3(ptr.Pointer(), C.int(x), C.int(y), C.int(width), C.int(height), C.int(startAngle), C.int(spanAngle))
	}
}

func (ptr *QPainter) DrawPixmap6(point core.QPoint_ITF, pixmap QPixmap_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawPixmap")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawPixmap6(ptr.Pointer(), core.PointerFromQPoint(point), PointerFromQPixmap(pixmap))
	}
}

func (ptr *QPainter) DrawPixmap4(point core.QPoint_ITF, pixmap QPixmap_ITF, source core.QRect_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawPixmap")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawPixmap4(ptr.Pointer(), core.PointerFromQPoint(point), PointerFromQPixmap(pixmap), core.PointerFromQRect(source))
	}
}

func (ptr *QPainter) DrawPixmap3(point core.QPointF_ITF, pixmap QPixmap_ITF, source core.QRectF_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawPixmap")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawPixmap3(ptr.Pointer(), core.PointerFromQPointF(point), PointerFromQPixmap(pixmap), core.PointerFromQRectF(source))
	}
}

func (ptr *QPainter) DrawPixmap8(rectangle core.QRect_ITF, pixmap QPixmap_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawPixmap")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawPixmap8(ptr.Pointer(), core.PointerFromQRect(rectangle), PointerFromQPixmap(pixmap))
	}
}

func (ptr *QPainter) DrawPixmap2(target core.QRect_ITF, pixmap QPixmap_ITF, source core.QRect_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawPixmap")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawPixmap2(ptr.Pointer(), core.PointerFromQRect(target), PointerFromQPixmap(pixmap), core.PointerFromQRect(source))
	}
}

func (ptr *QPainter) DrawPixmap7(x int, y int, pixmap QPixmap_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawPixmap")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawPixmap7(ptr.Pointer(), C.int(x), C.int(y), PointerFromQPixmap(pixmap))
	}
}

func (ptr *QPainter) DrawPixmap11(x int, y int, pixmap QPixmap_ITF, sx int, sy int, sw int, sh int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawPixmap")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawPixmap11(ptr.Pointer(), C.int(x), C.int(y), PointerFromQPixmap(pixmap), C.int(sx), C.int(sy), C.int(sw), C.int(sh))
	}
}

func (ptr *QPainter) DrawPixmap10(x int, y int, w int, h int, pixmap QPixmap_ITF, sx int, sy int, sw int, sh int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawPixmap")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawPixmap10(ptr.Pointer(), C.int(x), C.int(y), C.int(w), C.int(h), PointerFromQPixmap(pixmap), C.int(sx), C.int(sy), C.int(sw), C.int(sh))
	}
}

func (ptr *QPainter) DrawPixmap9(x int, y int, width int, height int, pixmap QPixmap_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawPixmap")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawPixmap9(ptr.Pointer(), C.int(x), C.int(y), C.int(width), C.int(height), PointerFromQPixmap(pixmap))
	}
}

func (ptr *QPainter) DrawPoint2(position core.QPoint_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawPoint")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawPoint2(ptr.Pointer(), core.PointerFromQPoint(position))
	}
}

func (ptr *QPainter) DrawPoint(position core.QPointF_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawPoint")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawPoint(ptr.Pointer(), core.PointerFromQPointF(position))
	}
}

func (ptr *QPainter) DrawPoint3(x int, y int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawPoint")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawPoint3(ptr.Pointer(), C.int(x), C.int(y))
	}
}

func (ptr *QPainter) DrawPoints2(points core.QPoint_ITF, pointCount int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawPoints")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawPoints2(ptr.Pointer(), core.PointerFromQPoint(points), C.int(pointCount))
	}
}

func (ptr *QPainter) DrawPoints(points core.QPointF_ITF, pointCount int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawPoints")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawPoints(ptr.Pointer(), core.PointerFromQPointF(points), C.int(pointCount))
	}
}

func (ptr *QPainter) DrawPoints4(points QPolygon_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawPoints")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawPoints4(ptr.Pointer(), PointerFromQPolygon(points))
	}
}

func (ptr *QPainter) DrawPoints3(points QPolygonF_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawPoints")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawPoints3(ptr.Pointer(), PointerFromQPolygonF(points))
	}
}

func (ptr *QPainter) DrawPolygon2(points core.QPoint_ITF, pointCount int, fillRule core.Qt__FillRule) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawPolygon")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawPolygon2(ptr.Pointer(), core.PointerFromQPoint(points), C.int(pointCount), C.int(fillRule))
	}
}

func (ptr *QPainter) DrawPolygon(points core.QPointF_ITF, pointCount int, fillRule core.Qt__FillRule) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawPolygon")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawPolygon(ptr.Pointer(), core.PointerFromQPointF(points), C.int(pointCount), C.int(fillRule))
	}
}

func (ptr *QPainter) DrawPolygon4(points QPolygon_ITF, fillRule core.Qt__FillRule) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawPolygon")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawPolygon4(ptr.Pointer(), PointerFromQPolygon(points), C.int(fillRule))
	}
}

func (ptr *QPainter) DrawPolygon3(points QPolygonF_ITF, fillRule core.Qt__FillRule) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawPolygon")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawPolygon3(ptr.Pointer(), PointerFromQPolygonF(points), C.int(fillRule))
	}
}

func (ptr *QPainter) DrawPolyline2(points core.QPoint_ITF, pointCount int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawPolyline")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawPolyline2(ptr.Pointer(), core.PointerFromQPoint(points), C.int(pointCount))
	}
}

func (ptr *QPainter) DrawPolyline(points core.QPointF_ITF, pointCount int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawPolyline")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawPolyline(ptr.Pointer(), core.PointerFromQPointF(points), C.int(pointCount))
	}
}

func (ptr *QPainter) DrawPolyline4(points QPolygon_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawPolyline")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawPolyline4(ptr.Pointer(), PointerFromQPolygon(points))
	}
}

func (ptr *QPainter) DrawPolyline3(points QPolygonF_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawPolyline")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawPolyline3(ptr.Pointer(), PointerFromQPolygonF(points))
	}
}

func (ptr *QPainter) DrawRect2(rectangle core.QRect_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawRect")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawRect2(ptr.Pointer(), core.PointerFromQRect(rectangle))
	}
}

func (ptr *QPainter) DrawRect(rectangle core.QRectF_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawRect")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawRect(ptr.Pointer(), core.PointerFromQRectF(rectangle))
	}
}

func (ptr *QPainter) DrawRect3(x int, y int, width int, height int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawRect")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawRect3(ptr.Pointer(), C.int(x), C.int(y), C.int(width), C.int(height))
	}
}

func (ptr *QPainter) DrawRoundedRect2(rect core.QRect_ITF, xRadius float64, yRadius float64, mode core.Qt__SizeMode) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawRoundedRect")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawRoundedRect2(ptr.Pointer(), core.PointerFromQRect(rect), C.double(xRadius), C.double(yRadius), C.int(mode))
	}
}

func (ptr *QPainter) DrawRoundedRect(rect core.QRectF_ITF, xRadius float64, yRadius float64, mode core.Qt__SizeMode) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawRoundedRect")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawRoundedRect(ptr.Pointer(), core.PointerFromQRectF(rect), C.double(xRadius), C.double(yRadius), C.int(mode))
	}
}

func (ptr *QPainter) DrawRoundedRect3(x int, y int, w int, h int, xRadius float64, yRadius float64, mode core.Qt__SizeMode) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawRoundedRect")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawRoundedRect3(ptr.Pointer(), C.int(x), C.int(y), C.int(w), C.int(h), C.double(xRadius), C.double(yRadius), C.int(mode))
	}
}

func (ptr *QPainter) DrawStaticText2(topLeftPosition core.QPoint_ITF, staticText QStaticText_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawStaticText")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawStaticText2(ptr.Pointer(), core.PointerFromQPoint(topLeftPosition), PointerFromQStaticText(staticText))
	}
}

func (ptr *QPainter) DrawStaticText(topLeftPosition core.QPointF_ITF, staticText QStaticText_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawStaticText")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawStaticText(ptr.Pointer(), core.PointerFromQPointF(topLeftPosition), PointerFromQStaticText(staticText))
	}
}

func (ptr *QPainter) DrawStaticText3(left int, top int, staticText QStaticText_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawStaticText")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawStaticText3(ptr.Pointer(), C.int(left), C.int(top), PointerFromQStaticText(staticText))
	}
}

func (ptr *QPainter) DrawText3(position core.QPoint_ITF, text string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawText")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawText3(ptr.Pointer(), core.PointerFromQPoint(position), C.CString(text))
	}
}

func (ptr *QPainter) DrawText6(x int, y int, text string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawText")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawText6(ptr.Pointer(), C.int(x), C.int(y), C.CString(text))
	}
}

func (ptr *QPainter) DrawText7(x int, y int, width int, height int, flags int, text string, boundingRect core.QRect_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawText")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawText7(ptr.Pointer(), C.int(x), C.int(y), C.int(width), C.int(height), C.int(flags), C.CString(text), core.PointerFromQRect(boundingRect))
	}
}

func (ptr *QPainter) DrawTiledPixmap2(rectangle core.QRect_ITF, pixmap QPixmap_ITF, position core.QPoint_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawTiledPixmap")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawTiledPixmap2(ptr.Pointer(), core.PointerFromQRect(rectangle), PointerFromQPixmap(pixmap), core.PointerFromQPoint(position))
	}
}

func (ptr *QPainter) DrawTiledPixmap3(x int, y int, width int, height int, pixmap QPixmap_ITF, sx int, sy int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::drawTiledPixmap")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DrawTiledPixmap3(ptr.Pointer(), C.int(x), C.int(y), C.int(width), C.int(height), PointerFromQPixmap(pixmap), C.int(sx), C.int(sy))
	}
}

func (ptr *QPainter) End() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::end")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QPainter_End(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPainter) EndNativePainting() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::endNativePainting")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_EndNativePainting(ptr.Pointer())
	}
}

func (ptr *QPainter) EraseRect2(rectangle core.QRect_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::eraseRect")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_EraseRect2(ptr.Pointer(), core.PointerFromQRect(rectangle))
	}
}

func (ptr *QPainter) EraseRect3(x int, y int, width int, height int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::eraseRect")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_EraseRect3(ptr.Pointer(), C.int(x), C.int(y), C.int(width), C.int(height))
	}
}

func (ptr *QPainter) FillPath(path QPainterPath_ITF, brush QBrush_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::fillPath")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_FillPath(ptr.Pointer(), PointerFromQPainterPath(path), PointerFromQBrush(brush))
	}
}

func (ptr *QPainter) FillRect3(rectangle core.QRect_ITF, style core.Qt__BrushStyle) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::fillRect")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_FillRect3(ptr.Pointer(), core.PointerFromQRect(rectangle), C.int(style))
	}
}

func (ptr *QPainter) FillRect11(rectangle core.QRect_ITF, color core.Qt__GlobalColor) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::fillRect")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_FillRect11(ptr.Pointer(), core.PointerFromQRect(rectangle), C.int(color))
	}
}

func (ptr *QPainter) FillRect4(rectangle core.QRectF_ITF, style core.Qt__BrushStyle) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::fillRect")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_FillRect4(ptr.Pointer(), core.PointerFromQRectF(rectangle), C.int(style))
	}
}

func (ptr *QPainter) FillRect12(rectangle core.QRectF_ITF, color core.Qt__GlobalColor) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::fillRect")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_FillRect12(ptr.Pointer(), core.PointerFromQRectF(rectangle), C.int(color))
	}
}

func (ptr *QPainter) FillRect2(x int, y int, width int, height int, style core.Qt__BrushStyle) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::fillRect")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_FillRect2(ptr.Pointer(), C.int(x), C.int(y), C.int(width), C.int(height), C.int(style))
	}
}

func (ptr *QPainter) FillRect10(x int, y int, width int, height int, color core.Qt__GlobalColor) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::fillRect")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_FillRect10(ptr.Pointer(), C.int(x), C.int(y), C.int(width), C.int(height), C.int(color))
	}
}

func (ptr *QPainter) FillRect8(x int, y int, width int, height int, brush QBrush_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::fillRect")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_FillRect8(ptr.Pointer(), C.int(x), C.int(y), C.int(width), C.int(height), PointerFromQBrush(brush))
	}
}

func (ptr *QPainter) FillRect9(x int, y int, width int, height int, color QColor_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::fillRect")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_FillRect9(ptr.Pointer(), C.int(x), C.int(y), C.int(width), C.int(height), PointerFromQColor(color))
	}
}

func (ptr *QPainter) HasClipping() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::hasClipping")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QPainter_HasClipping(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPainter) IsActive() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::isActive")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QPainter_IsActive(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPainter) LayoutDirection() core.Qt__LayoutDirection {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::layoutDirection")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__LayoutDirection(C.QPainter_LayoutDirection(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPainter) Opacity() float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::opacity")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QPainter_Opacity(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPainter) PaintEngine() *QPaintEngine {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::paintEngine")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQPaintEngineFromPointer(C.QPainter_PaintEngine(ptr.Pointer()))
	}
	return nil
}

func (ptr *QPainter) RenderHints() QPainter__RenderHint {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::renderHints")
		}
	}()

	if ptr.Pointer() != nil {
		return QPainter__RenderHint(C.QPainter_RenderHints(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPainter) ResetTransform() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::resetTransform")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_ResetTransform(ptr.Pointer())
	}
}

func (ptr *QPainter) Restore() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::restore")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_Restore(ptr.Pointer())
	}
}

func (ptr *QPainter) Save() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::save")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_Save(ptr.Pointer())
	}
}

func (ptr *QPainter) Scale(sx float64, sy float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::scale")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_Scale(ptr.Pointer(), C.double(sx), C.double(sy))
	}
}

func (ptr *QPainter) SetBackgroundMode(mode core.Qt__BGMode) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::setBackgroundMode")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_SetBackgroundMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QPainter) SetBrush2(style core.Qt__BrushStyle) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::setBrush")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_SetBrush2(ptr.Pointer(), C.int(style))
	}
}

func (ptr *QPainter) SetBrush(brush QBrush_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::setBrush")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_SetBrush(ptr.Pointer(), PointerFromQBrush(brush))
	}
}

func (ptr *QPainter) SetBrushOrigin2(position core.QPoint_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::setBrushOrigin")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_SetBrushOrigin2(ptr.Pointer(), core.PointerFromQPoint(position))
	}
}

func (ptr *QPainter) SetBrushOrigin3(x int, y int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::setBrushOrigin")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_SetBrushOrigin3(ptr.Pointer(), C.int(x), C.int(y))
	}
}

func (ptr *QPainter) SetClipRect2(x int, y int, width int, height int, operation core.Qt__ClipOperation) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::setClipRect")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_SetClipRect2(ptr.Pointer(), C.int(x), C.int(y), C.int(width), C.int(height), C.int(operation))
	}
}

func (ptr *QPainter) SetClipping(enable bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::setClipping")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_SetClipping(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QPainter) SetCompositionMode(mode QPainter__CompositionMode) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::setCompositionMode")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_SetCompositionMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QPainter) SetFont(font QFont_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::setFont")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_SetFont(ptr.Pointer(), PointerFromQFont(font))
	}
}

func (ptr *QPainter) SetLayoutDirection(direction core.Qt__LayoutDirection) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::setLayoutDirection")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_SetLayoutDirection(ptr.Pointer(), C.int(direction))
	}
}

func (ptr *QPainter) SetOpacity(opacity float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::setOpacity")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_SetOpacity(ptr.Pointer(), C.double(opacity))
	}
}

func (ptr *QPainter) SetPen3(style core.Qt__PenStyle) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::setPen")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_SetPen3(ptr.Pointer(), C.int(style))
	}
}

func (ptr *QPainter) SetPen2(color QColor_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::setPen")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_SetPen2(ptr.Pointer(), PointerFromQColor(color))
	}
}

func (ptr *QPainter) SetPen(pen QPen_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::setPen")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_SetPen(ptr.Pointer(), PointerFromQPen(pen))
	}
}

func (ptr *QPainter) SetRenderHint(hint QPainter__RenderHint, on bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::setRenderHint")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_SetRenderHint(ptr.Pointer(), C.int(hint), C.int(qt.GoBoolToInt(on)))
	}
}

func (ptr *QPainter) SetRenderHints(hints QPainter__RenderHint, on bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::setRenderHints")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_SetRenderHints(ptr.Pointer(), C.int(hints), C.int(qt.GoBoolToInt(on)))
	}
}

func (ptr *QPainter) SetTransform(transform QTransform_ITF, combine bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::setTransform")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_SetTransform(ptr.Pointer(), PointerFromQTransform(transform), C.int(qt.GoBoolToInt(combine)))
	}
}

func (ptr *QPainter) SetViewTransformEnabled(enable bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::setViewTransformEnabled")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_SetViewTransformEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QPainter) SetViewport2(x int, y int, width int, height int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::setViewport")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_SetViewport2(ptr.Pointer(), C.int(x), C.int(y), C.int(width), C.int(height))
	}
}

func (ptr *QPainter) SetWindow2(x int, y int, width int, height int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::setWindow")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_SetWindow2(ptr.Pointer(), C.int(x), C.int(y), C.int(width), C.int(height))
	}
}

func (ptr *QPainter) SetWorldMatrixEnabled(enable bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::setWorldMatrixEnabled")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_SetWorldMatrixEnabled(ptr.Pointer(), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QPainter) SetWorldTransform(matrix QTransform_ITF, combine bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::setWorldTransform")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_SetWorldTransform(ptr.Pointer(), PointerFromQTransform(matrix), C.int(qt.GoBoolToInt(combine)))
	}
}

func (ptr *QPainter) Shear(sh float64, sv float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::shear")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_Shear(ptr.Pointer(), C.double(sh), C.double(sv))
	}
}

func (ptr *QPainter) StrokePath(path QPainterPath_ITF, pen QPen_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::strokePath")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_StrokePath(ptr.Pointer(), PointerFromQPainterPath(path), PointerFromQPen(pen))
	}
}

func (ptr *QPainter) TestRenderHint(hint QPainter__RenderHint) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::testRenderHint")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QPainter_TestRenderHint(ptr.Pointer(), C.int(hint)) != 0
	}
	return false
}

func (ptr *QPainter) Translate2(offset core.QPoint_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::translate")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_Translate2(ptr.Pointer(), core.PointerFromQPoint(offset))
	}
}

func (ptr *QPainter) Translate(offset core.QPointF_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::translate")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_Translate(ptr.Pointer(), core.PointerFromQPointF(offset))
	}
}

func (ptr *QPainter) Translate3(dx float64, dy float64) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::translate")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_Translate3(ptr.Pointer(), C.double(dx), C.double(dy))
	}
}

func (ptr *QPainter) ViewTransformEnabled() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::viewTransformEnabled")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QPainter_ViewTransformEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPainter) WorldMatrixEnabled() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::worldMatrixEnabled")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QPainter_WorldMatrixEnabled(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QPainter) DestroyQPainter() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QPainter::~QPainter")
		}
	}()

	if ptr.Pointer() != nil {
		C.QPainter_DestroyQPainter(ptr.Pointer())
	}
}
