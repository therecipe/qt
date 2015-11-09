#ifdef __cplusplus
extern "C" {
#endif

int QToolBox_Count(void* ptr);
int QToolBox_CurrentIndex(void* ptr);
void QToolBox_SetCurrentIndex(void* ptr, int index);
void* QToolBox_NewQToolBox(void* parent, int f);
int QToolBox_AddItem2(void* ptr, void* w, char* text);
int QToolBox_AddItem(void* ptr, void* widget, void* iconSet, char* text);
void QToolBox_ConnectCurrentChanged(void* ptr);
void QToolBox_DisconnectCurrentChanged(void* ptr);
void* QToolBox_CurrentWidget(void* ptr);
int QToolBox_IndexOf(void* ptr, void* widget);
int QToolBox_InsertItem(void* ptr, int index, void* widget, void* icon, char* text);
int QToolBox_InsertItem2(void* ptr, int index, void* widget, char* text);
int QToolBox_IsItemEnabled(void* ptr, int index);
char* QToolBox_ItemText(void* ptr, int index);
char* QToolBox_ItemToolTip(void* ptr, int index);
void QToolBox_RemoveItem(void* ptr, int index);
void QToolBox_SetCurrentWidget(void* ptr, void* widget);
void QToolBox_SetItemEnabled(void* ptr, int index, int enabled);
void QToolBox_SetItemIcon(void* ptr, int index, void* icon);
void QToolBox_SetItemText(void* ptr, int index, char* text);
void QToolBox_SetItemToolTip(void* ptr, int index, char* toolTip);
void* QToolBox_Widget(void* ptr, int index);
void QToolBox_DestroyQToolBox(void* ptr);

#ifdef __cplusplus
}
#endif