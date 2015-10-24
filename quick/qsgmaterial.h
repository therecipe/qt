#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QSGMaterial_Compare(QtObjectPtr ptr, QtObjectPtr other);
QtObjectPtr QSGMaterial_CreateShader(QtObjectPtr ptr);
int QSGMaterial_Flags(QtObjectPtr ptr);
void QSGMaterial_SetFlag(QtObjectPtr ptr, int flags, int on);
QtObjectPtr QSGMaterial_Type(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif