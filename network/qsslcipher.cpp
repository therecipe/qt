#include "qsslcipher.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSslCipher>
#include "_cgo_export.h"

class MyQSslCipher: public QSslCipher {
public:
};

void* QSslCipher_NewQSslCipher(){
	return new QSslCipher();
}

void* QSslCipher_NewQSslCipher4(void* other){
	return new QSslCipher(*static_cast<QSslCipher*>(other));
}

void* QSslCipher_NewQSslCipher2(char* name){
	return new QSslCipher(QString(name));
}

char* QSslCipher_AuthenticationMethod(void* ptr){
	return static_cast<QSslCipher*>(ptr)->authenticationMethod().toUtf8().data();
}

char* QSslCipher_EncryptionMethod(void* ptr){
	return static_cast<QSslCipher*>(ptr)->encryptionMethod().toUtf8().data();
}

int QSslCipher_IsNull(void* ptr){
	return static_cast<QSslCipher*>(ptr)->isNull();
}

char* QSslCipher_KeyExchangeMethod(void* ptr){
	return static_cast<QSslCipher*>(ptr)->keyExchangeMethod().toUtf8().data();
}

char* QSslCipher_Name(void* ptr){
	return static_cast<QSslCipher*>(ptr)->name().toUtf8().data();
}

char* QSslCipher_ProtocolString(void* ptr){
	return static_cast<QSslCipher*>(ptr)->protocolString().toUtf8().data();
}

int QSslCipher_SupportedBits(void* ptr){
	return static_cast<QSslCipher*>(ptr)->supportedBits();
}

void QSslCipher_Swap(void* ptr, void* other){
	static_cast<QSslCipher*>(ptr)->swap(*static_cast<QSslCipher*>(other));
}

int QSslCipher_UsedBits(void* ptr){
	return static_cast<QSslCipher*>(ptr)->usedBits();
}

void QSslCipher_DestroyQSslCipher(void* ptr){
	static_cast<QSslCipher*>(ptr)->~QSslCipher();
}

