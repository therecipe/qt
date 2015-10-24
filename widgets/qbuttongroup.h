#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QButtonGroup_NewQButtonGroup(QtObjectPtr parent);
void QButtonGroup_AddButton(QtObjectPtr ptr, QtObjectPtr button, int id);
QtObjectPtr QButtonGroup_Button(QtObjectPtr ptr, int id);
QtObjectPtr QButtonGroup_CheckedButton(QtObjectPtr ptr);
int QButtonGroup_CheckedId(QtObjectPtr ptr);
int QButtonGroup_Exclusive(QtObjectPtr ptr);
int QButtonGroup_Id(QtObjectPtr ptr, QtObjectPtr button);
void QButtonGroup_RemoveButton(QtObjectPtr ptr, QtObjectPtr button);
void QButtonGroup_SetExclusive(QtObjectPtr ptr, int v);
void QButtonGroup_SetId(QtObjectPtr ptr, QtObjectPtr button, int id);
void QButtonGroup_DestroyQButtonGroup(QtObjectPtr ptr);
void QButtonGroup_ConnectButtonClicked(QtObjectPtr ptr);
void QButtonGroup_DisconnectButtonClicked(QtObjectPtr ptr);
void QButtonGroup_ConnectButtonPressed(QtObjectPtr ptr);
void QButtonGroup_DisconnectButtonPressed(QtObjectPtr ptr);
void QButtonGroup_ConnectButtonReleased(QtObjectPtr ptr);
void QButtonGroup_DisconnectButtonReleased(QtObjectPtr ptr);
void QButtonGroup_ConnectButtonToggled(QtObjectPtr ptr);
void QButtonGroup_DisconnectButtonToggled(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif