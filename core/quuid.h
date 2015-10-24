#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QUuid_Variant(QtObjectPtr ptr);
int QUuid_Version(QtObjectPtr ptr);
QtObjectPtr QUuid_NewQUuid();
QtObjectPtr QUuid_NewQUuid5(QtObjectPtr text);
QtObjectPtr QUuid_NewQUuid3(char* text);
int QUuid_IsNull(QtObjectPtr ptr);
char* QUuid_ToString(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif