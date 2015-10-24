#include "qdatastream.h"
#include <QUrl>
#include <QModelIndex>
#include <QByteArray>
#include <QIODevice>
#include <QString>
#include <QVariant>
#include <QDataStream>
#include "_cgo_export.h"

class MyQDataStream: public QDataStream {
public:
};

QtObjectPtr QDataStream_NewQDataStream3(QtObjectPtr a, int mode){
	return new QDataStream(static_cast<QByteArray*>(a), static_cast<QIODevice::OpenModeFlag>(mode));
}

int QDataStream_AtEnd(QtObjectPtr ptr){
	return static_cast<QDataStream*>(ptr)->atEnd();
}

QtObjectPtr QDataStream_NewQDataStream(){
	return new QDataStream();
}

QtObjectPtr QDataStream_NewQDataStream2(QtObjectPtr d){
	return new QDataStream(static_cast<QIODevice*>(d));
}

QtObjectPtr QDataStream_NewQDataStream4(QtObjectPtr a){
	return new QDataStream(*static_cast<QByteArray*>(a));
}

int QDataStream_ByteOrder(QtObjectPtr ptr){
	return static_cast<QDataStream*>(ptr)->byteOrder();
}

QtObjectPtr QDataStream_Device(QtObjectPtr ptr){
	return static_cast<QDataStream*>(ptr)->device();
}

int QDataStream_FloatingPointPrecision(QtObjectPtr ptr){
	return static_cast<QDataStream*>(ptr)->floatingPointPrecision();
}

int QDataStream_ReadRawData(QtObjectPtr ptr, char* s, int len){
	return static_cast<QDataStream*>(ptr)->readRawData(s, len);
}

void QDataStream_ResetStatus(QtObjectPtr ptr){
	static_cast<QDataStream*>(ptr)->resetStatus();
}

void QDataStream_SetByteOrder(QtObjectPtr ptr, int bo){
	static_cast<QDataStream*>(ptr)->setByteOrder(static_cast<QDataStream::ByteOrder>(bo));
}

void QDataStream_SetDevice(QtObjectPtr ptr, QtObjectPtr d){
	static_cast<QDataStream*>(ptr)->setDevice(static_cast<QIODevice*>(d));
}

void QDataStream_SetFloatingPointPrecision(QtObjectPtr ptr, int precision){
	static_cast<QDataStream*>(ptr)->setFloatingPointPrecision(static_cast<QDataStream::FloatingPointPrecision>(precision));
}

void QDataStream_SetStatus(QtObjectPtr ptr, int status){
	static_cast<QDataStream*>(ptr)->setStatus(static_cast<QDataStream::Status>(status));
}

void QDataStream_SetVersion(QtObjectPtr ptr, int v){
	static_cast<QDataStream*>(ptr)->setVersion(v);
}

int QDataStream_SkipRawData(QtObjectPtr ptr, int len){
	return static_cast<QDataStream*>(ptr)->skipRawData(len);
}

int QDataStream_Status(QtObjectPtr ptr){
	return static_cast<QDataStream*>(ptr)->status();
}

int QDataStream_Version(QtObjectPtr ptr){
	return static_cast<QDataStream*>(ptr)->version();
}

int QDataStream_WriteRawData(QtObjectPtr ptr, char* s, int len){
	return static_cast<QDataStream*>(ptr)->writeRawData(const_cast<const char*>(s), len);
}

void QDataStream_DestroyQDataStream(QtObjectPtr ptr){
	static_cast<QDataStream*>(ptr)->~QDataStream();
}

