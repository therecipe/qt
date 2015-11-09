#ifdef __cplusplus
extern "C" {
#endif

int QIODevice_GetChar(void* ptr, char* c);
int QIODevice_PutChar(void* ptr, char* c);
void QIODevice_ConnectAboutToClose(void* ptr);
void QIODevice_DisconnectAboutToClose(void* ptr);
int QIODevice_AtEnd(void* ptr);
int QIODevice_CanReadLine(void* ptr);
void QIODevice_Close(void* ptr);
char* QIODevice_ErrorString(void* ptr);
int QIODevice_IsOpen(void* ptr);
int QIODevice_IsReadable(void* ptr);
int QIODevice_IsSequential(void* ptr);
int QIODevice_IsTextModeEnabled(void* ptr);
int QIODevice_IsWritable(void* ptr);
int QIODevice_Open(void* ptr, int mode);
int QIODevice_OpenMode(void* ptr);
void* QIODevice_ReadAll(void* ptr);
void QIODevice_ConnectReadChannelFinished(void* ptr);
void QIODevice_DisconnectReadChannelFinished(void* ptr);
void QIODevice_ConnectReadyRead(void* ptr);
void QIODevice_DisconnectReadyRead(void* ptr);
int QIODevice_Reset(void* ptr);
void QIODevice_SetTextModeEnabled(void* ptr, int enabled);
void QIODevice_UngetChar(void* ptr, char* c);
int QIODevice_WaitForBytesWritten(void* ptr, int msecs);
int QIODevice_WaitForReadyRead(void* ptr, int msecs);
void QIODevice_DestroyQIODevice(void* ptr);

#ifdef __cplusplus
}
#endif