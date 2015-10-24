#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QAbstractVideoFilter_IsActive(QtObjectPtr ptr);
void QAbstractVideoFilter_SetActive(QtObjectPtr ptr, int v);
void QAbstractVideoFilter_ConnectActiveChanged(QtObjectPtr ptr);
void QAbstractVideoFilter_DisconnectActiveChanged(QtObjectPtr ptr);
QtObjectPtr QAbstractVideoFilter_CreateFilterRunnable(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif