#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QSignalBlocker_Reblock(QtObjectPtr ptr);
void QSignalBlocker_Unblock(QtObjectPtr ptr);
void QSignalBlocker_DestroyQSignalBlocker(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif