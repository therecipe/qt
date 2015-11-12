#include "qfileiconprovider.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QFile>
#include <QFileInfo>
#include <QFileIconProvider>
#include "_cgo_export.h"

class MyQFileIconProvider: public QFileIconProvider {
public:
};

void* QFileIconProvider_NewQFileIconProvider(){
	return new QFileIconProvider();
}

int QFileIconProvider_Options(void* ptr){
	return static_cast<QFileIconProvider*>(ptr)->options();
}

void QFileIconProvider_SetOptions(void* ptr, int options){
	static_cast<QFileIconProvider*>(ptr)->setOptions(static_cast<QFileIconProvider::Option>(options));
}

char* QFileIconProvider_Type(void* ptr, void* info){
	return static_cast<QFileIconProvider*>(ptr)->type(*static_cast<QFileInfo*>(info)).toUtf8().data();
}

void QFileIconProvider_DestroyQFileIconProvider(void* ptr){
	static_cast<QFileIconProvider*>(ptr)->~QFileIconProvider();
}

