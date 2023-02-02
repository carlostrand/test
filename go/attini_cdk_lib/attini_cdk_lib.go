package attini_cdk_lib

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"attini-cdk-lib.AttiniMerge",
		reflect.TypeOf((*AttiniMerge)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addBranch", GoMethod: "AddBranch"},
			_jsii_.MemberMethod{JsiiMethod: "addChoice", GoMethod: "AddChoice"},
			_jsii_.MemberMethod{JsiiMethod: "addIterator", GoMethod: "AddIterator"},
			_jsii_.MemberMethod{JsiiMethod: "addPrefix", GoMethod: "AddPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "bindToGraph", GoMethod: "BindToGraph"},
			_jsii_.MemberProperty{JsiiProperty: "branches", GoGetter: "Branches"},
			_jsii_.MemberProperty{JsiiProperty: "comment", GoGetter: "Comment"},
			_jsii_.MemberProperty{JsiiProperty: "defaultChoice", GoGetter: "DefaultChoice"},
			_jsii_.MemberProperty{JsiiProperty: "endStates", GoGetter: "EndStates"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "inputPath", GoGetter: "InputPath"},
			_jsii_.MemberProperty{JsiiProperty: "iteration", GoGetter: "Iteration"},
			_jsii_.MemberMethod{JsiiMethod: "makeDefault", GoMethod: "MakeDefault"},
			_jsii_.MemberMethod{JsiiMethod: "makeNext", GoMethod: "MakeNext"},
			_jsii_.MemberMethod{JsiiMethod: "next", GoMethod: "Next"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "outputPath", GoGetter: "OutputPath"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberMethod{JsiiMethod: "renderBranches", GoMethod: "RenderBranches"},
			_jsii_.MemberMethod{JsiiMethod: "renderChoices", GoMethod: "RenderChoices"},
			_jsii_.MemberMethod{JsiiMethod: "renderInputOutput", GoMethod: "RenderInputOutput"},
			_jsii_.MemberMethod{JsiiMethod: "renderIterator", GoMethod: "RenderIterator"},
			_jsii_.MemberMethod{JsiiMethod: "renderNextEnd", GoMethod: "RenderNextEnd"},
			_jsii_.MemberMethod{JsiiMethod: "renderProps", GoMethod: "RenderProps"},
			_jsii_.MemberMethod{JsiiMethod: "renderResultSelector", GoMethod: "RenderResultSelector"},
			_jsii_.MemberMethod{JsiiMethod: "renderRetryCatch", GoMethod: "RenderRetryCatch"},
			_jsii_.MemberProperty{JsiiProperty: "resultPath", GoGetter: "ResultPath"},
			_jsii_.MemberProperty{JsiiProperty: "resultSelector", GoGetter: "ResultSelector"},
			_jsii_.MemberProperty{JsiiProperty: "startState", GoGetter: "StartState"},
			_jsii_.MemberProperty{JsiiProperty: "stateId", GoGetter: "StateId"},
			_jsii_.MemberMethod{JsiiMethod: "toStateJson", GoMethod: "ToStateJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberMethod{JsiiMethod: "validateState", GoMethod: "ValidateState"},
			_jsii_.MemberMethod{JsiiMethod: "whenBoundToGraph", GoMethod: "WhenBoundToGraph"},
		},
		func() interface{} {
			j := jsiiProxy_AttiniMerge{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_AttiniState)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"attini-cdk-lib.AttiniRunner",
		reflect.TypeOf((*AttiniRunner)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "runnerName", GoGetter: "RunnerName"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AttiniRunner{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"attini-cdk-lib.AttiniRunnerJob",
		reflect.TypeOf((*AttiniRunnerJob)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addBranch", GoMethod: "AddBranch"},
			_jsii_.MemberMethod{JsiiMethod: "addChoice", GoMethod: "AddChoice"},
			_jsii_.MemberMethod{JsiiMethod: "addIterator", GoMethod: "AddIterator"},
			_jsii_.MemberMethod{JsiiMethod: "addPrefix", GoMethod: "AddPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "bindToGraph", GoMethod: "BindToGraph"},
			_jsii_.MemberProperty{JsiiProperty: "branches", GoGetter: "Branches"},
			_jsii_.MemberProperty{JsiiProperty: "comment", GoGetter: "Comment"},
			_jsii_.MemberProperty{JsiiProperty: "defaultChoice", GoGetter: "DefaultChoice"},
			_jsii_.MemberProperty{JsiiProperty: "endStates", GoGetter: "EndStates"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "inputPath", GoGetter: "InputPath"},
			_jsii_.MemberProperty{JsiiProperty: "iteration", GoGetter: "Iteration"},
			_jsii_.MemberMethod{JsiiMethod: "makeDefault", GoMethod: "MakeDefault"},
			_jsii_.MemberMethod{JsiiMethod: "makeNext", GoMethod: "MakeNext"},
			_jsii_.MemberMethod{JsiiMethod: "next", GoMethod: "Next"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "outputPath", GoGetter: "OutputPath"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberMethod{JsiiMethod: "renderBranches", GoMethod: "RenderBranches"},
			_jsii_.MemberMethod{JsiiMethod: "renderChoices", GoMethod: "RenderChoices"},
			_jsii_.MemberMethod{JsiiMethod: "renderInputOutput", GoMethod: "RenderInputOutput"},
			_jsii_.MemberMethod{JsiiMethod: "renderIterator", GoMethod: "RenderIterator"},
			_jsii_.MemberMethod{JsiiMethod: "renderNextEnd", GoMethod: "RenderNextEnd"},
			_jsii_.MemberMethod{JsiiMethod: "renderProps", GoMethod: "RenderProps"},
			_jsii_.MemberMethod{JsiiMethod: "renderResultSelector", GoMethod: "RenderResultSelector"},
			_jsii_.MemberMethod{JsiiMethod: "renderRetryCatch", GoMethod: "RenderRetryCatch"},
			_jsii_.MemberProperty{JsiiProperty: "resultPath", GoGetter: "ResultPath"},
			_jsii_.MemberProperty{JsiiProperty: "resultSelector", GoGetter: "ResultSelector"},
			_jsii_.MemberProperty{JsiiProperty: "startState", GoGetter: "StartState"},
			_jsii_.MemberProperty{JsiiProperty: "stateId", GoGetter: "StateId"},
			_jsii_.MemberMethod{JsiiMethod: "toStateJson", GoMethod: "ToStateJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberMethod{JsiiMethod: "validateState", GoMethod: "ValidateState"},
			_jsii_.MemberMethod{JsiiMethod: "whenBoundToGraph", GoMethod: "WhenBoundToGraph"},
		},
		func() interface{} {
			j := jsiiProxy_AttiniRunnerJob{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_AttiniState)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"attini-cdk-lib.AttiniRunnerJobProps",
		reflect.TypeOf((*AttiniRunnerJobProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"attini-cdk-lib.AttiniRunnerProps",
		reflect.TypeOf((*AttiniRunnerProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"attini-cdk-lib.AttiniState",
		reflect.TypeOf((*AttiniState)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addBranch", GoMethod: "AddBranch"},
			_jsii_.MemberMethod{JsiiMethod: "addChoice", GoMethod: "AddChoice"},
			_jsii_.MemberMethod{JsiiMethod: "addIterator", GoMethod: "AddIterator"},
			_jsii_.MemberMethod{JsiiMethod: "addPrefix", GoMethod: "AddPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "bindToGraph", GoMethod: "BindToGraph"},
			_jsii_.MemberProperty{JsiiProperty: "branches", GoGetter: "Branches"},
			_jsii_.MemberProperty{JsiiProperty: "comment", GoGetter: "Comment"},
			_jsii_.MemberProperty{JsiiProperty: "defaultChoice", GoGetter: "DefaultChoice"},
			_jsii_.MemberProperty{JsiiProperty: "endStates", GoGetter: "EndStates"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "inputPath", GoGetter: "InputPath"},
			_jsii_.MemberProperty{JsiiProperty: "iteration", GoGetter: "Iteration"},
			_jsii_.MemberMethod{JsiiMethod: "makeDefault", GoMethod: "MakeDefault"},
			_jsii_.MemberMethod{JsiiMethod: "makeNext", GoMethod: "MakeNext"},
			_jsii_.MemberMethod{JsiiMethod: "next", GoMethod: "Next"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "outputPath", GoGetter: "OutputPath"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberMethod{JsiiMethod: "renderBranches", GoMethod: "RenderBranches"},
			_jsii_.MemberMethod{JsiiMethod: "renderChoices", GoMethod: "RenderChoices"},
			_jsii_.MemberMethod{JsiiMethod: "renderInputOutput", GoMethod: "RenderInputOutput"},
			_jsii_.MemberMethod{JsiiMethod: "renderIterator", GoMethod: "RenderIterator"},
			_jsii_.MemberMethod{JsiiMethod: "renderNextEnd", GoMethod: "RenderNextEnd"},
			_jsii_.MemberMethod{JsiiMethod: "renderProps", GoMethod: "RenderProps"},
			_jsii_.MemberMethod{JsiiMethod: "renderResultSelector", GoMethod: "RenderResultSelector"},
			_jsii_.MemberMethod{JsiiMethod: "renderRetryCatch", GoMethod: "RenderRetryCatch"},
			_jsii_.MemberProperty{JsiiProperty: "resultPath", GoGetter: "ResultPath"},
			_jsii_.MemberProperty{JsiiProperty: "resultSelector", GoGetter: "ResultSelector"},
			_jsii_.MemberProperty{JsiiProperty: "startState", GoGetter: "StartState"},
			_jsii_.MemberProperty{JsiiProperty: "stateId", GoGetter: "StateId"},
			_jsii_.MemberMethod{JsiiMethod: "toStateJson", GoMethod: "ToStateJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberMethod{JsiiMethod: "validateState", GoMethod: "ValidateState"},
			_jsii_.MemberMethod{JsiiMethod: "whenBoundToGraph", GoMethod: "WhenBoundToGraph"},
		},
		func() interface{} {
			j := jsiiProxy_AttiniState{}
			_jsii_.InitJsiiProxy(&j.Type__awsstepfunctionsState)
			_jsii_.InitJsiiProxy(&j.Type__awsstepfunctionsIChainable)
			_jsii_.InitJsiiProxy(&j.Type__awsstepfunctionsINextable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"attini-cdk-lib.AwsVpcConfiguration",
		reflect.TypeOf((*AwsVpcConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"attini-cdk-lib.DeploymentPlan",
		reflect.TypeOf((*DeploymentPlan)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DeploymentPlan{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"attini-cdk-lib.DeploymentPlanProps",
		reflect.TypeOf((*DeploymentPlanProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"attini-cdk-lib.RunnerConfiguration",
		reflect.TypeOf((*RunnerConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"attini-cdk-lib.Startup",
		reflect.TypeOf((*Startup)(nil)).Elem(),
	)
}
