grammar Agpml;

// Lexer rules

BOOLEAN
    : 'true' | 'false'
    ;

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

PARAGRAPH
    : ';'
    ;

// Parser rules

program
    : '.header' headerConfigs? ('.vars' varConfigs?)? '.body' body? EOF
    ;

headerConfigs
    : headerConfig+
    ;

headerConfig
    : ATTRIBUTE ATRIBUTION Value=(BOOLEAN | STRING | NUMBER | COLOR)
    ;

varConfigs
    : varConfig+
    ;

varConfig
    : VARIABLE ATRIBUTION (BOOLEAN | STRING | NUMBER | COLOR | VARIABLE)
    ;
body
    : (element | idGroup | classGroup)+
    ;

idGroup
    : OPEN_ID (element)* CLOSE_ID
    ;

classGroup
    : OPEN_CLASS (element)* CLOSE_CLASS
    ;

element
    : DIVIDER
    | line
    ;

line
    : (FLAG_OLIST | FLAG_ULIST | FLAG_QUOTE | FLAG_TITLE | TEXT_CONTENT | style) style?  TEXT_CONTENT?
    ;

style
    : OPEN_CURLY_BRACKET styleConfig+ CLOSE_CURLY_BRACKET
    ;

styleConfig
    : ATTRIBUTE DEFINER (BOOLEAN | STRING | NUMBER | COLOR | VARIABLE) (',')?
    ;
