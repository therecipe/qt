#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QPropertyAnimation_SetPropertyName(QtObjectPtr ptr, QtObjectPtr propertyName);
void QPropertyAnimation_SetTargetObject(QtObjectPtr ptr, QtObjectPtr target);
QtObjectPtr QPropertyAnimation_TargetObject(QtObjectPtr ptr);
QtObjectPtr QPropertyAnimation_NewQPropertyAnimation(QtObjectPtr parent);
QtObjectPtr QPropertyAnimation_NewQPropertyAnimation2(QtObjectPtr target, QtObjectPtr propertyName, QtObjectPtr parent);
void QPropertyAnimation_DestroyQPropertyAnimation(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif