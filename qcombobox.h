#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"

//Public Functions
QtObjectPtr QComboBox_New_QWidget(QtObjectPtr parent);
void QComboBox_Destroy(QtObjectPtr ptr);
void QComboBox_AddItems_QStringList(QtObjectPtr ptr, char* texts);
int QComboBox_Count(QtObjectPtr ptr);
int QComboBox_CurrentIndex(QtObjectPtr ptr);
char* QComboBox_CurrentText(QtObjectPtr ptr);
int QComboBox_DuplicatesEnabled(QtObjectPtr ptr);
int QComboBox_HasFrame(QtObjectPtr ptr);
void QComboBox_HidePopup(QtObjectPtr ptr);
void QComboBox_InsertSeparator_Int(QtObjectPtr ptr, int index);
int QComboBox_IsEditable(QtObjectPtr ptr);
char* QComboBox_ItemText_Int(QtObjectPtr ptr, int index);
QtObjectPtr QComboBox_LineEdit(QtObjectPtr ptr);
int QComboBox_MaxCount(QtObjectPtr ptr);
int QComboBox_MaxVisibleItems(QtObjectPtr ptr);
int QComboBox_MinimumContentsLength(QtObjectPtr ptr);
int QComboBox_ModelColumn(QtObjectPtr ptr);
void QComboBox_RemoveItem_Int(QtObjectPtr ptr, int index);
void QComboBox_SetDuplicatesEnabled_Bool(QtObjectPtr ptr, int enable);
void QComboBox_SetEditable_Bool(QtObjectPtr ptr, int editable);
void QComboBox_SetFrame_Bool(QtObjectPtr ptr, int frame);
void QComboBox_SetItemText_Int_String(QtObjectPtr ptr, int index, char* text);
void QComboBox_SetLineEdit_QLineEdit(QtObjectPtr ptr, QtObjectPtr edit);
void QComboBox_SetMaxCount_Int(QtObjectPtr ptr, int max);
void QComboBox_SetMaxVisibleItems_Int(QtObjectPtr ptr, int maxItems);
void QComboBox_SetMinimumContentsLength_Int(QtObjectPtr ptr, int characters);
void QComboBox_SetModelColumn_Int(QtObjectPtr ptr, int visibleColumn);
void QComboBox_ShowPopup(QtObjectPtr ptr);
//Public Slots
void QComboBox_ConnectSlotClear(QtObjectPtr ptr);
void QComboBox_DisconnectSlotClear(QtObjectPtr ptr);
void QComboBox_Clear(QtObjectPtr ptr);
void QComboBox_ConnectSlotClearEditText(QtObjectPtr ptr);
void QComboBox_DisconnectSlotClearEditText(QtObjectPtr ptr);
void QComboBox_ClearEditText(QtObjectPtr ptr);
void QComboBox_ConnectSlotSetCurrentIndex(QtObjectPtr ptr);
void QComboBox_DisconnectSlotSetCurrentIndex(QtObjectPtr ptr);
void QComboBox_SetCurrentIndex_Int(QtObjectPtr ptr, int index);
//Signals
void QComboBox_ConnectSignalCurrentTextChanged(QtObjectPtr ptr);
void QComboBox_DisconnectSignalCurrentTextChanged(QtObjectPtr ptr);
void QComboBox_ConnectSignalEditTextChanged(QtObjectPtr ptr);
void QComboBox_DisconnectSignalEditTextChanged(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif
