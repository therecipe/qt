#include "qpaintengine.h"
#include <QImage>
#include <QString>
#include <QUrl>
#include <QPixmap>
#include <QPoint>
#include <QPointF>
#include <QPainter>
#include <QVariant>
#include <QModelIndex>
#include <QPaintEngineState>
#include <QRect>
#include <QTextItem>
#include <QPainterPath>
#include <QLineF>
#include <QPaintDevice>
#include <QRectF>
#include <QLine>
#include <QPaintEngine>
#include "_cgo_export.h"

class MyQPaintEngine: public QPaintEngine {
public:
};

void QPaintEngine_DrawEllipse(void* ptr, void* rect){
	static_cast<QPaintEngine*>(ptr)->drawEllipse(*static_cast<QRectF*>(rect));
}

void QPaintEngine_DrawImage(void* ptr, void* rectangle, void* image, void* sr, int flags){
	static_cast<QPaintEngine*>(ptr)->drawImage(*static_cast<QRectF*>(rectangle), *static_cast<QImage*>(image), *static_cast<QRectF*>(sr), static_cast<Qt::ImageConversionFlag>(flags));
}

void QPaintEngine_DrawPolygon(void* ptr, void* points, int pointCount, int mode){
	static_cast<QPaintEngine*>(ptr)->drawPolygon(static_cast<QPointF*>(points), pointCount, static_cast<QPaintEngine::PolygonDrawMode>(mode));
}

int QPaintEngine_Begin(void* ptr, void* pdev){
	return static_cast<QPaintEngine*>(ptr)->begin(static_cast<QPaintDevice*>(pdev));
}

void QPaintEngine_DrawEllipse2(void* ptr, void* rect){
	static_cast<QPaintEngine*>(ptr)->drawEllipse(*static_cast<QRect*>(rect));
}

void QPaintEngine_DrawLines2(void* ptr, void* lines, int lineCount){
	static_cast<QPaintEngine*>(ptr)->drawLines(static_cast<QLine*>(lines), lineCount);
}

void QPaintEngine_DrawLines(void* ptr, void* lines, int lineCount){
	static_cast<QPaintEngine*>(ptr)->drawLines(static_cast<QLineF*>(lines), lineCount);
}

void QPaintEngine_DrawPath(void* ptr, void* path){
	static_cast<QPaintEngine*>(ptr)->drawPath(*static_cast<QPainterPath*>(path));
}

void QPaintEngine_DrawPixmap(void* ptr, void* r, void* pm, void* sr){
	static_cast<QPaintEngine*>(ptr)->drawPixmap(*static_cast<QRectF*>(r), *static_cast<QPixmap*>(pm), *static_cast<QRectF*>(sr));
}

void QPaintEngine_DrawPoints2(void* ptr, void* points, int pointCount){
	static_cast<QPaintEngine*>(ptr)->drawPoints(static_cast<QPoint*>(points), pointCount);
}

void QPaintEngine_DrawPoints(void* ptr, void* points, int pointCount){
	static_cast<QPaintEngine*>(ptr)->drawPoints(static_cast<QPointF*>(points), pointCount);
}

void QPaintEngine_DrawPolygon2(void* ptr, void* points, int pointCount, int mode){
	static_cast<QPaintEngine*>(ptr)->drawPolygon(static_cast<QPoint*>(points), pointCount, static_cast<QPaintEngine::PolygonDrawMode>(mode));
}

void QPaintEngine_DrawRects2(void* ptr, void* rects, int rectCount){
	static_cast<QPaintEngine*>(ptr)->drawRects(static_cast<QRect*>(rects), rectCount);
}

void QPaintEngine_DrawRects(void* ptr, void* rects, int rectCount){
	static_cast<QPaintEngine*>(ptr)->drawRects(static_cast<QRectF*>(rects), rectCount);
}

void QPaintEngine_DrawTextItem(void* ptr, void* p, void* textItem){
	static_cast<QPaintEngine*>(ptr)->drawTextItem(*static_cast<QPointF*>(p), *static_cast<QTextItem*>(textItem));
}

void QPaintEngine_DrawTiledPixmap(void* ptr, void* rect, void* pixmap, void* p){
	static_cast<QPaintEngine*>(ptr)->drawTiledPixmap(*static_cast<QRectF*>(rect), *static_cast<QPixmap*>(pixmap), *static_cast<QPointF*>(p));
}

int QPaintEngine_End(void* ptr){
	return static_cast<QPaintEngine*>(ptr)->end();
}

int QPaintEngine_HasFeature(void* ptr, int feature){
	return static_cast<QPaintEngine*>(ptr)->hasFeature(static_cast<QPaintEngine::PaintEngineFeature>(feature));
}

int QPaintEngine_IsActive(void* ptr){
	return static_cast<QPaintEngine*>(ptr)->isActive();
}

void* QPaintEngine_PaintDevice(void* ptr){
	return static_cast<QPaintEngine*>(ptr)->paintDevice();
}

void* QPaintEngine_Painter(void* ptr){
	return static_cast<QPaintEngine*>(ptr)->painter();
}

void QPaintEngine_SetActive(void* ptr, int state){
	static_cast<QPaintEngine*>(ptr)->setActive(state != 0);
}

int QPaintEngine_Type(void* ptr){
	return static_cast<QPaintEngine*>(ptr)->type();
}

void QPaintEngine_UpdateState(void* ptr, void* state){
	static_cast<QPaintEngine*>(ptr)->updateState(*static_cast<QPaintEngineState*>(state));
}

void QPaintEngine_DestroyQPaintEngine(void* ptr){
	static_cast<QPaintEngine*>(ptr)->~QPaintEngine();
}

