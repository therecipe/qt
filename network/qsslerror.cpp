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

QtObjectPtr QSslError_NewQSslError(){
	return new QSslError();
}

QtObjectPtr QSslError_NewQSslError2(int error){
	return new QSslError(static_cast<QSslError::SslError>(error));
}

QtObjectPtr QSslError_NewQSslError3(int error, QtObjectPtr certificate){
	return new QSslError(static_cast<QSslError::SslError>(error), *static_cast<QSslCertificate*>(certificate));
}

QtObjectPtr QSslError_NewQSslError4(QtObjectPtr other){
	return new QSslError(*static_cast<QSslError*>(other));
}

int QSslError_Error(QtObjectPtr ptr){
	return static_cast<QSslError*>(ptr)->error();
}

char* QSslError_ErrorString(QtObjectPtr ptr){
	return static_cast<QSslError*>(ptr)->errorString().toUtf8().data();
}

void QSslError_Swap(QtObjectPtr ptr, QtObjectPtr other){
	static_cast<QSslError*>(ptr)->swap(*static_cast<QSslError*>(other));
}

void QSslError_DestroyQSslError(QtObjectPtr ptr){
	static_cast<QSslError*>(ptr)->~QSslError();
}

