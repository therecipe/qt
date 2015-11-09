#ifdef __cplusplus
extern "C" {
#endif

int QGraphicsBlurEffect_BlurHints(void* ptr);
double QGraphicsBlurEffect_BlurRadius(void* ptr);
void QGraphicsBlurEffect_SetBlurHints(void* ptr, int hints);
void QGraphicsBlurEffect_SetBlurRadius(void* ptr, double blurRadius);
void* QGraphicsBlurEffect_NewQGraphicsBlurEffect(void* parent);
void QGraphicsBlurEffect_ConnectBlurHintsChanged(void* ptr);
void QGraphicsBlurEffect_DisconnectBlurHintsChanged(void* ptr);
void QGraphicsBlurEffect_DestroyQGraphicsBlurEffect(void* ptr);

#ifdef __cplusplus
}
#endif