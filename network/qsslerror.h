#ifdef __cplusplus
extern "C" {
#endif

void* QSslError_NewQSslError();
void* QSslError_NewQSslError2(int error);
void* QSslError_NewQSslError3(int error, void* certificate);
void* QSslError_NewQSslError4(void* other);
int QSslError_Error(void* ptr);
char* QSslError_ErrorString(void* ptr);
void QSslError_Swap(void* ptr, void* other);
void QSslError_DestroyQSslError(void* ptr);

#ifdef __cplusplus
}
#endif