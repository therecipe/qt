#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QDialog_IsSizeGripEnabled(QtObjectPtr ptr);
void QDialog_SetModal(QtObjectPtr ptr, int modal);
void QDialog_SetResult(QtObjectPtr ptr, int i);
void QDialog_SetSizeGripEnabled(QtObjectPtr ptr, int v);
QtObjectPtr QDialog_NewQDialog(QtObjectPtr parent, int f);
void QDialog_Accept(QtObjectPtr ptr);
void QDialog_ConnectAccepted(QtObjectPtr ptr);
void QDialog_DisconnectAccepted(QtObjectPtr ptr);
void QDialog_Done(QtObjectPtr ptr, int r);
int QDialog_Exec(QtObjectPtr ptr);
void QDialog_ConnectFinished(QtObjectPtr ptr);
void QDialog_DisconnectFinished(QtObjectPtr ptr);
void QDialog_Open(QtObjectPtr ptr);
void QDialog_Reject(QtObjectPtr ptr);
void QDialog_ConnectRejected(QtObjectPtr ptr);
void QDialog_DisconnectRejected(QtObjectPtr ptr);
int QDialog_Result(QtObjectPtr ptr);
void QDialog_SetVisible(QtObjectPtr ptr, int visible);
void QDialog_DestroyQDialog(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif