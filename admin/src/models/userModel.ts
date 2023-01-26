import { useState, useEffect } from 'react';
import { getUserList, getUserInfo } from '@/services/user';
export default () => {
  const [userList, setUserList] = useState<any>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    getUserList().then((res) => {
      setUserList(res.data);
      setLoading(false);
    });
  }, []);
  const getUserInfoModel = (data: { id: string; username: string; address: string }) => {
    getUserInfo(data).then((res) => {
      let arr = [];
      arr.push(res.data);
      setUserList(arr);
      setLoading(false);
    });
  };

  return {
    getUserInfoModel,
    userList,
    loading,
  };
};
