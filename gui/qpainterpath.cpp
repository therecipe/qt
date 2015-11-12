#include "qpainterpath.h"
#include <QString>
#include <QUrl>
#include <QPainter>
#include <QPolygonF>
#include <QPoint>
#include <QPolygon>
#include <QVariant>
#include <QModelIndex>
#include <QRegion>
#include <QPointF>
#include <QRectF>
#include <QFont>
#include <QRect>
#include <QPainterPath>
#include "_cgo_export.h"

class MyQPainterPath: public QPainterPath {
public:
};

void* QPainterPath_NewQPainterPath3(void* path){
	return new QPainterPath(*static_cast<QPainterPath*>(path));
}

void QPainterPath_AddEllipse(void* ptr, void* boundingRectangle){
	static_cast<QPainterPath*>(ptr)->addEllipse(*static_cast<QRectF*>(boundingRectangle));
}

void QPainterPath_AddPath(void* ptr, void* path){
	static_cast<QPainterPath*>(ptr)->addPath(*static_cast<QPainterPath*>(path));
}

void QPainterPath_AddRect(void* ptr, void* rectangle){
	static_cast<QPainterPath*>(ptr)->addRect(*static_cast<QRectF*>(rectangle));
}

void QPainterPath_AddText(void* ptr, void* point, void* font, char* text){
	static_cast<QPainterPath*>(ptr)->addText(*static_cast<QPointF*>(point), *static_cast<QFont*>(font), QString(text));
}

void QPainterPath_ArcMoveTo(void* ptr, void* rectangle, double angle){
	static_cast<QPainterPath*>(ptr)->arcMoveTo(*static_cast<QRectF*>(rectangle), static_cast<qreal>(angle));
}

void QPainterPath_ArcTo(void* ptr, void* rectangle, double startAngle, double sweepLength){
	static_cast<QPainterPath*>(ptr)->arcTo(*static_cast<QRectF*>(rectangle), static_cast<qreal>(startAngle), static_cast<qreal>(sweepLength));
}

void QPainterPath_ConnectPath(void* ptr, void* path){
	static_cast<QPainterPath*>(ptr)->connectPath(*static_cast<QPainterPath*>(path));
}

int QPainterPath_Contains(void* ptr, void* point){
	return static_cast<QPainterPath*>(ptr)->contains(*static_cast<QPointF*>(point));
}

int QPainterPath_Contains2(void* ptr, void* rectangle){
	return static_cast<QPainterPath*>(ptr)->contains(*static_cast<QRectF*>(rectangle));
}

void QPainterPath_CubicTo(void* ptr, void* c1, void* c2, void* endPoint){
	static_cast<QPainterPath*>(ptr)->cubicTo(*static_cast<QPointF*>(c1), *static_cast<QPointF*>(c2), *static_cast<QPointF*>(endPoint));
}

int QPainterPath_ElementCount(void* ptr){
	return static_cast<QPainterPath*>(ptr)->elementCount();
}

int QPainterPath_Intersects(void* ptr, void* rectangle){
	return static_cast<QPainterPath*>(ptr)->intersects(*static_cast<QRectF*>(rectangle));
}

int QPainterPath_IsEmpty(void* ptr){
	return static_cast<QPainterPath*>(ptr)->isEmpty();
}

void QPainterPath_LineTo(void* ptr, void* endPoint){
	static_cast<QPainterPath*>(ptr)->lineTo(*static_cast<QPointF*>(endPoint));
}

void QPainterPath_MoveTo(void* ptr, void* point){
	static_cast<QPainterPath*>(ptr)->moveTo(*static_cast<QPointF*>(point));
}

void QPainterPath_QuadTo(void* ptr, void* c, void* endPoint){
	static_cast<QPainterPath*>(ptr)->quadTo(*static_cast<QPointF*>(c), *static_cast<QPointF*>(endPoint));
}

void QPainterPath_SetElementPositionAt(void* ptr, int index, double x, double y){
	static_cast<QPainterPath*>(ptr)->setElementPositionAt(index, static_cast<qreal>(x), static_cast<qreal>(y));
}

void QPainterPath_SetFillRule(void* ptr, int fillRule){
	static_cast<QPainterPath*>(ptr)->setFillRule(static_cast<Qt::FillRule>(fillRule));
}

void* QPainterPath_NewQPainterPath(){
	return new QPainterPath();
}

void* QPainterPath_NewQPainterPath2(void* startPoint){
	return new QPainterPath(*static_cast<QPointF*>(startPoint));
}

void QPainterPath_AddEllipse3(void* ptr, void* center, double rx, double ry){
	static_cast<QPainterPath*>(ptr)->addEllipse(*static_cast<QPointF*>(center), static_cast<qreal>(rx), static_cast<qreal>(ry));
}

void QPainterPath_AddEllipse2(void* ptr, double x, double y, double width, double height){
	static_cast<QPainterPath*>(ptr)->addEllipse(static_cast<qreal>(x), static_cast<qreal>(y), static_cast<qreal>(width), static_cast<qreal>(height));
}

void QPainterPath_AddPolygon(void* ptr, void* polygon){
	static_cast<QPainterPath*>(ptr)->addPolygon(*static_cast<QPolygonF*>(polygon));
}

void QPainterPath_AddRect2(void* ptr, double x, double y, double width, double height){
	static_cast<QPainterPath*>(ptr)->addRect(static_cast<qreal>(x), static_cast<qreal>(y), static_cast<qreal>(width), static_cast<qreal>(height));
}

void QPainterPath_AddRegion(void* ptr, void* region){
	static_cast<QPainterPath*>(ptr)->addRegion(*static_cast<QRegion*>(region));
}

void QPainterPath_AddRoundedRect(void* ptr, void* rect, double xRadius, double yRadius, int mode){
	static_cast<QPainterPath*>(ptr)->addRoundedRect(*static_cast<QRectF*>(rect), static_cast<qreal>(xRadius), static_cast<qreal>(yRadius), static_cast<Qt::SizeMode>(mode));
}

void QPainterPath_AddRoundedRect2(void* ptr, double x, double y, double w, double h, double xRadius, double yRadius, int mode){
	static_cast<QPainterPath*>(ptr)->addRoundedRect(static_cast<qreal>(x), static_cast<qreal>(y), static_cast<qreal>(w), static_cast<qreal>(h), static_cast<qreal>(xRadius), static_cast<qreal>(yRadius), static_cast<Qt::SizeMode>(mode));
}

void QPainterPath_AddText2(void* ptr, double x, double y, void* font, char* text){
	static_cast<QPainterPath*>(ptr)->addText(static_cast<qreal>(x), static_cast<qreal>(y), *static_cast<QFont*>(font), QString(text));
}

double QPainterPath_AngleAtPercent(void* ptr, double t){
	return static_cast<double>(static_cast<QPainterPath*>(ptr)->angleAtPercent(static_cast<qreal>(t)));
}

void QPainterPath_ArcMoveTo2(void* ptr, double x, double y, double width, double height, double angle){
	static_cast<QPainterPath*>(ptr)->arcMoveTo(static_cast<qreal>(x), static_cast<qreal>(y), static_cast<qreal>(width), static_cast<qreal>(height), static_cast<qreal>(angle));
}

void QPainterPath_ArcTo2(void* ptr, double x, double y, double width, double height, double startAngle, double sweepLength){
	static_cast<QPainterPath*>(ptr)->arcTo(static_cast<qreal>(x), static_cast<qreal>(y), static_cast<qreal>(width), static_cast<qreal>(height), static_cast<qreal>(startAngle), static_cast<qreal>(sweepLength));
}

void QPainterPath_CloseSubpath(void* ptr){
	static_cast<QPainterPath*>(ptr)->closeSubpath();
}

int QPainterPath_Contains3(void* ptr, void* p){
	return static_cast<QPainterPath*>(ptr)->contains(*static_cast<QPainterPath*>(p));
}

void QPainterPath_CubicTo2(void* ptr, double c1X, double c1Y, double c2X, double c2Y, double endPointX, double endPointY){
	static_cast<QPainterPath*>(ptr)->cubicTo(static_cast<qreal>(c1X), static_cast<qreal>(c1Y), static_cast<qreal>(c2X), static_cast<qreal>(c2Y), static_cast<qreal>(endPointX), static_cast<qreal>(endPointY));
}

int QPainterPath_FillRule(void* ptr){
	return static_cast<QPainterPath*>(ptr)->fillRule();
}

int QPainterPath_Intersects2(void* ptr, void* p){
	return static_cast<QPainterPath*>(ptr)->intersects(*static_cast<QPainterPath*>(p));
}

double QPainterPath_Length(void* ptr){
	return static_cast<double>(static_cast<QPainterPath*>(ptr)->length());
}

void QPainterPath_LineTo2(void* ptr, double x, double y){
	static_cast<QPainterPath*>(ptr)->lineTo(static_cast<qreal>(x), static_cast<qreal>(y));
}

void QPainterPath_MoveTo2(void* ptr, double x, double y){
	static_cast<QPainterPath*>(ptr)->moveTo(static_cast<qreal>(x), static_cast<qreal>(y));
}

double QPainterPath_PercentAtLength(void* ptr, double len){
	return static_cast<double>(static_cast<QPainterPath*>(ptr)->percentAtLength(static_cast<qreal>(len)));
}

void QPainterPath_QuadTo2(void* ptr, double cx, double cy, double endPointX, double endPointY){
	static_cast<QPainterPath*>(ptr)->quadTo(static_cast<qreal>(cx), static_cast<qreal>(cy), static_cast<qreal>(endPointX), static_cast<qreal>(endPointY));
}

double QPainterPath_SlopeAtPercent(void* ptr, double t){
	return static_cast<double>(static_cast<QPainterPath*>(ptr)->slopeAtPercent(static_cast<qreal>(t)));
}

void QPainterPath_Swap(void* ptr, void* other){
	static_cast<QPainterPath*>(ptr)->swap(*static_cast<QPainterPath*>(other));
}

void QPainterPath_Translate2(void* ptr, void* offset){
	static_cast<QPainterPath*>(ptr)->translate(*static_cast<QPointF*>(offset));
}

void QPainterPath_Translate(void* ptr, double dx, double dy){
	static_cast<QPainterPath*>(ptr)->translate(static_cast<qreal>(dx), static_cast<qreal>(dy));
}

void QPainterPath_DestroyQPainterPath(void* ptr){
	static_cast<QPainterPath*>(ptr)->~QPainterPath();
}

