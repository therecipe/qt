#include "qpaintengine.h"
#include <QString>
#include <QPainterPath>
#include <QRectF>
#include <QImage>
#include <QPaintEngineState>
#include <QLine>
#include <QLineF>
#include <QRect>
#include <QPaintDevice>
#include <QPoint>
#include <QModelIndex>
#include <QUrl>
#include <QPixmap>
#include <QPointF>
#include <QPainter>
#include <QTextItem>
#include <QVariant>
#include <QPaintEngine>
#include "_cgo_export.h"

class MyQPaintEngine: public QPaintEngine {
public:
};

void QPaintEngine_DrawEllipse(QtObjectPtr ptr, QtObjectPtr rect){
	static_cast<QPaintEngine*>(ptr)->drawEllipse(*static_cast<QRectF*>(rect));
}

void QPaintEngine_DrawImage(QtObjectPtr ptr, QtObjectPtr rectangle, QtObjectPtr image, QtObjectPtr sr, int flags){
	static_cast<QPaintEngine*>(ptr)->drawImage(*static_cast<QRectF*>(rectangle), *static_cast<QImage*>(image), *static_cast<QRectF*>(sr), static_cast<Qt::ImageConversionFlag>(flags));
}

void QPaintEngine_DrawPolygon(QtObjectPtr ptr, QtObjectPtr points, int pointCount, int mode){
	static_cast<QPaintEngine*>(ptr)->drawPolygon(static_cast<QPointF*>(points), pointCount, static_cast<QPaintEngine::PolygonDrawMode>(mode));
}

int QPaintEngine_Begin(QtObjectPtr ptr, QtObjectPtr pdev){
	return static_cast<QPaintEngine*>(ptr)->begin(static_cast<QPaintDevice*>(pdev));
}

void QPaintEngine_DrawEllipse2(QtObjectPtr ptr, QtObjectPtr rect){
	static_cast<QPaintEngine*>(ptr)->drawEllipse(*static_cast<QRect*>(rect));
}

void QPaintEngine_DrawLines2(QtObjectPtr ptr, QtObjectPtr lines, int lineCount){
	static_cast<QPaintEngine*>(ptr)->drawLines(static_cast<QLine*>(lines), lineCount);
}

void QPaintEngine_DrawLines(QtObjectPtr ptr, QtObjectPtr lines, int lineCount){
	static_cast<QPaintEngine*>(ptr)->drawLines(static_cast<QLineF*>(lines), lineCount);
}

void QPaintEngine_DrawPath(QtObjectPtr ptr, QtObjectPtr path){
	static_cast<QPaintEngine*>(ptr)->drawPath(*static_cast<QPainterPath*>(path));
}

void QPaintEngine_DrawPixmap(QtObjectPtr ptr, QtObjectPtr r, QtObjectPtr pm, QtObjectPtr sr){
	static_cast<QPaintEngine*>(ptr)->drawPixmap(*static_cast<QRectF*>(r), *static_cast<QPixmap*>(pm), *static_cast<QRectF*>(sr));
}

void QPaintEngine_DrawPoints2(QtObjectPtr ptr, QtObjectPtr points, int pointCount){
	static_cast<QPaintEngine*>(ptr)->drawPoints(static_cast<QPoint*>(points), pointCount);
}

void QPaintEngine_DrawPoints(QtObjectPtr ptr, QtObjectPtr points, int pointCount){
	static_cast<QPaintEngine*>(ptr)->drawPoints(static_cast<QPointF*>(points), pointCount);
}

void QPaintEngine_DrawPolygon2(QtObjectPtr ptr, QtObjectPtr points, int pointCount, int mode){
	static_cast<QPaintEngine*>(ptr)->drawPolygon(static_cast<QPoint*>(points), pointCount, static_cast<QPaintEngine::PolygonDrawMode>(mode));
}

void QPaintEngine_DrawRects2(QtObjectPtr ptr, QtObjectPtr rects, int rectCount){
	static_cast<QPaintEngine*>(ptr)->drawRects(static_cast<QRect*>(rects), rectCount);
}

void QPaintEngine_DrawRects(QtObjectPtr ptr, QtObjectPtr rects, int rectCount){
	static_cast<QPaintEngine*>(ptr)->drawRects(static_cast<QRectF*>(rects), rectCount);
}

void QPaintEngine_DrawTextItem(QtObjectPtr ptr, QtObjectPtr p, QtObjectPtr textItem){
	static_cast<QPaintEngine*>(ptr)->drawTextItem(*static_cast<QPointF*>(p), *static_cast<QTextItem*>(textItem));
}

void QPaintEngine_DrawTiledPixmap(QtObjectPtr ptr, QtObjectPtr rect, QtObjectPtr pixmap, QtObjectPtr p){
	static_cast<QPaintEngine*>(ptr)->drawTiledPixmap(*static_cast<QRectF*>(rect), *static_cast<QPixmap*>(pixmap), *static_cast<QPointF*>(p));
}

int QPaintEngine_End(QtObjectPtr ptr){
	return static_cast<QPaintEngine*>(ptr)->end();
}

int QPaintEngine_HasFeature(QtObjectPtr ptr, int feature){
	return static_cast<QPaintEngine*>(ptr)->hasFeature(static_cast<QPaintEngine::PaintEngineFeature>(feature));
}

int QPaintEngine_IsActive(QtObjectPtr ptr){
	return static_cast<QPaintEngine*>(ptr)->isActive();
}

QtObjectPtr QPaintEngine_PaintDevice(QtObjectPtr ptr){
	return static_cast<QPaintEngine*>(ptr)->paintDevice();
}

QtObjectPtr QPaintEngine_Painter(QtObjectPtr ptr){
	return static_cast<QPaintEngine*>(ptr)->painter();
}

void QPaintEngine_SetActive(QtObjectPtr ptr, int state){
	static_cast<QPaintEngine*>(ptr)->setActive(state != 0);
}

int QPaintEngine_Type(QtObjectPtr ptr){
	return static_cast<QPaintEngine*>(ptr)->type();
}

void QPaintEngine_UpdateState(QtObjectPtr ptr, QtObjectPtr state){
	static_cast<QPaintEngine*>(ptr)->updateState(*static_cast<QPaintEngineState*>(state));
}

void QPaintEngine_DestroyQPaintEngine(QtObjectPtr ptr){
	static_cast<QPaintEngine*>(ptr)->~QPaintEngine();
}

