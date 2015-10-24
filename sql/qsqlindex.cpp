#include "qsqlindex.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSqlField>
#include <QSqlIndex>
#include "_cgo_export.h"

class MyQSqlIndex: public QSqlIndex {
public:
};

QtObjectPtr QSqlIndex_NewQSqlIndex2(QtObjectPtr other){
	return new QSqlIndex(*static_cast<QSqlIndex*>(other));
}

QtObjectPtr QSqlIndex_NewQSqlIndex(char* cursorname, char* name){
	return new QSqlIndex(QString(cursorname), QString(name));
}

void QSqlIndex_Append(QtObjectPtr ptr, QtObjectPtr field){
	static_cast<QSqlIndex*>(ptr)->append(*static_cast<QSqlField*>(field));
}

void QSqlIndex_Append2(QtObjectPtr ptr, QtObjectPtr field, int desc){
	static_cast<QSqlIndex*>(ptr)->append(*static_cast<QSqlField*>(field), desc != 0);
}

char* QSqlIndex_CursorName(QtObjectPtr ptr){
	return static_cast<QSqlIndex*>(ptr)->cursorName().toUtf8().data();
}

int QSqlIndex_IsDescending(QtObjectPtr ptr, int i){
	return static_cast<QSqlIndex*>(ptr)->isDescending(i);
}

char* QSqlIndex_Name(QtObjectPtr ptr){
	return static_cast<QSqlIndex*>(ptr)->name().toUtf8().data();
}

void QSqlIndex_SetCursorName(QtObjectPtr ptr, char* cursorName){
	static_cast<QSqlIndex*>(ptr)->setCursorName(QString(cursorName));
}

void QSqlIndex_SetDescending(QtObjectPtr ptr, int i, int desc){
	static_cast<QSqlIndex*>(ptr)->setDescending(i, desc != 0);
}

void QSqlIndex_SetName(QtObjectPtr ptr, char* name){
	static_cast<QSqlIndex*>(ptr)->setName(QString(name));
}

void QSqlIndex_DestroyQSqlIndex(QtObjectPtr ptr){
	static_cast<QSqlIndex*>(ptr)->~QSqlIndex();
}

