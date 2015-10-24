#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QBluetoothHostInfo_NewQBluetoothHostInfo();
QtObjectPtr QBluetoothHostInfo_NewQBluetoothHostInfo2(QtObjectPtr other);
char* QBluetoothHostInfo_Name(QtObjectPtr ptr);
void QBluetoothHostInfo_SetAddress(QtObjectPtr ptr, QtObjectPtr address);
void QBluetoothHostInfo_SetName(QtObjectPtr ptr, char* name);
void QBluetoothHostInfo_DestroyQBluetoothHostInfo(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif