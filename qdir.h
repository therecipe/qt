#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"

//Public Types
int QDir_Dirs();
int QDir_AllDirs();
int QDir_Files();
int QDir_Drives();
int QDir_NoSymLinks();
int QDir_NoDotAndDotDot();
int QDir_NoDot();
int QDir_NoDotDot();
int QDir_AllEntries();
int QDir_Readable();
int QDir_Writable();
int QDir_Executable();
int QDir_Modified();
int QDir_Hidden();
int QDir_System();
int QDir_Name();
int QDir_Time();
int QDir_Size();
int QDir_Type();
int QDir_Unsorted();
int QDir_NoSort();
int QDir_DirsFirst();
int QDir_DirsLast();
int QDir_Reversed();
int QDir_IgnoreCase();
int QDir_LocaleAware();
//Static Public Members
void QDir_AddSearchPath_String_String(char* prefix, char* path);
char* QDir_CleanPath_String(char* path);
char* QDir_CurrentPath();
char* QDir_FromNativeSeparators_String(char* pathName);
char* QDir_HomePath();
int QDir_IsAbsolutePath_String(char* path);
int QDir_IsRelativePath_String(char* path);
int QDir_Match_String_String(char* filter, char* fileName);
int QDir_Match_QStringList_String(char* filters, char* fileName);
char* QDir_RootPath();
char* QDir_SearchPaths_String(char* prefix);
int QDir_SetCurrent_String(char* path);
void QDir_SetSearchPaths_String_QStringList(char* prefix, char* searchPaths);
char* QDir_TempPath();
char* QDir_ToNativeSeparators_String(char* pathName);

#ifdef __cplusplus
}
#endif
