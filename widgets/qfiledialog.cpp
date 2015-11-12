#include "qfiledialog.h"
#include <QWidget>
#include <QByteArray>
#include <QObject>
#include <QDir>
#include <QModelIndex>
#include <QAbstractItemDelegate>
#include <QUrl>
#include <QFileIconProvider>
#include <QStringList>
#include <QFile>
#include <QAbstractProxyModel>
#include <QString>
#include <QVariant>
#include <QFileDialog>
#include "_cgo_export.h"

class MyQFileDialog: public QFileDialog {
public:
void Signal_CurrentChanged(const QString & path){callbackQFileDialogCurrentChanged(this->objectName().toUtf8().data(), path.toUtf8().data());};
void Signal_DirectoryEntered(const QString & directory){callbackQFileDialogDirectoryEntered(this->objectName().toUtf8().data(), directory.toUtf8().data());};
void Signal_FileSelected(const QString & file){callbackQFileDialogFileSelected(this->objectName().toUtf8().data(), file.toUtf8().data());};
void Signal_FilesSelected(const QStringList & selected){callbackQFileDialogFilesSelected(this->objectName().toUtf8().data(), selected.join("|").toUtf8().data());};
void Signal_FilterSelected(const QString & filter){callbackQFileDialogFilterSelected(this->objectName().toUtf8().data(), filter.toUtf8().data());};
};

void* QFileDialog_NewQFileDialog(void* parent, int flags){
	return new QFileDialog(static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(flags));
}

int QFileDialog_AcceptMode(void* ptr){
	return static_cast<QFileDialog*>(ptr)->acceptMode();
}

int QFileDialog_ConfirmOverwrite(void* ptr){
	return static_cast<QFileDialog*>(ptr)->confirmOverwrite();
}

char* QFileDialog_DefaultSuffix(void* ptr){
	return static_cast<QFileDialog*>(ptr)->defaultSuffix().toUtf8().data();
}

int QFileDialog_FileMode(void* ptr){
	return static_cast<QFileDialog*>(ptr)->fileMode();
}

int QFileDialog_IsNameFilterDetailsVisible(void* ptr){
	return static_cast<QFileDialog*>(ptr)->isNameFilterDetailsVisible();
}

int QFileDialog_IsReadOnly(void* ptr){
	return static_cast<QFileDialog*>(ptr)->isReadOnly();
}

int QFileDialog_Options(void* ptr){
	return static_cast<QFileDialog*>(ptr)->options();
}

int QFileDialog_ResolveSymlinks(void* ptr){
	return static_cast<QFileDialog*>(ptr)->resolveSymlinks();
}

void QFileDialog_SetAcceptMode(void* ptr, int mode){
	static_cast<QFileDialog*>(ptr)->setAcceptMode(static_cast<QFileDialog::AcceptMode>(mode));
}

void QFileDialog_SetConfirmOverwrite(void* ptr, int enabled){
	static_cast<QFileDialog*>(ptr)->setConfirmOverwrite(enabled != 0);
}

void QFileDialog_SetDefaultSuffix(void* ptr, char* suffix){
	static_cast<QFileDialog*>(ptr)->setDefaultSuffix(QString(suffix));
}

void QFileDialog_SetFileMode(void* ptr, int mode){
	static_cast<QFileDialog*>(ptr)->setFileMode(static_cast<QFileDialog::FileMode>(mode));
}

void QFileDialog_SetNameFilterDetailsVisible(void* ptr, int enabled){
	static_cast<QFileDialog*>(ptr)->setNameFilterDetailsVisible(enabled != 0);
}

void QFileDialog_SetOptions(void* ptr, int options){
	static_cast<QFileDialog*>(ptr)->setOptions(static_cast<QFileDialog::Option>(options));
}

void QFileDialog_SetReadOnly(void* ptr, int enabled){
	static_cast<QFileDialog*>(ptr)->setReadOnly(enabled != 0);
}

void QFileDialog_SetResolveSymlinks(void* ptr, int enabled){
	static_cast<QFileDialog*>(ptr)->setResolveSymlinks(enabled != 0);
}

void QFileDialog_SetViewMode(void* ptr, int mode){
	static_cast<QFileDialog*>(ptr)->setViewMode(static_cast<QFileDialog::ViewMode>(mode));
}

int QFileDialog_ViewMode(void* ptr){
	return static_cast<QFileDialog*>(ptr)->viewMode();
}

void* QFileDialog_NewQFileDialog2(void* parent, char* caption, char* directory, char* filter){
	return new QFileDialog(static_cast<QWidget*>(parent), QString(caption), QString(directory), QString(filter));
}

void QFileDialog_ConnectCurrentChanged(void* ptr){
	QObject::connect(static_cast<QFileDialog*>(ptr), static_cast<void (QFileDialog::*)(const QString &)>(&QFileDialog::currentChanged), static_cast<MyQFileDialog*>(ptr), static_cast<void (MyQFileDialog::*)(const QString &)>(&MyQFileDialog::Signal_CurrentChanged));;
}

void QFileDialog_DisconnectCurrentChanged(void* ptr){
	QObject::disconnect(static_cast<QFileDialog*>(ptr), static_cast<void (QFileDialog::*)(const QString &)>(&QFileDialog::currentChanged), static_cast<MyQFileDialog*>(ptr), static_cast<void (MyQFileDialog::*)(const QString &)>(&MyQFileDialog::Signal_CurrentChanged));;
}

void* QFileDialog_Directory(void* ptr){
	return new QDir(static_cast<QFileDialog*>(ptr)->directory());
}

void QFileDialog_ConnectDirectoryEntered(void* ptr){
	QObject::connect(static_cast<QFileDialog*>(ptr), static_cast<void (QFileDialog::*)(const QString &)>(&QFileDialog::directoryEntered), static_cast<MyQFileDialog*>(ptr), static_cast<void (MyQFileDialog::*)(const QString &)>(&MyQFileDialog::Signal_DirectoryEntered));;
}

void QFileDialog_DisconnectDirectoryEntered(void* ptr){
	QObject::disconnect(static_cast<QFileDialog*>(ptr), static_cast<void (QFileDialog::*)(const QString &)>(&QFileDialog::directoryEntered), static_cast<MyQFileDialog*>(ptr), static_cast<void (MyQFileDialog::*)(const QString &)>(&MyQFileDialog::Signal_DirectoryEntered));;
}

void QFileDialog_ConnectFileSelected(void* ptr){
	QObject::connect(static_cast<QFileDialog*>(ptr), static_cast<void (QFileDialog::*)(const QString &)>(&QFileDialog::fileSelected), static_cast<MyQFileDialog*>(ptr), static_cast<void (MyQFileDialog::*)(const QString &)>(&MyQFileDialog::Signal_FileSelected));;
}

void QFileDialog_DisconnectFileSelected(void* ptr){
	QObject::disconnect(static_cast<QFileDialog*>(ptr), static_cast<void (QFileDialog::*)(const QString &)>(&QFileDialog::fileSelected), static_cast<MyQFileDialog*>(ptr), static_cast<void (MyQFileDialog::*)(const QString &)>(&MyQFileDialog::Signal_FileSelected));;
}

void QFileDialog_ConnectFilesSelected(void* ptr){
	QObject::connect(static_cast<QFileDialog*>(ptr), static_cast<void (QFileDialog::*)(const QStringList &)>(&QFileDialog::filesSelected), static_cast<MyQFileDialog*>(ptr), static_cast<void (MyQFileDialog::*)(const QStringList &)>(&MyQFileDialog::Signal_FilesSelected));;
}

void QFileDialog_DisconnectFilesSelected(void* ptr){
	QObject::disconnect(static_cast<QFileDialog*>(ptr), static_cast<void (QFileDialog::*)(const QStringList &)>(&QFileDialog::filesSelected), static_cast<MyQFileDialog*>(ptr), static_cast<void (MyQFileDialog::*)(const QStringList &)>(&MyQFileDialog::Signal_FilesSelected));;
}

int QFileDialog_Filter(void* ptr){
	return static_cast<QFileDialog*>(ptr)->filter();
}

void QFileDialog_ConnectFilterSelected(void* ptr){
	QObject::connect(static_cast<QFileDialog*>(ptr), static_cast<void (QFileDialog::*)(const QString &)>(&QFileDialog::filterSelected), static_cast<MyQFileDialog*>(ptr), static_cast<void (MyQFileDialog::*)(const QString &)>(&MyQFileDialog::Signal_FilterSelected));;
}

void QFileDialog_DisconnectFilterSelected(void* ptr){
	QObject::disconnect(static_cast<QFileDialog*>(ptr), static_cast<void (QFileDialog::*)(const QString &)>(&QFileDialog::filterSelected), static_cast<MyQFileDialog*>(ptr), static_cast<void (MyQFileDialog::*)(const QString &)>(&MyQFileDialog::Signal_FilterSelected));;
}

char* QFileDialog_QFileDialog_GetExistingDirectory(void* parent, char* caption, char* dir, int options){
	return QFileDialog::getExistingDirectory(static_cast<QWidget*>(parent), QString(caption), QString(dir), static_cast<QFileDialog::Option>(options)).toUtf8().data();
}

char* QFileDialog_QFileDialog_GetOpenFileName(void* parent, char* caption, char* dir, char* filter, char* selectedFilter, int options){
	return QFileDialog::getOpenFileName(static_cast<QWidget*>(parent), QString(caption), QString(dir), QString(filter), new QString(selectedFilter), static_cast<QFileDialog::Option>(options)).toUtf8().data();
}

char* QFileDialog_QFileDialog_GetOpenFileNames(void* parent, char* caption, char* dir, char* filter, char* selectedFilter, int options){
	return QFileDialog::getOpenFileNames(static_cast<QWidget*>(parent), QString(caption), QString(dir), QString(filter), new QString(selectedFilter), static_cast<QFileDialog::Option>(options)).join("|").toUtf8().data();
}

char* QFileDialog_QFileDialog_GetSaveFileName(void* parent, char* caption, char* dir, char* filter, char* selectedFilter, int options){
	return QFileDialog::getSaveFileName(static_cast<QWidget*>(parent), QString(caption), QString(dir), QString(filter), new QString(selectedFilter), static_cast<QFileDialog::Option>(options)).toUtf8().data();
}

char* QFileDialog_History(void* ptr){
	return static_cast<QFileDialog*>(ptr)->history().join("|").toUtf8().data();
}

void* QFileDialog_IconProvider(void* ptr){
	return static_cast<QFileDialog*>(ptr)->iconProvider();
}

void* QFileDialog_ItemDelegate(void* ptr){
	return static_cast<QFileDialog*>(ptr)->itemDelegate();
}

char* QFileDialog_LabelText(void* ptr, int label){
	return static_cast<QFileDialog*>(ptr)->labelText(static_cast<QFileDialog::DialogLabel>(label)).toUtf8().data();
}

char* QFileDialog_MimeTypeFilters(void* ptr){
	return static_cast<QFileDialog*>(ptr)->mimeTypeFilters().join("|").toUtf8().data();
}

char* QFileDialog_NameFilters(void* ptr){
	return static_cast<QFileDialog*>(ptr)->nameFilters().join("|").toUtf8().data();
}

void QFileDialog_Open(void* ptr, void* receiver, char* member){
	static_cast<QFileDialog*>(ptr)->open(static_cast<QObject*>(receiver), const_cast<const char*>(member));
}

void* QFileDialog_ProxyModel(void* ptr){
	return static_cast<QFileDialog*>(ptr)->proxyModel();
}

int QFileDialog_RestoreState(void* ptr, void* state){
	return static_cast<QFileDialog*>(ptr)->restoreState(*static_cast<QByteArray*>(state));
}

void* QFileDialog_SaveState(void* ptr){
	return new QByteArray(static_cast<QFileDialog*>(ptr)->saveState());
}

void QFileDialog_SelectFile(void* ptr, char* filename){
	static_cast<QFileDialog*>(ptr)->selectFile(QString(filename));
}

void QFileDialog_SelectMimeTypeFilter(void* ptr, char* filter){
	static_cast<QFileDialog*>(ptr)->selectMimeTypeFilter(QString(filter));
}

void QFileDialog_SelectNameFilter(void* ptr, char* filter){
	static_cast<QFileDialog*>(ptr)->selectNameFilter(QString(filter));
}

void QFileDialog_SelectUrl(void* ptr, void* url){
	static_cast<QFileDialog*>(ptr)->selectUrl(*static_cast<QUrl*>(url));
}

char* QFileDialog_SelectedFiles(void* ptr){
	return static_cast<QFileDialog*>(ptr)->selectedFiles().join("|").toUtf8().data();
}

char* QFileDialog_SelectedNameFilter(void* ptr){
	return static_cast<QFileDialog*>(ptr)->selectedNameFilter().toUtf8().data();
}

void QFileDialog_SetDirectory2(void* ptr, void* directory){
	static_cast<QFileDialog*>(ptr)->setDirectory(*static_cast<QDir*>(directory));
}

void QFileDialog_SetDirectory(void* ptr, char* directory){
	static_cast<QFileDialog*>(ptr)->setDirectory(QString(directory));
}

void QFileDialog_SetDirectoryUrl(void* ptr, void* directory){
	static_cast<QFileDialog*>(ptr)->setDirectoryUrl(*static_cast<QUrl*>(directory));
}

void QFileDialog_SetFilter(void* ptr, int filters){
	static_cast<QFileDialog*>(ptr)->setFilter(static_cast<QDir::Filter>(filters));
}

void QFileDialog_SetHistory(void* ptr, char* paths){
	static_cast<QFileDialog*>(ptr)->setHistory(QString(paths).split("|", QString::SkipEmptyParts));
}

void QFileDialog_SetIconProvider(void* ptr, void* provider){
	static_cast<QFileDialog*>(ptr)->setIconProvider(static_cast<QFileIconProvider*>(provider));
}

void QFileDialog_SetItemDelegate(void* ptr, void* delegate){
	static_cast<QFileDialog*>(ptr)->setItemDelegate(static_cast<QAbstractItemDelegate*>(delegate));
}

void QFileDialog_SetLabelText(void* ptr, int label, char* text){
	static_cast<QFileDialog*>(ptr)->setLabelText(static_cast<QFileDialog::DialogLabel>(label), QString(text));
}

void QFileDialog_SetMimeTypeFilters(void* ptr, char* filters){
	static_cast<QFileDialog*>(ptr)->setMimeTypeFilters(QString(filters).split("|", QString::SkipEmptyParts));
}

void QFileDialog_SetNameFilter(void* ptr, char* filter){
	static_cast<QFileDialog*>(ptr)->setNameFilter(QString(filter));
}

void QFileDialog_SetNameFilters(void* ptr, char* filters){
	static_cast<QFileDialog*>(ptr)->setNameFilters(QString(filters).split("|", QString::SkipEmptyParts));
}

void QFileDialog_SetOption(void* ptr, int option, int on){
	static_cast<QFileDialog*>(ptr)->setOption(static_cast<QFileDialog::Option>(option), on != 0);
}

void QFileDialog_SetProxyModel(void* ptr, void* proxyModel){
	static_cast<QFileDialog*>(ptr)->setProxyModel(static_cast<QAbstractProxyModel*>(proxyModel));
}

void QFileDialog_SetVisible(void* ptr, int visible){
	static_cast<QFileDialog*>(ptr)->setVisible(visible != 0);
}

int QFileDialog_TestOption(void* ptr, int option){
	return static_cast<QFileDialog*>(ptr)->testOption(static_cast<QFileDialog::Option>(option));
}

void QFileDialog_DestroyQFileDialog(void* ptr){
	static_cast<QFileDialog*>(ptr)->~QFileDialog();
}

