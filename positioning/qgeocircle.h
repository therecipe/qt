#ifdef __cplusplus
extern "C" {
#endif

void* QGeoCircle_NewQGeoCircle();
void* QGeoCircle_NewQGeoCircle3(void* other);
void* QGeoCircle_NewQGeoCircle2(void* center, double radius);
void* QGeoCircle_NewQGeoCircle4(void* other);
double QGeoCircle_Radius(void* ptr);
void QGeoCircle_SetCenter(void* ptr, void* center);
void QGeoCircle_SetRadius(void* ptr, double radius);
char* QGeoCircle_ToString(void* ptr);
void QGeoCircle_DestroyQGeoCircle(void* ptr);

#ifdef __cplusplus
}
#endif