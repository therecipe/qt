#include "qeventloop.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QEvent>
#include <QMetaObject>
#include <QObject>
#include <QEventLoop>
#include "_cgo_export.h"

class MyQEventLoop: public QEventLoop {
public:
};

QtObjectPtr QEventLoop_NewQEventLoop(QtObjectPtr parent){
	return new QEventLoop(static_cast<QObject*>(parent));
}

int QEventLoop_Event(QtObjectPtr ptr, QtObjectPtr event){
	return static_cast<QEventLoop*>(ptr)->event(static_cast<QEvent*>(event));
}

int QEventLoop_Exec(QtObjectPtr ptr, int flags){
	return static_cast<QEventLoop*>(ptr)->exec(static_cast<QEventLoop::ProcessEventsFlag>(flags));
}

void QEventLoop_Exit(QtObjectPtr ptr, int returnCode){
	static_cast<QEventLoop*>(ptr)->exit(returnCode);
}

int QEventLoop_IsRunning(QtObjectPtr ptr){
	return static_cast<QEventLoop*>(ptr)->isRunning();
}

int QEventLoop_ProcessEvents(QtObjectPtr ptr, int flags){
	return static_cast<QEventLoop*>(ptr)->processEvents(static_cast<QEventLoop::ProcessEventsFlag>(flags));
}

void QEventLoop_ProcessEvents2(QtObjectPtr ptr, int flags, int maxTime){
	static_cast<QEventLoop*>(ptr)->processEvents(static_cast<QEventLoop::ProcessEventsFlag>(flags), maxTime);
}

void QEventLoop_Quit(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QEventLoop*>(ptr), "quit");
}

void QEventLoop_WakeUp(QtObjectPtr ptr){
	static_cast<QEventLoop*>(ptr)->wakeUp();
}

void QEventLoop_DestroyQEventLoop(QtObjectPtr ptr){
	static_cast<QEventLoop*>(ptr)->~QEventLoop();
}

