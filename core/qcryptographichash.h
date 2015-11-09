#ifdef __cplusplus
extern "C" {
#endif

void* QCryptographicHash_NewQCryptographicHash(int method);
int QCryptographicHash_AddData2(void* ptr, void* device);
void QCryptographicHash_AddData3(void* ptr, void* data);
void QCryptographicHash_AddData(void* ptr, char* data, int length);
void* QCryptographicHash_QCryptographicHash_Hash(void* data, int method);
void QCryptographicHash_Reset(void* ptr);
void* QCryptographicHash_Result(void* ptr);
void QCryptographicHash_DestroyQCryptographicHash(void* ptr);

#ifdef __cplusplus
}
#endif