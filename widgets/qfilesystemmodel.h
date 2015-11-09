#ifdef __cplusplus
extern "C" {
#endif

int QFileSystemModel_FilePathRole_Type();
int QFileSystemModel_FileNameRole_Type();
int QFileSystemModel_FilePermissions_Type();
int QFileSystemModel_IsReadOnly(void* ptr);
int QFileSystemModel_NameFilterDisables(void* ptr);
int QFileSystemModel_ResolveSymlinks(void* ptr);
int QFileSystemModel_Rmdir(void* ptr, void* index);
void QFileSystemModel_SetNameFilterDisables(void* ptr, int enable);
void QFileSystemModel_SetReadOnly(void* ptr, int enable);
void QFileSystemModel_SetResolveSymlinks(void* ptr, int enable);
void* QFileSystemModel_NewQFileSystemModel(void* parent);
int QFileSystemModel_CanFetchMore(void* ptr, void* parent);
int QFileSystemModel_ColumnCount(void* ptr, void* parent);
void* QFileSystemModel_Data(void* ptr, void* index, int role);
void QFileSystemModel_ConnectDirectoryLoaded(void* ptr);
void QFileSystemModel_DisconnectDirectoryLoaded(void* ptr);
int QFileSystemModel_DropMimeData(void* ptr, void* data, int action, int row, int column, void* parent);
void QFileSystemModel_FetchMore(void* ptr, void* parent);
char* QFileSystemModel_FileName(void* ptr, void* index);
char* QFileSystemModel_FilePath(void* ptr, void* index);
void QFileSystemModel_ConnectFileRenamed(void* ptr);
void QFileSystemModel_DisconnectFileRenamed(void* ptr);
int QFileSystemModel_Filter(void* ptr);
int QFileSystemModel_Flags(void* ptr, void* index);
int QFileSystemModel_HasChildren(void* ptr, void* parent);
void* QFileSystemModel_HeaderData(void* ptr, int section, int orientation, int role);
void* QFileSystemModel_IconProvider(void* ptr);
void* QFileSystemModel_Index2(void* ptr, char* path, int column);
void* QFileSystemModel_Index(void* ptr, int row, int column, void* parent);
int QFileSystemModel_IsDir(void* ptr, void* index);
void* QFileSystemModel_LastModified(void* ptr, void* index);
char* QFileSystemModel_MimeTypes(void* ptr);
void* QFileSystemModel_Mkdir(void* ptr, void* parent, char* name);
void* QFileSystemModel_MyComputer(void* ptr, int role);
char* QFileSystemModel_NameFilters(void* ptr);
void* QFileSystemModel_Parent(void* ptr, void* index);
void* QFileSystemModel_RootDirectory(void* ptr);
char* QFileSystemModel_RootPath(void* ptr);
void QFileSystemModel_ConnectRootPathChanged(void* ptr);
void QFileSystemModel_DisconnectRootPathChanged(void* ptr);
int QFileSystemModel_RowCount(void* ptr, void* parent);
int QFileSystemModel_SetData(void* ptr, void* idx, void* value, int role);
void QFileSystemModel_SetFilter(void* ptr, int filters);
void QFileSystemModel_SetIconProvider(void* ptr, void* provider);
void QFileSystemModel_SetNameFilters(void* ptr, char* filters);
void* QFileSystemModel_SetRootPath(void* ptr, char* newPath);
void QFileSystemModel_Sort(void* ptr, int column, int order);
int QFileSystemModel_SupportedDropActions(void* ptr);
char* QFileSystemModel_Type(void* ptr, void* index);
void QFileSystemModel_DestroyQFileSystemModel(void* ptr);

#ifdef __cplusplus
}
#endif