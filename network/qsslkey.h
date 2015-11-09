#ifdef __cplusplus
extern "C" {
#endif

void* QSslKey_NewQSslKey();
void* QSslKey_NewQSslKey5(void* other);
void QSslKey_Clear(void* ptr);
int QSslKey_IsNull(void* ptr);
int QSslKey_Length(void* ptr);
void QSslKey_Swap(void* ptr, void* other);
void* QSslKey_ToDer(void* ptr, void* passPhrase);
void* QSslKey_ToPem(void* ptr, void* passPhrase);
void QSslKey_DestroyQSslKey(void* ptr);

#ifdef __cplusplus
}
#endif