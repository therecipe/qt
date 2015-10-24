#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QProxyStyle_BaseStyle(QtObjectPtr ptr);
void QProxyStyle_DrawComplexControl(QtObjectPtr ptr, int control, QtObjectPtr option, QtObjectPtr painter, QtObjectPtr widget);
void QProxyStyle_DrawControl(QtObjectPtr ptr, int element, QtObjectPtr option, QtObjectPtr painter, QtObjectPtr widget);
void QProxyStyle_DrawItemPixmap(QtObjectPtr ptr, QtObjectPtr painter, QtObjectPtr rect, int alignment, QtObjectPtr pixmap);
void QProxyStyle_DrawItemText(QtObjectPtr ptr, QtObjectPtr painter, QtObjectPtr rect, int flags, QtObjectPtr pal, int enabled, char* text, int textRole);
void QProxyStyle_DrawPrimitive(QtObjectPtr ptr, int element, QtObjectPtr option, QtObjectPtr painter, QtObjectPtr widget);
int QProxyStyle_HitTestComplexControl(QtObjectPtr ptr, int control, QtObjectPtr option, QtObjectPtr pos, QtObjectPtr widget);
int QProxyStyle_LayoutSpacing(QtObjectPtr ptr, int control1, int control2, int orientation, QtObjectPtr option, QtObjectPtr widget);
int QProxyStyle_PixelMetric(QtObjectPtr ptr, int metric, QtObjectPtr option, QtObjectPtr widget);
void QProxyStyle_Polish3(QtObjectPtr ptr, QtObjectPtr app);
void QProxyStyle_Polish2(QtObjectPtr ptr, QtObjectPtr pal);
void QProxyStyle_Polish(QtObjectPtr ptr, QtObjectPtr widget);
void QProxyStyle_SetBaseStyle(QtObjectPtr ptr, QtObjectPtr style);
int QProxyStyle_StyleHint(QtObjectPtr ptr, int hint, QtObjectPtr option, QtObjectPtr widget, QtObjectPtr returnData);
void QProxyStyle_Unpolish2(QtObjectPtr ptr, QtObjectPtr app);
void QProxyStyle_Unpolish(QtObjectPtr ptr, QtObjectPtr widget);
void QProxyStyle_DestroyQProxyStyle(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif