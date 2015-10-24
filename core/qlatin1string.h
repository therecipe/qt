#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QLatin1String_NewQLatin1String3(QtObjectPtr str);
QtObjectPtr QLatin1String_NewQLatin1String(char* str);
QtObjectPtr QLatin1String_NewQLatin1String2(char* str, int size);
int QLatin1String_Size(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif