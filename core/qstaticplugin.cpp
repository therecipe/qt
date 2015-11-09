#include "qstaticplugin.h"
#include <QModelIndex>
#include <QJsonObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QStaticPlugin>
#include "_cgo_export.h"

class MyQStaticPlugin: public QStaticPlugin {
public:
};

void* QStaticPlugin_Instance(void* ptr){
	return static_cast<QStaticPlugin*>(ptr)->instance();
}

void* QStaticPlugin_MetaData(void* ptr){
	return new QJsonObject(static_cast<QStaticPlugin*>(ptr)->metaData());
}

