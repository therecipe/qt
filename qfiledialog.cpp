#include "qfiledialog.h"
#include <QFileDialog>
#include "cgoexport.h"

//MyQFileDialog
class MyQFileDialog: public QFileDialog {
Q_OBJECT
public:
void Signal_CurrentChanged() { callbackQFileDialog(this, QString("currentChanged").toUtf8().data()); };
void Signal_DirectoryEntered() { callbackQFileDialog(this, QString("directoryEntered").toUtf8().data()); };
void Signal_FileSelected() { callbackQFileDialog(this, QString("fileSelected").toUtf8().data()); };
void Signal_FilesSelected() { callbackQFileDialog(this, QString("filesSelected").toUtf8().data()); };
void Signal_FilterSelected() { callbackQFileDialog(this, QString("filterSelected").toUtf8().data()); };

signals:

};
#include "qfiledialog.moc"

//Public Types
int QFileDialog_AcceptOpen() { return QFileDialog::AcceptOpen; }
int QFileDialog_AcceptSave() { return QFileDialog::AcceptSave; }
int QFileDialog_LookIn() { return QFileDialog::LookIn; }
int QFileDialog_FileName() { return QFileDialog::FileName; }
int QFileDialog_FileType() { return QFileDialog::FileType; }
int QFileDialog_Accept() { return QFileDialog::Accept; }
int QFileDialog_Reject() { return QFileDialog::Reject; }
int QFileDialog_AnyFile() { return QFileDialog::AnyFile; }
int QFileDialog_ExistingFile() { return QFileDialog::ExistingFile; }
int QFileDialog_Directory() { return QFileDialog::Directory; }
int QFileDialog_ExistingFiles() { return QFileDialog::ExistingFiles; }
int QFileDialog_DirectoryOnly() { return QFileDialog::DirectoryOnly; }
int QFileDialog_ShowDirsOnly() { return QFileDialog::ShowDirsOnly; }
int QFileDialog_DontResolveSymlinks() { return QFileDialog::DontResolveSymlinks; }
int QFileDialog_DontConfirmOverwrite() { return QFileDialog::DontConfirmOverwrite; }
int QFileDialog_DontUseNativeDialog() { return QFileDialog::DontUseNativeDialog; }
int QFileDialog_ReadOnly() { return QFileDialog::ReadOnly; }
int QFileDialog_HideNameFilterDetails() { return QFileDialog::HideNameFilterDetails; }
int QFileDialog_DontUseSheet() { return QFileDialog::DontUseSheet; }
int QFileDialog_DontUseCustomDirectoryIcons() { return QFileDialog::DontUseCustomDirectoryIcons; }
int QFileDialog_Detail() { return QFileDialog::Detail; }
int QFileDialog_List() { return QFileDialog::List; }

//Public Functions
QtObjectPtr QFileDialog_New_QWidget_String_String_String(QtObjectPtr parent, char* caption, char* directory, char* filter)
{
	return new QFileDialog(((QWidget*)(parent)), QString(caption), QString(directory), QString(filter));
}

void QFileDialog_Destroy(QtObjectPtr ptr)
{
	((QFileDialog*)(ptr))->~QFileDialog();
}

int QFileDialog_AcceptMode(QtObjectPtr ptr)
{
	return ((QFileDialog*)(ptr))->acceptMode();
}

char* QFileDialog_DefaultSuffix(QtObjectPtr ptr)
{
	return ((QFileDialog*)(ptr))->defaultSuffix().toUtf8().data();
}

int QFileDialog_FileMode(QtObjectPtr ptr)
{
	return ((QFileDialog*)(ptr))->fileMode();
}

char* QFileDialog_History(QtObjectPtr ptr)
{
	return ((QFileDialog*)(ptr))->history().join("|").toUtf8().data();
}

char* QFileDialog_LabelText_DialogLabel(QtObjectPtr ptr, int label)
{
	return ((QFileDialog*)(ptr))->labelText(((QFileDialog::DialogLabel)(label))).toUtf8().data();
}

char* QFileDialog_MimeTypeFilters(QtObjectPtr ptr)
{
	return ((QFileDialog*)(ptr))->mimeTypeFilters().join("|").toUtf8().data();
}

char* QFileDialog_NameFilters(QtObjectPtr ptr)
{
	return ((QFileDialog*)(ptr))->nameFilters().join("|").toUtf8().data();
}

void QFileDialog_Open_QObject_String(QtObjectPtr ptr, QtObjectPtr receiver, char* member)
{
	((QFileDialog*)(ptr))->open(((QObject*)(receiver)), member);
}

int QFileDialog_Options(QtObjectPtr ptr)
{
	return ((QFileDialog*)(ptr))->options();
}

void QFileDialog_SelectFile_String(QtObjectPtr ptr, char* filename)
{
	((QFileDialog*)(ptr))->selectFile(QString(filename));
}

void QFileDialog_SelectMimeTypeFilter_String(QtObjectPtr ptr, char* filter)
{
	((QFileDialog*)(ptr))->selectMimeTypeFilter(QString(filter));
}

void QFileDialog_SelectNameFilter_String(QtObjectPtr ptr, char* filter)
{
	((QFileDialog*)(ptr))->selectNameFilter(QString(filter));
}

char* QFileDialog_SelectedFiles(QtObjectPtr ptr)
{
	return ((QFileDialog*)(ptr))->selectedFiles().join("|").toUtf8().data();
}

char* QFileDialog_SelectedNameFilter(QtObjectPtr ptr)
{
	return ((QFileDialog*)(ptr))->selectedNameFilter().toUtf8().data();
}

void QFileDialog_SetAcceptMode_AcceptMode(QtObjectPtr ptr, int mode)
{
	((QFileDialog*)(ptr))->setAcceptMode(((QFileDialog::AcceptMode)(mode)));
}

void QFileDialog_SetDefaultSuffix_String(QtObjectPtr ptr, char* suffix)
{
	((QFileDialog*)(ptr))->setDefaultSuffix(QString(suffix));
}

void QFileDialog_SetDirectory_String(QtObjectPtr ptr, char* directory)
{
	((QFileDialog*)(ptr))->setDirectory(QString(directory));
}

void QFileDialog_SetFileMode_FileMode(QtObjectPtr ptr, int mode)
{
	((QFileDialog*)(ptr))->setFileMode(((QFileDialog::FileMode)(mode)));
}

void QFileDialog_SetHistory_QStringList(QtObjectPtr ptr, char* paths)
{
	((QFileDialog*)(ptr))->setHistory(QString(paths).split("|", QString::SkipEmptyParts));
}

void QFileDialog_SetLabelText_DialogLabel_String(QtObjectPtr ptr, int label, char* text)
{
	((QFileDialog*)(ptr))->setLabelText(((QFileDialog::DialogLabel)(label)), QString(text));
}

void QFileDialog_SetMimeTypeFilters_QStringList(QtObjectPtr ptr, char* filters)
{
	((QFileDialog*)(ptr))->setMimeTypeFilters(QString(filters).split("|", QString::SkipEmptyParts));
}

void QFileDialog_SetNameFilter_String(QtObjectPtr ptr, char* filter)
{
	((QFileDialog*)(ptr))->setNameFilter(QString(filter));
}

void QFileDialog_SetNameFilters_QStringList(QtObjectPtr ptr, char* filters)
{
	((QFileDialog*)(ptr))->setNameFilters(QString(filters).split("|", QString::SkipEmptyParts));
}

void QFileDialog_SetOption_Option_Bool(QtObjectPtr ptr, int option, int on)
{
	((QFileDialog*)(ptr))->setOption(((QFileDialog::Option)(option)), on != 0);
}

void QFileDialog_SetOptions_Option(QtObjectPtr ptr, int options)
{
	((QFileDialog*)(ptr))->setOptions(((QFileDialog::Option)(options)));
}

void QFileDialog_SetViewMode_ViewMode(QtObjectPtr ptr, int mode)
{
	((QFileDialog*)(ptr))->setViewMode(((QFileDialog::ViewMode)(mode)));
}

int QFileDialog_TestOption_Option(QtObjectPtr ptr, int option)
{
	return ((QFileDialog*)(ptr))->testOption(((QFileDialog::Option)(option)));
}

int QFileDialog_ViewMode(QtObjectPtr ptr)
{
	return ((QFileDialog*)(ptr))->viewMode();
}

//Signals
void QFileDialog_ConnectSignalCurrentChanged(QtObjectPtr ptr)
{
	QObject::connect(((QFileDialog*)(ptr)), &QFileDialog::currentChanged, ((MyQFileDialog*)(ptr)), &MyQFileDialog::Signal_CurrentChanged, Qt::QueuedConnection);
}

void QFileDialog_DisconnectSignalCurrentChanged(QtObjectPtr ptr)
{
	QObject::disconnect(((QFileDialog*)(ptr)), &QFileDialog::currentChanged, ((MyQFileDialog*)(ptr)), &MyQFileDialog::Signal_CurrentChanged);
}

void QFileDialog_ConnectSignalDirectoryEntered(QtObjectPtr ptr)
{
	QObject::connect(((QFileDialog*)(ptr)), &QFileDialog::directoryEntered, ((MyQFileDialog*)(ptr)), &MyQFileDialog::Signal_DirectoryEntered, Qt::QueuedConnection);
}

void QFileDialog_DisconnectSignalDirectoryEntered(QtObjectPtr ptr)
{
	QObject::disconnect(((QFileDialog*)(ptr)), &QFileDialog::directoryEntered, ((MyQFileDialog*)(ptr)), &MyQFileDialog::Signal_DirectoryEntered);
}

void QFileDialog_ConnectSignalFileSelected(QtObjectPtr ptr)
{
	QObject::connect(((QFileDialog*)(ptr)), &QFileDialog::fileSelected, ((MyQFileDialog*)(ptr)), &MyQFileDialog::Signal_FileSelected, Qt::QueuedConnection);
}

void QFileDialog_DisconnectSignalFileSelected(QtObjectPtr ptr)
{
	QObject::disconnect(((QFileDialog*)(ptr)), &QFileDialog::fileSelected, ((MyQFileDialog*)(ptr)), &MyQFileDialog::Signal_FileSelected);
}

void QFileDialog_ConnectSignalFilesSelected(QtObjectPtr ptr)
{
	QObject::connect(((QFileDialog*)(ptr)), &QFileDialog::filesSelected, ((MyQFileDialog*)(ptr)), &MyQFileDialog::Signal_FilesSelected, Qt::QueuedConnection);
}

void QFileDialog_DisconnectSignalFilesSelected(QtObjectPtr ptr)
{
	QObject::disconnect(((QFileDialog*)(ptr)), &QFileDialog::filesSelected, ((MyQFileDialog*)(ptr)), &MyQFileDialog::Signal_FilesSelected);
}

void QFileDialog_ConnectSignalFilterSelected(QtObjectPtr ptr)
{
	QObject::connect(((QFileDialog*)(ptr)), &QFileDialog::filterSelected, ((MyQFileDialog*)(ptr)), &MyQFileDialog::Signal_FilterSelected, Qt::QueuedConnection);
}

void QFileDialog_DisconnectSignalFilterSelected(QtObjectPtr ptr)
{
	QObject::disconnect(((QFileDialog*)(ptr)), &QFileDialog::filterSelected, ((MyQFileDialog*)(ptr)), &MyQFileDialog::Signal_FilterSelected);
}

//Static Public Members
char* QFileDialog_GetExistingDirectory_QWidget_String_String_Option(QtObjectPtr parent, char* caption, char* dir, int options)
{
	return QFileDialog::getExistingDirectory(((QWidget*)(parent)), QString(caption), QString(dir), ((QFileDialog::Option)(options))).toUtf8().data();
}

char* QFileDialog_GetOpenFileName_QWidget_String_String_String_String_Option(QtObjectPtr parent, char* caption, char* dir, char* filter, char* selectedFilter, int options)
{
	return QFileDialog::getOpenFileName(((QWidget*)(parent)), QString(caption), QString(dir), QString(filter), new QString(selectedFilter), ((QFileDialog::Option)(options))).toUtf8().data();
}

char* QFileDialog_GetOpenFileNames_QWidget_String_String_String_String_Option(QtObjectPtr parent, char* caption, char* dir, char* filter, char* selectedFilter, int options)
{
	return QFileDialog::getOpenFileNames(((QWidget*)(parent)), QString(caption), QString(dir), QString(filter), new QString(selectedFilter), ((QFileDialog::Option)(options))).join("|").toUtf8().data();
}

char* QFileDialog_GetSaveFileName_QWidget_String_String_String_String_Option(QtObjectPtr parent, char* caption, char* dir, char* filter, char* selectedFilter, int options)
{
	return QFileDialog::getSaveFileName(((QWidget*)(parent)), QString(caption), QString(dir), QString(filter), new QString(selectedFilter), ((QFileDialog::Option)(options))).toUtf8().data();
}

