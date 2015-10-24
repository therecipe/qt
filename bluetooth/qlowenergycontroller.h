#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QLowEnergyController_ConnectConnected(QtObjectPtr ptr);
void QLowEnergyController_DisconnectConnected(QtObjectPtr ptr);
void QLowEnergyController_ConnectDisconnected(QtObjectPtr ptr);
void QLowEnergyController_DisconnectDisconnected(QtObjectPtr ptr);
void QLowEnergyController_ConnectDiscoveryFinished(QtObjectPtr ptr);
void QLowEnergyController_DisconnectDiscoveryFinished(QtObjectPtr ptr);
void QLowEnergyController_ConnectStateChanged(QtObjectPtr ptr);
void QLowEnergyController_DisconnectStateChanged(QtObjectPtr ptr);
QtObjectPtr QLowEnergyController_NewQLowEnergyController(QtObjectPtr remoteDeviceInfo, QtObjectPtr parent);
void QLowEnergyController_ConnectToDevice(QtObjectPtr ptr);
QtObjectPtr QLowEnergyController_CreateServiceObject(QtObjectPtr ptr, QtObjectPtr serviceUuid, QtObjectPtr parent);
void QLowEnergyController_DisconnectFromDevice(QtObjectPtr ptr);
void QLowEnergyController_DiscoverServices(QtObjectPtr ptr);
int QLowEnergyController_Error(QtObjectPtr ptr);
char* QLowEnergyController_ErrorString(QtObjectPtr ptr);
int QLowEnergyController_RemoteAddressType(QtObjectPtr ptr);
char* QLowEnergyController_RemoteName(QtObjectPtr ptr);
void QLowEnergyController_SetRemoteAddressType(QtObjectPtr ptr, int ty);
void QLowEnergyController_DestroyQLowEnergyController(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif