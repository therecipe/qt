#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QScrollArea_Alignment(QtObjectPtr ptr);
void QScrollArea_SetAlignment(QtObjectPtr ptr, int v);
void QScrollArea_SetWidget(QtObjectPtr ptr, QtObjectPtr widget);
void QScrollArea_SetWidgetResizable(QtObjectPtr ptr, int resizable);
int QScrollArea_WidgetResizable(QtObjectPtr ptr);
QtObjectPtr QScrollArea_NewQScrollArea(QtObjectPtr parent);
void QScrollArea_EnsureVisible(QtObjectPtr ptr, int x, int y, int xmargin, int ymargin);
void QScrollArea_EnsureWidgetVisible(QtObjectPtr ptr, QtObjectPtr childWidget, int xmargin, int ymargin);
int QScrollArea_FocusNextPrevChild(QtObjectPtr ptr, int next);
QtObjectPtr QScrollArea_TakeWidget(QtObjectPtr ptr);
QtObjectPtr QScrollArea_Widget(QtObjectPtr ptr);
void QScrollArea_DestroyQScrollArea(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif