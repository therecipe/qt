#include "qresource.h"
#include <QModelIndex>
#include <QLocale>
#include <QString>
#include <QVariant>
#include <QUrl>
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

void* QResource_NewQResource(char* file, void* locale){
	return new QResource(QString(file), *static_cast<QLocale*>(locale));
}

char* QResource_AbsoluteFilePath(void* ptr){
	return static_cast<QResource*>(ptr)->absoluteFilePath().toUtf8().data();
}

char* QResource_FileName(void* ptr){
	return static_cast<QResource*>(ptr)->fileName().toUtf8().data();
}

int QResource_IsCompressed(void* ptr){
	return static_cast<QResource*>(ptr)->isCompressed();
}

int QResource_IsValid(void* ptr){
	return static_cast<QResource*>(ptr)->isValid();
}

void QResource_SetFileName(void* ptr, char* file){
	static_cast<QResource*>(ptr)->setFileName(QString(file));
}

void QResource_SetLocale(void* ptr, void* locale){
	static_cast<QResource*>(ptr)->setLocale(*static_cast<QLocale*>(locale));
}

void QResource_DestroyQResource(void* ptr){
	static_cast<QResource*>(ptr)->~QResource();
}

