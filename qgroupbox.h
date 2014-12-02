#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"

//Public Functions
QtObjectPtr QGroupBox_New_QWidget(QtObjectPtr parent);
QtObjectPtr QGroupBox_New_String_QWidget(char* title, QtObjectPtr parent);
void QGroupBox_Destroy(QtObjectPtr ptr);
int QGroupBox_Alignment(QtObjectPtr ptr);
int QGroupBox_IsCheckable(QtObjectPtr ptr);
int QGroupBox_IsChecked(QtObjectPtr ptr);
int QGroupBox_IsFlat(QtObjectPtr ptr);
void QGroupBox_SetAlignment_Int(QtObjectPtr ptr, int alignment);
void QGroupBox_SetCheckable_Bool(QtObjectPtr ptr, int checkable);
void QGroupBox_SetFlat_Bool(QtObjectPtr ptr, int flat);
void QGroupBox_SetTitle_String(QtObjectPtr ptr, char* title);
char* QGroupBox_Title(QtObjectPtr ptr);
//Public Slots
void QGroupBox_ConnectSlotSetChecked(QtObjectPtr ptr);
void QGroupBox_DisconnectSlotSetChecked(QtObjectPtr ptr);
void QGroupBox_SetChecked_Bool(QtObjectPtr ptr, int checked);
//Signals
void QGroupBox_ConnectSignalClicked(QtObjectPtr ptr);
void QGroupBox_DisconnectSignalClicked(QtObjectPtr ptr);
void QGroupBox_ConnectSignalToggled(QtObjectPtr ptr);
void QGroupBox_DisconnectSignalToggled(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif
