#ifdef __cplusplus
extern "C" {
#endif

void* QGeoRectangle_NewQGeoRectangle();
void* QGeoRectangle_NewQGeoRectangle3(void* topLeft, void* bottomRight);
void* QGeoRectangle_NewQGeoRectangle5(void* other);
void* QGeoRectangle_NewQGeoRectangle6(void* other);
int QGeoRectangle_Contains(void* ptr, void* rectangle);
int QGeoRectangle_Intersects(void* ptr, void* rectangle);
void QGeoRectangle_SetBottomLeft(void* ptr, void* bottomLeft);
void QGeoRectangle_SetBottomRight(void* ptr, void* bottomRight);
void QGeoRectangle_SetCenter(void* ptr, void* center);
void QGeoRectangle_SetTopLeft(void* ptr, void* topLeft);
void QGeoRectangle_SetTopRight(void* ptr, void* topRight);
char* QGeoRectangle_ToString(void* ptr);
void QGeoRectangle_DestroyQGeoRectangle(void* ptr);

#ifdef __cplusplus
}
#endif