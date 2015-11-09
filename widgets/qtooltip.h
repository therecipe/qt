#ifdef __cplusplus
extern "C" {
#endif

void QToolTip_QToolTip_HideText();
int QToolTip_QToolTip_IsVisible();
void QToolTip_QToolTip_SetFont(void* font);
void QToolTip_QToolTip_SetPalette(void* palette);
void QToolTip_QToolTip_ShowText3(void* pos, char* text, void* w);
void QToolTip_QToolTip_ShowText(void* pos, char* text, void* w, void* rect);
void QToolTip_QToolTip_ShowText2(void* pos, char* text, void* w, void* rect, int msecDisplayTime);
char* QToolTip_QToolTip_Text();

#ifdef __cplusplus
}
#endif