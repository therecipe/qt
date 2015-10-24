#include "qsocketnotifier.h"
#include <QUrl>
#include <QModelIndex>
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QSocketNotifier>
#include "_cgo_export.h"

class MyQSocketNotifier: public QSocketNotifier {
public:
void Signal_Activated(int socket){callbackQSocketNotifierActivated(this->objectName().toUtf8().data(), socket);};
};

void QSocketNotifier_ConnectActivated(QtObjectPtr ptr){
	QObject::connect(static_cast<QSocketNotifier*>(ptr), &QSocketNotifier::activated, static_cast<MyQSocketNotifier*>(ptr), static_cast<void (MyQSocketNotifier::*)(int)>(&MyQSocketNotifier::Signal_Activated));;
}

void QSocketNotifier_DisconnectActivated(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QSocketNotifier*>(ptr), &QSocketNotifier::activated, static_cast<MyQSocketNotifier*>(ptr), static_cast<void (MyQSocketNotifier::*)(int)>(&MyQSocketNotifier::Signal_Activated));;
}

int QSocketNotifier_IsEnabled(QtObjectPtr ptr){
	return static_cast<QSocketNotifier*>(ptr)->isEnabled();
}

void QSocketNotifier_SetEnabled(QtObjectPtr ptr, int enable){
	QMetaObject::invokeMethod(static_cast<QSocketNotifier*>(ptr), "setEnabled", Q_ARG(bool, enable != 0));
}

int QSocketNotifier_Type(QtObjectPtr ptr){
	return static_cast<QSocketNotifier*>(ptr)->type();
}

void QSocketNotifier_DestroyQSocketNotifier(QtObjectPtr ptr){
	static_cast<QSocketNotifier*>(ptr)->~QSocketNotifier();
}

