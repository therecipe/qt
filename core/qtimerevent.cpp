#include "qtimerevent.h"
#include <QModelIndex>
#include <QTimer>
#include <QTime>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QTimerEvent>
#include "_cgo_export.h"

class MyQTimerEvent: public QTimerEvent {
public:
};

QtObjectPtr QTimerEvent_NewQTimerEvent(int timerId){
	return new QTimerEvent(timerId);
}

int QTimerEvent_TimerId(QtObjectPtr ptr){
	return static_cast<QTimerEvent*>(ptr)->timerId();
}

