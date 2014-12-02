#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"

//Public Functions
QtObjectPtr QCheckBox_New_QWidget(QtObjectPtr parent);
QtObjectPtr QCheckBox_New_String_QWidget(char* text, QtObjectPtr parent);
void QCheckBox_Destroy(QtObjectPtr ptr);
int QCheckBox_CheckState(QtObjectPtr ptr);
int QCheckBox_IsTristate(QtObjectPtr ptr);
void QCheckBox_SetCheckState_CheckState(QtObjectPtr ptr, int state);
void QCheckBox_SetTristate_Bool(QtObjectPtr ptr, int y);
//Signals
void QCheckBox_ConnectSignalStateChanged(QtObjectPtr ptr);
void QCheckBox_DisconnectSignalStateChanged(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif
