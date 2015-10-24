#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

char* QVariantAnimation_CurrentValue(QtObjectPtr ptr);
int QVariantAnimation_Duration(QtObjectPtr ptr);
char* QVariantAnimation_EndValue(QtObjectPtr ptr);
void QVariantAnimation_SetDuration(QtObjectPtr ptr, int msecs);
void QVariantAnimation_SetEasingCurve(QtObjectPtr ptr, QtObjectPtr easing);
void QVariantAnimation_SetEndValue(QtObjectPtr ptr, char* value);
void QVariantAnimation_SetStartValue(QtObjectPtr ptr, char* value);
char* QVariantAnimation_StartValue(QtObjectPtr ptr);
QtObjectPtr QVariantAnimation_NewQVariantAnimation(QtObjectPtr parent);
void QVariantAnimation_ConnectValueChanged(QtObjectPtr ptr);
void QVariantAnimation_DisconnectValueChanged(QtObjectPtr ptr);
void QVariantAnimation_DestroyQVariantAnimation(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif