#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QSGAbstractRenderer_ClearMode(QtObjectPtr ptr);
void QSGAbstractRenderer_ConnectSceneGraphChanged(QtObjectPtr ptr);
void QSGAbstractRenderer_DisconnectSceneGraphChanged(QtObjectPtr ptr);
void QSGAbstractRenderer_SetClearColor(QtObjectPtr ptr, QtObjectPtr color);
void QSGAbstractRenderer_SetClearMode(QtObjectPtr ptr, int mode);
void QSGAbstractRenderer_SetDeviceRect(QtObjectPtr ptr, QtObjectPtr rect);
void QSGAbstractRenderer_SetDeviceRect2(QtObjectPtr ptr, QtObjectPtr size);
void QSGAbstractRenderer_SetProjectionMatrix(QtObjectPtr ptr, QtObjectPtr matrix);
void QSGAbstractRenderer_SetProjectionMatrixToRect(QtObjectPtr ptr, QtObjectPtr rect);
void QSGAbstractRenderer_SetViewportRect(QtObjectPtr ptr, QtObjectPtr rect);
void QSGAbstractRenderer_SetViewportRect2(QtObjectPtr ptr, QtObjectPtr size);

#ifdef __cplusplus
}
#endif