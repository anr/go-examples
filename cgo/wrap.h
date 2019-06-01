#pragma once

#ifdef __cplusplus
extern "C" {
#endif

typedef void* wcook;

wcook create_wcook(const char* name, int should_fail, char** error);
char* wcook_greet(wcook w, const char* start);
void destroy_wcook(wcook w);

#ifdef __cplusplus
}
#endif
