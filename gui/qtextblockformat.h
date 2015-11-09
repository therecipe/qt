#ifdef __cplusplus
extern "C" {
#endif

void* QTextBlockFormat_NewQTextBlockFormat();
int QTextBlockFormat_Alignment(void* ptr);
double QTextBlockFormat_BottomMargin(void* ptr);
int QTextBlockFormat_Indent(void* ptr);
int QTextBlockFormat_IsValid(void* ptr);
double QTextBlockFormat_LeftMargin(void* ptr);
double QTextBlockFormat_LineHeight2(void* ptr);
double QTextBlockFormat_LineHeight(void* ptr, double scriptLineHeight, double scaling);
int QTextBlockFormat_LineHeightType(void* ptr);
int QTextBlockFormat_NonBreakableLines(void* ptr);
int QTextBlockFormat_PageBreakPolicy(void* ptr);
double QTextBlockFormat_RightMargin(void* ptr);
void QTextBlockFormat_SetAlignment(void* ptr, int alignment);
void QTextBlockFormat_SetBottomMargin(void* ptr, double margin);
void QTextBlockFormat_SetIndent(void* ptr, int indentation);
void QTextBlockFormat_SetLeftMargin(void* ptr, double margin);
void QTextBlockFormat_SetLineHeight(void* ptr, double height, int heightType);
void QTextBlockFormat_SetNonBreakableLines(void* ptr, int b);
void QTextBlockFormat_SetPageBreakPolicy(void* ptr, int policy);
void QTextBlockFormat_SetRightMargin(void* ptr, double margin);
void QTextBlockFormat_SetTextIndent(void* ptr, double indent);
void QTextBlockFormat_SetTopMargin(void* ptr, double margin);
double QTextBlockFormat_TextIndent(void* ptr);
double QTextBlockFormat_TopMargin(void* ptr);

#ifdef __cplusplus
}
#endif