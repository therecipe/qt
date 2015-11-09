#ifdef __cplusplus
extern "C" {
#endif

void* QScrollerProperties_NewQScrollerProperties();
void* QScrollerProperties_NewQScrollerProperties2(void* sp);
void* QScrollerProperties_ScrollMetric(void* ptr, int metric);
void QScrollerProperties_QScrollerProperties_SetDefaultScrollerProperties(void* sp);
void QScrollerProperties_SetScrollMetric(void* ptr, int metric, void* value);
void QScrollerProperties_QScrollerProperties_UnsetDefaultScrollerProperties();
void QScrollerProperties_DestroyQScrollerProperties(void* ptr);

#ifdef __cplusplus
}
#endif