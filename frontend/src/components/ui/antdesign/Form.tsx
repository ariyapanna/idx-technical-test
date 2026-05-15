import React from 'react';
import { Form as AntForm } from 'antd';

interface FormProps {
  form?: any;
  layout?: 'horizontal' | 'vertical' | 'inline';
  initialValues?: any;
  onFinish?: (values: any) => void;
  children?: React.ReactNode;
  style?: React.CSSProperties;
}

interface FormItemProps {
  name?: string | number | (string | number)[];
  label?: React.ReactNode;
  rules?: any[];
  children?: React.ReactNode;
  style?: React.CSSProperties;
}

const Form: React.FC<FormProps> & { Item: React.FC<FormItemProps>; useForm: typeof AntForm.useForm } = ({ children, ...props }) => {
  return <AntForm {...props}>{children}</AntForm>;
};

Form.Item = ({ children, ...props }) => <AntForm.Item {...props}>{children}</AntForm.Item>;
Form.useForm = AntForm.useForm;

export default Form;
