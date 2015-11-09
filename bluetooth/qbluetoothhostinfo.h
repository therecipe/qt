#ifdef __cplusplus
extern "C" {
#endif

void* QBluetoothHostInfo_NewQBluetoothHostInfo();
void* QBluetoothHostInfo_NewQBluetoothHostInfo2(void* other);
char* QBluetoothHostInfo_Name(void* ptr);
void QBluetoothHostInfo_SetAddress(void* ptr, void* address);
void QBluetoothHostInfo_SetName(void* ptr, char* name);
void QBluetoothHostInfo_DestroyQBluetoothHostInfo(void* ptr);

#ifdef __cplusplus
}
#endif