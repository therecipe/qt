#include "qhttppart.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QNetworkRequest>
#include <QByteArray>
#include <QIODevice>
#include <QString>
#include <QHttpPart>
#include "_cgo_export.h"

class MyQHttpPart: public QHttpPart {
public:
};

QtObjectPtr QHttpPart_NewQHttpPart(){
	return new QHttpPart();
}

QtObjectPtr QHttpPart_NewQHttpPart2(QtObjectPtr other){
	return new QHttpPart(*static_cast<QHttpPart*>(other));
}

void QHttpPart_SetBody(QtObjectPtr ptr, QtObjectPtr body){
	static_cast<QHttpPart*>(ptr)->setBody(*static_cast<QByteArray*>(body));
}

void QHttpPart_SetBodyDevice(QtObjectPtr ptr, QtObjectPtr device){
	static_cast<QHttpPart*>(ptr)->setBodyDevice(static_cast<QIODevice*>(device));
}

void QHttpPart_SetHeader(QtObjectPtr ptr, int header, char* value){
	static_cast<QHttpPart*>(ptr)->setHeader(static_cast<QNetworkRequest::KnownHeaders>(header), QVariant(value));
}

void QHttpPart_SetRawHeader(QtObjectPtr ptr, QtObjectPtr headerName, QtObjectPtr headerValue){
	static_cast<QHttpPart*>(ptr)->setRawHeader(*static_cast<QByteArray*>(headerName), *static_cast<QByteArray*>(headerValue));
}

void QHttpPart_Swap(QtObjectPtr ptr, QtObjectPtr other){
	static_cast<QHttpPart*>(ptr)->swap(*static_cast<QHttpPart*>(other));
}

void QHttpPart_DestroyQHttpPart(QtObjectPtr ptr){
	static_cast<QHttpPart*>(ptr)->~QHttpPart();
}

