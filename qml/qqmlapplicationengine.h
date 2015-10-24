#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QQmlApplicationEngine_NewQQmlApplicationEngine(QtObjectPtr parent);
QtObjectPtr QQmlApplicationEngine_NewQQmlApplicationEngine3(char* filePath, QtObjectPtr parent);
QtObjectPtr QQmlApplicationEngine_NewQQmlApplicationEngine2(char* url, QtObjectPtr parent);
void QQmlApplicationEngine_Load2(QtObjectPtr ptr, char* filePath);
void QQmlApplicationEngine_Load(QtObjectPtr ptr, char* url);
void QQmlApplicationEngine_LoadData(QtObjectPtr ptr, QtObjectPtr data, char* url);
void QQmlApplicationEngine_ConnectObjectCreated(QtObjectPtr ptr);
void QQmlApplicationEngine_DisconnectObjectCreated(QtObjectPtr ptr);
void QQmlApplicationEngine_DestroyQQmlApplicationEngine(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif