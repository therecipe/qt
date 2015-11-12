#include "qeventloop.h"
#include <QEvent>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QMetaObject>
#include <QEventLoop>
#include "_cgo_export.h"

class MyQEventLoop: public QEventLoop {
public:
};

void* QEventLoop_NewQEventLoop(void* parent){
	return new QEventLoop(static_cast<QObject*>(parent));
}

int QEventLoop_Event(void* ptr, void* event){
	return static_cast<QEventLoop*>(ptr)->event(static_cast<QEvent*>(event));
}

int QEventLoop_Exec(void* ptr, int flags){
	return static_cast<QEventLoop*>(ptr)->exec(static_cast<QEventLoop::ProcessEventsFlag>(flags));
}

void QEventLoop_Exit(void* ptr, int returnCode){
	static_cast<QEventLoop*>(ptr)->exit(returnCode);
}

int QEventLoop_IsRunning(void* ptr){
	return static_cast<QEventLoop*>(ptr)->isRunning();
}

int QEventLoop_ProcessEvents(void* ptr, int flags){
	return static_cast<QEventLoop*>(ptr)->processEvents(static_cast<QEventLoop::ProcessEventsFlag>(flags));
}

void QEventLoop_ProcessEvents2(void* ptr, int flags, int maxTime){
	static_cast<QEventLoop*>(ptr)->processEvents(static_cast<QEventLoop::ProcessEventsFlag>(flags), maxTime);
}

void QEventLoop_Quit(void* ptr){
	QMetaObject::invokeMethod(static_cast<QEventLoop*>(ptr), "quit");
}

void QEventLoop_WakeUp(void* ptr){
	static_cast<QEventLoop*>(ptr)->wakeUp();
}

void QEventLoop_DestroyQEventLoop(void* ptr){
	static_cast<QEventLoop*>(ptr)->~QEventLoop();
}

