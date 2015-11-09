#include "qhttppart.h"
#include <QModelIndex>
#include <QNetworkRequest>
#include <QByteArray>
#include <QIODevice>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QHttpPart>
#include "_cgo_export.h"

class MyQHttpPart: public QHttpPart {
public:
};

void* QHttpPart_NewQHttpPart(){
	return new QHttpPart();
}

void* QHttpPart_NewQHttpPart2(void* other){
	return new QHttpPart(*static_cast<QHttpPart*>(other));
}

void QHttpPart_SetBody(void* ptr, void* body){
	static_cast<QHttpPart*>(ptr)->setBody(*static_cast<QByteArray*>(body));
}

void QHttpPart_SetBodyDevice(void* ptr, void* device){
	static_cast<QHttpPart*>(ptr)->setBodyDevice(static_cast<QIODevice*>(device));
}

void QHttpPart_SetHeader(void* ptr, int header, void* value){
	static_cast<QHttpPart*>(ptr)->setHeader(static_cast<QNetworkRequest::KnownHeaders>(header), *static_cast<QVariant*>(value));
}

void QHttpPart_SetRawHeader(void* ptr, void* headerName, void* headerValue){
	static_cast<QHttpPart*>(ptr)->setRawHeader(*static_cast<QByteArray*>(headerName), *static_cast<QByteArray*>(headerValue));
}

void QHttpPart_Swap(void* ptr, void* other){
	static_cast<QHttpPart*>(ptr)->swap(*static_cast<QHttpPart*>(other));
}

void QHttpPart_DestroyQHttpPart(void* ptr){
	static_cast<QHttpPart*>(ptr)->~QHttpPart();
}

