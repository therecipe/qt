#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QGeoCoordinate_NewQGeoCoordinate();
QtObjectPtr QGeoCoordinate_NewQGeoCoordinate4(QtObjectPtr other);
int QGeoCoordinate_IsValid(QtObjectPtr ptr);
char* QGeoCoordinate_ToString(QtObjectPtr ptr, int format);
int QGeoCoordinate_Type(QtObjectPtr ptr);
void QGeoCoordinate_DestroyQGeoCoordinate(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif