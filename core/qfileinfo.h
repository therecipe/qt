#ifdef __cplusplus
extern "C" {
#endif

void* QFileInfo_NewQFileInfo();
void* QFileInfo_NewQFileInfo5(void* dir, char* file);
void* QFileInfo_NewQFileInfo4(void* file);
void* QFileInfo_NewQFileInfo6(void* fileinfo);
void* QFileInfo_NewQFileInfo3(char* file);
void* QFileInfo_AbsoluteDir(void* ptr);
char* QFileInfo_AbsoluteFilePath(void* ptr);
char* QFileInfo_AbsolutePath(void* ptr);
char* QFileInfo_BaseName(void* ptr);
char* QFileInfo_BundleName(void* ptr);
int QFileInfo_Caching(void* ptr);
char* QFileInfo_CanonicalFilePath(void* ptr);
char* QFileInfo_CanonicalPath(void* ptr);
char* QFileInfo_CompleteBaseName(void* ptr);
char* QFileInfo_CompleteSuffix(void* ptr);
void* QFileInfo_Created(void* ptr);
void* QFileInfo_Dir(void* ptr);
int QFileInfo_QFileInfo_Exists2(char* file);
int QFileInfo_Exists(void* ptr);
char* QFileInfo_FileName(void* ptr);
char* QFileInfo_FilePath(void* ptr);
char* QFileInfo_Group(void* ptr);
int QFileInfo_IsAbsolute(void* ptr);
int QFileInfo_IsBundle(void* ptr);
int QFileInfo_IsDir(void* ptr);
int QFileInfo_IsExecutable(void* ptr);
int QFileInfo_IsFile(void* ptr);
int QFileInfo_IsHidden(void* ptr);
int QFileInfo_IsNativePath(void* ptr);
int QFileInfo_IsReadable(void* ptr);
int QFileInfo_IsRelative(void* ptr);
int QFileInfo_IsRoot(void* ptr);
int QFileInfo_IsSymLink(void* ptr);
int QFileInfo_IsWritable(void* ptr);
void* QFileInfo_LastModified(void* ptr);
void* QFileInfo_LastRead(void* ptr);
int QFileInfo_MakeAbsolute(void* ptr);
char* QFileInfo_Owner(void* ptr);
char* QFileInfo_Path(void* ptr);
void QFileInfo_Refresh(void* ptr);
void QFileInfo_SetCaching(void* ptr, int enable);
void QFileInfo_SetFile3(void* ptr, void* dir, char* file);
void QFileInfo_SetFile2(void* ptr, void* file);
void QFileInfo_SetFile(void* ptr, char* file);
char* QFileInfo_Suffix(void* ptr);
void QFileInfo_Swap(void* ptr, void* other);
char* QFileInfo_SymLinkTarget(void* ptr);
void QFileInfo_DestroyQFileInfo(void* ptr);

#ifdef __cplusplus
}
#endif