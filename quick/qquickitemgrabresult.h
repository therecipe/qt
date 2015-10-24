#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

char* QQuickItemGrabResult_Url(QtObjectPtr ptr);
void QQuickItemGrabResult_ConnectReady(QtObjectPtr ptr);
void QQuickItemGrabResult_DisconnectReady(QtObjectPtr ptr);
int QQuickItemGrabResult_SaveToFile(QtObjectPtr ptr, char* fileName);

#ifdef __cplusplus
}
#endif