Program: Statements
Statements: Statement | Statement Statements | ε
Statement: ExpressionStatement
ExpressionStatement: Expression
Expression: PrefixExpression | UnsignedInteger
PrefixExpression: SignedInteger
SignedInteger: "-" UnsignedInteger
UnsignedInteger: Digit | NonZeroDigit UnsignedInteger | Digit UnsignedInteger
Digit: "0" | NonZeroDigit
NonZeroDigit: "1" | "2" | "3" | "4" | "5" | "6" | "7" | "8" | "9"