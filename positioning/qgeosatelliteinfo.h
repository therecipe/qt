#ifdef __cplusplus
extern "C" {
#endif

void* QGeoSatelliteInfo_NewQGeoSatelliteInfo();
void* QGeoSatelliteInfo_NewQGeoSatelliteInfo2(void* other);
double QGeoSatelliteInfo_Attribute(void* ptr, int attribute);
int QGeoSatelliteInfo_HasAttribute(void* ptr, int attribute);
void QGeoSatelliteInfo_RemoveAttribute(void* ptr, int attribute);
int QGeoSatelliteInfo_SatelliteIdentifier(void* ptr);
int QGeoSatelliteInfo_SatelliteSystem(void* ptr);
void QGeoSatelliteInfo_SetAttribute(void* ptr, int attribute, double value);
void QGeoSatelliteInfo_SetSatelliteIdentifier(void* ptr, int satId);
void QGeoSatelliteInfo_SetSatelliteSystem(void* ptr, int system);
void QGeoSatelliteInfo_SetSignalStrength(void* ptr, int signalStrength);
int QGeoSatelliteInfo_SignalStrength(void* ptr);
void QGeoSatelliteInfo_DestroyQGeoSatelliteInfo(void* ptr);

#ifdef __cplusplus
}
#endif