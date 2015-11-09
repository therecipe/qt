#ifdef __cplusplus
extern "C" {
#endif

void* QProxyStyle_BaseStyle(void* ptr);
void QProxyStyle_DrawComplexControl(void* ptr, int control, void* option, void* painter, void* widget);
void QProxyStyle_DrawControl(void* ptr, int element, void* option, void* painter, void* widget);
void QProxyStyle_DrawItemPixmap(void* ptr, void* painter, void* rect, int alignment, void* pixmap);
void QProxyStyle_DrawItemText(void* ptr, void* painter, void* rect, int flags, void* pal, int enabled, char* text, int textRole);
void QProxyStyle_DrawPrimitive(void* ptr, int element, void* option, void* painter, void* widget);
int QProxyStyle_HitTestComplexControl(void* ptr, int control, void* option, void* pos, void* widget);
int QProxyStyle_LayoutSpacing(void* ptr, int control1, int control2, int orientation, void* option, void* widget);
int QProxyStyle_PixelMetric(void* ptr, int metric, void* option, void* widget);
void QProxyStyle_Polish3(void* ptr, void* app);
void QProxyStyle_Polish2(void* ptr, void* pal);
void QProxyStyle_Polish(void* ptr, void* widget);
void QProxyStyle_SetBaseStyle(void* ptr, void* style);
int QProxyStyle_StyleHint(void* ptr, int hint, void* option, void* widget, void* returnData);
void QProxyStyle_Unpolish2(void* ptr, void* app);
void QProxyStyle_Unpolish(void* ptr, void* widget);
void QProxyStyle_DestroyQProxyStyle(void* ptr);

#ifdef __cplusplus
}
#endif