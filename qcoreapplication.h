#ifdef __cplusplus
extern "C" {
#endif

#include "cgoutil.h"

//Public Functions
QtObjectPtr QCoreApplication_New_Int_String(int argc, char* argv);
void QCoreApplication_Destroy(QtObjectPtr ptr);
//Static Public Members
void QCoreApplication_AddLibraryPath_String(char* path);
char* QCoreApplication_ApplicationDirPath();
char* QCoreApplication_ApplicationFilePath();
char* QCoreApplication_ApplicationName();
char* QCoreApplication_ApplicationVersion();
int QCoreApplication_ClosingDown();
int QCoreApplication_Exec();
void QCoreApplication_Exit_Int(int returnCode);
QtObjectPtr QCoreApplication_Instance();
int QCoreApplication_IsQuitLockEnabled();
int QCoreApplication_IsSetuidAllowed();
char* QCoreApplication_OrganizationDomain();
char* QCoreApplication_OrganizationName();
void QCoreApplication_RemoveLibraryPath_String(char* path);
void QCoreApplication_RemovePostedEvents_QObject_Int(QtObjectPtr receiver, int eventType);
void QCoreApplication_SendPostedEvents_QObject_Int(QtObjectPtr receiver, int event_type);
void QCoreApplication_SetApplicationName_String(char* application);
void QCoreApplication_SetApplicationVersion_String(char* version);
void QCoreApplication_SetAttribute_ApplicationAttribute_Bool(int attribute, int on);
void QCoreApplication_SetOrganizationDomain_String(char* orgDomain);
void QCoreApplication_SetOrganizationName_String(char* orgName);
void QCoreApplication_SetQuitLockEnabled_Bool(int enabled);
void QCoreApplication_SetSetuidAllowed_Bool(int allow);
int QCoreApplication_StartingUp();
int QCoreApplication_TestAttribute_ApplicationAttribute(int attribute);
char* QCoreApplication_Translate_String_String_String_Int(char* context, char* sourceText, char* disambiguation, int n);

#ifdef __cplusplus
}
#endif
