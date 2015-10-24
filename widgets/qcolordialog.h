#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QColorDialog_Options(QtObjectPtr ptr);
void QColorDialog_SetCurrentColor(QtObjectPtr ptr, QtObjectPtr color);
void QColorDialog_SetOptions(QtObjectPtr ptr, int options);
QtObjectPtr QColorDialog_NewQColorDialog(QtObjectPtr parent);
QtObjectPtr QColorDialog_NewQColorDialog2(QtObjectPtr initial, QtObjectPtr parent);
int QColorDialog_QColorDialog_CustomCount();
void QColorDialog_Open(QtObjectPtr ptr, QtObjectPtr receiver, char* member);
void QColorDialog_QColorDialog_SetCustomColor(int index, QtObjectPtr color);
void QColorDialog_SetOption(QtObjectPtr ptr, int option, int on);
void QColorDialog_QColorDialog_SetStandardColor(int index, QtObjectPtr color);
void QColorDialog_SetVisible(QtObjectPtr ptr, int visible);
int QColorDialog_TestOption(QtObjectPtr ptr, int option);
void QColorDialog_DestroyQColorDialog(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif