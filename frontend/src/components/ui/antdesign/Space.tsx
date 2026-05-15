import React from 'react';
import { Space as AntSpace } from 'antd';

interface SpaceProps {
  orientation?: 'horizontal' | 'vertical';
  size?: number | 'small' | 'middle' | 'large';
  children?: React.ReactNode;
  style?: React.CSSProperties;
  className?: string;
  onClick?: (e: React.MouseEvent) => void;
}

const Space: React.FC<SpaceProps> = ({ children, orientation = 'horizontal', ...props }) => {
  return <AntSpace direction={orientation} {...props}>{children}</AntSpace>;
};

export default Space;
