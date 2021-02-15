import cdk = require('@aws-cdk/core');
import dynamodb = require('@aws-cdk/aws-dynamodb');

import {
  buildUser,
} from './dynamodb-table';

export class DynamoDBTableStack extends cdk.Stack {
  constructor(scope: cdk.Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    buildUser(this);
  }
}
