#include "qthread.h"
#include <QEvent>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMetaObject>
#include <QObject>
#include <QAbstractEventDispatcher>
#include <QThread>
#include "_cgo_export.h"

class MyQThread: public QThread {
public:
void Signal_Finished(){callbackQThreadFinished(this->objectName().toUtf8().data());};
void Signal_Started(){callbackQThreadStarted(this->objectName().toUtf8().data());};
};

void QThread_SetPriority(QtObjectPtr ptr, int priority){
	static_cast<QThread*>(ptr)->setPriority(static_cast<QThread::Priority>(priority));
}

QtObjectPtr QThread_NewQThread(QtObjectPtr parent){
	return new QThread(static_cast<QObject*>(parent));
}

QtObjectPtr QThread_QThread_CurrentThread(){
	return QThread::currentThread();
}

int QThread_Event(QtObjectPtr ptr, QtObjectPtr event){
	return static_cast<QThread*>(ptr)->event(static_cast<QEvent*>(event));
}

QtObjectPtr QThread_EventDispatcher(QtObjectPtr ptr){
	return static_cast<QThread*>(ptr)->eventDispatcher();
}

void QThread_Exit(QtObjectPtr ptr, int returnCode){
	static_cast<QThread*>(ptr)->exit(returnCode);
}

void QThread_ConnectFinished(QtObjectPtr ptr){
	QObject::connect(static_cast<QThread*>(ptr), &QThread::finished, static_cast<MyQThread*>(ptr), static_cast<void (MyQThread::*)()>(&MyQThread::Signal_Finished));;
}

void QThread_DisconnectFinished(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QThread*>(ptr), &QThread::finished, static_cast<MyQThread*>(ptr), static_cast<void (MyQThread::*)()>(&MyQThread::Signal_Finished));;
}

int QThread_IsFinished(QtObjectPtr ptr){
	return static_cast<QThread*>(ptr)->isFinished();
}

int QThread_IsInterruptionRequested(QtObjectPtr ptr){
	return static_cast<QThread*>(ptr)->isInterruptionRequested();
}

int QThread_IsRunning(QtObjectPtr ptr){
	return static_cast<QThread*>(ptr)->isRunning();
}

int QThread_LoopLevel(QtObjectPtr ptr){
	return static_cast<QThread*>(ptr)->loopLevel();
}

int QThread_Priority(QtObjectPtr ptr){
	return static_cast<QThread*>(ptr)->priority();
}

void QThread_Quit(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QThread*>(ptr), "quit");
}

void QThread_RequestInterruption(QtObjectPtr ptr){
	static_cast<QThread*>(ptr)->requestInterruption();
}

void QThread_SetEventDispatcher(QtObjectPtr ptr, QtObjectPtr eventDispatcher){
	static_cast<QThread*>(ptr)->setEventDispatcher(static_cast<QAbstractEventDispatcher*>(eventDispatcher));
}

void QThread_ConnectStarted(QtObjectPtr ptr){
	QObject::connect(static_cast<QThread*>(ptr), &QThread::started, static_cast<MyQThread*>(ptr), static_cast<void (MyQThread::*)()>(&MyQThread::Signal_Started));;
}

void QThread_DisconnectStarted(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QThread*>(ptr), &QThread::started, static_cast<MyQThread*>(ptr), static_cast<void (MyQThread::*)()>(&MyQThread::Signal_Started));;
}

void QThread_DestroyQThread(QtObjectPtr ptr){
	static_cast<QThread*>(ptr)->~QThread();
}

int QThread_QThread_IdealThreadCount(){
	return QThread::idealThreadCount();
}

void QThread_Start(QtObjectPtr ptr, int priority){
	QMetaObject::invokeMethod(static_cast<QThread*>(ptr), "start", Q_ARG(QThread::Priority, static_cast<QThread::Priority>(priority)));
}

void QThread_Terminate(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QThread*>(ptr), "terminate");
}

void QThread_QThread_YieldCurrentThread(){
	QThread::yieldCurrentThread();
}

