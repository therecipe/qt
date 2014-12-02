#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"

//Public Functions
QtObjectPtr QListWidgetItem_New_QListWidget_Int(QtObjectPtr parent, int typ);
QtObjectPtr QListWidgetItem_New_String_QListWidget_Int(char* text, QtObjectPtr parent, int typ);
void QListWidgetItem_Destroy(QtObjectPtr ptr);
int QListWidgetItem_CheckState(QtObjectPtr ptr);
QtObjectPtr QListWidgetItem_Clone(QtObjectPtr ptr);
int QListWidgetItem_Flags(QtObjectPtr ptr);
int QListWidgetItem_IsHidden(QtObjectPtr ptr);
int QListWidgetItem_IsSelected(QtObjectPtr ptr);
QtObjectPtr QListWidgetItem_ListWidget(QtObjectPtr ptr);
void QListWidgetItem_SetCheckState_CheckState(QtObjectPtr ptr, int state);
void QListWidgetItem_SetFlags_ItemFlag(QtObjectPtr ptr, int flags);
void QListWidgetItem_SetHidden_Bool(QtObjectPtr ptr, int hide);
void QListWidgetItem_SetSelected_Bool(QtObjectPtr ptr, int selected);
void QListWidgetItem_SetStatusTip_String(QtObjectPtr ptr, char* statusTip);
void QListWidgetItem_SetText_String(QtObjectPtr ptr, char* text);
void QListWidgetItem_SetTextAlignment_Int(QtObjectPtr ptr, int alignment);
void QListWidgetItem_SetToolTip_String(QtObjectPtr ptr, char* toolTip);
void QListWidgetItem_SetWhatsThis_String(QtObjectPtr ptr, char* whatsThis);
char* QListWidgetItem_StatusTip(QtObjectPtr ptr);
char* QListWidgetItem_Text(QtObjectPtr ptr);
int QListWidgetItem_TextAlignment(QtObjectPtr ptr);
char* QListWidgetItem_ToolTip(QtObjectPtr ptr);
int QListWidgetItem_Type(QtObjectPtr ptr);
char* QListWidgetItem_WhatsThis(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif
