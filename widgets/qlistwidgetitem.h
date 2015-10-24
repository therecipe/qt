#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QListWidgetItem_NewQListWidgetItem(QtObjectPtr parent, int ty);
QtObjectPtr QListWidgetItem_NewQListWidgetItem3(QtObjectPtr icon, char* text, QtObjectPtr parent, int ty);
QtObjectPtr QListWidgetItem_NewQListWidgetItem2(char* text, QtObjectPtr parent, int ty);
void QListWidgetItem_SetFlags(QtObjectPtr ptr, int flags);
QtObjectPtr QListWidgetItem_NewQListWidgetItem4(QtObjectPtr other);
int QListWidgetItem_CheckState(QtObjectPtr ptr);
QtObjectPtr QListWidgetItem_Clone(QtObjectPtr ptr);
char* QListWidgetItem_Data(QtObjectPtr ptr, int role);
int QListWidgetItem_Flags(QtObjectPtr ptr);
int QListWidgetItem_IsHidden(QtObjectPtr ptr);
int QListWidgetItem_IsSelected(QtObjectPtr ptr);
QtObjectPtr QListWidgetItem_ListWidget(QtObjectPtr ptr);
void QListWidgetItem_Read(QtObjectPtr ptr, QtObjectPtr in);
void QListWidgetItem_SetBackground(QtObjectPtr ptr, QtObjectPtr brush);
void QListWidgetItem_SetCheckState(QtObjectPtr ptr, int state);
void QListWidgetItem_SetData(QtObjectPtr ptr, int role, char* value);
void QListWidgetItem_SetFont(QtObjectPtr ptr, QtObjectPtr font);
void QListWidgetItem_SetForeground(QtObjectPtr ptr, QtObjectPtr brush);
void QListWidgetItem_SetHidden(QtObjectPtr ptr, int hide);
void QListWidgetItem_SetIcon(QtObjectPtr ptr, QtObjectPtr icon);
void QListWidgetItem_SetSelected(QtObjectPtr ptr, int sele);
void QListWidgetItem_SetSizeHint(QtObjectPtr ptr, QtObjectPtr size);
void QListWidgetItem_SetStatusTip(QtObjectPtr ptr, char* statusTip);
void QListWidgetItem_SetText(QtObjectPtr ptr, char* text);
void QListWidgetItem_SetTextAlignment(QtObjectPtr ptr, int alignment);
void QListWidgetItem_SetToolTip(QtObjectPtr ptr, char* toolTip);
void QListWidgetItem_SetWhatsThis(QtObjectPtr ptr, char* whatsThis);
char* QListWidgetItem_StatusTip(QtObjectPtr ptr);
char* QListWidgetItem_Text(QtObjectPtr ptr);
int QListWidgetItem_TextAlignment(QtObjectPtr ptr);
char* QListWidgetItem_ToolTip(QtObjectPtr ptr);
int QListWidgetItem_Type(QtObjectPtr ptr);
char* QListWidgetItem_WhatsThis(QtObjectPtr ptr);
void QListWidgetItem_Write(QtObjectPtr ptr, QtObjectPtr out);
void QListWidgetItem_DestroyQListWidgetItem(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif