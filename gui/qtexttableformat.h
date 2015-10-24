#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QTextTableFormat_NewQTextTableFormat();
int QTextTableFormat_Alignment(QtObjectPtr ptr);
void QTextTableFormat_ClearColumnWidthConstraints(QtObjectPtr ptr);
int QTextTableFormat_Columns(QtObjectPtr ptr);
int QTextTableFormat_HeaderRowCount(QtObjectPtr ptr);
int QTextTableFormat_IsValid(QtObjectPtr ptr);
void QTextTableFormat_SetAlignment(QtObjectPtr ptr, int alignment);
void QTextTableFormat_SetHeaderRowCount(QtObjectPtr ptr, int count);

#ifdef __cplusplus
}
#endif