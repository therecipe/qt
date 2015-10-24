#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QDomImplementation_NewQDomImplementation();
QtObjectPtr QDomImplementation_NewQDomImplementation2(QtObjectPtr x);
int QDomImplementation_HasFeature(QtObjectPtr ptr, char* feature, char* version);
int QDomImplementation_QDomImplementation_InvalidDataPolicy();
int QDomImplementation_IsNull(QtObjectPtr ptr);
void QDomImplementation_QDomImplementation_SetInvalidDataPolicy(int policy);
void QDomImplementation_DestroyQDomImplementation(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif