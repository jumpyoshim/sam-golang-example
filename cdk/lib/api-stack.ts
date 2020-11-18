import * as apigateway from '@aws-cdk/aws-apigateway';
import * as cdk from '@aws-cdk/core';

import {
  buildUser,
} from './api-stack/index'

export class APIStack extends cdk.Stack {
  constructor(scope: cdk.Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    const api = new apigateway.RestApi(this, "api", {
      restApiName: "api",
    });

    api.root.addMethod('GET')
    buildUser(this, api.root)
  }
}
