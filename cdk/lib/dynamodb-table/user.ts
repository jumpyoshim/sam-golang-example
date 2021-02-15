import cdk = require('@aws-cdk/core');
import dynamodb = require('@aws-cdk/aws-dynamodb');

export const buildUser = (cons: cdk.Construct) => {
  let tableProps: Partial<dynamodb.TableProps> = {
    partitionKey: { name: 'uuid', type: dynamodb.AttributeType.STRING },
    stream: dynamodb.StreamViewType.NEW_AND_OLD_IMAGES,
    billingMode: dynamodb.BillingMode.PAY_PER_REQUEST,
  };

  tableProps = {
    ...tableProps,
    tableName: `user`,
    removalPolicy: cdk.RemovalPolicy.DESTROY,
  };

  const table = new dynamodb.Table(cons, 'User', tableProps as dynamodb.TableProps);

  new cdk.CfnOutput(cons, 'UserTableName', {
    value: table.tableName,
    exportName: `dynamodb-user-table-name`,
  });
  new cdk.CfnOutput(cons, 'UserTableArn', {
    value: table.tableArn,
    exportName: `dynamodb-user-table-arn`,
  });
  new cdk.CfnOutput(cons, 'UserTableStreamArn', {
    value: table.tableStreamArn!,
    exportName: `dynamodb-user-table-stream-arn`,
  });
};
