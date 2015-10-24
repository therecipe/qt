#include "qtimer.h"
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMetaObject>
#include <QTime>
#include <QTimer>
#include "_cgo_export.h"

class MyQTimer: public QTimer {
public:
void Signal_Timeout(){callbackQTimerTimeout(this->objectName().toUtf8().data());};
};

int QTimer_RemainingTime(QtObjectPtr ptr){
	return static_cast<QTimer*>(ptr)->remainingTime();
}

void QTimer_SetInterval(QtObjectPtr ptr, int msec){
	static_cast<QTimer*>(ptr)->setInterval(msec);
}

QtObjectPtr QTimer_NewQTimer(QtObjectPtr parent){
	return new QTimer(static_cast<QObject*>(parent));
}

int QTimer_Interval(QtObjectPtr ptr){
	return static_cast<QTimer*>(ptr)->interval();
}

int QTimer_IsActive(QtObjectPtr ptr){
	return static_cast<QTimer*>(ptr)->isActive();
}

int QTimer_IsSingleShot(QtObjectPtr ptr){
	return static_cast<QTimer*>(ptr)->isSingleShot();
}

void QTimer_SetSingleShot(QtObjectPtr ptr, int singleShot){
	static_cast<QTimer*>(ptr)->setSingleShot(singleShot != 0);
}

void QTimer_SetTimerType(QtObjectPtr ptr, int atype){
	static_cast<QTimer*>(ptr)->setTimerType(static_cast<Qt::TimerType>(atype));
}

void QTimer_QTimer_SingleShot2(int msec, int timerType, QtObjectPtr receiver, char* member){
	QTimer::singleShot(msec, static_cast<Qt::TimerType>(timerType), static_cast<QObject*>(receiver), const_cast<const char*>(member));
}

void QTimer_QTimer_SingleShot(int msec, QtObjectPtr receiver, char* member){
	QTimer::singleShot(msec, static_cast<QObject*>(receiver), const_cast<const char*>(member));
}

void QTimer_Start2(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QTimer*>(ptr), "start");
}

void QTimer_Start(QtObjectPtr ptr, int msec){
	QMetaObject::invokeMethod(static_cast<QTimer*>(ptr), "start", Q_ARG(int, msec));
}

void QTimer_Stop(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QTimer*>(ptr), "stop");
}

void QTimer_ConnectTimeout(QtObjectPtr ptr){
	QObject::connect(static_cast<QTimer*>(ptr), &QTimer::timeout, static_cast<MyQTimer*>(ptr), static_cast<void (MyQTimer::*)()>(&MyQTimer::Signal_Timeout));;
}

void QTimer_DisconnectTimeout(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QTimer*>(ptr), &QTimer::timeout, static_cast<MyQTimer*>(ptr), static_cast<void (MyQTimer::*)()>(&MyQTimer::Signal_Timeout));;
}

int QTimer_TimerId(QtObjectPtr ptr){
	return static_cast<QTimer*>(ptr)->timerId();
}

int QTimer_TimerType(QtObjectPtr ptr){
	return static_cast<QTimer*>(ptr)->timerType();
}

void QTimer_DestroyQTimer(QtObjectPtr ptr){
	static_cast<QTimer*>(ptr)->~QTimer();
}

