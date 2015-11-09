#include "qdbuserror.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QDBusError>
#include "_cgo_export.h"

class MyQDBusError: public QDBusError {
public:
};

char* QDBusError_QDBusError_ErrorString(int error){
	return QDBusError::errorString(static_cast<QDBusError::ErrorType>(error)).toUtf8().data();
}

int QDBusError_IsValid(void* ptr){
	return static_cast<QDBusError*>(ptr)->isValid();
}

char* QDBusError_Message(void* ptr){
	return static_cast<QDBusError*>(ptr)->message().toUtf8().data();
}

char* QDBusError_Name(void* ptr){
	return static_cast<QDBusError*>(ptr)->name().toUtf8().data();
}

int QDBusError_Type(void* ptr){
	return static_cast<QDBusError*>(ptr)->type();
}

