#include "qscrollprepareevent.h"
#include <QModelIndex>
#include <QRect>
#include <QRectF>
#include <QSize>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QSizeF>
#include <QPoint>
#include <QPointF>
#include <QScrollPrepareEvent>
#include "_cgo_export.h"

class MyQScrollPrepareEvent: public QScrollPrepareEvent {
public:
};

void* QScrollPrepareEvent_NewQScrollPrepareEvent(void* startPos){
	return new QScrollPrepareEvent(*static_cast<QPointF*>(startPos));
}

void QScrollPrepareEvent_SetContentPos(void* ptr, void* pos){
	static_cast<QScrollPrepareEvent*>(ptr)->setContentPos(*static_cast<QPointF*>(pos));
}

void QScrollPrepareEvent_SetContentPosRange(void* ptr, void* rect){
	static_cast<QScrollPrepareEvent*>(ptr)->setContentPosRange(*static_cast<QRectF*>(rect));
}

void QScrollPrepareEvent_SetViewportSize(void* ptr, void* size){
	static_cast<QScrollPrepareEvent*>(ptr)->setViewportSize(*static_cast<QSizeF*>(size));
}

void QScrollPrepareEvent_DestroyQScrollPrepareEvent(void* ptr){
	static_cast<QScrollPrepareEvent*>(ptr)->~QScrollPrepareEvent();
}

