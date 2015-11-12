#include "qsqldriverplugin.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSqlDriver>
#include <QSqlDriverPlugin>
#include "_cgo_export.h"

class MyQSqlDriverPlugin: public QSqlDriverPlugin {
public:
};

void* QSqlDriverPlugin_Create(void* ptr, char* key){
	return static_cast<QSqlDriverPlugin*>(ptr)->create(QString(key));
}

void QSqlDriverPlugin_DestroyQSqlDriverPlugin(void* ptr){
	static_cast<QSqlDriverPlugin*>(ptr)->~QSqlDriverPlugin();
}

