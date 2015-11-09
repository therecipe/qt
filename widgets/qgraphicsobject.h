#ifdef __cplusplus
extern "C" {
#endif

void QGraphicsObject_ConnectEnabledChanged(void* ptr);
void QGraphicsObject_DisconnectEnabledChanged(void* ptr);
void QGraphicsObject_GrabGesture(void* ptr, int gesture, int flags);
void QGraphicsObject_ConnectOpacityChanged(void* ptr);
void QGraphicsObject_DisconnectOpacityChanged(void* ptr);
void QGraphicsObject_ConnectParentChanged(void* ptr);
void QGraphicsObject_DisconnectParentChanged(void* ptr);
void QGraphicsObject_ConnectRotationChanged(void* ptr);
void QGraphicsObject_DisconnectRotationChanged(void* ptr);
void QGraphicsObject_ConnectScaleChanged(void* ptr);
void QGraphicsObject_DisconnectScaleChanged(void* ptr);
void QGraphicsObject_UngrabGesture(void* ptr, int gesture);
void QGraphicsObject_ConnectVisibleChanged(void* ptr);
void QGraphicsObject_DisconnectVisibleChanged(void* ptr);
void QGraphicsObject_ConnectXChanged(void* ptr);
void QGraphicsObject_DisconnectXChanged(void* ptr);
void QGraphicsObject_ConnectYChanged(void* ptr);
void QGraphicsObject_DisconnectYChanged(void* ptr);
void QGraphicsObject_ConnectZChanged(void* ptr);
void QGraphicsObject_DisconnectZChanged(void* ptr);
void QGraphicsObject_DestroyQGraphicsObject(void* ptr);

#ifdef __cplusplus
}
#endif