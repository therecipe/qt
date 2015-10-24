#include "qfiledialog.h"
#include <QUrl>
#include <QModelIndex>
#include <QAbstractProxyModel>
#include <QAbstractItemDelegate>
#include <QStringList>
#include <QVariant>
#include <QObject>
#include <QFile>
#include <QDir>
#include <QFileIconProvider>
#include <QWidget>
#include <QByteArray>
#include <QString>
#include <QFileDialog>
#include "_cgo_export.h"

class MyQFileDialog: public QFileDialog {
public:
void Signal_CurrentChanged(const QString & path){callbackQFileDialogCurrentChanged(this->objectName().toUtf8().data(), path.toUtf8().data());};
void Signal_CurrentUrlChanged(const QUrl & url){callbackQFileDialogCurrentUrlChanged(this->objectName().toUtf8().data(), url.toString().toUtf8().data());};
void Signal_DirectoryEntered(const QString & directory){callbackQFileDialogDirectoryEntered(this->objectName().toUtf8().data(), directory.toUtf8().data());};
void Signal_DirectoryUrlEntered(const QUrl & directory){callbackQFileDialogDirectoryUrlEntered(this->objectName().toUtf8().data(), directory.toString().toUtf8().data());};
void Signal_FileSelected(const QString & file){callbackQFileDialogFileSelected(this->objectName().toUtf8().data(), file.toUtf8().data());};
void Signal_FilesSelected(const QStringList & selected){callbackQFileDialogFilesSelected(this->objectName().toUtf8().data(), selected.join("|").toUtf8().data());};
void Signal_FilterSelected(const QString & filter){callbackQFileDialogFilterSelected(this->objectName().toUtf8().data(), filter.toUtf8().data());};
void Signal_UrlSelected(const QUrl & url){callbackQFileDialogUrlSelected(this->objectName().toUtf8().data(), url.toString().toUtf8().data());};
};

QtObjectPtr QFileDialog_NewQFileDialog(QtObjectPtr parent, int flags){
	return new QFileDialog(static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(flags));
}

int QFileDialog_AcceptMode(QtObjectPtr ptr){
	return static_cast<QFileDialog*>(ptr)->acceptMode();
}

int QFileDialog_ConfirmOverwrite(QtObjectPtr ptr){
	return static_cast<QFileDialog*>(ptr)->confirmOverwrite();
}

char* QFileDialog_DefaultSuffix(QtObjectPtr ptr){
	return static_cast<QFileDialog*>(ptr)->defaultSuffix().toUtf8().data();
}

int QFileDialog_FileMode(QtObjectPtr ptr){
	return static_cast<QFileDialog*>(ptr)->fileMode();
}

int QFileDialog_IsNameFilterDetailsVisible(QtObjectPtr ptr){
	return static_cast<QFileDialog*>(ptr)->isNameFilterDetailsVisible();
}

int QFileDialog_IsReadOnly(QtObjectPtr ptr){
	return static_cast<QFileDialog*>(ptr)->isReadOnly();
}

int QFileDialog_Options(QtObjectPtr ptr){
	return static_cast<QFileDialog*>(ptr)->options();
}

int QFileDialog_ResolveSymlinks(QtObjectPtr ptr){
	return static_cast<QFileDialog*>(ptr)->resolveSymlinks();
}

void QFileDialog_SetAcceptMode(QtObjectPtr ptr, int mode){
	static_cast<QFileDialog*>(ptr)->setAcceptMode(static_cast<QFileDialog::AcceptMode>(mode));
}

void QFileDialog_SetConfirmOverwrite(QtObjectPtr ptr, int enabled){
	static_cast<QFileDialog*>(ptr)->setConfirmOverwrite(enabled != 0);
}

void QFileDialog_SetDefaultSuffix(QtObjectPtr ptr, char* suffix){
	static_cast<QFileDialog*>(ptr)->setDefaultSuffix(QString(suffix));
}

void QFileDialog_SetFileMode(QtObjectPtr ptr, int mode){
	static_cast<QFileDialog*>(ptr)->setFileMode(static_cast<QFileDialog::FileMode>(mode));
}

void QFileDialog_SetNameFilterDetailsVisible(QtObjectPtr ptr, int enabled){
	static_cast<QFileDialog*>(ptr)->setNameFilterDetailsVisible(enabled != 0);
}

void QFileDialog_SetOptions(QtObjectPtr ptr, int options){
	static_cast<QFileDialog*>(ptr)->setOptions(static_cast<QFileDialog::Option>(options));
}

void QFileDialog_SetReadOnly(QtObjectPtr ptr, int enabled){
	static_cast<QFileDialog*>(ptr)->setReadOnly(enabled != 0);
}

void QFileDialog_SetResolveSymlinks(QtObjectPtr ptr, int enabled){
	static_cast<QFileDialog*>(ptr)->setResolveSymlinks(enabled != 0);
}

void QFileDialog_SetViewMode(QtObjectPtr ptr, int mode){
	static_cast<QFileDialog*>(ptr)->setViewMode(static_cast<QFileDialog::ViewMode>(mode));
}

int QFileDialog_ViewMode(QtObjectPtr ptr){
	return static_cast<QFileDialog*>(ptr)->viewMode();
}

QtObjectPtr QFileDialog_NewQFileDialog2(QtObjectPtr parent, char* caption, char* directory, char* filter){
	return new QFileDialog(static_cast<QWidget*>(parent), QString(caption), QString(directory), QString(filter));
}

void QFileDialog_ConnectCurrentChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QFileDialog*>(ptr), static_cast<void (QFileDialog::*)(const QString &)>(&QFileDialog::currentChanged), static_cast<MyQFileDialog*>(ptr), static_cast<void (MyQFileDialog::*)(const QString &)>(&MyQFileDialog::Signal_CurrentChanged));;
}

void QFileDialog_DisconnectCurrentChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QFileDialog*>(ptr), static_cast<void (QFileDialog::*)(const QString &)>(&QFileDialog::currentChanged), static_cast<MyQFileDialog*>(ptr), static_cast<void (MyQFileDialog::*)(const QString &)>(&MyQFileDialog::Signal_CurrentChanged));;
}

void QFileDialog_ConnectCurrentUrlChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QFileDialog*>(ptr), static_cast<void (QFileDialog::*)(const QUrl &)>(&QFileDialog::currentUrlChanged), static_cast<MyQFileDialog*>(ptr), static_cast<void (MyQFileDialog::*)(const QUrl &)>(&MyQFileDialog::Signal_CurrentUrlChanged));;
}

void QFileDialog_DisconnectCurrentUrlChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QFileDialog*>(ptr), static_cast<void (QFileDialog::*)(const QUrl &)>(&QFileDialog::currentUrlChanged), static_cast<MyQFileDialog*>(ptr), static_cast<void (MyQFileDialog::*)(const QUrl &)>(&MyQFileDialog::Signal_CurrentUrlChanged));;
}

void QFileDialog_ConnectDirectoryEntered(QtObjectPtr ptr){
	QObject::connect(static_cast<QFileDialog*>(ptr), static_cast<void (QFileDialog::*)(const QString &)>(&QFileDialog::directoryEntered), static_cast<MyQFileDialog*>(ptr), static_cast<void (MyQFileDialog::*)(const QString &)>(&MyQFileDialog::Signal_DirectoryEntered));;
}

void QFileDialog_DisconnectDirectoryEntered(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QFileDialog*>(ptr), static_cast<void (QFileDialog::*)(const QString &)>(&QFileDialog::directoryEntered), static_cast<MyQFileDialog*>(ptr), static_cast<void (MyQFileDialog::*)(const QString &)>(&MyQFileDialog::Signal_DirectoryEntered));;
}

char* QFileDialog_DirectoryUrl(QtObjectPtr ptr){
	return static_cast<QFileDialog*>(ptr)->directoryUrl().toString().toUtf8().data();
}

void QFileDialog_ConnectDirectoryUrlEntered(QtObjectPtr ptr){
	QObject::connect(static_cast<QFileDialog*>(ptr), static_cast<void (QFileDialog::*)(const QUrl &)>(&QFileDialog::directoryUrlEntered), static_cast<MyQFileDialog*>(ptr), static_cast<void (MyQFileDialog::*)(const QUrl &)>(&MyQFileDialog::Signal_DirectoryUrlEntered));;
}

void QFileDialog_DisconnectDirectoryUrlEntered(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QFileDialog*>(ptr), static_cast<void (QFileDialog::*)(const QUrl &)>(&QFileDialog::directoryUrlEntered), static_cast<MyQFileDialog*>(ptr), static_cast<void (MyQFileDialog::*)(const QUrl &)>(&MyQFileDialog::Signal_DirectoryUrlEntered));;
}

void QFileDialog_ConnectFileSelected(QtObjectPtr ptr){
	QObject::connect(static_cast<QFileDialog*>(ptr), static_cast<void (QFileDialog::*)(const QString &)>(&QFileDialog::fileSelected), static_cast<MyQFileDialog*>(ptr), static_cast<void (MyQFileDialog::*)(const QString &)>(&MyQFileDialog::Signal_FileSelected));;
}

void QFileDialog_DisconnectFileSelected(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QFileDialog*>(ptr), static_cast<void (QFileDialog::*)(const QString &)>(&QFileDialog::fileSelected), static_cast<MyQFileDialog*>(ptr), static_cast<void (MyQFileDialog::*)(const QString &)>(&MyQFileDialog::Signal_FileSelected));;
}

void QFileDialog_ConnectFilesSelected(QtObjectPtr ptr){
	QObject::connect(static_cast<QFileDialog*>(ptr), static_cast<void (QFileDialog::*)(const QStringList &)>(&QFileDialog::filesSelected), static_cast<MyQFileDialog*>(ptr), static_cast<void (MyQFileDialog::*)(const QStringList &)>(&MyQFileDialog::Signal_FilesSelected));;
}

void QFileDialog_DisconnectFilesSelected(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QFileDialog*>(ptr), static_cast<void (QFileDialog::*)(const QStringList &)>(&QFileDialog::filesSelected), static_cast<MyQFileDialog*>(ptr), static_cast<void (MyQFileDialog::*)(const QStringList &)>(&MyQFileDialog::Signal_FilesSelected));;
}

int QFileDialog_Filter(QtObjectPtr ptr){
	return static_cast<QFileDialog*>(ptr)->filter();
}

void QFileDialog_ConnectFilterSelected(QtObjectPtr ptr){
	QObject::connect(static_cast<QFileDialog*>(ptr), static_cast<void (QFileDialog::*)(const QString &)>(&QFileDialog::filterSelected), static_cast<MyQFileDialog*>(ptr), static_cast<void (MyQFileDialog::*)(const QString &)>(&MyQFileDialog::Signal_FilterSelected));;
}

void QFileDialog_DisconnectFilterSelected(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QFileDialog*>(ptr), static_cast<void (QFileDialog::*)(const QString &)>(&QFileDialog::filterSelected), static_cast<MyQFileDialog*>(ptr), static_cast<void (MyQFileDialog::*)(const QString &)>(&MyQFileDialog::Signal_FilterSelected));;
}

char* QFileDialog_QFileDialog_GetExistingDirectory(QtObjectPtr parent, char* caption, char* dir, int options){
	return QFileDialog::getExistingDirectory(static_cast<QWidget*>(parent), QString(caption), QString(dir), static_cast<QFileDialog::Option>(options)).toUtf8().data();
}

char* QFileDialog_QFileDialog_GetExistingDirectoryUrl(QtObjectPtr parent, char* caption, char* dir, int options, char* supportedSchemes){
	return QFileDialog::getExistingDirectoryUrl(static_cast<QWidget*>(parent), QString(caption), QUrl(QString(dir)), static_cast<QFileDialog::Option>(options), QString(supportedSchemes).split("|", QString::SkipEmptyParts)).toString().toUtf8().data();
}

char* QFileDialog_QFileDialog_GetOpenFileName(QtObjectPtr parent, char* caption, char* dir, char* filter, char* selectedFilter, int options){
	return QFileDialog::getOpenFileName(static_cast<QWidget*>(parent), QString(caption), QString(dir), QString(filter), new QString(selectedFilter), static_cast<QFileDialog::Option>(options)).toUtf8().data();
}

char* QFileDialog_QFileDialog_GetOpenFileNames(QtObjectPtr parent, char* caption, char* dir, char* filter, char* selectedFilter, int options){
	return QFileDialog::getOpenFileNames(static_cast<QWidget*>(parent), QString(caption), QString(dir), QString(filter), new QString(selectedFilter), static_cast<QFileDialog::Option>(options)).join("|").toUtf8().data();
}

char* QFileDialog_QFileDialog_GetOpenFileUrl(QtObjectPtr parent, char* caption, char* dir, char* filter, char* selectedFilter, int options, char* supportedSchemes){
	return QFileDialog::getOpenFileUrl(static_cast<QWidget*>(parent), QString(caption), QUrl(QString(dir)), QString(filter), new QString(selectedFilter), static_cast<QFileDialog::Option>(options), QString(supportedSchemes).split("|", QString::SkipEmptyParts)).toString().toUtf8().data();
}

char* QFileDialog_QFileDialog_GetSaveFileName(QtObjectPtr parent, char* caption, char* dir, char* filter, char* selectedFilter, int options){
	return QFileDialog::getSaveFileName(static_cast<QWidget*>(parent), QString(caption), QString(dir), QString(filter), new QString(selectedFilter), static_cast<QFileDialog::Option>(options)).toUtf8().data();
}

char* QFileDialog_QFileDialog_GetSaveFileUrl(QtObjectPtr parent, char* caption, char* dir, char* filter, char* selectedFilter, int options, char* supportedSchemes){
	return QFileDialog::getSaveFileUrl(static_cast<QWidget*>(parent), QString(caption), QUrl(QString(dir)), QString(filter), new QString(selectedFilter), static_cast<QFileDialog::Option>(options), QString(supportedSchemes).split("|", QString::SkipEmptyParts)).toString().toUtf8().data();
}

char* QFileDialog_History(QtObjectPtr ptr){
	return static_cast<QFileDialog*>(ptr)->history().join("|").toUtf8().data();
}

QtObjectPtr QFileDialog_IconProvider(QtObjectPtr ptr){
	return static_cast<QFileDialog*>(ptr)->iconProvider();
}

QtObjectPtr QFileDialog_ItemDelegate(QtObjectPtr ptr){
	return static_cast<QFileDialog*>(ptr)->itemDelegate();
}

char* QFileDialog_LabelText(QtObjectPtr ptr, int label){
	return static_cast<QFileDialog*>(ptr)->labelText(static_cast<QFileDialog::DialogLabel>(label)).toUtf8().data();
}

char* QFileDialog_MimeTypeFilters(QtObjectPtr ptr){
	return static_cast<QFileDialog*>(ptr)->mimeTypeFilters().join("|").toUtf8().data();
}

char* QFileDialog_NameFilters(QtObjectPtr ptr){
	return static_cast<QFileDialog*>(ptr)->nameFilters().join("|").toUtf8().data();
}

void QFileDialog_Open(QtObjectPtr ptr, QtObjectPtr receiver, char* member){
	static_cast<QFileDialog*>(ptr)->open(static_cast<QObject*>(receiver), const_cast<const char*>(member));
}

QtObjectPtr QFileDialog_ProxyModel(QtObjectPtr ptr){
	return static_cast<QFileDialog*>(ptr)->proxyModel();
}

int QFileDialog_RestoreState(QtObjectPtr ptr, QtObjectPtr state){
	return static_cast<QFileDialog*>(ptr)->restoreState(*static_cast<QByteArray*>(state));
}

void QFileDialog_SelectFile(QtObjectPtr ptr, char* filename){
	static_cast<QFileDialog*>(ptr)->selectFile(QString(filename));
}

void QFileDialog_SelectMimeTypeFilter(QtObjectPtr ptr, char* filter){
	static_cast<QFileDialog*>(ptr)->selectMimeTypeFilter(QString(filter));
}

void QFileDialog_SelectNameFilter(QtObjectPtr ptr, char* filter){
	static_cast<QFileDialog*>(ptr)->selectNameFilter(QString(filter));
}

void QFileDialog_SelectUrl(QtObjectPtr ptr, char* url){
	static_cast<QFileDialog*>(ptr)->selectUrl(QUrl(QString(url)));
}

char* QFileDialog_SelectedFiles(QtObjectPtr ptr){
	return static_cast<QFileDialog*>(ptr)->selectedFiles().join("|").toUtf8().data();
}

char* QFileDialog_SelectedNameFilter(QtObjectPtr ptr){
	return static_cast<QFileDialog*>(ptr)->selectedNameFilter().toUtf8().data();
}

void QFileDialog_SetDirectory2(QtObjectPtr ptr, QtObjectPtr directory){
	static_cast<QFileDialog*>(ptr)->setDirectory(*static_cast<QDir*>(directory));
}

void QFileDialog_SetDirectory(QtObjectPtr ptr, char* directory){
	static_cast<QFileDialog*>(ptr)->setDirectory(QString(directory));
}

void QFileDialog_SetDirectoryUrl(QtObjectPtr ptr, char* directory){
	static_cast<QFileDialog*>(ptr)->setDirectoryUrl(QUrl(QString(directory)));
}

void QFileDialog_SetFilter(QtObjectPtr ptr, int filters){
	static_cast<QFileDialog*>(ptr)->setFilter(static_cast<QDir::Filter>(filters));
}

void QFileDialog_SetHistory(QtObjectPtr ptr, char* paths){
	static_cast<QFileDialog*>(ptr)->setHistory(QString(paths).split("|", QString::SkipEmptyParts));
}

void QFileDialog_SetIconProvider(QtObjectPtr ptr, QtObjectPtr provider){
	static_cast<QFileDialog*>(ptr)->setIconProvider(static_cast<QFileIconProvider*>(provider));
}

void QFileDialog_SetItemDelegate(QtObjectPtr ptr, QtObjectPtr delegate){
	static_cast<QFileDialog*>(ptr)->setItemDelegate(static_cast<QAbstractItemDelegate*>(delegate));
}

void QFileDialog_SetLabelText(QtObjectPtr ptr, int label, char* text){
	static_cast<QFileDialog*>(ptr)->setLabelText(static_cast<QFileDialog::DialogLabel>(label), QString(text));
}

void QFileDialog_SetMimeTypeFilters(QtObjectPtr ptr, char* filters){
	static_cast<QFileDialog*>(ptr)->setMimeTypeFilters(QString(filters).split("|", QString::SkipEmptyParts));
}

void QFileDialog_SetNameFilter(QtObjectPtr ptr, char* filter){
	static_cast<QFileDialog*>(ptr)->setNameFilter(QString(filter));
}

void QFileDialog_SetNameFilters(QtObjectPtr ptr, char* filters){
	static_cast<QFileDialog*>(ptr)->setNameFilters(QString(filters).split("|", QString::SkipEmptyParts));
}

void QFileDialog_SetOption(QtObjectPtr ptr, int option, int on){
	static_cast<QFileDialog*>(ptr)->setOption(static_cast<QFileDialog::Option>(option), on != 0);
}

void QFileDialog_SetProxyModel(QtObjectPtr ptr, QtObjectPtr proxyModel){
	static_cast<QFileDialog*>(ptr)->setProxyModel(static_cast<QAbstractProxyModel*>(proxyModel));
}

void QFileDialog_SetVisible(QtObjectPtr ptr, int visible){
	static_cast<QFileDialog*>(ptr)->setVisible(visible != 0);
}

int QFileDialog_TestOption(QtObjectPtr ptr, int option){
	return static_cast<QFileDialog*>(ptr)->testOption(static_cast<QFileDialog::Option>(option));
}

void QFileDialog_ConnectUrlSelected(QtObjectPtr ptr){
	QObject::connect(static_cast<QFileDialog*>(ptr), static_cast<void (QFileDialog::*)(const QUrl &)>(&QFileDialog::urlSelected), static_cast<MyQFileDialog*>(ptr), static_cast<void (MyQFileDialog::*)(const QUrl &)>(&MyQFileDialog::Signal_UrlSelected));;
}

void QFileDialog_DisconnectUrlSelected(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QFileDialog*>(ptr), static_cast<void (QFileDialog::*)(const QUrl &)>(&QFileDialog::urlSelected), static_cast<MyQFileDialog*>(ptr), static_cast<void (MyQFileDialog::*)(const QUrl &)>(&MyQFileDialog::Signal_UrlSelected));;
}

void QFileDialog_DestroyQFileDialog(QtObjectPtr ptr){
	static_cast<QFileDialog*>(ptr)->~QFileDialog();
}

