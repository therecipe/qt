#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QTapSensor_Reading(QtObjectPtr ptr);
int QTapSensor_ReturnDoubleTapEvents(QtObjectPtr ptr);
void QTapSensor_SetReturnDoubleTapEvents(QtObjectPtr ptr, int returnDoubleTapEvents);
QtObjectPtr QTapSensor_NewQTapSensor(QtObjectPtr parent);
void QTapSensor_ConnectReturnDoubleTapEventsChanged(QtObjectPtr ptr);
void QTapSensor_DisconnectReturnDoubleTapEventsChanged(QtObjectPtr ptr);
void QTapSensor_DestroyQTapSensor(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif