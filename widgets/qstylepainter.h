#ifdef __cplusplus
extern "C" {
#endif

void* QStylePainter_NewQStylePainter();
void* QStylePainter_NewQStylePainter3(void* pd, void* widget);
void* QStylePainter_NewQStylePainter2(void* widget);
int QStylePainter_Begin2(void* ptr, void* pd, void* widget);
int QStylePainter_Begin(void* ptr, void* widget);
void QStylePainter_DrawComplexControl(void* ptr, int cc, void* option);
void QStylePainter_DrawControl(void* ptr, int ce, void* option);
void QStylePainter_DrawItemPixmap(void* ptr, void* rect, int flags, void* pixmap);
void QStylePainter_DrawItemText(void* ptr, void* rect, int flags, void* pal, int enabled, char* text, int textRole);
void QStylePainter_DrawPrimitive(void* ptr, int pe, void* option);
void* QStylePainter_Style(void* ptr);

#ifdef __cplusplus
}
#endif