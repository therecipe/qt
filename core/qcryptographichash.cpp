#include "qcryptographichash.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QIODevice>
#include <QByteArray>
#include <QCryptographicHash>
#include "_cgo_export.h"

class MyQCryptographicHash: public QCryptographicHash {
public:
};

QtObjectPtr QCryptographicHash_NewQCryptographicHash(int method){
	return new QCryptographicHash(static_cast<QCryptographicHash::Algorithm>(method));
}

int QCryptographicHash_AddData2(QtObjectPtr ptr, QtObjectPtr device){
	return static_cast<QCryptographicHash*>(ptr)->addData(static_cast<QIODevice*>(device));
}

void QCryptographicHash_AddData3(QtObjectPtr ptr, QtObjectPtr data){
	static_cast<QCryptographicHash*>(ptr)->addData(*static_cast<QByteArray*>(data));
}

void QCryptographicHash_AddData(QtObjectPtr ptr, char* data, int length){
	static_cast<QCryptographicHash*>(ptr)->addData(const_cast<const char*>(data), length);
}

void QCryptographicHash_Reset(QtObjectPtr ptr){
	static_cast<QCryptographicHash*>(ptr)->reset();
}

void QCryptographicHash_DestroyQCryptographicHash(QtObjectPtr ptr){
	static_cast<QCryptographicHash*>(ptr)->~QCryptographicHash();
}

