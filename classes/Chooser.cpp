#include "../includes/Chooser.h"
#include "../includes/AgpmlToHtml.h"
#include <fstream>
#include <iostream>
#include <regex>
#include <string>
#include <new>

Chooser::Chooser(){}

void Chooser::sendFileToCorrect(char* filePath){

    std::string fileName(filePath);

    std::regex pattern("\\.([^\\.]+)$");
    std::smatch matches;

    if (std::regex_search(fileName, matches, pattern)) {
        std::string extension = matches[1].str();
        if(extension == "agpml"){
            AgpmlToHtml* agpmlToHtml = new AgpmlToHtml();
            agpmlToHtml->convertFile(filePath);
        }
    } else {
        std::cout << "File extension not found." << std::endl;
    }
}
