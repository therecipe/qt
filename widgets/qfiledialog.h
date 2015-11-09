#ifdef __cplusplus
extern "C" {
#endif

void* QFileDialog_NewQFileDialog(void* parent, int flags);
int QFileDialog_AcceptMode(void* ptr);
int QFileDialog_ConfirmOverwrite(void* ptr);
char* QFileDialog_DefaultSuffix(void* ptr);
int QFileDialog_FileMode(void* ptr);
int QFileDialog_IsNameFilterDetailsVisible(void* ptr);
int QFileDialog_IsReadOnly(void* ptr);
int QFileDialog_Options(void* ptr);
int QFileDialog_ResolveSymlinks(void* ptr);
void QFileDialog_SetAcceptMode(void* ptr, int mode);
void QFileDialog_SetConfirmOverwrite(void* ptr, int enabled);
void QFileDialog_SetDefaultSuffix(void* ptr, char* suffix);
void QFileDialog_SetFileMode(void* ptr, int mode);
void QFileDialog_SetNameFilterDetailsVisible(void* ptr, int enabled);
void QFileDialog_SetOptions(void* ptr, int options);
void QFileDialog_SetReadOnly(void* ptr, int enabled);
void QFileDialog_SetResolveSymlinks(void* ptr, int enabled);
void QFileDialog_SetViewMode(void* ptr, int mode);
int QFileDialog_ViewMode(void* ptr);
void* QFileDialog_NewQFileDialog2(void* parent, char* caption, char* directory, char* filter);
void QFileDialog_ConnectCurrentChanged(void* ptr);
void QFileDialog_DisconnectCurrentChanged(void* ptr);
void* QFileDialog_Directory(void* ptr);
void QFileDialog_ConnectDirectoryEntered(void* ptr);
void QFileDialog_DisconnectDirectoryEntered(void* ptr);
void QFileDialog_ConnectFileSelected(void* ptr);
void QFileDialog_DisconnectFileSelected(void* ptr);
void QFileDialog_ConnectFilesSelected(void* ptr);
void QFileDialog_DisconnectFilesSelected(void* ptr);
int QFileDialog_Filter(void* ptr);
void QFileDialog_ConnectFilterSelected(void* ptr);
void QFileDialog_DisconnectFilterSelected(void* ptr);
char* QFileDialog_QFileDialog_GetExistingDirectory(void* parent, char* caption, char* dir, int options);
char* QFileDialog_QFileDialog_GetOpenFileName(void* parent, char* caption, char* dir, char* filter, char* selectedFilter, int options);
char* QFileDialog_QFileDialog_GetOpenFileNames(void* parent, char* caption, char* dir, char* filter, char* selectedFilter, int options);
char* QFileDialog_QFileDialog_GetSaveFileName(void* parent, char* caption, char* dir, char* filter, char* selectedFilter, int options);
char* QFileDialog_History(void* ptr);
void* QFileDialog_IconProvider(void* ptr);
void* QFileDialog_ItemDelegate(void* ptr);
char* QFileDialog_LabelText(void* ptr, int label);
char* QFileDialog_MimeTypeFilters(void* ptr);
char* QFileDialog_NameFilters(void* ptr);
void QFileDialog_Open(void* ptr, void* receiver, char* member);
void* QFileDialog_ProxyModel(void* ptr);
int QFileDialog_RestoreState(void* ptr, void* state);
void* QFileDialog_SaveState(void* ptr);
void QFileDialog_SelectFile(void* ptr, char* filename);
void QFileDialog_SelectMimeTypeFilter(void* ptr, char* filter);
void QFileDialog_SelectNameFilter(void* ptr, char* filter);
void QFileDialog_SelectUrl(void* ptr, void* url);
char* QFileDialog_SelectedFiles(void* ptr);
char* QFileDialog_SelectedNameFilter(void* ptr);
void QFileDialog_SetDirectory2(void* ptr, void* directory);
void QFileDialog_SetDirectory(void* ptr, char* directory);
void QFileDialog_SetDirectoryUrl(void* ptr, void* directory);
void QFileDialog_SetFilter(void* ptr, int filters);
void QFileDialog_SetHistory(void* ptr, char* paths);
void QFileDialog_SetIconProvider(void* ptr, void* provider);
void QFileDialog_SetItemDelegate(void* ptr, void* delegate);
void QFileDialog_SetLabelText(void* ptr, int label, char* text);
void QFileDialog_SetMimeTypeFilters(void* ptr, char* filters);
void QFileDialog_SetNameFilter(void* ptr, char* filter);
void QFileDialog_SetNameFilters(void* ptr, char* filters);
void QFileDialog_SetOption(void* ptr, int option, int on);
void QFileDialog_SetProxyModel(void* ptr, void* proxyModel);
void QFileDialog_SetVisible(void* ptr, int visible);
int QFileDialog_TestOption(void* ptr, int option);
void QFileDialog_DestroyQFileDialog(void* ptr);

#ifdef __cplusplus
}
#endif