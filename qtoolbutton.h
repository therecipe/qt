#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"

//Public Functions
QtObjectPtr QToolButton_New_QWidget(QtObjectPtr parent);
void QToolButton_Destroy(QtObjectPtr ptr);
int QToolButton_ArrowType(QtObjectPtr ptr);
int QToolButton_AutoRaise(QtObjectPtr ptr);
QtObjectPtr QToolButton_Menu(QtObjectPtr ptr);
void QToolButton_SetArrowType_ArrowType(QtObjectPtr ptr, int typ);
void QToolButton_SetAutoRaise_Bool(QtObjectPtr ptr, int enable);
void QToolButton_SetMenu_QMenu(QtObjectPtr ptr, QtObjectPtr menu);
int QToolButton_ToolButtonStyle(QtObjectPtr ptr);
//Public Slots
void QToolButton_ConnectSlotShowMenu(QtObjectPtr ptr);
void QToolButton_DisconnectSlotShowMenu(QtObjectPtr ptr);
void QToolButton_ShowMenu(QtObjectPtr ptr);
//Signals
void QToolButton_ConnectSignalTriggered(QtObjectPtr ptr);
void QToolButton_DisconnectSignalTriggered(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif
