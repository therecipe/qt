#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QEnterEvent_NewQEnterEvent(QtObjectPtr localPos, QtObjectPtr windowPos, QtObjectPtr screenPos);
int QEnterEvent_GlobalX(QtObjectPtr ptr);
int QEnterEvent_GlobalY(QtObjectPtr ptr);
int QEnterEvent_X(QtObjectPtr ptr);
int QEnterEvent_Y(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif