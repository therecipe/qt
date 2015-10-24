#include "qdbuserror.h"
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QDBusError>
#include "_cgo_export.h"

class MyQDBusError: public QDBusError {
public:
};

char* QDBusError_QDBusError_ErrorString(int error){
	return QDBusError::errorString(static_cast<QDBusError::ErrorType>(error)).toUtf8().data();
}

int QDBusError_IsValid(QtObjectPtr ptr){
	return static_cast<QDBusError*>(ptr)->isValid();
}

char* QDBusError_Message(QtObjectPtr ptr){
	return static_cast<QDBusError*>(ptr)->message().toUtf8().data();
}

char* QDBusError_Name(QtObjectPtr ptr){
	return static_cast<QDBusError*>(ptr)->name().toUtf8().data();
}

int QDBusError_Type(QtObjectPtr ptr){
	return static_cast<QDBusError*>(ptr)->type();
}

