#include "qfilesystemmodel.h"
#include <QDir>
#include <QDate>
#include <QMimeData>
#include <QDateTime>
#include <QFile>
#include <QObject>
#include <QFileIconProvider>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
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

int QFileSystemModel_IsReadOnly(void* ptr){
	return static_cast<QFileSystemModel*>(ptr)->isReadOnly();
}

int QFileSystemModel_NameFilterDisables(void* ptr){
	return static_cast<QFileSystemModel*>(ptr)->nameFilterDisables();
}

int QFileSystemModel_ResolveSymlinks(void* ptr){
	return static_cast<QFileSystemModel*>(ptr)->resolveSymlinks();
}

int QFileSystemModel_Rmdir(void* ptr, void* index){
	return static_cast<QFileSystemModel*>(ptr)->rmdir(*static_cast<QModelIndex*>(index));
}

void QFileSystemModel_SetNameFilterDisables(void* ptr, int enable){
	static_cast<QFileSystemModel*>(ptr)->setNameFilterDisables(enable != 0);
}

void QFileSystemModel_SetReadOnly(void* ptr, int enable){
	static_cast<QFileSystemModel*>(ptr)->setReadOnly(enable != 0);
}

void QFileSystemModel_SetResolveSymlinks(void* ptr, int enable){
	static_cast<QFileSystemModel*>(ptr)->setResolveSymlinks(enable != 0);
}

void* QFileSystemModel_NewQFileSystemModel(void* parent){
	return new QFileSystemModel(static_cast<QObject*>(parent));
}

int QFileSystemModel_CanFetchMore(void* ptr, void* parent){
	return static_cast<QFileSystemModel*>(ptr)->canFetchMore(*static_cast<QModelIndex*>(parent));
}

int QFileSystemModel_ColumnCount(void* ptr, void* parent){
	return static_cast<QFileSystemModel*>(ptr)->columnCount(*static_cast<QModelIndex*>(parent));
}

void* QFileSystemModel_Data(void* ptr, void* index, int role){
	return new QVariant(static_cast<QFileSystemModel*>(ptr)->data(*static_cast<QModelIndex*>(index), role));
}

void QFileSystemModel_ConnectDirectoryLoaded(void* ptr){
	QObject::connect(static_cast<QFileSystemModel*>(ptr), static_cast<void (QFileSystemModel::*)(const QString &)>(&QFileSystemModel::directoryLoaded), static_cast<MyQFileSystemModel*>(ptr), static_cast<void (MyQFileSystemModel::*)(const QString &)>(&MyQFileSystemModel::Signal_DirectoryLoaded));;
}

void QFileSystemModel_DisconnectDirectoryLoaded(void* ptr){
	QObject::disconnect(static_cast<QFileSystemModel*>(ptr), static_cast<void (QFileSystemModel::*)(const QString &)>(&QFileSystemModel::directoryLoaded), static_cast<MyQFileSystemModel*>(ptr), static_cast<void (MyQFileSystemModel::*)(const QString &)>(&MyQFileSystemModel::Signal_DirectoryLoaded));;
}

int QFileSystemModel_DropMimeData(void* ptr, void* data, int action, int row, int column, void* parent){
	return static_cast<QFileSystemModel*>(ptr)->dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

void QFileSystemModel_FetchMore(void* ptr, void* parent){
	static_cast<QFileSystemModel*>(ptr)->fetchMore(*static_cast<QModelIndex*>(parent));
}

char* QFileSystemModel_FileName(void* ptr, void* index){
	return static_cast<QFileSystemModel*>(ptr)->fileName(*static_cast<QModelIndex*>(index)).toUtf8().data();
}

char* QFileSystemModel_FilePath(void* ptr, void* index){
	return static_cast<QFileSystemModel*>(ptr)->filePath(*static_cast<QModelIndex*>(index)).toUtf8().data();
}

void QFileSystemModel_ConnectFileRenamed(void* ptr){
	QObject::connect(static_cast<QFileSystemModel*>(ptr), static_cast<void (QFileSystemModel::*)(const QString &, const QString &, const QString &)>(&QFileSystemModel::fileRenamed), static_cast<MyQFileSystemModel*>(ptr), static_cast<void (MyQFileSystemModel::*)(const QString &, const QString &, const QString &)>(&MyQFileSystemModel::Signal_FileRenamed));;
}

void QFileSystemModel_DisconnectFileRenamed(void* ptr){
	QObject::disconnect(static_cast<QFileSystemModel*>(ptr), static_cast<void (QFileSystemModel::*)(const QString &, const QString &, const QString &)>(&QFileSystemModel::fileRenamed), static_cast<MyQFileSystemModel*>(ptr), static_cast<void (MyQFileSystemModel::*)(const QString &, const QString &, const QString &)>(&MyQFileSystemModel::Signal_FileRenamed));;
}

int QFileSystemModel_Filter(void* ptr){
	return static_cast<QFileSystemModel*>(ptr)->filter();
}

int QFileSystemModel_Flags(void* ptr, void* index){
	return static_cast<QFileSystemModel*>(ptr)->flags(*static_cast<QModelIndex*>(index));
}

int QFileSystemModel_HasChildren(void* ptr, void* parent){
	return static_cast<QFileSystemModel*>(ptr)->hasChildren(*static_cast<QModelIndex*>(parent));
}

void* QFileSystemModel_HeaderData(void* ptr, int section, int orientation, int role){
	return new QVariant(static_cast<QFileSystemModel*>(ptr)->headerData(section, static_cast<Qt::Orientation>(orientation), role));
}

void* QFileSystemModel_IconProvider(void* ptr){
	return static_cast<QFileSystemModel*>(ptr)->iconProvider();
}

void* QFileSystemModel_Index2(void* ptr, char* path, int column){
	return static_cast<QFileSystemModel*>(ptr)->index(QString(path), column).internalPointer();
}

void* QFileSystemModel_Index(void* ptr, int row, int column, void* parent){
	return static_cast<QFileSystemModel*>(ptr)->index(row, column, *static_cast<QModelIndex*>(parent)).internalPointer();
}

int QFileSystemModel_IsDir(void* ptr, void* index){
	return static_cast<QFileSystemModel*>(ptr)->isDir(*static_cast<QModelIndex*>(index));
}

void* QFileSystemModel_LastModified(void* ptr, void* index){
	return new QDateTime(static_cast<QFileSystemModel*>(ptr)->lastModified(*static_cast<QModelIndex*>(index)));
}

char* QFileSystemModel_MimeTypes(void* ptr){
	return static_cast<QFileSystemModel*>(ptr)->mimeTypes().join("|").toUtf8().data();
}

void* QFileSystemModel_Mkdir(void* ptr, void* parent, char* name){
	return static_cast<QFileSystemModel*>(ptr)->mkdir(*static_cast<QModelIndex*>(parent), QString(name)).internalPointer();
}

void* QFileSystemModel_MyComputer(void* ptr, int role){
	return new QVariant(static_cast<QFileSystemModel*>(ptr)->myComputer(role));
}

char* QFileSystemModel_NameFilters(void* ptr){
	return static_cast<QFileSystemModel*>(ptr)->nameFilters().join("|").toUtf8().data();
}

void* QFileSystemModel_Parent(void* ptr, void* index){
	return static_cast<QFileSystemModel*>(ptr)->parent(*static_cast<QModelIndex*>(index)).internalPointer();
}

void* QFileSystemModel_RootDirectory(void* ptr){
	return new QDir(static_cast<QFileSystemModel*>(ptr)->rootDirectory());
}

char* QFileSystemModel_RootPath(void* ptr){
	return static_cast<QFileSystemModel*>(ptr)->rootPath().toUtf8().data();
}

void QFileSystemModel_ConnectRootPathChanged(void* ptr){
	QObject::connect(static_cast<QFileSystemModel*>(ptr), static_cast<void (QFileSystemModel::*)(const QString &)>(&QFileSystemModel::rootPathChanged), static_cast<MyQFileSystemModel*>(ptr), static_cast<void (MyQFileSystemModel::*)(const QString &)>(&MyQFileSystemModel::Signal_RootPathChanged));;
}

void QFileSystemModel_DisconnectRootPathChanged(void* ptr){
	QObject::disconnect(static_cast<QFileSystemModel*>(ptr), static_cast<void (QFileSystemModel::*)(const QString &)>(&QFileSystemModel::rootPathChanged), static_cast<MyQFileSystemModel*>(ptr), static_cast<void (MyQFileSystemModel::*)(const QString &)>(&MyQFileSystemModel::Signal_RootPathChanged));;
}

int QFileSystemModel_RowCount(void* ptr, void* parent){
	return static_cast<QFileSystemModel*>(ptr)->rowCount(*static_cast<QModelIndex*>(parent));
}

int QFileSystemModel_SetData(void* ptr, void* idx, void* value, int role){
	return static_cast<QFileSystemModel*>(ptr)->setData(*static_cast<QModelIndex*>(idx), *static_cast<QVariant*>(value), role);
}

void QFileSystemModel_SetFilter(void* ptr, int filters){
	static_cast<QFileSystemModel*>(ptr)->setFilter(static_cast<QDir::Filter>(filters));
}

void QFileSystemModel_SetIconProvider(void* ptr, void* provider){
	static_cast<QFileSystemModel*>(ptr)->setIconProvider(static_cast<QFileIconProvider*>(provider));
}

void QFileSystemModel_SetNameFilters(void* ptr, char* filters){
	static_cast<QFileSystemModel*>(ptr)->setNameFilters(QString(filters).split("|", QString::SkipEmptyParts));
}

void* QFileSystemModel_SetRootPath(void* ptr, char* newPath){
	return static_cast<QFileSystemModel*>(ptr)->setRootPath(QString(newPath)).internalPointer();
}

void QFileSystemModel_Sort(void* ptr, int column, int order){
	static_cast<QFileSystemModel*>(ptr)->sort(column, static_cast<Qt::SortOrder>(order));
}

int QFileSystemModel_SupportedDropActions(void* ptr){
	return static_cast<QFileSystemModel*>(ptr)->supportedDropActions();
}

char* QFileSystemModel_Type(void* ptr, void* index){
	return static_cast<QFileSystemModel*>(ptr)->type(*static_cast<QModelIndex*>(index)).toUtf8().data();
}

void QFileSystemModel_DestroyQFileSystemModel(void* ptr){
	static_cast<QFileSystemModel*>(ptr)->~QFileSystemModel();
}

