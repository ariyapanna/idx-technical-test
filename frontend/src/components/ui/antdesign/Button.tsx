import React from 'react';
import { Button as AntButton } from 'antd';

interface ButtonProps {
  type?: 'primary' | 'default' | 'dashed' | 'link' | 'text';
  icon?: React.ReactNode;
  onClick?: () => void;
  danger?: boolean;
  size?: 'small' | 'middle' | 'large';
  children?: React.ReactNode;
  style?: React.CSSProperties;
  block?: boolean;
  loading?: boolean;
}

const Button: React.FC<ButtonProps> = ({ children, ...props }) => {
  return <AntButton {...props}>{children}</AntButton>;
};

export default Button;
