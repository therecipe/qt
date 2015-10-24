#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QGeoRectangle_NewQGeoRectangle();
QtObjectPtr QGeoRectangle_NewQGeoRectangle3(QtObjectPtr topLeft, QtObjectPtr bottomRight);
QtObjectPtr QGeoRectangle_NewQGeoRectangle5(QtObjectPtr other);
QtObjectPtr QGeoRectangle_NewQGeoRectangle6(QtObjectPtr other);
int QGeoRectangle_Contains(QtObjectPtr ptr, QtObjectPtr rectangle);
int QGeoRectangle_Intersects(QtObjectPtr ptr, QtObjectPtr rectangle);
void QGeoRectangle_SetBottomLeft(QtObjectPtr ptr, QtObjectPtr bottomLeft);
void QGeoRectangle_SetBottomRight(QtObjectPtr ptr, QtObjectPtr bottomRight);
void QGeoRectangle_SetCenter(QtObjectPtr ptr, QtObjectPtr center);
void QGeoRectangle_SetTopLeft(QtObjectPtr ptr, QtObjectPtr topLeft);
void QGeoRectangle_SetTopRight(QtObjectPtr ptr, QtObjectPtr topRight);
char* QGeoRectangle_ToString(QtObjectPtr ptr);
void QGeoRectangle_DestroyQGeoRectangle(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif