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

void QBasicTimer_Start(void* ptr, int msec, void* object){
	static_cast<QBasicTimer*>(ptr)->start(msec, static_cast<QObject*>(object));
}

void* QBasicTimer_NewQBasicTimer(){
	return new QBasicTimer();
}

int QBasicTimer_IsActive(void* ptr){
	return static_cast<QBasicTimer*>(ptr)->isActive();
}

void QBasicTimer_Start2(void* ptr, int msec, int timerType, void* obj){
	static_cast<QBasicTimer*>(ptr)->start(msec, static_cast<Qt::TimerType>(timerType), static_cast<QObject*>(obj));
}

void QBasicTimer_Stop(void* ptr){
	static_cast<QBasicTimer*>(ptr)->stop();
}

int QBasicTimer_TimerId(void* ptr){
	return static_cast<QBasicTimer*>(ptr)->timerId();
}

void QBasicTimer_DestroyQBasicTimer(void* ptr){
	static_cast<QBasicTimer*>(ptr)->~QBasicTimer();
}

