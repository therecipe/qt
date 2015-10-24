#include "qpainterpath.h"
#include <QUrl>
#include <QModelIndex>
#include <QRectF>
#include <QPointF>
#include <QPolygonF>
#include <QPoint>
#include <QFont>
#include <QString>
#include <QVariant>
#include <QPolygon>
#include <QRect>
#include <QRegion>
#include <QPainter>
#include <QPainterPath>
#include "_cgo_export.h"

class MyQPainterPath: public QPainterPath {
public:
};

QtObjectPtr QPainterPath_NewQPainterPath3(QtObjectPtr path){
	return new QPainterPath(*static_cast<QPainterPath*>(path));
}

void QPainterPath_AddEllipse(QtObjectPtr ptr, QtObjectPtr boundingRectangle){
	static_cast<QPainterPath*>(ptr)->addEllipse(*static_cast<QRectF*>(boundingRectangle));
}

void QPainterPath_AddPath(QtObjectPtr ptr, QtObjectPtr path){
	static_cast<QPainterPath*>(ptr)->addPath(*static_cast<QPainterPath*>(path));
}

void QPainterPath_AddRect(QtObjectPtr ptr, QtObjectPtr rectangle){
	static_cast<QPainterPath*>(ptr)->addRect(*static_cast<QRectF*>(rectangle));
}

void QPainterPath_AddText(QtObjectPtr ptr, QtObjectPtr point, QtObjectPtr font, char* text){
	static_cast<QPainterPath*>(ptr)->addText(*static_cast<QPointF*>(point), *static_cast<QFont*>(font), QString(text));
}

void QPainterPath_ConnectPath(QtObjectPtr ptr, QtObjectPtr path){
	static_cast<QPainterPath*>(ptr)->connectPath(*static_cast<QPainterPath*>(path));
}

int QPainterPath_Contains(QtObjectPtr ptr, QtObjectPtr point){
	return static_cast<QPainterPath*>(ptr)->contains(*static_cast<QPointF*>(point));
}

int QPainterPath_Contains2(QtObjectPtr ptr, QtObjectPtr rectangle){
	return static_cast<QPainterPath*>(ptr)->contains(*static_cast<QRectF*>(rectangle));
}

void QPainterPath_CubicTo(QtObjectPtr ptr, QtObjectPtr c1, QtObjectPtr c2, QtObjectPtr endPoint){
	static_cast<QPainterPath*>(ptr)->cubicTo(*static_cast<QPointF*>(c1), *static_cast<QPointF*>(c2), *static_cast<QPointF*>(endPoint));
}

int QPainterPath_ElementCount(QtObjectPtr ptr){
	return static_cast<QPainterPath*>(ptr)->elementCount();
}

int QPainterPath_Intersects(QtObjectPtr ptr, QtObjectPtr rectangle){
	return static_cast<QPainterPath*>(ptr)->intersects(*static_cast<QRectF*>(rectangle));
}

int QPainterPath_IsEmpty(QtObjectPtr ptr){
	return static_cast<QPainterPath*>(ptr)->isEmpty();
}

void QPainterPath_LineTo(QtObjectPtr ptr, QtObjectPtr endPoint){
	static_cast<QPainterPath*>(ptr)->lineTo(*static_cast<QPointF*>(endPoint));
}

void QPainterPath_MoveTo(QtObjectPtr ptr, QtObjectPtr point){
	static_cast<QPainterPath*>(ptr)->moveTo(*static_cast<QPointF*>(point));
}

void QPainterPath_QuadTo(QtObjectPtr ptr, QtObjectPtr c, QtObjectPtr endPoint){
	static_cast<QPainterPath*>(ptr)->quadTo(*static_cast<QPointF*>(c), *static_cast<QPointF*>(endPoint));
}

void QPainterPath_SetFillRule(QtObjectPtr ptr, int fillRule){
	static_cast<QPainterPath*>(ptr)->setFillRule(static_cast<Qt::FillRule>(fillRule));
}

QtObjectPtr QPainterPath_NewQPainterPath(){
	return new QPainterPath();
}

QtObjectPtr QPainterPath_NewQPainterPath2(QtObjectPtr startPoint){
	return new QPainterPath(*static_cast<QPointF*>(startPoint));
}

void QPainterPath_AddPolygon(QtObjectPtr ptr, QtObjectPtr polygon){
	static_cast<QPainterPath*>(ptr)->addPolygon(*static_cast<QPolygonF*>(polygon));
}

void QPainterPath_AddRegion(QtObjectPtr ptr, QtObjectPtr region){
	static_cast<QPainterPath*>(ptr)->addRegion(*static_cast<QRegion*>(region));
}

void QPainterPath_CloseSubpath(QtObjectPtr ptr){
	static_cast<QPainterPath*>(ptr)->closeSubpath();
}

int QPainterPath_Contains3(QtObjectPtr ptr, QtObjectPtr p){
	return static_cast<QPainterPath*>(ptr)->contains(*static_cast<QPainterPath*>(p));
}

int QPainterPath_FillRule(QtObjectPtr ptr){
	return static_cast<QPainterPath*>(ptr)->fillRule();
}

int QPainterPath_Intersects2(QtObjectPtr ptr, QtObjectPtr p){
	return static_cast<QPainterPath*>(ptr)->intersects(*static_cast<QPainterPath*>(p));
}

void QPainterPath_Swap(QtObjectPtr ptr, QtObjectPtr other){
	static_cast<QPainterPath*>(ptr)->swap(*static_cast<QPainterPath*>(other));
}

void QPainterPath_Translate2(QtObjectPtr ptr, QtObjectPtr offset){
	static_cast<QPainterPath*>(ptr)->translate(*static_cast<QPointF*>(offset));
}

void QPainterPath_DestroyQPainterPath(QtObjectPtr ptr){
	static_cast<QPainterPath*>(ptr)->~QPainterPath();
}

