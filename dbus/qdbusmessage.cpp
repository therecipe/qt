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

QtObjectPtr QDBusMessage_NewQDBusMessage(){
	return new QDBusMessage();
}

QtObjectPtr QDBusMessage_NewQDBusMessage2(QtObjectPtr other){
	return new QDBusMessage(*static_cast<QDBusMessage*>(other));
}

int QDBusMessage_AutoStartService(QtObjectPtr ptr){
	return static_cast<QDBusMessage*>(ptr)->autoStartService();
}

char* QDBusMessage_ErrorMessage(QtObjectPtr ptr){
	return static_cast<QDBusMessage*>(ptr)->errorMessage().toUtf8().data();
}

char* QDBusMessage_ErrorName(QtObjectPtr ptr){
	return static_cast<QDBusMessage*>(ptr)->errorName().toUtf8().data();
}

char* QDBusMessage_Interface(QtObjectPtr ptr){
	return static_cast<QDBusMessage*>(ptr)->interface().toUtf8().data();
}

int QDBusMessage_IsDelayedReply(QtObjectPtr ptr){
	return static_cast<QDBusMessage*>(ptr)->isDelayedReply();
}

int QDBusMessage_IsReplyRequired(QtObjectPtr ptr){
	return static_cast<QDBusMessage*>(ptr)->isReplyRequired();
}

char* QDBusMessage_Member(QtObjectPtr ptr){
	return static_cast<QDBusMessage*>(ptr)->member().toUtf8().data();
}

char* QDBusMessage_Path(QtObjectPtr ptr){
	return static_cast<QDBusMessage*>(ptr)->path().toUtf8().data();
}

char* QDBusMessage_Service(QtObjectPtr ptr){
	return static_cast<QDBusMessage*>(ptr)->service().toUtf8().data();
}

void QDBusMessage_SetAutoStartService(QtObjectPtr ptr, int enable){
	static_cast<QDBusMessage*>(ptr)->setAutoStartService(enable != 0);
}

void QDBusMessage_SetDelayedReply(QtObjectPtr ptr, int enable){
	static_cast<QDBusMessage*>(ptr)->setDelayedReply(enable != 0);
}

char* QDBusMessage_Signature(QtObjectPtr ptr){
	return static_cast<QDBusMessage*>(ptr)->signature().toUtf8().data();
}

int QDBusMessage_Type(QtObjectPtr ptr){
	return static_cast<QDBusMessage*>(ptr)->type();
}

void QDBusMessage_DestroyQDBusMessage(QtObjectPtr ptr){
	static_cast<QDBusMessage*>(ptr)->~QDBusMessage();
}

