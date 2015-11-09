#ifdef __cplusplus
extern "C" {
#endif

void* QLowEnergyDescriptor_NewQLowEnergyDescriptor();
void* QLowEnergyDescriptor_NewQLowEnergyDescriptor2(void* other);
int QLowEnergyDescriptor_IsValid(void* ptr);
char* QLowEnergyDescriptor_Name(void* ptr);
int QLowEnergyDescriptor_Type(void* ptr);
void* QLowEnergyDescriptor_Value(void* ptr);
void QLowEnergyDescriptor_DestroyQLowEnergyDescriptor(void* ptr);

#ifdef __cplusplus
}
#endif