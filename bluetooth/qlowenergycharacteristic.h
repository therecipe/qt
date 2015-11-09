#ifdef __cplusplus
extern "C" {
#endif

void* QLowEnergyCharacteristic_NewQLowEnergyCharacteristic();
void* QLowEnergyCharacteristic_NewQLowEnergyCharacteristic2(void* other);
int QLowEnergyCharacteristic_IsValid(void* ptr);
char* QLowEnergyCharacteristic_Name(void* ptr);
int QLowEnergyCharacteristic_Properties(void* ptr);
void* QLowEnergyCharacteristic_Value(void* ptr);
void QLowEnergyCharacteristic_DestroyQLowEnergyCharacteristic(void* ptr);

#ifdef __cplusplus
}
#endif