#ifdef __cplusplus
extern "C" {
#endif

void* QAbstractVideoBuffer_Handle(void* ptr);
int QAbstractVideoBuffer_HandleType(void* ptr);
int QAbstractVideoBuffer_MapMode(void* ptr);
void QAbstractVideoBuffer_Release(void* ptr);
void QAbstractVideoBuffer_Unmap(void* ptr);
void QAbstractVideoBuffer_DestroyQAbstractVideoBuffer(void* ptr);

#ifdef __cplusplus
}
#endif