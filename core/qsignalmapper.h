#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

QtObjectPtr QSignalMapper_NewQSignalMapper(QtObjectPtr parent);
void QSignalMapper_Map(QtObjectPtr ptr);
void QSignalMapper_Map2(QtObjectPtr ptr, QtObjectPtr sender);
void QSignalMapper_ConnectMapped(QtObjectPtr ptr);
void QSignalMapper_DisconnectMapped(QtObjectPtr ptr);
QtObjectPtr QSignalMapper_Mapping4(QtObjectPtr ptr, QtObjectPtr object);
QtObjectPtr QSignalMapper_Mapping2(QtObjectPtr ptr, char* id);
QtObjectPtr QSignalMapper_Mapping(QtObjectPtr ptr, int id);
void QSignalMapper_RemoveMappings(QtObjectPtr ptr, QtObjectPtr sender);
void QSignalMapper_SetMapping4(QtObjectPtr ptr, QtObjectPtr sender, QtObjectPtr object);
void QSignalMapper_SetMapping2(QtObjectPtr ptr, QtObjectPtr sender, char* text);
void QSignalMapper_SetMapping(QtObjectPtr ptr, QtObjectPtr sender, int id);
void QSignalMapper_DestroyQSignalMapper(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif