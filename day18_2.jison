/* lexical grammar */
%lex

%%
\s+                   /* skip whitespace */
[0-9]+("."[0-9]+)?\b  return 'NUMBER';
"*"                   return '*';
"+"                   return '+';
"("                   return '(';
")"                   return ')';

/lex

/* operator associations and precedence */

%left '*'
%left '+'

%start expressions

%% /* language grammar */

expressions : expr { return $$ };

expr : expr '+' expr { $$ = $1 + $3}
     | expr '*' expr { $$ = $1 * $3}
     | '(' expr ')' { $$ = $2 }
     | NUMBER { $$ = Number($1) };
