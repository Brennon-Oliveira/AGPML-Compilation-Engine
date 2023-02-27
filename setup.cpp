#include <iostream>
#include <fstream>
#include <string>
#include "./includes/Chooser.h"
#include "./includes/AgpmlToHtml.h"
#include <new>

using namespace std;

int main(int argc, char* argv[]){

    if(argc < 2){
        cout << "Input file not specified." << endl;
        return 0;
    }

    Chooser* chooser = new Chooser();
    chooser->sendFileToCorrect(argv[1]);
    

    return 0;
}
