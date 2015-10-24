#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QMediaTimeRange_NewQMediaTimeRange();
QtObjectPtr QMediaTimeRange_NewQMediaTimeRange3(QtObjectPtr interval);
QtObjectPtr QMediaTimeRange_NewQMediaTimeRange4(QtObjectPtr ran);
void QMediaTimeRange_AddInterval(QtObjectPtr ptr, QtObjectPtr interval);
void QMediaTimeRange_AddTimeRange(QtObjectPtr ptr, QtObjectPtr ran);
void QMediaTimeRange_Clear(QtObjectPtr ptr);
int QMediaTimeRange_IsContinuous(QtObjectPtr ptr);
int QMediaTimeRange_IsEmpty(QtObjectPtr ptr);
void QMediaTimeRange_RemoveInterval(QtObjectPtr ptr, QtObjectPtr interval);
void QMediaTimeRange_RemoveTimeRange(QtObjectPtr ptr, QtObjectPtr ran);
void QMediaTimeRange_DestroyQMediaTimeRange(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif