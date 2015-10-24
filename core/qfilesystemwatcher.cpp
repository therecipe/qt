#include "qfilesystemwatcher.h"
#include <QModelIndex>
#include <QFile>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QFileSystemWatcher>
#include "_cgo_export.h"

class MyQFileSystemWatcher: public QFileSystemWatcher {
public:
void Signal_DirectoryChanged(const QString & path){callbackQFileSystemWatcherDirectoryChanged(this->objectName().toUtf8().data(), path.toUtf8().data());};
void Signal_FileChanged(const QString & path){callbackQFileSystemWatcherFileChanged(this->objectName().toUtf8().data(), path.toUtf8().data());};
};

char* QFileSystemWatcher_Directories(QtObjectPtr ptr){
	return static_cast<QFileSystemWatcher*>(ptr)->directories().join("|").toUtf8().data();
}

char* QFileSystemWatcher_Files(QtObjectPtr ptr){
	return static_cast<QFileSystemWatcher*>(ptr)->files().join("|").toUtf8().data();
}

QtObjectPtr QFileSystemWatcher_NewQFileSystemWatcher(QtObjectPtr parent){
	return new QFileSystemWatcher(static_cast<QObject*>(parent));
}

QtObjectPtr QFileSystemWatcher_NewQFileSystemWatcher2(char* paths, QtObjectPtr parent){
	return new QFileSystemWatcher(QString(paths).split("|", QString::SkipEmptyParts), static_cast<QObject*>(parent));
}

int QFileSystemWatcher_AddPath(QtObjectPtr ptr, char* path){
	return static_cast<QFileSystemWatcher*>(ptr)->addPath(QString(path));
}

char* QFileSystemWatcher_AddPaths(QtObjectPtr ptr, char* paths){
	return static_cast<QFileSystemWatcher*>(ptr)->addPaths(QString(paths).split("|", QString::SkipEmptyParts)).join("|").toUtf8().data();
}

void QFileSystemWatcher_ConnectDirectoryChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QFileSystemWatcher*>(ptr), &QFileSystemWatcher::directoryChanged, static_cast<MyQFileSystemWatcher*>(ptr), static_cast<void (MyQFileSystemWatcher::*)(const QString &)>(&MyQFileSystemWatcher::Signal_DirectoryChanged));;
}

void QFileSystemWatcher_DisconnectDirectoryChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QFileSystemWatcher*>(ptr), &QFileSystemWatcher::directoryChanged, static_cast<MyQFileSystemWatcher*>(ptr), static_cast<void (MyQFileSystemWatcher::*)(const QString &)>(&MyQFileSystemWatcher::Signal_DirectoryChanged));;
}

void QFileSystemWatcher_ConnectFileChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QFileSystemWatcher*>(ptr), &QFileSystemWatcher::fileChanged, static_cast<MyQFileSystemWatcher*>(ptr), static_cast<void (MyQFileSystemWatcher::*)(const QString &)>(&MyQFileSystemWatcher::Signal_FileChanged));;
}

void QFileSystemWatcher_DisconnectFileChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QFileSystemWatcher*>(ptr), &QFileSystemWatcher::fileChanged, static_cast<MyQFileSystemWatcher*>(ptr), static_cast<void (MyQFileSystemWatcher::*)(const QString &)>(&MyQFileSystemWatcher::Signal_FileChanged));;
}

int QFileSystemWatcher_RemovePath(QtObjectPtr ptr, char* path){
	return static_cast<QFileSystemWatcher*>(ptr)->removePath(QString(path));
}

char* QFileSystemWatcher_RemovePaths(QtObjectPtr ptr, char* paths){
	return static_cast<QFileSystemWatcher*>(ptr)->removePaths(QString(paths).split("|", QString::SkipEmptyParts)).join("|").toUtf8().data();
}

void QFileSystemWatcher_DestroyQFileSystemWatcher(QtObjectPtr ptr){
	static_cast<QFileSystemWatcher*>(ptr)->~QFileSystemWatcher();
}

