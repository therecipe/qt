#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QChildEvent_NewQChildEvent(int ty, QtObjectPtr child);
int QChildEvent_Added(QtObjectPtr ptr);
QtObjectPtr QChildEvent_Child(QtObjectPtr ptr);
int QChildEvent_Polished(QtObjectPtr ptr);
int QChildEvent_Removed(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif