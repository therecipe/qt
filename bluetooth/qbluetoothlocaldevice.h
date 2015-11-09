#ifdef __cplusplus
extern "C" {
#endif

void QBluetoothLocalDevice_ConnectError(void* ptr);
void QBluetoothLocalDevice_DisconnectError(void* ptr);
void QBluetoothLocalDevice_ConnectHostModeStateChanged(void* ptr);
void QBluetoothLocalDevice_DisconnectHostModeStateChanged(void* ptr);
int QBluetoothLocalDevice_IsValid(void* ptr);
void QBluetoothLocalDevice_DestroyQBluetoothLocalDevice(void* ptr);
void* QBluetoothLocalDevice_NewQBluetoothLocalDevice(void* parent);
void* QBluetoothLocalDevice_NewQBluetoothLocalDevice2(void* address, void* parent);
int QBluetoothLocalDevice_HostMode(void* ptr);
char* QBluetoothLocalDevice_Name(void* ptr);
void QBluetoothLocalDevice_PairingConfirmation(void* ptr, int accept);
int QBluetoothLocalDevice_PairingStatus(void* ptr, void* address);
void QBluetoothLocalDevice_PowerOn(void* ptr);
void QBluetoothLocalDevice_RequestPairing(void* ptr, void* address, int pairing);
void QBluetoothLocalDevice_SetHostMode(void* ptr, int mode);

#ifdef __cplusplus
}
#endif