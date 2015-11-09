#include "qdbuspendingcallwatcher.h"
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QDBusPendingCall>
#include <QString>
#include <QVariant>
#include <QDBusPendingCallWatcher>
#include "_cgo_export.h"

class MyQDBusPendingCallWatcher: public QDBusPendingCallWatcher {
public:
void Signal_Finished(QDBusPendingCallWatcher * self){callbackQDBusPendingCallWatcherFinished(this->objectName().toUtf8().data(), self);};
};

void QDBusPendingCallWatcher_WaitForFinished(void* ptr){
	static_cast<QDBusPendingCallWatcher*>(ptr)->waitForFinished();
}

void* QDBusPendingCallWatcher_NewQDBusPendingCallWatcher(void* call, void* parent){
	return new QDBusPendingCallWatcher(*static_cast<QDBusPendingCall*>(call), static_cast<QObject*>(parent));
}

void QDBusPendingCallWatcher_ConnectFinished(void* ptr){
	QObject::connect(static_cast<QDBusPendingCallWatcher*>(ptr), static_cast<void (QDBusPendingCallWatcher::*)(QDBusPendingCallWatcher *)>(&QDBusPendingCallWatcher::finished), static_cast<MyQDBusPendingCallWatcher*>(ptr), static_cast<void (MyQDBusPendingCallWatcher::*)(QDBusPendingCallWatcher *)>(&MyQDBusPendingCallWatcher::Signal_Finished));;
}

void QDBusPendingCallWatcher_DisconnectFinished(void* ptr){
	QObject::disconnect(static_cast<QDBusPendingCallWatcher*>(ptr), static_cast<void (QDBusPendingCallWatcher::*)(QDBusPendingCallWatcher *)>(&QDBusPendingCallWatcher::finished), static_cast<MyQDBusPendingCallWatcher*>(ptr), static_cast<void (MyQDBusPendingCallWatcher::*)(QDBusPendingCallWatcher *)>(&MyQDBusPendingCallWatcher::Signal_Finished));;
}

int QDBusPendingCallWatcher_IsFinished(void* ptr){
	return static_cast<QDBusPendingCallWatcher*>(ptr)->isFinished();
}

void QDBusPendingCallWatcher_DestroyQDBusPendingCallWatcher(void* ptr){
	static_cast<QDBusPendingCallWatcher*>(ptr)->~QDBusPendingCallWatcher();
}

