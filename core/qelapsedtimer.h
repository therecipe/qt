#ifdef __cplusplus
extern "C" {
#endif

void* QElapsedTimer_NewQElapsedTimer();
void QElapsedTimer_Invalidate(void* ptr);
int QElapsedTimer_IsValid(void* ptr);
int QElapsedTimer_QElapsedTimer_ClockType();
int QElapsedTimer_QElapsedTimer_IsMonotonic();
void QElapsedTimer_Start(void* ptr);

#ifdef __cplusplus
}
#endif