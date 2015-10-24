#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QMouseEvent_NewQMouseEvent(int ty, QtObjectPtr localPos, int button, int buttons, int modifiers);
QtObjectPtr QMouseEvent_NewQMouseEvent2(int ty, QtObjectPtr localPos, QtObjectPtr screenPos, int button, int buttons, int modifiers);
QtObjectPtr QMouseEvent_NewQMouseEvent3(int ty, QtObjectPtr localPos, QtObjectPtr windowPos, QtObjectPtr screenPos, int button, int buttons, int modifiers);
int QMouseEvent_Button(QtObjectPtr ptr);
int QMouseEvent_Buttons(QtObjectPtr ptr);
int QMouseEvent_Flags(QtObjectPtr ptr);
int QMouseEvent_GlobalX(QtObjectPtr ptr);
int QMouseEvent_GlobalY(QtObjectPtr ptr);
int QMouseEvent_Source(QtObjectPtr ptr);
int QMouseEvent_X(QtObjectPtr ptr);
int QMouseEvent_Y(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif