#ifdef __cplusplus
extern "C" {
#endif

double QGraphicsDropShadowEffect_BlurRadius(void* ptr);
void* QGraphicsDropShadowEffect_Color(void* ptr);
void QGraphicsDropShadowEffect_SetBlurRadius(void* ptr, double blurRadius);
void QGraphicsDropShadowEffect_SetColor(void* ptr, void* color);
void QGraphicsDropShadowEffect_SetOffset(void* ptr, void* ofs);
void* QGraphicsDropShadowEffect_NewQGraphicsDropShadowEffect(void* parent);
void QGraphicsDropShadowEffect_ConnectColorChanged(void* ptr);
void QGraphicsDropShadowEffect_DisconnectColorChanged(void* ptr);
void QGraphicsDropShadowEffect_SetOffset3(void* ptr, double d);
void QGraphicsDropShadowEffect_SetOffset2(void* ptr, double dx, double dy);
void QGraphicsDropShadowEffect_SetXOffset(void* ptr, double dx);
void QGraphicsDropShadowEffect_SetYOffset(void* ptr, double dy);
double QGraphicsDropShadowEffect_XOffset(void* ptr);
double QGraphicsDropShadowEffect_YOffset(void* ptr);
void QGraphicsDropShadowEffect_DestroyQGraphicsDropShadowEffect(void* ptr);

#ifdef __cplusplus
}
#endif