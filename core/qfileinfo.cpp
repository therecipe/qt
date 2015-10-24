#include "qfileinfo.h"
#include <QDir>
#include <QFile>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QFileInfo>
#include "_cgo_export.h"

class MyQFileInfo: public QFileInfo {
public:
};

QtObjectPtr QFileInfo_NewQFileInfo(){
	return new QFileInfo();
}

QtObjectPtr QFileInfo_NewQFileInfo5(QtObjectPtr dir, char* file){
	return new QFileInfo(*static_cast<QDir*>(dir), QString(file));
}

QtObjectPtr QFileInfo_NewQFileInfo4(QtObjectPtr file){
	return new QFileInfo(*static_cast<QFile*>(file));
}

QtObjectPtr QFileInfo_NewQFileInfo6(QtObjectPtr fileinfo){
	return new QFileInfo(*static_cast<QFileInfo*>(fileinfo));
}

QtObjectPtr QFileInfo_NewQFileInfo3(char* file){
	return new QFileInfo(QString(file));
}

char* QFileInfo_AbsoluteFilePath(QtObjectPtr ptr){
	return static_cast<QFileInfo*>(ptr)->absoluteFilePath().toUtf8().data();
}

char* QFileInfo_AbsolutePath(QtObjectPtr ptr){
	return static_cast<QFileInfo*>(ptr)->absolutePath().toUtf8().data();
}

char* QFileInfo_BaseName(QtObjectPtr ptr){
	return static_cast<QFileInfo*>(ptr)->baseName().toUtf8().data();
}

char* QFileInfo_BundleName(QtObjectPtr ptr){
	return static_cast<QFileInfo*>(ptr)->bundleName().toUtf8().data();
}

int QFileInfo_Caching(QtObjectPtr ptr){
	return static_cast<QFileInfo*>(ptr)->caching();
}

char* QFileInfo_CanonicalFilePath(QtObjectPtr ptr){
	return static_cast<QFileInfo*>(ptr)->canonicalFilePath().toUtf8().data();
}

char* QFileInfo_CanonicalPath(QtObjectPtr ptr){
	return static_cast<QFileInfo*>(ptr)->canonicalPath().toUtf8().data();
}

char* QFileInfo_CompleteBaseName(QtObjectPtr ptr){
	return static_cast<QFileInfo*>(ptr)->completeBaseName().toUtf8().data();
}

char* QFileInfo_CompleteSuffix(QtObjectPtr ptr){
	return static_cast<QFileInfo*>(ptr)->completeSuffix().toUtf8().data();
}

int QFileInfo_QFileInfo_Exists2(char* file){
	return QFileInfo::exists(QString(file));
}

int QFileInfo_Exists(QtObjectPtr ptr){
	return static_cast<QFileInfo*>(ptr)->exists();
}

char* QFileInfo_FileName(QtObjectPtr ptr){
	return static_cast<QFileInfo*>(ptr)->fileName().toUtf8().data();
}

char* QFileInfo_FilePath(QtObjectPtr ptr){
	return static_cast<QFileInfo*>(ptr)->filePath().toUtf8().data();
}

char* QFileInfo_Group(QtObjectPtr ptr){
	return static_cast<QFileInfo*>(ptr)->group().toUtf8().data();
}

int QFileInfo_IsAbsolute(QtObjectPtr ptr){
	return static_cast<QFileInfo*>(ptr)->isAbsolute();
}

int QFileInfo_IsBundle(QtObjectPtr ptr){
	return static_cast<QFileInfo*>(ptr)->isBundle();
}

int QFileInfo_IsDir(QtObjectPtr ptr){
	return static_cast<QFileInfo*>(ptr)->isDir();
}

int QFileInfo_IsExecutable(QtObjectPtr ptr){
	return static_cast<QFileInfo*>(ptr)->isExecutable();
}

int QFileInfo_IsFile(QtObjectPtr ptr){
	return static_cast<QFileInfo*>(ptr)->isFile();
}

int QFileInfo_IsHidden(QtObjectPtr ptr){
	return static_cast<QFileInfo*>(ptr)->isHidden();
}

int QFileInfo_IsNativePath(QtObjectPtr ptr){
	return static_cast<QFileInfo*>(ptr)->isNativePath();
}

int QFileInfo_IsReadable(QtObjectPtr ptr){
	return static_cast<QFileInfo*>(ptr)->isReadable();
}

int QFileInfo_IsRelative(QtObjectPtr ptr){
	return static_cast<QFileInfo*>(ptr)->isRelative();
}

int QFileInfo_IsRoot(QtObjectPtr ptr){
	return static_cast<QFileInfo*>(ptr)->isRoot();
}

int QFileInfo_IsSymLink(QtObjectPtr ptr){
	return static_cast<QFileInfo*>(ptr)->isSymLink();
}

int QFileInfo_IsWritable(QtObjectPtr ptr){
	return static_cast<QFileInfo*>(ptr)->isWritable();
}

int QFileInfo_MakeAbsolute(QtObjectPtr ptr){
	return static_cast<QFileInfo*>(ptr)->makeAbsolute();
}

char* QFileInfo_Owner(QtObjectPtr ptr){
	return static_cast<QFileInfo*>(ptr)->owner().toUtf8().data();
}

char* QFileInfo_Path(QtObjectPtr ptr){
	return static_cast<QFileInfo*>(ptr)->path().toUtf8().data();
}

void QFileInfo_Refresh(QtObjectPtr ptr){
	static_cast<QFileInfo*>(ptr)->refresh();
}

void QFileInfo_SetCaching(QtObjectPtr ptr, int enable){
	static_cast<QFileInfo*>(ptr)->setCaching(enable != 0);
}

void QFileInfo_SetFile3(QtObjectPtr ptr, QtObjectPtr dir, char* file){
	static_cast<QFileInfo*>(ptr)->setFile(*static_cast<QDir*>(dir), QString(file));
}

void QFileInfo_SetFile2(QtObjectPtr ptr, QtObjectPtr file){
	static_cast<QFileInfo*>(ptr)->setFile(*static_cast<QFile*>(file));
}

void QFileInfo_SetFile(QtObjectPtr ptr, char* file){
	static_cast<QFileInfo*>(ptr)->setFile(QString(file));
}

char* QFileInfo_Suffix(QtObjectPtr ptr){
	return static_cast<QFileInfo*>(ptr)->suffix().toUtf8().data();
}

void QFileInfo_Swap(QtObjectPtr ptr, QtObjectPtr other){
	static_cast<QFileInfo*>(ptr)->swap(*static_cast<QFileInfo*>(other));
}

char* QFileInfo_SymLinkTarget(QtObjectPtr ptr){
	return static_cast<QFileInfo*>(ptr)->symLinkTarget().toUtf8().data();
}

void QFileInfo_DestroyQFileInfo(QtObjectPtr ptr){
	static_cast<QFileInfo*>(ptr)->~QFileInfo();
}

