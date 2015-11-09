#include "qtimer.h"
#include <QObject>
#include <QTime>
#include <QMetaObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QTimer>
#include "_cgo_export.h"

class MyQTimer: public QTimer {
public:
void Signal_Timeout(){callbackQTimerTimeout(this->objectName().toUtf8().data());};
};

int QTimer_RemainingTime(void* ptr){
	return static_cast<QTimer*>(ptr)->remainingTime();
}

void QTimer_SetInterval(void* ptr, int msec){
	static_cast<QTimer*>(ptr)->setInterval(msec);
}

void* QTimer_NewQTimer(void* parent){
	return new QTimer(static_cast<QObject*>(parent));
}

int QTimer_Interval(void* ptr){
	return static_cast<QTimer*>(ptr)->interval();
}

int QTimer_IsActive(void* ptr){
	return static_cast<QTimer*>(ptr)->isActive();
}

int QTimer_IsSingleShot(void* ptr){
	return static_cast<QTimer*>(ptr)->isSingleShot();
}

void QTimer_SetSingleShot(void* ptr, int singleShot){
	static_cast<QTimer*>(ptr)->setSingleShot(singleShot != 0);
}

void QTimer_SetTimerType(void* ptr, int atype){
	static_cast<QTimer*>(ptr)->setTimerType(static_cast<Qt::TimerType>(atype));
}

void QTimer_QTimer_SingleShot2(int msec, int timerType, void* receiver, char* member){
	QTimer::singleShot(msec, static_cast<Qt::TimerType>(timerType), static_cast<QObject*>(receiver), const_cast<const char*>(member));
}

void QTimer_QTimer_SingleShot(int msec, void* receiver, char* member){
	QTimer::singleShot(msec, static_cast<QObject*>(receiver), const_cast<const char*>(member));
}

void QTimer_Start2(void* ptr){
	QMetaObject::invokeMethod(static_cast<QTimer*>(ptr), "start");
}

void QTimer_Start(void* ptr, int msec){
	QMetaObject::invokeMethod(static_cast<QTimer*>(ptr), "start", Q_ARG(int, msec));
}

void QTimer_Stop(void* ptr){
	QMetaObject::invokeMethod(static_cast<QTimer*>(ptr), "stop");
}

void QTimer_ConnectTimeout(void* ptr){
	QObject::connect(static_cast<QTimer*>(ptr), &QTimer::timeout, static_cast<MyQTimer*>(ptr), static_cast<void (MyQTimer::*)()>(&MyQTimer::Signal_Timeout));;
}

void QTimer_DisconnectTimeout(void* ptr){
	QObject::disconnect(static_cast<QTimer*>(ptr), &QTimer::timeout, static_cast<MyQTimer*>(ptr), static_cast<void (MyQTimer::*)()>(&MyQTimer::Signal_Timeout));;
}

int QTimer_TimerId(void* ptr){
	return static_cast<QTimer*>(ptr)->timerId();
}

int QTimer_TimerType(void* ptr){
	return static_cast<QTimer*>(ptr)->timerType();
}

void QTimer_DestroyQTimer(void* ptr){
	static_cast<QTimer*>(ptr)->~QTimer();
}

