#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"

//Public Functions
QtObjectPtr QPushButton_New_QWidget(QtObjectPtr parent);
QtObjectPtr QPushButton_New_String_QWidget(char* text, QtObjectPtr parent);
void QPushButton_Destroy(QtObjectPtr ptr);
int QPushButton_AutoDefault(QtObjectPtr ptr);
int QPushButton_IsDefault(QtObjectPtr ptr);
int QPushButton_IsFlat(QtObjectPtr ptr);
QtObjectPtr QPushButton_Menu(QtObjectPtr ptr);
void QPushButton_SetAutoDefault_Bool(QtObjectPtr ptr, int autoDefault);
void QPushButton_SetDefault_Bool(QtObjectPtr ptr, int defaul);
void QPushButton_SetFlat_Bool(QtObjectPtr ptr, int flat);
void QPushButton_SetMenu_QMenu(QtObjectPtr ptr, QtObjectPtr menu);
//Public Slots
void QPushButton_ConnectSlotShowMenu(QtObjectPtr ptr);
void QPushButton_DisconnectSlotShowMenu(QtObjectPtr ptr);
void QPushButton_ShowMenu(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif
