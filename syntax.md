Program: Statements
Statements: Statement | Statement Statements | ε
Statement: ExpressionStatement ";"
ExpressionStatement: Expression
Expression: PrefixExpression | InfixExpression | GroupExpression | UnsignedInteger
PrefixExpression: SignedInteger
SignedInteger: "-" UnsignedInteger
InfixExpression: Expression InfixOperator Expression
InfixOperator: "+" | "-" | "*" | "/"
GroupExpression: "(" Expression ")"
UnsignedInteger: Digit | NonZeroDigit UnsignedInteger | Digit UnsignedInteger
Digit: "0" | NonZeroDigit
NonZeroDigit: "1" | "2" | "3" | "4" | "5" | "6" | "7" | "8" | "9"