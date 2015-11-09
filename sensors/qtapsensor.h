#ifdef __cplusplus
extern "C" {
#endif

void* QTapSensor_Reading(void* ptr);
int QTapSensor_ReturnDoubleTapEvents(void* ptr);
void QTapSensor_SetReturnDoubleTapEvents(void* ptr, int returnDoubleTapEvents);
void* QTapSensor_NewQTapSensor(void* parent);
void QTapSensor_ConnectReturnDoubleTapEventsChanged(void* ptr);
void QTapSensor_DisconnectReturnDoubleTapEventsChanged(void* ptr);
void QTapSensor_DestroyQTapSensor(void* ptr);

#ifdef __cplusplus
}
#endif