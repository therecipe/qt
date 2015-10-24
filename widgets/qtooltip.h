#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QToolTip_QToolTip_HideText();
int QToolTip_QToolTip_IsVisible();
void QToolTip_QToolTip_SetFont(QtObjectPtr font);
void QToolTip_QToolTip_SetPalette(QtObjectPtr palette);
void QToolTip_QToolTip_ShowText3(QtObjectPtr pos, char* text, QtObjectPtr w);
void QToolTip_QToolTip_ShowText(QtObjectPtr pos, char* text, QtObjectPtr w, QtObjectPtr rect);
void QToolTip_QToolTip_ShowText2(QtObjectPtr pos, char* text, QtObjectPtr w, QtObjectPtr rect, int msecDisplayTime);
char* QToolTip_QToolTip_Text();

#ifdef __cplusplus
}
#endif