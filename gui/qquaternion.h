#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QQuaternion_NewQQuaternion();
QtObjectPtr QQuaternion_NewQQuaternion5(QtObjectPtr vector);
void QQuaternion_GetAxes(QtObjectPtr ptr, QtObjectPtr xAxis, QtObjectPtr yAxis, QtObjectPtr zAxis);
int QQuaternion_IsIdentity(QtObjectPtr ptr);
int QQuaternion_IsNull(QtObjectPtr ptr);
void QQuaternion_Normalize(QtObjectPtr ptr);
void QQuaternion_SetVector(QtObjectPtr ptr, QtObjectPtr vector);

#ifdef __cplusplus
}
#endif