#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QEventLoopLocker_NewQEventLoopLocker();
QtObjectPtr QEventLoopLocker_NewQEventLoopLocker2(QtObjectPtr loop);
QtObjectPtr QEventLoopLocker_NewQEventLoopLocker3(QtObjectPtr thread);
void QEventLoopLocker_DestroyQEventLoopLocker(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif