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

void QAbstractUriResolver_DestroyQAbstractUriResolver(void* ptr){
	static_cast<QAbstractUriResolver*>(ptr)->~QAbstractUriResolver();
}

