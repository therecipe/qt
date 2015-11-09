#ifdef __cplusplus
extern "C" {
#endif

void* QTextFrameFormat_NewQTextFrameFormat();
double QTextFrameFormat_BottomMargin(void* ptr);
double QTextFrameFormat_LeftMargin(void* ptr);
double QTextFrameFormat_RightMargin(void* ptr);
void QTextFrameFormat_SetMargin(void* ptr, double margin);
double QTextFrameFormat_TopMargin(void* ptr);
double QTextFrameFormat_Border(void* ptr);
void* QTextFrameFormat_BorderBrush(void* ptr);
int QTextFrameFormat_BorderStyle(void* ptr);
int QTextFrameFormat_IsValid(void* ptr);
double QTextFrameFormat_Margin(void* ptr);
double QTextFrameFormat_Padding(void* ptr);
int QTextFrameFormat_PageBreakPolicy(void* ptr);
int QTextFrameFormat_Position(void* ptr);
void QTextFrameFormat_SetBorder(void* ptr, double width);
void QTextFrameFormat_SetBorderBrush(void* ptr, void* brush);
void QTextFrameFormat_SetBorderStyle(void* ptr, int style);
void QTextFrameFormat_SetBottomMargin(void* ptr, double margin);
void QTextFrameFormat_SetHeight(void* ptr, void* height);
void QTextFrameFormat_SetHeight2(void* ptr, double height);
void QTextFrameFormat_SetLeftMargin(void* ptr, double margin);
void QTextFrameFormat_SetPadding(void* ptr, double width);
void QTextFrameFormat_SetPageBreakPolicy(void* ptr, int policy);
void QTextFrameFormat_SetPosition(void* ptr, int policy);
void QTextFrameFormat_SetRightMargin(void* ptr, double margin);
void QTextFrameFormat_SetTopMargin(void* ptr, double margin);
void QTextFrameFormat_SetWidth(void* ptr, void* width);
void QTextFrameFormat_SetWidth2(void* ptr, double width);

#ifdef __cplusplus
}
#endif