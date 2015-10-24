#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QStylePainter_NewQStylePainter();
QtObjectPtr QStylePainter_NewQStylePainter3(QtObjectPtr pd, QtObjectPtr widget);
QtObjectPtr QStylePainter_NewQStylePainter2(QtObjectPtr widget);
int QStylePainter_Begin2(QtObjectPtr ptr, QtObjectPtr pd, QtObjectPtr widget);
int QStylePainter_Begin(QtObjectPtr ptr, QtObjectPtr widget);
void QStylePainter_DrawComplexControl(QtObjectPtr ptr, int cc, QtObjectPtr option);
void QStylePainter_DrawControl(QtObjectPtr ptr, int ce, QtObjectPtr option);
void QStylePainter_DrawItemPixmap(QtObjectPtr ptr, QtObjectPtr rect, int flags, QtObjectPtr pixmap);
void QStylePainter_DrawItemText(QtObjectPtr ptr, QtObjectPtr rect, int flags, QtObjectPtr pal, int enabled, char* text, int textRole);
void QStylePainter_DrawPrimitive(QtObjectPtr ptr, int pe, QtObjectPtr option);
QtObjectPtr QStylePainter_Style(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif