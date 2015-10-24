#include "qbasictimer.h"
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QBasicTimer>
#include "_cgo_export.h"

class MyQBasicTimer: public QBasicTimer {
public:
};

void QBasicTimer_Start(QtObjectPtr ptr, int msec, QtObjectPtr object){
	static_cast<QBasicTimer*>(ptr)->start(msec, static_cast<QObject*>(object));
}

QtObjectPtr QBasicTimer_NewQBasicTimer(){
	return new QBasicTimer();
}

int QBasicTimer_IsActive(QtObjectPtr ptr){
	return static_cast<QBasicTimer*>(ptr)->isActive();
}

void QBasicTimer_Start2(QtObjectPtr ptr, int msec, int timerType, QtObjectPtr obj){
	static_cast<QBasicTimer*>(ptr)->start(msec, static_cast<Qt::TimerType>(timerType), static_cast<QObject*>(obj));
}

void QBasicTimer_Stop(QtObjectPtr ptr){
	static_cast<QBasicTimer*>(ptr)->stop();
}

int QBasicTimer_TimerId(QtObjectPtr ptr){
	return static_cast<QBasicTimer*>(ptr)->timerId();
}

void QBasicTimer_DestroyQBasicTimer(QtObjectPtr ptr){
	static_cast<QBasicTimer*>(ptr)->~QBasicTimer();
}

