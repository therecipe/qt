#ifdef __cplusplus
extern "C" {
#endif

int QDialog_IsSizeGripEnabled(void* ptr);
void QDialog_SetModal(void* ptr, int modal);
void QDialog_SetResult(void* ptr, int i);
void QDialog_SetSizeGripEnabled(void* ptr, int v);
void* QDialog_NewQDialog(void* parent, int f);
void QDialog_Accept(void* ptr);
void QDialog_ConnectAccepted(void* ptr);
void QDialog_DisconnectAccepted(void* ptr);
void QDialog_Done(void* ptr, int r);
int QDialog_Exec(void* ptr);
void QDialog_ConnectFinished(void* ptr);
void QDialog_DisconnectFinished(void* ptr);
void QDialog_Open(void* ptr);
void QDialog_Reject(void* ptr);
void QDialog_ConnectRejected(void* ptr);
void QDialog_DisconnectRejected(void* ptr);
int QDialog_Result(void* ptr);
void QDialog_SetVisible(void* ptr, int visible);
void QDialog_DestroyQDialog(void* ptr);

#ifdef __cplusplus
}
#endif