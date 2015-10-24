#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QTiltSensor_NewQTiltSensor(QtObjectPtr parent);
QtObjectPtr QTiltSensor_Reading(QtObjectPtr ptr);
void QTiltSensor_DestroyQTiltSensor(QtObjectPtr ptr);
void QTiltSensor_Calibrate(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif