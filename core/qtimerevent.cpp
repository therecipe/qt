#include "qtimerevent.h"
#include <QTimer>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QTime>
#include <QTimerEvent>
#include "_cgo_export.h"

class MyQTimerEvent: public QTimerEvent {
public:
};

void* QTimerEvent_NewQTimerEvent(int timerId){
	return new QTimerEvent(timerId);
}

int QTimerEvent_TimerId(void* ptr){
	return static_cast<QTimerEvent*>(ptr)->timerId();
}

