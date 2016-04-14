#ifndef BRIDGE_H
#define BRIDGE_H

#ifdef __cplusplus
extern "C" {
#endif

void initSignalHandler(void* rootContext);

void sendSignalToQml(char* value);

#ifdef __cplusplus
}
#endif

#endif // BRIDGE_H
