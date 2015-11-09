#include "qmoveevent.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QPoint>
#include <QMoveEvent>
#include "_cgo_export.h"

class MyQMoveEvent: public QMoveEvent {
public:
};

void* QMoveEvent_NewQMoveEvent(void* pos, void* oldPos){
	return new QMoveEvent(*static_cast<QPoint*>(pos), *static_cast<QPoint*>(oldPos));
}

