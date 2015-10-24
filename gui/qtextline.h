#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QTextLine_NewQTextLine();
int QTextLine_IsValid(QtObjectPtr ptr);
int QTextLine_LeadingIncluded(QtObjectPtr ptr);
int QTextLine_LineNumber(QtObjectPtr ptr);
void QTextLine_SetLeadingIncluded(QtObjectPtr ptr, int included);
void QTextLine_SetNumColumns(QtObjectPtr ptr, int numColumns);
void QTextLine_SetPosition(QtObjectPtr ptr, QtObjectPtr pos);
int QTextLine_TextLength(QtObjectPtr ptr);
int QTextLine_TextStart(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif