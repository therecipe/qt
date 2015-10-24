#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QSizeF_NewQSizeF();
QtObjectPtr QSizeF_NewQSizeF2(QtObjectPtr size);
int QSizeF_IsEmpty(QtObjectPtr ptr);
int QSizeF_IsNull(QtObjectPtr ptr);
int QSizeF_IsValid(QtObjectPtr ptr);
void QSizeF_Scale2(QtObjectPtr ptr, QtObjectPtr size, int mode);
void QSizeF_Transpose(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif