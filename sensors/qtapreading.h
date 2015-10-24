#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QTapReading_IsDoubleTap(QtObjectPtr ptr);
int QTapReading_TapDirection(QtObjectPtr ptr);
void QTapReading_SetDoubleTap(QtObjectPtr ptr, int doubleTap);
void QTapReading_SetTapDirection(QtObjectPtr ptr, int tapDirection);

#ifdef __cplusplus
}
#endif