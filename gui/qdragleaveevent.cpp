#include "qdragleaveevent.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QDrag>
#include <QDragLeaveEvent>
#include "_cgo_export.h"

class MyQDragLeaveEvent: public QDragLeaveEvent {
public:
};

QtObjectPtr QDragLeaveEvent_NewQDragLeaveEvent(){
	return new QDragLeaveEvent();
}

