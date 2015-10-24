#include "qpainter.h"
#include <QPaintDevice>
#include <QRegion>
#include <QPixmap>
#include <QPen>
#include <QPoint>
#include <QPicture>
#include <QGlyphRun>
#include <QPolygonF>
#include <QLineF>
#include <QRectF>
#include <QVariant>
#include <QModelIndex>
#include <QPointF>
#include <QTextOption>
#include <QBrush>
#include <QColor>
#include <QTransform>
#include <QStaticText>
#include <QPolygon>
#include <QRect>
#include <QString>
#include <QUrl>
#include <QFont>
#include <QImage>
#include <QLine>
#include <QPainterPath>
#include <QPainter>
#include "_cgo_export.h"

class MyQPainter: public QPainter {
public:
};

QtObjectPtr QPainter_NewQPainter2(QtObjectPtr device){
	return new QPainter(static_cast<QPaintDevice*>(device));
}

int QPainter_Begin(QtObjectPtr ptr, QtObjectPtr device){
	return static_cast<QPainter*>(ptr)->begin(static_cast<QPaintDevice*>(device));
}

void QPainter_DrawArc(QtObjectPtr ptr, QtObjectPtr rectangle, int startAngle, int spanAngle){
	static_cast<QPainter*>(ptr)->drawArc(*static_cast<QRectF*>(rectangle), startAngle, spanAngle);
}

void QPainter_DrawChord(QtObjectPtr ptr, QtObjectPtr rectangle, int startAngle, int spanAngle){
	static_cast<QPainter*>(ptr)->drawChord(*static_cast<QRectF*>(rectangle), startAngle, spanAngle);
}

void QPainter_DrawConvexPolygon2(QtObjectPtr ptr, QtObjectPtr points, int pointCount){
	static_cast<QPainter*>(ptr)->drawConvexPolygon(static_cast<QPoint*>(points), pointCount);
}

void QPainter_DrawConvexPolygon(QtObjectPtr ptr, QtObjectPtr points, int pointCount){
	static_cast<QPainter*>(ptr)->drawConvexPolygon(static_cast<QPointF*>(points), pointCount);
}

void QPainter_DrawEllipse2(QtObjectPtr ptr, QtObjectPtr rectangle){
	static_cast<QPainter*>(ptr)->drawEllipse(*static_cast<QRect*>(rectangle));
}

void QPainter_DrawEllipse(QtObjectPtr ptr, QtObjectPtr rectangle){
	static_cast<QPainter*>(ptr)->drawEllipse(*static_cast<QRectF*>(rectangle));
}

void QPainter_DrawGlyphRun(QtObjectPtr ptr, QtObjectPtr position, QtObjectPtr glyphs){
	static_cast<QPainter*>(ptr)->drawGlyphRun(*static_cast<QPointF*>(position), *static_cast<QGlyphRun*>(glyphs));
}

void QPainter_DrawImage3(QtObjectPtr ptr, QtObjectPtr point, QtObjectPtr image){
	static_cast<QPainter*>(ptr)->drawImage(*static_cast<QPointF*>(point), *static_cast<QImage*>(image));
}

void QPainter_DrawImage(QtObjectPtr ptr, QtObjectPtr target, QtObjectPtr image, QtObjectPtr source, int flags){
	static_cast<QPainter*>(ptr)->drawImage(*static_cast<QRectF*>(target), *static_cast<QImage*>(image), *static_cast<QRectF*>(source), static_cast<Qt::ImageConversionFlag>(flags));
}

void QPainter_DrawLines2(QtObjectPtr ptr, QtObjectPtr lines, int lineCount){
	static_cast<QPainter*>(ptr)->drawLines(static_cast<QLine*>(lines), lineCount);
}

void QPainter_DrawPicture(QtObjectPtr ptr, QtObjectPtr point, QtObjectPtr picture){
	static_cast<QPainter*>(ptr)->drawPicture(*static_cast<QPointF*>(point), *static_cast<QPicture*>(picture));
}

void QPainter_DrawPie(QtObjectPtr ptr, QtObjectPtr rectangle, int startAngle, int spanAngle){
	static_cast<QPainter*>(ptr)->drawPie(*static_cast<QRectF*>(rectangle), startAngle, spanAngle);
}

void QPainter_DrawPixmap5(QtObjectPtr ptr, QtObjectPtr point, QtObjectPtr pixmap){
	static_cast<QPainter*>(ptr)->drawPixmap(*static_cast<QPointF*>(point), *static_cast<QPixmap*>(pixmap));
}

void QPainter_DrawPixmap(QtObjectPtr ptr, QtObjectPtr target, QtObjectPtr pixmap, QtObjectPtr source){
	static_cast<QPainter*>(ptr)->drawPixmap(*static_cast<QRectF*>(target), *static_cast<QPixmap*>(pixmap), *static_cast<QRectF*>(source));
}

void QPainter_DrawRects2(QtObjectPtr ptr, QtObjectPtr rectangles, int rectCount){
	static_cast<QPainter*>(ptr)->drawRects(static_cast<QRect*>(rectangles), rectCount);
}

void QPainter_DrawRects(QtObjectPtr ptr, QtObjectPtr rectangles, int rectCount){
	static_cast<QPainter*>(ptr)->drawRects(static_cast<QRectF*>(rectangles), rectCount);
}

void QPainter_DrawText(QtObjectPtr ptr, QtObjectPtr position, char* text){
	static_cast<QPainter*>(ptr)->drawText(*static_cast<QPointF*>(position), QString(text));
}

void QPainter_DrawText5(QtObjectPtr ptr, QtObjectPtr rectangle, int flags, char* text, QtObjectPtr boundingRect){
	static_cast<QPainter*>(ptr)->drawText(*static_cast<QRect*>(rectangle), flags, QString(text), static_cast<QRect*>(boundingRect));
}

void QPainter_DrawText8(QtObjectPtr ptr, QtObjectPtr rectangle, char* text, QtObjectPtr option){
	static_cast<QPainter*>(ptr)->drawText(*static_cast<QRectF*>(rectangle), QString(text), *static_cast<QTextOption*>(option));
}

void QPainter_DrawText4(QtObjectPtr ptr, QtObjectPtr rectangle, int flags, char* text, QtObjectPtr boundingRect){
	static_cast<QPainter*>(ptr)->drawText(*static_cast<QRectF*>(rectangle), flags, QString(text), static_cast<QRectF*>(boundingRect));
}

void QPainter_DrawTiledPixmap(QtObjectPtr ptr, QtObjectPtr rectangle, QtObjectPtr pixmap, QtObjectPtr position){
	static_cast<QPainter*>(ptr)->drawTiledPixmap(*static_cast<QRectF*>(rectangle), *static_cast<QPixmap*>(pixmap), *static_cast<QPointF*>(position));
}

void QPainter_EraseRect(QtObjectPtr ptr, QtObjectPtr rectangle){
	static_cast<QPainter*>(ptr)->eraseRect(*static_cast<QRectF*>(rectangle));
}

void QPainter_FillRect5(QtObjectPtr ptr, QtObjectPtr rectangle, QtObjectPtr brush){
	static_cast<QPainter*>(ptr)->fillRect(*static_cast<QRect*>(rectangle), *static_cast<QBrush*>(brush));
}

void QPainter_FillRect6(QtObjectPtr ptr, QtObjectPtr rectangle, QtObjectPtr color){
	static_cast<QPainter*>(ptr)->fillRect(*static_cast<QRect*>(rectangle), *static_cast<QColor*>(color));
}

void QPainter_FillRect(QtObjectPtr ptr, QtObjectPtr rectangle, QtObjectPtr brush){
	static_cast<QPainter*>(ptr)->fillRect(*static_cast<QRectF*>(rectangle), *static_cast<QBrush*>(brush));
}

void QPainter_FillRect7(QtObjectPtr ptr, QtObjectPtr rectangle, QtObjectPtr color){
	static_cast<QPainter*>(ptr)->fillRect(*static_cast<QRectF*>(rectangle), *static_cast<QColor*>(color));
}

void QPainter_SetBackground(QtObjectPtr ptr, QtObjectPtr brush){
	static_cast<QPainter*>(ptr)->setBackground(*static_cast<QBrush*>(brush));
}

void QPainter_SetBrushOrigin(QtObjectPtr ptr, QtObjectPtr position){
	static_cast<QPainter*>(ptr)->setBrushOrigin(*static_cast<QPointF*>(position));
}

void QPainter_SetClipPath(QtObjectPtr ptr, QtObjectPtr path, int operation){
	static_cast<QPainter*>(ptr)->setClipPath(*static_cast<QPainterPath*>(path), static_cast<Qt::ClipOperation>(operation));
}

void QPainter_SetClipRect3(QtObjectPtr ptr, QtObjectPtr rectangle, int operation){
	static_cast<QPainter*>(ptr)->setClipRect(*static_cast<QRect*>(rectangle), static_cast<Qt::ClipOperation>(operation));
}

void QPainter_SetClipRect(QtObjectPtr ptr, QtObjectPtr rectangle, int operation){
	static_cast<QPainter*>(ptr)->setClipRect(*static_cast<QRectF*>(rectangle), static_cast<Qt::ClipOperation>(operation));
}

void QPainter_SetClipRegion(QtObjectPtr ptr, QtObjectPtr region, int operation){
	static_cast<QPainter*>(ptr)->setClipRegion(*static_cast<QRegion*>(region), static_cast<Qt::ClipOperation>(operation));
}

void QPainter_SetViewport(QtObjectPtr ptr, QtObjectPtr rectangle){
	static_cast<QPainter*>(ptr)->setViewport(*static_cast<QRect*>(rectangle));
}

void QPainter_SetWindow(QtObjectPtr ptr, QtObjectPtr rectangle){
	static_cast<QPainter*>(ptr)->setWindow(*static_cast<QRect*>(rectangle));
}

QtObjectPtr QPainter_NewQPainter(){
	return new QPainter();
}

int QPainter_BackgroundMode(QtObjectPtr ptr){
	return static_cast<QPainter*>(ptr)->backgroundMode();
}

void QPainter_BeginNativePainting(QtObjectPtr ptr){
	static_cast<QPainter*>(ptr)->beginNativePainting();
}

int QPainter_CompositionMode(QtObjectPtr ptr){
	return static_cast<QPainter*>(ptr)->compositionMode();
}

QtObjectPtr QPainter_Device(QtObjectPtr ptr){
	return static_cast<QPainter*>(ptr)->device();
}

void QPainter_DrawArc2(QtObjectPtr ptr, QtObjectPtr rectangle, int startAngle, int spanAngle){
	static_cast<QPainter*>(ptr)->drawArc(*static_cast<QRect*>(rectangle), startAngle, spanAngle);
}

void QPainter_DrawArc3(QtObjectPtr ptr, int x, int y, int width, int height, int startAngle, int spanAngle){
	static_cast<QPainter*>(ptr)->drawArc(x, y, width, height, startAngle, spanAngle);
}

void QPainter_DrawChord2(QtObjectPtr ptr, QtObjectPtr rectangle, int startAngle, int spanAngle){
	static_cast<QPainter*>(ptr)->drawChord(*static_cast<QRect*>(rectangle), startAngle, spanAngle);
}

void QPainter_DrawChord3(QtObjectPtr ptr, int x, int y, int width, int height, int startAngle, int spanAngle){
	static_cast<QPainter*>(ptr)->drawChord(x, y, width, height, startAngle, spanAngle);
}

void QPainter_DrawConvexPolygon4(QtObjectPtr ptr, QtObjectPtr polygon){
	static_cast<QPainter*>(ptr)->drawConvexPolygon(*static_cast<QPolygon*>(polygon));
}

void QPainter_DrawConvexPolygon3(QtObjectPtr ptr, QtObjectPtr polygon){
	static_cast<QPainter*>(ptr)->drawConvexPolygon(*static_cast<QPolygonF*>(polygon));
}

void QPainter_DrawEllipse5(QtObjectPtr ptr, QtObjectPtr center, int rx, int ry){
	static_cast<QPainter*>(ptr)->drawEllipse(*static_cast<QPoint*>(center), rx, ry);
}

void QPainter_DrawEllipse3(QtObjectPtr ptr, int x, int y, int width, int height){
	static_cast<QPainter*>(ptr)->drawEllipse(x, y, width, height);
}

void QPainter_DrawImage4(QtObjectPtr ptr, QtObjectPtr point, QtObjectPtr image){
	static_cast<QPainter*>(ptr)->drawImage(*static_cast<QPoint*>(point), *static_cast<QImage*>(image));
}

void QPainter_DrawImage6(QtObjectPtr ptr, QtObjectPtr point, QtObjectPtr image, QtObjectPtr source, int flags){
	static_cast<QPainter*>(ptr)->drawImage(*static_cast<QPoint*>(point), *static_cast<QImage*>(image), *static_cast<QRect*>(source), static_cast<Qt::ImageConversionFlag>(flags));
}

void QPainter_DrawImage5(QtObjectPtr ptr, QtObjectPtr point, QtObjectPtr image, QtObjectPtr source, int flags){
	static_cast<QPainter*>(ptr)->drawImage(*static_cast<QPointF*>(point), *static_cast<QImage*>(image), *static_cast<QRectF*>(source), static_cast<Qt::ImageConversionFlag>(flags));
}

void QPainter_DrawImage8(QtObjectPtr ptr, QtObjectPtr rectangle, QtObjectPtr image){
	static_cast<QPainter*>(ptr)->drawImage(*static_cast<QRect*>(rectangle), *static_cast<QImage*>(image));
}

void QPainter_DrawImage2(QtObjectPtr ptr, QtObjectPtr target, QtObjectPtr image, QtObjectPtr source, int flags){
	static_cast<QPainter*>(ptr)->drawImage(*static_cast<QRect*>(target), *static_cast<QImage*>(image), *static_cast<QRect*>(source), static_cast<Qt::ImageConversionFlag>(flags));
}

void QPainter_DrawImage7(QtObjectPtr ptr, QtObjectPtr rectangle, QtObjectPtr image){
	static_cast<QPainter*>(ptr)->drawImage(*static_cast<QRectF*>(rectangle), *static_cast<QImage*>(image));
}

void QPainter_DrawImage9(QtObjectPtr ptr, int x, int y, QtObjectPtr image, int sx, int sy, int sw, int sh, int flags){
	static_cast<QPainter*>(ptr)->drawImage(x, y, *static_cast<QImage*>(image), sx, sy, sw, sh, static_cast<Qt::ImageConversionFlag>(flags));
}

void QPainter_DrawLine2(QtObjectPtr ptr, QtObjectPtr line){
	static_cast<QPainter*>(ptr)->drawLine(*static_cast<QLine*>(line));
}

void QPainter_DrawLine(QtObjectPtr ptr, QtObjectPtr line){
	static_cast<QPainter*>(ptr)->drawLine(*static_cast<QLineF*>(line));
}

void QPainter_DrawLine3(QtObjectPtr ptr, QtObjectPtr p1, QtObjectPtr p2){
	static_cast<QPainter*>(ptr)->drawLine(*static_cast<QPoint*>(p1), *static_cast<QPoint*>(p2));
}

void QPainter_DrawLine4(QtObjectPtr ptr, QtObjectPtr p1, QtObjectPtr p2){
	static_cast<QPainter*>(ptr)->drawLine(*static_cast<QPointF*>(p1), *static_cast<QPointF*>(p2));
}

void QPainter_DrawLine5(QtObjectPtr ptr, int x1, int y1, int x2, int y2){
	static_cast<QPainter*>(ptr)->drawLine(x1, y1, x2, y2);
}

void QPainter_DrawLines(QtObjectPtr ptr, QtObjectPtr lines, int lineCount){
	static_cast<QPainter*>(ptr)->drawLines(static_cast<QLineF*>(lines), lineCount);
}

void QPainter_DrawLines4(QtObjectPtr ptr, QtObjectPtr pointPairs, int lineCount){
	static_cast<QPainter*>(ptr)->drawLines(static_cast<QPoint*>(pointPairs), lineCount);
}

void QPainter_DrawLines3(QtObjectPtr ptr, QtObjectPtr pointPairs, int lineCount){
	static_cast<QPainter*>(ptr)->drawLines(static_cast<QPointF*>(pointPairs), lineCount);
}

void QPainter_DrawPath(QtObjectPtr ptr, QtObjectPtr path){
	static_cast<QPainter*>(ptr)->drawPath(*static_cast<QPainterPath*>(path));
}

void QPainter_DrawPicture2(QtObjectPtr ptr, QtObjectPtr point, QtObjectPtr picture){
	static_cast<QPainter*>(ptr)->drawPicture(*static_cast<QPoint*>(point), *static_cast<QPicture*>(picture));
}

void QPainter_DrawPicture3(QtObjectPtr ptr, int x, int y, QtObjectPtr picture){
	static_cast<QPainter*>(ptr)->drawPicture(x, y, *static_cast<QPicture*>(picture));
}

void QPainter_DrawPie2(QtObjectPtr ptr, QtObjectPtr rectangle, int startAngle, int spanAngle){
	static_cast<QPainter*>(ptr)->drawPie(*static_cast<QRect*>(rectangle), startAngle, spanAngle);
}

void QPainter_DrawPie3(QtObjectPtr ptr, int x, int y, int width, int height, int startAngle, int spanAngle){
	static_cast<QPainter*>(ptr)->drawPie(x, y, width, height, startAngle, spanAngle);
}

void QPainter_DrawPixmap6(QtObjectPtr ptr, QtObjectPtr point, QtObjectPtr pixmap){
	static_cast<QPainter*>(ptr)->drawPixmap(*static_cast<QPoint*>(point), *static_cast<QPixmap*>(pixmap));
}

void QPainter_DrawPixmap4(QtObjectPtr ptr, QtObjectPtr point, QtObjectPtr pixmap, QtObjectPtr source){
	static_cast<QPainter*>(ptr)->drawPixmap(*static_cast<QPoint*>(point), *static_cast<QPixmap*>(pixmap), *static_cast<QRect*>(source));
}

void QPainter_DrawPixmap3(QtObjectPtr ptr, QtObjectPtr point, QtObjectPtr pixmap, QtObjectPtr source){
	static_cast<QPainter*>(ptr)->drawPixmap(*static_cast<QPointF*>(point), *static_cast<QPixmap*>(pixmap), *static_cast<QRectF*>(source));
}

void QPainter_DrawPixmap8(QtObjectPtr ptr, QtObjectPtr rectangle, QtObjectPtr pixmap){
	static_cast<QPainter*>(ptr)->drawPixmap(*static_cast<QRect*>(rectangle), *static_cast<QPixmap*>(pixmap));
}

void QPainter_DrawPixmap2(QtObjectPtr ptr, QtObjectPtr target, QtObjectPtr pixmap, QtObjectPtr source){
	static_cast<QPainter*>(ptr)->drawPixmap(*static_cast<QRect*>(target), *static_cast<QPixmap*>(pixmap), *static_cast<QRect*>(source));
}

void QPainter_DrawPixmap7(QtObjectPtr ptr, int x, int y, QtObjectPtr pixmap){
	static_cast<QPainter*>(ptr)->drawPixmap(x, y, *static_cast<QPixmap*>(pixmap));
}

void QPainter_DrawPixmap11(QtObjectPtr ptr, int x, int y, QtObjectPtr pixmap, int sx, int sy, int sw, int sh){
	static_cast<QPainter*>(ptr)->drawPixmap(x, y, *static_cast<QPixmap*>(pixmap), sx, sy, sw, sh);
}

void QPainter_DrawPixmap10(QtObjectPtr ptr, int x, int y, int w, int h, QtObjectPtr pixmap, int sx, int sy, int sw, int sh){
	static_cast<QPainter*>(ptr)->drawPixmap(x, y, w, h, *static_cast<QPixmap*>(pixmap), sx, sy, sw, sh);
}

void QPainter_DrawPixmap9(QtObjectPtr ptr, int x, int y, int width, int height, QtObjectPtr pixmap){
	static_cast<QPainter*>(ptr)->drawPixmap(x, y, width, height, *static_cast<QPixmap*>(pixmap));
}

void QPainter_DrawPoint2(QtObjectPtr ptr, QtObjectPtr position){
	static_cast<QPainter*>(ptr)->drawPoint(*static_cast<QPoint*>(position));
}

void QPainter_DrawPoint(QtObjectPtr ptr, QtObjectPtr position){
	static_cast<QPainter*>(ptr)->drawPoint(*static_cast<QPointF*>(position));
}

void QPainter_DrawPoint3(QtObjectPtr ptr, int x, int y){
	static_cast<QPainter*>(ptr)->drawPoint(x, y);
}

void QPainter_DrawPoints2(QtObjectPtr ptr, QtObjectPtr points, int pointCount){
	static_cast<QPainter*>(ptr)->drawPoints(static_cast<QPoint*>(points), pointCount);
}

void QPainter_DrawPoints(QtObjectPtr ptr, QtObjectPtr points, int pointCount){
	static_cast<QPainter*>(ptr)->drawPoints(static_cast<QPointF*>(points), pointCount);
}

void QPainter_DrawPoints4(QtObjectPtr ptr, QtObjectPtr points){
	static_cast<QPainter*>(ptr)->drawPoints(*static_cast<QPolygon*>(points));
}

void QPainter_DrawPoints3(QtObjectPtr ptr, QtObjectPtr points){
	static_cast<QPainter*>(ptr)->drawPoints(*static_cast<QPolygonF*>(points));
}

void QPainter_DrawPolygon2(QtObjectPtr ptr, QtObjectPtr points, int pointCount, int fillRule){
	static_cast<QPainter*>(ptr)->drawPolygon(static_cast<QPoint*>(points), pointCount, static_cast<Qt::FillRule>(fillRule));
}

void QPainter_DrawPolygon(QtObjectPtr ptr, QtObjectPtr points, int pointCount, int fillRule){
	static_cast<QPainter*>(ptr)->drawPolygon(static_cast<QPointF*>(points), pointCount, static_cast<Qt::FillRule>(fillRule));
}

void QPainter_DrawPolygon4(QtObjectPtr ptr, QtObjectPtr points, int fillRule){
	static_cast<QPainter*>(ptr)->drawPolygon(*static_cast<QPolygon*>(points), static_cast<Qt::FillRule>(fillRule));
}

void QPainter_DrawPolygon3(QtObjectPtr ptr, QtObjectPtr points, int fillRule){
	static_cast<QPainter*>(ptr)->drawPolygon(*static_cast<QPolygonF*>(points), static_cast<Qt::FillRule>(fillRule));
}

void QPainter_DrawPolyline2(QtObjectPtr ptr, QtObjectPtr points, int pointCount){
	static_cast<QPainter*>(ptr)->drawPolyline(static_cast<QPoint*>(points), pointCount);
}

void QPainter_DrawPolyline(QtObjectPtr ptr, QtObjectPtr points, int pointCount){
	static_cast<QPainter*>(ptr)->drawPolyline(static_cast<QPointF*>(points), pointCount);
}

void QPainter_DrawPolyline4(QtObjectPtr ptr, QtObjectPtr points){
	static_cast<QPainter*>(ptr)->drawPolyline(*static_cast<QPolygon*>(points));
}

void QPainter_DrawPolyline3(QtObjectPtr ptr, QtObjectPtr points){
	static_cast<QPainter*>(ptr)->drawPolyline(*static_cast<QPolygonF*>(points));
}

void QPainter_DrawRect2(QtObjectPtr ptr, QtObjectPtr rectangle){
	static_cast<QPainter*>(ptr)->drawRect(*static_cast<QRect*>(rectangle));
}

void QPainter_DrawRect(QtObjectPtr ptr, QtObjectPtr rectangle){
	static_cast<QPainter*>(ptr)->drawRect(*static_cast<QRectF*>(rectangle));
}

void QPainter_DrawRect3(QtObjectPtr ptr, int x, int y, int width, int height){
	static_cast<QPainter*>(ptr)->drawRect(x, y, width, height);
}

void QPainter_DrawStaticText2(QtObjectPtr ptr, QtObjectPtr topLeftPosition, QtObjectPtr staticText){
	static_cast<QPainter*>(ptr)->drawStaticText(*static_cast<QPoint*>(topLeftPosition), *static_cast<QStaticText*>(staticText));
}

void QPainter_DrawStaticText(QtObjectPtr ptr, QtObjectPtr topLeftPosition, QtObjectPtr staticText){
	static_cast<QPainter*>(ptr)->drawStaticText(*static_cast<QPointF*>(topLeftPosition), *static_cast<QStaticText*>(staticText));
}

void QPainter_DrawStaticText3(QtObjectPtr ptr, int left, int top, QtObjectPtr staticText){
	static_cast<QPainter*>(ptr)->drawStaticText(left, top, *static_cast<QStaticText*>(staticText));
}

void QPainter_DrawText3(QtObjectPtr ptr, QtObjectPtr position, char* text){
	static_cast<QPainter*>(ptr)->drawText(*static_cast<QPoint*>(position), QString(text));
}

void QPainter_DrawText6(QtObjectPtr ptr, int x, int y, char* text){
	static_cast<QPainter*>(ptr)->drawText(x, y, QString(text));
}

void QPainter_DrawText7(QtObjectPtr ptr, int x, int y, int width, int height, int flags, char* text, QtObjectPtr boundingRect){
	static_cast<QPainter*>(ptr)->drawText(x, y, width, height, flags, QString(text), static_cast<QRect*>(boundingRect));
}

void QPainter_DrawTiledPixmap2(QtObjectPtr ptr, QtObjectPtr rectangle, QtObjectPtr pixmap, QtObjectPtr position){
	static_cast<QPainter*>(ptr)->drawTiledPixmap(*static_cast<QRect*>(rectangle), *static_cast<QPixmap*>(pixmap), *static_cast<QPoint*>(position));
}

void QPainter_DrawTiledPixmap3(QtObjectPtr ptr, int x, int y, int width, int height, QtObjectPtr pixmap, int sx, int sy){
	static_cast<QPainter*>(ptr)->drawTiledPixmap(x, y, width, height, *static_cast<QPixmap*>(pixmap), sx, sy);
}

int QPainter_End(QtObjectPtr ptr){
	return static_cast<QPainter*>(ptr)->end();
}

void QPainter_EndNativePainting(QtObjectPtr ptr){
	static_cast<QPainter*>(ptr)->endNativePainting();
}

void QPainter_EraseRect2(QtObjectPtr ptr, QtObjectPtr rectangle){
	static_cast<QPainter*>(ptr)->eraseRect(*static_cast<QRect*>(rectangle));
}

void QPainter_EraseRect3(QtObjectPtr ptr, int x, int y, int width, int height){
	static_cast<QPainter*>(ptr)->eraseRect(x, y, width, height);
}

void QPainter_FillPath(QtObjectPtr ptr, QtObjectPtr path, QtObjectPtr brush){
	static_cast<QPainter*>(ptr)->fillPath(*static_cast<QPainterPath*>(path), *static_cast<QBrush*>(brush));
}

void QPainter_FillRect3(QtObjectPtr ptr, QtObjectPtr rectangle, int style){
	static_cast<QPainter*>(ptr)->fillRect(*static_cast<QRect*>(rectangle), static_cast<Qt::BrushStyle>(style));
}

void QPainter_FillRect11(QtObjectPtr ptr, QtObjectPtr rectangle, int color){
	static_cast<QPainter*>(ptr)->fillRect(*static_cast<QRect*>(rectangle), static_cast<Qt::GlobalColor>(color));
}

void QPainter_FillRect4(QtObjectPtr ptr, QtObjectPtr rectangle, int style){
	static_cast<QPainter*>(ptr)->fillRect(*static_cast<QRectF*>(rectangle), static_cast<Qt::BrushStyle>(style));
}

void QPainter_FillRect12(QtObjectPtr ptr, QtObjectPtr rectangle, int color){
	static_cast<QPainter*>(ptr)->fillRect(*static_cast<QRectF*>(rectangle), static_cast<Qt::GlobalColor>(color));
}

void QPainter_FillRect2(QtObjectPtr ptr, int x, int y, int width, int height, int style){
	static_cast<QPainter*>(ptr)->fillRect(x, y, width, height, static_cast<Qt::BrushStyle>(style));
}

void QPainter_FillRect10(QtObjectPtr ptr, int x, int y, int width, int height, int color){
	static_cast<QPainter*>(ptr)->fillRect(x, y, width, height, static_cast<Qt::GlobalColor>(color));
}

void QPainter_FillRect8(QtObjectPtr ptr, int x, int y, int width, int height, QtObjectPtr brush){
	static_cast<QPainter*>(ptr)->fillRect(x, y, width, height, *static_cast<QBrush*>(brush));
}

void QPainter_FillRect9(QtObjectPtr ptr, int x, int y, int width, int height, QtObjectPtr color){
	static_cast<QPainter*>(ptr)->fillRect(x, y, width, height, *static_cast<QColor*>(color));
}

int QPainter_HasClipping(QtObjectPtr ptr){
	return static_cast<QPainter*>(ptr)->hasClipping();
}

int QPainter_IsActive(QtObjectPtr ptr){
	return static_cast<QPainter*>(ptr)->isActive();
}

int QPainter_LayoutDirection(QtObjectPtr ptr){
	return static_cast<QPainter*>(ptr)->layoutDirection();
}

QtObjectPtr QPainter_PaintEngine(QtObjectPtr ptr){
	return static_cast<QPainter*>(ptr)->paintEngine();
}

int QPainter_RenderHints(QtObjectPtr ptr){
	return static_cast<QPainter*>(ptr)->renderHints();
}

void QPainter_ResetTransform(QtObjectPtr ptr){
	static_cast<QPainter*>(ptr)->resetTransform();
}

void QPainter_Restore(QtObjectPtr ptr){
	static_cast<QPainter*>(ptr)->restore();
}

void QPainter_Save(QtObjectPtr ptr){
	static_cast<QPainter*>(ptr)->save();
}

void QPainter_SetBackgroundMode(QtObjectPtr ptr, int mode){
	static_cast<QPainter*>(ptr)->setBackgroundMode(static_cast<Qt::BGMode>(mode));
}

void QPainter_SetBrush2(QtObjectPtr ptr, int style){
	static_cast<QPainter*>(ptr)->setBrush(static_cast<Qt::BrushStyle>(style));
}

void QPainter_SetBrush(QtObjectPtr ptr, QtObjectPtr brush){
	static_cast<QPainter*>(ptr)->setBrush(*static_cast<QBrush*>(brush));
}

void QPainter_SetBrushOrigin2(QtObjectPtr ptr, QtObjectPtr position){
	static_cast<QPainter*>(ptr)->setBrushOrigin(*static_cast<QPoint*>(position));
}

void QPainter_SetBrushOrigin3(QtObjectPtr ptr, int x, int y){
	static_cast<QPainter*>(ptr)->setBrushOrigin(x, y);
}

void QPainter_SetClipRect2(QtObjectPtr ptr, int x, int y, int width, int height, int operation){
	static_cast<QPainter*>(ptr)->setClipRect(x, y, width, height, static_cast<Qt::ClipOperation>(operation));
}

void QPainter_SetClipping(QtObjectPtr ptr, int enable){
	static_cast<QPainter*>(ptr)->setClipping(enable != 0);
}

void QPainter_SetCompositionMode(QtObjectPtr ptr, int mode){
	static_cast<QPainter*>(ptr)->setCompositionMode(static_cast<QPainter::CompositionMode>(mode));
}

void QPainter_SetFont(QtObjectPtr ptr, QtObjectPtr font){
	static_cast<QPainter*>(ptr)->setFont(*static_cast<QFont*>(font));
}

void QPainter_SetLayoutDirection(QtObjectPtr ptr, int direction){
	static_cast<QPainter*>(ptr)->setLayoutDirection(static_cast<Qt::LayoutDirection>(direction));
}

void QPainter_SetPen3(QtObjectPtr ptr, int style){
	static_cast<QPainter*>(ptr)->setPen(static_cast<Qt::PenStyle>(style));
}

void QPainter_SetPen2(QtObjectPtr ptr, QtObjectPtr color){
	static_cast<QPainter*>(ptr)->setPen(*static_cast<QColor*>(color));
}

void QPainter_SetPen(QtObjectPtr ptr, QtObjectPtr pen){
	static_cast<QPainter*>(ptr)->setPen(*static_cast<QPen*>(pen));
}

void QPainter_SetRenderHint(QtObjectPtr ptr, int hint, int on){
	static_cast<QPainter*>(ptr)->setRenderHint(static_cast<QPainter::RenderHint>(hint), on != 0);
}

void QPainter_SetRenderHints(QtObjectPtr ptr, int hints, int on){
	static_cast<QPainter*>(ptr)->setRenderHints(static_cast<QPainter::RenderHint>(hints), on != 0);
}

void QPainter_SetTransform(QtObjectPtr ptr, QtObjectPtr transform, int combine){
	static_cast<QPainter*>(ptr)->setTransform(*static_cast<QTransform*>(transform), combine != 0);
}

void QPainter_SetViewTransformEnabled(QtObjectPtr ptr, int enable){
	static_cast<QPainter*>(ptr)->setViewTransformEnabled(enable != 0);
}

void QPainter_SetViewport2(QtObjectPtr ptr, int x, int y, int width, int height){
	static_cast<QPainter*>(ptr)->setViewport(x, y, width, height);
}

void QPainter_SetWindow2(QtObjectPtr ptr, int x, int y, int width, int height){
	static_cast<QPainter*>(ptr)->setWindow(x, y, width, height);
}

void QPainter_SetWorldMatrixEnabled(QtObjectPtr ptr, int enable){
	static_cast<QPainter*>(ptr)->setWorldMatrixEnabled(enable != 0);
}

void QPainter_SetWorldTransform(QtObjectPtr ptr, QtObjectPtr matrix, int combine){
	static_cast<QPainter*>(ptr)->setWorldTransform(*static_cast<QTransform*>(matrix), combine != 0);
}

void QPainter_StrokePath(QtObjectPtr ptr, QtObjectPtr path, QtObjectPtr pen){
	static_cast<QPainter*>(ptr)->strokePath(*static_cast<QPainterPath*>(path), *static_cast<QPen*>(pen));
}

int QPainter_TestRenderHint(QtObjectPtr ptr, int hint){
	return static_cast<QPainter*>(ptr)->testRenderHint(static_cast<QPainter::RenderHint>(hint));
}

void QPainter_Translate2(QtObjectPtr ptr, QtObjectPtr offset){
	static_cast<QPainter*>(ptr)->translate(*static_cast<QPoint*>(offset));
}

void QPainter_Translate(QtObjectPtr ptr, QtObjectPtr offset){
	static_cast<QPainter*>(ptr)->translate(*static_cast<QPointF*>(offset));
}

int QPainter_ViewTransformEnabled(QtObjectPtr ptr){
	return static_cast<QPainter*>(ptr)->viewTransformEnabled();
}

int QPainter_WorldMatrixEnabled(QtObjectPtr ptr){
	return static_cast<QPainter*>(ptr)->worldMatrixEnabled();
}

void QPainter_DestroyQPainter(QtObjectPtr ptr){
	static_cast<QPainter*>(ptr)->~QPainter();
}

