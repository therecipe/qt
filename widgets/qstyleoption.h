#ifdef __cplusplus
extern "C" {
#endif

int QStyleOption_SO_Slider_Type();
int QStyleOption_SO_SpinBox_Type();
int QStyleOption_SO_ToolButton_Type();
int QStyleOption_SO_ComboBox_Type();
int QStyleOption_SO_TitleBar_Type();
int QStyleOption_SO_GroupBox_Type();
int QStyleOption_SO_SizeGrip_Type();
void* QStyleOption_NewQStyleOption2(void* other);
void* QStyleOption_NewQStyleOption(int version, int ty);
void QStyleOption_InitFrom(void* ptr, void* widget);
void QStyleOption_DestroyQStyleOption(void* ptr);

#ifdef __cplusplus
}
#endif