import * as cdk from '@aws-cdk/core';
import * as apigateway from '@aws-cdk/aws-apigateway';
import * as lambda from '@aws-cdk/aws-lambda';
import * as iam from '@aws-cdk/aws-iam';

export const buildUser = (
  stack: cdk.Stack,
  path: apigateway.IResource,
) => {
  const role = new iam.Role(stack, 'UserRole', {
    assumedBy: new iam.ServicePrincipal('lambda.amazonaws.com'),
    managedPolicies: [
      iam.ManagedPolicy.fromAwsManagedPolicyName('service-role/AWSLambdaBasicExecutionRole'),
      iam.ManagedPolicy.fromAwsManagedPolicyName('AWSXRayDaemonWriteAccess'),
      iam.ManagedPolicy.fromAwsManagedPolicyName('CloudWatchLambdaInsightsExecutionRolePolicy'),
    ],
  });
  role.addToPolicy(new iam.PolicyStatement({
    actions: ['dynamodb:*'],
    resources: ['*'],
  }));

  const handler = new lambda.Function(stack, 'UserFunction', {
    functionName: `PrivateAPIUser`,
    runtime: lambda.Runtime.GO_1_X,
    handler: 'main',
    code: lambda.Code.fromAsset('../build/funcs/api-private-user'),
    memorySize: 512,
    timeout: cdk.Duration.seconds(120),
    tracing: lambda.Tracing.ACTIVE,
    role: role,
    environment: {
      APP_ENV: "prod"
    }
  });

  const resource = path.addResource('user')
  const integration = new apigateway.LambdaIntegration(handler)
  resource.addMethod('ANY', integration)
  resource.addProxy({ defaultIntegration: integration });
};
