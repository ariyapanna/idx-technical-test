import React from 'react';
import { Popconfirm as AntPopconfirm } from 'antd';

interface PopconfirmProps {
  title: React.ReactNode;
  onConfirm?: () => void;
  okText?: string;
  cancelText?: string;
  children?: React.ReactNode;
}

const Popconfirm: React.FC<PopconfirmProps> = (props) => {
  return <AntPopconfirm {...props} />;
};

export default Popconfirm;
