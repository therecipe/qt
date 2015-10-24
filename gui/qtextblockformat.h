#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QTextBlockFormat_NewQTextBlockFormat();
int QTextBlockFormat_Alignment(QtObjectPtr ptr);
int QTextBlockFormat_Indent(QtObjectPtr ptr);
int QTextBlockFormat_IsValid(QtObjectPtr ptr);
int QTextBlockFormat_LineHeightType(QtObjectPtr ptr);
int QTextBlockFormat_NonBreakableLines(QtObjectPtr ptr);
int QTextBlockFormat_PageBreakPolicy(QtObjectPtr ptr);
void QTextBlockFormat_SetAlignment(QtObjectPtr ptr, int alignment);
void QTextBlockFormat_SetIndent(QtObjectPtr ptr, int indentation);
void QTextBlockFormat_SetNonBreakableLines(QtObjectPtr ptr, int b);
void QTextBlockFormat_SetPageBreakPolicy(QtObjectPtr ptr, int policy);

#ifdef __cplusplus
}
#endif