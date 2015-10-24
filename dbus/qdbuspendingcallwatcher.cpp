#include "qdbuspendingcallwatcher.h"
#include <QModelIndex>
#include <QObject>
#include <QDBusPendingCall>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QDBusPendingCallWatcher>
#include "_cgo_export.h"

class MyQDBusPendingCallWatcher: public QDBusPendingCallWatcher {
public:
void Signal_Finished(QDBusPendingCallWatcher * self){callbackQDBusPendingCallWatcherFinished(this->objectName().toUtf8().data(), self);};
};

void QDBusPendingCallWatcher_WaitForFinished(QtObjectPtr ptr){
	static_cast<QDBusPendingCallWatcher*>(ptr)->waitForFinished();
}

QtObjectPtr QDBusPendingCallWatcher_NewQDBusPendingCallWatcher(QtObjectPtr call, QtObjectPtr parent){
	return new QDBusPendingCallWatcher(*static_cast<QDBusPendingCall*>(call), static_cast<QObject*>(parent));
}

void QDBusPendingCallWatcher_ConnectFinished(QtObjectPtr ptr){
	QObject::connect(static_cast<QDBusPendingCallWatcher*>(ptr), static_cast<void (QDBusPendingCallWatcher::*)(QDBusPendingCallWatcher *)>(&QDBusPendingCallWatcher::finished), static_cast<MyQDBusPendingCallWatcher*>(ptr), static_cast<void (MyQDBusPendingCallWatcher::*)(QDBusPendingCallWatcher *)>(&MyQDBusPendingCallWatcher::Signal_Finished));;
}

void QDBusPendingCallWatcher_DisconnectFinished(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QDBusPendingCallWatcher*>(ptr), static_cast<void (QDBusPendingCallWatcher::*)(QDBusPendingCallWatcher *)>(&QDBusPendingCallWatcher::finished), static_cast<MyQDBusPendingCallWatcher*>(ptr), static_cast<void (MyQDBusPendingCallWatcher::*)(QDBusPendingCallWatcher *)>(&MyQDBusPendingCallWatcher::Signal_Finished));;
}

int QDBusPendingCallWatcher_IsFinished(QtObjectPtr ptr){
	return static_cast<QDBusPendingCallWatcher*>(ptr)->isFinished();
}

void QDBusPendingCallWatcher_DestroyQDBusPendingCallWatcher(QtObjectPtr ptr){
	static_cast<QDBusPendingCallWatcher*>(ptr)->~QDBusPendingCallWatcher();
}

