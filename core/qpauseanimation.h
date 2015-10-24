#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QPauseAnimation_Duration(QtObjectPtr ptr);
void QPauseAnimation_SetDuration(QtObjectPtr ptr, int msecs);
QtObjectPtr QPauseAnimation_NewQPauseAnimation(QtObjectPtr parent);
QtObjectPtr QPauseAnimation_NewQPauseAnimation2(int msecs, QtObjectPtr parent);
void QPauseAnimation_DestroyQPauseAnimation(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif