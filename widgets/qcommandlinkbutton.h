#ifdef __cplusplus
extern "C" {
#endif

char* QCommandLinkButton_Description(void* ptr);
void QCommandLinkButton_SetDescription(void* ptr, char* description);
void* QCommandLinkButton_NewQCommandLinkButton(void* parent);
void* QCommandLinkButton_NewQCommandLinkButton2(char* text, void* parent);
void* QCommandLinkButton_NewQCommandLinkButton3(char* text, char* description, void* parent);
void QCommandLinkButton_DestroyQCommandLinkButton(void* ptr);

#ifdef __cplusplus
}
#endif