#include "qsslkey.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QByteArray>
#include <QSslKey>
#include "_cgo_export.h"

class MyQSslKey: public QSslKey {
public:
};

void* QSslKey_NewQSslKey(){
	return new QSslKey();
}

void* QSslKey_NewQSslKey5(void* other){
	return new QSslKey(*static_cast<QSslKey*>(other));
}

void QSslKey_Clear(void* ptr){
	static_cast<QSslKey*>(ptr)->clear();
}

int QSslKey_IsNull(void* ptr){
	return static_cast<QSslKey*>(ptr)->isNull();
}

int QSslKey_Length(void* ptr){
	return static_cast<QSslKey*>(ptr)->length();
}

void QSslKey_Swap(void* ptr, void* other){
	static_cast<QSslKey*>(ptr)->swap(*static_cast<QSslKey*>(other));
}

void* QSslKey_ToDer(void* ptr, void* passPhrase){
	return new QByteArray(static_cast<QSslKey*>(ptr)->toDer(*static_cast<QByteArray*>(passPhrase)));
}

void* QSslKey_ToPem(void* ptr, void* passPhrase){
	return new QByteArray(static_cast<QSslKey*>(ptr)->toPem(*static_cast<QByteArray*>(passPhrase)));
}

void QSslKey_DestroyQSslKey(void* ptr){
	static_cast<QSslKey*>(ptr)->~QSslKey();
}

