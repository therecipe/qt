#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QTextDocumentFragment_NewQTextDocumentFragment4(QtObjectPtr other);
QtObjectPtr QTextDocumentFragment_NewQTextDocumentFragment();
QtObjectPtr QTextDocumentFragment_NewQTextDocumentFragment3(QtObjectPtr cursor);
QtObjectPtr QTextDocumentFragment_NewQTextDocumentFragment2(QtObjectPtr document);
int QTextDocumentFragment_IsEmpty(QtObjectPtr ptr);
char* QTextDocumentFragment_ToHtml(QtObjectPtr ptr, QtObjectPtr encoding);
char* QTextDocumentFragment_ToPlainText(QtObjectPtr ptr);
void QTextDocumentFragment_DestroyQTextDocumentFragment(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif