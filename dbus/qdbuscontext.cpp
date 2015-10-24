#include "qdbuscontext.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QDBusError>
#include <QDBusContext>
#include "_cgo_export.h"

class MyQDBusContext: public QDBusContext {
public:
};

QtObjectPtr QDBusContext_NewQDBusContext(){
	return new QDBusContext();
}

int QDBusContext_CalledFromDBus(QtObjectPtr ptr){
	return static_cast<QDBusContext*>(ptr)->calledFromDBus();
}

int QDBusContext_IsDelayedReply(QtObjectPtr ptr){
	return static_cast<QDBusContext*>(ptr)->isDelayedReply();
}

void QDBusContext_SendErrorReply2(QtObjectPtr ptr, int ty, char* msg){
	static_cast<QDBusContext*>(ptr)->sendErrorReply(static_cast<QDBusError::ErrorType>(ty), QString(msg));
}

void QDBusContext_SendErrorReply(QtObjectPtr ptr, char* name, char* msg){
	static_cast<QDBusContext*>(ptr)->sendErrorReply(QString(name), QString(msg));
}

void QDBusContext_SetDelayedReply(QtObjectPtr ptr, int enable){
	static_cast<QDBusContext*>(ptr)->setDelayedReply(enable != 0);
}

void QDBusContext_DestroyQDBusContext(QtObjectPtr ptr){
	static_cast<QDBusContext*>(ptr)->~QDBusContext();
}

