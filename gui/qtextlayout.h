#ifdef __cplusplus
extern "C" {
#endif

void QTextLayout_DrawCursor2(void* ptr, void* painter, void* position, int cursorPosition);
void QTextLayout_DrawCursor(void* ptr, void* painter, void* position, int cursorPosition, int width);
void* QTextLayout_NewQTextLayout();
void* QTextLayout_NewQTextLayout2(char* text);
void* QTextLayout_NewQTextLayout3(char* text, void* font, void* paintdevice);
void QTextLayout_BeginLayout(void* ptr);
int QTextLayout_CacheEnabled(void* ptr);
void QTextLayout_ClearAdditionalFormats(void* ptr);
void QTextLayout_ClearLayout(void* ptr);
int QTextLayout_CursorMoveStyle(void* ptr);
void QTextLayout_EndLayout(void* ptr);
int QTextLayout_IsValidCursorPosition(void* ptr, int pos);
int QTextLayout_LeftCursorPosition(void* ptr, int oldPos);
int QTextLayout_LineCount(void* ptr);
double QTextLayout_MaximumWidth(void* ptr);
double QTextLayout_MinimumWidth(void* ptr);
int QTextLayout_NextCursorPosition(void* ptr, int oldPos, int mode);
int QTextLayout_PreeditAreaPosition(void* ptr);
char* QTextLayout_PreeditAreaText(void* ptr);
int QTextLayout_PreviousCursorPosition(void* ptr, int oldPos, int mode);
int QTextLayout_RightCursorPosition(void* ptr, int oldPos);
void QTextLayout_SetCacheEnabled(void* ptr, int enable);
void QTextLayout_SetCursorMoveStyle(void* ptr, int style);
void QTextLayout_SetFont(void* ptr, void* font);
void QTextLayout_SetPosition(void* ptr, void* p);
void QTextLayout_SetPreeditArea(void* ptr, int position, char* text);
void QTextLayout_SetText(void* ptr, char* stri);
void QTextLayout_SetTextOption(void* ptr, void* option);
char* QTextLayout_Text(void* ptr);
void QTextLayout_DestroyQTextLayout(void* ptr);

#ifdef __cplusplus
}
#endif