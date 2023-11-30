grammar agpml;










VARIABLES
    : '$' (LETTERS DIGITS)*
    ;


COLOR
    : '#' [a-fA-F0-9]*
    ;
LETTERS
    : [a-zA-Z]+ ;
DIGITS
    : [0-9]+ ;
COMMENTS
    : '//' .*? '\n' -> skip ;
WS
    : [ \t\r\n]+ -> skip ;





program
    : '.header' headerConfigs ('.vars' varConfigs)? '.body' dataConfigs EOF TESTE
    ;

headerConfigs
    : '1'
    ;

varConfigs
    : '2'
    ;

dataConfigs
    : '3'
    ;
