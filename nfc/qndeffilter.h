#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QNdefFilter_NewQNdefFilter();
QtObjectPtr QNdefFilter_NewQNdefFilter2(QtObjectPtr other);
void QNdefFilter_Clear(QtObjectPtr ptr);
int QNdefFilter_OrderMatch(QtObjectPtr ptr);
int QNdefFilter_RecordCount(QtObjectPtr ptr);
void QNdefFilter_SetOrderMatch(QtObjectPtr ptr, int on);
void QNdefFilter_DestroyQNdefFilter(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif