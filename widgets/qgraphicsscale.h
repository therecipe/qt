#ifdef __cplusplus
extern "C" {
#endif

void QGraphicsScale_SetOrigin(void* ptr, void* point);
void QGraphicsScale_SetXScale(void* ptr, double v);
void QGraphicsScale_SetYScale(void* ptr, double v);
void QGraphicsScale_SetZScale(void* ptr, double v);
double QGraphicsScale_XScale(void* ptr);
double QGraphicsScale_YScale(void* ptr);
double QGraphicsScale_ZScale(void* ptr);
void* QGraphicsScale_NewQGraphicsScale(void* parent);
void QGraphicsScale_ApplyTo(void* ptr, void* matrix);
void QGraphicsScale_ConnectOriginChanged(void* ptr);
void QGraphicsScale_DisconnectOriginChanged(void* ptr);
void QGraphicsScale_ConnectScaleChanged(void* ptr);
void QGraphicsScale_DisconnectScaleChanged(void* ptr);
void QGraphicsScale_ConnectXScaleChanged(void* ptr);
void QGraphicsScale_DisconnectXScaleChanged(void* ptr);
void QGraphicsScale_ConnectYScaleChanged(void* ptr);
void QGraphicsScale_DisconnectYScaleChanged(void* ptr);
void QGraphicsScale_ConnectZScaleChanged(void* ptr);
void QGraphicsScale_DisconnectZScaleChanged(void* ptr);
void QGraphicsScale_DestroyQGraphicsScale(void* ptr);

#ifdef __cplusplus
}
#endif