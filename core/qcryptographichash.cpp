#include "qcryptographichash.h"
#include <QIODevice>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QByteArray>
#include <QCryptographicHash>
#include "_cgo_export.h"

class MyQCryptographicHash: public QCryptographicHash {
public:
};

void* QCryptographicHash_NewQCryptographicHash(int method){
	return new QCryptographicHash(static_cast<QCryptographicHash::Algorithm>(method));
}

int QCryptographicHash_AddData2(void* ptr, void* device){
	return static_cast<QCryptographicHash*>(ptr)->addData(static_cast<QIODevice*>(device));
}

void QCryptographicHash_AddData3(void* ptr, void* data){
	static_cast<QCryptographicHash*>(ptr)->addData(*static_cast<QByteArray*>(data));
}

void QCryptographicHash_AddData(void* ptr, char* data, int length){
	static_cast<QCryptographicHash*>(ptr)->addData(const_cast<const char*>(data), length);
}

void* QCryptographicHash_QCryptographicHash_Hash(void* data, int method){
	return new QByteArray(QCryptographicHash::hash(*static_cast<QByteArray*>(data), static_cast<QCryptographicHash::Algorithm>(method)));
}

void QCryptographicHash_Reset(void* ptr){
	static_cast<QCryptographicHash*>(ptr)->reset();
}

void* QCryptographicHash_Result(void* ptr){
	return new QByteArray(static_cast<QCryptographicHash*>(ptr)->result());
}

void QCryptographicHash_DestroyQCryptographicHash(void* ptr){
	static_cast<QCryptographicHash*>(ptr)->~QCryptographicHash();
}

