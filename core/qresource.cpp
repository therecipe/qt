#include "qresource.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QLocale>
#include <QResource>
#include "_cgo_export.h"

class MyQResource: public QResource {
public:
};

int QResource_QResource_RegisterResource(char* rccFileName, char* mapRoot){
	return QResource::registerResource(QString(rccFileName), QString(mapRoot));
}

int QResource_QResource_UnregisterResource(char* rccFileName, char* mapRoot){
	return QResource::unregisterResource(QString(rccFileName), QString(mapRoot));
}

QtObjectPtr QResource_NewQResource(char* file, QtObjectPtr locale){
	return new QResource(QString(file), *static_cast<QLocale*>(locale));
}

char* QResource_AbsoluteFilePath(QtObjectPtr ptr){
	return static_cast<QResource*>(ptr)->absoluteFilePath().toUtf8().data();
}

char* QResource_FileName(QtObjectPtr ptr){
	return static_cast<QResource*>(ptr)->fileName().toUtf8().data();
}

int QResource_IsCompressed(QtObjectPtr ptr){
	return static_cast<QResource*>(ptr)->isCompressed();
}

int QResource_IsValid(QtObjectPtr ptr){
	return static_cast<QResource*>(ptr)->isValid();
}

void QResource_SetFileName(QtObjectPtr ptr, char* file){
	static_cast<QResource*>(ptr)->setFileName(QString(file));
}

void QResource_SetLocale(QtObjectPtr ptr, QtObjectPtr locale){
	static_cast<QResource*>(ptr)->setLocale(*static_cast<QLocale*>(locale));
}

void QResource_DestroyQResource(QtObjectPtr ptr){
	static_cast<QResource*>(ptr)->~QResource();
}

