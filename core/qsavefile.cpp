#include "qsavefile.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QIODevice>
#include <QObject>
#include <QSaveFile>
#include "_cgo_export.h"

class MyQSaveFile: public QSaveFile {
public:
};

void* QSaveFile_NewQSaveFile2(void* parent){
	return new QSaveFile(static_cast<QObject*>(parent));
}

void* QSaveFile_NewQSaveFile(char* name){
	return new QSaveFile(QString(name));
}

void* QSaveFile_NewQSaveFile3(char* name, void* parent){
	return new QSaveFile(QString(name), static_cast<QObject*>(parent));
}

void QSaveFile_CancelWriting(void* ptr){
	static_cast<QSaveFile*>(ptr)->cancelWriting();
}

int QSaveFile_Commit(void* ptr){
	return static_cast<QSaveFile*>(ptr)->commit();
}

int QSaveFile_DirectWriteFallback(void* ptr){
	return static_cast<QSaveFile*>(ptr)->directWriteFallback();
}

char* QSaveFile_FileName(void* ptr){
	return static_cast<QSaveFile*>(ptr)->fileName().toUtf8().data();
}

int QSaveFile_Open(void* ptr, int mode){
	return static_cast<QSaveFile*>(ptr)->open(static_cast<QIODevice::OpenModeFlag>(mode));
}

void QSaveFile_SetDirectWriteFallback(void* ptr, int enabled){
	static_cast<QSaveFile*>(ptr)->setDirectWriteFallback(enabled != 0);
}

void QSaveFile_SetFileName(void* ptr, char* name){
	static_cast<QSaveFile*>(ptr)->setFileName(QString(name));
}

void QSaveFile_DestroyQSaveFile(void* ptr){
	static_cast<QSaveFile*>(ptr)->~QSaveFile();
}

