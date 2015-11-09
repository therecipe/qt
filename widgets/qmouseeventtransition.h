#ifdef __cplusplus
extern "C" {
#endif

void* QMouseEventTransition_NewQMouseEventTransition2(void* object, int ty, int button, void* sourceState);
void* QMouseEventTransition_NewQMouseEventTransition(void* sourceState);
int QMouseEventTransition_Button(void* ptr);
int QMouseEventTransition_ModifierMask(void* ptr);
void QMouseEventTransition_SetButton(void* ptr, int button);
void QMouseEventTransition_SetHitTestPath(void* ptr, void* path);
void QMouseEventTransition_SetModifierMask(void* ptr, int modifierMask);
void QMouseEventTransition_DestroyQMouseEventTransition(void* ptr);

#ifdef __cplusplus
}
#endif