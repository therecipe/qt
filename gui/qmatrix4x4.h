#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QMatrix4x4_NewQMatrix4x4();
QtObjectPtr QMatrix4x4_NewQMatrix4x47(QtObjectPtr transform);
int QMatrix4x4_IsAffine(QtObjectPtr ptr);
int QMatrix4x4_IsIdentity(QtObjectPtr ptr);
void QMatrix4x4_LookAt(QtObjectPtr ptr, QtObjectPtr eye, QtObjectPtr center, QtObjectPtr up);
void QMatrix4x4_Optimize(QtObjectPtr ptr);
void QMatrix4x4_Ortho2(QtObjectPtr ptr, QtObjectPtr rect);
void QMatrix4x4_Ortho3(QtObjectPtr ptr, QtObjectPtr rect);
void QMatrix4x4_Rotate2(QtObjectPtr ptr, QtObjectPtr quaternion);
void QMatrix4x4_Scale(QtObjectPtr ptr, QtObjectPtr vector);
void QMatrix4x4_SetColumn(QtObjectPtr ptr, int index, QtObjectPtr value);
void QMatrix4x4_SetRow(QtObjectPtr ptr, int index, QtObjectPtr value);
void QMatrix4x4_SetToIdentity(QtObjectPtr ptr);
void QMatrix4x4_Translate(QtObjectPtr ptr, QtObjectPtr vector);
void QMatrix4x4_Viewport2(QtObjectPtr ptr, QtObjectPtr rect);

#ifdef __cplusplus
}
#endif