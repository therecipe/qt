#include "qpainterpathstroker.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QPainterPath>
#include <QPainter>
#include <QPen>
#include <QString>
#include <QPainterPathStroker>
#include "_cgo_export.h"

class MyQPainterPathStroker: public QPainterPathStroker {
public:
};

QtObjectPtr QPainterPathStroker_NewQPainterPathStroker(){
	return new QPainterPathStroker();
}

QtObjectPtr QPainterPathStroker_NewQPainterPathStroker2(QtObjectPtr pen){
	return new QPainterPathStroker(*static_cast<QPen*>(pen));
}

int QPainterPathStroker_CapStyle(QtObjectPtr ptr){
	return static_cast<QPainterPathStroker*>(ptr)->capStyle();
}

int QPainterPathStroker_JoinStyle(QtObjectPtr ptr){
	return static_cast<QPainterPathStroker*>(ptr)->joinStyle();
}

void QPainterPathStroker_SetCapStyle(QtObjectPtr ptr, int style){
	static_cast<QPainterPathStroker*>(ptr)->setCapStyle(static_cast<Qt::PenCapStyle>(style));
}

void QPainterPathStroker_SetDashPattern(QtObjectPtr ptr, int style){
	static_cast<QPainterPathStroker*>(ptr)->setDashPattern(static_cast<Qt::PenStyle>(style));
}

void QPainterPathStroker_SetJoinStyle(QtObjectPtr ptr, int style){
	static_cast<QPainterPathStroker*>(ptr)->setJoinStyle(static_cast<Qt::PenJoinStyle>(style));
}

void QPainterPathStroker_DestroyQPainterPathStroker(QtObjectPtr ptr){
	static_cast<QPainterPathStroker*>(ptr)->~QPainterPathStroker();
}

