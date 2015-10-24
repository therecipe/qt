#include "qquickimageprovider.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QQmlImageProviderBase>
#include <QQuickImageProvider>
#include "_cgo_export.h"

class MyQQuickImageProvider: public QQuickImageProvider {
public:
};

QtObjectPtr QQuickImageProvider_NewQQuickImageProvider(int ty, int flags){
	return new QQuickImageProvider(static_cast<QQmlImageProviderBase::ImageType>(ty), static_cast<QQmlImageProviderBase::Flag>(flags));
}

int QQuickImageProvider_Flags(QtObjectPtr ptr){
	return static_cast<QQuickImageProvider*>(ptr)->flags();
}

int QQuickImageProvider_ImageType(QtObjectPtr ptr){
	return static_cast<QQuickImageProvider*>(ptr)->imageType();
}

void QQuickImageProvider_DestroyQQuickImageProvider(QtObjectPtr ptr){
	static_cast<QQuickImageProvider*>(ptr)->~QQuickImageProvider();
}

