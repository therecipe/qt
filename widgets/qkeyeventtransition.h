#ifdef __cplusplus
extern "C" {
#endif

void* QKeyEventTransition_NewQKeyEventTransition2(void* object, int ty, int key, void* sourceState);
void* QKeyEventTransition_NewQKeyEventTransition(void* sourceState);
int QKeyEventTransition_Key(void* ptr);
int QKeyEventTransition_ModifierMask(void* ptr);
void QKeyEventTransition_SetKey(void* ptr, int key);
void QKeyEventTransition_SetModifierMask(void* ptr, int modifierMask);
void QKeyEventTransition_DestroyQKeyEventTransition(void* ptr);

#ifdef __cplusplus
}
#endif