import React from 'react';
import { Tooltip as AntTooltip } from 'antd';

interface TooltipProps {
  title: React.ReactNode;
  children?: React.ReactNode;
}

const Tooltip: React.FC<TooltipProps> = (props) => {
  return <AntTooltip {...props} />;
};

export default Tooltip;
