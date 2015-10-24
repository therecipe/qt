#include "qsavefile.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QIODevice>
#include <QString>
#include <QSaveFile>
#include "_cgo_export.h"

class MyQSaveFile: public QSaveFile {
public:
};

QtObjectPtr QSaveFile_NewQSaveFile2(QtObjectPtr parent){
	return new QSaveFile(static_cast<QObject*>(parent));
}

QtObjectPtr QSaveFile_NewQSaveFile(char* name){
	return new QSaveFile(QString(name));
}

QtObjectPtr QSaveFile_NewQSaveFile3(char* name, QtObjectPtr parent){
	return new QSaveFile(QString(name), static_cast<QObject*>(parent));
}

void QSaveFile_CancelWriting(QtObjectPtr ptr){
	static_cast<QSaveFile*>(ptr)->cancelWriting();
}

int QSaveFile_Commit(QtObjectPtr ptr){
	return static_cast<QSaveFile*>(ptr)->commit();
}

int QSaveFile_DirectWriteFallback(QtObjectPtr ptr){
	return static_cast<QSaveFile*>(ptr)->directWriteFallback();
}

char* QSaveFile_FileName(QtObjectPtr ptr){
	return static_cast<QSaveFile*>(ptr)->fileName().toUtf8().data();
}

int QSaveFile_Open(QtObjectPtr ptr, int mode){
	return static_cast<QSaveFile*>(ptr)->open(static_cast<QIODevice::OpenModeFlag>(mode));
}

void QSaveFile_SetDirectWriteFallback(QtObjectPtr ptr, int enabled){
	static_cast<QSaveFile*>(ptr)->setDirectWriteFallback(enabled != 0);
}

void QSaveFile_SetFileName(QtObjectPtr ptr, char* name){
	static_cast<QSaveFile*>(ptr)->setFileName(QString(name));
}

void QSaveFile_DestroyQSaveFile(QtObjectPtr ptr){
	static_cast<QSaveFile*>(ptr)->~QSaveFile();
}

