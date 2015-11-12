#include "qhttpmultipart.h"
#include <QModelIndex>
#include <QHttpPart>
#include <QByteArray>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QHttpMultiPart>
#include "_cgo_export.h"

class MyQHttpMultiPart: public QHttpMultiPart {
public:
};

void* QHttpMultiPart_NewQHttpMultiPart2(int contentType, void* parent){
	return new QHttpMultiPart(static_cast<QHttpMultiPart::ContentType>(contentType), static_cast<QObject*>(parent));
}

void* QHttpMultiPart_NewQHttpMultiPart(void* parent){
	return new QHttpMultiPart(static_cast<QObject*>(parent));
}

void QHttpMultiPart_Append(void* ptr, void* httpPart){
	static_cast<QHttpMultiPart*>(ptr)->append(*static_cast<QHttpPart*>(httpPart));
}

void* QHttpMultiPart_Boundary(void* ptr){
	return new QByteArray(static_cast<QHttpMultiPart*>(ptr)->boundary());
}

void QHttpMultiPart_SetBoundary(void* ptr, void* boundary){
	static_cast<QHttpMultiPart*>(ptr)->setBoundary(*static_cast<QByteArray*>(boundary));
}

void QHttpMultiPart_SetContentType(void* ptr, int contentType){
	static_cast<QHttpMultiPart*>(ptr)->setContentType(static_cast<QHttpMultiPart::ContentType>(contentType));
}

void QHttpMultiPart_DestroyQHttpMultiPart(void* ptr){
	static_cast<QHttpMultiPart*>(ptr)->~QHttpMultiPart();
}

