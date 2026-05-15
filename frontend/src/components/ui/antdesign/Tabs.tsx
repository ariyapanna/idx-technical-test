import React from 'react';
import { Tabs as AntTabs, type TabsProps as AntTabsProps } from 'antd';

export interface TabItem {
  key: string;
  label: string;
  icon?: React.ReactNode;
  children: React.ReactNode;
}

interface CustomTabsProps extends Omit<AntTabsProps, 'items'> {
  items: TabItem[];
}

const Tabs: React.FC<CustomTabsProps> = ({ items, style, ...rest }) => {
  const themedItems = items.map((item) => ({
    key: item.key,
    label: (
      <div style={{ display: 'flex', alignItems: 'center', gap: '8px' }}>
        {item.icon}
        <span>{item.label}</span>
      </div>
    ),
    children: item.children,
  }));

  return (
    <AntTabs
      items={themedItems}
      style={{
        ...style,
      }}
      className="custom-tabs"
      {...rest}
    />
  );
};

export default Tabs;
