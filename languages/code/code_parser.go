// Code generated from Code.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Code

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 14, 83, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 3, 2, 5,
	2, 26, 10, 2, 3, 2, 7, 2, 29, 10, 2, 12, 2, 14, 2, 32, 11, 2, 3, 2, 7,
	2, 35, 10, 2, 12, 2, 14, 2, 38, 11, 2, 3, 2, 7, 2, 41, 10, 2, 12, 2, 14,
	2, 44, 11, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 5, 3,
	5, 3, 5, 5, 5, 57, 10, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	8, 7, 8, 67, 10, 8, 12, 8, 14, 8, 70, 11, 8, 3, 9, 3, 9, 3, 10, 3, 10,
	3, 10, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 2, 2, 13, 2, 4,
	6, 8, 10, 12, 14, 16, 18, 20, 22, 2, 3, 4, 2, 10, 10, 14, 14, 2, 78, 2,
	25, 3, 2, 2, 2, 4, 47, 3, 2, 2, 2, 6, 50, 3, 2, 2, 2, 8, 56, 3, 2, 2, 2,
	10, 58, 3, 2, 2, 2, 12, 60, 3, 2, 2, 2, 14, 68, 3, 2, 2, 2, 16, 71, 3,
	2, 2, 2, 18, 73, 3, 2, 2, 2, 20, 76, 3, 2, 2, 2, 22, 79, 3, 2, 2, 2, 24,
	26, 5, 4, 3, 2, 25, 24, 3, 2, 2, 2, 25, 26, 3, 2, 2, 2, 26, 30, 3, 2, 2,
	2, 27, 29, 5, 6, 4, 2, 28, 27, 3, 2, 2, 2, 29, 32, 3, 2, 2, 2, 30, 28,
	3, 2, 2, 2, 30, 31, 3, 2, 2, 2, 31, 36, 3, 2, 2, 2, 32, 30, 3, 2, 2, 2,
	33, 35, 5, 8, 5, 2, 34, 33, 3, 2, 2, 2, 35, 38, 3, 2, 2, 2, 36, 34, 3,
	2, 2, 2, 36, 37, 3, 2, 2, 2, 37, 42, 3, 2, 2, 2, 38, 36, 3, 2, 2, 2, 39,
	41, 5, 10, 6, 2, 40, 39, 3, 2, 2, 2, 41, 44, 3, 2, 2, 2, 42, 40, 3, 2,
	2, 2, 42, 43, 3, 2, 2, 2, 43, 45, 3, 2, 2, 2, 44, 42, 3, 2, 2, 2, 45, 46,
	7, 2, 2, 3, 46, 3, 3, 2, 2, 2, 47, 48, 7, 5, 2, 2, 48, 49, 7, 10, 2, 2,
	49, 5, 3, 2, 2, 2, 50, 51, 7, 6, 2, 2, 51, 52, 7, 10, 2, 2, 52, 7, 3, 2,
	2, 2, 53, 57, 5, 18, 10, 2, 54, 57, 5, 20, 11, 2, 55, 57, 5, 22, 12, 2,
	56, 53, 3, 2, 2, 2, 56, 54, 3, 2, 2, 2, 56, 55, 3, 2, 2, 2, 57, 9, 3, 2,
	2, 2, 58, 59, 5, 12, 7, 2, 59, 11, 3, 2, 2, 2, 60, 61, 7, 10, 2, 2, 61,
	62, 7, 3, 2, 2, 62, 63, 5, 14, 8, 2, 63, 64, 7, 4, 2, 2, 64, 13, 3, 2,
	2, 2, 65, 67, 5, 16, 9, 2, 66, 65, 3, 2, 2, 2, 67, 70, 3, 2, 2, 2, 68,
	66, 3, 2, 2, 2, 68, 69, 3, 2, 2, 2, 69, 15, 3, 2, 2, 2, 70, 68, 3, 2, 2,
	2, 71, 72, 9, 2, 2, 2, 72, 17, 3, 2, 2, 2, 73, 74, 7, 7, 2, 2, 74, 75,
	7, 10, 2, 2, 75, 19, 3, 2, 2, 2, 76, 77, 7, 8, 2, 2, 77, 78, 7, 10, 2,
	2, 78, 21, 3, 2, 2, 2, 79, 80, 7, 9, 2, 2, 80, 81, 7, 10, 2, 2, 81, 23,
	3, 2, 2, 2, 8, 25, 30, 36, 42, 56, 68,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'('", "')'", "'package'", "'import'", "'struct'", "'member'", "'function'",
}
var symbolicNames = []string{
	"", "", "", "PACKAGE", "IMPORT", "DATA_STRUCT", "MEMBER", "FUNCTION", "IDENTIFIER",
	"WS", "COMMENT", "LINE_COMMENT", "STRING_LITERAL",
}

var ruleNames = []string{
	"compilationUnit", "packageDeclaration", "importDeclaration", "typeDeclaration",
	"expressDeclaration", "methodCallDeclaration", "parameterList", "parameter",
	"dataStructDeclaration", "memberDeclaration", "functionDeclaration",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type CodeParser struct {
	*antlr.BaseParser
}

func NewCodeParser(input antlr.TokenStream) *CodeParser {
	this := new(CodeParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Code.g4"

	return this
}

// CodeParser tokens.
const (
	CodeParserEOF            = antlr.TokenEOF
	CodeParserT__0           = 1
	CodeParserT__1           = 2
	CodeParserPACKAGE        = 3
	CodeParserIMPORT         = 4
	CodeParserDATA_STRUCT    = 5
	CodeParserMEMBER         = 6
	CodeParserFUNCTION       = 7
	CodeParserIDENTIFIER     = 8
	CodeParserWS             = 9
	CodeParserCOMMENT        = 10
	CodeParserLINE_COMMENT   = 11
	CodeParserSTRING_LITERAL = 12
)

// CodeParser rules.
const (
	CodeParserRULE_compilationUnit       = 0
	CodeParserRULE_packageDeclaration    = 1
	CodeParserRULE_importDeclaration     = 2
	CodeParserRULE_typeDeclaration       = 3
	CodeParserRULE_expressDeclaration    = 4
	CodeParserRULE_methodCallDeclaration = 5
	CodeParserRULE_parameterList         = 6
	CodeParserRULE_parameter             = 7
	CodeParserRULE_dataStructDeclaration = 8
	CodeParserRULE_memberDeclaration     = 9
	CodeParserRULE_functionDeclaration   = 10
)

// ICompilationUnitContext is an interface to support dynamic dispatch.
type ICompilationUnitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCompilationUnitContext differentiates from other interfaces.
	IsCompilationUnitContext()
}

type CompilationUnitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompilationUnitContext() *CompilationUnitContext {
	var p = new(CompilationUnitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CodeParserRULE_compilationUnit
	return p
}

func (*CompilationUnitContext) IsCompilationUnitContext() {}

func NewCompilationUnitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompilationUnitContext {
	var p = new(CompilationUnitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CodeParserRULE_compilationUnit

	return p
}

func (s *CompilationUnitContext) GetParser() antlr.Parser { return s.parser }

func (s *CompilationUnitContext) EOF() antlr.TerminalNode {
	return s.GetToken(CodeParserEOF, 0)
}

func (s *CompilationUnitContext) PackageDeclaration() IPackageDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPackageDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPackageDeclarationContext)
}

func (s *CompilationUnitContext) AllImportDeclaration() []IImportDeclarationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IImportDeclarationContext)(nil)).Elem())
	var tst = make([]IImportDeclarationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IImportDeclarationContext)
		}
	}

	return tst
}

func (s *CompilationUnitContext) ImportDeclaration(i int) IImportDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImportDeclarationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IImportDeclarationContext)
}

func (s *CompilationUnitContext) AllTypeDeclaration() []ITypeDeclarationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypeDeclarationContext)(nil)).Elem())
	var tst = make([]ITypeDeclarationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeDeclarationContext)
		}
	}

	return tst
}

func (s *CompilationUnitContext) TypeDeclaration(i int) ITypeDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeDeclarationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeDeclarationContext)
}

func (s *CompilationUnitContext) AllExpressDeclaration() []IExpressDeclarationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressDeclarationContext)(nil)).Elem())
	var tst = make([]IExpressDeclarationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressDeclarationContext)
		}
	}

	return tst
}

func (s *CompilationUnitContext) ExpressDeclaration(i int) IExpressDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressDeclarationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressDeclarationContext)
}

func (s *CompilationUnitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompilationUnitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CompilationUnitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeListener); ok {
		listenerT.EnterCompilationUnit(s)
	}
}

func (s *CompilationUnitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeListener); ok {
		listenerT.ExitCompilationUnit(s)
	}
}

func (s *CompilationUnitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CodeVisitor:
		return t.VisitCompilationUnit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CodeParser) CompilationUnit() (localctx ICompilationUnitContext) {
	localctx = NewCompilationUnitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CodeParserRULE_compilationUnit)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(23)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CodeParserPACKAGE {
		{
			p.SetState(22)
			p.PackageDeclaration()
		}

	}
	p.SetState(28)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CodeParserIMPORT {
		{
			p.SetState(25)
			p.ImportDeclaration()
		}

		p.SetState(30)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(34)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CodeParserDATA_STRUCT)|(1<<CodeParserMEMBER)|(1<<CodeParserFUNCTION))) != 0 {
		{
			p.SetState(31)
			p.TypeDeclaration()
		}

		p.SetState(36)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(40)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CodeParserIDENTIFIER {
		{
			p.SetState(37)
			p.ExpressDeclaration()
		}

		p.SetState(42)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(43)
		p.Match(CodeParserEOF)
	}

	return localctx
}

// IPackageDeclarationContext is an interface to support dynamic dispatch.
type IPackageDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPackageDeclarationContext differentiates from other interfaces.
	IsPackageDeclarationContext()
}

type PackageDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPackageDeclarationContext() *PackageDeclarationContext {
	var p = new(PackageDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CodeParserRULE_packageDeclaration
	return p
}

func (*PackageDeclarationContext) IsPackageDeclarationContext() {}

func NewPackageDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PackageDeclarationContext {
	var p = new(PackageDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CodeParserRULE_packageDeclaration

	return p
}

func (s *PackageDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *PackageDeclarationContext) PACKAGE() antlr.TerminalNode {
	return s.GetToken(CodeParserPACKAGE, 0)
}

func (s *PackageDeclarationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CodeParserIDENTIFIER, 0)
}

func (s *PackageDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PackageDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PackageDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeListener); ok {
		listenerT.EnterPackageDeclaration(s)
	}
}

func (s *PackageDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeListener); ok {
		listenerT.ExitPackageDeclaration(s)
	}
}

func (s *PackageDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CodeVisitor:
		return t.VisitPackageDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CodeParser) PackageDeclaration() (localctx IPackageDeclarationContext) {
	localctx = NewPackageDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CodeParserRULE_packageDeclaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(45)
		p.Match(CodeParserPACKAGE)
	}
	{
		p.SetState(46)
		p.Match(CodeParserIDENTIFIER)
	}

	return localctx
}

// IImportDeclarationContext is an interface to support dynamic dispatch.
type IImportDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImportDeclarationContext differentiates from other interfaces.
	IsImportDeclarationContext()
}

type ImportDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportDeclarationContext() *ImportDeclarationContext {
	var p = new(ImportDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CodeParserRULE_importDeclaration
	return p
}

func (*ImportDeclarationContext) IsImportDeclarationContext() {}

func NewImportDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportDeclarationContext {
	var p = new(ImportDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CodeParserRULE_importDeclaration

	return p
}

func (s *ImportDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportDeclarationContext) IMPORT() antlr.TerminalNode {
	return s.GetToken(CodeParserIMPORT, 0)
}

func (s *ImportDeclarationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CodeParserIDENTIFIER, 0)
}

func (s *ImportDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeListener); ok {
		listenerT.EnterImportDeclaration(s)
	}
}

func (s *ImportDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeListener); ok {
		listenerT.ExitImportDeclaration(s)
	}
}

func (s *ImportDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CodeVisitor:
		return t.VisitImportDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CodeParser) ImportDeclaration() (localctx IImportDeclarationContext) {
	localctx = NewImportDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, CodeParserRULE_importDeclaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(48)
		p.Match(CodeParserIMPORT)
	}
	{
		p.SetState(49)
		p.Match(CodeParserIDENTIFIER)
	}

	return localctx
}

// ITypeDeclarationContext is an interface to support dynamic dispatch.
type ITypeDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeDeclarationContext differentiates from other interfaces.
	IsTypeDeclarationContext()
}

type TypeDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeDeclarationContext() *TypeDeclarationContext {
	var p = new(TypeDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CodeParserRULE_typeDeclaration
	return p
}

func (*TypeDeclarationContext) IsTypeDeclarationContext() {}

func NewTypeDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeDeclarationContext {
	var p = new(TypeDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CodeParserRULE_typeDeclaration

	return p
}

func (s *TypeDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeDeclarationContext) DataStructDeclaration() IDataStructDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDataStructDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDataStructDeclarationContext)
}

func (s *TypeDeclarationContext) MemberDeclaration() IMemberDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMemberDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMemberDeclarationContext)
}

func (s *TypeDeclarationContext) FunctionDeclaration() IFunctionDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionDeclarationContext)
}

func (s *TypeDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeListener); ok {
		listenerT.EnterTypeDeclaration(s)
	}
}

func (s *TypeDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeListener); ok {
		listenerT.ExitTypeDeclaration(s)
	}
}

func (s *TypeDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CodeVisitor:
		return t.VisitTypeDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CodeParser) TypeDeclaration() (localctx ITypeDeclarationContext) {
	localctx = NewTypeDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, CodeParserRULE_typeDeclaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(54)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CodeParserDATA_STRUCT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(51)
			p.DataStructDeclaration()
		}

	case CodeParserMEMBER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(52)
			p.MemberDeclaration()
		}

	case CodeParserFUNCTION:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(53)
			p.FunctionDeclaration()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IExpressDeclarationContext is an interface to support dynamic dispatch.
type IExpressDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressDeclarationContext differentiates from other interfaces.
	IsExpressDeclarationContext()
}

type ExpressDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressDeclarationContext() *ExpressDeclarationContext {
	var p = new(ExpressDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CodeParserRULE_expressDeclaration
	return p
}

func (*ExpressDeclarationContext) IsExpressDeclarationContext() {}

func NewExpressDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressDeclarationContext {
	var p = new(ExpressDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CodeParserRULE_expressDeclaration

	return p
}

func (s *ExpressDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressDeclarationContext) MethodCallDeclaration() IMethodCallDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMethodCallDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMethodCallDeclarationContext)
}

func (s *ExpressDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeListener); ok {
		listenerT.EnterExpressDeclaration(s)
	}
}

func (s *ExpressDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeListener); ok {
		listenerT.ExitExpressDeclaration(s)
	}
}

func (s *ExpressDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CodeVisitor:
		return t.VisitExpressDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CodeParser) ExpressDeclaration() (localctx IExpressDeclarationContext) {
	localctx = NewExpressDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, CodeParserRULE_expressDeclaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(56)
		p.MethodCallDeclaration()
	}

	return localctx
}

// IMethodCallDeclarationContext is an interface to support dynamic dispatch.
type IMethodCallDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMethodCallDeclarationContext differentiates from other interfaces.
	IsMethodCallDeclarationContext()
}

type MethodCallDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMethodCallDeclarationContext() *MethodCallDeclarationContext {
	var p = new(MethodCallDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CodeParserRULE_methodCallDeclaration
	return p
}

func (*MethodCallDeclarationContext) IsMethodCallDeclarationContext() {}

func NewMethodCallDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodCallDeclarationContext {
	var p = new(MethodCallDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CodeParserRULE_methodCallDeclaration

	return p
}

func (s *MethodCallDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *MethodCallDeclarationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CodeParserIDENTIFIER, 0)
}

func (s *MethodCallDeclarationContext) ParameterList() IParameterListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParameterListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParameterListContext)
}

func (s *MethodCallDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodCallDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MethodCallDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeListener); ok {
		listenerT.EnterMethodCallDeclaration(s)
	}
}

func (s *MethodCallDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeListener); ok {
		listenerT.ExitMethodCallDeclaration(s)
	}
}

func (s *MethodCallDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CodeVisitor:
		return t.VisitMethodCallDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CodeParser) MethodCallDeclaration() (localctx IMethodCallDeclarationContext) {
	localctx = NewMethodCallDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, CodeParserRULE_methodCallDeclaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(58)
		p.Match(CodeParserIDENTIFIER)
	}
	{
		p.SetState(59)
		p.Match(CodeParserT__0)
	}
	{
		p.SetState(60)
		p.ParameterList()
	}
	{
		p.SetState(61)
		p.Match(CodeParserT__1)
	}

	return localctx
}

// IParameterListContext is an interface to support dynamic dispatch.
type IParameterListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParameterListContext differentiates from other interfaces.
	IsParameterListContext()
}

type ParameterListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterListContext() *ParameterListContext {
	var p = new(ParameterListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CodeParserRULE_parameterList
	return p
}

func (*ParameterListContext) IsParameterListContext() {}

func NewParameterListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterListContext {
	var p = new(ParameterListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CodeParserRULE_parameterList

	return p
}

func (s *ParameterListContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterListContext) AllParameter() []IParameterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IParameterContext)(nil)).Elem())
	var tst = make([]IParameterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IParameterContext)
		}
	}

	return tst
}

func (s *ParameterListContext) Parameter(i int) IParameterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParameterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IParameterContext)
}

func (s *ParameterListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeListener); ok {
		listenerT.EnterParameterList(s)
	}
}

func (s *ParameterListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeListener); ok {
		listenerT.ExitParameterList(s)
	}
}

func (s *ParameterListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CodeVisitor:
		return t.VisitParameterList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CodeParser) ParameterList() (localctx IParameterListContext) {
	localctx = NewParameterListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, CodeParserRULE_parameterList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(66)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CodeParserIDENTIFIER || _la == CodeParserSTRING_LITERAL {
		{
			p.SetState(63)
			p.Parameter()
		}

		p.SetState(68)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IParameterContext is an interface to support dynamic dispatch.
type IParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParameterContext differentiates from other interfaces.
	IsParameterContext()
}

type ParameterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterContext() *ParameterContext {
	var p = new(ParameterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CodeParserRULE_parameter
	return p
}

func (*ParameterContext) IsParameterContext() {}

func NewParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterContext {
	var p = new(ParameterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CodeParserRULE_parameter

	return p
}

func (s *ParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CodeParserIDENTIFIER, 0)
}

func (s *ParameterContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(CodeParserSTRING_LITERAL, 0)
}

func (s *ParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeListener); ok {
		listenerT.EnterParameter(s)
	}
}

func (s *ParameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeListener); ok {
		listenerT.ExitParameter(s)
	}
}

func (s *ParameterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CodeVisitor:
		return t.VisitParameter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CodeParser) Parameter() (localctx IParameterContext) {
	localctx = NewParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, CodeParserRULE_parameter)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(69)
		_la = p.GetTokenStream().LA(1)

		if !(_la == CodeParserIDENTIFIER || _la == CodeParserSTRING_LITERAL) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IDataStructDeclarationContext is an interface to support dynamic dispatch.
type IDataStructDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDataStructDeclarationContext differentiates from other interfaces.
	IsDataStructDeclarationContext()
}

type DataStructDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDataStructDeclarationContext() *DataStructDeclarationContext {
	var p = new(DataStructDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CodeParserRULE_dataStructDeclaration
	return p
}

func (*DataStructDeclarationContext) IsDataStructDeclarationContext() {}

func NewDataStructDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DataStructDeclarationContext {
	var p = new(DataStructDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CodeParserRULE_dataStructDeclaration

	return p
}

func (s *DataStructDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DataStructDeclarationContext) DATA_STRUCT() antlr.TerminalNode {
	return s.GetToken(CodeParserDATA_STRUCT, 0)
}

func (s *DataStructDeclarationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CodeParserIDENTIFIER, 0)
}

func (s *DataStructDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DataStructDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DataStructDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeListener); ok {
		listenerT.EnterDataStructDeclaration(s)
	}
}

func (s *DataStructDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeListener); ok {
		listenerT.ExitDataStructDeclaration(s)
	}
}

func (s *DataStructDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CodeVisitor:
		return t.VisitDataStructDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CodeParser) DataStructDeclaration() (localctx IDataStructDeclarationContext) {
	localctx = NewDataStructDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, CodeParserRULE_dataStructDeclaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(71)
		p.Match(CodeParserDATA_STRUCT)
	}
	{
		p.SetState(72)
		p.Match(CodeParserIDENTIFIER)
	}

	return localctx
}

// IMemberDeclarationContext is an interface to support dynamic dispatch.
type IMemberDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMemberDeclarationContext differentiates from other interfaces.
	IsMemberDeclarationContext()
}

type MemberDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMemberDeclarationContext() *MemberDeclarationContext {
	var p = new(MemberDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CodeParserRULE_memberDeclaration
	return p
}

func (*MemberDeclarationContext) IsMemberDeclarationContext() {}

func NewMemberDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MemberDeclarationContext {
	var p = new(MemberDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CodeParserRULE_memberDeclaration

	return p
}

func (s *MemberDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *MemberDeclarationContext) MEMBER() antlr.TerminalNode {
	return s.GetToken(CodeParserMEMBER, 0)
}

func (s *MemberDeclarationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CodeParserIDENTIFIER, 0)
}

func (s *MemberDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemberDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MemberDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeListener); ok {
		listenerT.EnterMemberDeclaration(s)
	}
}

func (s *MemberDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeListener); ok {
		listenerT.ExitMemberDeclaration(s)
	}
}

func (s *MemberDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CodeVisitor:
		return t.VisitMemberDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CodeParser) MemberDeclaration() (localctx IMemberDeclarationContext) {
	localctx = NewMemberDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, CodeParserRULE_memberDeclaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(74)
		p.Match(CodeParserMEMBER)
	}
	{
		p.SetState(75)
		p.Match(CodeParserIDENTIFIER)
	}

	return localctx
}

// IFunctionDeclarationContext is an interface to support dynamic dispatch.
type IFunctionDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionDeclarationContext differentiates from other interfaces.
	IsFunctionDeclarationContext()
}

type FunctionDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionDeclarationContext() *FunctionDeclarationContext {
	var p = new(FunctionDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CodeParserRULE_functionDeclaration
	return p
}

func (*FunctionDeclarationContext) IsFunctionDeclarationContext() {}

func NewFunctionDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionDeclarationContext {
	var p = new(FunctionDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CodeParserRULE_functionDeclaration

	return p
}

func (s *FunctionDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionDeclarationContext) FUNCTION() antlr.TerminalNode {
	return s.GetToken(CodeParserFUNCTION, 0)
}

func (s *FunctionDeclarationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CodeParserIDENTIFIER, 0)
}

func (s *FunctionDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeListener); ok {
		listenerT.EnterFunctionDeclaration(s)
	}
}

func (s *FunctionDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CodeListener); ok {
		listenerT.ExitFunctionDeclaration(s)
	}
}

func (s *FunctionDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CodeVisitor:
		return t.VisitFunctionDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CodeParser) FunctionDeclaration() (localctx IFunctionDeclarationContext) {
	localctx = NewFunctionDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, CodeParserRULE_functionDeclaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(77)
		p.Match(CodeParserFUNCTION)
	}
	{
		p.SetState(78)
		p.Match(CodeParserIDENTIFIER)
	}

	return localctx
}