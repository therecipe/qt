#include "qabstracteventdispatcher.h"
#include <QObject>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSocketNotifier>
#include <QThread>
#include <QEventLoop>
#include <QAbstractNativeEventFilter>
#include <QEvent>
#include <QString>
#include <QAbstractEventDispatcher>
#include "_cgo_export.h"

class MyQAbstractEventDispatcher: public QAbstractEventDispatcher {
public:
void Signal_AboutToBlock(){callbackQAbstractEventDispatcherAboutToBlock(this->objectName().toUtf8().data());};
void Signal_Awake(){callbackQAbstractEventDispatcherAwake(this->objectName().toUtf8().data());};
};

void QAbstractEventDispatcher_ConnectAboutToBlock(QtObjectPtr ptr){
	QObject::connect(static_cast<QAbstractEventDispatcher*>(ptr), static_cast<void (QAbstractEventDispatcher::*)()>(&QAbstractEventDispatcher::aboutToBlock), static_cast<MyQAbstractEventDispatcher*>(ptr), static_cast<void (MyQAbstractEventDispatcher::*)()>(&MyQAbstractEventDispatcher::Signal_AboutToBlock));;
}

void QAbstractEventDispatcher_DisconnectAboutToBlock(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAbstractEventDispatcher*>(ptr), static_cast<void (QAbstractEventDispatcher::*)()>(&QAbstractEventDispatcher::aboutToBlock), static_cast<MyQAbstractEventDispatcher*>(ptr), static_cast<void (MyQAbstractEventDispatcher::*)()>(&MyQAbstractEventDispatcher::Signal_AboutToBlock));;
}

void QAbstractEventDispatcher_ConnectAwake(QtObjectPtr ptr){
	QObject::connect(static_cast<QAbstractEventDispatcher*>(ptr), static_cast<void (QAbstractEventDispatcher::*)()>(&QAbstractEventDispatcher::awake), static_cast<MyQAbstractEventDispatcher*>(ptr), static_cast<void (MyQAbstractEventDispatcher::*)()>(&MyQAbstractEventDispatcher::Signal_Awake));;
}

void QAbstractEventDispatcher_DisconnectAwake(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAbstractEventDispatcher*>(ptr), static_cast<void (QAbstractEventDispatcher::*)()>(&QAbstractEventDispatcher::awake), static_cast<MyQAbstractEventDispatcher*>(ptr), static_cast<void (MyQAbstractEventDispatcher::*)()>(&MyQAbstractEventDispatcher::Signal_Awake));;
}

void QAbstractEventDispatcher_Flush(QtObjectPtr ptr){
	static_cast<QAbstractEventDispatcher*>(ptr)->flush();
}

void QAbstractEventDispatcher_InstallNativeEventFilter(QtObjectPtr ptr, QtObjectPtr filterObj){
	static_cast<QAbstractEventDispatcher*>(ptr)->installNativeEventFilter(static_cast<QAbstractNativeEventFilter*>(filterObj));
}

QtObjectPtr QAbstractEventDispatcher_QAbstractEventDispatcher_Instance(QtObjectPtr thread){
	return QAbstractEventDispatcher::instance(static_cast<QThread*>(thread));
}

void QAbstractEventDispatcher_Interrupt(QtObjectPtr ptr){
	static_cast<QAbstractEventDispatcher*>(ptr)->interrupt();
}

int QAbstractEventDispatcher_ProcessEvents(QtObjectPtr ptr, int flags){
	return static_cast<QAbstractEventDispatcher*>(ptr)->processEvents(static_cast<QEventLoop::ProcessEventsFlag>(flags));
}

void QAbstractEventDispatcher_RegisterSocketNotifier(QtObjectPtr ptr, QtObjectPtr notifier){
	static_cast<QAbstractEventDispatcher*>(ptr)->registerSocketNotifier(static_cast<QSocketNotifier*>(notifier));
}

int QAbstractEventDispatcher_RemainingTime(QtObjectPtr ptr, int timerId){
	return static_cast<QAbstractEventDispatcher*>(ptr)->remainingTime(timerId);
}

void QAbstractEventDispatcher_RemoveNativeEventFilter(QtObjectPtr ptr, QtObjectPtr filter){
	static_cast<QAbstractEventDispatcher*>(ptr)->removeNativeEventFilter(static_cast<QAbstractNativeEventFilter*>(filter));
}

void QAbstractEventDispatcher_UnregisterSocketNotifier(QtObjectPtr ptr, QtObjectPtr notifier){
	static_cast<QAbstractEventDispatcher*>(ptr)->unregisterSocketNotifier(static_cast<QSocketNotifier*>(notifier));
}

int QAbstractEventDispatcher_UnregisterTimer(QtObjectPtr ptr, int timerId){
	return static_cast<QAbstractEventDispatcher*>(ptr)->unregisterTimer(timerId);
}

int QAbstractEventDispatcher_UnregisterTimers(QtObjectPtr ptr, QtObjectPtr object){
	return static_cast<QAbstractEventDispatcher*>(ptr)->unregisterTimers(static_cast<QObject*>(object));
}

void QAbstractEventDispatcher_WakeUp(QtObjectPtr ptr){
	static_cast<QAbstractEventDispatcher*>(ptr)->wakeUp();
}

void QAbstractEventDispatcher_DestroyQAbstractEventDispatcher(QtObjectPtr ptr){
	static_cast<QAbstractEventDispatcher*>(ptr)->~QAbstractEventDispatcher();
}

