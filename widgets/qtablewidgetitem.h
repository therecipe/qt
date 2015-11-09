#ifdef __cplusplus
extern "C" {
#endif

void QTableWidgetItem_SetFlags(void* ptr, int flags);
void* QTableWidgetItem_NewQTableWidgetItem3(void* icon, char* text, int ty);
void* QTableWidgetItem_NewQTableWidgetItem2(char* text, int ty);
void* QTableWidgetItem_NewQTableWidgetItem4(void* other);
void* QTableWidgetItem_NewQTableWidgetItem(int ty);
void* QTableWidgetItem_Background(void* ptr);
int QTableWidgetItem_CheckState(void* ptr);
void* QTableWidgetItem_Clone(void* ptr);
int QTableWidgetItem_Column(void* ptr);
void* QTableWidgetItem_Data(void* ptr, int role);
int QTableWidgetItem_Flags(void* ptr);
void* QTableWidgetItem_Foreground(void* ptr);
int QTableWidgetItem_IsSelected(void* ptr);
void QTableWidgetItem_Read(void* ptr, void* in);
int QTableWidgetItem_Row(void* ptr);
void QTableWidgetItem_SetBackground(void* ptr, void* brush);
void QTableWidgetItem_SetCheckState(void* ptr, int state);
void QTableWidgetItem_SetData(void* ptr, int role, void* value);
void QTableWidgetItem_SetFont(void* ptr, void* font);
void QTableWidgetItem_SetForeground(void* ptr, void* brush);
void QTableWidgetItem_SetIcon(void* ptr, void* icon);
void QTableWidgetItem_SetSelected(void* ptr, int sele);
void QTableWidgetItem_SetSizeHint(void* ptr, void* size);
void QTableWidgetItem_SetStatusTip(void* ptr, char* statusTip);
void QTableWidgetItem_SetText(void* ptr, char* text);
void QTableWidgetItem_SetTextAlignment(void* ptr, int alignment);
void QTableWidgetItem_SetToolTip(void* ptr, char* toolTip);
void QTableWidgetItem_SetWhatsThis(void* ptr, char* whatsThis);
char* QTableWidgetItem_StatusTip(void* ptr);
void* QTableWidgetItem_TableWidget(void* ptr);
char* QTableWidgetItem_Text(void* ptr);
int QTableWidgetItem_TextAlignment(void* ptr);
char* QTableWidgetItem_ToolTip(void* ptr);
int QTableWidgetItem_Type(void* ptr);
char* QTableWidgetItem_WhatsThis(void* ptr);
void QTableWidgetItem_Write(void* ptr, void* out);
void QTableWidgetItem_DestroyQTableWidgetItem(void* ptr);

#ifdef __cplusplus
}
#endif