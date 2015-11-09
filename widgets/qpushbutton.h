#ifdef __cplusplus
extern "C" {
#endif

int QPushButton_AutoDefault(void* ptr);
int QPushButton_IsDefault(void* ptr);
int QPushButton_IsFlat(void* ptr);
void QPushButton_SetAutoDefault(void* ptr, int v);
void QPushButton_SetDefault(void* ptr, int v);
void QPushButton_SetFlat(void* ptr, int v);
void* QPushButton_NewQPushButton(void* parent);
void* QPushButton_NewQPushButton3(void* icon, char* text, void* parent);
void* QPushButton_NewQPushButton2(char* text, void* parent);
void* QPushButton_Menu(void* ptr);
void QPushButton_SetMenu(void* ptr, void* menu);
void QPushButton_ShowMenu(void* ptr);
void QPushButton_DestroyQPushButton(void* ptr);

#ifdef __cplusplus
}
#endif