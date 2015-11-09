#include "qxmlentityresolver.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QXmlEntityResolver>
#include "_cgo_export.h"

class MyQXmlEntityResolver: public QXmlEntityResolver {
public:
};

char* QXmlEntityResolver_ErrorString(void* ptr){
	return static_cast<QXmlEntityResolver*>(ptr)->errorString().toUtf8().data();
}

void QXmlEntityResolver_DestroyQXmlEntityResolver(void* ptr){
	static_cast<QXmlEntityResolver*>(ptr)->~QXmlEntityResolver();
}

