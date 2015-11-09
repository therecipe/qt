#ifdef __cplusplus
extern "C" {
#endif

void* QBluetoothTransferManager_Put(void* ptr, void* request, void* data);
void* QBluetoothTransferManager_NewQBluetoothTransferManager(void* parent);
void QBluetoothTransferManager_ConnectFinished(void* ptr);
void QBluetoothTransferManager_DisconnectFinished(void* ptr);
void QBluetoothTransferManager_DestroyQBluetoothTransferManager(void* ptr);

#ifdef __cplusplus
}
#endif