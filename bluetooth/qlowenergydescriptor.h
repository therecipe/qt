#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QLowEnergyDescriptor_NewQLowEnergyDescriptor();
QtObjectPtr QLowEnergyDescriptor_NewQLowEnergyDescriptor2(QtObjectPtr other);
int QLowEnergyDescriptor_IsValid(QtObjectPtr ptr);
char* QLowEnergyDescriptor_Name(QtObjectPtr ptr);
int QLowEnergyDescriptor_Type(QtObjectPtr ptr);
void QLowEnergyDescriptor_DestroyQLowEnergyDescriptor(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif