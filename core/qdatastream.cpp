#include "qdatastream.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QIODevice>
#include <QByteArray>
#include <QDataStream>
#include "_cgo_export.h"

class MyQDataStream: public QDataStream {
public:
};

void* QDataStream_NewQDataStream3(void* a, int mode){
	return new QDataStream(static_cast<QByteArray*>(a), static_cast<QIODevice::OpenModeFlag>(mode));
}

int QDataStream_AtEnd(void* ptr){
	return static_cast<QDataStream*>(ptr)->atEnd();
}

void* QDataStream_NewQDataStream(){
	return new QDataStream();
}

void* QDataStream_NewQDataStream2(void* d){
	return new QDataStream(static_cast<QIODevice*>(d));
}

void* QDataStream_NewQDataStream4(void* a){
	return new QDataStream(*static_cast<QByteArray*>(a));
}

int QDataStream_ByteOrder(void* ptr){
	return static_cast<QDataStream*>(ptr)->byteOrder();
}

void* QDataStream_Device(void* ptr){
	return static_cast<QDataStream*>(ptr)->device();
}

int QDataStream_FloatingPointPrecision(void* ptr){
	return static_cast<QDataStream*>(ptr)->floatingPointPrecision();
}

int QDataStream_ReadRawData(void* ptr, char* s, int len){
	return static_cast<QDataStream*>(ptr)->readRawData(s, len);
}

void QDataStream_ResetStatus(void* ptr){
	static_cast<QDataStream*>(ptr)->resetStatus();
}

void QDataStream_SetByteOrder(void* ptr, int bo){
	static_cast<QDataStream*>(ptr)->setByteOrder(static_cast<QDataStream::ByteOrder>(bo));
}

void QDataStream_SetDevice(void* ptr, void* d){
	static_cast<QDataStream*>(ptr)->setDevice(static_cast<QIODevice*>(d));
}

void QDataStream_SetFloatingPointPrecision(void* ptr, int precision){
	static_cast<QDataStream*>(ptr)->setFloatingPointPrecision(static_cast<QDataStream::FloatingPointPrecision>(precision));
}

void QDataStream_SetStatus(void* ptr, int status){
	static_cast<QDataStream*>(ptr)->setStatus(static_cast<QDataStream::Status>(status));
}

void QDataStream_SetVersion(void* ptr, int v){
	static_cast<QDataStream*>(ptr)->setVersion(v);
}

int QDataStream_SkipRawData(void* ptr, int len){
	return static_cast<QDataStream*>(ptr)->skipRawData(len);
}

int QDataStream_Status(void* ptr){
	return static_cast<QDataStream*>(ptr)->status();
}

int QDataStream_Version(void* ptr){
	return static_cast<QDataStream*>(ptr)->version();
}

int QDataStream_WriteRawData(void* ptr, char* s, int len){
	return static_cast<QDataStream*>(ptr)->writeRawData(const_cast<const char*>(s), len);
}

void QDataStream_DestroyQDataStream(void* ptr){
	static_cast<QDataStream*>(ptr)->~QDataStream();
}

