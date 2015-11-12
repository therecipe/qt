#include "qdbusvirtualobject.h"
#include <QModelIndex>
#include <QDBusConnection>
#include <QDBusMessage>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QDBusVirtualObject>
#include "_cgo_export.h"

class MyQDBusVirtualObject: public QDBusVirtualObject {
public:
};

int QDBusVirtualObject_HandleMessage(void* ptr, void* message, void* connection){
	return static_cast<QDBusVirtualObject*>(ptr)->handleMessage(*static_cast<QDBusMessage*>(message), *static_cast<QDBusConnection*>(connection));
}

char* QDBusVirtualObject_Introspect(void* ptr, char* path){
	return static_cast<QDBusVirtualObject*>(ptr)->introspect(QString(path)).toUtf8().data();
}

void QDBusVirtualObject_DestroyQDBusVirtualObject(void* ptr){
	static_cast<QDBusVirtualObject*>(ptr)->~QDBusVirtualObject();
}

