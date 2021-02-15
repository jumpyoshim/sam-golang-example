#!/usr/bin/env node
import 'source-map-support/register';
import * as cdk from '@aws-cdk/core';
import { PrivateAPIStack } from '../lib/api-private-stack';
import { DynamoDBTableStack } from '../lib/dynamodb-table-stack';

const app = new cdk.App();
new PrivateAPIStack(app, 'api-private');
new DynamoDBTableStack(app, 'dynamodb-table');
