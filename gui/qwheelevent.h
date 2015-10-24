#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QWheelEvent_NewQWheelEvent(QtObjectPtr pos, QtObjectPtr globalPos, QtObjectPtr pixelDelta, QtObjectPtr angleDelta, int qt4Delta, int qt4Orientation, int buttons, int modifiers);
QtObjectPtr QWheelEvent_NewQWheelEvent4(QtObjectPtr pos, QtObjectPtr globalPos, QtObjectPtr pixelDelta, QtObjectPtr angleDelta, int qt4Delta, int qt4Orientation, int buttons, int modifiers, int phase);
QtObjectPtr QWheelEvent_NewQWheelEvent5(QtObjectPtr pos, QtObjectPtr globalPos, QtObjectPtr pixelDelta, QtObjectPtr angleDelta, int qt4Delta, int qt4Orientation, int buttons, int modifiers, int phase, int source);
int QWheelEvent_Buttons(QtObjectPtr ptr);
int QWheelEvent_GlobalX(QtObjectPtr ptr);
int QWheelEvent_GlobalY(QtObjectPtr ptr);
int QWheelEvent_Phase(QtObjectPtr ptr);
int QWheelEvent_Source(QtObjectPtr ptr);
int QWheelEvent_X(QtObjectPtr ptr);
int QWheelEvent_Y(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif