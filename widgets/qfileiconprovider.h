#ifdef __cplusplus
extern "C" {
#endif

void* QFileIconProvider_NewQFileIconProvider();
int QFileIconProvider_Options(void* ptr);
void QFileIconProvider_SetOptions(void* ptr, int options);
char* QFileIconProvider_Type(void* ptr, void* info);
void QFileIconProvider_DestroyQFileIconProvider(void* ptr);

#ifdef __cplusplus
}
#endif