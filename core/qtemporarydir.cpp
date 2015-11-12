#include "qtemporarydir.h"
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QTemporaryDir>
#include "_cgo_export.h"

class MyQTemporaryDir: public QTemporaryDir {
public:
};

void* QTemporaryDir_NewQTemporaryDir(){
	return new QTemporaryDir();
}

void* QTemporaryDir_NewQTemporaryDir2(char* templatePath){
	return new QTemporaryDir(QString(templatePath));
}

int QTemporaryDir_AutoRemove(void* ptr){
	return static_cast<QTemporaryDir*>(ptr)->autoRemove();
}

int QTemporaryDir_IsValid(void* ptr){
	return static_cast<QTemporaryDir*>(ptr)->isValid();
}

char* QTemporaryDir_Path(void* ptr){
	return static_cast<QTemporaryDir*>(ptr)->path().toUtf8().data();
}

void QTemporaryDir_SetAutoRemove(void* ptr, int b){
	static_cast<QTemporaryDir*>(ptr)->setAutoRemove(b != 0);
}

void QTemporaryDir_DestroyQTemporaryDir(void* ptr){
	static_cast<QTemporaryDir*>(ptr)->~QTemporaryDir();
}

