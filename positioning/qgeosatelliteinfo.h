#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QGeoSatelliteInfo_NewQGeoSatelliteInfo();
QtObjectPtr QGeoSatelliteInfo_NewQGeoSatelliteInfo2(QtObjectPtr other);
int QGeoSatelliteInfo_HasAttribute(QtObjectPtr ptr, int attribute);
void QGeoSatelliteInfo_RemoveAttribute(QtObjectPtr ptr, int attribute);
int QGeoSatelliteInfo_SatelliteIdentifier(QtObjectPtr ptr);
int QGeoSatelliteInfo_SatelliteSystem(QtObjectPtr ptr);
void QGeoSatelliteInfo_SetSatelliteIdentifier(QtObjectPtr ptr, int satId);
void QGeoSatelliteInfo_SetSatelliteSystem(QtObjectPtr ptr, int system);
void QGeoSatelliteInfo_SetSignalStrength(QtObjectPtr ptr, int signalStrength);
int QGeoSatelliteInfo_SignalStrength(QtObjectPtr ptr);
void QGeoSatelliteInfo_DestroyQGeoSatelliteInfo(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif