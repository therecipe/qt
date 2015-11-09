#ifdef __cplusplus
extern "C" {
#endif

void* QListWidgetItem_NewQListWidgetItem(void* parent, int ty);
void* QListWidgetItem_NewQListWidgetItem3(void* icon, char* text, void* parent, int ty);
void* QListWidgetItem_NewQListWidgetItem2(char* text, void* parent, int ty);
void QListWidgetItem_SetFlags(void* ptr, int flags);
void* QListWidgetItem_NewQListWidgetItem4(void* other);
void* QListWidgetItem_Background(void* ptr);
int QListWidgetItem_CheckState(void* ptr);
void* QListWidgetItem_Clone(void* ptr);
void* QListWidgetItem_Data(void* ptr, int role);
int QListWidgetItem_Flags(void* ptr);
void* QListWidgetItem_Foreground(void* ptr);
int QListWidgetItem_IsHidden(void* ptr);
int QListWidgetItem_IsSelected(void* ptr);
void* QListWidgetItem_ListWidget(void* ptr);
void QListWidgetItem_Read(void* ptr, void* in);
void QListWidgetItem_SetBackground(void* ptr, void* brush);
void QListWidgetItem_SetCheckState(void* ptr, int state);
void QListWidgetItem_SetData(void* ptr, int role, void* value);
void QListWidgetItem_SetFont(void* ptr, void* font);
void QListWidgetItem_SetForeground(void* ptr, void* brush);
void QListWidgetItem_SetHidden(void* ptr, int hide);
void QListWidgetItem_SetIcon(void* ptr, void* icon);
void QListWidgetItem_SetSelected(void* ptr, int sele);
void QListWidgetItem_SetSizeHint(void* ptr, void* size);
void QListWidgetItem_SetStatusTip(void* ptr, char* statusTip);
void QListWidgetItem_SetText(void* ptr, char* text);
void QListWidgetItem_SetTextAlignment(void* ptr, int alignment);
void QListWidgetItem_SetToolTip(void* ptr, char* toolTip);
void QListWidgetItem_SetWhatsThis(void* ptr, char* whatsThis);
char* QListWidgetItem_StatusTip(void* ptr);
char* QListWidgetItem_Text(void* ptr);
int QListWidgetItem_TextAlignment(void* ptr);
char* QListWidgetItem_ToolTip(void* ptr);
int QListWidgetItem_Type(void* ptr);
char* QListWidgetItem_WhatsThis(void* ptr);
void QListWidgetItem_Write(void* ptr, void* out);
void QListWidgetItem_DestroyQListWidgetItem(void* ptr);

#ifdef __cplusplus
}
#endif