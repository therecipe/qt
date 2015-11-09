#ifdef __cplusplus
extern "C" {
#endif

void* QAccessibleValueInterface_CurrentValue(void* ptr);
void* QAccessibleValueInterface_MaximumValue(void* ptr);
void* QAccessibleValueInterface_MinimumStepSize(void* ptr);
void* QAccessibleValueInterface_MinimumValue(void* ptr);
void QAccessibleValueInterface_SetCurrentValue(void* ptr, void* value);
void QAccessibleValueInterface_DestroyQAccessibleValueInterface(void* ptr);

#ifdef __cplusplus
}
#endif