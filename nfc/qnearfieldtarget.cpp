#include "qnearfieldtarget.h"
#include <QModelIndex>
#include <QObject>
#include <QByteArray>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QNearFieldTarget>
#include "_cgo_export.h"

class MyQNearFieldTarget: public QNearFieldTarget {
public:
void Signal_Disconnected(){callbackQNearFieldTargetDisconnected(this->objectName().toUtf8().data());};
void Signal_NdefMessagesWritten(){callbackQNearFieldTargetNdefMessagesWritten(this->objectName().toUtf8().data());};
};

int QNearFieldTarget_AccessMethods(void* ptr){
	return static_cast<QNearFieldTarget*>(ptr)->accessMethods();
}

void QNearFieldTarget_ConnectDisconnected(void* ptr){
	QObject::connect(static_cast<QNearFieldTarget*>(ptr), static_cast<void (QNearFieldTarget::*)()>(&QNearFieldTarget::disconnected), static_cast<MyQNearFieldTarget*>(ptr), static_cast<void (MyQNearFieldTarget::*)()>(&MyQNearFieldTarget::Signal_Disconnected));;
}

void QNearFieldTarget_DisconnectDisconnected(void* ptr){
	QObject::disconnect(static_cast<QNearFieldTarget*>(ptr), static_cast<void (QNearFieldTarget::*)()>(&QNearFieldTarget::disconnected), static_cast<MyQNearFieldTarget*>(ptr), static_cast<void (MyQNearFieldTarget::*)()>(&MyQNearFieldTarget::Signal_Disconnected));;
}

int QNearFieldTarget_HasNdefMessage(void* ptr){
	return static_cast<QNearFieldTarget*>(ptr)->hasNdefMessage();
}

int QNearFieldTarget_IsProcessingCommand(void* ptr){
	return static_cast<QNearFieldTarget*>(ptr)->isProcessingCommand();
}

void QNearFieldTarget_ConnectNdefMessagesWritten(void* ptr){
	QObject::connect(static_cast<QNearFieldTarget*>(ptr), static_cast<void (QNearFieldTarget::*)()>(&QNearFieldTarget::ndefMessagesWritten), static_cast<MyQNearFieldTarget*>(ptr), static_cast<void (MyQNearFieldTarget::*)()>(&MyQNearFieldTarget::Signal_NdefMessagesWritten));;
}

void QNearFieldTarget_DisconnectNdefMessagesWritten(void* ptr){
	QObject::disconnect(static_cast<QNearFieldTarget*>(ptr), static_cast<void (QNearFieldTarget::*)()>(&QNearFieldTarget::ndefMessagesWritten), static_cast<MyQNearFieldTarget*>(ptr), static_cast<void (MyQNearFieldTarget::*)()>(&MyQNearFieldTarget::Signal_NdefMessagesWritten));;
}

int QNearFieldTarget_Type(void* ptr){
	return static_cast<QNearFieldTarget*>(ptr)->type();
}

void* QNearFieldTarget_Uid(void* ptr){
	return new QByteArray(static_cast<QNearFieldTarget*>(ptr)->uid());
}

void QNearFieldTarget_DestroyQNearFieldTarget(void* ptr){
	static_cast<QNearFieldTarget*>(ptr)->~QNearFieldTarget();
}

