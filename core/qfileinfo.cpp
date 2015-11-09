#include "qfileinfo.h"
#include <QUrl>
#include <QModelIndex>
#include <QDateTime>
#include <QDir>
#include <QDate>
#include <QFile>
#include <QString>
#include <QVariant>
#include <QFileInfo>
#include "_cgo_export.h"

class MyQFileInfo: public QFileInfo {
public:
};

void* QFileInfo_NewQFileInfo(){
	return new QFileInfo();
}

void* QFileInfo_NewQFileInfo5(void* dir, char* file){
	return new QFileInfo(*static_cast<QDir*>(dir), QString(file));
}

void* QFileInfo_NewQFileInfo4(void* file){
	return new QFileInfo(*static_cast<QFile*>(file));
}

void* QFileInfo_NewQFileInfo6(void* fileinfo){
	return new QFileInfo(*static_cast<QFileInfo*>(fileinfo));
}

void* QFileInfo_NewQFileInfo3(char* file){
	return new QFileInfo(QString(file));
}

void* QFileInfo_AbsoluteDir(void* ptr){
	return new QDir(static_cast<QFileInfo*>(ptr)->absoluteDir());
}

char* QFileInfo_AbsoluteFilePath(void* ptr){
	return static_cast<QFileInfo*>(ptr)->absoluteFilePath().toUtf8().data();
}

char* QFileInfo_AbsolutePath(void* ptr){
	return static_cast<QFileInfo*>(ptr)->absolutePath().toUtf8().data();
}

char* QFileInfo_BaseName(void* ptr){
	return static_cast<QFileInfo*>(ptr)->baseName().toUtf8().data();
}

char* QFileInfo_BundleName(void* ptr){
	return static_cast<QFileInfo*>(ptr)->bundleName().toUtf8().data();
}

int QFileInfo_Caching(void* ptr){
	return static_cast<QFileInfo*>(ptr)->caching();
}

char* QFileInfo_CanonicalFilePath(void* ptr){
	return static_cast<QFileInfo*>(ptr)->canonicalFilePath().toUtf8().data();
}

char* QFileInfo_CanonicalPath(void* ptr){
	return static_cast<QFileInfo*>(ptr)->canonicalPath().toUtf8().data();
}

char* QFileInfo_CompleteBaseName(void* ptr){
	return static_cast<QFileInfo*>(ptr)->completeBaseName().toUtf8().data();
}

char* QFileInfo_CompleteSuffix(void* ptr){
	return static_cast<QFileInfo*>(ptr)->completeSuffix().toUtf8().data();
}

void* QFileInfo_Created(void* ptr){
	return new QDateTime(static_cast<QFileInfo*>(ptr)->created());
}

void* QFileInfo_Dir(void* ptr){
	return new QDir(static_cast<QFileInfo*>(ptr)->dir());
}

int QFileInfo_QFileInfo_Exists2(char* file){
	return QFileInfo::exists(QString(file));
}

int QFileInfo_Exists(void* ptr){
	return static_cast<QFileInfo*>(ptr)->exists();
}

char* QFileInfo_FileName(void* ptr){
	return static_cast<QFileInfo*>(ptr)->fileName().toUtf8().data();
}

char* QFileInfo_FilePath(void* ptr){
	return static_cast<QFileInfo*>(ptr)->filePath().toUtf8().data();
}

char* QFileInfo_Group(void* ptr){
	return static_cast<QFileInfo*>(ptr)->group().toUtf8().data();
}

int QFileInfo_IsAbsolute(void* ptr){
	return static_cast<QFileInfo*>(ptr)->isAbsolute();
}

int QFileInfo_IsBundle(void* ptr){
	return static_cast<QFileInfo*>(ptr)->isBundle();
}

int QFileInfo_IsDir(void* ptr){
	return static_cast<QFileInfo*>(ptr)->isDir();
}

int QFileInfo_IsExecutable(void* ptr){
	return static_cast<QFileInfo*>(ptr)->isExecutable();
}

int QFileInfo_IsFile(void* ptr){
	return static_cast<QFileInfo*>(ptr)->isFile();
}

int QFileInfo_IsHidden(void* ptr){
	return static_cast<QFileInfo*>(ptr)->isHidden();
}

int QFileInfo_IsNativePath(void* ptr){
	return static_cast<QFileInfo*>(ptr)->isNativePath();
}

int QFileInfo_IsReadable(void* ptr){
	return static_cast<QFileInfo*>(ptr)->isReadable();
}

int QFileInfo_IsRelative(void* ptr){
	return static_cast<QFileInfo*>(ptr)->isRelative();
}

int QFileInfo_IsRoot(void* ptr){
	return static_cast<QFileInfo*>(ptr)->isRoot();
}

int QFileInfo_IsSymLink(void* ptr){
	return static_cast<QFileInfo*>(ptr)->isSymLink();
}

int QFileInfo_IsWritable(void* ptr){
	return static_cast<QFileInfo*>(ptr)->isWritable();
}

void* QFileInfo_LastModified(void* ptr){
	return new QDateTime(static_cast<QFileInfo*>(ptr)->lastModified());
}

void* QFileInfo_LastRead(void* ptr){
	return new QDateTime(static_cast<QFileInfo*>(ptr)->lastRead());
}

int QFileInfo_MakeAbsolute(void* ptr){
	return static_cast<QFileInfo*>(ptr)->makeAbsolute();
}

char* QFileInfo_Owner(void* ptr){
	return static_cast<QFileInfo*>(ptr)->owner().toUtf8().data();
}

char* QFileInfo_Path(void* ptr){
	return static_cast<QFileInfo*>(ptr)->path().toUtf8().data();
}

void QFileInfo_Refresh(void* ptr){
	static_cast<QFileInfo*>(ptr)->refresh();
}

void QFileInfo_SetCaching(void* ptr, int enable){
	static_cast<QFileInfo*>(ptr)->setCaching(enable != 0);
}

void QFileInfo_SetFile3(void* ptr, void* dir, char* file){
	static_cast<QFileInfo*>(ptr)->setFile(*static_cast<QDir*>(dir), QString(file));
}

void QFileInfo_SetFile2(void* ptr, void* file){
	static_cast<QFileInfo*>(ptr)->setFile(*static_cast<QFile*>(file));
}

void QFileInfo_SetFile(void* ptr, char* file){
	static_cast<QFileInfo*>(ptr)->setFile(QString(file));
}

char* QFileInfo_Suffix(void* ptr){
	return static_cast<QFileInfo*>(ptr)->suffix().toUtf8().data();
}

void QFileInfo_Swap(void* ptr, void* other){
	static_cast<QFileInfo*>(ptr)->swap(*static_cast<QFileInfo*>(other));
}

char* QFileInfo_SymLinkTarget(void* ptr){
	return static_cast<QFileInfo*>(ptr)->symLinkTarget().toUtf8().data();
}

void QFileInfo_DestroyQFileInfo(void* ptr){
	static_cast<QFileInfo*>(ptr)->~QFileInfo();
}

