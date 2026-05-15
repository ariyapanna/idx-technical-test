import React from 'react';
import { Select as AntSelect } from 'antd';

interface SelectProps {
  placeholder?: string;
  value?: any;
  defaultValue?: any;
  onChange?: (value: any) => void;
  style?: React.CSSProperties;
  className?: string;
  children?: React.ReactNode;
  options?: { label: string; value: any }[];
  allowClear?: boolean;
}

const Select: React.FC<SelectProps> & { Option: typeof AntSelect.Option } = ({ children, ...props }) => {
  return <AntSelect {...props}>{children}</AntSelect>;
};

Select.Option = AntSelect.Option;

export default Select;
