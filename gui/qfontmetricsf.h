#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QFontMetricsF_NewQFontMetricsF(QtObjectPtr font);
QtObjectPtr QFontMetricsF_NewQFontMetricsF2(QtObjectPtr font, QtObjectPtr paintdevice);
QtObjectPtr QFontMetricsF_NewQFontMetricsF3(QtObjectPtr fontMetrics);
QtObjectPtr QFontMetricsF_NewQFontMetricsF4(QtObjectPtr fm);
int QFontMetricsF_InFont(QtObjectPtr ptr, QtObjectPtr ch);
void QFontMetricsF_Swap(QtObjectPtr ptr, QtObjectPtr other);
void QFontMetricsF_DestroyQFontMetricsF(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif