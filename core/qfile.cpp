#include "qfile.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QByteArray>
#include <QIODevice>
#include <QFileDevice>
#include <QFile>
#include "_cgo_export.h"

class MyQFile: public QFile {
public:
};

QtObjectPtr QFile_NewQFile3(QtObjectPtr parent){
	return new QFile(static_cast<QObject*>(parent));
}

QtObjectPtr QFile_NewQFile(char* name){
	return new QFile(QString(name));
}

QtObjectPtr QFile_NewQFile4(char* name, QtObjectPtr parent){
	return new QFile(QString(name), static_cast<QObject*>(parent));
}

int QFile_QFile_Copy2(char* fileName, char* newName){
	return QFile::copy(QString(fileName), QString(newName));
}

int QFile_Copy(QtObjectPtr ptr, char* newName){
	return static_cast<QFile*>(ptr)->copy(QString(newName));
}

char* QFile_QFile_DecodeName(QtObjectPtr localFileName){
	return QFile::decodeName(*static_cast<QByteArray*>(localFileName)).toUtf8().data();
}

char* QFile_QFile_DecodeName2(char* localFileName){
	return QFile::decodeName(const_cast<const char*>(localFileName)).toUtf8().data();
}

int QFile_QFile_Exists(char* fileName){
	return QFile::exists(QString(fileName));
}

int QFile_Exists2(QtObjectPtr ptr){
	return static_cast<QFile*>(ptr)->exists();
}

char* QFile_FileName(QtObjectPtr ptr){
	return static_cast<QFile*>(ptr)->fileName().toUtf8().data();
}

int QFile_QFile_Link2(char* fileName, char* linkName){
	return QFile::link(QString(fileName), QString(linkName));
}

int QFile_Link(QtObjectPtr ptr, char* linkName){
	return static_cast<QFile*>(ptr)->link(QString(linkName));
}

int QFile_Open(QtObjectPtr ptr, int mode){
	return static_cast<QFile*>(ptr)->open(static_cast<QIODevice::OpenModeFlag>(mode));
}

int QFile_Open3(QtObjectPtr ptr, int fd, int mode, int handleFlags){
	return static_cast<QFile*>(ptr)->open(fd, static_cast<QIODevice::OpenModeFlag>(mode), static_cast<QFileDevice::FileHandleFlag>(handleFlags));
}

int QFile_QFile_Permissions2(char* fileName){
	return QFile::permissions(QString(fileName));
}

int QFile_Permissions(QtObjectPtr ptr){
	return static_cast<QFile*>(ptr)->permissions();
}

int QFile_Rename(QtObjectPtr ptr, char* newName){
	return static_cast<QFile*>(ptr)->rename(QString(newName));
}

int QFile_QFile_Rename2(char* oldName, char* newName){
	return QFile::rename(QString(oldName), QString(newName));
}

void QFile_SetFileName(QtObjectPtr ptr, char* name){
	static_cast<QFile*>(ptr)->setFileName(QString(name));
}

int QFile_SetPermissions(QtObjectPtr ptr, int permissions){
	return static_cast<QFile*>(ptr)->setPermissions(static_cast<QFileDevice::Permission>(permissions));
}

int QFile_QFile_SetPermissions2(char* fileName, int permissions){
	return QFile::setPermissions(QString(fileName), static_cast<QFileDevice::Permission>(permissions));
}

char* QFile_QFile_SymLinkTarget(char* fileName){
	return QFile::symLinkTarget(QString(fileName)).toUtf8().data();
}

char* QFile_SymLinkTarget2(QtObjectPtr ptr){
	return static_cast<QFile*>(ptr)->symLinkTarget().toUtf8().data();
}

void QFile_DestroyQFile(QtObjectPtr ptr){
	static_cast<QFile*>(ptr)->~QFile();
}

