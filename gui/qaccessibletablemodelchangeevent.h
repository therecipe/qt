#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QAccessibleTableModelChangeEvent_NewQAccessibleTableModelChangeEvent2(QtObjectPtr iface, int changeType);
QtObjectPtr QAccessibleTableModelChangeEvent_NewQAccessibleTableModelChangeEvent(QtObjectPtr object, int changeType);
int QAccessibleTableModelChangeEvent_FirstColumn(QtObjectPtr ptr);
int QAccessibleTableModelChangeEvent_FirstRow(QtObjectPtr ptr);
int QAccessibleTableModelChangeEvent_LastColumn(QtObjectPtr ptr);
int QAccessibleTableModelChangeEvent_LastRow(QtObjectPtr ptr);
int QAccessibleTableModelChangeEvent_ModelChangeType(QtObjectPtr ptr);
void QAccessibleTableModelChangeEvent_SetFirstColumn(QtObjectPtr ptr, int column);
void QAccessibleTableModelChangeEvent_SetFirstRow(QtObjectPtr ptr, int row);
void QAccessibleTableModelChangeEvent_SetLastColumn(QtObjectPtr ptr, int column);
void QAccessibleTableModelChangeEvent_SetLastRow(QtObjectPtr ptr, int row);
void QAccessibleTableModelChangeEvent_SetModelChangeType(QtObjectPtr ptr, int changeType);

#ifdef __cplusplus
}
#endif