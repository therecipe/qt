#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QBluetoothAddress_NewQBluetoothAddress();
QtObjectPtr QBluetoothAddress_NewQBluetoothAddress4(QtObjectPtr other);
QtObjectPtr QBluetoothAddress_NewQBluetoothAddress3(char* address);
void QBluetoothAddress_Clear(QtObjectPtr ptr);
int QBluetoothAddress_IsNull(QtObjectPtr ptr);
char* QBluetoothAddress_ToString(QtObjectPtr ptr);
void QBluetoothAddress_DestroyQBluetoothAddress(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif