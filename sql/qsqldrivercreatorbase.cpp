#include "qsqldrivercreatorbase.h"
#include <QModelIndex>
#include <QSqlDriver>
#include <QSqlDriverCreator>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QSqlDriverCreatorBase>
#include "_cgo_export.h"

class MyQSqlDriverCreatorBase: public QSqlDriverCreatorBase {
public:
};

QtObjectPtr QSqlDriverCreatorBase_CreateObject(QtObjectPtr ptr){
	return static_cast<QSqlDriverCreatorBase*>(ptr)->createObject();
}

void QSqlDriverCreatorBase_DestroyQSqlDriverCreatorBase(QtObjectPtr ptr){
	static_cast<QSqlDriverCreatorBase*>(ptr)->~QSqlDriverCreatorBase();
}

