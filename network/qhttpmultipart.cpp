#include "qhttpmultipart.h"
#include <QHttpPart>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QByteArray>
#include <QHttpMultiPart>
#include "_cgo_export.h"

class MyQHttpMultiPart: public QHttpMultiPart {
public:
};

QtObjectPtr QHttpMultiPart_NewQHttpMultiPart2(int contentType, QtObjectPtr parent){
	return new QHttpMultiPart(static_cast<QHttpMultiPart::ContentType>(contentType), static_cast<QObject*>(parent));
}

QtObjectPtr QHttpMultiPart_NewQHttpMultiPart(QtObjectPtr parent){
	return new QHttpMultiPart(static_cast<QObject*>(parent));
}

void QHttpMultiPart_Append(QtObjectPtr ptr, QtObjectPtr httpPart){
	static_cast<QHttpMultiPart*>(ptr)->append(*static_cast<QHttpPart*>(httpPart));
}

void QHttpMultiPart_SetBoundary(QtObjectPtr ptr, QtObjectPtr boundary){
	static_cast<QHttpMultiPart*>(ptr)->setBoundary(*static_cast<QByteArray*>(boundary));
}

void QHttpMultiPart_SetContentType(QtObjectPtr ptr, int contentType){
	static_cast<QHttpMultiPart*>(ptr)->setContentType(static_cast<QHttpMultiPart::ContentType>(contentType));
}

void QHttpMultiPart_DestroyQHttpMultiPart(QtObjectPtr ptr){
	static_cast<QHttpMultiPart*>(ptr)->~QHttpMultiPart();
}

