#ifdef __cplusplus
extern "C" {
#endif

void QLowEnergyController_ConnectConnected(void* ptr);
void QLowEnergyController_DisconnectConnected(void* ptr);
void QLowEnergyController_ConnectDisconnected(void* ptr);
void QLowEnergyController_DisconnectDisconnected(void* ptr);
void QLowEnergyController_ConnectDiscoveryFinished(void* ptr);
void QLowEnergyController_DisconnectDiscoveryFinished(void* ptr);
void QLowEnergyController_ConnectStateChanged(void* ptr);
void QLowEnergyController_DisconnectStateChanged(void* ptr);
void* QLowEnergyController_NewQLowEnergyController(void* remoteDeviceInfo, void* parent);
void QLowEnergyController_ConnectToDevice(void* ptr);
void* QLowEnergyController_CreateServiceObject(void* ptr, void* serviceUuid, void* parent);
void QLowEnergyController_DisconnectFromDevice(void* ptr);
void QLowEnergyController_DiscoverServices(void* ptr);
int QLowEnergyController_Error(void* ptr);
char* QLowEnergyController_ErrorString(void* ptr);
int QLowEnergyController_RemoteAddressType(void* ptr);
char* QLowEnergyController_RemoteName(void* ptr);
void QLowEnergyController_SetRemoteAddressType(void* ptr, int ty);
void QLowEnergyController_DestroyQLowEnergyController(void* ptr);

#ifdef __cplusplus
}
#endif