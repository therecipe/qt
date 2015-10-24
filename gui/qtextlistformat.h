#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QTextListFormat_NewQTextListFormat();
int QTextListFormat_Indent(QtObjectPtr ptr);
int QTextListFormat_IsValid(QtObjectPtr ptr);
char* QTextListFormat_NumberPrefix(QtObjectPtr ptr);
char* QTextListFormat_NumberSuffix(QtObjectPtr ptr);
void QTextListFormat_SetIndent(QtObjectPtr ptr, int indentation);
void QTextListFormat_SetNumberPrefix(QtObjectPtr ptr, char* numberPrefix);
void QTextListFormat_SetNumberSuffix(QtObjectPtr ptr, char* numberSuffix);
void QTextListFormat_SetStyle(QtObjectPtr ptr, int style);
int QTextListFormat_Style(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif