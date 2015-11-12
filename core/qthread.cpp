#include "qthread.h"
#include <QUrl>
#include <QModelIndex>
#include <QMetaObject>
#include <QEvent>
#include <QObject>
#include <QAbstractEventDispatcher>
#include <QString>
#include <QVariant>
#include <QThread>
#include "_cgo_export.h"

class MyQThread: public QThread {
public:
void Signal_Finished(){callbackQThreadFinished(this->objectName().toUtf8().data());};
void Signal_Started(){callbackQThreadStarted(this->objectName().toUtf8().data());};
};

void QThread_SetPriority(void* ptr, int priority){
	static_cast<QThread*>(ptr)->setPriority(static_cast<QThread::Priority>(priority));
}

void* QThread_NewQThread(void* parent){
	return new QThread(static_cast<QObject*>(parent));
}

void* QThread_QThread_CurrentThread(){
	return QThread::currentThread();
}

int QThread_Event(void* ptr, void* event){
	return static_cast<QThread*>(ptr)->event(static_cast<QEvent*>(event));
}

void* QThread_EventDispatcher(void* ptr){
	return static_cast<QThread*>(ptr)->eventDispatcher();
}

void QThread_Exit(void* ptr, int returnCode){
	static_cast<QThread*>(ptr)->exit(returnCode);
}

void QThread_ConnectFinished(void* ptr){
	QObject::connect(static_cast<QThread*>(ptr), &QThread::finished, static_cast<MyQThread*>(ptr), static_cast<void (MyQThread::*)()>(&MyQThread::Signal_Finished));;
}

void QThread_DisconnectFinished(void* ptr){
	QObject::disconnect(static_cast<QThread*>(ptr), &QThread::finished, static_cast<MyQThread*>(ptr), static_cast<void (MyQThread::*)()>(&MyQThread::Signal_Finished));;
}

int QThread_IsFinished(void* ptr){
	return static_cast<QThread*>(ptr)->isFinished();
}

int QThread_IsInterruptionRequested(void* ptr){
	return static_cast<QThread*>(ptr)->isInterruptionRequested();
}

int QThread_IsRunning(void* ptr){
	return static_cast<QThread*>(ptr)->isRunning();
}

int QThread_LoopLevel(void* ptr){
	return static_cast<QThread*>(ptr)->loopLevel();
}

int QThread_Priority(void* ptr){
	return static_cast<QThread*>(ptr)->priority();
}

void QThread_Quit(void* ptr){
	QMetaObject::invokeMethod(static_cast<QThread*>(ptr), "quit");
}

void QThread_RequestInterruption(void* ptr){
	static_cast<QThread*>(ptr)->requestInterruption();
}

void QThread_SetEventDispatcher(void* ptr, void* eventDispatcher){
	static_cast<QThread*>(ptr)->setEventDispatcher(static_cast<QAbstractEventDispatcher*>(eventDispatcher));
}

void QThread_ConnectStarted(void* ptr){
	QObject::connect(static_cast<QThread*>(ptr), &QThread::started, static_cast<MyQThread*>(ptr), static_cast<void (MyQThread::*)()>(&MyQThread::Signal_Started));;
}

void QThread_DisconnectStarted(void* ptr){
	QObject::disconnect(static_cast<QThread*>(ptr), &QThread::started, static_cast<MyQThread*>(ptr), static_cast<void (MyQThread::*)()>(&MyQThread::Signal_Started));;
}

void QThread_DestroyQThread(void* ptr){
	static_cast<QThread*>(ptr)->~QThread();
}

int QThread_QThread_IdealThreadCount(){
	return QThread::idealThreadCount();
}

void QThread_Start(void* ptr, int priority){
	QMetaObject::invokeMethod(static_cast<QThread*>(ptr), "start", Q_ARG(QThread::Priority, static_cast<QThread::Priority>(priority)));
}

void QThread_Terminate(void* ptr){
	QMetaObject::invokeMethod(static_cast<QThread*>(ptr), "terminate");
}

void QThread_QThread_YieldCurrentThread(){
	QThread::yieldCurrentThread();
}

