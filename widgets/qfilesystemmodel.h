#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QFileSystemModel_FilePathRole_Type();
int QFileSystemModel_FileNameRole_Type();
int QFileSystemModel_FilePermissions_Type();
int QFileSystemModel_IsReadOnly(QtObjectPtr ptr);
int QFileSystemModel_NameFilterDisables(QtObjectPtr ptr);
int QFileSystemModel_ResolveSymlinks(QtObjectPtr ptr);
int QFileSystemModel_Rmdir(QtObjectPtr ptr, QtObjectPtr index);
void QFileSystemModel_SetNameFilterDisables(QtObjectPtr ptr, int enable);
void QFileSystemModel_SetReadOnly(QtObjectPtr ptr, int enable);
void QFileSystemModel_SetResolveSymlinks(QtObjectPtr ptr, int enable);
QtObjectPtr QFileSystemModel_NewQFileSystemModel(QtObjectPtr parent);
int QFileSystemModel_CanFetchMore(QtObjectPtr ptr, QtObjectPtr parent);
int QFileSystemModel_ColumnCount(QtObjectPtr ptr, QtObjectPtr parent);
char* QFileSystemModel_Data(QtObjectPtr ptr, QtObjectPtr index, int role);
void QFileSystemModel_ConnectDirectoryLoaded(QtObjectPtr ptr);
void QFileSystemModel_DisconnectDirectoryLoaded(QtObjectPtr ptr);
int QFileSystemModel_DropMimeData(QtObjectPtr ptr, QtObjectPtr data, int action, int row, int column, QtObjectPtr parent);
void QFileSystemModel_FetchMore(QtObjectPtr ptr, QtObjectPtr parent);
char* QFileSystemModel_FileName(QtObjectPtr ptr, QtObjectPtr index);
char* QFileSystemModel_FilePath(QtObjectPtr ptr, QtObjectPtr index);
void QFileSystemModel_ConnectFileRenamed(QtObjectPtr ptr);
void QFileSystemModel_DisconnectFileRenamed(QtObjectPtr ptr);
int QFileSystemModel_Filter(QtObjectPtr ptr);
int QFileSystemModel_Flags(QtObjectPtr ptr, QtObjectPtr index);
int QFileSystemModel_HasChildren(QtObjectPtr ptr, QtObjectPtr parent);
char* QFileSystemModel_HeaderData(QtObjectPtr ptr, int section, int orientation, int role);
QtObjectPtr QFileSystemModel_IconProvider(QtObjectPtr ptr);
QtObjectPtr QFileSystemModel_Index2(QtObjectPtr ptr, char* path, int column);
QtObjectPtr QFileSystemModel_Index(QtObjectPtr ptr, int row, int column, QtObjectPtr parent);
int QFileSystemModel_IsDir(QtObjectPtr ptr, QtObjectPtr index);
char* QFileSystemModel_MimeTypes(QtObjectPtr ptr);
QtObjectPtr QFileSystemModel_Mkdir(QtObjectPtr ptr, QtObjectPtr parent, char* name);
char* QFileSystemModel_MyComputer(QtObjectPtr ptr, int role);
char* QFileSystemModel_NameFilters(QtObjectPtr ptr);
QtObjectPtr QFileSystemModel_Parent(QtObjectPtr ptr, QtObjectPtr index);
char* QFileSystemModel_RootPath(QtObjectPtr ptr);
void QFileSystemModel_ConnectRootPathChanged(QtObjectPtr ptr);
void QFileSystemModel_DisconnectRootPathChanged(QtObjectPtr ptr);
int QFileSystemModel_RowCount(QtObjectPtr ptr, QtObjectPtr parent);
int QFileSystemModel_SetData(QtObjectPtr ptr, QtObjectPtr idx, char* value, int role);
void QFileSystemModel_SetFilter(QtObjectPtr ptr, int filters);
void QFileSystemModel_SetIconProvider(QtObjectPtr ptr, QtObjectPtr provider);
void QFileSystemModel_SetNameFilters(QtObjectPtr ptr, char* filters);
QtObjectPtr QFileSystemModel_SetRootPath(QtObjectPtr ptr, char* newPath);
void QFileSystemModel_Sort(QtObjectPtr ptr, int column, int order);
int QFileSystemModel_SupportedDropActions(QtObjectPtr ptr);
char* QFileSystemModel_Type(QtObjectPtr ptr, QtObjectPtr index);
void QFileSystemModel_DestroyQFileSystemModel(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif