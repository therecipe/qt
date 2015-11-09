#ifdef __cplusplus
extern "C" {
#endif

double QGraphicsOpacityEffect_Opacity(void* ptr);
void* QGraphicsOpacityEffect_OpacityMask(void* ptr);
void QGraphicsOpacityEffect_SetOpacity(void* ptr, double opacity);
void QGraphicsOpacityEffect_SetOpacityMask(void* ptr, void* mask);
void* QGraphicsOpacityEffect_NewQGraphicsOpacityEffect(void* parent);
void QGraphicsOpacityEffect_ConnectOpacityMaskChanged(void* ptr);
void QGraphicsOpacityEffect_DisconnectOpacityMaskChanged(void* ptr);
void QGraphicsOpacityEffect_DestroyQGraphicsOpacityEffect(void* ptr);

#ifdef __cplusplus
}
#endif