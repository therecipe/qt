#include "qmessageauthenticationcode.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QCryptographicHash>
#include <QIODevice>
#include <QByteArray>
#include <QMessageAuthenticationCode>
#include "_cgo_export.h"

class MyQMessageAuthenticationCode: public QMessageAuthenticationCode {
public:
};

void* QMessageAuthenticationCode_NewQMessageAuthenticationCode(int method, void* key){
	return new QMessageAuthenticationCode(static_cast<QCryptographicHash::Algorithm>(method), *static_cast<QByteArray*>(key));
}

int QMessageAuthenticationCode_AddData2(void* ptr, void* device){
	return static_cast<QMessageAuthenticationCode*>(ptr)->addData(static_cast<QIODevice*>(device));
}

void QMessageAuthenticationCode_AddData3(void* ptr, void* data){
	static_cast<QMessageAuthenticationCode*>(ptr)->addData(*static_cast<QByteArray*>(data));
}

void QMessageAuthenticationCode_AddData(void* ptr, char* data, int length){
	static_cast<QMessageAuthenticationCode*>(ptr)->addData(const_cast<const char*>(data), length);
}

void* QMessageAuthenticationCode_QMessageAuthenticationCode_Hash(void* message, void* key, int method){
	return new QByteArray(QMessageAuthenticationCode::hash(*static_cast<QByteArray*>(message), *static_cast<QByteArray*>(key), static_cast<QCryptographicHash::Algorithm>(method)));
}

void QMessageAuthenticationCode_Reset(void* ptr){
	static_cast<QMessageAuthenticationCode*>(ptr)->reset();
}

void* QMessageAuthenticationCode_Result(void* ptr){
	return new QByteArray(static_cast<QMessageAuthenticationCode*>(ptr)->result());
}

void QMessageAuthenticationCode_SetKey(void* ptr, void* key){
	static_cast<QMessageAuthenticationCode*>(ptr)->setKey(*static_cast<QByteArray*>(key));
}

void QMessageAuthenticationCode_DestroyQMessageAuthenticationCode(void* ptr){
	static_cast<QMessageAuthenticationCode*>(ptr)->~QMessageAuthenticationCode();
}

