#ifdef __cplusplus
extern "C" {
#endif

void* QGeoShape_NewQGeoShape();
void* QGeoShape_NewQGeoShape2(void* other);
int QGeoShape_Contains(void* ptr, void* coordinate);
void QGeoShape_ExtendShape(void* ptr, void* coordinate);
int QGeoShape_IsEmpty(void* ptr);
int QGeoShape_IsValid(void* ptr);
char* QGeoShape_ToString(void* ptr);
int QGeoShape_Type(void* ptr);
void QGeoShape_DestroyQGeoShape(void* ptr);

#ifdef __cplusplus
}
#endif