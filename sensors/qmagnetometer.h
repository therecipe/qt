#ifdef __cplusplus
extern "C" {
#endif

void* QMagnetometer_Reading(void* ptr);
int QMagnetometer_ReturnGeoValues(void* ptr);
void QMagnetometer_SetReturnGeoValues(void* ptr, int returnGeoValues);
void* QMagnetometer_NewQMagnetometer(void* parent);
void QMagnetometer_ConnectReturnGeoValuesChanged(void* ptr);
void QMagnetometer_DisconnectReturnGeoValuesChanged(void* ptr);
void QMagnetometer_DestroyQMagnetometer(void* ptr);

#ifdef __cplusplus
}
#endif