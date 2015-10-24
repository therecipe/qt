package gui

//#include "qpainter.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QPainter struct {
	ptr unsafe.Pointer
}

type QPainterITF interface {
	QPainterPTR() *QPainter
}

func (p *QPainter) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QPainter) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQPainter(ptr QPainterITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPainterPTR().Pointer()
	}
	return nil
}

func QPainterFromPointer(ptr unsafe.Pointer) *QPainter {
	var n = new(QPainter)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPainter) QPainterPTR() *QPainter {
	return ptr
}

//QPainter::CompositionMode
type QPainter__CompositionMode int

var (
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
type QPainter__PixmapFragmentHint int

var (
	QPainter__OpaqueHint = QPainter__PixmapFragmentHint(0x01)
)

//QPainter::RenderHint
type QPainter__RenderHint int

var (
	QPainter__Antialiasing            = QPainter__RenderHint(0x01)
	QPainter__TextAntialiasing        = QPainter__RenderHint(0x02)
	QPainter__SmoothPixmapTransform   = QPainter__RenderHint(0x04)
	QPainter__HighQualityAntialiasing = QPainter__RenderHint(0x08)
	QPainter__NonCosmeticDefaultPen   = QPainter__RenderHint(0x10)
	QPainter__Qt4CompatiblePainting   = QPainter__RenderHint(0x20)
)

func NewQPainter2(device QPaintDeviceITF) *QPainter {
	return QPainterFromPointer(unsafe.Pointer(C.QPainter_NewQPainter2(C.QtObjectPtr(PointerFromQPaintDevice(device)))))
}

func (ptr *QPainter) Begin(device QPaintDeviceITF) bool {
	if ptr.Pointer() != nil {
		return C.QPainter_Begin(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPaintDevice(device))) != 0
	}
	return false
}

func (ptr *QPainter) DrawArc(rectangle core.QRectFITF, startAngle int, spanAngle int) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawArc(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRectF(rectangle)), C.int(startAngle), C.int(spanAngle))
	}
}

func (ptr *QPainter) DrawChord(rectangle core.QRectFITF, startAngle int, spanAngle int) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawChord(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRectF(rectangle)), C.int(startAngle), C.int(spanAngle))
	}
}

func (ptr *QPainter) DrawConvexPolygon2(points core.QPointITF, pointCount int) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawConvexPolygon2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPoint(points)), C.int(pointCount))
	}
}

func (ptr *QPainter) DrawConvexPolygon(points core.QPointFITF, pointCount int) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawConvexPolygon(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPointF(points)), C.int(pointCount))
	}
}

func (ptr *QPainter) DrawEllipse2(rectangle core.QRectITF) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawEllipse2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRect(rectangle)))
	}
}

func (ptr *QPainter) DrawEllipse(rectangle core.QRectFITF) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawEllipse(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRectF(rectangle)))
	}
}

func (ptr *QPainter) DrawGlyphRun(position core.QPointFITF, glyphs QGlyphRunITF) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawGlyphRun(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPointF(position)), C.QtObjectPtr(PointerFromQGlyphRun(glyphs)))
	}
}

func (ptr *QPainter) DrawImage3(point core.QPointFITF, image QImageITF) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawImage3(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPointF(point)), C.QtObjectPtr(PointerFromQImage(image)))
	}
}

func (ptr *QPainter) DrawImage(target core.QRectFITF, image QImageITF, source core.QRectFITF, flags core.Qt__ImageConversionFlag) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawImage(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRectF(target)), C.QtObjectPtr(PointerFromQImage(image)), C.QtObjectPtr(core.PointerFromQRectF(source)), C.int(flags))
	}
}

func (ptr *QPainter) DrawLines2(lines core.QLineITF, lineCount int) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawLines2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQLine(lines)), C.int(lineCount))
	}
}

func (ptr *QPainter) DrawPicture(point core.QPointFITF, picture QPictureITF) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawPicture(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPointF(point)), C.QtObjectPtr(PointerFromQPicture(picture)))
	}
}

func (ptr *QPainter) DrawPie(rectangle core.QRectFITF, startAngle int, spanAngle int) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawPie(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRectF(rectangle)), C.int(startAngle), C.int(spanAngle))
	}
}

func (ptr *QPainter) DrawPixmap5(point core.QPointFITF, pixmap QPixmapITF) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawPixmap5(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPointF(point)), C.QtObjectPtr(PointerFromQPixmap(pixmap)))
	}
}

func (ptr *QPainter) DrawPixmap(target core.QRectFITF, pixmap QPixmapITF, source core.QRectFITF) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawPixmap(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRectF(target)), C.QtObjectPtr(PointerFromQPixmap(pixmap)), C.QtObjectPtr(core.PointerFromQRectF(source)))
	}
}

func (ptr *QPainter) DrawRects2(rectangles core.QRectITF, rectCount int) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawRects2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRect(rectangles)), C.int(rectCount))
	}
}

func (ptr *QPainter) DrawRects(rectangles core.QRectFITF, rectCount int) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawRects(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRectF(rectangles)), C.int(rectCount))
	}
}

func (ptr *QPainter) DrawText(position core.QPointFITF, text string) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawText(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPointF(position)), C.CString(text))
	}
}

func (ptr *QPainter) DrawText5(rectangle core.QRectITF, flags int, text string, boundingRect core.QRectITF) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawText5(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRect(rectangle)), C.int(flags), C.CString(text), C.QtObjectPtr(core.PointerFromQRect(boundingRect)))
	}
}

func (ptr *QPainter) DrawText8(rectangle core.QRectFITF, text string, option QTextOptionITF) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawText8(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRectF(rectangle)), C.CString(text), C.QtObjectPtr(PointerFromQTextOption(option)))
	}
}

func (ptr *QPainter) DrawText4(rectangle core.QRectFITF, flags int, text string, boundingRect core.QRectFITF) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawText4(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRectF(rectangle)), C.int(flags), C.CString(text), C.QtObjectPtr(core.PointerFromQRectF(boundingRect)))
	}
}

func (ptr *QPainter) DrawTiledPixmap(rectangle core.QRectFITF, pixmap QPixmapITF, position core.QPointFITF) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawTiledPixmap(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRectF(rectangle)), C.QtObjectPtr(PointerFromQPixmap(pixmap)), C.QtObjectPtr(core.PointerFromQPointF(position)))
	}
}

func (ptr *QPainter) EraseRect(rectangle core.QRectFITF) {
	if ptr.Pointer() != nil {
		C.QPainter_EraseRect(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRectF(rectangle)))
	}
}

func (ptr *QPainter) FillRect5(rectangle core.QRectITF, brush QBrushITF) {
	if ptr.Pointer() != nil {
		C.QPainter_FillRect5(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRect(rectangle)), C.QtObjectPtr(PointerFromQBrush(brush)))
	}
}

func (ptr *QPainter) FillRect6(rectangle core.QRectITF, color QColorITF) {
	if ptr.Pointer() != nil {
		C.QPainter_FillRect6(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRect(rectangle)), C.QtObjectPtr(PointerFromQColor(color)))
	}
}

func (ptr *QPainter) FillRect(rectangle core.QRectFITF, brush QBrushITF) {
	if ptr.Pointer() != nil {
		C.QPainter_FillRect(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRectF(rectangle)), C.QtObjectPtr(PointerFromQBrush(brush)))
	}
}

func (ptr *QPainter) FillRect7(rectangle core.QRectFITF, color QColorITF) {
	if ptr.Pointer() != nil {
		C.QPainter_FillRect7(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRectF(rectangle)), C.QtObjectPtr(PointerFromQColor(color)))
	}
}

func (ptr *QPainter) SetBackground(brush QBrushITF) {
	if ptr.Pointer() != nil {
		C.QPainter_SetBackground(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQBrush(brush)))
	}
}

func (ptr *QPainter) SetBrushOrigin(position core.QPointFITF) {
	if ptr.Pointer() != nil {
		C.QPainter_SetBrushOrigin(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPointF(position)))
	}
}

func (ptr *QPainter) SetClipPath(path QPainterPathITF, operation core.Qt__ClipOperation) {
	if ptr.Pointer() != nil {
		C.QPainter_SetClipPath(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPainterPath(path)), C.int(operation))
	}
}

func (ptr *QPainter) SetClipRect3(rectangle core.QRectITF, operation core.Qt__ClipOperation) {
	if ptr.Pointer() != nil {
		C.QPainter_SetClipRect3(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRect(rectangle)), C.int(operation))
	}
}

func (ptr *QPainter) SetClipRect(rectangle core.QRectFITF, operation core.Qt__ClipOperation) {
	if ptr.Pointer() != nil {
		C.QPainter_SetClipRect(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRectF(rectangle)), C.int(operation))
	}
}

func (ptr *QPainter) SetClipRegion(region QRegionITF, operation core.Qt__ClipOperation) {
	if ptr.Pointer() != nil {
		C.QPainter_SetClipRegion(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQRegion(region)), C.int(operation))
	}
}

func (ptr *QPainter) SetViewport(rectangle core.QRectITF) {
	if ptr.Pointer() != nil {
		C.QPainter_SetViewport(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRect(rectangle)))
	}
}

func (ptr *QPainter) SetWindow(rectangle core.QRectITF) {
	if ptr.Pointer() != nil {
		C.QPainter_SetWindow(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRect(rectangle)))
	}
}

func NewQPainter() *QPainter {
	return QPainterFromPointer(unsafe.Pointer(C.QPainter_NewQPainter()))
}

func (ptr *QPainter) BackgroundMode() core.Qt__BGMode {
	if ptr.Pointer() != nil {
		return core.Qt__BGMode(C.QPainter_BackgroundMode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QPainter) BeginNativePainting() {
	if ptr.Pointer() != nil {
		C.QPainter_BeginNativePainting(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QPainter) CompositionMode() QPainter__CompositionMode {
	if ptr.Pointer() != nil {
		return QPainter__CompositionMode(C.QPainter_CompositionMode(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QPainter) Device() *QPaintDevice {
	if ptr.Pointer() != nil {
		return QPaintDeviceFromPointer(unsafe.Pointer(C.QPainter_Device(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QPainter) DrawArc2(rectangle core.QRectITF, startAngle int, spanAngle int) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawArc2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRect(rectangle)), C.int(startAngle), C.int(spanAngle))
	}
}

func (ptr *QPainter) DrawArc3(x int, y int, width int, height int, startAngle int, spanAngle int) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawArc3(C.QtObjectPtr(ptr.Pointer()), C.int(x), C.int(y), C.int(width), C.int(height), C.int(startAngle), C.int(spanAngle))
	}
}

func (ptr *QPainter) DrawChord2(rectangle core.QRectITF, startAngle int, spanAngle int) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawChord2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRect(rectangle)), C.int(startAngle), C.int(spanAngle))
	}
}

func (ptr *QPainter) DrawChord3(x int, y int, width int, height int, startAngle int, spanAngle int) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawChord3(C.QtObjectPtr(ptr.Pointer()), C.int(x), C.int(y), C.int(width), C.int(height), C.int(startAngle), C.int(spanAngle))
	}
}

func (ptr *QPainter) DrawConvexPolygon4(polygon QPolygonITF) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawConvexPolygon4(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPolygon(polygon)))
	}
}

func (ptr *QPainter) DrawConvexPolygon3(polygon QPolygonFITF) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawConvexPolygon3(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPolygonF(polygon)))
	}
}

func (ptr *QPainter) DrawEllipse5(center core.QPointITF, rx int, ry int) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawEllipse5(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPoint(center)), C.int(rx), C.int(ry))
	}
}

func (ptr *QPainter) DrawEllipse3(x int, y int, width int, height int) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawEllipse3(C.QtObjectPtr(ptr.Pointer()), C.int(x), C.int(y), C.int(width), C.int(height))
	}
}

func (ptr *QPainter) DrawImage4(point core.QPointITF, image QImageITF) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawImage4(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPoint(point)), C.QtObjectPtr(PointerFromQImage(image)))
	}
}

func (ptr *QPainter) DrawImage6(point core.QPointITF, image QImageITF, source core.QRectITF, flags core.Qt__ImageConversionFlag) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawImage6(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPoint(point)), C.QtObjectPtr(PointerFromQImage(image)), C.QtObjectPtr(core.PointerFromQRect(source)), C.int(flags))
	}
}

func (ptr *QPainter) DrawImage5(point core.QPointFITF, image QImageITF, source core.QRectFITF, flags core.Qt__ImageConversionFlag) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawImage5(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPointF(point)), C.QtObjectPtr(PointerFromQImage(image)), C.QtObjectPtr(core.PointerFromQRectF(source)), C.int(flags))
	}
}

func (ptr *QPainter) DrawImage8(rectangle core.QRectITF, image QImageITF) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawImage8(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRect(rectangle)), C.QtObjectPtr(PointerFromQImage(image)))
	}
}

func (ptr *QPainter) DrawImage2(target core.QRectITF, image QImageITF, source core.QRectITF, flags core.Qt__ImageConversionFlag) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawImage2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRect(target)), C.QtObjectPtr(PointerFromQImage(image)), C.QtObjectPtr(core.PointerFromQRect(source)), C.int(flags))
	}
}

func (ptr *QPainter) DrawImage7(rectangle core.QRectFITF, image QImageITF) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawImage7(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRectF(rectangle)), C.QtObjectPtr(PointerFromQImage(image)))
	}
}

func (ptr *QPainter) DrawImage9(x int, y int, image QImageITF, sx int, sy int, sw int, sh int, flags core.Qt__ImageConversionFlag) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawImage9(C.QtObjectPtr(ptr.Pointer()), C.int(x), C.int(y), C.QtObjectPtr(PointerFromQImage(image)), C.int(sx), C.int(sy), C.int(sw), C.int(sh), C.int(flags))
	}
}

func (ptr *QPainter) DrawLine2(line core.QLineITF) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawLine2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQLine(line)))
	}
}

func (ptr *QPainter) DrawLine(line core.QLineFITF) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawLine(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQLineF(line)))
	}
}

func (ptr *QPainter) DrawLine3(p1 core.QPointITF, p2 core.QPointITF) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawLine3(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPoint(p1)), C.QtObjectPtr(core.PointerFromQPoint(p2)))
	}
}

func (ptr *QPainter) DrawLine4(p1 core.QPointFITF, p2 core.QPointFITF) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawLine4(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPointF(p1)), C.QtObjectPtr(core.PointerFromQPointF(p2)))
	}
}

func (ptr *QPainter) DrawLine5(x1 int, y1 int, x2 int, y2 int) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawLine5(C.QtObjectPtr(ptr.Pointer()), C.int(x1), C.int(y1), C.int(x2), C.int(y2))
	}
}

func (ptr *QPainter) DrawLines(lines core.QLineFITF, lineCount int) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawLines(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQLineF(lines)), C.int(lineCount))
	}
}

func (ptr *QPainter) DrawLines4(pointPairs core.QPointITF, lineCount int) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawLines4(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPoint(pointPairs)), C.int(lineCount))
	}
}

func (ptr *QPainter) DrawLines3(pointPairs core.QPointFITF, lineCount int) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawLines3(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPointF(pointPairs)), C.int(lineCount))
	}
}

func (ptr *QPainter) DrawPath(path QPainterPathITF) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawPath(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPainterPath(path)))
	}
}

func (ptr *QPainter) DrawPicture2(point core.QPointITF, picture QPictureITF) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawPicture2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPoint(point)), C.QtObjectPtr(PointerFromQPicture(picture)))
	}
}

func (ptr *QPainter) DrawPicture3(x int, y int, picture QPictureITF) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawPicture3(C.QtObjectPtr(ptr.Pointer()), C.int(x), C.int(y), C.QtObjectPtr(PointerFromQPicture(picture)))
	}
}

func (ptr *QPainter) DrawPie2(rectangle core.QRectITF, startAngle int, spanAngle int) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawPie2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRect(rectangle)), C.int(startAngle), C.int(spanAngle))
	}
}

func (ptr *QPainter) DrawPie3(x int, y int, width int, height int, startAngle int, spanAngle int) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawPie3(C.QtObjectPtr(ptr.Pointer()), C.int(x), C.int(y), C.int(width), C.int(height), C.int(startAngle), C.int(spanAngle))
	}
}

func (ptr *QPainter) DrawPixmap6(point core.QPointITF, pixmap QPixmapITF) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawPixmap6(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPoint(point)), C.QtObjectPtr(PointerFromQPixmap(pixmap)))
	}
}

func (ptr *QPainter) DrawPixmap4(point core.QPointITF, pixmap QPixmapITF, source core.QRectITF) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawPixmap4(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPoint(point)), C.QtObjectPtr(PointerFromQPixmap(pixmap)), C.QtObjectPtr(core.PointerFromQRect(source)))
	}
}

func (ptr *QPainter) DrawPixmap3(point core.QPointFITF, pixmap QPixmapITF, source core.QRectFITF) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawPixmap3(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPointF(point)), C.QtObjectPtr(PointerFromQPixmap(pixmap)), C.QtObjectPtr(core.PointerFromQRectF(source)))
	}
}

func (ptr *QPainter) DrawPixmap8(rectangle core.QRectITF, pixmap QPixmapITF) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawPixmap8(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRect(rectangle)), C.QtObjectPtr(PointerFromQPixmap(pixmap)))
	}
}

func (ptr *QPainter) DrawPixmap2(target core.QRectITF, pixmap QPixmapITF, source core.QRectITF) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawPixmap2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRect(target)), C.QtObjectPtr(PointerFromQPixmap(pixmap)), C.QtObjectPtr(core.PointerFromQRect(source)))
	}
}

func (ptr *QPainter) DrawPixmap7(x int, y int, pixmap QPixmapITF) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawPixmap7(C.QtObjectPtr(ptr.Pointer()), C.int(x), C.int(y), C.QtObjectPtr(PointerFromQPixmap(pixmap)))
	}
}

func (ptr *QPainter) DrawPixmap11(x int, y int, pixmap QPixmapITF, sx int, sy int, sw int, sh int) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawPixmap11(C.QtObjectPtr(ptr.Pointer()), C.int(x), C.int(y), C.QtObjectPtr(PointerFromQPixmap(pixmap)), C.int(sx), C.int(sy), C.int(sw), C.int(sh))
	}
}

func (ptr *QPainter) DrawPixmap10(x int, y int, w int, h int, pixmap QPixmapITF, sx int, sy int, sw int, sh int) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawPixmap10(C.QtObjectPtr(ptr.Pointer()), C.int(x), C.int(y), C.int(w), C.int(h), C.QtObjectPtr(PointerFromQPixmap(pixmap)), C.int(sx), C.int(sy), C.int(sw), C.int(sh))
	}
}

func (ptr *QPainter) DrawPixmap9(x int, y int, width int, height int, pixmap QPixmapITF) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawPixmap9(C.QtObjectPtr(ptr.Pointer()), C.int(x), C.int(y), C.int(width), C.int(height), C.QtObjectPtr(PointerFromQPixmap(pixmap)))
	}
}

func (ptr *QPainter) DrawPoint2(position core.QPointITF) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawPoint2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPoint(position)))
	}
}

func (ptr *QPainter) DrawPoint(position core.QPointFITF) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawPoint(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPointF(position)))
	}
}

func (ptr *QPainter) DrawPoint3(x int, y int) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawPoint3(C.QtObjectPtr(ptr.Pointer()), C.int(x), C.int(y))
	}
}

func (ptr *QPainter) DrawPoints2(points core.QPointITF, pointCount int) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawPoints2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPoint(points)), C.int(pointCount))
	}
}

func (ptr *QPainter) DrawPoints(points core.QPointFITF, pointCount int) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawPoints(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPointF(points)), C.int(pointCount))
	}
}

func (ptr *QPainter) DrawPoints4(points QPolygonITF) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawPoints4(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPolygon(points)))
	}
}

func (ptr *QPainter) DrawPoints3(points QPolygonFITF) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawPoints3(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPolygonF(points)))
	}
}

func (ptr *QPainter) DrawPolygon2(points core.QPointITF, pointCount int, fillRule core.Qt__FillRule) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawPolygon2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPoint(points)), C.int(pointCount), C.int(fillRule))
	}
}

func (ptr *QPainter) DrawPolygon(points core.QPointFITF, pointCount int, fillRule core.Qt__FillRule) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawPolygon(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPointF(points)), C.int(pointCount), C.int(fillRule))
	}
}

func (ptr *QPainter) DrawPolygon4(points QPolygonITF, fillRule core.Qt__FillRule) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawPolygon4(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPolygon(points)), C.int(fillRule))
	}
}

func (ptr *QPainter) DrawPolygon3(points QPolygonFITF, fillRule core.Qt__FillRule) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawPolygon3(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPolygonF(points)), C.int(fillRule))
	}
}

func (ptr *QPainter) DrawPolyline2(points core.QPointITF, pointCount int) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawPolyline2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPoint(points)), C.int(pointCount))
	}
}

func (ptr *QPainter) DrawPolyline(points core.QPointFITF, pointCount int) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawPolyline(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPointF(points)), C.int(pointCount))
	}
}

func (ptr *QPainter) DrawPolyline4(points QPolygonITF) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawPolyline4(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPolygon(points)))
	}
}

func (ptr *QPainter) DrawPolyline3(points QPolygonFITF) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawPolyline3(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPolygonF(points)))
	}
}

func (ptr *QPainter) DrawRect2(rectangle core.QRectITF) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawRect2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRect(rectangle)))
	}
}

func (ptr *QPainter) DrawRect(rectangle core.QRectFITF) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawRect(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRectF(rectangle)))
	}
}

func (ptr *QPainter) DrawRect3(x int, y int, width int, height int) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawRect3(C.QtObjectPtr(ptr.Pointer()), C.int(x), C.int(y), C.int(width), C.int(height))
	}
}

func (ptr *QPainter) DrawStaticText2(topLeftPosition core.QPointITF, staticText QStaticTextITF) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawStaticText2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPoint(topLeftPosition)), C.QtObjectPtr(PointerFromQStaticText(staticText)))
	}
}

func (ptr *QPainter) DrawStaticText(topLeftPosition core.QPointFITF, staticText QStaticTextITF) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawStaticText(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPointF(topLeftPosition)), C.QtObjectPtr(PointerFromQStaticText(staticText)))
	}
}

func (ptr *QPainter) DrawStaticText3(left int, top int, staticText QStaticTextITF) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawStaticText3(C.QtObjectPtr(ptr.Pointer()), C.int(left), C.int(top), C.QtObjectPtr(PointerFromQStaticText(staticText)))
	}
}

func (ptr *QPainter) DrawText3(position core.QPointITF, text string) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawText3(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPoint(position)), C.CString(text))
	}
}

func (ptr *QPainter) DrawText6(x int, y int, text string) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawText6(C.QtObjectPtr(ptr.Pointer()), C.int(x), C.int(y), C.CString(text))
	}
}

func (ptr *QPainter) DrawText7(x int, y int, width int, height int, flags int, text string, boundingRect core.QRectITF) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawText7(C.QtObjectPtr(ptr.Pointer()), C.int(x), C.int(y), C.int(width), C.int(height), C.int(flags), C.CString(text), C.QtObjectPtr(core.PointerFromQRect(boundingRect)))
	}
}

func (ptr *QPainter) DrawTiledPixmap2(rectangle core.QRectITF, pixmap QPixmapITF, position core.QPointITF) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawTiledPixmap2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRect(rectangle)), C.QtObjectPtr(PointerFromQPixmap(pixmap)), C.QtObjectPtr(core.PointerFromQPoint(position)))
	}
}

func (ptr *QPainter) DrawTiledPixmap3(x int, y int, width int, height int, pixmap QPixmapITF, sx int, sy int) {
	if ptr.Pointer() != nil {
		C.QPainter_DrawTiledPixmap3(C.QtObjectPtr(ptr.Pointer()), C.int(x), C.int(y), C.int(width), C.int(height), C.QtObjectPtr(PointerFromQPixmap(pixmap)), C.int(sx), C.int(sy))
	}
}

func (ptr *QPainter) End() bool {
	if ptr.Pointer() != nil {
		return C.QPainter_End(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QPainter) EndNativePainting() {
	if ptr.Pointer() != nil {
		C.QPainter_EndNativePainting(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QPainter) EraseRect2(rectangle core.QRectITF) {
	if ptr.Pointer() != nil {
		C.QPainter_EraseRect2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRect(rectangle)))
	}
}

func (ptr *QPainter) EraseRect3(x int, y int, width int, height int) {
	if ptr.Pointer() != nil {
		C.QPainter_EraseRect3(C.QtObjectPtr(ptr.Pointer()), C.int(x), C.int(y), C.int(width), C.int(height))
	}
}

func (ptr *QPainter) FillPath(path QPainterPathITF, brush QBrushITF) {
	if ptr.Pointer() != nil {
		C.QPainter_FillPath(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPainterPath(path)), C.QtObjectPtr(PointerFromQBrush(brush)))
	}
}

func (ptr *QPainter) FillRect3(rectangle core.QRectITF, style core.Qt__BrushStyle) {
	if ptr.Pointer() != nil {
		C.QPainter_FillRect3(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRect(rectangle)), C.int(style))
	}
}

func (ptr *QPainter) FillRect11(rectangle core.QRectITF, color core.Qt__GlobalColor) {
	if ptr.Pointer() != nil {
		C.QPainter_FillRect11(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRect(rectangle)), C.int(color))
	}
}

func (ptr *QPainter) FillRect4(rectangle core.QRectFITF, style core.Qt__BrushStyle) {
	if ptr.Pointer() != nil {
		C.QPainter_FillRect4(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRectF(rectangle)), C.int(style))
	}
}

func (ptr *QPainter) FillRect12(rectangle core.QRectFITF, color core.Qt__GlobalColor) {
	if ptr.Pointer() != nil {
		C.QPainter_FillRect12(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQRectF(rectangle)), C.int(color))
	}
}

func (ptr *QPainter) FillRect2(x int, y int, width int, height int, style core.Qt__BrushStyle) {
	if ptr.Pointer() != nil {
		C.QPainter_FillRect2(C.QtObjectPtr(ptr.Pointer()), C.int(x), C.int(y), C.int(width), C.int(height), C.int(style))
	}
}

func (ptr *QPainter) FillRect10(x int, y int, width int, height int, color core.Qt__GlobalColor) {
	if ptr.Pointer() != nil {
		C.QPainter_FillRect10(C.QtObjectPtr(ptr.Pointer()), C.int(x), C.int(y), C.int(width), C.int(height), C.int(color))
	}
}

func (ptr *QPainter) FillRect8(x int, y int, width int, height int, brush QBrushITF) {
	if ptr.Pointer() != nil {
		C.QPainter_FillRect8(C.QtObjectPtr(ptr.Pointer()), C.int(x), C.int(y), C.int(width), C.int(height), C.QtObjectPtr(PointerFromQBrush(brush)))
	}
}

func (ptr *QPainter) FillRect9(x int, y int, width int, height int, color QColorITF) {
	if ptr.Pointer() != nil {
		C.QPainter_FillRect9(C.QtObjectPtr(ptr.Pointer()), C.int(x), C.int(y), C.int(width), C.int(height), C.QtObjectPtr(PointerFromQColor(color)))
	}
}

func (ptr *QPainter) HasClipping() bool {
	if ptr.Pointer() != nil {
		return C.QPainter_HasClipping(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QPainter) IsActive() bool {
	if ptr.Pointer() != nil {
		return C.QPainter_IsActive(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QPainter) LayoutDirection() core.Qt__LayoutDirection {
	if ptr.Pointer() != nil {
		return core.Qt__LayoutDirection(C.QPainter_LayoutDirection(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QPainter) PaintEngine() *QPaintEngine {
	if ptr.Pointer() != nil {
		return QPaintEngineFromPointer(unsafe.Pointer(C.QPainter_PaintEngine(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QPainter) RenderHints() QPainter__RenderHint {
	if ptr.Pointer() != nil {
		return QPainter__RenderHint(C.QPainter_RenderHints(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QPainter) ResetTransform() {
	if ptr.Pointer() != nil {
		C.QPainter_ResetTransform(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QPainter) Restore() {
	if ptr.Pointer() != nil {
		C.QPainter_Restore(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QPainter) Save() {
	if ptr.Pointer() != nil {
		C.QPainter_Save(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QPainter) SetBackgroundMode(mode core.Qt__BGMode) {
	if ptr.Pointer() != nil {
		C.QPainter_SetBackgroundMode(C.QtObjectPtr(ptr.Pointer()), C.int(mode))
	}
}

func (ptr *QPainter) SetBrush2(style core.Qt__BrushStyle) {
	if ptr.Pointer() != nil {
		C.QPainter_SetBrush2(C.QtObjectPtr(ptr.Pointer()), C.int(style))
	}
}

func (ptr *QPainter) SetBrush(brush QBrushITF) {
	if ptr.Pointer() != nil {
		C.QPainter_SetBrush(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQBrush(brush)))
	}
}

func (ptr *QPainter) SetBrushOrigin2(position core.QPointITF) {
	if ptr.Pointer() != nil {
		C.QPainter_SetBrushOrigin2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPoint(position)))
	}
}

func (ptr *QPainter) SetBrushOrigin3(x int, y int) {
	if ptr.Pointer() != nil {
		C.QPainter_SetBrushOrigin3(C.QtObjectPtr(ptr.Pointer()), C.int(x), C.int(y))
	}
}

func (ptr *QPainter) SetClipRect2(x int, y int, width int, height int, operation core.Qt__ClipOperation) {
	if ptr.Pointer() != nil {
		C.QPainter_SetClipRect2(C.QtObjectPtr(ptr.Pointer()), C.int(x), C.int(y), C.int(width), C.int(height), C.int(operation))
	}
}

func (ptr *QPainter) SetClipping(enable bool) {
	if ptr.Pointer() != nil {
		C.QPainter_SetClipping(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QPainter) SetCompositionMode(mode QPainter__CompositionMode) {
	if ptr.Pointer() != nil {
		C.QPainter_SetCompositionMode(C.QtObjectPtr(ptr.Pointer()), C.int(mode))
	}
}

func (ptr *QPainter) SetFont(font QFontITF) {
	if ptr.Pointer() != nil {
		C.QPainter_SetFont(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQFont(font)))
	}
}

func (ptr *QPainter) SetLayoutDirection(direction core.Qt__LayoutDirection) {
	if ptr.Pointer() != nil {
		C.QPainter_SetLayoutDirection(C.QtObjectPtr(ptr.Pointer()), C.int(direction))
	}
}

func (ptr *QPainter) SetPen3(style core.Qt__PenStyle) {
	if ptr.Pointer() != nil {
		C.QPainter_SetPen3(C.QtObjectPtr(ptr.Pointer()), C.int(style))
	}
}

func (ptr *QPainter) SetPen2(color QColorITF) {
	if ptr.Pointer() != nil {
		C.QPainter_SetPen2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQColor(color)))
	}
}

func (ptr *QPainter) SetPen(pen QPenITF) {
	if ptr.Pointer() != nil {
		C.QPainter_SetPen(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPen(pen)))
	}
}

func (ptr *QPainter) SetRenderHint(hint QPainter__RenderHint, on bool) {
	if ptr.Pointer() != nil {
		C.QPainter_SetRenderHint(C.QtObjectPtr(ptr.Pointer()), C.int(hint), C.int(qt.GoBoolToInt(on)))
	}
}

func (ptr *QPainter) SetRenderHints(hints QPainter__RenderHint, on bool) {
	if ptr.Pointer() != nil {
		C.QPainter_SetRenderHints(C.QtObjectPtr(ptr.Pointer()), C.int(hints), C.int(qt.GoBoolToInt(on)))
	}
}

func (ptr *QPainter) SetTransform(transform QTransformITF, combine bool) {
	if ptr.Pointer() != nil {
		C.QPainter_SetTransform(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTransform(transform)), C.int(qt.GoBoolToInt(combine)))
	}
}

func (ptr *QPainter) SetViewTransformEnabled(enable bool) {
	if ptr.Pointer() != nil {
		C.QPainter_SetViewTransformEnabled(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QPainter) SetViewport2(x int, y int, width int, height int) {
	if ptr.Pointer() != nil {
		C.QPainter_SetViewport2(C.QtObjectPtr(ptr.Pointer()), C.int(x), C.int(y), C.int(width), C.int(height))
	}
}

func (ptr *QPainter) SetWindow2(x int, y int, width int, height int) {
	if ptr.Pointer() != nil {
		C.QPainter_SetWindow2(C.QtObjectPtr(ptr.Pointer()), C.int(x), C.int(y), C.int(width), C.int(height))
	}
}

func (ptr *QPainter) SetWorldMatrixEnabled(enable bool) {
	if ptr.Pointer() != nil {
		C.QPainter_SetWorldMatrixEnabled(C.QtObjectPtr(ptr.Pointer()), C.int(qt.GoBoolToInt(enable)))
	}
}

func (ptr *QPainter) SetWorldTransform(matrix QTransformITF, combine bool) {
	if ptr.Pointer() != nil {
		C.QPainter_SetWorldTransform(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQTransform(matrix)), C.int(qt.GoBoolToInt(combine)))
	}
}

func (ptr *QPainter) StrokePath(path QPainterPathITF, pen QPenITF) {
	if ptr.Pointer() != nil {
		C.QPainter_StrokePath(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQPainterPath(path)), C.QtObjectPtr(PointerFromQPen(pen)))
	}
}

func (ptr *QPainter) TestRenderHint(hint QPainter__RenderHint) bool {
	if ptr.Pointer() != nil {
		return C.QPainter_TestRenderHint(C.QtObjectPtr(ptr.Pointer()), C.int(hint)) != 0
	}
	return false
}

func (ptr *QPainter) Translate2(offset core.QPointITF) {
	if ptr.Pointer() != nil {
		C.QPainter_Translate2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPoint(offset)))
	}
}

func (ptr *QPainter) Translate(offset core.QPointFITF) {
	if ptr.Pointer() != nil {
		C.QPainter_Translate(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQPointF(offset)))
	}
}

func (ptr *QPainter) ViewTransformEnabled() bool {
	if ptr.Pointer() != nil {
		return C.QPainter_ViewTransformEnabled(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QPainter) WorldMatrixEnabled() bool {
	if ptr.Pointer() != nil {
		return C.QPainter_WorldMatrixEnabled(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QPainter) DestroyQPainter() {
	if ptr.Pointer() != nil {
		C.QPainter_DestroyQPainter(C.QtObjectPtr(ptr.Pointer()))
	}
}
