#include "qdbusmessage.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QDBusMessage>
#include "_cgo_export.h"

class MyQDBusMessage: public QDBusMessage {
public:
};

void* QDBusMessage_NewQDBusMessage(){
	return new QDBusMessage();
}

void* QDBusMessage_NewQDBusMessage2(void* other){
	return new QDBusMessage(*static_cast<QDBusMessage*>(other));
}

int QDBusMessage_AutoStartService(void* ptr){
	return static_cast<QDBusMessage*>(ptr)->autoStartService();
}

char* QDBusMessage_ErrorMessage(void* ptr){
	return static_cast<QDBusMessage*>(ptr)->errorMessage().toUtf8().data();
}

char* QDBusMessage_ErrorName(void* ptr){
	return static_cast<QDBusMessage*>(ptr)->errorName().toUtf8().data();
}

char* QDBusMessage_Interface(void* ptr){
	return static_cast<QDBusMessage*>(ptr)->interface().toUtf8().data();
}

int QDBusMessage_IsDelayedReply(void* ptr){
	return static_cast<QDBusMessage*>(ptr)->isDelayedReply();
}

int QDBusMessage_IsReplyRequired(void* ptr){
	return static_cast<QDBusMessage*>(ptr)->isReplyRequired();
}

char* QDBusMessage_Member(void* ptr){
	return static_cast<QDBusMessage*>(ptr)->member().toUtf8().data();
}

char* QDBusMessage_Path(void* ptr){
	return static_cast<QDBusMessage*>(ptr)->path().toUtf8().data();
}

char* QDBusMessage_Service(void* ptr){
	return static_cast<QDBusMessage*>(ptr)->service().toUtf8().data();
}

void QDBusMessage_SetAutoStartService(void* ptr, int enable){
	static_cast<QDBusMessage*>(ptr)->setAutoStartService(enable != 0);
}

void QDBusMessage_SetDelayedReply(void* ptr, int enable){
	static_cast<QDBusMessage*>(ptr)->setDelayedReply(enable != 0);
}

char* QDBusMessage_Signature(void* ptr){
	return static_cast<QDBusMessage*>(ptr)->signature().toUtf8().data();
}

int QDBusMessage_Type(void* ptr){
	return static_cast<QDBusMessage*>(ptr)->type();
}

void QDBusMessage_DestroyQDBusMessage(void* ptr){
	static_cast<QDBusMessage*>(ptr)->~QDBusMessage();
}

