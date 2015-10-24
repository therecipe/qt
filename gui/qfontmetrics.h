#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QFontMetrics_NewQFontMetrics(QtObjectPtr font);
QtObjectPtr QFontMetrics_NewQFontMetrics2(QtObjectPtr font, QtObjectPtr paintdevice);
QtObjectPtr QFontMetrics_NewQFontMetrics3(QtObjectPtr fm);
int QFontMetrics_Ascent(QtObjectPtr ptr);
int QFontMetrics_AverageCharWidth(QtObjectPtr ptr);
int QFontMetrics_Descent(QtObjectPtr ptr);
char* QFontMetrics_ElidedText(QtObjectPtr ptr, char* text, int mode, int width, int flags);
int QFontMetrics_Height(QtObjectPtr ptr);
int QFontMetrics_InFont(QtObjectPtr ptr, QtObjectPtr ch);
int QFontMetrics_Leading(QtObjectPtr ptr);
int QFontMetrics_LeftBearing(QtObjectPtr ptr, QtObjectPtr ch);
int QFontMetrics_LineSpacing(QtObjectPtr ptr);
int QFontMetrics_LineWidth(QtObjectPtr ptr);
int QFontMetrics_MaxWidth(QtObjectPtr ptr);
int QFontMetrics_MinLeftBearing(QtObjectPtr ptr);
int QFontMetrics_MinRightBearing(QtObjectPtr ptr);
int QFontMetrics_OverlinePos(QtObjectPtr ptr);
int QFontMetrics_RightBearing(QtObjectPtr ptr, QtObjectPtr ch);
int QFontMetrics_StrikeOutPos(QtObjectPtr ptr);
void QFontMetrics_Swap(QtObjectPtr ptr, QtObjectPtr other);
int QFontMetrics_UnderlinePos(QtObjectPtr ptr);
int QFontMetrics_Width3(QtObjectPtr ptr, QtObjectPtr ch);
int QFontMetrics_Width(QtObjectPtr ptr, char* text, int len);
int QFontMetrics_XHeight(QtObjectPtr ptr);
void QFontMetrics_DestroyQFontMetrics(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif