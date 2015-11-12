#include "qpainter.h"
#include <QRect>
#include <QBrush>
#include <QStaticText>
#include <QPolygon>
#include <QLine>
#include <QRegion>
#include <QTextOption>
#include <QTransform>
#include <QLineF>
#include <QFont>
#include <QImage>
#include <QPen>
#include <QPaintDevice>
#include <QModelIndex>
#include <QPointF>
#include <QUrl>
#include <QPainterPath>
#include <QRectF>
#include <QColor>
#include <QPolygonF>
#include <QGlyphRun>
#include <QString>
#include <QVariant>
#include <QPixmap>
#include <QPoint>
#include <QPicture>
#include <QPainter>
#include "_cgo_export.h"

class MyQPainter: public QPainter {
public:
};

void* QPainter_NewQPainter2(void* device){
	return new QPainter(static_cast<QPaintDevice*>(device));
}

int QPainter_Begin(void* ptr, void* device){
	return static_cast<QPainter*>(ptr)->begin(static_cast<QPaintDevice*>(device));
}

void QPainter_DrawArc(void* ptr, void* rectangle, int startAngle, int spanAngle){
	static_cast<QPainter*>(ptr)->drawArc(*static_cast<QRectF*>(rectangle), startAngle, spanAngle);
}

void QPainter_DrawChord(void* ptr, void* rectangle, int startAngle, int spanAngle){
	static_cast<QPainter*>(ptr)->drawChord(*static_cast<QRectF*>(rectangle), startAngle, spanAngle);
}

void QPainter_DrawConvexPolygon2(void* ptr, void* points, int pointCount){
	static_cast<QPainter*>(ptr)->drawConvexPolygon(static_cast<QPoint*>(points), pointCount);
}

void QPainter_DrawConvexPolygon(void* ptr, void* points, int pointCount){
	static_cast<QPainter*>(ptr)->drawConvexPolygon(static_cast<QPointF*>(points), pointCount);
}

void QPainter_DrawEllipse2(void* ptr, void* rectangle){
	static_cast<QPainter*>(ptr)->drawEllipse(*static_cast<QRect*>(rectangle));
}

void QPainter_DrawEllipse(void* ptr, void* rectangle){
	static_cast<QPainter*>(ptr)->drawEllipse(*static_cast<QRectF*>(rectangle));
}

void QPainter_DrawGlyphRun(void* ptr, void* position, void* glyphs){
	static_cast<QPainter*>(ptr)->drawGlyphRun(*static_cast<QPointF*>(position), *static_cast<QGlyphRun*>(glyphs));
}

void QPainter_DrawImage3(void* ptr, void* point, void* image){
	static_cast<QPainter*>(ptr)->drawImage(*static_cast<QPointF*>(point), *static_cast<QImage*>(image));
}

void QPainter_DrawImage(void* ptr, void* target, void* image, void* source, int flags){
	static_cast<QPainter*>(ptr)->drawImage(*static_cast<QRectF*>(target), *static_cast<QImage*>(image), *static_cast<QRectF*>(source), static_cast<Qt::ImageConversionFlag>(flags));
}

void QPainter_DrawLines2(void* ptr, void* lines, int lineCount){
	static_cast<QPainter*>(ptr)->drawLines(static_cast<QLine*>(lines), lineCount);
}

void QPainter_DrawPicture(void* ptr, void* point, void* picture){
	static_cast<QPainter*>(ptr)->drawPicture(*static_cast<QPointF*>(point), *static_cast<QPicture*>(picture));
}

void QPainter_DrawPie(void* ptr, void* rectangle, int startAngle, int spanAngle){
	static_cast<QPainter*>(ptr)->drawPie(*static_cast<QRectF*>(rectangle), startAngle, spanAngle);
}

void QPainter_DrawPixmap5(void* ptr, void* point, void* pixmap){
	static_cast<QPainter*>(ptr)->drawPixmap(*static_cast<QPointF*>(point), *static_cast<QPixmap*>(pixmap));
}

void QPainter_DrawPixmap(void* ptr, void* target, void* pixmap, void* source){
	static_cast<QPainter*>(ptr)->drawPixmap(*static_cast<QRectF*>(target), *static_cast<QPixmap*>(pixmap), *static_cast<QRectF*>(source));
}

void QPainter_DrawRects2(void* ptr, void* rectangles, int rectCount){
	static_cast<QPainter*>(ptr)->drawRects(static_cast<QRect*>(rectangles), rectCount);
}

void QPainter_DrawRects(void* ptr, void* rectangles, int rectCount){
	static_cast<QPainter*>(ptr)->drawRects(static_cast<QRectF*>(rectangles), rectCount);
}

void QPainter_DrawText(void* ptr, void* position, char* text){
	static_cast<QPainter*>(ptr)->drawText(*static_cast<QPointF*>(position), QString(text));
}

void QPainter_DrawText5(void* ptr, void* rectangle, int flags, char* text, void* boundingRect){
	static_cast<QPainter*>(ptr)->drawText(*static_cast<QRect*>(rectangle), flags, QString(text), static_cast<QRect*>(boundingRect));
}

void QPainter_DrawText8(void* ptr, void* rectangle, char* text, void* option){
	static_cast<QPainter*>(ptr)->drawText(*static_cast<QRectF*>(rectangle), QString(text), *static_cast<QTextOption*>(option));
}

void QPainter_DrawText4(void* ptr, void* rectangle, int flags, char* text, void* boundingRect){
	static_cast<QPainter*>(ptr)->drawText(*static_cast<QRectF*>(rectangle), flags, QString(text), static_cast<QRectF*>(boundingRect));
}

void QPainter_DrawTiledPixmap(void* ptr, void* rectangle, void* pixmap, void* position){
	static_cast<QPainter*>(ptr)->drawTiledPixmap(*static_cast<QRectF*>(rectangle), *static_cast<QPixmap*>(pixmap), *static_cast<QPointF*>(position));
}

void QPainter_EraseRect(void* ptr, void* rectangle){
	static_cast<QPainter*>(ptr)->eraseRect(*static_cast<QRectF*>(rectangle));
}

void QPainter_FillRect5(void* ptr, void* rectangle, void* brush){
	static_cast<QPainter*>(ptr)->fillRect(*static_cast<QRect*>(rectangle), *static_cast<QBrush*>(brush));
}

void QPainter_FillRect6(void* ptr, void* rectangle, void* color){
	static_cast<QPainter*>(ptr)->fillRect(*static_cast<QRect*>(rectangle), *static_cast<QColor*>(color));
}

void QPainter_FillRect(void* ptr, void* rectangle, void* brush){
	static_cast<QPainter*>(ptr)->fillRect(*static_cast<QRectF*>(rectangle), *static_cast<QBrush*>(brush));
}

void QPainter_FillRect7(void* ptr, void* rectangle, void* color){
	static_cast<QPainter*>(ptr)->fillRect(*static_cast<QRectF*>(rectangle), *static_cast<QColor*>(color));
}

void QPainter_Rotate(void* ptr, double angle){
	static_cast<QPainter*>(ptr)->rotate(static_cast<qreal>(angle));
}

void QPainter_SetBackground(void* ptr, void* brush){
	static_cast<QPainter*>(ptr)->setBackground(*static_cast<QBrush*>(brush));
}

void QPainter_SetBrushOrigin(void* ptr, void* position){
	static_cast<QPainter*>(ptr)->setBrushOrigin(*static_cast<QPointF*>(position));
}

void QPainter_SetClipPath(void* ptr, void* path, int operation){
	static_cast<QPainter*>(ptr)->setClipPath(*static_cast<QPainterPath*>(path), static_cast<Qt::ClipOperation>(operation));
}

void QPainter_SetClipRect3(void* ptr, void* rectangle, int operation){
	static_cast<QPainter*>(ptr)->setClipRect(*static_cast<QRect*>(rectangle), static_cast<Qt::ClipOperation>(operation));
}

void QPainter_SetClipRect(void* ptr, void* rectangle, int operation){
	static_cast<QPainter*>(ptr)->setClipRect(*static_cast<QRectF*>(rectangle), static_cast<Qt::ClipOperation>(operation));
}

void QPainter_SetClipRegion(void* ptr, void* region, int operation){
	static_cast<QPainter*>(ptr)->setClipRegion(*static_cast<QRegion*>(region), static_cast<Qt::ClipOperation>(operation));
}

void QPainter_SetViewport(void* ptr, void* rectangle){
	static_cast<QPainter*>(ptr)->setViewport(*static_cast<QRect*>(rectangle));
}

void QPainter_SetWindow(void* ptr, void* rectangle){
	static_cast<QPainter*>(ptr)->setWindow(*static_cast<QRect*>(rectangle));
}

void* QPainter_NewQPainter(){
	return new QPainter();
}

void* QPainter_Background(void* ptr){
	return new QBrush(static_cast<QPainter*>(ptr)->background());
}

int QPainter_BackgroundMode(void* ptr){
	return static_cast<QPainter*>(ptr)->backgroundMode();
}

void QPainter_BeginNativePainting(void* ptr){
	static_cast<QPainter*>(ptr)->beginNativePainting();
}

void* QPainter_Brush(void* ptr){
	return new QBrush(static_cast<QPainter*>(ptr)->brush());
}

void* QPainter_ClipRegion(void* ptr){
	return new QRegion(static_cast<QPainter*>(ptr)->clipRegion());
}

int QPainter_CompositionMode(void* ptr){
	return static_cast<QPainter*>(ptr)->compositionMode();
}

void* QPainter_Device(void* ptr){
	return static_cast<QPainter*>(ptr)->device();
}

void QPainter_DrawArc2(void* ptr, void* rectangle, int startAngle, int spanAngle){
	static_cast<QPainter*>(ptr)->drawArc(*static_cast<QRect*>(rectangle), startAngle, spanAngle);
}

void QPainter_DrawArc3(void* ptr, int x, int y, int width, int height, int startAngle, int spanAngle){
	static_cast<QPainter*>(ptr)->drawArc(x, y, width, height, startAngle, spanAngle);
}

void QPainter_DrawChord2(void* ptr, void* rectangle, int startAngle, int spanAngle){
	static_cast<QPainter*>(ptr)->drawChord(*static_cast<QRect*>(rectangle), startAngle, spanAngle);
}

void QPainter_DrawChord3(void* ptr, int x, int y, int width, int height, int startAngle, int spanAngle){
	static_cast<QPainter*>(ptr)->drawChord(x, y, width, height, startAngle, spanAngle);
}

void QPainter_DrawConvexPolygon4(void* ptr, void* polygon){
	static_cast<QPainter*>(ptr)->drawConvexPolygon(*static_cast<QPolygon*>(polygon));
}

void QPainter_DrawConvexPolygon3(void* ptr, void* polygon){
	static_cast<QPainter*>(ptr)->drawConvexPolygon(*static_cast<QPolygonF*>(polygon));
}

void QPainter_DrawEllipse5(void* ptr, void* center, int rx, int ry){
	static_cast<QPainter*>(ptr)->drawEllipse(*static_cast<QPoint*>(center), rx, ry);
}

void QPainter_DrawEllipse4(void* ptr, void* center, double rx, double ry){
	static_cast<QPainter*>(ptr)->drawEllipse(*static_cast<QPointF*>(center), static_cast<qreal>(rx), static_cast<qreal>(ry));
}

void QPainter_DrawEllipse3(void* ptr, int x, int y, int width, int height){
	static_cast<QPainter*>(ptr)->drawEllipse(x, y, width, height);
}

void QPainter_DrawImage4(void* ptr, void* point, void* image){
	static_cast<QPainter*>(ptr)->drawImage(*static_cast<QPoint*>(point), *static_cast<QImage*>(image));
}

void QPainter_DrawImage6(void* ptr, void* point, void* image, void* source, int flags){
	static_cast<QPainter*>(ptr)->drawImage(*static_cast<QPoint*>(point), *static_cast<QImage*>(image), *static_cast<QRect*>(source), static_cast<Qt::ImageConversionFlag>(flags));
}

void QPainter_DrawImage5(void* ptr, void* point, void* image, void* source, int flags){
	static_cast<QPainter*>(ptr)->drawImage(*static_cast<QPointF*>(point), *static_cast<QImage*>(image), *static_cast<QRectF*>(source), static_cast<Qt::ImageConversionFlag>(flags));
}

void QPainter_DrawImage8(void* ptr, void* rectangle, void* image){
	static_cast<QPainter*>(ptr)->drawImage(*static_cast<QRect*>(rectangle), *static_cast<QImage*>(image));
}

void QPainter_DrawImage2(void* ptr, void* target, void* image, void* source, int flags){
	static_cast<QPainter*>(ptr)->drawImage(*static_cast<QRect*>(target), *static_cast<QImage*>(image), *static_cast<QRect*>(source), static_cast<Qt::ImageConversionFlag>(flags));
}

void QPainter_DrawImage7(void* ptr, void* rectangle, void* image){
	static_cast<QPainter*>(ptr)->drawImage(*static_cast<QRectF*>(rectangle), *static_cast<QImage*>(image));
}

void QPainter_DrawImage9(void* ptr, int x, int y, void* image, int sx, int sy, int sw, int sh, int flags){
	static_cast<QPainter*>(ptr)->drawImage(x, y, *static_cast<QImage*>(image), sx, sy, sw, sh, static_cast<Qt::ImageConversionFlag>(flags));
}

void QPainter_DrawLine2(void* ptr, void* line){
	static_cast<QPainter*>(ptr)->drawLine(*static_cast<QLine*>(line));
}

void QPainter_DrawLine(void* ptr, void* line){
	static_cast<QPainter*>(ptr)->drawLine(*static_cast<QLineF*>(line));
}

void QPainter_DrawLine3(void* ptr, void* p1, void* p2){
	static_cast<QPainter*>(ptr)->drawLine(*static_cast<QPoint*>(p1), *static_cast<QPoint*>(p2));
}

void QPainter_DrawLine4(void* ptr, void* p1, void* p2){
	static_cast<QPainter*>(ptr)->drawLine(*static_cast<QPointF*>(p1), *static_cast<QPointF*>(p2));
}

void QPainter_DrawLine5(void* ptr, int x1, int y1, int x2, int y2){
	static_cast<QPainter*>(ptr)->drawLine(x1, y1, x2, y2);
}

void QPainter_DrawLines(void* ptr, void* lines, int lineCount){
	static_cast<QPainter*>(ptr)->drawLines(static_cast<QLineF*>(lines), lineCount);
}

void QPainter_DrawLines4(void* ptr, void* pointPairs, int lineCount){
	static_cast<QPainter*>(ptr)->drawLines(static_cast<QPoint*>(pointPairs), lineCount);
}

void QPainter_DrawLines3(void* ptr, void* pointPairs, int lineCount){
	static_cast<QPainter*>(ptr)->drawLines(static_cast<QPointF*>(pointPairs), lineCount);
}

void QPainter_DrawPath(void* ptr, void* path){
	static_cast<QPainter*>(ptr)->drawPath(*static_cast<QPainterPath*>(path));
}

void QPainter_DrawPicture2(void* ptr, void* point, void* picture){
	static_cast<QPainter*>(ptr)->drawPicture(*static_cast<QPoint*>(point), *static_cast<QPicture*>(picture));
}

void QPainter_DrawPicture3(void* ptr, int x, int y, void* picture){
	static_cast<QPainter*>(ptr)->drawPicture(x, y, *static_cast<QPicture*>(picture));
}

void QPainter_DrawPie2(void* ptr, void* rectangle, int startAngle, int spanAngle){
	static_cast<QPainter*>(ptr)->drawPie(*static_cast<QRect*>(rectangle), startAngle, spanAngle);
}

void QPainter_DrawPie3(void* ptr, int x, int y, int width, int height, int startAngle, int spanAngle){
	static_cast<QPainter*>(ptr)->drawPie(x, y, width, height, startAngle, spanAngle);
}

void QPainter_DrawPixmap6(void* ptr, void* point, void* pixmap){
	static_cast<QPainter*>(ptr)->drawPixmap(*static_cast<QPoint*>(point), *static_cast<QPixmap*>(pixmap));
}

void QPainter_DrawPixmap4(void* ptr, void* point, void* pixmap, void* source){
	static_cast<QPainter*>(ptr)->drawPixmap(*static_cast<QPoint*>(point), *static_cast<QPixmap*>(pixmap), *static_cast<QRect*>(source));
}

void QPainter_DrawPixmap3(void* ptr, void* point, void* pixmap, void* source){
	static_cast<QPainter*>(ptr)->drawPixmap(*static_cast<QPointF*>(point), *static_cast<QPixmap*>(pixmap), *static_cast<QRectF*>(source));
}

void QPainter_DrawPixmap8(void* ptr, void* rectangle, void* pixmap){
	static_cast<QPainter*>(ptr)->drawPixmap(*static_cast<QRect*>(rectangle), *static_cast<QPixmap*>(pixmap));
}

void QPainter_DrawPixmap2(void* ptr, void* target, void* pixmap, void* source){
	static_cast<QPainter*>(ptr)->drawPixmap(*static_cast<QRect*>(target), *static_cast<QPixmap*>(pixmap), *static_cast<QRect*>(source));
}

void QPainter_DrawPixmap7(void* ptr, int x, int y, void* pixmap){
	static_cast<QPainter*>(ptr)->drawPixmap(x, y, *static_cast<QPixmap*>(pixmap));
}

void QPainter_DrawPixmap11(void* ptr, int x, int y, void* pixmap, int sx, int sy, int sw, int sh){
	static_cast<QPainter*>(ptr)->drawPixmap(x, y, *static_cast<QPixmap*>(pixmap), sx, sy, sw, sh);
}

void QPainter_DrawPixmap10(void* ptr, int x, int y, int w, int h, void* pixmap, int sx, int sy, int sw, int sh){
	static_cast<QPainter*>(ptr)->drawPixmap(x, y, w, h, *static_cast<QPixmap*>(pixmap), sx, sy, sw, sh);
}

void QPainter_DrawPixmap9(void* ptr, int x, int y, int width, int height, void* pixmap){
	static_cast<QPainter*>(ptr)->drawPixmap(x, y, width, height, *static_cast<QPixmap*>(pixmap));
}

void QPainter_DrawPoint2(void* ptr, void* position){
	static_cast<QPainter*>(ptr)->drawPoint(*static_cast<QPoint*>(position));
}

void QPainter_DrawPoint(void* ptr, void* position){
	static_cast<QPainter*>(ptr)->drawPoint(*static_cast<QPointF*>(position));
}

void QPainter_DrawPoint3(void* ptr, int x, int y){
	static_cast<QPainter*>(ptr)->drawPoint(x, y);
}

void QPainter_DrawPoints2(void* ptr, void* points, int pointCount){
	static_cast<QPainter*>(ptr)->drawPoints(static_cast<QPoint*>(points), pointCount);
}

void QPainter_DrawPoints(void* ptr, void* points, int pointCount){
	static_cast<QPainter*>(ptr)->drawPoints(static_cast<QPointF*>(points), pointCount);
}

void QPainter_DrawPoints4(void* ptr, void* points){
	static_cast<QPainter*>(ptr)->drawPoints(*static_cast<QPolygon*>(points));
}

void QPainter_DrawPoints3(void* ptr, void* points){
	static_cast<QPainter*>(ptr)->drawPoints(*static_cast<QPolygonF*>(points));
}

void QPainter_DrawPolygon2(void* ptr, void* points, int pointCount, int fillRule){
	static_cast<QPainter*>(ptr)->drawPolygon(static_cast<QPoint*>(points), pointCount, static_cast<Qt::FillRule>(fillRule));
}

void QPainter_DrawPolygon(void* ptr, void* points, int pointCount, int fillRule){
	static_cast<QPainter*>(ptr)->drawPolygon(static_cast<QPointF*>(points), pointCount, static_cast<Qt::FillRule>(fillRule));
}

void QPainter_DrawPolygon4(void* ptr, void* points, int fillRule){
	static_cast<QPainter*>(ptr)->drawPolygon(*static_cast<QPolygon*>(points), static_cast<Qt::FillRule>(fillRule));
}

void QPainter_DrawPolygon3(void* ptr, void* points, int fillRule){
	static_cast<QPainter*>(ptr)->drawPolygon(*static_cast<QPolygonF*>(points), static_cast<Qt::FillRule>(fillRule));
}

void QPainter_DrawPolyline2(void* ptr, void* points, int pointCount){
	static_cast<QPainter*>(ptr)->drawPolyline(static_cast<QPoint*>(points), pointCount);
}

void QPainter_DrawPolyline(void* ptr, void* points, int pointCount){
	static_cast<QPainter*>(ptr)->drawPolyline(static_cast<QPointF*>(points), pointCount);
}

void QPainter_DrawPolyline4(void* ptr, void* points){
	static_cast<QPainter*>(ptr)->drawPolyline(*static_cast<QPolygon*>(points));
}

void QPainter_DrawPolyline3(void* ptr, void* points){
	static_cast<QPainter*>(ptr)->drawPolyline(*static_cast<QPolygonF*>(points));
}

void QPainter_DrawRect2(void* ptr, void* rectangle){
	static_cast<QPainter*>(ptr)->drawRect(*static_cast<QRect*>(rectangle));
}

void QPainter_DrawRect(void* ptr, void* rectangle){
	static_cast<QPainter*>(ptr)->drawRect(*static_cast<QRectF*>(rectangle));
}

void QPainter_DrawRect3(void* ptr, int x, int y, int width, int height){
	static_cast<QPainter*>(ptr)->drawRect(x, y, width, height);
}

void QPainter_DrawRoundedRect2(void* ptr, void* rect, double xRadius, double yRadius, int mode){
	static_cast<QPainter*>(ptr)->drawRoundedRect(*static_cast<QRect*>(rect), static_cast<qreal>(xRadius), static_cast<qreal>(yRadius), static_cast<Qt::SizeMode>(mode));
}

void QPainter_DrawRoundedRect(void* ptr, void* rect, double xRadius, double yRadius, int mode){
	static_cast<QPainter*>(ptr)->drawRoundedRect(*static_cast<QRectF*>(rect), static_cast<qreal>(xRadius), static_cast<qreal>(yRadius), static_cast<Qt::SizeMode>(mode));
}

void QPainter_DrawRoundedRect3(void* ptr, int x, int y, int w, int h, double xRadius, double yRadius, int mode){
	static_cast<QPainter*>(ptr)->drawRoundedRect(x, y, w, h, static_cast<qreal>(xRadius), static_cast<qreal>(yRadius), static_cast<Qt::SizeMode>(mode));
}

void QPainter_DrawStaticText2(void* ptr, void* topLeftPosition, void* staticText){
	static_cast<QPainter*>(ptr)->drawStaticText(*static_cast<QPoint*>(topLeftPosition), *static_cast<QStaticText*>(staticText));
}

void QPainter_DrawStaticText(void* ptr, void* topLeftPosition, void* staticText){
	static_cast<QPainter*>(ptr)->drawStaticText(*static_cast<QPointF*>(topLeftPosition), *static_cast<QStaticText*>(staticText));
}

void QPainter_DrawStaticText3(void* ptr, int left, int top, void* staticText){
	static_cast<QPainter*>(ptr)->drawStaticText(left, top, *static_cast<QStaticText*>(staticText));
}

void QPainter_DrawText3(void* ptr, void* position, char* text){
	static_cast<QPainter*>(ptr)->drawText(*static_cast<QPoint*>(position), QString(text));
}

void QPainter_DrawText6(void* ptr, int x, int y, char* text){
	static_cast<QPainter*>(ptr)->drawText(x, y, QString(text));
}

void QPainter_DrawText7(void* ptr, int x, int y, int width, int height, int flags, char* text, void* boundingRect){
	static_cast<QPainter*>(ptr)->drawText(x, y, width, height, flags, QString(text), static_cast<QRect*>(boundingRect));
}

void QPainter_DrawTiledPixmap2(void* ptr, void* rectangle, void* pixmap, void* position){
	static_cast<QPainter*>(ptr)->drawTiledPixmap(*static_cast<QRect*>(rectangle), *static_cast<QPixmap*>(pixmap), *static_cast<QPoint*>(position));
}

void QPainter_DrawTiledPixmap3(void* ptr, int x, int y, int width, int height, void* pixmap, int sx, int sy){
	static_cast<QPainter*>(ptr)->drawTiledPixmap(x, y, width, height, *static_cast<QPixmap*>(pixmap), sx, sy);
}

int QPainter_End(void* ptr){
	return static_cast<QPainter*>(ptr)->end();
}

void QPainter_EndNativePainting(void* ptr){
	static_cast<QPainter*>(ptr)->endNativePainting();
}

void QPainter_EraseRect2(void* ptr, void* rectangle){
	static_cast<QPainter*>(ptr)->eraseRect(*static_cast<QRect*>(rectangle));
}

void QPainter_EraseRect3(void* ptr, int x, int y, int width, int height){
	static_cast<QPainter*>(ptr)->eraseRect(x, y, width, height);
}

void QPainter_FillPath(void* ptr, void* path, void* brush){
	static_cast<QPainter*>(ptr)->fillPath(*static_cast<QPainterPath*>(path), *static_cast<QBrush*>(brush));
}

void QPainter_FillRect3(void* ptr, void* rectangle, int style){
	static_cast<QPainter*>(ptr)->fillRect(*static_cast<QRect*>(rectangle), static_cast<Qt::BrushStyle>(style));
}

void QPainter_FillRect11(void* ptr, void* rectangle, int color){
	static_cast<QPainter*>(ptr)->fillRect(*static_cast<QRect*>(rectangle), static_cast<Qt::GlobalColor>(color));
}

void QPainter_FillRect4(void* ptr, void* rectangle, int style){
	static_cast<QPainter*>(ptr)->fillRect(*static_cast<QRectF*>(rectangle), static_cast<Qt::BrushStyle>(style));
}

void QPainter_FillRect12(void* ptr, void* rectangle, int color){
	static_cast<QPainter*>(ptr)->fillRect(*static_cast<QRectF*>(rectangle), static_cast<Qt::GlobalColor>(color));
}

void QPainter_FillRect2(void* ptr, int x, int y, int width, int height, int style){
	static_cast<QPainter*>(ptr)->fillRect(x, y, width, height, static_cast<Qt::BrushStyle>(style));
}

void QPainter_FillRect10(void* ptr, int x, int y, int width, int height, int color){
	static_cast<QPainter*>(ptr)->fillRect(x, y, width, height, static_cast<Qt::GlobalColor>(color));
}

void QPainter_FillRect8(void* ptr, int x, int y, int width, int height, void* brush){
	static_cast<QPainter*>(ptr)->fillRect(x, y, width, height, *static_cast<QBrush*>(brush));
}

void QPainter_FillRect9(void* ptr, int x, int y, int width, int height, void* color){
	static_cast<QPainter*>(ptr)->fillRect(x, y, width, height, *static_cast<QColor*>(color));
}

int QPainter_HasClipping(void* ptr){
	return static_cast<QPainter*>(ptr)->hasClipping();
}

int QPainter_IsActive(void* ptr){
	return static_cast<QPainter*>(ptr)->isActive();
}

int QPainter_LayoutDirection(void* ptr){
	return static_cast<QPainter*>(ptr)->layoutDirection();
}

double QPainter_Opacity(void* ptr){
	return static_cast<double>(static_cast<QPainter*>(ptr)->opacity());
}

void* QPainter_PaintEngine(void* ptr){
	return static_cast<QPainter*>(ptr)->paintEngine();
}

int QPainter_RenderHints(void* ptr){
	return static_cast<QPainter*>(ptr)->renderHints();
}

void QPainter_ResetTransform(void* ptr){
	static_cast<QPainter*>(ptr)->resetTransform();
}

void QPainter_Restore(void* ptr){
	static_cast<QPainter*>(ptr)->restore();
}

void QPainter_Save(void* ptr){
	static_cast<QPainter*>(ptr)->save();
}

void QPainter_Scale(void* ptr, double sx, double sy){
	static_cast<QPainter*>(ptr)->scale(static_cast<qreal>(sx), static_cast<qreal>(sy));
}

void QPainter_SetBackgroundMode(void* ptr, int mode){
	static_cast<QPainter*>(ptr)->setBackgroundMode(static_cast<Qt::BGMode>(mode));
}

void QPainter_SetBrush2(void* ptr, int style){
	static_cast<QPainter*>(ptr)->setBrush(static_cast<Qt::BrushStyle>(style));
}

void QPainter_SetBrush(void* ptr, void* brush){
	static_cast<QPainter*>(ptr)->setBrush(*static_cast<QBrush*>(brush));
}

void QPainter_SetBrushOrigin2(void* ptr, void* position){
	static_cast<QPainter*>(ptr)->setBrushOrigin(*static_cast<QPoint*>(position));
}

void QPainter_SetBrushOrigin3(void* ptr, int x, int y){
	static_cast<QPainter*>(ptr)->setBrushOrigin(x, y);
}

void QPainter_SetClipRect2(void* ptr, int x, int y, int width, int height, int operation){
	static_cast<QPainter*>(ptr)->setClipRect(x, y, width, height, static_cast<Qt::ClipOperation>(operation));
}

void QPainter_SetClipping(void* ptr, int enable){
	static_cast<QPainter*>(ptr)->setClipping(enable != 0);
}

void QPainter_SetCompositionMode(void* ptr, int mode){
	static_cast<QPainter*>(ptr)->setCompositionMode(static_cast<QPainter::CompositionMode>(mode));
}

void QPainter_SetFont(void* ptr, void* font){
	static_cast<QPainter*>(ptr)->setFont(*static_cast<QFont*>(font));
}

void QPainter_SetLayoutDirection(void* ptr, int direction){
	static_cast<QPainter*>(ptr)->setLayoutDirection(static_cast<Qt::LayoutDirection>(direction));
}

void QPainter_SetOpacity(void* ptr, double opacity){
	static_cast<QPainter*>(ptr)->setOpacity(static_cast<qreal>(opacity));
}

void QPainter_SetPen3(void* ptr, int style){
	static_cast<QPainter*>(ptr)->setPen(static_cast<Qt::PenStyle>(style));
}

void QPainter_SetPen2(void* ptr, void* color){
	static_cast<QPainter*>(ptr)->setPen(*static_cast<QColor*>(color));
}

void QPainter_SetPen(void* ptr, void* pen){
	static_cast<QPainter*>(ptr)->setPen(*static_cast<QPen*>(pen));
}

void QPainter_SetRenderHint(void* ptr, int hint, int on){
	static_cast<QPainter*>(ptr)->setRenderHint(static_cast<QPainter::RenderHint>(hint), on != 0);
}

void QPainter_SetRenderHints(void* ptr, int hints, int on){
	static_cast<QPainter*>(ptr)->setRenderHints(static_cast<QPainter::RenderHint>(hints), on != 0);
}

void QPainter_SetTransform(void* ptr, void* transform, int combine){
	static_cast<QPainter*>(ptr)->setTransform(*static_cast<QTransform*>(transform), combine != 0);
}

void QPainter_SetViewTransformEnabled(void* ptr, int enable){
	static_cast<QPainter*>(ptr)->setViewTransformEnabled(enable != 0);
}

void QPainter_SetViewport2(void* ptr, int x, int y, int width, int height){
	static_cast<QPainter*>(ptr)->setViewport(x, y, width, height);
}

void QPainter_SetWindow2(void* ptr, int x, int y, int width, int height){
	static_cast<QPainter*>(ptr)->setWindow(x, y, width, height);
}

void QPainter_SetWorldMatrixEnabled(void* ptr, int enable){
	static_cast<QPainter*>(ptr)->setWorldMatrixEnabled(enable != 0);
}

void QPainter_SetWorldTransform(void* ptr, void* matrix, int combine){
	static_cast<QPainter*>(ptr)->setWorldTransform(*static_cast<QTransform*>(matrix), combine != 0);
}

void QPainter_Shear(void* ptr, double sh, double sv){
	static_cast<QPainter*>(ptr)->shear(static_cast<qreal>(sh), static_cast<qreal>(sv));
}

void QPainter_StrokePath(void* ptr, void* path, void* pen){
	static_cast<QPainter*>(ptr)->strokePath(*static_cast<QPainterPath*>(path), *static_cast<QPen*>(pen));
}

int QPainter_TestRenderHint(void* ptr, int hint){
	return static_cast<QPainter*>(ptr)->testRenderHint(static_cast<QPainter::RenderHint>(hint));
}

void QPainter_Translate2(void* ptr, void* offset){
	static_cast<QPainter*>(ptr)->translate(*static_cast<QPoint*>(offset));
}

void QPainter_Translate(void* ptr, void* offset){
	static_cast<QPainter*>(ptr)->translate(*static_cast<QPointF*>(offset));
}

void QPainter_Translate3(void* ptr, double dx, double dy){
	static_cast<QPainter*>(ptr)->translate(static_cast<qreal>(dx), static_cast<qreal>(dy));
}

int QPainter_ViewTransformEnabled(void* ptr){
	return static_cast<QPainter*>(ptr)->viewTransformEnabled();
}

int QPainter_WorldMatrixEnabled(void* ptr){
	return static_cast<QPainter*>(ptr)->worldMatrixEnabled();
}

void QPainter_DestroyQPainter(void* ptr){
	static_cast<QPainter*>(ptr)->~QPainter();
}

