#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QPushButton_AutoDefault(QtObjectPtr ptr);
int QPushButton_IsDefault(QtObjectPtr ptr);
int QPushButton_IsFlat(QtObjectPtr ptr);
void QPushButton_SetAutoDefault(QtObjectPtr ptr, int v);
void QPushButton_SetDefault(QtObjectPtr ptr, int v);
void QPushButton_SetFlat(QtObjectPtr ptr, int v);
QtObjectPtr QPushButton_NewQPushButton(QtObjectPtr parent);
QtObjectPtr QPushButton_NewQPushButton3(QtObjectPtr icon, char* text, QtObjectPtr parent);
QtObjectPtr QPushButton_NewQPushButton2(char* text, QtObjectPtr parent);
QtObjectPtr QPushButton_Menu(QtObjectPtr ptr);
void QPushButton_SetMenu(QtObjectPtr ptr, QtObjectPtr menu);
void QPushButton_ShowMenu(QtObjectPtr ptr);
void QPushButton_DestroyQPushButton(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif