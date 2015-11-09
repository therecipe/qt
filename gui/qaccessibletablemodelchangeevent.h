#ifdef __cplusplus
extern "C" {
#endif

void* QAccessibleTableModelChangeEvent_NewQAccessibleTableModelChangeEvent2(void* iface, int changeType);
void* QAccessibleTableModelChangeEvent_NewQAccessibleTableModelChangeEvent(void* object, int changeType);
int QAccessibleTableModelChangeEvent_FirstColumn(void* ptr);
int QAccessibleTableModelChangeEvent_FirstRow(void* ptr);
int QAccessibleTableModelChangeEvent_LastColumn(void* ptr);
int QAccessibleTableModelChangeEvent_LastRow(void* ptr);
int QAccessibleTableModelChangeEvent_ModelChangeType(void* ptr);
void QAccessibleTableModelChangeEvent_SetFirstColumn(void* ptr, int column);
void QAccessibleTableModelChangeEvent_SetFirstRow(void* ptr, int row);
void QAccessibleTableModelChangeEvent_SetLastColumn(void* ptr, int column);
void QAccessibleTableModelChangeEvent_SetLastRow(void* ptr, int row);
void QAccessibleTableModelChangeEvent_SetModelChangeType(void* ptr, int changeType);

#ifdef __cplusplus
}
#endif