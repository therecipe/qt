#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QTime_NewQTime();
QtObjectPtr QTime_NewQTime3(int h, int m, int s, int ms);
int QTime_Elapsed(QtObjectPtr ptr);
int QTime_Hour(QtObjectPtr ptr);
int QTime_IsNull(QtObjectPtr ptr);
int QTime_QTime_IsValid2(int h, int m, int s, int ms);
int QTime_IsValid(QtObjectPtr ptr);
int QTime_Minute(QtObjectPtr ptr);
int QTime_Msec(QtObjectPtr ptr);
int QTime_MsecsSinceStartOfDay(QtObjectPtr ptr);
int QTime_MsecsTo(QtObjectPtr ptr, QtObjectPtr t);
int QTime_Restart(QtObjectPtr ptr);
int QTime_Second(QtObjectPtr ptr);
int QTime_SecsTo(QtObjectPtr ptr, QtObjectPtr t);
int QTime_SetHMS(QtObjectPtr ptr, int h, int m, int s, int ms);
void QTime_Start(QtObjectPtr ptr);
char* QTime_ToString2(QtObjectPtr ptr, int format);
char* QTime_ToString(QtObjectPtr ptr, char* format);

#ifdef __cplusplus
}
#endif