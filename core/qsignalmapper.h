#ifdef __cplusplus
extern "C" {
#endif

void* QSignalMapper_NewQSignalMapper(void* parent);
void QSignalMapper_Map(void* ptr);
void QSignalMapper_Map2(void* ptr, void* sender);
void QSignalMapper_ConnectMapped(void* ptr);
void QSignalMapper_DisconnectMapped(void* ptr);
void* QSignalMapper_Mapping4(void* ptr, void* object);
void* QSignalMapper_Mapping2(void* ptr, char* id);
void* QSignalMapper_Mapping(void* ptr, int id);
void QSignalMapper_RemoveMappings(void* ptr, void* sender);
void QSignalMapper_SetMapping4(void* ptr, void* sender, void* object);
void QSignalMapper_SetMapping2(void* ptr, void* sender, char* text);
void QSignalMapper_SetMapping(void* ptr, void* sender, int id);
void QSignalMapper_DestroyQSignalMapper(void* ptr);

#ifdef __cplusplus
}
#endif