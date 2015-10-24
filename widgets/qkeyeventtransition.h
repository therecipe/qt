#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QKeyEventTransition_NewQKeyEventTransition2(QtObjectPtr object, int ty, int key, QtObjectPtr sourceState);
QtObjectPtr QKeyEventTransition_NewQKeyEventTransition(QtObjectPtr sourceState);
int QKeyEventTransition_Key(QtObjectPtr ptr);
int QKeyEventTransition_ModifierMask(QtObjectPtr ptr);
void QKeyEventTransition_SetKey(QtObjectPtr ptr, int key);
void QKeyEventTransition_SetModifierMask(QtObjectPtr ptr, int modifierMask);
void QKeyEventTransition_DestroyQKeyEventTransition(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif