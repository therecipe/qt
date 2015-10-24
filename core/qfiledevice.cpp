#include "qfiledevice.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QFile>
#include <QFileDevice>
#include "_cgo_export.h"

class MyQFileDevice: public QFileDevice {
public:
};

int QFileDevice_AtEnd(QtObjectPtr ptr){
	return static_cast<QFileDevice*>(ptr)->atEnd();
}

void QFileDevice_Close(QtObjectPtr ptr){
	static_cast<QFileDevice*>(ptr)->close();
}

int QFileDevice_Error(QtObjectPtr ptr){
	return static_cast<QFileDevice*>(ptr)->error();
}

char* QFileDevice_FileName(QtObjectPtr ptr){
	return static_cast<QFileDevice*>(ptr)->fileName().toUtf8().data();
}

int QFileDevice_Flush(QtObjectPtr ptr){
	return static_cast<QFileDevice*>(ptr)->flush();
}

int QFileDevice_Handle(QtObjectPtr ptr){
	return static_cast<QFileDevice*>(ptr)->handle();
}

int QFileDevice_IsSequential(QtObjectPtr ptr){
	return static_cast<QFileDevice*>(ptr)->isSequential();
}

int QFileDevice_Permissions(QtObjectPtr ptr){
	return static_cast<QFileDevice*>(ptr)->permissions();
}

int QFileDevice_SetPermissions(QtObjectPtr ptr, int permissions){
	return static_cast<QFileDevice*>(ptr)->setPermissions(static_cast<QFileDevice::Permission>(permissions));
}

void QFileDevice_UnsetError(QtObjectPtr ptr){
	static_cast<QFileDevice*>(ptr)->unsetError();
}

void QFileDevice_DestroyQFileDevice(QtObjectPtr ptr){
	static_cast<QFileDevice*>(ptr)->~QFileDevice();
}

