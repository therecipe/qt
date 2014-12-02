#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"

//Public Functions
QtObjectPtr QDialog_New_QWidget_WindowType(QtObjectPtr parent, int f);
void QDialog_Destroy(QtObjectPtr ptr);
int QDialog_IsSizeGripEnabled(QtObjectPtr ptr);
int QDialog_Result(QtObjectPtr ptr);
void QDialog_SetModal_Bool(QtObjectPtr ptr, int modal);
void QDialog_SetResult_Int(QtObjectPtr ptr, int i);
void QDialog_SetSizeGripEnabled_Bool(QtObjectPtr ptr, int sizeGripEnabled);
//Public Slots
void QDialog_ConnectSlotAccept(QtObjectPtr ptr);
void QDialog_DisconnectSlotAccept(QtObjectPtr ptr);
void QDialog_Accept(QtObjectPtr ptr);
void QDialog_ConnectSlotDone(QtObjectPtr ptr);
void QDialog_DisconnectSlotDone(QtObjectPtr ptr);
void QDialog_Done_Int(QtObjectPtr ptr, int r);
void QDialog_ConnectSlotExec(QtObjectPtr ptr);
void QDialog_DisconnectSlotExec(QtObjectPtr ptr);
void QDialog_Exec(QtObjectPtr ptr);
void QDialog_ConnectSlotOpen(QtObjectPtr ptr);
void QDialog_DisconnectSlotOpen(QtObjectPtr ptr);
void QDialog_Open(QtObjectPtr ptr);
void QDialog_ConnectSlotReject(QtObjectPtr ptr);
void QDialog_DisconnectSlotReject(QtObjectPtr ptr);
void QDialog_Reject(QtObjectPtr ptr);
//Signals
void QDialog_ConnectSignalAccepted(QtObjectPtr ptr);
void QDialog_DisconnectSignalAccepted(QtObjectPtr ptr);
void QDialog_ConnectSignalFinished(QtObjectPtr ptr);
void QDialog_DisconnectSignalFinished(QtObjectPtr ptr);
void QDialog_ConnectSignalRejected(QtObjectPtr ptr);
void QDialog_DisconnectSignalRejected(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif
