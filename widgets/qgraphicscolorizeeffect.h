#ifdef __cplusplus
extern "C" {
#endif

void* QGraphicsColorizeEffect_Color(void* ptr);
void QGraphicsColorizeEffect_SetColor(void* ptr, void* c);
void QGraphicsColorizeEffect_SetStrength(void* ptr, double strength);
double QGraphicsColorizeEffect_Strength(void* ptr);
void* QGraphicsColorizeEffect_NewQGraphicsColorizeEffect(void* parent);
void QGraphicsColorizeEffect_ConnectColorChanged(void* ptr);
void QGraphicsColorizeEffect_DisconnectColorChanged(void* ptr);
void QGraphicsColorizeEffect_DestroyQGraphicsColorizeEffect(void* ptr);

#ifdef __cplusplus
}
#endif