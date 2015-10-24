#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QPinchGesture_ChangeFlags(QtObjectPtr ptr);
void QPinchGesture_SetCenterPoint(QtObjectPtr ptr, QtObjectPtr value);
void QPinchGesture_SetChangeFlags(QtObjectPtr ptr, int value);
void QPinchGesture_SetLastCenterPoint(QtObjectPtr ptr, QtObjectPtr value);
void QPinchGesture_SetStartCenterPoint(QtObjectPtr ptr, QtObjectPtr value);
void QPinchGesture_SetTotalChangeFlags(QtObjectPtr ptr, int value);
int QPinchGesture_TotalChangeFlags(QtObjectPtr ptr);
void QPinchGesture_DestroyQPinchGesture(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif