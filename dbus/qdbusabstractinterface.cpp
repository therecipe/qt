#include "qdbusabstractinterface.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QDBusAbstractInterface>
#include "_cgo_export.h"

class MyQDBusAbstractInterface: public QDBusAbstractInterface {
public:
};

char* QDBusAbstractInterface_Interface(QtObjectPtr ptr){
	return static_cast<QDBusAbstractInterface*>(ptr)->interface().toUtf8().data();
}

int QDBusAbstractInterface_IsValid(QtObjectPtr ptr){
	return static_cast<QDBusAbstractInterface*>(ptr)->isValid();
}

char* QDBusAbstractInterface_Path(QtObjectPtr ptr){
	return static_cast<QDBusAbstractInterface*>(ptr)->path().toUtf8().data();
}

char* QDBusAbstractInterface_Service(QtObjectPtr ptr){
	return static_cast<QDBusAbstractInterface*>(ptr)->service().toUtf8().data();
}

void QDBusAbstractInterface_SetTimeout(QtObjectPtr ptr, int timeout){
	static_cast<QDBusAbstractInterface*>(ptr)->setTimeout(timeout);
}

int QDBusAbstractInterface_Timeout(QtObjectPtr ptr){
	return static_cast<QDBusAbstractInterface*>(ptr)->timeout();
}

void QDBusAbstractInterface_DestroyQDBusAbstractInterface(QtObjectPtr ptr){
	static_cast<QDBusAbstractInterface*>(ptr)->~QDBusAbstractInterface();
}

