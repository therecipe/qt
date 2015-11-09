#ifdef __cplusplus
extern "C" {
#endif

void QKeySequenceEdit_SetKeySequence(void* ptr, void* keySequence);
void* QKeySequenceEdit_NewQKeySequenceEdit(void* parent);
void* QKeySequenceEdit_NewQKeySequenceEdit2(void* keySequence, void* parent);
void QKeySequenceEdit_Clear(void* ptr);
void QKeySequenceEdit_ConnectEditingFinished(void* ptr);
void QKeySequenceEdit_DisconnectEditingFinished(void* ptr);
void QKeySequenceEdit_DestroyQKeySequenceEdit(void* ptr);

#ifdef __cplusplus
}
#endif