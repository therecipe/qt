#include "qsignaltransition.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QByteArray>
#include <QState>
#include <QSignalTransition>
#include "_cgo_export.h"

class MyQSignalTransition: public QSignalTransition {
public:
void Signal_SenderObjectChanged(){callbackQSignalTransitionSenderObjectChanged(this->objectName().toUtf8().data());};
void Signal_SignalChanged(){callbackQSignalTransitionSignalChanged(this->objectName().toUtf8().data());};
};

QtObjectPtr QSignalTransition_NewQSignalTransition(QtObjectPtr sourceState){
	return new QSignalTransition(static_cast<QState*>(sourceState));
}

QtObjectPtr QSignalTransition_NewQSignalTransition2(QtObjectPtr sender, char* signal, QtObjectPtr sourceState){
	return new QSignalTransition(static_cast<QObject*>(sender), const_cast<const char*>(signal), static_cast<QState*>(sourceState));
}

QtObjectPtr QSignalTransition_SenderObject(QtObjectPtr ptr){
	return static_cast<QSignalTransition*>(ptr)->senderObject();
}

void QSignalTransition_ConnectSenderObjectChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QSignalTransition*>(ptr), &QSignalTransition::senderObjectChanged, static_cast<MyQSignalTransition*>(ptr), static_cast<void (MyQSignalTransition::*)()>(&MyQSignalTransition::Signal_SenderObjectChanged));;
}

void QSignalTransition_DisconnectSenderObjectChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QSignalTransition*>(ptr), &QSignalTransition::senderObjectChanged, static_cast<MyQSignalTransition*>(ptr), static_cast<void (MyQSignalTransition::*)()>(&MyQSignalTransition::Signal_SenderObjectChanged));;
}

void QSignalTransition_SetSenderObject(QtObjectPtr ptr, QtObjectPtr sender){
	static_cast<QSignalTransition*>(ptr)->setSenderObject(static_cast<QObject*>(sender));
}

void QSignalTransition_SetSignal(QtObjectPtr ptr, QtObjectPtr signal){
	static_cast<QSignalTransition*>(ptr)->setSignal(*static_cast<QByteArray*>(signal));
}

void QSignalTransition_ConnectSignalChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QSignalTransition*>(ptr), &QSignalTransition::signalChanged, static_cast<MyQSignalTransition*>(ptr), static_cast<void (MyQSignalTransition::*)()>(&MyQSignalTransition::Signal_SignalChanged));;
}

void QSignalTransition_DisconnectSignalChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QSignalTransition*>(ptr), &QSignalTransition::signalChanged, static_cast<MyQSignalTransition*>(ptr), static_cast<void (MyQSignalTransition::*)()>(&MyQSignalTransition::Signal_SignalChanged));;
}

void QSignalTransition_DestroyQSignalTransition(QtObjectPtr ptr){
	static_cast<QSignalTransition*>(ptr)->~QSignalTransition();
}

