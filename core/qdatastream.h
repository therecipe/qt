#ifdef __cplusplus
extern "C" {
#endif

void* QDataStream_NewQDataStream3(void* a, int mode);
int QDataStream_AtEnd(void* ptr);
void* QDataStream_NewQDataStream();
void* QDataStream_NewQDataStream2(void* d);
void* QDataStream_NewQDataStream4(void* a);
int QDataStream_ByteOrder(void* ptr);
void* QDataStream_Device(void* ptr);
int QDataStream_FloatingPointPrecision(void* ptr);
int QDataStream_ReadRawData(void* ptr, char* s, int len);
void QDataStream_ResetStatus(void* ptr);
void QDataStream_SetByteOrder(void* ptr, int bo);
void QDataStream_SetDevice(void* ptr, void* d);
void QDataStream_SetFloatingPointPrecision(void* ptr, int precision);
void QDataStream_SetStatus(void* ptr, int status);
void QDataStream_SetVersion(void* ptr, int v);
int QDataStream_SkipRawData(void* ptr, int len);
int QDataStream_Status(void* ptr);
int QDataStream_Version(void* ptr);
int QDataStream_WriteRawData(void* ptr, char* s, int len);
void QDataStream_DestroyQDataStream(void* ptr);

#ifdef __cplusplus
}
#endif