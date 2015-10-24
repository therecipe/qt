#include "qsslcipher.h"
#include <QUrl>
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QSslCipher>
#include "_cgo_export.h"

class MyQSslCipher: public QSslCipher {
public:
};

QtObjectPtr QSslCipher_NewQSslCipher(){
	return new QSslCipher();
}

QtObjectPtr QSslCipher_NewQSslCipher4(QtObjectPtr other){
	return new QSslCipher(*static_cast<QSslCipher*>(other));
}

QtObjectPtr QSslCipher_NewQSslCipher2(char* name){
	return new QSslCipher(QString(name));
}

char* QSslCipher_AuthenticationMethod(QtObjectPtr ptr){
	return static_cast<QSslCipher*>(ptr)->authenticationMethod().toUtf8().data();
}

char* QSslCipher_EncryptionMethod(QtObjectPtr ptr){
	return static_cast<QSslCipher*>(ptr)->encryptionMethod().toUtf8().data();
}

int QSslCipher_IsNull(QtObjectPtr ptr){
	return static_cast<QSslCipher*>(ptr)->isNull();
}

char* QSslCipher_KeyExchangeMethod(QtObjectPtr ptr){
	return static_cast<QSslCipher*>(ptr)->keyExchangeMethod().toUtf8().data();
}

char* QSslCipher_Name(QtObjectPtr ptr){
	return static_cast<QSslCipher*>(ptr)->name().toUtf8().data();
}

char* QSslCipher_ProtocolString(QtObjectPtr ptr){
	return static_cast<QSslCipher*>(ptr)->protocolString().toUtf8().data();
}

int QSslCipher_SupportedBits(QtObjectPtr ptr){
	return static_cast<QSslCipher*>(ptr)->supportedBits();
}

void QSslCipher_Swap(QtObjectPtr ptr, QtObjectPtr other){
	static_cast<QSslCipher*>(ptr)->swap(*static_cast<QSslCipher*>(other));
}

int QSslCipher_UsedBits(QtObjectPtr ptr){
	return static_cast<QSslCipher*>(ptr)->usedBits();
}

void QSslCipher_DestroyQSslCipher(QtObjectPtr ptr){
	static_cast<QSslCipher*>(ptr)->~QSslCipher();
}

