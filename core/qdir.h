#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QDir_NewQDir(QtObjectPtr dir);
QtObjectPtr QDir_NewQDir2(char* path);
QtObjectPtr QDir_NewQDir3(char* path, char* nameFilter, int sort, int filters);
char* QDir_AbsoluteFilePath(QtObjectPtr ptr, char* fileName);
char* QDir_AbsolutePath(QtObjectPtr ptr);
void QDir_QDir_AddSearchPath(char* prefix, char* path);
char* QDir_CanonicalPath(QtObjectPtr ptr);
int QDir_Cd(QtObjectPtr ptr, char* dirName);
int QDir_CdUp(QtObjectPtr ptr);
char* QDir_QDir_CleanPath(char* path);
char* QDir_QDir_CurrentPath();
char* QDir_DirName(QtObjectPtr ptr);
char* QDir_EntryList2(QtObjectPtr ptr, int filters, int sort);
char* QDir_EntryList(QtObjectPtr ptr, char* nameFilters, int filters, int sort);
int QDir_Exists2(QtObjectPtr ptr);
int QDir_Exists(QtObjectPtr ptr, char* name);
char* QDir_FilePath(QtObjectPtr ptr, char* fileName);
int QDir_Filter(QtObjectPtr ptr);
char* QDir_QDir_FromNativeSeparators(char* pathName);
char* QDir_QDir_HomePath();
int QDir_IsAbsolute(QtObjectPtr ptr);
int QDir_QDir_IsAbsolutePath(char* path);
int QDir_IsReadable(QtObjectPtr ptr);
int QDir_IsRelative(QtObjectPtr ptr);
int QDir_QDir_IsRelativePath(char* path);
int QDir_IsRoot(QtObjectPtr ptr);
int QDir_MakeAbsolute(QtObjectPtr ptr);
int QDir_QDir_Match(char* filter, char* fileName);
int QDir_QDir_Match2(char* filters, char* fileName);
int QDir_Mkdir(QtObjectPtr ptr, char* dirName);
int QDir_Mkpath(QtObjectPtr ptr, char* dirPath);
char* QDir_NameFilters(QtObjectPtr ptr);
char* QDir_Path(QtObjectPtr ptr);
void QDir_Refresh(QtObjectPtr ptr);
char* QDir_RelativeFilePath(QtObjectPtr ptr, char* fileName);
int QDir_RemoveRecursively(QtObjectPtr ptr);
int QDir_Rename(QtObjectPtr ptr, char* oldName, char* newName);
int QDir_Rmdir(QtObjectPtr ptr, char* dirName);
int QDir_Rmpath(QtObjectPtr ptr, char* dirPath);
char* QDir_QDir_RootPath();
char* QDir_QDir_SearchPaths(char* prefix);
int QDir_QDir_SetCurrent(char* path);
void QDir_SetFilter(QtObjectPtr ptr, int filters);
void QDir_SetNameFilters(QtObjectPtr ptr, char* nameFilters);
void QDir_SetPath(QtObjectPtr ptr, char* path);
void QDir_QDir_SetSearchPaths(char* prefix, char* searchPaths);
void QDir_SetSorting(QtObjectPtr ptr, int sort);
int QDir_Sorting(QtObjectPtr ptr);
void QDir_Swap(QtObjectPtr ptr, QtObjectPtr other);
char* QDir_QDir_TempPath();
char* QDir_QDir_ToNativeSeparators(char* pathName);
void QDir_DestroyQDir(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif