#!/usr/bin/env node
import 'source-map-support/register';
import * as cdk from '@aws-cdk/core';
import { APIStack } from '../lib/api-stack';

const app = new cdk.App();
new APIStack(app, 'api');
