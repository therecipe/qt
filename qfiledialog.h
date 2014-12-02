#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"

//Public Types
int QFileDialog_AcceptOpen();
int QFileDialog_AcceptSave();
int QFileDialog_LookIn();
int QFileDialog_FileName();
int QFileDialog_FileType();
int QFileDialog_Accept();
int QFileDialog_Reject();
int QFileDialog_AnyFile();
int QFileDialog_ExistingFile();
int QFileDialog_Directory();
int QFileDialog_ExistingFiles();
int QFileDialog_DirectoryOnly();
int QFileDialog_ShowDirsOnly();
int QFileDialog_DontResolveSymlinks();
int QFileDialog_DontConfirmOverwrite();
int QFileDialog_DontUseNativeDialog();
int QFileDialog_ReadOnly();
int QFileDialog_HideNameFilterDetails();
int QFileDialog_DontUseSheet();
int QFileDialog_DontUseCustomDirectoryIcons();
int QFileDialog_Detail();
int QFileDialog_List();
//Public Functions
QtObjectPtr QFileDialog_New_QWidget_String_String_String(QtObjectPtr parent, char* caption, char* directory, char* filter);
void QFileDialog_Destroy(QtObjectPtr ptr);
int QFileDialog_AcceptMode(QtObjectPtr ptr);
char* QFileDialog_DefaultSuffix(QtObjectPtr ptr);
int QFileDialog_FileMode(QtObjectPtr ptr);
char* QFileDialog_History(QtObjectPtr ptr);
char* QFileDialog_LabelText_DialogLabel(QtObjectPtr ptr, int label);
char* QFileDialog_MimeTypeFilters(QtObjectPtr ptr);
char* QFileDialog_NameFilters(QtObjectPtr ptr);
void QFileDialog_Open_QObject_String(QtObjectPtr ptr, QtObjectPtr receiver, char* member);
int QFileDialog_Options(QtObjectPtr ptr);
void QFileDialog_SelectFile_String(QtObjectPtr ptr, char* filename);
void QFileDialog_SelectMimeTypeFilter_String(QtObjectPtr ptr, char* filter);
void QFileDialog_SelectNameFilter_String(QtObjectPtr ptr, char* filter);
char* QFileDialog_SelectedFiles(QtObjectPtr ptr);
char* QFileDialog_SelectedNameFilter(QtObjectPtr ptr);
void QFileDialog_SetAcceptMode_AcceptMode(QtObjectPtr ptr, int mode);
void QFileDialog_SetDefaultSuffix_String(QtObjectPtr ptr, char* suffix);
void QFileDialog_SetDirectory_String(QtObjectPtr ptr, char* directory);
void QFileDialog_SetFileMode_FileMode(QtObjectPtr ptr, int mode);
void QFileDialog_SetHistory_QStringList(QtObjectPtr ptr, char* paths);
void QFileDialog_SetLabelText_DialogLabel_String(QtObjectPtr ptr, int label, char* text);
void QFileDialog_SetMimeTypeFilters_QStringList(QtObjectPtr ptr, char* filters);
void QFileDialog_SetNameFilter_String(QtObjectPtr ptr, char* filter);
void QFileDialog_SetNameFilters_QStringList(QtObjectPtr ptr, char* filters);
void QFileDialog_SetOption_Option_Bool(QtObjectPtr ptr, int option, int on);
void QFileDialog_SetOptions_Option(QtObjectPtr ptr, int options);
void QFileDialog_SetViewMode_ViewMode(QtObjectPtr ptr, int mode);
int QFileDialog_TestOption_Option(QtObjectPtr ptr, int option);
int QFileDialog_ViewMode(QtObjectPtr ptr);
//Signals
void QFileDialog_ConnectSignalCurrentChanged(QtObjectPtr ptr);
void QFileDialog_DisconnectSignalCurrentChanged(QtObjectPtr ptr);
void QFileDialog_ConnectSignalDirectoryEntered(QtObjectPtr ptr);
void QFileDialog_DisconnectSignalDirectoryEntered(QtObjectPtr ptr);
void QFileDialog_ConnectSignalFileSelected(QtObjectPtr ptr);
void QFileDialog_DisconnectSignalFileSelected(QtObjectPtr ptr);
void QFileDialog_ConnectSignalFilesSelected(QtObjectPtr ptr);
void QFileDialog_DisconnectSignalFilesSelected(QtObjectPtr ptr);
void QFileDialog_ConnectSignalFilterSelected(QtObjectPtr ptr);
void QFileDialog_DisconnectSignalFilterSelected(QtObjectPtr ptr);
//Static Public Members
char* QFileDialog_GetExistingDirectory_QWidget_String_String_Option(QtObjectPtr parent, char* caption, char* dir, int options);
char* QFileDialog_GetOpenFileName_QWidget_String_String_String_String_Option(QtObjectPtr parent, char* caption, char* dir, char* filter, char* selectedFilter, int options);
char* QFileDialog_GetOpenFileNames_QWidget_String_String_String_String_Option(QtObjectPtr parent, char* caption, char* dir, char* filter, char* selectedFilter, int options);
char* QFileDialog_GetSaveFileName_QWidget_String_String_String_String_Option(QtObjectPtr parent, char* caption, char* dir, char* filter, char* selectedFilter, int options);

#ifdef __cplusplus
}
#endif
