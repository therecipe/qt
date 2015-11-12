#include "qsqldrivercreatorbase.h"
#include <QSqlDriver>
#include <QSqlDriverCreator>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSqlDriverCreatorBase>
#include "_cgo_export.h"

class MyQSqlDriverCreatorBase: public QSqlDriverCreatorBase {
public:
};

void* QSqlDriverCreatorBase_CreateObject(void* ptr){
	return static_cast<QSqlDriverCreatorBase*>(ptr)->createObject();
}

void QSqlDriverCreatorBase_DestroyQSqlDriverCreatorBase(void* ptr){
	static_cast<QSqlDriverCreatorBase*>(ptr)->~QSqlDriverCreatorBase();
}

