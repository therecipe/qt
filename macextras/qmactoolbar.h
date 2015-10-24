#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QMacToolBar_NewQMacToolBar(QtObjectPtr parent);
QtObjectPtr QMacToolBar_NewQMacToolBar2(char* identifier, QtObjectPtr parent);
QtObjectPtr QMacToolBar_AddAllowedItem(QtObjectPtr ptr, QtObjectPtr icon, char* text);
QtObjectPtr QMacToolBar_AddItem(QtObjectPtr ptr, QtObjectPtr icon, char* text);
void QMacToolBar_AddSeparator(QtObjectPtr ptr);
void QMacToolBar_AttachToWindow(QtObjectPtr ptr, QtObjectPtr window);
void QMacToolBar_DetachFromWindow(QtObjectPtr ptr);
void QMacToolBar_DestroyQMacToolBar(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif