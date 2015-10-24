#include "qsqldriverplugin.h"
#include <QModelIndex>
#include <QSqlDriver>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QSqlDriverPlugin>
#include "_cgo_export.h"

class MyQSqlDriverPlugin: public QSqlDriverPlugin {
public:
};

QtObjectPtr QSqlDriverPlugin_Create(QtObjectPtr ptr, char* key){
	return static_cast<QSqlDriverPlugin*>(ptr)->create(QString(key));
}

void QSqlDriverPlugin_DestroyQSqlDriverPlugin(QtObjectPtr ptr){
	static_cast<QSqlDriverPlugin*>(ptr)->~QSqlDriverPlugin();
}

