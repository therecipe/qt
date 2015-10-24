#include "qnearfieldtarget.h"
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QNearFieldTarget>
#include "_cgo_export.h"

class MyQNearFieldTarget: public QNearFieldTarget {
public:
void Signal_Disconnected(){callbackQNearFieldTargetDisconnected(this->objectName().toUtf8().data());};
void Signal_NdefMessagesWritten(){callbackQNearFieldTargetNdefMessagesWritten(this->objectName().toUtf8().data());};
};

int QNearFieldTarget_AccessMethods(QtObjectPtr ptr){
	return static_cast<QNearFieldTarget*>(ptr)->accessMethods();
}

void QNearFieldTarget_ConnectDisconnected(QtObjectPtr ptr){
	QObject::connect(static_cast<QNearFieldTarget*>(ptr), static_cast<void (QNearFieldTarget::*)()>(&QNearFieldTarget::disconnected), static_cast<MyQNearFieldTarget*>(ptr), static_cast<void (MyQNearFieldTarget::*)()>(&MyQNearFieldTarget::Signal_Disconnected));;
}

void QNearFieldTarget_DisconnectDisconnected(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QNearFieldTarget*>(ptr), static_cast<void (QNearFieldTarget::*)()>(&QNearFieldTarget::disconnected), static_cast<MyQNearFieldTarget*>(ptr), static_cast<void (MyQNearFieldTarget::*)()>(&MyQNearFieldTarget::Signal_Disconnected));;
}

int QNearFieldTarget_HasNdefMessage(QtObjectPtr ptr){
	return static_cast<QNearFieldTarget*>(ptr)->hasNdefMessage();
}

int QNearFieldTarget_IsProcessingCommand(QtObjectPtr ptr){
	return static_cast<QNearFieldTarget*>(ptr)->isProcessingCommand();
}

void QNearFieldTarget_ConnectNdefMessagesWritten(QtObjectPtr ptr){
	QObject::connect(static_cast<QNearFieldTarget*>(ptr), static_cast<void (QNearFieldTarget::*)()>(&QNearFieldTarget::ndefMessagesWritten), static_cast<MyQNearFieldTarget*>(ptr), static_cast<void (MyQNearFieldTarget::*)()>(&MyQNearFieldTarget::Signal_NdefMessagesWritten));;
}

void QNearFieldTarget_DisconnectNdefMessagesWritten(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QNearFieldTarget*>(ptr), static_cast<void (QNearFieldTarget::*)()>(&QNearFieldTarget::ndefMessagesWritten), static_cast<MyQNearFieldTarget*>(ptr), static_cast<void (MyQNearFieldTarget::*)()>(&MyQNearFieldTarget::Signal_NdefMessagesWritten));;
}

int QNearFieldTarget_Type(QtObjectPtr ptr){
	return static_cast<QNearFieldTarget*>(ptr)->type();
}

char* QNearFieldTarget_Url(QtObjectPtr ptr){
	return static_cast<QNearFieldTarget*>(ptr)->url().toString().toUtf8().data();
}

void QNearFieldTarget_DestroyQNearFieldTarget(QtObjectPtr ptr){
	static_cast<QNearFieldTarget*>(ptr)->~QNearFieldTarget();
}

