import * as apigateway from '@aws-cdk/aws-apigateway';
import * as cdk from '@aws-cdk/core';

import {
  buildUser,
} from './api-private-stack/index'

export class PrivateAPIStack extends cdk.Stack {
  constructor(scope: cdk.Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    const api = new apigateway.RestApi(this, "PrivateAPI", {
      restApiName: "PrivateAPI",
    });

    api.root.addMethod('GET')
    buildUser(this, api.root)
  }
}
