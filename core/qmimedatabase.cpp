#include "qmimedatabase.h"
#include <QUrl>
#include <QModelIndex>
#include <QMimeData>
#include <QString>
#include <QVariant>
#include <QMimeDatabase>
#include "_cgo_export.h"

class MyQMimeDatabase: public QMimeDatabase {
public:
};

void* QMimeDatabase_NewQMimeDatabase(){
	return new QMimeDatabase();
}

void QMimeDatabase_DestroyQMimeDatabase(void* ptr){
	static_cast<QMimeDatabase*>(ptr)->~QMimeDatabase();
}

char* QMimeDatabase_SuffixForFileName(void* ptr, char* fileName){
	return static_cast<QMimeDatabase*>(ptr)->suffixForFileName(QString(fileName)).toUtf8().data();
}

