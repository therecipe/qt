#include "qtemporaryfile.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QFile>
#include <QObject>
#include <QTemporaryFile>
#include "_cgo_export.h"

class MyQTemporaryFile: public QTemporaryFile {
public:
};

QtObjectPtr QTemporaryFile_NewQTemporaryFile(){
	return new QTemporaryFile();
}

QtObjectPtr QTemporaryFile_NewQTemporaryFile3(QtObjectPtr parent){
	return new QTemporaryFile(static_cast<QObject*>(parent));
}

QtObjectPtr QTemporaryFile_NewQTemporaryFile2(char* templateName){
	return new QTemporaryFile(QString(templateName));
}

QtObjectPtr QTemporaryFile_NewQTemporaryFile4(char* templateName, QtObjectPtr parent){
	return new QTemporaryFile(QString(templateName), static_cast<QObject*>(parent));
}

int QTemporaryFile_AutoRemove(QtObjectPtr ptr){
	return static_cast<QTemporaryFile*>(ptr)->autoRemove();
}

QtObjectPtr QTemporaryFile_QTemporaryFile_CreateNativeFile(QtObjectPtr file){
	return QTemporaryFile::createNativeFile(*static_cast<QFile*>(file));
}

QtObjectPtr QTemporaryFile_QTemporaryFile_CreateNativeFile2(char* fileName){
	return QTemporaryFile::createNativeFile(QString(fileName));
}

char* QTemporaryFile_FileName(QtObjectPtr ptr){
	return static_cast<QTemporaryFile*>(ptr)->fileName().toUtf8().data();
}

char* QTemporaryFile_FileTemplate(QtObjectPtr ptr){
	return static_cast<QTemporaryFile*>(ptr)->fileTemplate().toUtf8().data();
}

int QTemporaryFile_Open(QtObjectPtr ptr){
	return static_cast<QTemporaryFile*>(ptr)->open();
}

void QTemporaryFile_SetAutoRemove(QtObjectPtr ptr, int b){
	static_cast<QTemporaryFile*>(ptr)->setAutoRemove(b != 0);
}

void QTemporaryFile_SetFileTemplate(QtObjectPtr ptr, char* name){
	static_cast<QTemporaryFile*>(ptr)->setFileTemplate(QString(name));
}

void QTemporaryFile_DestroyQTemporaryFile(QtObjectPtr ptr){
	static_cast<QTemporaryFile*>(ptr)->~QTemporaryFile();
}

