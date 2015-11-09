#ifdef __cplusplus
extern "C" {
#endif

void* QItemEditorFactory_NewQItemEditorFactory();
void* QItemEditorFactory_CreateEditor(void* ptr, int userType, void* parent);
void* QItemEditorFactory_QItemEditorFactory_DefaultFactory();
void QItemEditorFactory_RegisterEditor(void* ptr, int userType, void* creator);
void QItemEditorFactory_QItemEditorFactory_SetDefaultFactory(void* factory);
void* QItemEditorFactory_ValuePropertyName(void* ptr, int userType);
void QItemEditorFactory_DestroyQItemEditorFactory(void* ptr);

#ifdef __cplusplus
}
#endif