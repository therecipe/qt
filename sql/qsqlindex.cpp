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

void* QSqlIndex_NewQSqlIndex2(void* other){
	return new QSqlIndex(*static_cast<QSqlIndex*>(other));
}

void* QSqlIndex_NewQSqlIndex(char* cursorname, char* name){
	return new QSqlIndex(QString(cursorname), QString(name));
}

void QSqlIndex_Append(void* ptr, void* field){
	static_cast<QSqlIndex*>(ptr)->append(*static_cast<QSqlField*>(field));
}

void QSqlIndex_Append2(void* ptr, void* field, int desc){
	static_cast<QSqlIndex*>(ptr)->append(*static_cast<QSqlField*>(field), desc != 0);
}

char* QSqlIndex_CursorName(void* ptr){
	return static_cast<QSqlIndex*>(ptr)->cursorName().toUtf8().data();
}

int QSqlIndex_IsDescending(void* ptr, int i){
	return static_cast<QSqlIndex*>(ptr)->isDescending(i);
}

char* QSqlIndex_Name(void* ptr){
	return static_cast<QSqlIndex*>(ptr)->name().toUtf8().data();
}

void QSqlIndex_SetCursorName(void* ptr, char* cursorName){
	static_cast<QSqlIndex*>(ptr)->setCursorName(QString(cursorName));
}

void QSqlIndex_SetDescending(void* ptr, int i, int desc){
	static_cast<QSqlIndex*>(ptr)->setDescending(i, desc != 0);
}

void QSqlIndex_SetName(void* ptr, char* name){
	static_cast<QSqlIndex*>(ptr)->setName(QString(name));
}

void QSqlIndex_DestroyQSqlIndex(void* ptr){
	static_cast<QSqlIndex*>(ptr)->~QSqlIndex();
}

