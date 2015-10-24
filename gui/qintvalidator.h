#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QIntValidator_SetBottom(QtObjectPtr ptr, int v);
void QIntValidator_SetTop(QtObjectPtr ptr, int v);
QtObjectPtr QIntValidator_NewQIntValidator(QtObjectPtr parent);
QtObjectPtr QIntValidator_NewQIntValidator2(int minimum, int maximum, QtObjectPtr parent);
int QIntValidator_Bottom(QtObjectPtr ptr);
void QIntValidator_SetRange(QtObjectPtr ptr, int bottom, int top);
int QIntValidator_Top(QtObjectPtr ptr);
void QIntValidator_DestroyQIntValidator(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif