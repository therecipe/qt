#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QCommonStyle_DrawControl(QtObjectPtr ptr, int element, QtObjectPtr opt, QtObjectPtr p, QtObjectPtr widget);
void QCommonStyle_DrawPrimitive(QtObjectPtr ptr, int pe, QtObjectPtr opt, QtObjectPtr p, QtObjectPtr widget);
void QCommonStyle_DrawComplexControl(QtObjectPtr ptr, int cc, QtObjectPtr opt, QtObjectPtr p, QtObjectPtr widget);
int QCommonStyle_HitTestComplexControl(QtObjectPtr ptr, int cc, QtObjectPtr opt, QtObjectPtr pt, QtObjectPtr widget);
int QCommonStyle_LayoutSpacing(QtObjectPtr ptr, int control1, int control2, int orientation, QtObjectPtr option, QtObjectPtr widget);
int QCommonStyle_PixelMetric(QtObjectPtr ptr, int m, QtObjectPtr opt, QtObjectPtr widget);
void QCommonStyle_Polish2(QtObjectPtr ptr, QtObjectPtr app);
void QCommonStyle_Polish(QtObjectPtr ptr, QtObjectPtr pal);
void QCommonStyle_Polish3(QtObjectPtr ptr, QtObjectPtr widget);
int QCommonStyle_StyleHint(QtObjectPtr ptr, int sh, QtObjectPtr opt, QtObjectPtr widget, QtObjectPtr hret);
void QCommonStyle_Unpolish2(QtObjectPtr ptr, QtObjectPtr application);
void QCommonStyle_Unpolish(QtObjectPtr ptr, QtObjectPtr widget);
void QCommonStyle_DestroyQCommonStyle(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif