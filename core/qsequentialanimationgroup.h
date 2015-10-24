#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QSequentialAnimationGroup_CurrentAnimation(QtObjectPtr ptr);
QtObjectPtr QSequentialAnimationGroup_AddPause(QtObjectPtr ptr, int msecs);
void QSequentialAnimationGroup_ConnectCurrentAnimationChanged(QtObjectPtr ptr);
void QSequentialAnimationGroup_DisconnectCurrentAnimationChanged(QtObjectPtr ptr);
int QSequentialAnimationGroup_Duration(QtObjectPtr ptr);
QtObjectPtr QSequentialAnimationGroup_InsertPause(QtObjectPtr ptr, int index, int msecs);
void QSequentialAnimationGroup_DestroyQSequentialAnimationGroup(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif