#ifdef __cplusplus
extern "C" {
#endif

void* QFontMetricsF_NewQFontMetricsF(void* font);
void* QFontMetricsF_NewQFontMetricsF2(void* font, void* paintdevice);
void* QFontMetricsF_NewQFontMetricsF3(void* fontMetrics);
void* QFontMetricsF_NewQFontMetricsF4(void* fm);
double QFontMetricsF_Ascent(void* ptr);
double QFontMetricsF_AverageCharWidth(void* ptr);
double QFontMetricsF_Descent(void* ptr);
char* QFontMetricsF_ElidedText(void* ptr, char* text, int mode, double width, int flags);
double QFontMetricsF_Height(void* ptr);
int QFontMetricsF_InFont(void* ptr, void* ch);
double QFontMetricsF_Leading(void* ptr);
double QFontMetricsF_LeftBearing(void* ptr, void* ch);
double QFontMetricsF_LineSpacing(void* ptr);
double QFontMetricsF_LineWidth(void* ptr);
double QFontMetricsF_MaxWidth(void* ptr);
double QFontMetricsF_MinLeftBearing(void* ptr);
double QFontMetricsF_MinRightBearing(void* ptr);
double QFontMetricsF_OverlinePos(void* ptr);
double QFontMetricsF_RightBearing(void* ptr, void* ch);
double QFontMetricsF_StrikeOutPos(void* ptr);
void QFontMetricsF_Swap(void* ptr, void* other);
double QFontMetricsF_UnderlinePos(void* ptr);
double QFontMetricsF_Width2(void* ptr, void* ch);
double QFontMetricsF_Width(void* ptr, char* text);
double QFontMetricsF_XHeight(void* ptr);
void QFontMetricsF_DestroyQFontMetricsF(void* ptr);

#ifdef __cplusplus
}
#endif