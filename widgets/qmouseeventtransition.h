#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QMouseEventTransition_NewQMouseEventTransition2(QtObjectPtr object, int ty, int button, QtObjectPtr sourceState);
QtObjectPtr QMouseEventTransition_NewQMouseEventTransition(QtObjectPtr sourceState);
int QMouseEventTransition_Button(QtObjectPtr ptr);
int QMouseEventTransition_ModifierMask(QtObjectPtr ptr);
void QMouseEventTransition_SetButton(QtObjectPtr ptr, int button);
void QMouseEventTransition_SetHitTestPath(QtObjectPtr ptr, QtObjectPtr path);
void QMouseEventTransition_SetModifierMask(QtObjectPtr ptr, int modifierMask);
void QMouseEventTransition_DestroyQMouseEventTransition(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif