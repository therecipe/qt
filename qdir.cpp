#include "qdir.h"
#include <QDir>
#include "cgoexport.h"


//Public Types
int QDir_Dirs() { return QDir::Dirs; }
int QDir_AllDirs() { return QDir::AllDirs; }
int QDir_Files() { return QDir::Files; }
int QDir_Drives() { return QDir::Drives; }
int QDir_NoSymLinks() { return QDir::NoSymLinks; }
int QDir_NoDotAndDotDot() { return QDir::NoDotAndDotDot; }
int QDir_NoDot() { return QDir::NoDot; }
int QDir_NoDotDot() { return QDir::NoDotDot; }
int QDir_AllEntries() { return QDir::AllEntries; }
int QDir_Readable() { return QDir::Readable; }
int QDir_Writable() { return QDir::Writable; }
int QDir_Executable() { return QDir::Executable; }
int QDir_Modified() { return QDir::Modified; }
int QDir_Hidden() { return QDir::Hidden; }
int QDir_System() { return QDir::System; }
int QDir_Name() { return QDir::Name; }
int QDir_Time() { return QDir::Time; }
int QDir_Size() { return QDir::Size; }
int QDir_Type() { return QDir::Type; }
int QDir_Unsorted() { return QDir::Unsorted; }
int QDir_NoSort() { return QDir::NoSort; }
int QDir_DirsFirst() { return QDir::DirsFirst; }
int QDir_DirsLast() { return QDir::DirsLast; }
int QDir_Reversed() { return QDir::Reversed; }
int QDir_IgnoreCase() { return QDir::IgnoreCase; }
int QDir_LocaleAware() { return QDir::LocaleAware; }
//Static Public Members
void QDir_AddSearchPath_String_String(char* prefix, char* path)
{
	QDir::addSearchPath(QString(prefix), QString(path));
}

char* QDir_CleanPath_String(char* path)
{
	return QDir::cleanPath(QString(path)).toUtf8().data();
}

char* QDir_CurrentPath()
{
	return QDir::currentPath().toUtf8().data();
}

char* QDir_FromNativeSeparators_String(char* pathName)
{
	return QDir::fromNativeSeparators(QString(pathName)).toUtf8().data();
}

char* QDir_HomePath()
{
	return QDir::homePath().toUtf8().data();
}

int QDir_IsAbsolutePath_String(char* path)
{
	return QDir::isAbsolutePath(QString(path));
}

int QDir_IsRelativePath_String(char* path)
{
	return QDir::isRelativePath(QString(path));
}

int QDir_Match_String_String(char* filter, char* fileName)
{
	return QDir::match(QString(filter), QString(fileName));
}

int QDir_Match_QStringList_String(char* filters, char* fileName)
{
	return QDir::match(QString(filters).split("|", QString::SkipEmptyParts), QString(fileName));
}

char* QDir_RootPath()
{
	return QDir::rootPath().toUtf8().data();
}

char* QDir_SearchPaths_String(char* prefix)
{
	return QDir::searchPaths(QString(prefix)).join("|").toUtf8().data();
}

int QDir_SetCurrent_String(char* path)
{
	return QDir::setCurrent(QString(path));
}

void QDir_SetSearchPaths_String_QStringList(char* prefix, char* searchPaths)
{
	QDir::setSearchPaths(QString(prefix), QString(searchPaths).split("|", QString::SkipEmptyParts));
}

char* QDir_TempPath()
{
	return QDir::tempPath().toUtf8().data();
}

char* QDir_ToNativeSeparators_String(char* pathName)
{
	return QDir::toNativeSeparators(QString(pathName)).toUtf8().data();
}

