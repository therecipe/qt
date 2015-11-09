#ifdef __cplusplus
extern "C" {
#endif

double QTextInlineObject_Ascent(void* ptr);
double QTextInlineObject_Descent(void* ptr);
int QTextInlineObject_FormatIndex(void* ptr);
double QTextInlineObject_Height(void* ptr);
int QTextInlineObject_IsValid(void* ptr);
void QTextInlineObject_SetAscent(void* ptr, double a);
void QTextInlineObject_SetDescent(void* ptr, double d);
void QTextInlineObject_SetWidth(void* ptr, double w);
int QTextInlineObject_TextDirection(void* ptr);
int QTextInlineObject_TextPosition(void* ptr);
double QTextInlineObject_Width(void* ptr);

#ifdef __cplusplus
}
#endif