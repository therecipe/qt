#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QAbstractVideoSurface_ConnectActiveChanged(QtObjectPtr ptr);
void QAbstractVideoSurface_DisconnectActiveChanged(QtObjectPtr ptr);
int QAbstractVideoSurface_Error(QtObjectPtr ptr);
int QAbstractVideoSurface_IsActive(QtObjectPtr ptr);
int QAbstractVideoSurface_IsFormatSupported(QtObjectPtr ptr, QtObjectPtr format);
int QAbstractVideoSurface_Present(QtObjectPtr ptr, QtObjectPtr frame);
int QAbstractVideoSurface_Start(QtObjectPtr ptr, QtObjectPtr format);
void QAbstractVideoSurface_Stop(QtObjectPtr ptr);
void QAbstractVideoSurface_ConnectSupportedFormatsChanged(QtObjectPtr ptr);
void QAbstractVideoSurface_DisconnectSupportedFormatsChanged(QtObjectPtr ptr);
void QAbstractVideoSurface_DestroyQAbstractVideoSurface(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif