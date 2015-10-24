#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QTextLayout_DrawCursor2(QtObjectPtr ptr, QtObjectPtr painter, QtObjectPtr position, int cursorPosition);
void QTextLayout_DrawCursor(QtObjectPtr ptr, QtObjectPtr painter, QtObjectPtr position, int cursorPosition, int width);
QtObjectPtr QTextLayout_NewQTextLayout();
QtObjectPtr QTextLayout_NewQTextLayout2(char* text);
QtObjectPtr QTextLayout_NewQTextLayout3(char* text, QtObjectPtr font, QtObjectPtr paintdevice);
void QTextLayout_BeginLayout(QtObjectPtr ptr);
int QTextLayout_CacheEnabled(QtObjectPtr ptr);
void QTextLayout_ClearAdditionalFormats(QtObjectPtr ptr);
void QTextLayout_ClearLayout(QtObjectPtr ptr);
int QTextLayout_CursorMoveStyle(QtObjectPtr ptr);
void QTextLayout_EndLayout(QtObjectPtr ptr);
int QTextLayout_IsValidCursorPosition(QtObjectPtr ptr, int pos);
int QTextLayout_LeftCursorPosition(QtObjectPtr ptr, int oldPos);
int QTextLayout_LineCount(QtObjectPtr ptr);
int QTextLayout_NextCursorPosition(QtObjectPtr ptr, int oldPos, int mode);
int QTextLayout_PreeditAreaPosition(QtObjectPtr ptr);
char* QTextLayout_PreeditAreaText(QtObjectPtr ptr);
int QTextLayout_PreviousCursorPosition(QtObjectPtr ptr, int oldPos, int mode);
int QTextLayout_RightCursorPosition(QtObjectPtr ptr, int oldPos);
void QTextLayout_SetCacheEnabled(QtObjectPtr ptr, int enable);
void QTextLayout_SetCursorMoveStyle(QtObjectPtr ptr, int style);
void QTextLayout_SetFont(QtObjectPtr ptr, QtObjectPtr font);
void QTextLayout_SetPosition(QtObjectPtr ptr, QtObjectPtr p);
void QTextLayout_SetPreeditArea(QtObjectPtr ptr, int position, char* text);
void QTextLayout_SetText(QtObjectPtr ptr, char* stri);
void QTextLayout_SetTextOption(QtObjectPtr ptr, QtObjectPtr option);
char* QTextLayout_Text(QtObjectPtr ptr);
void QTextLayout_DestroyQTextLayout(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif