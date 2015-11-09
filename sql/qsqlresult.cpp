#include "qsqlresult.h"
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QSqlResult>
#include "_cgo_export.h"

class MyQSqlResult: public QSqlResult {
public:
};

void* QSqlResult_Handle(void* ptr){
	return new QVariant(static_cast<QSqlResult*>(ptr)->handle());
}

void QSqlResult_DestroyQSqlResult(void* ptr){
	static_cast<QSqlResult*>(ptr)->~QSqlResult();
}

