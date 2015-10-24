#include "qxmlstreamentityresolver.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QXmlStreamEntityResolver>
#include "_cgo_export.h"

class MyQXmlStreamEntityResolver: public QXmlStreamEntityResolver {
public:
};

char* QXmlStreamEntityResolver_ResolveUndeclaredEntity(QtObjectPtr ptr, char* name){
	return static_cast<QXmlStreamEntityResolver*>(ptr)->resolveUndeclaredEntity(QString(name)).toUtf8().data();
}

void QXmlStreamEntityResolver_DestroyQXmlStreamEntityResolver(QtObjectPtr ptr){
	static_cast<QXmlStreamEntityResolver*>(ptr)->~QXmlStreamEntityResolver();
}

