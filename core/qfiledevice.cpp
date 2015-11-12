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

int QFileDevice_AtEnd(void* ptr){
	return static_cast<QFileDevice*>(ptr)->atEnd();
}

void QFileDevice_Close(void* ptr){
	static_cast<QFileDevice*>(ptr)->close();
}

int QFileDevice_Error(void* ptr){
	return static_cast<QFileDevice*>(ptr)->error();
}

char* QFileDevice_FileName(void* ptr){
	return static_cast<QFileDevice*>(ptr)->fileName().toUtf8().data();
}

int QFileDevice_Flush(void* ptr){
	return static_cast<QFileDevice*>(ptr)->flush();
}

int QFileDevice_Handle(void* ptr){
	return static_cast<QFileDevice*>(ptr)->handle();
}

int QFileDevice_IsSequential(void* ptr){
	return static_cast<QFileDevice*>(ptr)->isSequential();
}

int QFileDevice_Permissions(void* ptr){
	return static_cast<QFileDevice*>(ptr)->permissions();
}

int QFileDevice_SetPermissions(void* ptr, int permissions){
	return static_cast<QFileDevice*>(ptr)->setPermissions(static_cast<QFileDevice::Permission>(permissions));
}

void QFileDevice_UnsetError(void* ptr){
	static_cast<QFileDevice*>(ptr)->unsetError();
}

void QFileDevice_DestroyQFileDevice(void* ptr){
	static_cast<QFileDevice*>(ptr)->~QFileDevice();
}

