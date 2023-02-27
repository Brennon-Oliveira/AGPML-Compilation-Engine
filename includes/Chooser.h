#pragma once

class Chooser {
    
    public:
        Chooser();

        void sendFileToCorrect(char* file);
        
        std::string convertLine(std::string line);
};
