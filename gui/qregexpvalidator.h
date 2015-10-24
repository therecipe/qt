#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QRegExpValidator_SetRegExp(QtObjectPtr ptr, QtObjectPtr rx);
QtObjectPtr QRegExpValidator_NewQRegExpValidator(QtObjectPtr parent);
QtObjectPtr QRegExpValidator_NewQRegExpValidator2(QtObjectPtr rx, QtObjectPtr parent);
void QRegExpValidator_DestroyQRegExpValidator(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif