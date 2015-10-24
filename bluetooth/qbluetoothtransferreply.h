#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QBluetoothTransferReply_Abort(QtObjectPtr ptr);
int QBluetoothTransferReply_Error(QtObjectPtr ptr);
char* QBluetoothTransferReply_ErrorString(QtObjectPtr ptr);
void QBluetoothTransferReply_ConnectFinished(QtObjectPtr ptr);
void QBluetoothTransferReply_DisconnectFinished(QtObjectPtr ptr);
int QBluetoothTransferReply_IsFinished(QtObjectPtr ptr);
int QBluetoothTransferReply_IsRunning(QtObjectPtr ptr);
QtObjectPtr QBluetoothTransferReply_Manager(QtObjectPtr ptr);
void QBluetoothTransferReply_DestroyQBluetoothTransferReply(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif