#ifdef __cplusplus
extern "C" {
#endif

void* QSequentialAnimationGroup_CurrentAnimation(void* ptr);
void* QSequentialAnimationGroup_AddPause(void* ptr, int msecs);
void QSequentialAnimationGroup_ConnectCurrentAnimationChanged(void* ptr);
void QSequentialAnimationGroup_DisconnectCurrentAnimationChanged(void* ptr);
int QSequentialAnimationGroup_Duration(void* ptr);
void* QSequentialAnimationGroup_InsertPause(void* ptr, int index, int msecs);
void QSequentialAnimationGroup_DestroyQSequentialAnimationGroup(void* ptr);

#ifdef __cplusplus
}
#endif