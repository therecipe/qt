#include "qbuffer.h"
#include <QUrl>
#include <QModelIndex>
#include <QByteArray>
#include <QObject>
#include <QIODevice>
#include <QString>
#include <QVariant>
#include <QBuffer>
#include "_cgo_export.h"

class MyQBuffer: public QBuffer {
public:
};

void* QBuffer_NewQBuffer2(void* byteArray, void* parent){
	return new QBuffer(static_cast<QByteArray*>(byteArray), static_cast<QObject*>(parent));
}

void* QBuffer_NewQBuffer(void* parent){
	return new QBuffer(static_cast<QObject*>(parent));
}

int QBuffer_AtEnd(void* ptr){
	return static_cast<QBuffer*>(ptr)->atEnd();
}

void* QBuffer_Buffer(void* ptr){
	return new QByteArray(static_cast<QBuffer*>(ptr)->buffer());
}

void* QBuffer_Buffer2(void* ptr){
	return new QByteArray(static_cast<QBuffer*>(ptr)->buffer());
}

int QBuffer_CanReadLine(void* ptr){
	return static_cast<QBuffer*>(ptr)->canReadLine();
}

void QBuffer_Close(void* ptr){
	static_cast<QBuffer*>(ptr)->close();
}

void* QBuffer_Data(void* ptr){
	return new QByteArray(static_cast<QBuffer*>(ptr)->data());
}

int QBuffer_Open(void* ptr, int flags){
	return static_cast<QBuffer*>(ptr)->open(static_cast<QIODevice::OpenModeFlag>(flags));
}

void QBuffer_SetBuffer(void* ptr, void* byteArray){
	static_cast<QBuffer*>(ptr)->setBuffer(static_cast<QByteArray*>(byteArray));
}

void QBuffer_SetData(void* ptr, void* data){
	static_cast<QBuffer*>(ptr)->setData(*static_cast<QByteArray*>(data));
}

void QBuffer_SetData2(void* ptr, char* data, int size){
	static_cast<QBuffer*>(ptr)->setData(const_cast<const char*>(data), size);
}

void QBuffer_DestroyQBuffer(void* ptr){
	static_cast<QBuffer*>(ptr)->~QBuffer();
}

