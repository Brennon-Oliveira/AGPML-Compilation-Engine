grammar Agpml;

// Lexer rules

TEXT_CONTENT
	:	';' ( ESC_SEQ | ~(';'|'\\'|'\n') )* ('\n' | EOF | ';')
     	;

STRING
	:	'"' ( ESC_SEQ | ~('"'|'\\') )* '"'
	;

DIVIDER
    : ([\-])+
    ;

FLAG_TITLE
    : '#[' DIGIT+ ']'
    ;

FLAG_OLIST
    : DIGIT '.'
    ;

FLAG_ULIST
    : '*'
    ;

FLAG_QUOTE
    : '>'
    ;

ATTRIBUTE
    : [a-zA-Z] [a-zA-Z]*
    ;

VARIABLE
    : '$' ([a-zA-Z] | DIGIT)*
    ;

ATRIBUTION
    : '='
    ;

DEFINER
    : ':'
    ;

SEPARATOR
    : ','
    ;

fragment
ESC_SEQ
	:	'\\"' | '\\;'
	;

COLOR
    : '#'  [a-fA-F0-9] [a-fA-F0-9] [a-fA-F0-9] ([a-fA-F0-9] [a-fA-F0-9] [a-fA-F0-9])?
    ;
NUMBER
    : DIGIT+ ('.' DIGIT+)?
    ;
DIGIT
    : [0-9]
    ;

COMMENT
    : '//' .*? '\n' -> skip
    ;

WS
    : [ \t\r\n]+ -> skip
    ;

OPEN_CURLY_BRACKET
    : '{'
    ;

CLOSE_CURLY_BRACKET
    : '}'
    ;

OPEN_CLASS
    : '..' [a-zA-Z] [a-zA-Z0-9]*
    ;

CLOSE_CLASS
    : '..'
    ;

OPEN_ID
    : '##' [a-zA-Z] [a-zA-Z0-9]*
    ;

CLOSE_ID
    : '##'
    ;

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

style
    : '{' styleConfig? '}'
    ;

styleConfig
    : '1'
    ;
