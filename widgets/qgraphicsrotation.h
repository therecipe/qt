#ifdef __cplusplus
extern "C" {
#endif

double QGraphicsRotation_Angle(void* ptr);
void QGraphicsRotation_SetAngle(void* ptr, double v);
void QGraphicsRotation_SetAxis2(void* ptr, int axis);
void QGraphicsRotation_SetAxis(void* ptr, void* axis);
void QGraphicsRotation_SetOrigin(void* ptr, void* point);
void* QGraphicsRotation_NewQGraphicsRotation(void* parent);
void QGraphicsRotation_ConnectAngleChanged(void* ptr);
void QGraphicsRotation_DisconnectAngleChanged(void* ptr);
void QGraphicsRotation_ApplyTo(void* ptr, void* matrix);
void QGraphicsRotation_ConnectAxisChanged(void* ptr);
void QGraphicsRotation_DisconnectAxisChanged(void* ptr);
void QGraphicsRotation_ConnectOriginChanged(void* ptr);
void QGraphicsRotation_DisconnectOriginChanged(void* ptr);
void QGraphicsRotation_DestroyQGraphicsRotation(void* ptr);

#ifdef __cplusplus
}
#endif