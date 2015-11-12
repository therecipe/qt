#include "qpaintevent.h"
#include <QRegion>
#include <QRect>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QPaintEvent>
#include "_cgo_export.h"

class MyQPaintEvent: public QPaintEvent {
public:
};

void* QPaintEvent_NewQPaintEvent2(void* paintRect){
	return new QPaintEvent(*static_cast<QRect*>(paintRect));
}

void* QPaintEvent_NewQPaintEvent(void* paintRegion){
	return new QPaintEvent(*static_cast<QRegion*>(paintRegion));
}

void* QPaintEvent_Region(void* ptr){
	return new QRegion(static_cast<QPaintEvent*>(ptr)->region());
}

