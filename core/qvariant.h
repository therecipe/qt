#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

char* QVariant_ToStringList(QtObjectPtr ptr);
char* QVariant_ToUrl(QtObjectPtr ptr);
void QVariant_DestroyQVariant(QtObjectPtr ptr);
void QVariant_Clear(QtObjectPtr ptr);
int QVariant_Convert(QtObjectPtr ptr, int targetTypeId);
int QVariant_IsNull(QtObjectPtr ptr);
int QVariant_IsValid(QtObjectPtr ptr);
int QVariant_ToBool(QtObjectPtr ptr);
int QVariant_ToInt(QtObjectPtr ptr, int ok);
QtObjectPtr QVariant_ToModelIndex(QtObjectPtr ptr);
char* QVariant_ToString(QtObjectPtr ptr);
int QVariant_UserType(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif