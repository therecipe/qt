#include "qtemporaryfile.h"
#include <QObject>
#include <QFile>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QTemporaryFile>
#include "_cgo_export.h"

class MyQTemporaryFile: public QTemporaryFile {
public:
};

void* QTemporaryFile_NewQTemporaryFile(){
	return new QTemporaryFile();
}

void* QTemporaryFile_NewQTemporaryFile3(void* parent){
	return new QTemporaryFile(static_cast<QObject*>(parent));
}

void* QTemporaryFile_NewQTemporaryFile2(char* templateName){
	return new QTemporaryFile(QString(templateName));
}

void* QTemporaryFile_NewQTemporaryFile4(char* templateName, void* parent){
	return new QTemporaryFile(QString(templateName), static_cast<QObject*>(parent));
}

int QTemporaryFile_AutoRemove(void* ptr){
	return static_cast<QTemporaryFile*>(ptr)->autoRemove();
}

void* QTemporaryFile_QTemporaryFile_CreateNativeFile(void* file){
	return QTemporaryFile::createNativeFile(*static_cast<QFile*>(file));
}

void* QTemporaryFile_QTemporaryFile_CreateNativeFile2(char* fileName){
	return QTemporaryFile::createNativeFile(QString(fileName));
}

char* QTemporaryFile_FileName(void* ptr){
	return static_cast<QTemporaryFile*>(ptr)->fileName().toUtf8().data();
}

char* QTemporaryFile_FileTemplate(void* ptr){
	return static_cast<QTemporaryFile*>(ptr)->fileTemplate().toUtf8().data();
}

int QTemporaryFile_Open(void* ptr){
	return static_cast<QTemporaryFile*>(ptr)->open();
}

void QTemporaryFile_SetAutoRemove(void* ptr, int b){
	static_cast<QTemporaryFile*>(ptr)->setAutoRemove(b != 0);
}

void QTemporaryFile_SetFileTemplate(void* ptr, char* name){
	static_cast<QTemporaryFile*>(ptr)->setFileTemplate(QString(name));
}

void QTemporaryFile_DestroyQTemporaryFile(void* ptr){
	static_cast<QTemporaryFile*>(ptr)->~QTemporaryFile();
}

