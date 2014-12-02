#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"

//Public Functions
QtObjectPtr QButtonGroup_New_QObject(QtObjectPtr parent);
void QButtonGroup_Destroy(QtObjectPtr ptr);
void QButtonGroup_AddButton_QAbstractButton_Int(QtObjectPtr ptr, QtObjectPtr button, int id);
QtObjectPtr QButtonGroup_Button_Int(QtObjectPtr ptr, int id);
QtObjectPtr QButtonGroup_CheckedButton(QtObjectPtr ptr);
int QButtonGroup_CheckedId(QtObjectPtr ptr);
int QButtonGroup_Exclusive(QtObjectPtr ptr);
int QButtonGroup_Id_QAbstractButton(QtObjectPtr ptr, QtObjectPtr button);
void QButtonGroup_RemoveButton_QAbstractButton(QtObjectPtr ptr, QtObjectPtr button);
void QButtonGroup_SetExclusive_Bool(QtObjectPtr ptr, int exclusive);
void QButtonGroup_SetId_QAbstractButton_Int(QtObjectPtr ptr, QtObjectPtr button, int id);

#ifdef __cplusplus
}
#endif
