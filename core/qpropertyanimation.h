#ifdef __cplusplus
extern "C" {
#endif

void* QPropertyAnimation_PropertyName(void* ptr);
void QPropertyAnimation_SetPropertyName(void* ptr, void* propertyName);
void QPropertyAnimation_SetTargetObject(void* ptr, void* target);
void* QPropertyAnimation_TargetObject(void* ptr);
void* QPropertyAnimation_NewQPropertyAnimation(void* parent);
void* QPropertyAnimation_NewQPropertyAnimation2(void* target, void* propertyName, void* parent);
void QPropertyAnimation_DestroyQPropertyAnimation(void* ptr);

#ifdef __cplusplus
}
#endif