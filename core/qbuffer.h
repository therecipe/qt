#ifdef __cplusplus
extern "C" {
#endif

void* QBuffer_NewQBuffer2(void* byteArray, void* parent);
void* QBuffer_NewQBuffer(void* parent);
int QBuffer_AtEnd(void* ptr);
void* QBuffer_Buffer(void* ptr);
void* QBuffer_Buffer2(void* ptr);
int QBuffer_CanReadLine(void* ptr);
void QBuffer_Close(void* ptr);
void* QBuffer_Data(void* ptr);
int QBuffer_Open(void* ptr, int flags);
void QBuffer_SetBuffer(void* ptr, void* byteArray);
void QBuffer_SetData(void* ptr, void* data);
void QBuffer_SetData2(void* ptr, char* data, int size);
void QBuffer_DestroyQBuffer(void* ptr);

#ifdef __cplusplus
}
#endif