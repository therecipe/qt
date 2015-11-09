#ifdef __cplusplus
extern "C" {
#endif

void* QTime_NewQTime();
void* QTime_NewQTime3(int h, int m, int s, int ms);
int QTime_Elapsed(void* ptr);
int QTime_Hour(void* ptr);
int QTime_IsNull(void* ptr);
int QTime_QTime_IsValid2(int h, int m, int s, int ms);
int QTime_IsValid(void* ptr);
int QTime_Minute(void* ptr);
int QTime_Msec(void* ptr);
int QTime_MsecsSinceStartOfDay(void* ptr);
int QTime_MsecsTo(void* ptr, void* t);
int QTime_Restart(void* ptr);
int QTime_Second(void* ptr);
int QTime_SecsTo(void* ptr, void* t);
int QTime_SetHMS(void* ptr, int h, int m, int s, int ms);
void QTime_Start(void* ptr);
char* QTime_ToString2(void* ptr, int format);
char* QTime_ToString(void* ptr, char* format);

#ifdef __cplusplus
}
#endif