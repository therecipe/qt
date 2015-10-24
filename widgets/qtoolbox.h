#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QToolBox_Count(QtObjectPtr ptr);
int QToolBox_CurrentIndex(QtObjectPtr ptr);
void QToolBox_SetCurrentIndex(QtObjectPtr ptr, int index);
QtObjectPtr QToolBox_NewQToolBox(QtObjectPtr parent, int f);
int QToolBox_AddItem2(QtObjectPtr ptr, QtObjectPtr w, char* text);
int QToolBox_AddItem(QtObjectPtr ptr, QtObjectPtr widget, QtObjectPtr iconSet, char* text);
void QToolBox_ConnectCurrentChanged(QtObjectPtr ptr);
void QToolBox_DisconnectCurrentChanged(QtObjectPtr ptr);
QtObjectPtr QToolBox_CurrentWidget(QtObjectPtr ptr);
int QToolBox_IndexOf(QtObjectPtr ptr, QtObjectPtr widget);
int QToolBox_InsertItem(QtObjectPtr ptr, int index, QtObjectPtr widget, QtObjectPtr icon, char* text);
int QToolBox_InsertItem2(QtObjectPtr ptr, int index, QtObjectPtr widget, char* text);
int QToolBox_IsItemEnabled(QtObjectPtr ptr, int index);
char* QToolBox_ItemText(QtObjectPtr ptr, int index);
char* QToolBox_ItemToolTip(QtObjectPtr ptr, int index);
void QToolBox_RemoveItem(QtObjectPtr ptr, int index);
void QToolBox_SetCurrentWidget(QtObjectPtr ptr, QtObjectPtr widget);
void QToolBox_SetItemEnabled(QtObjectPtr ptr, int index, int enabled);
void QToolBox_SetItemIcon(QtObjectPtr ptr, int index, QtObjectPtr icon);
void QToolBox_SetItemText(QtObjectPtr ptr, int index, char* text);
void QToolBox_SetItemToolTip(QtObjectPtr ptr, int index, char* toolTip);
QtObjectPtr QToolBox_Widget(QtObjectPtr ptr, int index);
void QToolBox_DestroyQToolBox(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif