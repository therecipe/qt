#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QTableWidgetItem_SetFlags(QtObjectPtr ptr, int flags);
QtObjectPtr QTableWidgetItem_NewQTableWidgetItem3(QtObjectPtr icon, char* text, int ty);
QtObjectPtr QTableWidgetItem_NewQTableWidgetItem2(char* text, int ty);
QtObjectPtr QTableWidgetItem_NewQTableWidgetItem4(QtObjectPtr other);
QtObjectPtr QTableWidgetItem_NewQTableWidgetItem(int ty);
int QTableWidgetItem_CheckState(QtObjectPtr ptr);
QtObjectPtr QTableWidgetItem_Clone(QtObjectPtr ptr);
int QTableWidgetItem_Column(QtObjectPtr ptr);
char* QTableWidgetItem_Data(QtObjectPtr ptr, int role);
int QTableWidgetItem_Flags(QtObjectPtr ptr);
int QTableWidgetItem_IsSelected(QtObjectPtr ptr);
void QTableWidgetItem_Read(QtObjectPtr ptr, QtObjectPtr in);
int QTableWidgetItem_Row(QtObjectPtr ptr);
void QTableWidgetItem_SetBackground(QtObjectPtr ptr, QtObjectPtr brush);
void QTableWidgetItem_SetCheckState(QtObjectPtr ptr, int state);
void QTableWidgetItem_SetData(QtObjectPtr ptr, int role, char* value);
void QTableWidgetItem_SetFont(QtObjectPtr ptr, QtObjectPtr font);
void QTableWidgetItem_SetForeground(QtObjectPtr ptr, QtObjectPtr brush);
void QTableWidgetItem_SetIcon(QtObjectPtr ptr, QtObjectPtr icon);
void QTableWidgetItem_SetSelected(QtObjectPtr ptr, int sele);
void QTableWidgetItem_SetSizeHint(QtObjectPtr ptr, QtObjectPtr size);
void QTableWidgetItem_SetStatusTip(QtObjectPtr ptr, char* statusTip);
void QTableWidgetItem_SetText(QtObjectPtr ptr, char* text);
void QTableWidgetItem_SetTextAlignment(QtObjectPtr ptr, int alignment);
void QTableWidgetItem_SetToolTip(QtObjectPtr ptr, char* toolTip);
void QTableWidgetItem_SetWhatsThis(QtObjectPtr ptr, char* whatsThis);
char* QTableWidgetItem_StatusTip(QtObjectPtr ptr);
QtObjectPtr QTableWidgetItem_TableWidget(QtObjectPtr ptr);
char* QTableWidgetItem_Text(QtObjectPtr ptr);
int QTableWidgetItem_TextAlignment(QtObjectPtr ptr);
char* QTableWidgetItem_ToolTip(QtObjectPtr ptr);
int QTableWidgetItem_Type(QtObjectPtr ptr);
char* QTableWidgetItem_WhatsThis(QtObjectPtr ptr);
void QTableWidgetItem_Write(QtObjectPtr ptr, QtObjectPtr out);
void QTableWidgetItem_DestroyQTableWidgetItem(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif