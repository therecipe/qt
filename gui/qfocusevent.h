#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QFocusEvent_NewQFocusEvent(int ty, int reason);
int QFocusEvent_GotFocus(QtObjectPtr ptr);
int QFocusEvent_LostFocus(QtObjectPtr ptr);
int QFocusEvent_Reason(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif