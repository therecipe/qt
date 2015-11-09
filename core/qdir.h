#ifdef __cplusplus
extern "C" {
#endif

void* QDir_NewQDir(void* dir);
void* QDir_NewQDir2(char* path);
void* QDir_NewQDir3(char* path, char* nameFilter, int sort, int filters);
char* QDir_AbsoluteFilePath(void* ptr, char* fileName);
char* QDir_AbsolutePath(void* ptr);
void QDir_QDir_AddSearchPath(char* prefix, char* path);
char* QDir_CanonicalPath(void* ptr);
int QDir_Cd(void* ptr, char* dirName);
int QDir_CdUp(void* ptr);
char* QDir_QDir_CleanPath(char* path);
void* QDir_QDir_Current();
char* QDir_QDir_CurrentPath();
char* QDir_DirName(void* ptr);
char* QDir_EntryList2(void* ptr, int filters, int sort);
char* QDir_EntryList(void* ptr, char* nameFilters, int filters, int sort);
int QDir_Exists2(void* ptr);
int QDir_Exists(void* ptr, char* name);
char* QDir_FilePath(void* ptr, char* fileName);
int QDir_Filter(void* ptr);
char* QDir_QDir_FromNativeSeparators(char* pathName);
void* QDir_QDir_Home();
char* QDir_QDir_HomePath();
int QDir_IsAbsolute(void* ptr);
int QDir_QDir_IsAbsolutePath(char* path);
int QDir_IsReadable(void* ptr);
int QDir_IsRelative(void* ptr);
int QDir_QDir_IsRelativePath(char* path);
int QDir_IsRoot(void* ptr);
int QDir_MakeAbsolute(void* ptr);
int QDir_QDir_Match(char* filter, char* fileName);
int QDir_QDir_Match2(char* filters, char* fileName);
int QDir_Mkdir(void* ptr, char* dirName);
int QDir_Mkpath(void* ptr, char* dirPath);
char* QDir_NameFilters(void* ptr);
char* QDir_Path(void* ptr);
void QDir_Refresh(void* ptr);
char* QDir_RelativeFilePath(void* ptr, char* fileName);
int QDir_RemoveRecursively(void* ptr);
int QDir_Rename(void* ptr, char* oldName, char* newName);
int QDir_Rmdir(void* ptr, char* dirName);
int QDir_Rmpath(void* ptr, char* dirPath);
void* QDir_QDir_Root();
char* QDir_QDir_RootPath();
char* QDir_QDir_SearchPaths(char* prefix);
int QDir_QDir_SetCurrent(char* path);
void QDir_SetFilter(void* ptr, int filters);
void QDir_SetNameFilters(void* ptr, char* nameFilters);
void QDir_SetPath(void* ptr, char* path);
void QDir_QDir_SetSearchPaths(char* prefix, char* searchPaths);
void QDir_SetSorting(void* ptr, int sort);
int QDir_Sorting(void* ptr);
void QDir_Swap(void* ptr, void* other);
void* QDir_QDir_Temp();
char* QDir_QDir_TempPath();
char* QDir_QDir_ToNativeSeparators(char* pathName);
void QDir_DestroyQDir(void* ptr);

#ifdef __cplusplus
}
#endif