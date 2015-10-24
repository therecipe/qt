#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QWaitCondition_NewQWaitCondition();
void QWaitCondition_WakeAll(QtObjectPtr ptr);
void QWaitCondition_WakeOne(QtObjectPtr ptr);
void QWaitCondition_DestroyQWaitCondition(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif