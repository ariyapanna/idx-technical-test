import React from 'react';
import { Input as AntInput } from 'antd';

interface BaseInputProps {
  placeholder?: string;
  value?: string;
  onChange?: (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => void;
  style?: React.CSSProperties;
  className?: string;
  allowClear?: boolean;
}

interface InputProps extends BaseInputProps {}

interface TextAreaProps extends BaseInputProps {
  rows?: number;
}

const Input: React.FC<InputProps> & {
  TextArea: React.FC<TextAreaProps>;
} = (props) => {
  return <AntInput {...props} />;
};

Input.TextArea = (props) => <AntInput.TextArea {...props} />;

export default Input;
