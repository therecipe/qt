#include "qqmlimageproviderbase.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QQmlImageProviderBase>
#include "_cgo_export.h"

class MyQQmlImageProviderBase: public QQmlImageProviderBase {
public:
};

int QQmlImageProviderBase_Flags(QtObjectPtr ptr){
	return static_cast<QQmlImageProviderBase*>(ptr)->flags();
}

int QQmlImageProviderBase_ImageType(QtObjectPtr ptr){
	return static_cast<QQmlImageProviderBase*>(ptr)->imageType();
}

