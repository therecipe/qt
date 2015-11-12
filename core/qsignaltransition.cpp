#include "qsignaltransition.h"
#include <QUrl>
#include <QModelIndex>
#include <QState>
#include <QByteArray>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QSignalTransition>
#include "_cgo_export.h"

class MyQSignalTransition: public QSignalTransition {
public:
void Signal_SenderObjectChanged(){callbackQSignalTransitionSenderObjectChanged(this->objectName().toUtf8().data());};
void Signal_SignalChanged(){callbackQSignalTransitionSignalChanged(this->objectName().toUtf8().data());};
};

void* QSignalTransition_NewQSignalTransition(void* sourceState){
	return new QSignalTransition(static_cast<QState*>(sourceState));
}

void* QSignalTransition_NewQSignalTransition2(void* sender, char* signal, void* sourceState){
	return new QSignalTransition(static_cast<QObject*>(sender), const_cast<const char*>(signal), static_cast<QState*>(sourceState));
}

void* QSignalTransition_SenderObject(void* ptr){
	return static_cast<QSignalTransition*>(ptr)->senderObject();
}

void QSignalTransition_ConnectSenderObjectChanged(void* ptr){
	QObject::connect(static_cast<QSignalTransition*>(ptr), &QSignalTransition::senderObjectChanged, static_cast<MyQSignalTransition*>(ptr), static_cast<void (MyQSignalTransition::*)()>(&MyQSignalTransition::Signal_SenderObjectChanged));;
}

void QSignalTransition_DisconnectSenderObjectChanged(void* ptr){
	QObject::disconnect(static_cast<QSignalTransition*>(ptr), &QSignalTransition::senderObjectChanged, static_cast<MyQSignalTransition*>(ptr), static_cast<void (MyQSignalTransition::*)()>(&MyQSignalTransition::Signal_SenderObjectChanged));;
}

void QSignalTransition_SetSenderObject(void* ptr, void* sender){
	static_cast<QSignalTransition*>(ptr)->setSenderObject(static_cast<QObject*>(sender));
}

void QSignalTransition_SetSignal(void* ptr, void* signal){
	static_cast<QSignalTransition*>(ptr)->setSignal(*static_cast<QByteArray*>(signal));
}

void* QSignalTransition_Signal(void* ptr){
	return new QByteArray(static_cast<QSignalTransition*>(ptr)->signal());
}

void QSignalTransition_ConnectSignalChanged(void* ptr){
	QObject::connect(static_cast<QSignalTransition*>(ptr), &QSignalTransition::signalChanged, static_cast<MyQSignalTransition*>(ptr), static_cast<void (MyQSignalTransition::*)()>(&MyQSignalTransition::Signal_SignalChanged));;
}

void QSignalTransition_DisconnectSignalChanged(void* ptr){
	QObject::disconnect(static_cast<QSignalTransition*>(ptr), &QSignalTransition::signalChanged, static_cast<MyQSignalTransition*>(ptr), static_cast<void (MyQSignalTransition::*)()>(&MyQSignalTransition::Signal_SignalChanged));;
}

void QSignalTransition_DestroyQSignalTransition(void* ptr){
	static_cast<QSignalTransition*>(ptr)->~QSignalTransition();
}

