#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QKeySequenceEdit_SetKeySequence(QtObjectPtr ptr, QtObjectPtr keySequence);
QtObjectPtr QKeySequenceEdit_NewQKeySequenceEdit(QtObjectPtr parent);
QtObjectPtr QKeySequenceEdit_NewQKeySequenceEdit2(QtObjectPtr keySequence, QtObjectPtr parent);
void QKeySequenceEdit_Clear(QtObjectPtr ptr);
void QKeySequenceEdit_ConnectEditingFinished(QtObjectPtr ptr);
void QKeySequenceEdit_DisconnectEditingFinished(QtObjectPtr ptr);
void QKeySequenceEdit_DestroyQKeySequenceEdit(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif