#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QBluetoothTransferRequest_NewQBluetoothTransferRequest(QtObjectPtr address);
QtObjectPtr QBluetoothTransferRequest_NewQBluetoothTransferRequest2(QtObjectPtr other);
char* QBluetoothTransferRequest_Attribute(QtObjectPtr ptr, int code, char* defaultValue);
void QBluetoothTransferRequest_SetAttribute(QtObjectPtr ptr, int code, char* value);
void QBluetoothTransferRequest_DestroyQBluetoothTransferRequest(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif