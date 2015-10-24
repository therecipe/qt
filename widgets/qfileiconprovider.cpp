#include "qfileiconprovider.h"
#include <QFile>
#include <QFileInfo>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QFileIconProvider>
#include "_cgo_export.h"

class MyQFileIconProvider: public QFileIconProvider {
public:
};

QtObjectPtr QFileIconProvider_NewQFileIconProvider(){
	return new QFileIconProvider();
}

int QFileIconProvider_Options(QtObjectPtr ptr){
	return static_cast<QFileIconProvider*>(ptr)->options();
}

void QFileIconProvider_SetOptions(QtObjectPtr ptr, int options){
	static_cast<QFileIconProvider*>(ptr)->setOptions(static_cast<QFileIconProvider::Option>(options));
}

char* QFileIconProvider_Type(QtObjectPtr ptr, QtObjectPtr info){
	return static_cast<QFileIconProvider*>(ptr)->type(*static_cast<QFileInfo*>(info)).toUtf8().data();
}

void QFileIconProvider_DestroyQFileIconProvider(QtObjectPtr ptr){
	static_cast<QFileIconProvider*>(ptr)->~QFileIconProvider();
}

