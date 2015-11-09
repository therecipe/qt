#include "qfilesystemwatcher.h"
#include <QFile>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QFileSystemWatcher>
#include "_cgo_export.h"

class MyQFileSystemWatcher: public QFileSystemWatcher {
public:
void Signal_DirectoryChanged(const QString & path){callbackQFileSystemWatcherDirectoryChanged(this->objectName().toUtf8().data(), path.toUtf8().data());};
void Signal_FileChanged(const QString & path){callbackQFileSystemWatcherFileChanged(this->objectName().toUtf8().data(), path.toUtf8().data());};
};

char* QFileSystemWatcher_Directories(void* ptr){
	return static_cast<QFileSystemWatcher*>(ptr)->directories().join("|").toUtf8().data();
}

char* QFileSystemWatcher_Files(void* ptr){
	return static_cast<QFileSystemWatcher*>(ptr)->files().join("|").toUtf8().data();
}

void* QFileSystemWatcher_NewQFileSystemWatcher(void* parent){
	return new QFileSystemWatcher(static_cast<QObject*>(parent));
}

void* QFileSystemWatcher_NewQFileSystemWatcher2(char* paths, void* parent){
	return new QFileSystemWatcher(QString(paths).split("|", QString::SkipEmptyParts), static_cast<QObject*>(parent));
}

int QFileSystemWatcher_AddPath(void* ptr, char* path){
	return static_cast<QFileSystemWatcher*>(ptr)->addPath(QString(path));
}

char* QFileSystemWatcher_AddPaths(void* ptr, char* paths){
	return static_cast<QFileSystemWatcher*>(ptr)->addPaths(QString(paths).split("|", QString::SkipEmptyParts)).join("|").toUtf8().data();
}

void QFileSystemWatcher_ConnectDirectoryChanged(void* ptr){
	QObject::connect(static_cast<QFileSystemWatcher*>(ptr), &QFileSystemWatcher::directoryChanged, static_cast<MyQFileSystemWatcher*>(ptr), static_cast<void (MyQFileSystemWatcher::*)(const QString &)>(&MyQFileSystemWatcher::Signal_DirectoryChanged));;
}

void QFileSystemWatcher_DisconnectDirectoryChanged(void* ptr){
	QObject::disconnect(static_cast<QFileSystemWatcher*>(ptr), &QFileSystemWatcher::directoryChanged, static_cast<MyQFileSystemWatcher*>(ptr), static_cast<void (MyQFileSystemWatcher::*)(const QString &)>(&MyQFileSystemWatcher::Signal_DirectoryChanged));;
}

void QFileSystemWatcher_ConnectFileChanged(void* ptr){
	QObject::connect(static_cast<QFileSystemWatcher*>(ptr), &QFileSystemWatcher::fileChanged, static_cast<MyQFileSystemWatcher*>(ptr), static_cast<void (MyQFileSystemWatcher::*)(const QString &)>(&MyQFileSystemWatcher::Signal_FileChanged));;
}

void QFileSystemWatcher_DisconnectFileChanged(void* ptr){
	QObject::disconnect(static_cast<QFileSystemWatcher*>(ptr), &QFileSystemWatcher::fileChanged, static_cast<MyQFileSystemWatcher*>(ptr), static_cast<void (MyQFileSystemWatcher::*)(const QString &)>(&MyQFileSystemWatcher::Signal_FileChanged));;
}

int QFileSystemWatcher_RemovePath(void* ptr, char* path){
	return static_cast<QFileSystemWatcher*>(ptr)->removePath(QString(path));
}

char* QFileSystemWatcher_RemovePaths(void* ptr, char* paths){
	return static_cast<QFileSystemWatcher*>(ptr)->removePaths(QString(paths).split("|", QString::SkipEmptyParts)).join("|").toUtf8().data();
}

void QFileSystemWatcher_DestroyQFileSystemWatcher(void* ptr){
	static_cast<QFileSystemWatcher*>(ptr)->~QFileSystemWatcher();
}

