import React from 'react';
import { Typography as AntTypography } from 'antd';

interface TextProps {
  children?: React.ReactNode;
  delete?: boolean;
  strong?: boolean;
  type?: 'secondary' | 'success' | 'warning' | 'danger';
  style?: React.CSSProperties;
}

interface TitleProps {
  children?: React.ReactNode;
  level?: 1 | 2 | 3 | 4 | 5;
  style?: React.CSSProperties;
}

const Typography = {
  Text: ({ children, ...props }: TextProps) => <AntTypography.Text {...props}>{children}</AntTypography.Text>,
  Title: ({ children, ...props }: TitleProps) => <AntTypography.Title {...props}>{children}</AntTypography.Title>,
};

export default Typography;
export const { Text, Title } = Typography;
