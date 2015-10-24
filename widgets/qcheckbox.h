#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QCheckBox_IsTristate(QtObjectPtr ptr);
void QCheckBox_SetTristate(QtObjectPtr ptr, int y);
QtObjectPtr QCheckBox_NewQCheckBox(QtObjectPtr parent);
QtObjectPtr QCheckBox_NewQCheckBox2(char* text, QtObjectPtr parent);
int QCheckBox_CheckState(QtObjectPtr ptr);
void QCheckBox_SetCheckState(QtObjectPtr ptr, int state);
void QCheckBox_ConnectStateChanged(QtObjectPtr ptr);
void QCheckBox_DisconnectStateChanged(QtObjectPtr ptr);
void QCheckBox_DestroyQCheckBox(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif