#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

int QStyleOption_SO_Slider_Type();
int QStyleOption_SO_SpinBox_Type();
int QStyleOption_SO_ToolButton_Type();
int QStyleOption_SO_ComboBox_Type();
int QStyleOption_SO_TitleBar_Type();
int QStyleOption_SO_GroupBox_Type();
int QStyleOption_SO_SizeGrip_Type();
QtObjectPtr QStyleOption_NewQStyleOption2(QtObjectPtr other);
QtObjectPtr QStyleOption_NewQStyleOption(int version, int ty);
void QStyleOption_InitFrom(QtObjectPtr ptr, QtObjectPtr widget);
void QStyleOption_DestroyQStyleOption(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif