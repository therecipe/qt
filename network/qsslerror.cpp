#include "qsslerror.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSslCertificate>
#include <QSslError>
#include "_cgo_export.h"

class MyQSslError: public QSslError {
public:
};

void* QSslError_NewQSslError(){
	return new QSslError();
}

void* QSslError_NewQSslError2(int error){
	return new QSslError(static_cast<QSslError::SslError>(error));
}

void* QSslError_NewQSslError3(int error, void* certificate){
	return new QSslError(static_cast<QSslError::SslError>(error), *static_cast<QSslCertificate*>(certificate));
}

void* QSslError_NewQSslError4(void* other){
	return new QSslError(*static_cast<QSslError*>(other));
}

int QSslError_Error(void* ptr){
	return static_cast<QSslError*>(ptr)->error();
}

char* QSslError_ErrorString(void* ptr){
	return static_cast<QSslError*>(ptr)->errorString().toUtf8().data();
}

void QSslError_Swap(void* ptr, void* other){
	static_cast<QSslError*>(ptr)->swap(*static_cast<QSslError*>(other));
}

void QSslError_DestroyQSslError(void* ptr){
	static_cast<QSslError*>(ptr)->~QSslError();
}

