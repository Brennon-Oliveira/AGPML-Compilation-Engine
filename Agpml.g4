grammar Agpml;

// Lexer rules

STRING
	:	'"' ( ESC_SEQ | ~('"'|'\\') )* '"'
	;

DIVIDER
    : '-'+ {p.GetInputStream().LA(1) == '\n'}?
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

VARIABLE
    : '$' ([a-zA-Z] | DIGIT)*
    ;

TEXT
    : ~(
        '#' | '*' | '$' | '\n' | '\r' | '\t' | ' ' | '"' | '\\' | [1.] | '>'
        ) ('\\$' | ~[$\r\n])+ {p.GetInputStream().LA(1) != '\n'}?
    ;


ATRIBUTION
    : '='
    ;

fragment
ESC_SEQ
	:	'\\"'
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
