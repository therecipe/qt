#ifdef __cplusplus
extern "C" {
#endif

void* QColorDialog_CurrentColor(void* ptr);
int QColorDialog_Options(void* ptr);
void QColorDialog_SetCurrentColor(void* ptr, void* color);
void QColorDialog_SetOptions(void* ptr, int options);
void* QColorDialog_NewQColorDialog(void* parent);
void* QColorDialog_NewQColorDialog2(void* initial, void* parent);
void QColorDialog_ConnectColorSelected(void* ptr);
void QColorDialog_DisconnectColorSelected(void* ptr);
void QColorDialog_ConnectCurrentColorChanged(void* ptr);
void QColorDialog_DisconnectCurrentColorChanged(void* ptr);
void* QColorDialog_QColorDialog_CustomColor(int index);
int QColorDialog_QColorDialog_CustomCount();
void* QColorDialog_QColorDialog_GetColor(void* initial, void* parent, char* title, int options);
void QColorDialog_Open(void* ptr, void* receiver, char* member);
void* QColorDialog_SelectedColor(void* ptr);
void QColorDialog_QColorDialog_SetCustomColor(int index, void* color);
void QColorDialog_SetOption(void* ptr, int option, int on);
void QColorDialog_QColorDialog_SetStandardColor(int index, void* color);
void QColorDialog_SetVisible(void* ptr, int visible);
void* QColorDialog_QColorDialog_StandardColor(int index);
int QColorDialog_TestOption(void* ptr, int option);
void QColorDialog_DestroyQColorDialog(void* ptr);

#ifdef __cplusplus
}
#endif