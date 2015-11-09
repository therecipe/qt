#ifdef __cplusplus
extern "C" {
#endif

int QDataWidgetMapper_CurrentIndex(void* ptr);
int QDataWidgetMapper_Orientation(void* ptr);
void QDataWidgetMapper_SetCurrentIndex(void* ptr, int index);
void QDataWidgetMapper_SetOrientation(void* ptr, int aOrientation);
void QDataWidgetMapper_SetSubmitPolicy(void* ptr, int policy);
int QDataWidgetMapper_SubmitPolicy(void* ptr);
void* QDataWidgetMapper_NewQDataWidgetMapper(void* parent);
void QDataWidgetMapper_AddMapping(void* ptr, void* widget, int section);
void QDataWidgetMapper_AddMapping2(void* ptr, void* widget, int section, void* propertyName);
void QDataWidgetMapper_ClearMapping(void* ptr);
void QDataWidgetMapper_ConnectCurrentIndexChanged(void* ptr);
void QDataWidgetMapper_DisconnectCurrentIndexChanged(void* ptr);
void* QDataWidgetMapper_ItemDelegate(void* ptr);
void* QDataWidgetMapper_MappedPropertyName(void* ptr, void* widget);
int QDataWidgetMapper_MappedSection(void* ptr, void* widget);
void* QDataWidgetMapper_MappedWidgetAt(void* ptr, int section);
void* QDataWidgetMapper_Model(void* ptr);
void QDataWidgetMapper_RemoveMapping(void* ptr, void* widget);
void QDataWidgetMapper_Revert(void* ptr);
void* QDataWidgetMapper_RootIndex(void* ptr);
void QDataWidgetMapper_SetCurrentModelIndex(void* ptr, void* index);
void QDataWidgetMapper_SetItemDelegate(void* ptr, void* delegate);
void QDataWidgetMapper_SetModel(void* ptr, void* model);
void QDataWidgetMapper_SetRootIndex(void* ptr, void* index);
int QDataWidgetMapper_Submit(void* ptr);
void QDataWidgetMapper_ToFirst(void* ptr);
void QDataWidgetMapper_ToLast(void* ptr);
void QDataWidgetMapper_ToNext(void* ptr);
void QDataWidgetMapper_ToPrevious(void* ptr);
void QDataWidgetMapper_DestroyQDataWidgetMapper(void* ptr);

#ifdef __cplusplus
}
#endif