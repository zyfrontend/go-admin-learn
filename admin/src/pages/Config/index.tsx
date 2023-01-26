import React, { memo } from 'react';
import type { ReactNode, FC } from 'react';
import { Card, Form, Input, Tabs } from 'antd';
import { useIntl } from '@umijs/max';
import { PageContainer } from '@ant-design/pro-components';
// import type { TabsProps } from 'antd';

interface IProps {
  children?: ReactNode;
}
const USDTConfig = () => {
  return (
    <Form>
      <Form.Item label="环境名称">
        <Input placeholder="测试网" />
      </Form.Item>
      <Form.Item label="合约地址">
        <Input placeholder="Contract" />
      </Form.Item>
      <Form.Item label="ABI地址">
        <Input placeholder="ABI" />
      </Form.Item>
      <Form.Item label="单位">
        <Input placeholder="代币单位" />
      </Form.Item>
      <Form.Item label="代币小数位">
        <Input placeholder="代币小数位" />
      </Form.Item>
    </Form>
  );
};
const DefaultConfig = () => {
  return (
    <Form>
      <Form.Item label="环境名称">
        <Input placeholder="测试网" />
      </Form.Item>
      <Form.Item label="合约地址">
        <Input placeholder="Contract" />
      </Form.Item>
      <Form.Item label="ABI地址">
        <Input placeholder="ABI" />
      </Form.Item>
      <Form.Item label="单位">
        <Input placeholder="代币单位" />
      </Form.Item>
      <Form.Item label="代币小数位">
        <Input placeholder="代币小数位" />
      </Form.Item>
    </Form>
  );
};
const Config: FC<IProps> = () => {
  const intl = useIntl();
  const onChange = () => {
    console.log('aaa');
  };
  return (
    <PageContainer
      content={intl.formatMessage({
        id: 'pages.admin.subPage.config',
        defaultMessage: 'config',
      })}
    >
      <Card size="small">
        <Tabs defaultActiveKey="1" onChange={onChange}></Tabs>
        <DefaultConfig />
      </Card>
      <Card size="small">
        <br />
        <USDTConfig />
      </Card>
    </PageContainer>
  );
};

export default memo(Config);
