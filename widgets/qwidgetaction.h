#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QWidgetAction_NewQWidgetAction(QtObjectPtr parent);
QtObjectPtr QWidgetAction_DefaultWidget(QtObjectPtr ptr);
void QWidgetAction_ReleaseWidget(QtObjectPtr ptr, QtObjectPtr widget);
QtObjectPtr QWidgetAction_RequestWidget(QtObjectPtr ptr, QtObjectPtr parent);
void QWidgetAction_SetDefaultWidget(QtObjectPtr ptr, QtObjectPtr widget);
void QWidgetAction_DestroyQWidgetAction(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif