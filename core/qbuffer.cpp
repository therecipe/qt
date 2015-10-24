#include "qbuffer.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QByteArray>
#include <QIODevice>
#include <QString>
#include <QBuffer>
#include "_cgo_export.h"

class MyQBuffer: public QBuffer {
public:
};

QtObjectPtr QBuffer_NewQBuffer2(QtObjectPtr byteArray, QtObjectPtr parent){
	return new QBuffer(static_cast<QByteArray*>(byteArray), static_cast<QObject*>(parent));
}

QtObjectPtr QBuffer_NewQBuffer(QtObjectPtr parent){
	return new QBuffer(static_cast<QObject*>(parent));
}

int QBuffer_AtEnd(QtObjectPtr ptr){
	return static_cast<QBuffer*>(ptr)->atEnd();
}

int QBuffer_CanReadLine(QtObjectPtr ptr){
	return static_cast<QBuffer*>(ptr)->canReadLine();
}

void QBuffer_Close(QtObjectPtr ptr){
	static_cast<QBuffer*>(ptr)->close();
}

int QBuffer_Open(QtObjectPtr ptr, int flags){
	return static_cast<QBuffer*>(ptr)->open(static_cast<QIODevice::OpenModeFlag>(flags));
}

void QBuffer_SetBuffer(QtObjectPtr ptr, QtObjectPtr byteArray){
	static_cast<QBuffer*>(ptr)->setBuffer(static_cast<QByteArray*>(byteArray));
}

void QBuffer_SetData(QtObjectPtr ptr, QtObjectPtr data){
	static_cast<QBuffer*>(ptr)->setData(*static_cast<QByteArray*>(data));
}

void QBuffer_SetData2(QtObjectPtr ptr, char* data, int size){
	static_cast<QBuffer*>(ptr)->setData(const_cast<const char*>(data), size);
}

void QBuffer_DestroyQBuffer(QtObjectPtr ptr){
	static_cast<QBuffer*>(ptr)->~QBuffer();
}

