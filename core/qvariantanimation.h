#ifdef __cplusplus
extern "C" {
#endif

void* QVariantAnimation_CurrentValue(void* ptr);
int QVariantAnimation_Duration(void* ptr);
void* QVariantAnimation_EasingCurve(void* ptr);
void* QVariantAnimation_EndValue(void* ptr);
void QVariantAnimation_SetDuration(void* ptr, int msecs);
void QVariantAnimation_SetEasingCurve(void* ptr, void* easing);
void QVariantAnimation_SetEndValue(void* ptr, void* value);
void QVariantAnimation_SetStartValue(void* ptr, void* value);
void* QVariantAnimation_StartValue(void* ptr);
void* QVariantAnimation_NewQVariantAnimation(void* parent);
void* QVariantAnimation_KeyValueAt(void* ptr, double step);
void QVariantAnimation_SetKeyValueAt(void* ptr, double step, void* value);
void QVariantAnimation_ConnectValueChanged(void* ptr);
void QVariantAnimation_DisconnectValueChanged(void* ptr);
void QVariantAnimation_DestroyQVariantAnimation(void* ptr);

#ifdef __cplusplus
}
#endif