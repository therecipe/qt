#ifdef __cplusplus
extern "C" {
#endif

void* QWheelEvent_NewQWheelEvent(void* pos, void* globalPos, void* pixelDelta, void* angleDelta, int qt4Delta, int qt4Orientation, int buttons, int modifiers);
void* QWheelEvent_NewQWheelEvent4(void* pos, void* globalPos, void* pixelDelta, void* angleDelta, int qt4Delta, int qt4Orientation, int buttons, int modifiers, int phase);
void* QWheelEvent_NewQWheelEvent5(void* pos, void* globalPos, void* pixelDelta, void* angleDelta, int qt4Delta, int qt4Orientation, int buttons, int modifiers, int phase, int source);
int QWheelEvent_Buttons(void* ptr);
int QWheelEvent_GlobalX(void* ptr);
int QWheelEvent_GlobalY(void* ptr);
int QWheelEvent_Phase(void* ptr);
int QWheelEvent_Source(void* ptr);
int QWheelEvent_X(void* ptr);
int QWheelEvent_Y(void* ptr);

#ifdef __cplusplus
}
#endif