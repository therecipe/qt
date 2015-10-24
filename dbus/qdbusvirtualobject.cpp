#include "qdbusvirtualobject.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QDBusConnection>
#include <QDBusMessage>
#include <QDBusVirtualObject>
#include "_cgo_export.h"

class MyQDBusVirtualObject: public QDBusVirtualObject {
public:
};

int QDBusVirtualObject_HandleMessage(QtObjectPtr ptr, QtObjectPtr message, QtObjectPtr connection){
	return static_cast<QDBusVirtualObject*>(ptr)->handleMessage(*static_cast<QDBusMessage*>(message), *static_cast<QDBusConnection*>(connection));
}

char* QDBusVirtualObject_Introspect(QtObjectPtr ptr, char* path){
	return static_cast<QDBusVirtualObject*>(ptr)->introspect(QString(path)).toUtf8().data();
}

void QDBusVirtualObject_DestroyQDBusVirtualObject(QtObjectPtr ptr){
	static_cast<QDBusVirtualObject*>(ptr)->~QDBusVirtualObject();
}

