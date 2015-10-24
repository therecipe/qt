#include "qtemporarydir.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QTemporaryDir>
#include "_cgo_export.h"

class MyQTemporaryDir: public QTemporaryDir {
public:
};

QtObjectPtr QTemporaryDir_NewQTemporaryDir(){
	return new QTemporaryDir();
}

QtObjectPtr QTemporaryDir_NewQTemporaryDir2(char* templatePath){
	return new QTemporaryDir(QString(templatePath));
}

int QTemporaryDir_AutoRemove(QtObjectPtr ptr){
	return static_cast<QTemporaryDir*>(ptr)->autoRemove();
}

int QTemporaryDir_IsValid(QtObjectPtr ptr){
	return static_cast<QTemporaryDir*>(ptr)->isValid();
}

char* QTemporaryDir_Path(QtObjectPtr ptr){
	return static_cast<QTemporaryDir*>(ptr)->path().toUtf8().data();
}

void QTemporaryDir_SetAutoRemove(QtObjectPtr ptr, int b){
	static_cast<QTemporaryDir*>(ptr)->setAutoRemove(b != 0);
}

void QTemporaryDir_DestroyQTemporaryDir(QtObjectPtr ptr){
	static_cast<QTemporaryDir*>(ptr)->~QTemporaryDir();
}

