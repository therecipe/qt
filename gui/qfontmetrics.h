#ifdef __cplusplus
extern "C" {
#endif

void* QFontMetrics_NewQFontMetrics(void* font);
void* QFontMetrics_NewQFontMetrics2(void* font, void* paintdevice);
void* QFontMetrics_NewQFontMetrics3(void* fm);
int QFontMetrics_Ascent(void* ptr);
int QFontMetrics_AverageCharWidth(void* ptr);
int QFontMetrics_Descent(void* ptr);
char* QFontMetrics_ElidedText(void* ptr, char* text, int mode, int width, int flags);
int QFontMetrics_Height(void* ptr);
int QFontMetrics_InFont(void* ptr, void* ch);
int QFontMetrics_Leading(void* ptr);
int QFontMetrics_LeftBearing(void* ptr, void* ch);
int QFontMetrics_LineSpacing(void* ptr);
int QFontMetrics_LineWidth(void* ptr);
int QFontMetrics_MaxWidth(void* ptr);
int QFontMetrics_MinLeftBearing(void* ptr);
int QFontMetrics_MinRightBearing(void* ptr);
int QFontMetrics_OverlinePos(void* ptr);
int QFontMetrics_RightBearing(void* ptr, void* ch);
int QFontMetrics_StrikeOutPos(void* ptr);
void QFontMetrics_Swap(void* ptr, void* other);
int QFontMetrics_UnderlinePos(void* ptr);
int QFontMetrics_Width3(void* ptr, void* ch);
int QFontMetrics_Width(void* ptr, char* text, int len);
int QFontMetrics_XHeight(void* ptr);
void QFontMetrics_DestroyQFontMetrics(void* ptr);

#ifdef __cplusplus
}
#endif