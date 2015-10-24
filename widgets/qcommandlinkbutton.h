#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

char* QCommandLinkButton_Description(QtObjectPtr ptr);
void QCommandLinkButton_SetDescription(QtObjectPtr ptr, char* description);
QtObjectPtr QCommandLinkButton_NewQCommandLinkButton(QtObjectPtr parent);
QtObjectPtr QCommandLinkButton_NewQCommandLinkButton2(char* text, QtObjectPtr parent);
QtObjectPtr QCommandLinkButton_NewQCommandLinkButton3(char* text, char* description, QtObjectPtr parent);
void QCommandLinkButton_DestroyQCommandLinkButton(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif