#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QTextCharFormat_NewQTextCharFormat();
char* QTextCharFormat_AnchorNames(QtObjectPtr ptr);
int QTextCharFormat_FontUnderline(QtObjectPtr ptr);
void QTextCharFormat_SetUnderlineStyle(QtObjectPtr ptr, int style);
char* QTextCharFormat_AnchorHref(QtObjectPtr ptr);
int QTextCharFormat_FontCapitalization(QtObjectPtr ptr);
char* QTextCharFormat_FontFamily(QtObjectPtr ptr);
int QTextCharFormat_FontFixedPitch(QtObjectPtr ptr);
int QTextCharFormat_FontHintingPreference(QtObjectPtr ptr);
int QTextCharFormat_FontItalic(QtObjectPtr ptr);
int QTextCharFormat_FontKerning(QtObjectPtr ptr);
int QTextCharFormat_FontLetterSpacingType(QtObjectPtr ptr);
int QTextCharFormat_FontOverline(QtObjectPtr ptr);
int QTextCharFormat_FontStretch(QtObjectPtr ptr);
int QTextCharFormat_FontStrikeOut(QtObjectPtr ptr);
int QTextCharFormat_FontStyleHint(QtObjectPtr ptr);
int QTextCharFormat_FontStyleStrategy(QtObjectPtr ptr);
int QTextCharFormat_FontWeight(QtObjectPtr ptr);
int QTextCharFormat_IsAnchor(QtObjectPtr ptr);
int QTextCharFormat_IsValid(QtObjectPtr ptr);
void QTextCharFormat_SetAnchor(QtObjectPtr ptr, int anchor);
void QTextCharFormat_SetAnchorHref(QtObjectPtr ptr, char* value);
void QTextCharFormat_SetAnchorNames(QtObjectPtr ptr, char* names);
void QTextCharFormat_SetFont2(QtObjectPtr ptr, QtObjectPtr font);
void QTextCharFormat_SetFont(QtObjectPtr ptr, QtObjectPtr font, int behavior);
void QTextCharFormat_SetFontCapitalization(QtObjectPtr ptr, int capitalization);
void QTextCharFormat_SetFontFamily(QtObjectPtr ptr, char* family);
void QTextCharFormat_SetFontFixedPitch(QtObjectPtr ptr, int fixedPitch);
void QTextCharFormat_SetFontHintingPreference(QtObjectPtr ptr, int hintingPreference);
void QTextCharFormat_SetFontItalic(QtObjectPtr ptr, int italic);
void QTextCharFormat_SetFontKerning(QtObjectPtr ptr, int enable);
void QTextCharFormat_SetFontLetterSpacingType(QtObjectPtr ptr, int letterSpacingType);
void QTextCharFormat_SetFontOverline(QtObjectPtr ptr, int overline);
void QTextCharFormat_SetFontStretch(QtObjectPtr ptr, int factor);
void QTextCharFormat_SetFontStrikeOut(QtObjectPtr ptr, int strikeOut);
void QTextCharFormat_SetFontStyleHint(QtObjectPtr ptr, int hint, int strategy);
void QTextCharFormat_SetFontStyleStrategy(QtObjectPtr ptr, int strategy);
void QTextCharFormat_SetFontUnderline(QtObjectPtr ptr, int underline);
void QTextCharFormat_SetFontWeight(QtObjectPtr ptr, int weight);
void QTextCharFormat_SetTextOutline(QtObjectPtr ptr, QtObjectPtr pen);
void QTextCharFormat_SetToolTip(QtObjectPtr ptr, char* text);
void QTextCharFormat_SetUnderlineColor(QtObjectPtr ptr, QtObjectPtr color);
void QTextCharFormat_SetVerticalAlignment(QtObjectPtr ptr, int alignment);
char* QTextCharFormat_ToolTip(QtObjectPtr ptr);
int QTextCharFormat_UnderlineStyle(QtObjectPtr ptr);
int QTextCharFormat_VerticalAlignment(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif