#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QTextFrame_NewQTextFrame(QtObjectPtr document);
int QTextFrame_FirstPosition(QtObjectPtr ptr);
int QTextFrame_LastPosition(QtObjectPtr ptr);
QtObjectPtr QTextFrame_ParentFrame(QtObjectPtr ptr);
void QTextFrame_SetFrameFormat(QtObjectPtr ptr, QtObjectPtr format);
void QTextFrame_DestroyQTextFrame(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif