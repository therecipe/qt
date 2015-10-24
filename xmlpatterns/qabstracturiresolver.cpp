#include "qabstracturiresolver.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QAbstractUriResolver>
#include "_cgo_export.h"

class MyQAbstractUriResolver: public QAbstractUriResolver {
public:
};

char* QAbstractUriResolver_Resolve(QtObjectPtr ptr, char* relative, char* baseURI){
	return static_cast<QAbstractUriResolver*>(ptr)->resolve(QUrl(QString(relative)), QUrl(QString(baseURI))).toString().toUtf8().data();
}

void QAbstractUriResolver_DestroyQAbstractUriResolver(QtObjectPtr ptr){
	static_cast<QAbstractUriResolver*>(ptr)->~QAbstractUriResolver();
}

