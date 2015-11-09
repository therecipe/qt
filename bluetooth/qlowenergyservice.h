#ifdef __cplusplus
extern "C" {
#endif

void QLowEnergyService_ConnectStateChanged(void* ptr);
void QLowEnergyService_DisconnectStateChanged(void* ptr);
int QLowEnergyService_Contains(void* ptr, void* characteristic);
int QLowEnergyService_Contains2(void* ptr, void* descriptor);
void QLowEnergyService_DiscoverDetails(void* ptr);
int QLowEnergyService_Error(void* ptr);
void QLowEnergyService_ReadCharacteristic(void* ptr, void* characteristic);
void QLowEnergyService_ReadDescriptor(void* ptr, void* descriptor);
char* QLowEnergyService_ServiceName(void* ptr);
int QLowEnergyService_Type(void* ptr);
void QLowEnergyService_WriteCharacteristic(void* ptr, void* characteristic, void* newValue, int mode);
void QLowEnergyService_WriteDescriptor(void* ptr, void* descriptor, void* newValue);
void QLowEnergyService_DestroyQLowEnergyService(void* ptr);

#ifdef __cplusplus
}
#endif