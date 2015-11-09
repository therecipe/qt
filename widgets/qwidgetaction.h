#ifdef __cplusplus
extern "C" {
#endif

void* QWidgetAction_NewQWidgetAction(void* parent);
void* QWidgetAction_DefaultWidget(void* ptr);
void QWidgetAction_ReleaseWidget(void* ptr, void* widget);
void* QWidgetAction_RequestWidget(void* ptr, void* parent);
void QWidgetAction_SetDefaultWidget(void* ptr, void* widget);
void QWidgetAction_DestroyQWidgetAction(void* ptr);

#ifdef __cplusplus
}
#endif