#ifdef __cplusplus
extern "C" {
#endif

int QScrollArea_Alignment(void* ptr);
void QScrollArea_SetAlignment(void* ptr, int v);
void QScrollArea_SetWidget(void* ptr, void* widget);
void QScrollArea_SetWidgetResizable(void* ptr, int resizable);
int QScrollArea_WidgetResizable(void* ptr);
void* QScrollArea_NewQScrollArea(void* parent);
void QScrollArea_EnsureVisible(void* ptr, int x, int y, int xmargin, int ymargin);
void QScrollArea_EnsureWidgetVisible(void* ptr, void* childWidget, int xmargin, int ymargin);
int QScrollArea_FocusNextPrevChild(void* ptr, int next);
void* QScrollArea_TakeWidget(void* ptr);
void* QScrollArea_Widget(void* ptr);
void QScrollArea_DestroyQScrollArea(void* ptr);

#ifdef __cplusplus
}
#endif