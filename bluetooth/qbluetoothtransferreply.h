#ifdef __cplusplus
extern "C" {
#endif

void QBluetoothTransferReply_Abort(void* ptr);
int QBluetoothTransferReply_Error(void* ptr);
char* QBluetoothTransferReply_ErrorString(void* ptr);
void QBluetoothTransferReply_ConnectFinished(void* ptr);
void QBluetoothTransferReply_DisconnectFinished(void* ptr);
int QBluetoothTransferReply_IsFinished(void* ptr);
int QBluetoothTransferReply_IsRunning(void* ptr);
void* QBluetoothTransferReply_Manager(void* ptr);
void QBluetoothTransferReply_DestroyQBluetoothTransferReply(void* ptr);

#ifdef __cplusplus
}
#endif