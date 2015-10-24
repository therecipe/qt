#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QBluetoothUuid_NewQBluetoothUuid();
QtObjectPtr QBluetoothUuid_NewQBluetoothUuid4(int uuid);
QtObjectPtr QBluetoothUuid_NewQBluetoothUuid5(int uuid);
QtObjectPtr QBluetoothUuid_NewQBluetoothUuid2(int uuid);
QtObjectPtr QBluetoothUuid_NewQBluetoothUuid3(int uuid);
QtObjectPtr QBluetoothUuid_NewQBluetoothUuid10(QtObjectPtr uuid);
QtObjectPtr QBluetoothUuid_NewQBluetoothUuid9(char* uuid);
QtObjectPtr QBluetoothUuid_NewQBluetoothUuid11(QtObjectPtr uuid);
char* QBluetoothUuid_QBluetoothUuid_CharacteristicToString(int uuid);
char* QBluetoothUuid_QBluetoothUuid_DescriptorToString(int uuid);
int QBluetoothUuid_MinimumSize(QtObjectPtr ptr);
char* QBluetoothUuid_QBluetoothUuid_ProtocolToString(int uuid);
char* QBluetoothUuid_QBluetoothUuid_ServiceClassToString(int uuid);
void QBluetoothUuid_DestroyQBluetoothUuid(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif