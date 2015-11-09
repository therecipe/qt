#include "qabstracteventdispatcher.h"
#include <QThread>
#include <QAbstractNativeEventFilter>
#include <QSocketNotifier>
#include <QUrl>
#include <QObject>
#include <QModelIndex>
#include <QEventLoop>
#include <QEvent>
#include <QString>
#include <QVariant>
#include <QAbstractEventDispatcher>
#include "_cgo_export.h"

class MyQAbstractEventDispatcher: public QAbstractEventDispatcher {
public:
void Signal_AboutToBlock(){callbackQAbstractEventDispatcherAboutToBlock(this->objectName().toUtf8().data());};
void Signal_Awake(){callbackQAbstractEventDispatcherAwake(this->objectName().toUtf8().data());};
};

void QAbstractEventDispatcher_ConnectAboutToBlock(void* ptr){
	QObject::connect(static_cast<QAbstractEventDispatcher*>(ptr), static_cast<void (QAbstractEventDispatcher::*)()>(&QAbstractEventDispatcher::aboutToBlock), static_cast<MyQAbstractEventDispatcher*>(ptr), static_cast<void (MyQAbstractEventDispatcher::*)()>(&MyQAbstractEventDispatcher::Signal_AboutToBlock));;
}

void QAbstractEventDispatcher_DisconnectAboutToBlock(void* ptr){
	QObject::disconnect(static_cast<QAbstractEventDispatcher*>(ptr), static_cast<void (QAbstractEventDispatcher::*)()>(&QAbstractEventDispatcher::aboutToBlock), static_cast<MyQAbstractEventDispatcher*>(ptr), static_cast<void (MyQAbstractEventDispatcher::*)()>(&MyQAbstractEventDispatcher::Signal_AboutToBlock));;
}

void QAbstractEventDispatcher_ConnectAwake(void* ptr){
	QObject::connect(static_cast<QAbstractEventDispatcher*>(ptr), static_cast<void (QAbstractEventDispatcher::*)()>(&QAbstractEventDispatcher::awake), static_cast<MyQAbstractEventDispatcher*>(ptr), static_cast<void (MyQAbstractEventDispatcher::*)()>(&MyQAbstractEventDispatcher::Signal_Awake));;
}

void QAbstractEventDispatcher_DisconnectAwake(void* ptr){
	QObject::disconnect(static_cast<QAbstractEventDispatcher*>(ptr), static_cast<void (QAbstractEventDispatcher::*)()>(&QAbstractEventDispatcher::awake), static_cast<MyQAbstractEventDispatcher*>(ptr), static_cast<void (MyQAbstractEventDispatcher::*)()>(&MyQAbstractEventDispatcher::Signal_Awake));;
}

void QAbstractEventDispatcher_Flush(void* ptr){
	static_cast<QAbstractEventDispatcher*>(ptr)->flush();
}

void QAbstractEventDispatcher_InstallNativeEventFilter(void* ptr, void* filterObj){
	static_cast<QAbstractEventDispatcher*>(ptr)->installNativeEventFilter(static_cast<QAbstractNativeEventFilter*>(filterObj));
}

void* QAbstractEventDispatcher_QAbstractEventDispatcher_Instance(void* thread){
	return QAbstractEventDispatcher::instance(static_cast<QThread*>(thread));
}

void QAbstractEventDispatcher_Interrupt(void* ptr){
	static_cast<QAbstractEventDispatcher*>(ptr)->interrupt();
}

int QAbstractEventDispatcher_ProcessEvents(void* ptr, int flags){
	return static_cast<QAbstractEventDispatcher*>(ptr)->processEvents(static_cast<QEventLoop::ProcessEventsFlag>(flags));
}

void QAbstractEventDispatcher_RegisterSocketNotifier(void* ptr, void* notifier){
	static_cast<QAbstractEventDispatcher*>(ptr)->registerSocketNotifier(static_cast<QSocketNotifier*>(notifier));
}

int QAbstractEventDispatcher_RemainingTime(void* ptr, int timerId){
	return static_cast<QAbstractEventDispatcher*>(ptr)->remainingTime(timerId);
}

void QAbstractEventDispatcher_RemoveNativeEventFilter(void* ptr, void* filter){
	static_cast<QAbstractEventDispatcher*>(ptr)->removeNativeEventFilter(static_cast<QAbstractNativeEventFilter*>(filter));
}

void QAbstractEventDispatcher_UnregisterSocketNotifier(void* ptr, void* notifier){
	static_cast<QAbstractEventDispatcher*>(ptr)->unregisterSocketNotifier(static_cast<QSocketNotifier*>(notifier));
}

int QAbstractEventDispatcher_UnregisterTimer(void* ptr, int timerId){
	return static_cast<QAbstractEventDispatcher*>(ptr)->unregisterTimer(timerId);
}

int QAbstractEventDispatcher_UnregisterTimers(void* ptr, void* object){
	return static_cast<QAbstractEventDispatcher*>(ptr)->unregisterTimers(static_cast<QObject*>(object));
}

void QAbstractEventDispatcher_WakeUp(void* ptr){
	static_cast<QAbstractEventDispatcher*>(ptr)->wakeUp();
}

void QAbstractEventDispatcher_DestroyQAbstractEventDispatcher(void* ptr){
	static_cast<QAbstractEventDispatcher*>(ptr)->~QAbstractEventDispatcher();
}

