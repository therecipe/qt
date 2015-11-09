#ifdef __cplusplus
extern "C" {
#endif

void* QMediaTimeRange_NewQMediaTimeRange();
void* QMediaTimeRange_NewQMediaTimeRange3(void* interval);
void* QMediaTimeRange_NewQMediaTimeRange4(void* ran);
void QMediaTimeRange_AddInterval(void* ptr, void* interval);
void QMediaTimeRange_AddTimeRange(void* ptr, void* ran);
void QMediaTimeRange_Clear(void* ptr);
int QMediaTimeRange_IsContinuous(void* ptr);
int QMediaTimeRange_IsEmpty(void* ptr);
void QMediaTimeRange_RemoveInterval(void* ptr, void* interval);
void QMediaTimeRange_RemoveTimeRange(void* ptr, void* ran);
void QMediaTimeRange_DestroyQMediaTimeRange(void* ptr);

#ifdef __cplusplus
}
#endif