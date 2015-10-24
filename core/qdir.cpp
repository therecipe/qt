#include "qdir.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QDir>
#include "_cgo_export.h"

class MyQDir: public QDir {
public:
};

QtObjectPtr QDir_NewQDir(QtObjectPtr dir){
	return new QDir(*static_cast<QDir*>(dir));
}

QtObjectPtr QDir_NewQDir2(char* path){
	return new QDir(QString(path));
}

QtObjectPtr QDir_NewQDir3(char* path, char* nameFilter, int sort, int filters){
	return new QDir(QString(path), QString(nameFilter), static_cast<QDir::SortFlag>(sort), static_cast<QDir::Filter>(filters));
}

char* QDir_AbsoluteFilePath(QtObjectPtr ptr, char* fileName){
	return static_cast<QDir*>(ptr)->absoluteFilePath(QString(fileName)).toUtf8().data();
}

char* QDir_AbsolutePath(QtObjectPtr ptr){
	return static_cast<QDir*>(ptr)->absolutePath().toUtf8().data();
}

void QDir_QDir_AddSearchPath(char* prefix, char* path){
	QDir::addSearchPath(QString(prefix), QString(path));
}

char* QDir_CanonicalPath(QtObjectPtr ptr){
	return static_cast<QDir*>(ptr)->canonicalPath().toUtf8().data();
}

int QDir_Cd(QtObjectPtr ptr, char* dirName){
	return static_cast<QDir*>(ptr)->cd(QString(dirName));
}

int QDir_CdUp(QtObjectPtr ptr){
	return static_cast<QDir*>(ptr)->cdUp();
}

char* QDir_QDir_CleanPath(char* path){
	return QDir::cleanPath(QString(path)).toUtf8().data();
}

char* QDir_QDir_CurrentPath(){
	return QDir::currentPath().toUtf8().data();
}

char* QDir_DirName(QtObjectPtr ptr){
	return static_cast<QDir*>(ptr)->dirName().toUtf8().data();
}

char* QDir_EntryList2(QtObjectPtr ptr, int filters, int sort){
	return static_cast<QDir*>(ptr)->entryList(static_cast<QDir::Filter>(filters), static_cast<QDir::SortFlag>(sort)).join("|").toUtf8().data();
}

char* QDir_EntryList(QtObjectPtr ptr, char* nameFilters, int filters, int sort){
	return static_cast<QDir*>(ptr)->entryList(QString(nameFilters).split("|", QString::SkipEmptyParts), static_cast<QDir::Filter>(filters), static_cast<QDir::SortFlag>(sort)).join("|").toUtf8().data();
}

int QDir_Exists2(QtObjectPtr ptr){
	return static_cast<QDir*>(ptr)->exists();
}

int QDir_Exists(QtObjectPtr ptr, char* name){
	return static_cast<QDir*>(ptr)->exists(QString(name));
}

char* QDir_FilePath(QtObjectPtr ptr, char* fileName){
	return static_cast<QDir*>(ptr)->filePath(QString(fileName)).toUtf8().data();
}

int QDir_Filter(QtObjectPtr ptr){
	return static_cast<QDir*>(ptr)->filter();
}

char* QDir_QDir_FromNativeSeparators(char* pathName){
	return QDir::fromNativeSeparators(QString(pathName)).toUtf8().data();
}

char* QDir_QDir_HomePath(){
	return QDir::homePath().toUtf8().data();
}

int QDir_IsAbsolute(QtObjectPtr ptr){
	return static_cast<QDir*>(ptr)->isAbsolute();
}

int QDir_QDir_IsAbsolutePath(char* path){
	return QDir::isAbsolutePath(QString(path));
}

int QDir_IsReadable(QtObjectPtr ptr){
	return static_cast<QDir*>(ptr)->isReadable();
}

int QDir_IsRelative(QtObjectPtr ptr){
	return static_cast<QDir*>(ptr)->isRelative();
}

int QDir_QDir_IsRelativePath(char* path){
	return QDir::isRelativePath(QString(path));
}

int QDir_IsRoot(QtObjectPtr ptr){
	return static_cast<QDir*>(ptr)->isRoot();
}

int QDir_MakeAbsolute(QtObjectPtr ptr){
	return static_cast<QDir*>(ptr)->makeAbsolute();
}

int QDir_QDir_Match(char* filter, char* fileName){
	return QDir::match(QString(filter), QString(fileName));
}

int QDir_QDir_Match2(char* filters, char* fileName){
	return QDir::match(QString(filters).split("|", QString::SkipEmptyParts), QString(fileName));
}

int QDir_Mkdir(QtObjectPtr ptr, char* dirName){
	return static_cast<QDir*>(ptr)->mkdir(QString(dirName));
}

int QDir_Mkpath(QtObjectPtr ptr, char* dirPath){
	return static_cast<QDir*>(ptr)->mkpath(QString(dirPath));
}

char* QDir_NameFilters(QtObjectPtr ptr){
	return static_cast<QDir*>(ptr)->nameFilters().join("|").toUtf8().data();
}

char* QDir_Path(QtObjectPtr ptr){
	return static_cast<QDir*>(ptr)->path().toUtf8().data();
}

void QDir_Refresh(QtObjectPtr ptr){
	static_cast<QDir*>(ptr)->refresh();
}

char* QDir_RelativeFilePath(QtObjectPtr ptr, char* fileName){
	return static_cast<QDir*>(ptr)->relativeFilePath(QString(fileName)).toUtf8().data();
}

int QDir_RemoveRecursively(QtObjectPtr ptr){
	return static_cast<QDir*>(ptr)->removeRecursively();
}

int QDir_Rename(QtObjectPtr ptr, char* oldName, char* newName){
	return static_cast<QDir*>(ptr)->rename(QString(oldName), QString(newName));
}

int QDir_Rmdir(QtObjectPtr ptr, char* dirName){
	return static_cast<QDir*>(ptr)->rmdir(QString(dirName));
}

int QDir_Rmpath(QtObjectPtr ptr, char* dirPath){
	return static_cast<QDir*>(ptr)->rmpath(QString(dirPath));
}

char* QDir_QDir_RootPath(){
	return QDir::rootPath().toUtf8().data();
}

char* QDir_QDir_SearchPaths(char* prefix){
	return QDir::searchPaths(QString(prefix)).join("|").toUtf8().data();
}

int QDir_QDir_SetCurrent(char* path){
	return QDir::setCurrent(QString(path));
}

void QDir_SetFilter(QtObjectPtr ptr, int filters){
	static_cast<QDir*>(ptr)->setFilter(static_cast<QDir::Filter>(filters));
}

void QDir_SetNameFilters(QtObjectPtr ptr, char* nameFilters){
	static_cast<QDir*>(ptr)->setNameFilters(QString(nameFilters).split("|", QString::SkipEmptyParts));
}

void QDir_SetPath(QtObjectPtr ptr, char* path){
	static_cast<QDir*>(ptr)->setPath(QString(path));
}

void QDir_QDir_SetSearchPaths(char* prefix, char* searchPaths){
	QDir::setSearchPaths(QString(prefix), QString(searchPaths).split("|", QString::SkipEmptyParts));
}

void QDir_SetSorting(QtObjectPtr ptr, int sort){
	static_cast<QDir*>(ptr)->setSorting(static_cast<QDir::SortFlag>(sort));
}

int QDir_Sorting(QtObjectPtr ptr){
	return static_cast<QDir*>(ptr)->sorting();
}

void QDir_Swap(QtObjectPtr ptr, QtObjectPtr other){
	static_cast<QDir*>(ptr)->swap(*static_cast<QDir*>(other));
}

char* QDir_QDir_TempPath(){
	return QDir::tempPath().toUtf8().data();
}

char* QDir_QDir_ToNativeSeparators(char* pathName){
	return QDir::toNativeSeparators(QString(pathName)).toUtf8().data();
}

void QDir_DestroyQDir(QtObjectPtr ptr){
	static_cast<QDir*>(ptr)->~QDir();
}

