#include "qpaintevent.h"
#include <QRegion>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QRect>
#include <QPaintEvent>
#include "_cgo_export.h"

class MyQPaintEvent: public QPaintEvent {
public:
};

QtObjectPtr QPaintEvent_NewQPaintEvent2(QtObjectPtr paintRect){
	return new QPaintEvent(*static_cast<QRect*>(paintRect));
}

QtObjectPtr QPaintEvent_NewQPaintEvent(QtObjectPtr paintRegion){
	return new QPaintEvent(*static_cast<QRegion*>(paintRegion));
}

