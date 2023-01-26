import { request } from '@umijs/max';

export async function getUserList() {
  return request<API.LoginResult>('/api/users/list', {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
    },
  });
}

export async function getUserInfo(options?: { id: string; username: string; address: string }) {
  return request<API.LoginResult>('/api/userinfo', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: {
      ...(options || {}),
    },
  });
}
