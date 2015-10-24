#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QLinearGradient_NewQLinearGradient();
QtObjectPtr QLinearGradient_NewQLinearGradient2(QtObjectPtr start, QtObjectPtr finalStop);
void QLinearGradient_SetFinalStop(QtObjectPtr ptr, QtObjectPtr stop);
void QLinearGradient_SetStart(QtObjectPtr ptr, QtObjectPtr start);

#ifdef __cplusplus
}
#endif