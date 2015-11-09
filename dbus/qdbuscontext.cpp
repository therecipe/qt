#include "qdbuscontext.h"
#include <QUrl>
#include <QModelIndex>
#include <QDBusError>
#include <QString>
#include <QVariant>
#include <QDBusContext>
#include "_cgo_export.h"

class MyQDBusContext: public QDBusContext {
public:
};

void* QDBusContext_NewQDBusContext(){
	return new QDBusContext();
}

int QDBusContext_CalledFromDBus(void* ptr){
	return static_cast<QDBusContext*>(ptr)->calledFromDBus();
}

int QDBusContext_IsDelayedReply(void* ptr){
	return static_cast<QDBusContext*>(ptr)->isDelayedReply();
}

void QDBusContext_SendErrorReply2(void* ptr, int ty, char* msg){
	static_cast<QDBusContext*>(ptr)->sendErrorReply(static_cast<QDBusError::ErrorType>(ty), QString(msg));
}

void QDBusContext_SendErrorReply(void* ptr, char* name, char* msg){
	static_cast<QDBusContext*>(ptr)->sendErrorReply(QString(name), QString(msg));
}

void QDBusContext_SetDelayedReply(void* ptr, int enable){
	static_cast<QDBusContext*>(ptr)->setDelayedReply(enable != 0);
}

void QDBusContext_DestroyQDBusContext(void* ptr){
	static_cast<QDBusContext*>(ptr)->~QDBusContext();
}

