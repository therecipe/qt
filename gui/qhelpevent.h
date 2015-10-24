#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QHelpEvent_NewQHelpEvent(int ty, QtObjectPtr pos, QtObjectPtr globalPos);
int QHelpEvent_GlobalX(QtObjectPtr ptr);
int QHelpEvent_GlobalY(QtObjectPtr ptr);
int QHelpEvent_X(QtObjectPtr ptr);
int QHelpEvent_Y(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif