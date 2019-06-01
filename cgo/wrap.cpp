#include "wrap.h"

#include "string.h"
#include "kitchen.h"

wcook create_wcook(const char* name, int should_fail, char** error) {
    try {
        kitchen::Cook* cook = new kitchen::Cook(name, should_fail == 1);
        return (wcook)cook;
    } catch (std::string e) {
        *error = strdup(e.c_str());
       return nullptr;
    }
}

void destroy_wcook(wcook w) {
    delete (kitchen::Cook*)w;
}

char* wcook_greet(wcook w, const char* start) {
    std::string s = ((kitchen::Cook*)w)->greet(start);

    return strdup(s.c_str());
}
