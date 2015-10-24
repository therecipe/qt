#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QGroupBox_Alignment(QtObjectPtr ptr);
int QGroupBox_IsCheckable(QtObjectPtr ptr);
int QGroupBox_IsChecked(QtObjectPtr ptr);
int QGroupBox_IsFlat(QtObjectPtr ptr);
void QGroupBox_SetAlignment(QtObjectPtr ptr, int alignment);
void QGroupBox_SetCheckable(QtObjectPtr ptr, int checkable);
void QGroupBox_SetChecked(QtObjectPtr ptr, int checked);
void QGroupBox_SetFlat(QtObjectPtr ptr, int flat);
void QGroupBox_SetTitle(QtObjectPtr ptr, char* title);
char* QGroupBox_Title(QtObjectPtr ptr);
QtObjectPtr QGroupBox_NewQGroupBox(QtObjectPtr parent);
QtObjectPtr QGroupBox_NewQGroupBox2(char* title, QtObjectPtr parent);
void QGroupBox_ConnectClicked(QtObjectPtr ptr);
void QGroupBox_DisconnectClicked(QtObjectPtr ptr);
void QGroupBox_ConnectToggled(QtObjectPtr ptr);
void QGroupBox_DisconnectToggled(QtObjectPtr ptr);
void QGroupBox_DestroyQGroupBox(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif