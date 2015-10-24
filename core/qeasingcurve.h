#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QEasingCurve_NewQEasingCurve3(QtObjectPtr other);
QtObjectPtr QEasingCurve_NewQEasingCurve(int ty);
QtObjectPtr QEasingCurve_NewQEasingCurve2(QtObjectPtr other);
void QEasingCurve_AddCubicBezierSegment(QtObjectPtr ptr, QtObjectPtr c1, QtObjectPtr c2, QtObjectPtr endPoint);
void QEasingCurve_SetType(QtObjectPtr ptr, int ty);
void QEasingCurve_Swap(QtObjectPtr ptr, QtObjectPtr other);
int QEasingCurve_Type(QtObjectPtr ptr);
void QEasingCurve_DestroyQEasingCurve(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif