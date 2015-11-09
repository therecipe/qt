#include "qdbusabstractinterface.h"
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QDBusAbstractInterface>
#include "_cgo_export.h"

class MyQDBusAbstractInterface: public QDBusAbstractInterface {
public:
};

char* QDBusAbstractInterface_Interface(void* ptr){
	return static_cast<QDBusAbstractInterface*>(ptr)->interface().toUtf8().data();
}

int QDBusAbstractInterface_IsValid(void* ptr){
	return static_cast<QDBusAbstractInterface*>(ptr)->isValid();
}

char* QDBusAbstractInterface_Path(void* ptr){
	return static_cast<QDBusAbstractInterface*>(ptr)->path().toUtf8().data();
}

char* QDBusAbstractInterface_Service(void* ptr){
	return static_cast<QDBusAbstractInterface*>(ptr)->service().toUtf8().data();
}

void QDBusAbstractInterface_SetTimeout(void* ptr, int timeout){
	static_cast<QDBusAbstractInterface*>(ptr)->setTimeout(timeout);
}

int QDBusAbstractInterface_Timeout(void* ptr){
	return static_cast<QDBusAbstractInterface*>(ptr)->timeout();
}

void QDBusAbstractInterface_DestroyQDBusAbstractInterface(void* ptr){
	static_cast<QDBusAbstractInterface*>(ptr)->~QDBusAbstractInterface();
}

