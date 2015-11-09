#ifdef __cplusplus
extern "C" {
#endif

int QResource_QResource_RegisterResource(char* rccFileName, char* mapRoot);
int QResource_QResource_UnregisterResource(char* rccFileName, char* mapRoot);
void* QResource_NewQResource(char* file, void* locale);
char* QResource_AbsoluteFilePath(void* ptr);
char* QResource_FileName(void* ptr);
int QResource_IsCompressed(void* ptr);
int QResource_IsValid(void* ptr);
void QResource_SetFileName(void* ptr, char* file);
void QResource_SetLocale(void* ptr, void* locale);
void QResource_DestroyQResource(void* ptr);

#ifdef __cplusplus
}
#endif