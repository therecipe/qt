#include "qmessageauthenticationcode.h"
#include <QIODevice>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QByteArray>
#include <QCryptographicHash>
#include <QMessageAuthenticationCode>
#include "_cgo_export.h"

class MyQMessageAuthenticationCode: public QMessageAuthenticationCode {
public:
};

QtObjectPtr QMessageAuthenticationCode_NewQMessageAuthenticationCode(int method, QtObjectPtr key){
	return new QMessageAuthenticationCode(static_cast<QCryptographicHash::Algorithm>(method), *static_cast<QByteArray*>(key));
}

int QMessageAuthenticationCode_AddData2(QtObjectPtr ptr, QtObjectPtr device){
	return static_cast<QMessageAuthenticationCode*>(ptr)->addData(static_cast<QIODevice*>(device));
}

void QMessageAuthenticationCode_AddData3(QtObjectPtr ptr, QtObjectPtr data){
	static_cast<QMessageAuthenticationCode*>(ptr)->addData(*static_cast<QByteArray*>(data));
}

void QMessageAuthenticationCode_AddData(QtObjectPtr ptr, char* data, int length){
	static_cast<QMessageAuthenticationCode*>(ptr)->addData(const_cast<const char*>(data), length);
}

void QMessageAuthenticationCode_Reset(QtObjectPtr ptr){
	static_cast<QMessageAuthenticationCode*>(ptr)->reset();
}

void QMessageAuthenticationCode_SetKey(QtObjectPtr ptr, QtObjectPtr key){
	static_cast<QMessageAuthenticationCode*>(ptr)->setKey(*static_cast<QByteArray*>(key));
}

void QMessageAuthenticationCode_DestroyQMessageAuthenticationCode(QtObjectPtr ptr){
	static_cast<QMessageAuthenticationCode*>(ptr)->~QMessageAuthenticationCode();
}

