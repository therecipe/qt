#include "qlibrary.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QString>
#include <QLibrary>
#include "_cgo_export.h"

class MyQLibrary: public QLibrary {
public:
};

char* QLibrary_FileName(void* ptr){
	return static_cast<QLibrary*>(ptr)->fileName().toUtf8().data();
}

int QLibrary_LoadHints(void* ptr){
	return static_cast<QLibrary*>(ptr)->loadHints();
}

void QLibrary_SetFileName(void* ptr, char* fileName){
	static_cast<QLibrary*>(ptr)->setFileName(QString(fileName));
}

void QLibrary_SetFileNameAndVersion(void* ptr, char* fileName, int versionNumber){
	static_cast<QLibrary*>(ptr)->setFileNameAndVersion(QString(fileName), versionNumber);
}

void QLibrary_SetLoadHints(void* ptr, int hints){
	static_cast<QLibrary*>(ptr)->setLoadHints(static_cast<QLibrary::LoadHint>(hints));
}

void* QLibrary_NewQLibrary(void* parent){
	return new QLibrary(static_cast<QObject*>(parent));
}

void* QLibrary_NewQLibrary2(char* fileName, void* parent){
	return new QLibrary(QString(fileName), static_cast<QObject*>(parent));
}

void* QLibrary_NewQLibrary4(char* fileName, char* version, void* parent){
	return new QLibrary(QString(fileName), QString(version), static_cast<QObject*>(parent));
}

void* QLibrary_NewQLibrary3(char* fileName, int verNum, void* parent){
	return new QLibrary(QString(fileName), verNum, static_cast<QObject*>(parent));
}

char* QLibrary_ErrorString(void* ptr){
	return static_cast<QLibrary*>(ptr)->errorString().toUtf8().data();
}

int QLibrary_QLibrary_IsLibrary(char* fileName){
	return QLibrary::isLibrary(QString(fileName));
}

int QLibrary_IsLoaded(void* ptr){
	return static_cast<QLibrary*>(ptr)->isLoaded();
}

int QLibrary_Load(void* ptr){
	return static_cast<QLibrary*>(ptr)->load();
}

void QLibrary_SetFileNameAndVersion2(void* ptr, char* fileName, char* version){
	static_cast<QLibrary*>(ptr)->setFileNameAndVersion(QString(fileName), QString(version));
}

int QLibrary_Unload(void* ptr){
	return static_cast<QLibrary*>(ptr)->unload();
}

void QLibrary_DestroyQLibrary(void* ptr){
	static_cast<QLibrary*>(ptr)->~QLibrary();
}

