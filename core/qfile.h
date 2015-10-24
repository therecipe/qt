#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QFile_NewQFile3(QtObjectPtr parent);
QtObjectPtr QFile_NewQFile(char* name);
QtObjectPtr QFile_NewQFile4(char* name, QtObjectPtr parent);
int QFile_QFile_Copy2(char* fileName, char* newName);
int QFile_Copy(QtObjectPtr ptr, char* newName);
char* QFile_QFile_DecodeName(QtObjectPtr localFileName);
char* QFile_QFile_DecodeName2(char* localFileName);
int QFile_QFile_Exists(char* fileName);
int QFile_Exists2(QtObjectPtr ptr);
char* QFile_FileName(QtObjectPtr ptr);
int QFile_QFile_Link2(char* fileName, char* linkName);
int QFile_Link(QtObjectPtr ptr, char* linkName);
int QFile_Open(QtObjectPtr ptr, int mode);
int QFile_Open3(QtObjectPtr ptr, int fd, int mode, int handleFlags);
int QFile_QFile_Permissions2(char* fileName);
int QFile_Permissions(QtObjectPtr ptr);
int QFile_Rename(QtObjectPtr ptr, char* newName);
int QFile_QFile_Rename2(char* oldName, char* newName);
void QFile_SetFileName(QtObjectPtr ptr, char* name);
int QFile_SetPermissions(QtObjectPtr ptr, int permissions);
int QFile_QFile_SetPermissions2(char* fileName, int permissions);
char* QFile_QFile_SymLinkTarget(char* fileName);
char* QFile_SymLinkTarget2(QtObjectPtr ptr);
void QFile_DestroyQFile(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif