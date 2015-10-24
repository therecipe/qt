#include "qlibrary.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QLibrary>
#include "_cgo_export.h"

class MyQLibrary: public QLibrary {
public:
};

char* QLibrary_FileName(QtObjectPtr ptr){
	return static_cast<QLibrary*>(ptr)->fileName().toUtf8().data();
}

int QLibrary_LoadHints(QtObjectPtr ptr){
	return static_cast<QLibrary*>(ptr)->loadHints();
}

void QLibrary_SetFileName(QtObjectPtr ptr, char* fileName){
	static_cast<QLibrary*>(ptr)->setFileName(QString(fileName));
}

void QLibrary_SetFileNameAndVersion(QtObjectPtr ptr, char* fileName, int versionNumber){
	static_cast<QLibrary*>(ptr)->setFileNameAndVersion(QString(fileName), versionNumber);
}

void QLibrary_SetLoadHints(QtObjectPtr ptr, int hints){
	static_cast<QLibrary*>(ptr)->setLoadHints(static_cast<QLibrary::LoadHint>(hints));
}

QtObjectPtr QLibrary_NewQLibrary(QtObjectPtr parent){
	return new QLibrary(static_cast<QObject*>(parent));
}

QtObjectPtr QLibrary_NewQLibrary2(char* fileName, QtObjectPtr parent){
	return new QLibrary(QString(fileName), static_cast<QObject*>(parent));
}

QtObjectPtr QLibrary_NewQLibrary4(char* fileName, char* version, QtObjectPtr parent){
	return new QLibrary(QString(fileName), QString(version), static_cast<QObject*>(parent));
}

QtObjectPtr QLibrary_NewQLibrary3(char* fileName, int verNum, QtObjectPtr parent){
	return new QLibrary(QString(fileName), verNum, static_cast<QObject*>(parent));
}

char* QLibrary_ErrorString(QtObjectPtr ptr){
	return static_cast<QLibrary*>(ptr)->errorString().toUtf8().data();
}

int QLibrary_QLibrary_IsLibrary(char* fileName){
	return QLibrary::isLibrary(QString(fileName));
}

int QLibrary_IsLoaded(QtObjectPtr ptr){
	return static_cast<QLibrary*>(ptr)->isLoaded();
}

int QLibrary_Load(QtObjectPtr ptr){
	return static_cast<QLibrary*>(ptr)->load();
}

void QLibrary_SetFileNameAndVersion2(QtObjectPtr ptr, char* fileName, char* version){
	static_cast<QLibrary*>(ptr)->setFileNameAndVersion(QString(fileName), QString(version));
}

int QLibrary_Unload(QtObjectPtr ptr){
	return static_cast<QLibrary*>(ptr)->unload();
}

void QLibrary_DestroyQLibrary(QtObjectPtr ptr){
	static_cast<QLibrary*>(ptr)->~QLibrary();
}

