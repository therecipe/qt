#ifdef __cplusplus
extern "C" {
#endif

void* QStaticText_NewQStaticText();
void* QStaticText_NewQStaticText3(void* other);
void* QStaticText_NewQStaticText2(char* text);
int QStaticText_PerformanceHint(void* ptr);
void QStaticText_Prepare(void* ptr, void* matrix, void* font);
void QStaticText_SetPerformanceHint(void* ptr, int performanceHint);
void QStaticText_SetText(void* ptr, char* text);
void QStaticText_SetTextFormat(void* ptr, int textFormat);
void QStaticText_SetTextOption(void* ptr, void* textOption);
void QStaticText_SetTextWidth(void* ptr, double textWidth);
void QStaticText_Swap(void* ptr, void* other);
char* QStaticText_Text(void* ptr);
int QStaticText_TextFormat(void* ptr);
double QStaticText_TextWidth(void* ptr);
void QStaticText_DestroyQStaticText(void* ptr);

#ifdef __cplusplus
}
#endif