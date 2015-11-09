#ifdef __cplusplus
extern "C" {
#endif

void QCommonStyle_DrawControl(void* ptr, int element, void* opt, void* p, void* widget);
void QCommonStyle_DrawPrimitive(void* ptr, int pe, void* opt, void* p, void* widget);
void QCommonStyle_DrawComplexControl(void* ptr, int cc, void* opt, void* p, void* widget);
int QCommonStyle_HitTestComplexControl(void* ptr, int cc, void* opt, void* pt, void* widget);
int QCommonStyle_LayoutSpacing(void* ptr, int control1, int control2, int orientation, void* option, void* widget);
int QCommonStyle_PixelMetric(void* ptr, int m, void* opt, void* widget);
void QCommonStyle_Polish2(void* ptr, void* app);
void QCommonStyle_Polish(void* ptr, void* pal);
void QCommonStyle_Polish3(void* ptr, void* widget);
int QCommonStyle_StyleHint(void* ptr, int sh, void* opt, void* widget, void* hret);
void QCommonStyle_Unpolish2(void* ptr, void* application);
void QCommonStyle_Unpolish(void* ptr, void* widget);
void QCommonStyle_DestroyQCommonStyle(void* ptr);

#ifdef __cplusplus
}
#endif