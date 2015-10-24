#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QGraphicsRotation_SetAxis2(QtObjectPtr ptr, int axis);
void QGraphicsRotation_SetAxis(QtObjectPtr ptr, QtObjectPtr axis);
void QGraphicsRotation_SetOrigin(QtObjectPtr ptr, QtObjectPtr point);
QtObjectPtr QGraphicsRotation_NewQGraphicsRotation(QtObjectPtr parent);
void QGraphicsRotation_ConnectAngleChanged(QtObjectPtr ptr);
void QGraphicsRotation_DisconnectAngleChanged(QtObjectPtr ptr);
void QGraphicsRotation_ApplyTo(QtObjectPtr ptr, QtObjectPtr matrix);
void QGraphicsRotation_ConnectAxisChanged(QtObjectPtr ptr);
void QGraphicsRotation_DisconnectAxisChanged(QtObjectPtr ptr);
void QGraphicsRotation_ConnectOriginChanged(QtObjectPtr ptr);
void QGraphicsRotation_DisconnectOriginChanged(QtObjectPtr ptr);
void QGraphicsRotation_DestroyQGraphicsRotation(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif