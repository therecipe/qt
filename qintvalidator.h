#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"

//Public Functions
QtObjectPtr QIntValidator_New_QObject(QtObjectPtr parent);
QtObjectPtr QIntValidator_New_Int_Int_QObject(int minimum, int maximum, QtObjectPtr parent);
void QIntValidator_Destroy(QtObjectPtr ptr);
int QIntValidator_Bottom(QtObjectPtr ptr);
void QIntValidator_SetBottom_Int(QtObjectPtr ptr, int in);
void QIntValidator_SetRange_Int_Int(QtObjectPtr ptr, int bottom, int top);
void QIntValidator_SetTop_Int(QtObjectPtr ptr, int in);
int QIntValidator_Top(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif
