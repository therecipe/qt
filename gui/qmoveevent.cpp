#include "qmoveevent.h"
#include <QModelIndex>
#include <QPoint>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QMoveEvent>
#include "_cgo_export.h"

class MyQMoveEvent: public QMoveEvent {
public:
};

QtObjectPtr QMoveEvent_NewQMoveEvent(QtObjectPtr pos, QtObjectPtr oldPos){
	return new QMoveEvent(*static_cast<QPoint*>(pos), *static_cast<QPoint*>(oldPos));
}

