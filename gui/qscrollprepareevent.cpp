#include "qscrollprepareevent.h"
#include <QUrl>
#include <QSizeF>
#include <QPoint>
#include <QString>
#include <QVariant>
#include <QModelIndex>
#include <QRect>
#include <QRectF>
#include <QPointF>
#include <QSize>
#include <QScrollPrepareEvent>
#include "_cgo_export.h"

class MyQScrollPrepareEvent: public QScrollPrepareEvent {
public:
};

QtObjectPtr QScrollPrepareEvent_NewQScrollPrepareEvent(QtObjectPtr startPos){
	return new QScrollPrepareEvent(*static_cast<QPointF*>(startPos));
}

void QScrollPrepareEvent_SetContentPos(QtObjectPtr ptr, QtObjectPtr pos){
	static_cast<QScrollPrepareEvent*>(ptr)->setContentPos(*static_cast<QPointF*>(pos));
}

void QScrollPrepareEvent_SetContentPosRange(QtObjectPtr ptr, QtObjectPtr rect){
	static_cast<QScrollPrepareEvent*>(ptr)->setContentPosRange(*static_cast<QRectF*>(rect));
}

void QScrollPrepareEvent_SetViewportSize(QtObjectPtr ptr, QtObjectPtr size){
	static_cast<QScrollPrepareEvent*>(ptr)->setViewportSize(*static_cast<QSizeF*>(size));
}

void QScrollPrepareEvent_DestroyQScrollPrepareEvent(QtObjectPtr ptr){
	static_cast<QScrollPrepareEvent*>(ptr)->~QScrollPrepareEvent();
}

