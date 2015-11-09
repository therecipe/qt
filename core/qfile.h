#ifdef __cplusplus
extern "C" {
#endif

void* QFile_NewQFile3(void* parent);
void* QFile_NewQFile(char* name);
void* QFile_NewQFile4(char* name, void* parent);
int QFile_QFile_Copy2(char* fileName, char* newName);
int QFile_Copy(void* ptr, char* newName);
char* QFile_QFile_DecodeName(void* localFileName);
char* QFile_QFile_DecodeName2(char* localFileName);
void* QFile_QFile_EncodeName(char* fileName);
int QFile_QFile_Exists(char* fileName);
int QFile_Exists2(void* ptr);
char* QFile_FileName(void* ptr);
int QFile_QFile_Link2(char* fileName, char* linkName);
int QFile_Link(void* ptr, char* linkName);
int QFile_Open(void* ptr, int mode);
int QFile_Open3(void* ptr, int fd, int mode, int handleFlags);
int QFile_QFile_Permissions2(char* fileName);
int QFile_Permissions(void* ptr);
int QFile_Rename(void* ptr, char* newName);
int QFile_QFile_Rename2(char* oldName, char* newName);
void QFile_SetFileName(void* ptr, char* name);
int QFile_SetPermissions(void* ptr, int permissions);
int QFile_QFile_SetPermissions2(char* fileName, int permissions);
char* QFile_QFile_SymLinkTarget(char* fileName);
char* QFile_SymLinkTarget2(void* ptr);
void QFile_DestroyQFile(void* ptr);

#ifdef __cplusplus
}
#endif