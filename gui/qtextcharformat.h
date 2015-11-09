#ifdef __cplusplus
extern "C" {
#endif

void* QTextCharFormat_NewQTextCharFormat();
char* QTextCharFormat_AnchorNames(void* ptr);
int QTextCharFormat_FontUnderline(void* ptr);
void QTextCharFormat_SetUnderlineStyle(void* ptr, int style);
char* QTextCharFormat_AnchorHref(void* ptr);
int QTextCharFormat_FontCapitalization(void* ptr);
char* QTextCharFormat_FontFamily(void* ptr);
int QTextCharFormat_FontFixedPitch(void* ptr);
int QTextCharFormat_FontHintingPreference(void* ptr);
int QTextCharFormat_FontItalic(void* ptr);
int QTextCharFormat_FontKerning(void* ptr);
double QTextCharFormat_FontLetterSpacing(void* ptr);
int QTextCharFormat_FontLetterSpacingType(void* ptr);
int QTextCharFormat_FontOverline(void* ptr);
double QTextCharFormat_FontPointSize(void* ptr);
int QTextCharFormat_FontStretch(void* ptr);
int QTextCharFormat_FontStrikeOut(void* ptr);
int QTextCharFormat_FontStyleHint(void* ptr);
int QTextCharFormat_FontStyleStrategy(void* ptr);
int QTextCharFormat_FontWeight(void* ptr);
double QTextCharFormat_FontWordSpacing(void* ptr);
int QTextCharFormat_IsAnchor(void* ptr);
int QTextCharFormat_IsValid(void* ptr);
void QTextCharFormat_SetAnchor(void* ptr, int anchor);
void QTextCharFormat_SetAnchorHref(void* ptr, char* value);
void QTextCharFormat_SetAnchorNames(void* ptr, char* names);
void QTextCharFormat_SetFont2(void* ptr, void* font);
void QTextCharFormat_SetFont(void* ptr, void* font, int behavior);
void QTextCharFormat_SetFontCapitalization(void* ptr, int capitalization);
void QTextCharFormat_SetFontFamily(void* ptr, char* family);
void QTextCharFormat_SetFontFixedPitch(void* ptr, int fixedPitch);
void QTextCharFormat_SetFontHintingPreference(void* ptr, int hintingPreference);
void QTextCharFormat_SetFontItalic(void* ptr, int italic);
void QTextCharFormat_SetFontKerning(void* ptr, int enable);
void QTextCharFormat_SetFontLetterSpacing(void* ptr, double spacing);
void QTextCharFormat_SetFontLetterSpacingType(void* ptr, int letterSpacingType);
void QTextCharFormat_SetFontOverline(void* ptr, int overline);
void QTextCharFormat_SetFontPointSize(void* ptr, double size);
void QTextCharFormat_SetFontStretch(void* ptr, int factor);
void QTextCharFormat_SetFontStrikeOut(void* ptr, int strikeOut);
void QTextCharFormat_SetFontStyleHint(void* ptr, int hint, int strategy);
void QTextCharFormat_SetFontStyleStrategy(void* ptr, int strategy);
void QTextCharFormat_SetFontUnderline(void* ptr, int underline);
void QTextCharFormat_SetFontWeight(void* ptr, int weight);
void QTextCharFormat_SetFontWordSpacing(void* ptr, double spacing);
void QTextCharFormat_SetTextOutline(void* ptr, void* pen);
void QTextCharFormat_SetToolTip(void* ptr, char* text);
void QTextCharFormat_SetUnderlineColor(void* ptr, void* color);
void QTextCharFormat_SetVerticalAlignment(void* ptr, int alignment);
char* QTextCharFormat_ToolTip(void* ptr);
void* QTextCharFormat_UnderlineColor(void* ptr);
int QTextCharFormat_UnderlineStyle(void* ptr);
int QTextCharFormat_VerticalAlignment(void* ptr);

#ifdef __cplusplus
}
#endif