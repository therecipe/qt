#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QLowEnergyService_ConnectStateChanged(QtObjectPtr ptr);
void QLowEnergyService_DisconnectStateChanged(QtObjectPtr ptr);
int QLowEnergyService_Contains(QtObjectPtr ptr, QtObjectPtr characteristic);
int QLowEnergyService_Contains2(QtObjectPtr ptr, QtObjectPtr descriptor);
void QLowEnergyService_DiscoverDetails(QtObjectPtr ptr);
int QLowEnergyService_Error(QtObjectPtr ptr);
void QLowEnergyService_ReadCharacteristic(QtObjectPtr ptr, QtObjectPtr characteristic);
void QLowEnergyService_ReadDescriptor(QtObjectPtr ptr, QtObjectPtr descriptor);
char* QLowEnergyService_ServiceName(QtObjectPtr ptr);
int QLowEnergyService_Type(QtObjectPtr ptr);
void QLowEnergyService_WriteCharacteristic(QtObjectPtr ptr, QtObjectPtr characteristic, QtObjectPtr newValue, int mode);
void QLowEnergyService_WriteDescriptor(QtObjectPtr ptr, QtObjectPtr descriptor, QtObjectPtr newValue);
void QLowEnergyService_DestroyQLowEnergyService(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif