#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QGeoCircle_NewQGeoCircle();
QtObjectPtr QGeoCircle_NewQGeoCircle3(QtObjectPtr other);
QtObjectPtr QGeoCircle_NewQGeoCircle4(QtObjectPtr other);
void QGeoCircle_SetCenter(QtObjectPtr ptr, QtObjectPtr center);
char* QGeoCircle_ToString(QtObjectPtr ptr);
void QGeoCircle_DestroyQGeoCircle(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif