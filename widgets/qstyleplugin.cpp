#include "qstyleplugin.h"
#include <QUrl>
#include <QModelIndex>
#include <QStyle>
#include <QString>
#include <QVariant>
#include <QStylePlugin>
#include "_cgo_export.h"

class MyQStylePlugin: public QStylePlugin {
public:
};

void* QStylePlugin_Create(void* ptr, char* key){
	return static_cast<QStylePlugin*>(ptr)->create(QString(key));
}

void QStylePlugin_DestroyQStylePlugin(void* ptr){
	static_cast<QStylePlugin*>(ptr)->~QStylePlugin();
}

