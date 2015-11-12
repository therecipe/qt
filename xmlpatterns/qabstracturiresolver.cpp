#include "qabstracturiresolver.h"
#include <QUrl>
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QAbstractUriResolver>
#include "_cgo_export.h"

class MyQAbstractUriResolver: public QAbstractUriResolver {
public:
};

void QAbstractUriResolver_DestroyQAbstractUriResolver(void* ptr){
	static_cast<QAbstractUriResolver*>(ptr)->~QAbstractUriResolver();
}

