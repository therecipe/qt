#include "qsqlresult.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSqlResult>
#include "_cgo_export.h"

class MyQSqlResult: public QSqlResult {
public:
};

char* QSqlResult_Handle(QtObjectPtr ptr){
	return static_cast<QSqlResult*>(ptr)->handle().toString().toUtf8().data();
}

void QSqlResult_DestroyQSqlResult(QtObjectPtr ptr){
	static_cast<QSqlResult*>(ptr)->~QSqlResult();
}

