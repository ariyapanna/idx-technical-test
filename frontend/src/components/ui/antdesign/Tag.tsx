import React from 'react';
import { Tag as AntTag } from 'antd';

interface TagProps {
  color?: string;
  children?: React.ReactNode;
  style?: React.CSSProperties;
}

const Tag: React.FC<TagProps> = (props) => {
  return <AntTag {...props} />;
};

export default Tag;
