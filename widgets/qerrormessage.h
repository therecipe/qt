#ifdef __cplusplus
extern "C" {
#endif

void* QErrorMessage_NewQErrorMessage(void* parent);
void* QErrorMessage_QErrorMessage_QtHandler();
void QErrorMessage_ShowMessage(void* ptr, char* message);
void QErrorMessage_ShowMessage2(void* ptr, char* message, char* ty);
void QErrorMessage_DestroyQErrorMessage(void* ptr);

#ifdef __cplusplus
}
#endif