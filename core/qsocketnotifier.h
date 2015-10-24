#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QSocketNotifier_ConnectActivated(QtObjectPtr ptr);
void QSocketNotifier_DisconnectActivated(QtObjectPtr ptr);
int QSocketNotifier_IsEnabled(QtObjectPtr ptr);
void QSocketNotifier_SetEnabled(QtObjectPtr ptr, int enable);
int QSocketNotifier_Type(QtObjectPtr ptr);
void QSocketNotifier_DestroyQSocketNotifier(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif