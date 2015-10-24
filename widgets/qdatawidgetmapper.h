#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QDataWidgetMapper_CurrentIndex(QtObjectPtr ptr);
int QDataWidgetMapper_Orientation(QtObjectPtr ptr);
void QDataWidgetMapper_SetCurrentIndex(QtObjectPtr ptr, int index);
void QDataWidgetMapper_SetOrientation(QtObjectPtr ptr, int aOrientation);
void QDataWidgetMapper_SetSubmitPolicy(QtObjectPtr ptr, int policy);
int QDataWidgetMapper_SubmitPolicy(QtObjectPtr ptr);
QtObjectPtr QDataWidgetMapper_NewQDataWidgetMapper(QtObjectPtr parent);
void QDataWidgetMapper_AddMapping(QtObjectPtr ptr, QtObjectPtr widget, int section);
void QDataWidgetMapper_AddMapping2(QtObjectPtr ptr, QtObjectPtr widget, int section, QtObjectPtr propertyName);
void QDataWidgetMapper_ClearMapping(QtObjectPtr ptr);
void QDataWidgetMapper_ConnectCurrentIndexChanged(QtObjectPtr ptr);
void QDataWidgetMapper_DisconnectCurrentIndexChanged(QtObjectPtr ptr);
QtObjectPtr QDataWidgetMapper_ItemDelegate(QtObjectPtr ptr);
int QDataWidgetMapper_MappedSection(QtObjectPtr ptr, QtObjectPtr widget);
QtObjectPtr QDataWidgetMapper_MappedWidgetAt(QtObjectPtr ptr, int section);
QtObjectPtr QDataWidgetMapper_Model(QtObjectPtr ptr);
void QDataWidgetMapper_RemoveMapping(QtObjectPtr ptr, QtObjectPtr widget);
void QDataWidgetMapper_Revert(QtObjectPtr ptr);
QtObjectPtr QDataWidgetMapper_RootIndex(QtObjectPtr ptr);
void QDataWidgetMapper_SetCurrentModelIndex(QtObjectPtr ptr, QtObjectPtr index);
void QDataWidgetMapper_SetItemDelegate(QtObjectPtr ptr, QtObjectPtr delegate);
void QDataWidgetMapper_SetModel(QtObjectPtr ptr, QtObjectPtr model);
void QDataWidgetMapper_SetRootIndex(QtObjectPtr ptr, QtObjectPtr index);
int QDataWidgetMapper_Submit(QtObjectPtr ptr);
void QDataWidgetMapper_ToFirst(QtObjectPtr ptr);
void QDataWidgetMapper_ToLast(QtObjectPtr ptr);
void QDataWidgetMapper_ToNext(QtObjectPtr ptr);
void QDataWidgetMapper_ToPrevious(QtObjectPtr ptr);
void QDataWidgetMapper_DestroyQDataWidgetMapper(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif