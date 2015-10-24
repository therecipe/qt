#include "qmimedatabase.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMimeData>
#include <QString>
#include <QMimeDatabase>
#include "_cgo_export.h"

class MyQMimeDatabase: public QMimeDatabase {
public:
};

QtObjectPtr QMimeDatabase_NewQMimeDatabase(){
	return new QMimeDatabase();
}

void QMimeDatabase_DestroyQMimeDatabase(QtObjectPtr ptr){
	static_cast<QMimeDatabase*>(ptr)->~QMimeDatabase();
}

char* QMimeDatabase_SuffixForFileName(QtObjectPtr ptr, char* fileName){
	return static_cast<QMimeDatabase*>(ptr)->suffixForFileName(QString(fileName)).toUtf8().data();
}

