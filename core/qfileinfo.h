#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QFileInfo_NewQFileInfo();
QtObjectPtr QFileInfo_NewQFileInfo5(QtObjectPtr dir, char* file);
QtObjectPtr QFileInfo_NewQFileInfo4(QtObjectPtr file);
QtObjectPtr QFileInfo_NewQFileInfo6(QtObjectPtr fileinfo);
QtObjectPtr QFileInfo_NewQFileInfo3(char* file);
char* QFileInfo_AbsoluteFilePath(QtObjectPtr ptr);
char* QFileInfo_AbsolutePath(QtObjectPtr ptr);
char* QFileInfo_BaseName(QtObjectPtr ptr);
char* QFileInfo_BundleName(QtObjectPtr ptr);
int QFileInfo_Caching(QtObjectPtr ptr);
char* QFileInfo_CanonicalFilePath(QtObjectPtr ptr);
char* QFileInfo_CanonicalPath(QtObjectPtr ptr);
char* QFileInfo_CompleteBaseName(QtObjectPtr ptr);
char* QFileInfo_CompleteSuffix(QtObjectPtr ptr);
int QFileInfo_QFileInfo_Exists2(char* file);
int QFileInfo_Exists(QtObjectPtr ptr);
char* QFileInfo_FileName(QtObjectPtr ptr);
char* QFileInfo_FilePath(QtObjectPtr ptr);
char* QFileInfo_Group(QtObjectPtr ptr);
int QFileInfo_IsAbsolute(QtObjectPtr ptr);
int QFileInfo_IsBundle(QtObjectPtr ptr);
int QFileInfo_IsDir(QtObjectPtr ptr);
int QFileInfo_IsExecutable(QtObjectPtr ptr);
int QFileInfo_IsFile(QtObjectPtr ptr);
int QFileInfo_IsHidden(QtObjectPtr ptr);
int QFileInfo_IsNativePath(QtObjectPtr ptr);
int QFileInfo_IsReadable(QtObjectPtr ptr);
int QFileInfo_IsRelative(QtObjectPtr ptr);
int QFileInfo_IsRoot(QtObjectPtr ptr);
int QFileInfo_IsSymLink(QtObjectPtr ptr);
int QFileInfo_IsWritable(QtObjectPtr ptr);
int QFileInfo_MakeAbsolute(QtObjectPtr ptr);
char* QFileInfo_Owner(QtObjectPtr ptr);
char* QFileInfo_Path(QtObjectPtr ptr);
void QFileInfo_Refresh(QtObjectPtr ptr);
void QFileInfo_SetCaching(QtObjectPtr ptr, int enable);
void QFileInfo_SetFile3(QtObjectPtr ptr, QtObjectPtr dir, char* file);
void QFileInfo_SetFile2(QtObjectPtr ptr, QtObjectPtr file);
void QFileInfo_SetFile(QtObjectPtr ptr, char* file);
char* QFileInfo_Suffix(QtObjectPtr ptr);
void QFileInfo_Swap(QtObjectPtr ptr, QtObjectPtr other);
char* QFileInfo_SymLinkTarget(QtObjectPtr ptr);
void QFileInfo_DestroyQFileInfo(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif