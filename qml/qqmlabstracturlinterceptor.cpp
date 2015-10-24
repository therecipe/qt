#include "qqmlabstracturlinterceptor.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QQmlAbstractUrlInterceptor>
#include "_cgo_export.h"

class MyQQmlAbstractUrlInterceptor: public QQmlAbstractUrlInterceptor {
public:
};

char* QQmlAbstractUrlInterceptor_Intercept(QtObjectPtr ptr, char* url, int ty){
	return static_cast<QQmlAbstractUrlInterceptor*>(ptr)->intercept(QUrl(QString(url)), static_cast<QQmlAbstractUrlInterceptor::DataType>(ty)).toString().toUtf8().data();
}

void QQmlAbstractUrlInterceptor_DestroyQQmlAbstractUrlInterceptor(QtObjectPtr ptr){
	static_cast<QQmlAbstractUrlInterceptor*>(ptr)->~QQmlAbstractUrlInterceptor();
}

