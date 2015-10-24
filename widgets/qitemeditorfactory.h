#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QItemEditorFactory_NewQItemEditorFactory();
QtObjectPtr QItemEditorFactory_CreateEditor(QtObjectPtr ptr, int userType, QtObjectPtr parent);
QtObjectPtr QItemEditorFactory_QItemEditorFactory_DefaultFactory();
void QItemEditorFactory_RegisterEditor(QtObjectPtr ptr, int userType, QtObjectPtr creator);
void QItemEditorFactory_QItemEditorFactory_SetDefaultFactory(QtObjectPtr factory);
void QItemEditorFactory_DestroyQItemEditorFactory(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif