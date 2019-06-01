#include "kitchen.h"

namespace kitchen {

Cook::Cook(const std::string& name, bool should_fail) : name(name) {
    if (should_fail) {
        throw std::string("Too many cooks spoil the broth");
    }
}

std::string Cook::greet(const std::string& start) const {
    return start + name;
}

} // namespace kitchen
