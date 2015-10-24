#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QBluetoothLocalDevice_ConnectError(QtObjectPtr ptr);
void QBluetoothLocalDevice_DisconnectError(QtObjectPtr ptr);
void QBluetoothLocalDevice_ConnectHostModeStateChanged(QtObjectPtr ptr);
void QBluetoothLocalDevice_DisconnectHostModeStateChanged(QtObjectPtr ptr);
int QBluetoothLocalDevice_IsValid(QtObjectPtr ptr);
void QBluetoothLocalDevice_DestroyQBluetoothLocalDevice(QtObjectPtr ptr);
QtObjectPtr QBluetoothLocalDevice_NewQBluetoothLocalDevice(QtObjectPtr parent);
QtObjectPtr QBluetoothLocalDevice_NewQBluetoothLocalDevice2(QtObjectPtr address, QtObjectPtr parent);
int QBluetoothLocalDevice_HostMode(QtObjectPtr ptr);
char* QBluetoothLocalDevice_Name(QtObjectPtr ptr);
void QBluetoothLocalDevice_PairingConfirmation(QtObjectPtr ptr, int accept);
int QBluetoothLocalDevice_PairingStatus(QtObjectPtr ptr, QtObjectPtr address);
void QBluetoothLocalDevice_PowerOn(QtObjectPtr ptr);
void QBluetoothLocalDevice_RequestPairing(QtObjectPtr ptr, QtObjectPtr address, int pairing);
void QBluetoothLocalDevice_SetHostMode(QtObjectPtr ptr, int mode);

#ifdef __cplusplus
}
#endif