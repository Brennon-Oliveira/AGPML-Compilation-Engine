grammar agpml;

// Lexer rules

TEXT_CONTENT
    : '\n' ~[ \t\r\n]+ '\n'
    | FLAG ~[ \t\r\n]+ '\n'
    | '}' ~[ \t\r\n]+ '\n'
    ;

CADEIA
	:	'"' ( ESC_SEQ | ~('"'|'\\') )* '"'
	;

FLAG
    : FLAG_TITLE
    | FLAG_OLIST
    | FLAG_ULIST
    | FLAG_QUOTE
    ;

DIVIDER
    : '\n' '-'* WS* '\n'
    ;

FLAG_TITLE
    : '\n#[' DIGIT+ ']'
    ;

FLAG_OLIST
    : '\n' DIGIT '.'
    ;

FLAG_ULIST
    : '\n*'
    ;

FLAG_QUOTE
    : '\n>'
    ;

VARIABLE
    : '$' (LETTER | DIGIT)*
    ;

ATTRIBUTE_NAME
    : LETTER (LETTER | DIGIT)*
    ;

ATRIBUTION
    : '='
    ;

fragment
ESC_SEQ
	:	'\\"';

COLOR
    : '#' [a-fA-F0-9]*
    ;
LETTER
    : [a-zA-Z] ;
DIGIT
    : [0-9] ;
COMMENT
    : '//' .*? '\n' -> skip ;
WS
    : [ \t\r\n]+ -> skip ;

// Parser rules

program
    : '.header' headerConfigs? ('.vars' varConfigs?)? '.body' dataConfigs? EOF
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
