package code

import (
	. "../../languages/code"
	. "../../model"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"reflect"
)

var currentCodeModel CodeModel
var currentFunction = CreateFunction("")
var varMaps = make(map[string]string)

func NewCodeAppListener() *CodeAppListener {
	currentCodeModel = *&CodeModel{nil, nil, nil}
	return &CodeAppListener{}
}

type CodeAppListener struct {
	BaseCodeListener
}

func (s *CodeAppListener) EnterMethodCallDeclaration(ctx *MethodCallDeclarationContext) {
	allParameters := ctx.AllParameter()
	functionName := ctx.IDENTIFIER().GetText()

	parentParentType := reflect.TypeOf(ctx.GetParent().GetParent()).String()
	switch parentParentType {
	case "*parser.TypeDeclarationContext":
		functionCall := BuildFunctionCall(allParameters, functionName)
		currentCodeModel.FunctionCalls = append(currentCodeModel.FunctionCalls, functionCall)
	case "*parser.FunctionBodyContext":
		functionCall := BuildFunctionCall(allParameters, functionName)
		currentFunction.CodeFunctionCalls = append(currentFunction.CodeFunctionCalls, functionCall)
	}
}

func BuildFunctionCall(allParameters []IParameterContext, functionName string) CodeFunctionCall {
	var parameters []CodeParameter
	for _, parameter := range allParameters {
		childType := reflect.TypeOf(parameter.GetChild(0)).String()
		paraCodeType := &CodeType{
			Type: "string",
		}

		switch childType {
		case "*antlr.TerminalNodeImpl":
			paraCodeType.Type = "type"
		default:
		}

		var paramValue = &CodeParameterValue{Value: parameter.GetText()}
		parameter := &CodeParameter{*paraCodeType, *paramValue}
		parameters = append(parameters, *parameter)
	}
	functionCall := CreateFunctionCall(functionName, parameters)
	return functionCall
}

func (s *CodeAppListener) EnterFunctionDeclaration(ctx *FunctionDeclarationContext) {
	function := CreateFunction(ctx.IDENTIFIER().GetText())
	currentFunction = function

	function.Position.StartLine = ctx.GetStart().GetLine()
	function.Position.StartLineColumn = ctx.IDENTIFIER().GetSymbol().GetColumn()
	function.Position.StopLine = ctx.GetStop().GetLine()
	function.Position.StopLineColumn = ctx.IDENTIFIER().GetSymbol().GetColumn()

	allExpressDeclaration := ctx.FunctionBody().(*FunctionBodyContext).AllExpression()
	for _, express := range allExpressDeclaration {
		expressCtx := express.(*ExpressionContext)

		firstChildCtx := expressCtx.GetChild(0)

		switch reflect.TypeOf(firstChildCtx).String() {
		case "*parser.MethodCallDeclarationContext":
			context := firstChildCtx.(*MethodCallDeclarationContext)
			functionCall := BuildFunctionCall(context.AllParameter(), context.IDENTIFIER().GetText())
			function.CodeFunctionCalls = append(function.CodeFunctionCalls, functionCall)

		case "*parser.BlockStatementContext":
			child := firstChildCtx.GetChild(0)
			childType := reflect.TypeOf(child).String()
			switch childType {
			case "*parser.LocalVariableDeclarationContext":
				context := child.(*LocalVariableDeclarationContext).GetChild(0).(*VariableDeclaratorsContext)

				s.handleLocalVariable(context)
			case "*parser.StatementContext":
				statement := s.handleStatement(child.(*StatementContext))
				fmt.Println(statement)
			}
		}
	}

	currentCodeModel.Functions = append(currentCodeModel.Functions, function)
}

type BlockStatement struct {
	Condition      string
	BlockStatement []string
}

func (s *CodeAppListener) handleStatement(statementCtx *StatementContext) BlockStatement {
	blockStatement := &BlockStatement{
		Condition:      "",
		BlockStatement: nil,
	}

	child := statementCtx.GetChild(0)
	childType := reflect.TypeOf(child).String()
	switch childType {
	case "*antlr.TerminalNodeImpl":
		nodeType := child.(*antlr.TerminalNodeImpl)
		switch nodeType.GetText() {
		case "for":
			FOR_STATEMENT_INDEX := 4
			FOR_CONTROL_INDEX := 2

			forControlText := statementCtx.GetChild(FOR_CONTROL_INDEX).(*ForControlContext).GetText()
			blockStatement.Condition = forControlText
			blockCtx := statementCtx.GetChild(FOR_STATEMENT_INDEX).GetChild(0).(*BlockContext)
			for _, statement := range blockCtx.AllBlockStatement() {
				blockStatement.BlockStatement = append(blockStatement.BlockStatement, statement.GetText())
			}
		}
	}

	return *blockStatement
}

func (s *CodeAppListener) handleLocalVariable(context *VariableDeclaratorsContext) {
	for _, varDeclarator := range context.AllVariableDeclarator() {
		varCtx := varDeclarator.(*VariableDeclaratorContext)
		ident := varCtx.VariableDeclaratorId().GetText()
		value := ""

		if varCtx.VariableInitializer() != nil {
			value = varCtx.VariableInitializer().GetText()
		}

		currentFunction.Variables[ident] = value
	}
}

func (s *CodeAppListener) EnterVariableDeclarators(ctx *VariableDeclaratorsContext) {
	for _, varDeclarator := range ctx.AllVariableDeclarator() {
		varCtx := varDeclarator.(*VariableDeclaratorContext)
		ident := varCtx.VariableDeclaratorId().GetText()
		value := ""

		if varCtx.VariableInitializer() != nil {
			value = varCtx.VariableInitializer().GetText()
		}

		varMaps[ident] = value
	}
}

func (s *CodeAppListener) getCode() CodeModel {
	currentCodeModel.Variables = varMaps
	return currentCodeModel
}
