#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QPlainTextDocumentLayout_CursorWidth(QtObjectPtr ptr);
void QPlainTextDocumentLayout_SetCursorWidth(QtObjectPtr ptr, int width);
QtObjectPtr QPlainTextDocumentLayout_NewQPlainTextDocumentLayout(QtObjectPtr document);
void QPlainTextDocumentLayout_EnsureBlockLayout(QtObjectPtr ptr, QtObjectPtr block);
int QPlainTextDocumentLayout_PageCount(QtObjectPtr ptr);
void QPlainTextDocumentLayout_RequestUpdate(QtObjectPtr ptr);
void QPlainTextDocumentLayout_DestroyQPlainTextDocumentLayout(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif