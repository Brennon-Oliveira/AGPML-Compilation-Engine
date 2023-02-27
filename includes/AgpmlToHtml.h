#pragma once
#include <string>

class AgpmlToHtml {
    public:
        AgpmlToHtml();

        void convertFile(char* file_path);
        
        std::string firstStepInLine(std::string);


};
