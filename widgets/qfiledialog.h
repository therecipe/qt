#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QFileDialog_NewQFileDialog(QtObjectPtr parent, int flags);
int QFileDialog_AcceptMode(QtObjectPtr ptr);
int QFileDialog_ConfirmOverwrite(QtObjectPtr ptr);
char* QFileDialog_DefaultSuffix(QtObjectPtr ptr);
int QFileDialog_FileMode(QtObjectPtr ptr);
int QFileDialog_IsNameFilterDetailsVisible(QtObjectPtr ptr);
int QFileDialog_IsReadOnly(QtObjectPtr ptr);
int QFileDialog_Options(QtObjectPtr ptr);
int QFileDialog_ResolveSymlinks(QtObjectPtr ptr);
void QFileDialog_SetAcceptMode(QtObjectPtr ptr, int mode);
void QFileDialog_SetConfirmOverwrite(QtObjectPtr ptr, int enabled);
void QFileDialog_SetDefaultSuffix(QtObjectPtr ptr, char* suffix);
void QFileDialog_SetFileMode(QtObjectPtr ptr, int mode);
void QFileDialog_SetNameFilterDetailsVisible(QtObjectPtr ptr, int enabled);
void QFileDialog_SetOptions(QtObjectPtr ptr, int options);
void QFileDialog_SetReadOnly(QtObjectPtr ptr, int enabled);
void QFileDialog_SetResolveSymlinks(QtObjectPtr ptr, int enabled);
void QFileDialog_SetViewMode(QtObjectPtr ptr, int mode);
int QFileDialog_ViewMode(QtObjectPtr ptr);
QtObjectPtr QFileDialog_NewQFileDialog2(QtObjectPtr parent, char* caption, char* directory, char* filter);
void QFileDialog_ConnectCurrentChanged(QtObjectPtr ptr);
void QFileDialog_DisconnectCurrentChanged(QtObjectPtr ptr);
void QFileDialog_ConnectCurrentUrlChanged(QtObjectPtr ptr);
void QFileDialog_DisconnectCurrentUrlChanged(QtObjectPtr ptr);
void QFileDialog_ConnectDirectoryEntered(QtObjectPtr ptr);
void QFileDialog_DisconnectDirectoryEntered(QtObjectPtr ptr);
char* QFileDialog_DirectoryUrl(QtObjectPtr ptr);
void QFileDialog_ConnectDirectoryUrlEntered(QtObjectPtr ptr);
void QFileDialog_DisconnectDirectoryUrlEntered(QtObjectPtr ptr);
void QFileDialog_ConnectFileSelected(QtObjectPtr ptr);
void QFileDialog_DisconnectFileSelected(QtObjectPtr ptr);
void QFileDialog_ConnectFilesSelected(QtObjectPtr ptr);
void QFileDialog_DisconnectFilesSelected(QtObjectPtr ptr);
int QFileDialog_Filter(QtObjectPtr ptr);
void QFileDialog_ConnectFilterSelected(QtObjectPtr ptr);
void QFileDialog_DisconnectFilterSelected(QtObjectPtr ptr);
char* QFileDialog_QFileDialog_GetExistingDirectory(QtObjectPtr parent, char* caption, char* dir, int options);
char* QFileDialog_QFileDialog_GetExistingDirectoryUrl(QtObjectPtr parent, char* caption, char* dir, int options, char* supportedSchemes);
char* QFileDialog_QFileDialog_GetOpenFileName(QtObjectPtr parent, char* caption, char* dir, char* filter, char* selectedFilter, int options);
char* QFileDialog_QFileDialog_GetOpenFileNames(QtObjectPtr parent, char* caption, char* dir, char* filter, char* selectedFilter, int options);
char* QFileDialog_QFileDialog_GetOpenFileUrl(QtObjectPtr parent, char* caption, char* dir, char* filter, char* selectedFilter, int options, char* supportedSchemes);
char* QFileDialog_QFileDialog_GetSaveFileName(QtObjectPtr parent, char* caption, char* dir, char* filter, char* selectedFilter, int options);
char* QFileDialog_QFileDialog_GetSaveFileUrl(QtObjectPtr parent, char* caption, char* dir, char* filter, char* selectedFilter, int options, char* supportedSchemes);
char* QFileDialog_History(QtObjectPtr ptr);
QtObjectPtr QFileDialog_IconProvider(QtObjectPtr ptr);
QtObjectPtr QFileDialog_ItemDelegate(QtObjectPtr ptr);
char* QFileDialog_LabelText(QtObjectPtr ptr, int label);
char* QFileDialog_MimeTypeFilters(QtObjectPtr ptr);
char* QFileDialog_NameFilters(QtObjectPtr ptr);
void QFileDialog_Open(QtObjectPtr ptr, QtObjectPtr receiver, char* member);
QtObjectPtr QFileDialog_ProxyModel(QtObjectPtr ptr);
int QFileDialog_RestoreState(QtObjectPtr ptr, QtObjectPtr state);
void QFileDialog_SelectFile(QtObjectPtr ptr, char* filename);
void QFileDialog_SelectMimeTypeFilter(QtObjectPtr ptr, char* filter);
void QFileDialog_SelectNameFilter(QtObjectPtr ptr, char* filter);
void QFileDialog_SelectUrl(QtObjectPtr ptr, char* url);
char* QFileDialog_SelectedFiles(QtObjectPtr ptr);
char* QFileDialog_SelectedNameFilter(QtObjectPtr ptr);
void QFileDialog_SetDirectory2(QtObjectPtr ptr, QtObjectPtr directory);
void QFileDialog_SetDirectory(QtObjectPtr ptr, char* directory);
void QFileDialog_SetDirectoryUrl(QtObjectPtr ptr, char* directory);
void QFileDialog_SetFilter(QtObjectPtr ptr, int filters);
void QFileDialog_SetHistory(QtObjectPtr ptr, char* paths);
void QFileDialog_SetIconProvider(QtObjectPtr ptr, QtObjectPtr provider);
void QFileDialog_SetItemDelegate(QtObjectPtr ptr, QtObjectPtr delegate);
void QFileDialog_SetLabelText(QtObjectPtr ptr, int label, char* text);
void QFileDialog_SetMimeTypeFilters(QtObjectPtr ptr, char* filters);
void QFileDialog_SetNameFilter(QtObjectPtr ptr, char* filter);
void QFileDialog_SetNameFilters(QtObjectPtr ptr, char* filters);
void QFileDialog_SetOption(QtObjectPtr ptr, int option, int on);
void QFileDialog_SetProxyModel(QtObjectPtr ptr, QtObjectPtr proxyModel);
void QFileDialog_SetVisible(QtObjectPtr ptr, int visible);
int QFileDialog_TestOption(QtObjectPtr ptr, int option);
void QFileDialog_ConnectUrlSelected(QtObjectPtr ptr);
void QFileDialog_DisconnectUrlSelected(QtObjectPtr ptr);
void QFileDialog_DestroyQFileDialog(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif