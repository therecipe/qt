#include "qpainterpathstroker.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QPen>
#include <QPainterPath>
#include <QPainter>
#include <QString>
#include <QPainterPathStroker>
#include "_cgo_export.h"

class MyQPainterPathStroker: public QPainterPathStroker {
public:
};

void* QPainterPathStroker_NewQPainterPathStroker(){
	return new QPainterPathStroker();
}

void* QPainterPathStroker_NewQPainterPathStroker2(void* pen){
	return new QPainterPathStroker(*static_cast<QPen*>(pen));
}

int QPainterPathStroker_CapStyle(void* ptr){
	return static_cast<QPainterPathStroker*>(ptr)->capStyle();
}

double QPainterPathStroker_CurveThreshold(void* ptr){
	return static_cast<double>(static_cast<QPainterPathStroker*>(ptr)->curveThreshold());
}

double QPainterPathStroker_DashOffset(void* ptr){
	return static_cast<double>(static_cast<QPainterPathStroker*>(ptr)->dashOffset());
}

int QPainterPathStroker_JoinStyle(void* ptr){
	return static_cast<QPainterPathStroker*>(ptr)->joinStyle();
}

double QPainterPathStroker_MiterLimit(void* ptr){
	return static_cast<double>(static_cast<QPainterPathStroker*>(ptr)->miterLimit());
}

void QPainterPathStroker_SetCapStyle(void* ptr, int style){
	static_cast<QPainterPathStroker*>(ptr)->setCapStyle(static_cast<Qt::PenCapStyle>(style));
}

void QPainterPathStroker_SetCurveThreshold(void* ptr, double threshold){
	static_cast<QPainterPathStroker*>(ptr)->setCurveThreshold(static_cast<qreal>(threshold));
}

void QPainterPathStroker_SetDashOffset(void* ptr, double offset){
	static_cast<QPainterPathStroker*>(ptr)->setDashOffset(static_cast<qreal>(offset));
}

void QPainterPathStroker_SetDashPattern(void* ptr, int style){
	static_cast<QPainterPathStroker*>(ptr)->setDashPattern(static_cast<Qt::PenStyle>(style));
}

void QPainterPathStroker_SetJoinStyle(void* ptr, int style){
	static_cast<QPainterPathStroker*>(ptr)->setJoinStyle(static_cast<Qt::PenJoinStyle>(style));
}

void QPainterPathStroker_SetMiterLimit(void* ptr, double limit){
	static_cast<QPainterPathStroker*>(ptr)->setMiterLimit(static_cast<qreal>(limit));
}

void QPainterPathStroker_SetWidth(void* ptr, double width){
	static_cast<QPainterPathStroker*>(ptr)->setWidth(static_cast<qreal>(width));
}

double QPainterPathStroker_Width(void* ptr){
	return static_cast<double>(static_cast<QPainterPathStroker*>(ptr)->width());
}

void QPainterPathStroker_DestroyQPainterPathStroker(void* ptr){
	static_cast<QPainterPathStroker*>(ptr)->~QPainterPathStroker();
}

