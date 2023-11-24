del /F GoParser.g4
copy GoParser_LL.g4 GoParser.g4
java -jar antlr-4.9.3-complete.jar -Dlanguage=CSharp GoLexer.g4
java -jar antlr-4.9.3-complete.jar -Dlanguage=CSharp GoParser.g4

pause