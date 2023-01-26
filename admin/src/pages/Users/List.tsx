import React, { memo, useRef } from 'react';
import { PageContainer } from '@ant-design/pro-components';
import { useIntl } from '@umijs/max';
import { Card, Table, Button, Form, Input } from 'antd';
import type { ColumnsType } from 'antd/es/table';
import type { FormInstance } from 'antd/es/form';
import type { FC } from 'react';
import { SyncOutlined } from '@ant-design/icons';
import userModel from '@/models/userModel';
const UserList: FC = () => {
  const intl = useIntl();
  const formRef = useRef<FormInstance>(null);
  const { userList, getUserInfoModel } = userModel();
  const columns: ColumnsType<DataType> = [
    {
      title: 'ID',
      dataIndex: 'ID',
      width: 50,
      render: (text) => <a>{text}</a>,
    },
    {
      title: '用户名',
      dataIndex: 'userName',
      width: 100,
      render: (text) => <a>{text}</a>,
    },
    {
      title: '钱包地址',
      dataIndex: 'address',
      width: 200,
      render: (text) => (text ? <a>{text}</a> : '暂无'),
    },
    {
      title: '余额',
      dataIndex: 'amount',
      width: 150,
      render: (text) => <a>{text}</a>,
    },
    {
      title: 'LP 余额',
      dataIndex: 'amountLp',
      width: 150,
      render: (text) => <a>{text}</a>,
    },
    {
      title: 'USDT 余额',
      dataIndex: 'amountUsdt',
      width: 150,
      render: (text) => <a>{text}</a>,
    },
    {
      title: '注册时间',
      dataIndex: 'CreatedAt',
      width: 250,
      render: (text) => <a>{text}</a>,
    },
    {
      title: '操作',
      dataIndex: '',
      fixed: 'right',
      width: 150,
      render: () => <Button type="primary">编辑</Button>,
    },
  ];
  const onFinish = (values: { id: string; username: string; address: string }) => {
    getUserInfoModel(values);
  };

  const onFinishFailed = (errorInfo: any) => {
    console.log('Failed:', errorInfo);
  };
  const syncTable = () => {
    window.location.reload();
  };
  return (
    <PageContainer
      content={intl.formatMessage({
        id: 'pages.admin.subPage.userlist',
        defaultMessage: 'userlist',
      })}
    >
      <Card size="small">
        <Form
          ref={formRef}
          name="basic"
          initialValues={{ remember: true }}
          autoComplete="off"
          layout="inline"
          onFinish={onFinish}
          onFinishFailed={onFinishFailed}
        >
          <Form.Item name={'id'} label="ID">
            <Input />
          </Form.Item>
          <Form.Item name={'username'} label="用户名">
            <Input />
          </Form.Item>
          <Form.Item name={'address'} label="钱包地址">
            <Input />
          </Form.Item>
          <Form.Item>
            <Button type="primary" htmlType="submit">
              查找
            </Button>
          </Form.Item>
          <Form.Item>
            <Button type="primary" onClick={syncTable} icon={<SyncOutlined />} />
          </Form.Item>
        </Form>
      </Card>
      <br />
      <Card size="small">
        <Table
          rowKey={(r: DataType) => r.ID}
          scroll={{ x: 240 }}
          columns={columns}
          dataSource={userList}
        />
      </Card>
    </PageContainer>
  );
};

interface DataType {
  CreatedAt?: string | null;
  DeletedAt?: string | null;
  UpdatedAt?: string | null;
  ID?: number;
  address?: string | null;
  amount?: number;
  amountLp?: number;
  amountUsdt?: number;
  chain?: string | null;
  isBroker?: number;
  isMember?: number;
  level?: number;
  password?: string | null;
  payU?: number;
  recommendUserId?: number;
  salt?: string | null;
  status?: number;
  userName?: string | null;
}

export default memo(UserList);
