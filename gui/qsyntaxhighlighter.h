#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QSyntaxHighlighter_Document(QtObjectPtr ptr);
void QSyntaxHighlighter_Rehighlight(QtObjectPtr ptr);
void QSyntaxHighlighter_RehighlightBlock(QtObjectPtr ptr, QtObjectPtr block);
void QSyntaxHighlighter_SetDocument(QtObjectPtr ptr, QtObjectPtr doc);
void QSyntaxHighlighter_DestroyQSyntaxHighlighter(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif