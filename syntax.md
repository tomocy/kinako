Program: Statements  
Statements: Statement | Statement Statements | ε
Statement: ExpressionStatement ";" | VariableDeclaration ";"  
ExpressionStatement: Expression  
Expression: PrefixExpression | InfixExpression | GroupExpression | Identifier | UnsignedInteger  
PrefixExpression: SignedInteger  
SignedInteger: "-" UnsignedInteger  
InfixExpression: Expression InfixOperator Expression  
InfixOperator: "+" | "-" | "*" | "/"  
GroupExpression: "(" Expression ")"  
UnsignedInteger: Digit | NonZeroDigit UnsignedInteger | Digit UnsignedInteger  
Digit: "0" | NonZeroDigit  
NonZeroDigit: "1" | "2" | "3" | "4" | "5" | "6" | "7" | "8" | "9"  
VariableDeclaration: "var" Identifier Type "=" Expression  
Identifier: Letter  
Type: Identifier  
Letter: /* a to z or A to Z */  