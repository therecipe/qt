#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QGeoShape_NewQGeoShape();
QtObjectPtr QGeoShape_NewQGeoShape2(QtObjectPtr other);
int QGeoShape_Contains(QtObjectPtr ptr, QtObjectPtr coordinate);
void QGeoShape_ExtendShape(QtObjectPtr ptr, QtObjectPtr coordinate);
int QGeoShape_IsEmpty(QtObjectPtr ptr);
int QGeoShape_IsValid(QtObjectPtr ptr);
char* QGeoShape_ToString(QtObjectPtr ptr);
int QGeoShape_Type(QtObjectPtr ptr);
void QGeoShape_DestroyQGeoShape(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif