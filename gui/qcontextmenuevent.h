#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QContextMenuEvent_NewQContextMenuEvent3(int reason, QtObjectPtr pos);
QtObjectPtr QContextMenuEvent_NewQContextMenuEvent2(int reason, QtObjectPtr pos, QtObjectPtr globalPos);
QtObjectPtr QContextMenuEvent_NewQContextMenuEvent(int reason, QtObjectPtr pos, QtObjectPtr globalPos, int modifiers);
int QContextMenuEvent_GlobalX(QtObjectPtr ptr);
int QContextMenuEvent_GlobalY(QtObjectPtr ptr);
int QContextMenuEvent_Reason(QtObjectPtr ptr);
int QContextMenuEvent_X(QtObjectPtr ptr);
int QContextMenuEvent_Y(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif