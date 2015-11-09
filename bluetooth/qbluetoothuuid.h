#ifdef __cplusplus
extern "C" {
#endif

void* QBluetoothUuid_NewQBluetoothUuid();
void* QBluetoothUuid_NewQBluetoothUuid4(int uuid);
void* QBluetoothUuid_NewQBluetoothUuid5(int uuid);
void* QBluetoothUuid_NewQBluetoothUuid2(int uuid);
void* QBluetoothUuid_NewQBluetoothUuid3(int uuid);
void* QBluetoothUuid_NewQBluetoothUuid10(void* uuid);
void* QBluetoothUuid_NewQBluetoothUuid9(char* uuid);
void* QBluetoothUuid_NewQBluetoothUuid11(void* uuid);
char* QBluetoothUuid_QBluetoothUuid_CharacteristicToString(int uuid);
char* QBluetoothUuid_QBluetoothUuid_DescriptorToString(int uuid);
int QBluetoothUuid_MinimumSize(void* ptr);
char* QBluetoothUuid_QBluetoothUuid_ProtocolToString(int uuid);
char* QBluetoothUuid_QBluetoothUuid_ServiceClassToString(int uuid);
void QBluetoothUuid_DestroyQBluetoothUuid(void* ptr);

#ifdef __cplusplus
}
#endif