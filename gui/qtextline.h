#ifdef __cplusplus
extern "C" {
#endif

int QTextLine_XToCursor(void* ptr, double x, int cpos);
void* QTextLine_NewQTextLine();
double QTextLine_Ascent(void* ptr);
double QTextLine_CursorToX(void* ptr, int cursorPos, int edge);
double QTextLine_CursorToX2(void* ptr, int cursorPos, int edge);
double QTextLine_Descent(void* ptr);
double QTextLine_Height(void* ptr);
double QTextLine_HorizontalAdvance(void* ptr);
int QTextLine_IsValid(void* ptr);
double QTextLine_Leading(void* ptr);
int QTextLine_LeadingIncluded(void* ptr);
int QTextLine_LineNumber(void* ptr);
double QTextLine_NaturalTextWidth(void* ptr);
void QTextLine_SetLeadingIncluded(void* ptr, int included);
void QTextLine_SetLineWidth(void* ptr, double width);
void QTextLine_SetNumColumns(void* ptr, int numColumns);
void QTextLine_SetNumColumns2(void* ptr, int numColumns, double alignmentWidth);
void QTextLine_SetPosition(void* ptr, void* pos);
int QTextLine_TextLength(void* ptr);
int QTextLine_TextStart(void* ptr);
double QTextLine_Width(void* ptr);
double QTextLine_X(void* ptr);
double QTextLine_Y(void* ptr);

#ifdef __cplusplus
}
#endif