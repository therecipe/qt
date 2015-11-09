#ifdef __cplusplus
extern "C" {
#endif

void* QBluetoothAddress_NewQBluetoothAddress();
void* QBluetoothAddress_NewQBluetoothAddress4(void* other);
void* QBluetoothAddress_NewQBluetoothAddress3(char* address);
void QBluetoothAddress_Clear(void* ptr);
int QBluetoothAddress_IsNull(void* ptr);
char* QBluetoothAddress_ToString(void* ptr);
void QBluetoothAddress_DestroyQBluetoothAddress(void* ptr);

#ifdef __cplusplus
}
#endif