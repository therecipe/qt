#include "qsocketnotifier.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QMetaObject>
#include <QSocketNotifier>
#include "_cgo_export.h"

class MyQSocketNotifier: public QSocketNotifier {
public:
void Signal_Activated(int socket){callbackQSocketNotifierActivated(this->objectName().toUtf8().data(), socket);};
};

void QSocketNotifier_ConnectActivated(void* ptr){
	QObject::connect(static_cast<QSocketNotifier*>(ptr), &QSocketNotifier::activated, static_cast<MyQSocketNotifier*>(ptr), static_cast<void (MyQSocketNotifier::*)(int)>(&MyQSocketNotifier::Signal_Activated));;
}

void QSocketNotifier_DisconnectActivated(void* ptr){
	QObject::disconnect(static_cast<QSocketNotifier*>(ptr), &QSocketNotifier::activated, static_cast<MyQSocketNotifier*>(ptr), static_cast<void (MyQSocketNotifier::*)(int)>(&MyQSocketNotifier::Signal_Activated));;
}

int QSocketNotifier_IsEnabled(void* ptr){
	return static_cast<QSocketNotifier*>(ptr)->isEnabled();
}

void QSocketNotifier_SetEnabled(void* ptr, int enable){
	QMetaObject::invokeMethod(static_cast<QSocketNotifier*>(ptr), "setEnabled", Q_ARG(bool, enable != 0));
}

int QSocketNotifier_Type(void* ptr){
	return static_cast<QSocketNotifier*>(ptr)->type();
}

void QSocketNotifier_DestroyQSocketNotifier(void* ptr){
	static_cast<QSocketNotifier*>(ptr)->~QSocketNotifier();
}

