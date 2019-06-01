#pragma once

#include <string>

namespace kitchen {

class Cook {
    std::string name;

public:
    Cook(const std::string& name, bool should_fail=false);
    std::string greet(const std::string& start) const;
};

} // namespace kitchen
