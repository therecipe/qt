#include "qfilesystemmodel.h"
#include <QDir>
#include <QUrl>
#include <QModelIndex>
#include <QFileIconProvider>
#include <QObject>
#include <QMimeData>
#include <QFile>
#include <QString>
#include <QVariant>
#include <QFileSystemModel>
#include "_cgo_export.h"

class MyQFileSystemModel: public QFileSystemModel {
public:
void Signal_DirectoryLoaded(const QString & path){callbackQFileSystemModelDirectoryLoaded(this->objectName().toUtf8().data(), path.toUtf8().data());};
void Signal_FileRenamed(const QString & path, const QString & oldName, const QString & newName){callbackQFileSystemModelFileRenamed(this->objectName().toUtf8().data(), path.toUtf8().data(), oldName.toUtf8().data(), newName.toUtf8().data());};
void Signal_RootPathChanged(const QString & newPath){callbackQFileSystemModelRootPathChanged(this->objectName().toUtf8().data(), newPath.toUtf8().data());};
};

int QFileSystemModel_FilePathRole_Type(){
	return QFileSystemModel::FilePathRole;
}

int QFileSystemModel_FileNameRole_Type(){
	return QFileSystemModel::FileNameRole;
}

int QFileSystemModel_FilePermissions_Type(){
	return QFileSystemModel::FilePermissions;
}

int QFileSystemModel_IsReadOnly(QtObjectPtr ptr){
	return static_cast<QFileSystemModel*>(ptr)->isReadOnly();
}

int QFileSystemModel_NameFilterDisables(QtObjectPtr ptr){
	return static_cast<QFileSystemModel*>(ptr)->nameFilterDisables();
}

int QFileSystemModel_ResolveSymlinks(QtObjectPtr ptr){
	return static_cast<QFileSystemModel*>(ptr)->resolveSymlinks();
}

int QFileSystemModel_Rmdir(QtObjectPtr ptr, QtObjectPtr index){
	return static_cast<QFileSystemModel*>(ptr)->rmdir(*static_cast<QModelIndex*>(index));
}

void QFileSystemModel_SetNameFilterDisables(QtObjectPtr ptr, int enable){
	static_cast<QFileSystemModel*>(ptr)->setNameFilterDisables(enable != 0);
}

void QFileSystemModel_SetReadOnly(QtObjectPtr ptr, int enable){
	static_cast<QFileSystemModel*>(ptr)->setReadOnly(enable != 0);
}

void QFileSystemModel_SetResolveSymlinks(QtObjectPtr ptr, int enable){
	static_cast<QFileSystemModel*>(ptr)->setResolveSymlinks(enable != 0);
}

QtObjectPtr QFileSystemModel_NewQFileSystemModel(QtObjectPtr parent){
	return new QFileSystemModel(static_cast<QObject*>(parent));
}

int QFileSystemModel_CanFetchMore(QtObjectPtr ptr, QtObjectPtr parent){
	return static_cast<QFileSystemModel*>(ptr)->canFetchMore(*static_cast<QModelIndex*>(parent));
}

int QFileSystemModel_ColumnCount(QtObjectPtr ptr, QtObjectPtr parent){
	return static_cast<QFileSystemModel*>(ptr)->columnCount(*static_cast<QModelIndex*>(parent));
}

char* QFileSystemModel_Data(QtObjectPtr ptr, QtObjectPtr index, int role){
	return static_cast<QFileSystemModel*>(ptr)->data(*static_cast<QModelIndex*>(index), role).toString().toUtf8().data();
}

void QFileSystemModel_ConnectDirectoryLoaded(QtObjectPtr ptr){
	QObject::connect(static_cast<QFileSystemModel*>(ptr), static_cast<void (QFileSystemModel::*)(const QString &)>(&QFileSystemModel::directoryLoaded), static_cast<MyQFileSystemModel*>(ptr), static_cast<void (MyQFileSystemModel::*)(const QString &)>(&MyQFileSystemModel::Signal_DirectoryLoaded));;
}

void QFileSystemModel_DisconnectDirectoryLoaded(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QFileSystemModel*>(ptr), static_cast<void (QFileSystemModel::*)(const QString &)>(&QFileSystemModel::directoryLoaded), static_cast<MyQFileSystemModel*>(ptr), static_cast<void (MyQFileSystemModel::*)(const QString &)>(&MyQFileSystemModel::Signal_DirectoryLoaded));;
}

int QFileSystemModel_DropMimeData(QtObjectPtr ptr, QtObjectPtr data, int action, int row, int column, QtObjectPtr parent){
	return static_cast<QFileSystemModel*>(ptr)->dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

void QFileSystemModel_FetchMore(QtObjectPtr ptr, QtObjectPtr parent){
	static_cast<QFileSystemModel*>(ptr)->fetchMore(*static_cast<QModelIndex*>(parent));
}

char* QFileSystemModel_FileName(QtObjectPtr ptr, QtObjectPtr index){
	return static_cast<QFileSystemModel*>(ptr)->fileName(*static_cast<QModelIndex*>(index)).toUtf8().data();
}

char* QFileSystemModel_FilePath(QtObjectPtr ptr, QtObjectPtr index){
	return static_cast<QFileSystemModel*>(ptr)->filePath(*static_cast<QModelIndex*>(index)).toUtf8().data();
}

void QFileSystemModel_ConnectFileRenamed(QtObjectPtr ptr){
	QObject::connect(static_cast<QFileSystemModel*>(ptr), static_cast<void (QFileSystemModel::*)(const QString &, const QString &, const QString &)>(&QFileSystemModel::fileRenamed), static_cast<MyQFileSystemModel*>(ptr), static_cast<void (MyQFileSystemModel::*)(const QString &, const QString &, const QString &)>(&MyQFileSystemModel::Signal_FileRenamed));;
}

void QFileSystemModel_DisconnectFileRenamed(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QFileSystemModel*>(ptr), static_cast<void (QFileSystemModel::*)(const QString &, const QString &, const QString &)>(&QFileSystemModel::fileRenamed), static_cast<MyQFileSystemModel*>(ptr), static_cast<void (MyQFileSystemModel::*)(const QString &, const QString &, const QString &)>(&MyQFileSystemModel::Signal_FileRenamed));;
}

int QFileSystemModel_Filter(QtObjectPtr ptr){
	return static_cast<QFileSystemModel*>(ptr)->filter();
}

int QFileSystemModel_Flags(QtObjectPtr ptr, QtObjectPtr index){
	return static_cast<QFileSystemModel*>(ptr)->flags(*static_cast<QModelIndex*>(index));
}

int QFileSystemModel_HasChildren(QtObjectPtr ptr, QtObjectPtr parent){
	return static_cast<QFileSystemModel*>(ptr)->hasChildren(*static_cast<QModelIndex*>(parent));
}

char* QFileSystemModel_HeaderData(QtObjectPtr ptr, int section, int orientation, int role){
	return static_cast<QFileSystemModel*>(ptr)->headerData(section, static_cast<Qt::Orientation>(orientation), role).toString().toUtf8().data();
}

QtObjectPtr QFileSystemModel_IconProvider(QtObjectPtr ptr){
	return static_cast<QFileSystemModel*>(ptr)->iconProvider();
}

QtObjectPtr QFileSystemModel_Index2(QtObjectPtr ptr, char* path, int column){
	return static_cast<QFileSystemModel*>(ptr)->index(QString(path), column).internalPointer();
}

QtObjectPtr QFileSystemModel_Index(QtObjectPtr ptr, int row, int column, QtObjectPtr parent){
	return static_cast<QFileSystemModel*>(ptr)->index(row, column, *static_cast<QModelIndex*>(parent)).internalPointer();
}

int QFileSystemModel_IsDir(QtObjectPtr ptr, QtObjectPtr index){
	return static_cast<QFileSystemModel*>(ptr)->isDir(*static_cast<QModelIndex*>(index));
}

char* QFileSystemModel_MimeTypes(QtObjectPtr ptr){
	return static_cast<QFileSystemModel*>(ptr)->mimeTypes().join("|").toUtf8().data();
}

QtObjectPtr QFileSystemModel_Mkdir(QtObjectPtr ptr, QtObjectPtr parent, char* name){
	return static_cast<QFileSystemModel*>(ptr)->mkdir(*static_cast<QModelIndex*>(parent), QString(name)).internalPointer();
}

char* QFileSystemModel_MyComputer(QtObjectPtr ptr, int role){
	return static_cast<QFileSystemModel*>(ptr)->myComputer(role).toString().toUtf8().data();
}

char* QFileSystemModel_NameFilters(QtObjectPtr ptr){
	return static_cast<QFileSystemModel*>(ptr)->nameFilters().join("|").toUtf8().data();
}

QtObjectPtr QFileSystemModel_Parent(QtObjectPtr ptr, QtObjectPtr index){
	return static_cast<QFileSystemModel*>(ptr)->parent(*static_cast<QModelIndex*>(index)).internalPointer();
}

char* QFileSystemModel_RootPath(QtObjectPtr ptr){
	return static_cast<QFileSystemModel*>(ptr)->rootPath().toUtf8().data();
}

void QFileSystemModel_ConnectRootPathChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QFileSystemModel*>(ptr), static_cast<void (QFileSystemModel::*)(const QString &)>(&QFileSystemModel::rootPathChanged), static_cast<MyQFileSystemModel*>(ptr), static_cast<void (MyQFileSystemModel::*)(const QString &)>(&MyQFileSystemModel::Signal_RootPathChanged));;
}

void QFileSystemModel_DisconnectRootPathChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QFileSystemModel*>(ptr), static_cast<void (QFileSystemModel::*)(const QString &)>(&QFileSystemModel::rootPathChanged), static_cast<MyQFileSystemModel*>(ptr), static_cast<void (MyQFileSystemModel::*)(const QString &)>(&MyQFileSystemModel::Signal_RootPathChanged));;
}

int QFileSystemModel_RowCount(QtObjectPtr ptr, QtObjectPtr parent){
	return static_cast<QFileSystemModel*>(ptr)->rowCount(*static_cast<QModelIndex*>(parent));
}

int QFileSystemModel_SetData(QtObjectPtr ptr, QtObjectPtr idx, char* value, int role){
	return static_cast<QFileSystemModel*>(ptr)->setData(*static_cast<QModelIndex*>(idx), QVariant(value), role);
}

void QFileSystemModel_SetFilter(QtObjectPtr ptr, int filters){
	static_cast<QFileSystemModel*>(ptr)->setFilter(static_cast<QDir::Filter>(filters));
}

void QFileSystemModel_SetIconProvider(QtObjectPtr ptr, QtObjectPtr provider){
	static_cast<QFileSystemModel*>(ptr)->setIconProvider(static_cast<QFileIconProvider*>(provider));
}

void QFileSystemModel_SetNameFilters(QtObjectPtr ptr, char* filters){
	static_cast<QFileSystemModel*>(ptr)->setNameFilters(QString(filters).split("|", QString::SkipEmptyParts));
}

QtObjectPtr QFileSystemModel_SetRootPath(QtObjectPtr ptr, char* newPath){
	return static_cast<QFileSystemModel*>(ptr)->setRootPath(QString(newPath)).internalPointer();
}

void QFileSystemModel_Sort(QtObjectPtr ptr, int column, int order){
	static_cast<QFileSystemModel*>(ptr)->sort(column, static_cast<Qt::SortOrder>(order));
}

int QFileSystemModel_SupportedDropActions(QtObjectPtr ptr){
	return static_cast<QFileSystemModel*>(ptr)->supportedDropActions();
}

char* QFileSystemModel_Type(QtObjectPtr ptr, QtObjectPtr index){
	return static_cast<QFileSystemModel*>(ptr)->type(*static_cast<QModelIndex*>(index)).toUtf8().data();
}

void QFileSystemModel_DestroyQFileSystemModel(QtObjectPtr ptr){
	static_cast<QFileSystemModel*>(ptr)->~QFileSystemModel();
}

