#include "../includes/AgpmlToHtml.h"
#include <fstream>
#include <iostream>


AgpmlToHtml::AgpmlToHtml(){}

void AgpmlToHtml::convertFile(char* filePath){
    std::string line;
    std::fstream file(filePath);    

    if(file.is_open()){
        while(getline(file, line)){
            std::string convertLine(line);
        }
    } else {
        std::cout << "Error opening file." << std::endl;
    }
}

