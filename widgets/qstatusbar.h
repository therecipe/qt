#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QStatusBar_IsSizeGripEnabled(QtObjectPtr ptr);
void QStatusBar_SetSizeGripEnabled(QtObjectPtr ptr, int v);
QtObjectPtr QStatusBar_NewQStatusBar(QtObjectPtr parent);
void QStatusBar_AddPermanentWidget(QtObjectPtr ptr, QtObjectPtr widget, int stretch);
void QStatusBar_AddWidget(QtObjectPtr ptr, QtObjectPtr widget, int stretch);
void QStatusBar_ClearMessage(QtObjectPtr ptr);
char* QStatusBar_CurrentMessage(QtObjectPtr ptr);
int QStatusBar_InsertPermanentWidget(QtObjectPtr ptr, int index, QtObjectPtr widget, int stretch);
int QStatusBar_InsertWidget(QtObjectPtr ptr, int index, QtObjectPtr widget, int stretch);
void QStatusBar_ConnectMessageChanged(QtObjectPtr ptr);
void QStatusBar_DisconnectMessageChanged(QtObjectPtr ptr);
void QStatusBar_RemoveWidget(QtObjectPtr ptr, QtObjectPtr widget);
void QStatusBar_ShowMessage(QtObjectPtr ptr, char* message, int timeout);
void QStatusBar_DestroyQStatusBar(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif