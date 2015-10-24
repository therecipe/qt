#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QGradient_CoordinateMode(QtObjectPtr ptr);
void QGradient_SetCoordinateMode(QtObjectPtr ptr, int mode);
void QGradient_SetSpread(QtObjectPtr ptr, int method);
int QGradient_Spread(QtObjectPtr ptr);
int QGradient_Type(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif