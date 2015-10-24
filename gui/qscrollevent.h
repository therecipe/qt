#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QScrollEvent_NewQScrollEvent(QtObjectPtr contentPos, QtObjectPtr overshootDistance, int scrollState);
int QScrollEvent_ScrollState(QtObjectPtr ptr);
void QScrollEvent_DestroyQScrollEvent(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif