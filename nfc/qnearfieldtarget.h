#ifdef __cplusplus
extern "C" {
#endif

int QNearFieldTarget_AccessMethods(void* ptr);
void QNearFieldTarget_ConnectDisconnected(void* ptr);
void QNearFieldTarget_DisconnectDisconnected(void* ptr);
int QNearFieldTarget_HasNdefMessage(void* ptr);
int QNearFieldTarget_IsProcessingCommand(void* ptr);
void QNearFieldTarget_ConnectNdefMessagesWritten(void* ptr);
void QNearFieldTarget_DisconnectNdefMessagesWritten(void* ptr);
int QNearFieldTarget_Type(void* ptr);
void* QNearFieldTarget_Uid(void* ptr);
void QNearFieldTarget_DestroyQNearFieldTarget(void* ptr);

#ifdef __cplusplus
}
#endif