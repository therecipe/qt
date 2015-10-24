#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QGraphicsScale_SetOrigin(QtObjectPtr ptr, QtObjectPtr point);
QtObjectPtr QGraphicsScale_NewQGraphicsScale(QtObjectPtr parent);
void QGraphicsScale_ApplyTo(QtObjectPtr ptr, QtObjectPtr matrix);
void QGraphicsScale_ConnectOriginChanged(QtObjectPtr ptr);
void QGraphicsScale_DisconnectOriginChanged(QtObjectPtr ptr);
void QGraphicsScale_ConnectScaleChanged(QtObjectPtr ptr);
void QGraphicsScale_DisconnectScaleChanged(QtObjectPtr ptr);
void QGraphicsScale_ConnectXScaleChanged(QtObjectPtr ptr);
void QGraphicsScale_DisconnectXScaleChanged(QtObjectPtr ptr);
void QGraphicsScale_ConnectYScaleChanged(QtObjectPtr ptr);
void QGraphicsScale_DisconnectYScaleChanged(QtObjectPtr ptr);
void QGraphicsScale_ConnectZScaleChanged(QtObjectPtr ptr);
void QGraphicsScale_DisconnectZScaleChanged(QtObjectPtr ptr);
void QGraphicsScale_DestroyQGraphicsScale(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif