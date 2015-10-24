#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QElapsedTimer_NewQElapsedTimer();
void QElapsedTimer_Invalidate(QtObjectPtr ptr);
int QElapsedTimer_IsValid(QtObjectPtr ptr);
int QElapsedTimer_QElapsedTimer_ClockType();
int QElapsedTimer_QElapsedTimer_IsMonotonic();
void QElapsedTimer_Start(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif