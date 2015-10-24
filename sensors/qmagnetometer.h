#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QMagnetometer_Reading(QtObjectPtr ptr);
int QMagnetometer_ReturnGeoValues(QtObjectPtr ptr);
void QMagnetometer_SetReturnGeoValues(QtObjectPtr ptr, int returnGeoValues);
QtObjectPtr QMagnetometer_NewQMagnetometer(QtObjectPtr parent);
void QMagnetometer_ConnectReturnGeoValuesChanged(QtObjectPtr ptr);
void QMagnetometer_DisconnectReturnGeoValuesChanged(QtObjectPtr ptr);
void QMagnetometer_DestroyQMagnetometer(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif