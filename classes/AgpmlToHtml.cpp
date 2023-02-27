#include "../includes/AgpmlToHtml.h"
#include <fstream>
#include <iostream>
#include <regex>

AgpmlToHtml::AgpmlToHtml(){}

void AgpmlToHtml::convertFile(char* filePath){
    std::string line;
    std::fstream file(filePath);    

    if(file.is_open()){
        while(getline(file, line)){
            this->firstStepInLine(line.c_str());
        }
    } else {
        std::cout << "Error opening file." << std::endl;
    }
}

std::string AgpmlToHtml::firstStepInLine(std::string line){
    std::cout << line << std::endl;

    std::regex getFirstLetter("//S");
    std::smatch results;
    std::regex_search(line, results, getFirstLetter);

    std::string letter = results[1].str();

    return line;
}

