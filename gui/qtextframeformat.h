#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QTextFrameFormat_NewQTextFrameFormat();
int QTextFrameFormat_BorderStyle(QtObjectPtr ptr);
int QTextFrameFormat_IsValid(QtObjectPtr ptr);
int QTextFrameFormat_PageBreakPolicy(QtObjectPtr ptr);
int QTextFrameFormat_Position(QtObjectPtr ptr);
void QTextFrameFormat_SetBorderBrush(QtObjectPtr ptr, QtObjectPtr brush);
void QTextFrameFormat_SetBorderStyle(QtObjectPtr ptr, int style);
void QTextFrameFormat_SetHeight(QtObjectPtr ptr, QtObjectPtr height);
void QTextFrameFormat_SetPageBreakPolicy(QtObjectPtr ptr, int policy);
void QTextFrameFormat_SetPosition(QtObjectPtr ptr, int policy);
void QTextFrameFormat_SetWidth(QtObjectPtr ptr, QtObjectPtr width);

#ifdef __cplusplus
}
#endif