#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"

//Public Functions
QtObjectPtr QToolBox_New_QWidget_WindowType(QtObjectPtr parent, int f);
void QToolBox_Destroy(QtObjectPtr ptr);
int QToolBox_AddItem_QWidget_String(QtObjectPtr ptr, QtObjectPtr w, char* text);
int QToolBox_Count(QtObjectPtr ptr);
int QToolBox_CurrentIndex(QtObjectPtr ptr);
QtObjectPtr QToolBox_CurrentWidget(QtObjectPtr ptr);
int QToolBox_IndexOf_QWidget(QtObjectPtr ptr, QtObjectPtr widget);
int QToolBox_InsertItem_Int_QWidget_String(QtObjectPtr ptr, int index, QtObjectPtr widget, char* text);
int QToolBox_IsItemEnabled_Int(QtObjectPtr ptr, int index);
char* QToolBox_ItemText_Int(QtObjectPtr ptr, int index);
char* QToolBox_ItemToolTip_Int(QtObjectPtr ptr, int index);
void QToolBox_RemoveItem_Int(QtObjectPtr ptr, int index);
void QToolBox_SetItemEnabled_Int_Bool(QtObjectPtr ptr, int index, int enabled);
void QToolBox_SetItemText_Int_String(QtObjectPtr ptr, int index, char* text);
void QToolBox_SetItemToolTip_Int_String(QtObjectPtr ptr, int index, char* toolTip);
QtObjectPtr QToolBox_Widget_Int(QtObjectPtr ptr, int index);
//Public Slots
void QToolBox_ConnectSlotSetCurrentIndex(QtObjectPtr ptr);
void QToolBox_DisconnectSlotSetCurrentIndex(QtObjectPtr ptr);
void QToolBox_SetCurrentIndex_Int(QtObjectPtr ptr, int index);
//Signals
void QToolBox_ConnectSignalCurrentChanged(QtObjectPtr ptr);
void QToolBox_DisconnectSignalCurrentChanged(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif
