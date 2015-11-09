#include "qquickimageprovider.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QQmlImageProviderBase>
#include <QString>
#include <QQuickImageProvider>
#include "_cgo_export.h"

class MyQQuickImageProvider: public QQuickImageProvider {
public:
};

void* QQuickImageProvider_NewQQuickImageProvider(int ty, int flags){
	return new QQuickImageProvider(static_cast<QQmlImageProviderBase::ImageType>(ty), static_cast<QQmlImageProviderBase::Flag>(flags));
}

int QQuickImageProvider_Flags(void* ptr){
	return static_cast<QQuickImageProvider*>(ptr)->flags();
}

int QQuickImageProvider_ImageType(void* ptr){
	return static_cast<QQuickImageProvider*>(ptr)->imageType();
}

void QQuickImageProvider_DestroyQQuickImageProvider(void* ptr){
	static_cast<QQuickImageProvider*>(ptr)->~QQuickImageProvider();
}

