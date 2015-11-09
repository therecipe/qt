#ifdef __cplusplus
extern "C" {
#endif

void* QBluetoothTransferRequest_NewQBluetoothTransferRequest(void* address);
void* QBluetoothTransferRequest_NewQBluetoothTransferRequest2(void* other);
void* QBluetoothTransferRequest_Attribute(void* ptr, int code, void* defaultValue);
void QBluetoothTransferRequest_SetAttribute(void* ptr, int code, void* value);
void QBluetoothTransferRequest_DestroyQBluetoothTransferRequest(void* ptr);

#ifdef __cplusplus
}
#endif