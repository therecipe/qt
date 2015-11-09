#ifdef __cplusplus
extern "C" {
#endif

void* QGeoCoordinate_NewQGeoCoordinate();
void* QGeoCoordinate_NewQGeoCoordinate4(void* other);
double QGeoCoordinate_AzimuthTo(void* ptr, void* other);
double QGeoCoordinate_DistanceTo(void* ptr, void* other);
int QGeoCoordinate_IsValid(void* ptr);
char* QGeoCoordinate_ToString(void* ptr, int format);
int QGeoCoordinate_Type(void* ptr);
void QGeoCoordinate_DestroyQGeoCoordinate(void* ptr);

#ifdef __cplusplus
}
#endif