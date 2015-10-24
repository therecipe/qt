#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QTouchEvent_Device(QtObjectPtr ptr);
QtObjectPtr QTouchEvent_Target(QtObjectPtr ptr);
int QTouchEvent_TouchPointStates(QtObjectPtr ptr);
QtObjectPtr QTouchEvent_Window(QtObjectPtr ptr);
void QTouchEvent_DestroyQTouchEvent(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif