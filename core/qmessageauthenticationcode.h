#ifdef __cplusplus
extern "C" {
#endif

void* QMessageAuthenticationCode_NewQMessageAuthenticationCode(int method, void* key);
int QMessageAuthenticationCode_AddData2(void* ptr, void* device);
void QMessageAuthenticationCode_AddData3(void* ptr, void* data);
void QMessageAuthenticationCode_AddData(void* ptr, char* data, int length);
void* QMessageAuthenticationCode_QMessageAuthenticationCode_Hash(void* message, void* key, int method);
void QMessageAuthenticationCode_Reset(void* ptr);
void* QMessageAuthenticationCode_Result(void* ptr);
void QMessageAuthenticationCode_SetKey(void* ptr, void* key);
void QMessageAuthenticationCode_DestroyQMessageAuthenticationCode(void* ptr);

#ifdef __cplusplus
}
#endif